/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package system.basic

import akka.stream.scaladsl.{Sink, Source}
import common.rest.WskRestOperations
import common.{TestHelpers, TestUtils, Wsk, WskProps, WskTestHelpers}
import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner
import spray.json.DefaultJsonProtocol._
import spray.json._

import scala.concurrent._
import scala.concurrent.duration._
import scala.concurrent.Future

@RunWith(classOf[JUnitRunner])
class WskCliActivationTests extends TestHelpers with WskTestHelpers with HttpProxy {
  val wsk = new Wsk
  val wskRest = new WskRestOperations
  val defaultAction = Some(TestUtils.getTestActionFilename("hello.js"))

  behavior of "Wsk poll"

  implicit val wskprops: WskProps = WskProps()

  it should "change the since time as it polls" in withAssetCleaner(wskprops) {
    val name = "pollTest"
    (wp, assetHelper) =>
      assetHelper.withCleaner(wsk.action, name) { (action, _) =>
        action.create(name, Some(TestUtils.getTestActionFilename("hello.js")))
      }

      val args = Map("payload" -> "test".toJson)
      //This test spin up 2 parallel tasks
      // 1. Perform blocking invocations with 1 second interval
      // 2. Perform poll to pick up those activation results
      // For poll it inserts a proxy which intercepts the request sent to server
      // and then it asserts if the request sent have there `since` time getting changed or not
      withProxy { (proxyProps, requests) =>
        //It may taken some time for the activations to show up in poll result
        //based on view lag. So keep a bit longer time span for the poll
        val pollDuration = 10.seconds
        println(s"Running poll for $pollDuration")

        val consoleFuture = Future {
          //pass the `proxyProps` such that calls from poll cli command go via our proxy
          wsk.activation.console(pollDuration, actionName = Some(name))(proxyProps)
        }

        val runsFuture = Source(1 to 5)
          .map { _ =>
            val r = wskRest.action.invoke(name, args, blocking = true)
            Thread.sleep(2.second.toMillis)
            r
          }
          .runWith(Sink.seq)

        val f = for {
          rr <- runsFuture
          cr <- consoleFuture
        } yield (cr, rr)

        val (consoleResult, runResult) = Await.result(f, 1.minute)

        val activations = runResult.filter(_.statusCode.isSuccess()).map(_.respData.parseJson.asJsObject)
        val ids = activations.flatMap(_.fields.get("activationId").map(_.convertTo[String]))
        val idsInPoll = ids.filter(consoleResult.stdout.contains(_))

        //There should be more than 1 activationId in common between poll output
        //and actual invoked actions output
        //This is required to ensure that since time can change which would only
        //happen if more than one activation result is picked up in poll
        withClue(
          s"activations received ${activations.mkString("\n")}, console output $consoleResult. Expecting" +
            s"more than one matching activation between these 2") {
          idsInPoll.size should be > 1

          //Collect the 'since' value passed during poll requests
          val sinceTimes = requests.map(_._1.uri.query()).flatMap(_.get("since")).toSet

          withClue(s"value of 'since' $sinceTimes should have changed") {
            sinceTimes.size should be > 1
          }
        }
      }
  }

}
