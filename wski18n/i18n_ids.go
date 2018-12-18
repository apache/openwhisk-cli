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

package wski18n

// DO NOT TRANSLATE
// i18n Identifiers
const (
	// Cobra command descriptions
	ID_CMD_DESC_LONG_DEPLOY   = "msg_cmd_desc_long_deploy"
	ID_CMD_DESC_LONG_SYNC     = "msg_cmd_desc_long_sync"
	ID_CMD_DESC_LONG_UNDEPLOY = "msg_cmd_desc_long_undeploy"
	ID_CMD_DESC_LONG_EXPORT   = "msg_cmd_desc_long_export"
)

// DO NOT TRANSLATE
// Used to unit test that translations exist with these IDs
var I18N_ID_SET = [](string){
	ID_CMD_DESC_LONG_DEPLOY,
	ID_CMD_DESC_LONG_SYNC,
	ID_CMD_DESC_LONG_UNDEPLOY,
	ID_CMD_DESC_LONG_EXPORT,
}
