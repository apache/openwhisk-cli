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

import com.jayway.restassured.RestAssured

import common.{TestUtils, Wsk}
import common.TestUtils._

import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner

import scala.concurrent.duration.DurationInt
import scala.util.parsing.json.JSON

/**
  * Tests for basic CLI usage. Some of these tests require a deployed backend.
  */
@RunWith(classOf[JUnitRunner])
class ApiGwCliTests extends ApiGwTests {
  override lazy val wsk: common.Wsk = new Wsk
  override lazy val createCode = SUCCESS_EXIT
  behavior of "Cli Wsk api creation with path parameters no swagger"

  it should "fail to create an API if the relative path contains invalid path parameters" in withAssetCleaner(wskprops) {(wp, assetHelper) =>
    val actionName = "APIGWTEST_BAD_RELATIVE_PATH_ACTION"
    val basePath = "/mybase/path"
    val file = TestUtils.getTestActionFilename(s"echo-web-http.js")
    assetHelper.withCleaner(wsk.action, actionName, confirmDelete = true) {
      (action, _) =>
        action.create(actionName, Some(file), web = Some("true"))
    }
    var relPath = "/bad/{path/value"
    var rr = apiCreate(basepath = Some(basePath),
      relpath = Some(relPath),
      operation = Some("GET"),
      action = Some(actionName),
      expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"Relative path '${relPath}' does not include valid path parameters. Each parameter must be enclosed in curly braces '{}'.")

    relPath = "/bad/path}/value"
    rr = apiCreate(basepath = Some(basePath),
      relpath = Some(relPath),
      operation = Some("GET"),
      action = Some(actionName),
      expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"Relative path '${relPath}' does not include valid path parameters. Each parameter must be enclosed in curly braces '{}'.")

    relPath = "/bad/{path/va}lue"
    rr = apiCreate(basepath = Some(basePath),
      relpath = Some(relPath),
      operation = Some("GET"),
      action = Some(actionName),
      expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"Relative path '${relPath}' does not include valid path parameters. Each parameter must be enclosed in curly braces '{}'.")

    relPath = "/ba}d/{path/value"
    rr = apiCreate(basepath = Some(basePath),
      relpath = Some(relPath),
      operation = Some("GET"),
      action = Some(actionName),
      expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"Relative path '${relPath}' does not include valid path parameters. Each parameter must be enclosed in curly braces '{}'.")

    relPath = "/ba}d/{p{at}h/value"
    rr = apiCreate(basepath = Some(basePath),
      relpath = Some(relPath),
      operation = Some("GET"),
      action = Some(actionName),
      expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"Relative path '${relPath}' does not include valid path parameters. Each parameter must be enclosed in curly braces '{}'.")
  }

  it should "fail to create an API if the base path contains path parameters" in withAssetCleaner(wskprops) {(wp, assetHelper) =>
    val actionName = "APIGWTEST_BAD_BASE_PATH_ACTION"
    val basePath = "/mybase/{path}"
    val file = TestUtils.getTestActionFilename(s"echo-web-http.js")
    assetHelper.withCleaner(wsk.action, actionName, confirmDelete = true) {
      (action, _) =>
        action.create(actionName, Some(file), web = Some("true"))
    }
    val relPath = "/bad/{path}/value"
    val rr = apiCreate(basepath = Some(basePath),
      relpath = Some(relPath),
      operation = Some("GET"),
      action = Some(actionName),
      expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"The base path '${basePath}' cannot have parameters. Only the relative path supports path parameters.")
  }

