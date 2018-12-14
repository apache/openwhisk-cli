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
import java.net.ServerSocket

import akka.http.scaladsl.{Http, HttpsConnectionContext}
import akka.http.scaladsl.model.{HttpRequest, HttpResponse, Uri}
import akka.http.scaladsl.model.Uri.Authority
import akka.http.scaladsl.server.Route
import akka.stream.ActorMaterializer
import akka.stream.scaladsl.{Sink, Source}
import com.typesafe.sslconfig.akka.AkkaSSLConfig
import common.{WskActorSystem, WskProps}
import common.rest.{AcceptAllHostNameVerifier, SSL}
import javax.net.ssl.HostnameVerifier
import org.scalatest.Suite
import org.scalatest.concurrent.ScalaFutures

import scala.collection.mutable.ListBuffer
import scala.concurrent.duration._

/**
 * A minimal reverse proxy implementation for test purpose which intercepts the
 * request and responses and then make them available to test for validation.
 *
 * It also allows connecting to https endpoint while still expose a http endpoint
 * to local client
 */
trait HttpProxy extends WskActorSystem with ScalaFutures {
  self: Suite =>

  implicit val materializer: ActorMaterializer = ActorMaterializer()
  implicit val testConfig: PatienceConfig = PatienceConfig(1.minute)

  def withProxy(check: (WskProps, ListBuffer[(HttpRequest, HttpResponse)]) => Unit)(implicit wp: WskProps): Unit = {
    val uri = getTargetUri(wp)
    val requests = new ListBuffer[(HttpRequest, HttpResponse)]
    val port = freePort()
    val proxy = Route { context =>
      val request = context.request
      val handler = Source
        .single(proxyRequest(request, uri))
        .via(makeHttpFlow(uri))
        .runWith(Sink.head)
        .map { response =>
          requests += ((request, response))
          response
        }
        .flatMap(context.complete(_))
      handler
    }

    val binding = Http(actorSystem).bindAndHandle(handler = proxy, interface = "localhost", port = port)
    binding.map { b =>
      val proxyProps = wp.copy(apihost = s"http://localhost:$port")
      check(proxyProps, requests)
      b.unbind()
    }.futureValue
  }

  private def getTargetUri(wp: WskProps) = {
    // startsWith(http) includes https
    if (wp.apihost.startsWith("http")) {
      Uri(wp.apihost)
    } else {
      Uri().withScheme("https").withHost(wp.apihost)
    }
  }

  private def makeHttpFlow(uri: Uri) = {
    if (uri.scheme == "https") {
      //Use ssl config which does not validate anything
      Http(actorSystem).outgoingConnectionHttps(
        uri.authority.host.address(),
        uri.effectivePort,
        connectionContext = httpsConnectionContext())
    } else {
      Http(actorSystem).outgoingConnection(uri.authority.host.address(), uri.effectivePort)
    }
  }

  private def httpsConnectionContext() = {
    val sslConfig = AkkaSSLConfig().mapSettings { s =>
      s.withHostnameVerifierClass(classOf[AcceptAllHostNameVerifier].asInstanceOf[Class[HostnameVerifier]])
    }
    //SSL.httpsConnectionContext initializes config which is not there in cli test
    //So inline the flow as we do not need client auth for this case
    new HttpsConnectionContext(SSL.nonValidatingContext(false), Some(sslConfig))
  }

  private def proxyRequest(req: HttpRequest, uri: Uri): HttpRequest = {
    //https://github.com/akka/akka-http/issues/64
    req
      .copy(headers = req.headers.filterNot(h => h.is("timeout-access")))
      .copy(uri = req.uri.copy(scheme = "", authority = Authority.Empty)) //Strip the authority as it refers to proxy
  }

  private def freePort(): Int = {
    val socket = new ServerSocket(0)
    try socket.getLocalPort
    finally if (socket != null) socket.close()
  }
}
