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

import common.TestUtils
import common.TestUtils.NOT_ALLOWED
import common.Wsk

@RunWith(classOf[JUnitRunner])
class WskCliActionTests extends WskActionTests {
  override val wsk = new Wsk

  it should "not be able to use --kind and --docker at the same time when running action create or update" in {
    val file = TestUtils.getTestActionFilename(s"echo.js")
    Seq(false, true).foreach { updateValue =>
      val out = wsk.action.create(
        name = "kindAndDockerAction",
        artifact = Some(file),
        expectedExitCode = NOT_ALLOWED,
        kind = Some("nodejs:6"),
        docker = Some("mydockerimagename"),
        update = updateValue)
      out.stderr should include("Cannot specify both --kind and --docker at the same time")
    }
  }
}
