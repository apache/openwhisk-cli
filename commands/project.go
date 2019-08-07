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

package commands

import (
	"github.com/apache/openwhisk-cli/wski18n"
	"github.com/apache/openwhisk-wskdeploy/cmd"
	"github.com/apache/openwhisk-wskdeploy/utils"
	wskdeploy_wski18n "github.com/apache/openwhisk-wskdeploy/wski18n"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "The OpenWhisk Project Management Tool",
}

var projectDeployCmd = &cobra.Command{
	Use:           "deploy",
	Short:         wski18n.T(wskdeploy_wski18n.ID_CMD_DESC_SHORT_ROOT),
	Long:          wski18n.T(CMD_DESC_LONG_DEPLOY),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cobraCMD *cobra.Command, args []string) error {
		return cmd.Deploy(cobraCMD)
	},
}

var projectUnDeployCmd = &cobra.Command{
	Use:           "undeploy",
	Short:         wski18n.T(wskdeploy_wski18n.ID_CMD_DESC_SHORT_UNDEPLOY),
	Long:          wski18n.T(CMD_DESC_LONG_UNDEPLOY),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cobraCMD *cobra.Command, args []string) error {
		return cmd.Undeploy(cobraCMD)
	},
}

var projectSyncCmd = &cobra.Command{
	Use:           "sync",
	Short:         wski18n.T(wskdeploy_wski18n.ID_CMD_DESC_SHORT_SYNC),
	Long:          wski18n.T(CMD_DESC_LONG_SYNC),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cobraCMD *cobra.Command, args []string) error {
		utils.Flags.Sync = true
		return cmd.Deploy(cobraCMD)
	},
}

var projectExportCmd = &cobra.Command{
	Use:           "export",
	Short:         wski18n.T(wskdeploy_wski18n.ID_CMD_DESC_SHORT_EXPORT),
	Long:          wski18n.T(CMD_DESC_LONG_EXPORT),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cobraCMD *cobra.Command, args []string) error {
		return cmd.ExportCmdImp(cobraCMD, args)
	},
}

func init() {
	projectCmd.PersistentFlags().StringVar(&utils.Flags.CfgFile, cmd.FLAG_CONFIG, "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_CONFIG))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.ManifestPath, cmd.FLAG_MANIFEST, "", "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_MANIFEST))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.ProjectPath, cmd.FLAG_PROJECT, "", ".", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_PROJECT))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.DeploymentPath, cmd.FLAG_DEPLOYMENT, "", "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_DEPLOYMENT))
	projectCmd.PersistentFlags().BoolVarP(&utils.Flags.Strict, cmd.FLAG_STRICT, "", false, wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_STRICT))
	projectCmd.PersistentFlags().BoolVarP(&utils.Flags.Preview, cmd.FLAG_PREVIEW, "", false, wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_PREVIEW))
	projectCmd.PersistentFlags().StringSliceVarP(&utils.Flags.Param, cmd.FLAG_PARAM, "", []string{}, wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_PARAM))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.ParamFile, cmd.FLAG_PARAMFILE, "", "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_PARAM_FILE))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.ApiHost, cmd.FLAG_API_HOST, "", "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_API_HOST))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.Namespace, cmd.FLAG_NAMESPACE, cmd.FLAG_NAMESPACE_SHORT, "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_NAMESPACE))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.Auth, cmd.FLAG_AUTH, cmd.FLAG_AUTH_SHORT, "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_AUTH_KEY))
	projectCmd.PersistentFlags().BoolVarP(&utils.Flags.Verbose, cmd.FLAG_VERBOSE, cmd.FLAG_VERBOSE_SHORT, false, wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_VERBOSE))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.Key, cmd.FLAG_KEY, cmd.FLAG_KEY_SHORT, "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_KEY_FILE))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.Cert, cmd.FLAG_CERT, cmd.FLAG_CERT_SHORT, "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_CERT_FILE))
	projectCmd.PersistentFlags().StringVarP(&utils.Flags.ProjectName, cmd.FLAG_PROJECTNAME, "", "", wski18n.T(wskdeploy_wski18n.ID_CMD_FLAG_PROJECTNAME))

	projectCmd.AddCommand(projectDeployCmd)
	projectCmd.AddCommand(projectUnDeployCmd)
	projectCmd.AddCommand(projectSyncCmd)
	projectCmd.AddCommand(projectExportCmd)
}
