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

package whisk.core.cli.test

import scala.language.postfixOps
import scala.concurrent.duration.DurationInt
import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner
import common.TestHelpers
import common.TestUtils
import common.TestUtils._
import common.Wsk
import common.WskProps
import common.WskTestHelpers
import spray.json.DefaultJsonProtocol._
import spray.json._
import whisk.core.entity._
import whisk.core.cli.test.TestJsonArgs._
import whisk.http.Messages

/**
  * Tests for basic CLI usage. Some of these tests require a deployed backend.
  */
@RunWith(classOf[JUnitRunner])
class WskCliBasicUsageTests extends TestHelpers with WskTestHelpers {

  implicit val wskprops = WskProps()
  val wsk = new Wsk
  val defaultAction = Some(TestUtils.getTestActionFilename("hello.js"))
  val usrAgentHeaderRegEx = """\bUser-Agent\b": \[\s+"OpenWhisk\-CLI/1.\d+.*"""
  // certain environments may return router IP address instead of api_host string causing a failure
  // Set apiHostCheck to false to avoid apihost check
  val apiHostCheck = true

  behavior of "Wsk CLI usage"

  it should "show help and usage info" in {
    val stdout = wsk.cli(Seq()).stdout
    stdout should include regex ("""(?i)Usage:""")
    stdout should include regex ("""(?i)Flags""")
    stdout should include regex ("""(?i)Available commands""")
    stdout should include regex ("""(?i)--help""")
  }

  it should "show help and usage info using the default language" in {
    val env = Map("LANG" -> "de_DE")
    // Call will fail with exit code 2 if language not supported
    wsk.cli(Seq("-h"), env = env)
  }

  it should "show cli build version" in {
    val stdout = wsk.cli(Seq("property", "get", "--cliversion")).stdout
    stdout should include regex ("""(?i)whisk CLI version\s+201.*""")
  }

  it should "show api version" in {
    val stdout = wsk.cli(Seq("property", "get", "--apiversion")).stdout
    stdout should include regex ("""(?i)whisk API version\s+v1""")
  }

  it should "reject bad command" in {
    val result = wsk.cli(Seq("bogus"), expectedExitCode = ERROR_EXIT)
    result.stderr should include regex ("""(?i)Run 'wsk --help' for usage""")
  }

  it should "allow a 3 part Fully Qualified Name (FQN) without a leading '/'" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val guestNamespace = wsk.namespace.whois()
    val packageName = withTimestamp("packageName3ptFQN")
    val actionName = withTimestamp("actionName3ptFQN")
    val triggerName = withTimestamp("triggerName3ptFQN")
    val ruleName = withTimestamp("ruleName3ptFQN")
    val fullQualifiedName = s"${guestNamespace}/${packageName}/${actionName}"
    // Used for action and rule creation below
    assetHelper.withCleaner(wsk.pkg, packageName) { (pkg, _) =>
      pkg.create(packageName)
    }
    assetHelper.withCleaner(wsk.trigger, triggerName) { (trigger, _) =>
      trigger.create(triggerName)
    }
    // Test action and rule creation where action name is 3 part FQN w/out leading slash
    assetHelper.withCleaner(wsk.action, fullQualifiedName) { (action, _) =>
      action.create(fullQualifiedName, defaultAction)
    }
    assetHelper.withCleaner(wsk.rule, ruleName) { (rule, _) =>
      rule.create(ruleName, trigger = triggerName, action = fullQualifiedName)
    }