  it should "fail to create an Api if path parameters are specified but http response type is not given" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val actionName = "CLI_APIGWTEST_PATH_param_fail1_action"
    val file = TestUtils.getTestActionFilename(s"echo-web-http.js")
    assetHelper.withCleaner(wsk.action, actionName, confirmDelete = true) {
      (action, _) =>
        action.create(actionName, Some(file), web = Some("true"))
    }
    val relPath = "/bad/{path}/value"
    val rr = apiCreate(basepath = Some("/mybase"),
                       relpath = Some(relPath),
                       operation = Some("GET"),
                       action = Some(actionName),
                       expectedExitCode = ANY_ERROR_EXIT)
    rr.stderr should include(
      s"A response type of 'http' is required when using path parameters.")
  }

  it should "create api with path parameters for the verb" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val testName = "CLI_APIGWTEST_PATH_PARAMS1"
    val testBasePath = s"/${testName}_bp"
    val testUrlName1 = "scooby"
    val testUrlName2 = "doo"
    val testRelPath = "/path/{with}/some/{path}/params"
    val testRelPathGet = s"/path/${testUrlName1}/some/${testUrlName2}/params"
    val testUrlOp = "get"
    val testApiName = testName + " API Name"
    val actionName = testName + "_action"
    var exception: Throwable = null
    val reqPath = "\\$\\(request.path\\)"

    // Create the action for the API.  It must be a "web-action" action.
    val file = TestUtils.getTestActionFilename(s"echo-web-http.js")
    assetHelper.withCleaner(wsk.action, actionName, confirmDelete = true) {
      (action, _) =>
        action.create(actionName, Some(file), web = Some("true"))
    }
    try {
      var rr = apiCreate(
        basepath = Some(testBasePath),
        relpath = Some(testRelPath),
        operation = Some(testUrlOp),
        action = Some(actionName),
        apiname = Some(testApiName),
        responsetype = Some("http")
      )
      verifyApiCreated(rr)
      val swaggerApiUrl = getSwaggerUrl(rr).replace("{with}", testUrlName1).replace("{path}", testUrlName2)

      //Validate the api created contained parameters and they were correct
      rr = apiGet(basepathOrApiName = Some(testApiName))
      rr.stdout should include(testBasePath)
      rr.stdout should include(s"${actionName}")
      rr.stdout should include regex (""""cors":\s*\{\s*\n\s*"enabled":\s*true""")
      rr.stdout should include regex (s"""target-url.*${actionName}.http${reqPath}""")
      val params = getParametersFromJson(rr, testRelPath)
      params.size should be(2)
      validateParameter(params(0), "with", "path", true, "string", "Default description for 'with'")
      validateParameter(params(1), "path", "path", true, "string", "Default description for 'path'")

      //Lets call the swagger url so we can make sure the response is valid and contains our path in the ow path
      val apiToInvoke = s"$swaggerApiUrl"
      println(s"Invoking: '${apiToInvoke}'")
      val response = whisk.utils.retry({
        val response = RestAssured.given().config(getSslConfig()).get(s"$apiToInvoke")
        response.statusCode should be(200)
        response
      }, 6, Some(2.second))
      val jsonReponse = JSON.parseFull(response.asString()).asInstanceOf[Option[Map[String, String]]].get
      jsonReponse.get("__ow_path").get should not be ("")
      jsonReponse.get("__ow_path").get should include (testRelPathGet)
    } catch {
      case unknown: Throwable => exception = unknown
    } finally {
      apiDelete(basepathOrApiName = testBasePath)
    }
    assert(exception == null)
  }

  it should "create api with path parameters and pass them into the action bound to the api" in withAssetCleaner(
    wskprops) { (wp, assetHelper) =>
    val testName = "CLI_APIGWTEST_PATH_PARAMS2"
    val testBasePath = "/" + testName + "_bp"
    val testRelPath = "/path/{with}/some/{double}/{extra}/{extra}/{path}"
    val testUrlName1 = "scooby"
    val testUrlName2 = "doo"
    val testUrlName3 = "shaggy"
    val testUrlName4 = "velma"
    val testRelPathGet = s"/path/$testUrlName1/some/$testUrlName3/$testUrlName4/$testUrlName4/$testUrlName2"
    val testUrlOp = "get"
    val testApiName = testName + " API Name"
    val actionName = testName + "_action"
    var exception: Throwable = null
    val reqPath = "\\$\\(request.path\\)"
    // Create the action for the API.  It must be a "web-action" action.
    val file = TestUtils.getTestActionFilename(s"echo-web-http.js")
    assetHelper.withCleaner(wsk.action, actionName, confirmDelete = true) {
      (action, _) =>
        action.create(actionName, Some(file), web = Some("true"))
    }
    try {
      var rr = apiCreate(
        basepath = Some(testBasePath),
        relpath = Some(testRelPath),
        operation = Some(testUrlOp),
        action = Some(actionName),
        apiname = Some(testApiName),
        responsetype = Some("http")
      )
      verifyApiCreated(rr)
      val swaggerApiUrl = getSwaggerUrl(rr).replace("{with}", testUrlName1).replace("{path}", testUrlName2)
          .replace("{double}", testUrlName3).replace("{extra}", testUrlName4)

      //Validate the api created contained parameters and they were correct
      rr = apiGet(basepathOrApiName = Some(testApiName))
      rr.stdout should include(testBasePath)
      rr.stdout should include(s"${actionName}")
      rr.stdout should include regex (""""cors":\s*\{\s*\n\s*"enabled":\s*true""")
      rr.stdout should include regex (s"""target-url.*${actionName}.http${reqPath}""")
      val params = getParametersFromJson(rr, testRelPath)

      // should have 4, not 5 parameter definitions (i.e. don't define "extra" twice
      params.size should be(4)
      validateParameter(params(0), "with", "path", true, "string", "Default description for 'with'")
      validateParameter(params(1), "double", "path", true, "string", "Default description for 'double'")
      validateParameter(params(2), "extra", "path", true, "string", "Default description for 'extra'")
      validateParameter(params(3), "path", "path", true, "string", "Default description for 'path'")

      //Lets call the swagger url so we can make sure the response is valid and contains our path in the ow path
      val apiToInvoke = s"$swaggerApiUrl"
      println(s"Invoking: '${apiToInvoke}'")
      val response = whisk.utils.retry({
        val response = RestAssured.given().config(getSslConfig()).get(s"$apiToInvoke")
        response.statusCode should be(200)
        response
      }, 6, Some(2.second))
      val jsonReponse = JSON.parseFull(response.asString()).asInstanceOf[Option[Map[String, String]]].get
      jsonReponse.get("__ow_path").get should not be ("")
      jsonReponse.get("__ow_path").get should include (testRelPathGet)
    } catch {
      case unknown: Throwable => exception = unknown
    } finally {
      apiDelete(basepathOrApiName = testBasePath)
    }
    assert(exception == null)
  }
}
