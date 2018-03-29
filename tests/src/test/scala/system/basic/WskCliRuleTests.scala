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

import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner

import common.Wsk

import spray.json._
import spray.json.DefaultJsonProtocol._

@RunWith(classOf[JUnitRunner])
class WskCliRuleTests extends WskRuleTests {
  override val wsk: Wsk = new Wsk

  it should "preserve rule status when a rule is updated" in withAssetCleaner(wskprops) { (wp, assetHelper) =>
    val ruleName = withTimestamp("r1to1")
    val triggerName = withTimestamp("t1to1")
    val actionName = withTimestamp("a1 to 1")
    val triggerName2 = withTimestamp("t2to1")
    val active = Some("active".toJson)
    val inactive = Some("inactive".toJson)
    val statusPermutations =
      Seq((triggerName, active), (triggerName, inactive), (triggerName2, active), (triggerName2, inactive))

    ruleSetup(Seq((ruleName, triggerName, (actionName, actionName, defaultAction))), assetHelper)
    assetHelper.withCleaner(wsk.trigger, triggerName2) { (trigger, name) =>
      trigger.create(name)
    }

    statusPermutations.foreach {
      case (trigger, status) =>
        if (status == active) wsk.rule.enable(ruleName) else wsk.rule.disable(ruleName)
        // CLI stdout must strip out the preamble text (i.e. "ok: got rule XXXXX") to get at the JSON
        wsk.rule.create(ruleName, trigger, actionName, update = true)
        val getStdout = wsk.rule.get(ruleName).stdout
        getStdout.substring(getStdout.indexOf('{')).parseJson.asJsObject.fields.get("status") shouldBe status
    }
  }
}