    //wsk.action.invoke(fullQualifiedName).stdout should include(
    //  s"ok: invoked /$fullQualifiedName")
    //wsk.action.get(fullQualifiedName).stdout should include(
     // s"ok: got action ${packageName}/${actionName}")
  }

  it should "include CLI user agent headers with outbound requests" in {
    val stdout = wsk
      .cli(
        Seq("action", "list", "--auth", wskprops.authKey) ++ wskprops.overrides,
        verbose = true)
      .stdout
    stdout should include regex (usrAgentHeaderRegEx)
  }

  behavior of "Wsk actions"

  it should "reject creating entities with invalid names" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val names = Seq(("", ERROR_EXIT),
                    (" ", BAD_REQUEST),
                    ("hi+there", BAD_REQUEST),
                    ("$hola", BAD_REQUEST),
                    ("dora?", BAD_REQUEST),
                    ("|dora|dora?", BAD_REQUEST))

    names foreach {
      case (name, ec) =>
        assetHelper.withCleaner(wsk.action, name, confirmDelete = false) {
          (action, _) =>
            action.create(name, defaultAction, expectedExitCode = ec)
        }
    }
  }

  it should "reject create with missing file" in {
    val name = "notfound"
    wsk.action
      .create("missingFile", Some(name), expectedExitCode = MISUSE_EXIT)
      .stderr should include(
      s"File '$name' is not a valid file or it does not exist")
  }

  it should "reject action update when specified file is missing" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    // Create dummy action to update
    val name = "updateMissingFile"
    val file = Some(TestUtils.getTestActionFilename("hello.js"))
    assetHelper.withCleaner(wsk.action, name) { (action, name) =>
      action.create(name, file)
    }
    // Update it with a missing file
    wsk.action.create(name,
                      Some("notfound"),
                      update = true,
                      expectedExitCode = MISUSE_EXIT)
  }

  it should "reject action update for sequence with no components" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val name = "updateMissingComponents"
    val file = Some(TestUtils.getTestActionFilename("hello.js"))
    assetHelper.withCleaner(wsk.action, name) { (action, name) =>
      action.create(name, file)
    }
    wsk.action.create(name,
                      None,
                      update = true,
                      kind = Some("sequence"),
                      expectedExitCode = MISUSE_EXIT)
  }

  it should "create, and get an action to verify parameter and annotation parsing" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val name = "actionAnnotations"
    val file = Some(TestUtils.getTestActionFilename("hello.js"))

    assetHelper.withCleaner(wsk.action, name) { (action, _) =>
      action.create(name,
                    file,
                    annotations = getValidJSONTestArgInput,
                    parameters = getValidJSONTestArgInput)
    }

    val stdout = wsk.action.get(name).stdout
    assert(stdout.startsWith(s"ok: got action $name\n"))

    val receivedParams = wsk
      .parseJsonString(stdout)
      .fields("parameters")
      .convertTo[JsArray]
      .elements
    val receivedAnnots = wsk
      .parseJsonString(stdout)
      .fields("annotations")
      .convertTo[JsArray]
      .elements
    val escapedJSONArr = getValidJSONTestArgOutput.convertTo[JsArray].elements

    for (expectedItem <- escapedJSONArr) {
      receivedParams should contain(expectedItem)
      receivedAnnots should contain(expectedItem)
    }
  }

  it should "create, and get an action to verify file parameter and annotation parsing" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val name = "actionAnnotAndParamParsing"
    val file = Some(TestUtils.getTestActionFilename("hello.js"))
    val argInput = Some(TestUtils.getTestActionFilename("validInput1.json"))

    assetHelper.withCleaner(wsk.action, name) { (action, _) =>
      action.create(name,
                    file,
                    annotationFile = argInput,
                    parameterFile = argInput)
    }

    val stdout = wsk.action.get(name).stdout
    assert(stdout.startsWith(s"ok: got action $name\n"))

    val receivedParams = wsk
      .parseJsonString(stdout)
      .fields("parameters")
      .convertTo[JsArray]
      .elements
    val receivedAnnots = wsk
      .parseJsonString(stdout)
      .fields("annotations")
      .convertTo[JsArray]
      .elements
    val escapedJSONArr = getJSONFileOutput.convertTo[JsArray].elements

    for (expectedItem <- escapedJSONArr) {
      receivedParams should contain(expectedItem)
      receivedAnnots should contain(expectedItem)
    }
  }

  it should "create an action with the proper parameter and annotation escapes" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val name = "actionEscapes"
    val file = Some(TestUtils.getTestActionFilename("hello.js"))

    assetHelper.withCleaner(wsk.action, name) { (action, _) =>
      action.create(name,
                    file,
                    parameters = getEscapedJSONTestArgInput,
                    annotations = getEscapedJSONTestArgInput)
    }

    val stdout = wsk.action.get(name).stdout
    assert(stdout.startsWith(s"ok: got action $name\n"))

    val receivedParams = wsk
      .parseJsonString(stdout)
      .fields("parameters")
      .convertTo[JsArray]
      .elements
    val receivedAnnots = wsk
      .parseJsonString(stdout)
      .fields("annotations")
      .convertTo[JsArray]
      .elements
    val escapedJSONArr = getEscapedJSONTestArgOutput.convertTo[JsArray].elements

    for (expectedItem <- escapedJSONArr) {
      receivedParams should contain(expectedItem)
      receivedAnnots should contain(expectedItem)
    }
  }

  it should "invoke an action that exits during initialization and get appropriate error" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val name = "abort init"
    assetHelper.withCleaner(wsk.action, name) { (action, _) =>
      action.create(name,
                    Some(TestUtils.getTestActionFilename("initexit.js")))
    }

    withActivation(wsk.activation, wsk.action.invoke(name)) { activation =>
      val response = activation.response
      response.result.get
        .fields("error") shouldBe Messages.abnormalInitialization.toJson
      response.status shouldBe ActivationResponse.messageForCode(
        ActivationResponse.ContainerError)
    }
  }

  it should "invoke an action that hangs during initialization and get appropriate error" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val name = "hang init"
    assetHelper.withCleaner(wsk.action, name) { (action, _) =>
      action.create(name,
                    Some(TestUtils.getTestActionFilename("initforever.js")),
                    timeout = Some(3 seconds))
    }

    withActivation(wsk.activation, wsk.action.invoke(name)) { activation =>
      val response = activation.response
      response.result.get.fields("error") shouldBe Messages
        .timedoutActivation(3 seconds, true).toJson
      response.status shouldBe ActivationResponse.messageForCode(
        ActivationResponse.ApplicationError)
    }
  }

}
