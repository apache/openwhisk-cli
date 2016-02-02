package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// sdkCmd represents the sdk command
var sdkCmd = &cobra.Command{
	Use:   "sdk",
	Short: "work with the sdk",
}

var sdkInstallCmd = &cobra.Command{
	Use:   "install <component string:{docker,swift,iOS}>",
	Short: "install artifacts",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err := errors.New("Invalid component argument")
			fmt.Println(err)
			return
		}
		component := args[0]
		switch component {
		case "docker":
			err = dockerInstall()
		case "swift":
			err = swiftInstall()
		case "iOS":
			err = iOSInstall()
		default:
			err = errors.New("Invalid component argument")
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func dockerInstall() error {
	var err error
	tarFile := "blackbox-0.1.0.tar.gz"
	blackboxDir := "dockerSkeleton"
	if _, err = os.Stat(tarFile); err == nil {
		err = fmt.Errorf("The path %s already exists.  Please deleted it and retry.", tarFile)
		return err
	}

	if _, err = os.Stat(blackboxDir); err == nil {
		err = fmt.Errorf("The path %s already exists.  Please deleted it and retry.", blackboxDir)
		return err
	}

	downloadCmd := exec.Command("wget", "--quiet", "--no-check-certificate", "https://whisk.sl.cloud9.ibm.com/"+tarFile)

	if err = downloadCmd.Run(); err != nil {
		err = errors.New("Download of docker skeleton failed")
		return err
	}

	installCmd := exec.Command("tar", "pxf", tarFile)

	if err = installCmd.Run(); err != nil {
		err = errors.New("Could not install docker skeleton")
		return err
	}

	rmCmd := exec.Command("rm", tarFile)
	if err = rmCmd.Run(); err != nil {
		// Don't really care...
	}

	fmt.Println("The docker skeleton is now installed at the current directory.")

	return nil
}

func swiftInstall() error {
	fmt.Println("swift SDK coming soon")
	return nil
}

func iOSInstall() error {
	var err error
	zipFile := "WhiskIOSStarterApp.zip"
	if _, err = os.Stat(zipFile); err == nil {
		err = fmt.Errorf("The path %s already exists.  Please delete it and try again", zipFile)
		return err
	}

	url := fmt.Sprintf("https://%s/%s", Properties.APIHost, zipFile)
	downloadCmd := exec.Command("wget", "--quiet", "--no-check-certificate", url)
	if err = downloadCmd.Run(); err != nil {
		err = errors.New("Download of iOS Whisk starter app failed.")
		return err
	}

	fmt.Printf("Downloaded iOS whisk starter app. Unzip %s and open the project in Xcode", zipFile)

	return nil
}

func init() {
	sdkCmd.AddCommand(sdkInstallCmd)
}
