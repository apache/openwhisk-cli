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

import java.io.File

import org.junit.runner.RunWith
import org.scalatest.junit.JUnitRunner

import common.Wsk

import whisk.core.WhiskConfig

/**
 * Tests sequence execution
 */

@RunWith(classOf[JUnitRunner])
class WskCliSequenceTests extends WskSequenceTests {
  override val wsk = new Wsk
  val owHome = System.getenv("OPENWHISK_HOME") match {
    case home: String if !home.isEmpty => home
    case _                             => "../../../../../../incubator-openwhisk"
  }
  val propertiesFile = new File(s"$owHome/whisk.properties")
  override val whiskConfig =
    new WhiskConfig(Map(WhiskConfig.actionSequenceMaxLimit -> null), propertiesFile = propertiesFile)
}
