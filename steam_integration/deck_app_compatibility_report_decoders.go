package steam_integration

import (
	"strings"
)

const (
	steamDeckLocTokenPrefix = "#SteamDeckVerified_TestResult_"
	steamOsLocTokenPrefix   = "#SteamOS_TestResult_"
)

// Source: https://store.fastly.steamstatic.com/public/javascript/applications/store/shared_english-json.js
var steamDeckLocTokensStrings = map[string]string{
	"AudioOutputHasNonblockingIssues":                      "This game has minor audio issues on Steam Deck",
	"AuxFunctionalityNotAccessibleMapEditor":               "Some auxilliary functionality is not accessible on Steam Deck: map editor",
	"AuxFunctionalityNotAccessible_MapEditor":              "Some auxilliary functionality is not accessible on Steam Deck: map editor",
	"CloudSavesNotEnabledByDefault":                        "This game requires manually enabling Steam Cloud support for saved games using in-game settings",
	"ControllerGlyphsMatchDeckDevice":                      "This game shows Steam Deck controller icons",
	"ControllerGlyphsDoNotMatchDeckDevice":                 "This game sometimes shows mouse, keyboard, or non-Steam-Deck controller icons",
	"CrossPlatformCloudSavesNotSupported":                  "This game does not support cross-platform saved games",
	"DefaultConfigurationIsPerformant":                     "This game's default graphics configuration performs well on Steam Deck",
	"DefaultConfigurationIsNotPerformant":                  "This game requires manual configuration of graphics settings to perform well on Steam Deck",
	"DefaultControllerConfigFullyFunctional":               "All functionality is accessible when using the default controller configuration",
	"DefaultControllerConfigNotFullyFunctional":            "Some functionality is not accessible when using the default controller configuration, requiring use of the touchscreen or virtual keyboard, or a community configuration",
	"DeviceCompatibilityWarningsShown":                     "This game displays compatibility warnings when running on Steam Deck, but runs fine",
	"DisplayOutputHasNonblockingIssues":                    "This game has minor graphics/display issues on Steam Deck",
	"DisplayOutputNotCorrectlyScaled":                      "This game is incorrectly scaled on Steam Deck and may require you to configure the display resolution manually",
	"ExternalControllersNotSupportedLocalMultiplayer":      "This game does not support external Bluetooth/USB controllers on Deck for local multiplayer",
	"ExternalControllersNotSupportedPrimaryPlayer":         "This game does not default to external Bluetooth/USB controllers on Deck, and may require manually switching the active controller via the Quick Access Menu",
	"FirstTimeSetupRequiresActiveInternetConnection":       "This game's first-time setup requires an active Internet connection",
	"GameOrLauncherDoesntExitCleanly":                      "This game does not exit cleanly, and may require you to manually quit via the Steam overlay",
	"GamepadNotEnabledByDefault":                           "This game requires manually enabling controller support using in-game settings",
	"InterfaceTextIsLegible":                               "In-game interface text is legible on Steam Deck",
	"InterfaceTextIsNotLegible":                            "Some in-game text is small and may be difficult to read",
	"LauncherInteractionIssues":                            "This game's launcher/setup tool may require the touchscreen or virtual keyboard, or have difficult to read text",
	"MultiWindowAppAutomaticallySetsFocus":                 "This game uses multiple windows and may require you to focus the appropriate windows manually via the Steam overlay",
	"NativeResolutionNotDefault":                           "This game supports Steam Deck's native display resolution but does not set it by default and may require you to configure the display resolution manually",
	"NativeResolutionNotSupported":                         "This game doesn't support Steam Deck's native display resolution and may experience degraded performance",
	"NotFullyFunctionalWithoutExternalKeyboard":            "Parts of this game would benefit from using an external keyboard",
	"ResumeFromSleepNotFunctional":                         "This game may experience temporary issues after sleeping on Steam Deck",
	"SimultaneousInputGyroTrackpadFriendly":                "This game supports using the gyro/trackpad in mouse mode for camera controls with gamepad controls for movement",
	"SingleplayerGameplayRequiresActiveInternetConnection": "Singleplayer gameplay requires an active Internet connection",
	"SteamOSDoesNotSupport":                                "Valve is still working on adding support for this game on Steam Deck",
	"SteamOSDoesNotSupport_Software":                       "Valve is still working on adding support for this game on Steam Deck",
	"SteamOSDoesNotSupportVR":                              "Steam Deck does not support VR games",
	"SteamOSDoesNotSupport_VR":                             "Steam Deck does not support VR games",
	"TextInputDoesNotAutomaticallyInvokesKeyboard":         "Entering some text requires manually invoking the on-screen keyboard",
	"UnsupportedAntiCheatConfiguration":                    "This game's anti-cheat is not configured to support Steam Deck",
	"UnsupportedAntiCheatOther":                            "This game is unsupported on Steam Deck due to use of an unsupported anti-cheat",
	"UnsupportedGraphicsPerformance":                       "This game's graphics settings cannot be configured to run well on Steam Deck",
	"VideoPlaybackHasNonblockingIssues":                    "Some in-game movie content may be missing",
}

