package steamcmd

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/arelate/southern_light/vangogh_integration"
)

const (
	steamCmdLoginCommand           = "+login"
	steamCmdForceInstallDirCommand = "+force_install_dir"
	steamCmdAppUpdateCommand       = "+app_update"
	steamCmdAppInfoPrintCommand    = "+app_info_print"
	steamCmdAppUninstallCommand    = "+app_uninstall"
	steamCmdLogoutCommand          = "+logout"
	steamCmdQuitCommand            = "+quit"
)

const (
	steamCmdShutdownOnFailedCommandVariable = "+@ShutdownOnFailedCommand"
	steamCmdNoPromptForPasswordVariable     = "+@NoPromptForPassword"
	steamCmdForcePlatformTypeVariable       = "+@sSteamCmdForcePlatformType"
)

const (
	steamCmdTrueValue      = "1"
	steamCmdFalseValue     = "0"
	steamCmdAnonymousValue = "anonymous"
)

func runSteamCmdCommands(absSteamCmdPath string, commands ...string) (*exec.Cmd, error) {

	cmd := exec.Command(absSteamCmdPath, commands...)

	// always enable i/o to be able to input Steam password for a given username
	// and see Steam Guard prompt
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil
}

func Login(absSteamCmdPath string, username string) error {

	cmd, err := runSteamCmdCommands(absSteamCmdPath, steamCmdLoginCommand, username, steamCmdQuitCommand)
	if err != nil {
		return err
	}

	return cmd.Run()
}

func Logout(absSteamCmdPath string) error {

	cmd, err := runSteamCmdCommands(absSteamCmdPath, steamCmdLogoutCommand, steamCmdQuitCommand)
	if err != nil {
		return err
	}

	return cmd.Run()
}

func AppInfoPrint(absSteamCmdPath string, id string) (string, error) {

	appInfoPrintCmd, err := runSteamCmdCommands(absSteamCmdPath,
		steamCmdShutdownOnFailedCommandVariable, steamCmdTrueValue,
		steamCmdNoPromptForPasswordVariable, steamCmdTrueValue,
		steamCmdLoginCommand, steamCmdAnonymousValue,
		steamCmdAppInfoPrintCommand, id,
		steamCmdQuitCommand)
	if err != nil {
		return "", err
	}

	stdout := bytes.NewBuffer(nil)

	appInfoPrintCmd.Stdout = stdout
	appInfoPrintCmd.Stderr = stdout

	if err = appInfoPrintCmd.Run(); err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(stdout)
	sb := new(strings.Builder)
	appinfo := false

	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "\"" + id + "\"":
			appinfo = true
		case "}":
			sb.WriteString(line)
			appinfo = false
		default:
			// do nothing
		}

		if appinfo {
			sb.WriteString(line)
		}
	}

	if scanner.Err() != nil {
		return "", err
	}

	return sb.String(), nil
}

func AppUpdate(absSteamCmdPath string, id string, operatingSystem vangogh_integration.OperatingSystem, absInstallDir, username string) error {

	steamAppUpdateCmd, err := runSteamCmdCommands(absSteamCmdPath,
		steamCmdShutdownOnFailedCommandVariable, steamCmdTrueValue,
		steamCmdNoPromptForPasswordVariable, steamCmdTrueValue,
		steamCmdForcePlatformTypeVariable, strings.ToLower(operatingSystem.String()),
		steamCmdForceInstallDirCommand, absInstallDir,
		steamCmdLoginCommand, username,
		steamCmdAppUpdateCommand, id,
		steamCmdQuitCommand)
	if err != nil {
		return err
	}

	return steamAppUpdateCmd.Run()
}

func AppUninstall(absSteamCmdPath string, id string, operatingSystem vangogh_integration.OperatingSystem, absInstallDir string) error {

	steamAppUninstallCmd, err := runSteamCmdCommands(absSteamCmdPath,
		steamCmdShutdownOnFailedCommandVariable, steamCmdTrueValue,
		steamCmdNoPromptForPasswordVariable, steamCmdTrueValue,
		steamCmdForcePlatformTypeVariable, strings.ToLower(operatingSystem.String()),
		steamCmdForceInstallDirCommand, absInstallDir,
		steamCmdLoginCommand, steamCmdAnonymousValue,
		steamCmdAppUninstallCommand, id,
		steamCmdQuitCommand)
	if err != nil {
		return err
	}

	return steamAppUninstallCmd.Run()

}
