package steam_integration

import (
	"github.com/arelate/southern_light"
	"strings"
)

const locTokenPrefix = "#SteamDeckVerified_TestResult_"

var locTokensStrings = map[string]string{
	"audioOutputHasNonblockingIssues":                      "This game has minor audio issues on Steam Deck",
	"auxFunctionalityNotAccessibleMapEditor":               "Some auxilliary functionality is not accessible on Steam Deck: map editor",
	"auxFunctionalityNotAccessible_MapEditor":              "Some auxilliary functionality is not accessible on Steam Deck: map editor",
	"cloudSavesNotEnabledByDefault":                        "This game requires manually enabling Steam Cloud support for saved games using in-game settings",
	"controllerGlyphsMatchDeckDevice":                      "This game shows Steam Deck controller icons",
	"controllerGlyphsDoNotMatchDeckDevice":                 "This game sometimes shows mouse, keyboard, or non-Steam-Deck controller icons",
	"crossPlatformCloudSavesNotSupported":                  "This game does not support cross-platform saved games",
	"defaultConfigurationIsPerformant":                     "This game's default graphics configuration performs well on Steam Deck",
	"defaultConfigurationIsNotPerformant":                  "This game requires manual configuration of graphics settings to perform well on Steam Deck",
	"defaultControllerConfigFullyFunctional":               "All functionality is accessible when using the default controller configuration",
	"defaultControllerConfigNotFullyFunctional":            "Some functionality is not accessible when using the default controller configuration, requiring use of the touchscreen or virtual keyboard, or a community configuration",
	"deviceCompatibilityWarningsShown":                     "This game displays compatibility warnings when running on Steam Deck, but runs fine",
	"displayOutputHasNonblockingIssues":                    "This game has minor graphics/display issues on Steam Deck",
	"displayOutputNotCorrectlyScaled":                      "This game is incorrectly scaled on Steam Deck and may require you to configure the display resolution manually",
	"externalControllersNotSupportedLocalMultiplayer":      "This game does not support external Bluetooth/USB controllers on Deck for local multiplayer",
	"externalControllersNotSupportedPrimaryPlayer":         "This game does not default to external Bluetooth/USB controllers on Deck, and may require manually switching the active controller via the Quick Access Menu",
	"firstTimeSetupRequiresActiveInternetConnection":       "This game's first-time setup requires an active Internet connection",
	"gameOrLauncherDoesntExitCleanly":                      "This game does not exit cleanly, and may require you to manually quit via the Steam overlay",
	"gamepadNotEnabledByDefault":                           "This game requires manually enabling controller support using in-game settings",
	"interfaceTextIsLegible":                               "In-game interface text is legible on Steam Deck",
	"interfaceTextIsNotLegible":                            "Some in-game text is small and may be difficult to read",
	"launcherInteractionIssues":                            "This game's launcher/setup tool may require the touchscreen or virtual keyboard, or have difficult to read text",
	"multiWindowAppAutomaticallySetsFocus":                 "This game uses multiple windows and may require you to focus the appropriate windows manually via the Steam overlay",
	"nativeResolutionNotDefault":                           "This game supports Steam Deck's native display resolution but does not set it by default and may require you to configure the display resolution manually",
	"nativeResolutionNotSupported":                         "This game doesn't support Steam Deck's native display resolution and may experience degraded performance",
	"notFullyFunctionalWithoutExternalKeyboard":            "Parts of this game would benefit from using an external keyboard",
	"resumeFromSleepNotFunctional":                         "This game may experience temporary issues after sleeping on Steam Deck",
	"simultaneousInputGyroTrackpadFriendly":                "This game supports using the gyro/trackpad in mouse mode for camera controls with gamepad controls for movement",
	"singleplayerGameplayRequiresActiveInternetConnection": "Singleplayer gameplay requires an active Internet connection",
	"steamOSDoesNotSupport":                                "Valve is still working on adding support for this game on Steam Deck",
	"steamOSDoesNotSupport_Software":                       "Valve is still working on adding support for this game on Steam Deck",
	"steamOSDoesNotSupportVR":                              "Steam Deck does not support VR games",
	"steamOSDoesNotSupport_VR":                             "Steam Deck does not support VR games",
	"textInputDoesNotAutomaticallyInvokesKeyboard":         "Entering some text requires manually invoking the on-screen keyboard",
	"unsupportedAntiCheatConfiguration":                    "This game's anti-cheat is not configured to support Steam Deck",
	"unsupportedAntiCheatOther":                            "This game is unsupported on Steam Deck due to use of an unsupported anti-cheat",
	"unsupportedGraphicsPerformance":                       "This game's graphics settings cannot be configured to run well on Steam Deck",
	"videoPlaybackHasNonblockingIssues":                    "Some in-game movie content may be missing",
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

func TrimLocToken(token string) string {
	return strings.TrimPrefix(token, locTokenPrefix)
}

func DecodeLocToken(token string) string {
	token = southern_light.FirstToLower(token)
	return locTokensStrings[token]
}