// Source: same link as above
var steamOsLocTokensStrings = map[string]string{
	"GameStartupFunctional":                                "This game runs successfully on SteamOS",
	"SteamOSDoesNotSupport":                                "Valve is still working on adding support for this game on SteamOS",
	"UnsupportedAntiCheatConfiguration":                    "This game's anti-cheat is not configured to support SteamOS",
	"UnsupportedAntiCheat_Other":                           "This game is unsupported on SteamOS due to use of an unsupported anti-cheat or multiplayer service",
	"SteamOSDoesNotSupport_Software":                       "This is a Software title and is not supported on SteamOS",
	"SteamOSDoesNotSupport_Retired":                        "This game has been retired or is no longer in a playable state and is not supported on SteamOS",
	"SteamOSDoesNotSupport_OperatingSystem":                "This game requires an operating system that is not currently supported on SteamOS",
	"LauncherInteractionIssues":                            "This game's launcher/setup tool may require a touchscreen or virtual keyboard, or have difficult to read text",
	"TextInputDoesNotAutomaticallyInvokesKeyboard":         "Entering some text requires manually invoking the on-screen keyboard",
	"DefaultControllerConfigNotFullyFunctional":            "Some functionality is not accessible when using the default controller configuration, requiring use of the touchscreen or virtual keyboard, or a community configuration",
	"GameOrLauncherDoesntExitCleanly":                      "This game does not exit cleanly and may require you to manually quit via the Steam overlay",
	"VideoPlaybackHasNonblockingIssues":                    "Some in-game movie content may be missing or have playback issues",
	"DisplayOutputHasNonblockingIssues":                    "This game has minor graphics/display issues on SteamOS",
	"AudioOutputHasNonblockingIssues":                      "This game has minor audio issues on SteamOS",
	"AuxFunctionalityNotAccessible_MapEditor":              "Some auxilliary functionality is not accessible on SteamOS: map editor",
	"MultiWindowAppAutomaticallySetsFocus":                 "This game uses multiple windows and may require you to focus the appropriate windows manually via the Steam overlay",
	"DisplayOutputNotCorrectlyScaled":                      "This game is incorrectly scaled on SteamOS and may require you to configure the display resolution manually",
	"ResumeFromSleepNotFunctional":                         "This game may experience temporary issues after sleeping on SteamOS",
	"GamepadNotEnabledByDefault":                           "This game requires manually enabling controller support using in-game settings",
	"CloudSavesNotEnabledByDefault":                        "This game requires manually enabling Steam Cloud support for saved games using in-game settings",
	"NotFullyFunctionalWithoutExternalKeyboard":            "Parts of this game would benefit from using an external keyboard",
	"NotFullyFunctionalWithoutExternalWebcam":              "This game requires certain external devices: Webcam",
	"NotFullyFunctionalWithoutExternalUSBGuitar":           "This game requires certain external devices: USB Guitar",
	"FirstTimeSetupRequiresActiveInternetConnection":       "This game's first-time setup requires an active Internet connection",
	"SingleplayerGameplayRequiresActiveInternetConnection": "Singleplayer gameplay requires an active Internet connection",
	"CrossPlatformCloudSavesNotSupported":                  "This game does not support cross-platform saved games",
	"ExternalControllersNotSupportedPrimaryPlayer":         "This game does not default to external Bluetooth/USB controllers on SteamOS, and may require manually switching the active controller via the Quick Access Menu",
	"ExternalControllersNotSupportedLocalMultiplayer":      "This game does not support external Bluetooth/USB controllers on SteamOS for local multiplayer",
	"SimultaneousInputGyroTrackpadFriendly":                "This game supports using the gyro/trackpad in mouse mode for camera controls with gamepad controls for movement",
	"HDRMustBeManuallyEnabled":                             "This game supports HDR but it must be manually enabled using in-game settings",
}

var categoryTitles = map[int]string{
	0: "Unknown",
	1: "Unsupported",
	2: "Playable",
	3: "Verified",
}

func DecodeCategory(category int) string {
	return categoryTitles[category]
}

func SteamDeckTrimLocToken(token string) string {
	return strings.TrimPrefix(token, steamDeckLocTokenPrefix)
}

func SteamDeckDecodeLocToken(token string) string {
	return steamDeckLocTokensStrings[token]
}

func SteamOsTrimLocToken(token string) string {
	return strings.TrimPrefix(token, steamOsLocTokenPrefix)
}

func SteamOsDecodeLocToken(token string) string {
	return steamOsLocTokensStrings[token]
}
