package wine_integration

import (
	"maps"
	"slices"
)

const (
	ProtonEnableWayland = "enable-wayland"
	ProtonPreferSdl     = "prefer-sdl"
	ProtonDisableHidRaw = "disable-hidraw"
	ProtonEnableHdr     = "enable-hdr"
	ProtonNoSteamInput  = "no-steaminput"
	ProtonUseD7Vk       = "use-d7vk"
)

var ProtonOptionsEnv = map[string]string{
	ProtonEnableWayland: "PROTON_ENABLE_WAYLAND",
	ProtonPreferSdl:     "PROTON_PREFER_SDL",
	ProtonDisableHidRaw: "PROTON_DISABLE_HIDRAW",
	ProtonEnableHdr:     "PROTON_ENABLE_HDR",
	ProtonNoSteamInput:  "PROTON_NO_STEAMINPUT",
	ProtonUseD7Vk:       "PROTON_USE_D7VK",
}

const (
	SteamProton316                  = "proton-3.16"
	SteamProton370                  = "proton-3.7"
	SteamProton411                  = "proton-4.11"
	SteamProton420                  = "proton-4.2"
	SteamProton500                  = "proton-5.0"
	SteamProton513                  = "proton-5.13"
	SteamProton630                  = "proton-6.3"
	SteamProton700                  = "proton-7.0"
	SteamProton800                  = "proton-8.0"
	SteamProton900                  = "proton-9.0"
	SteamProton1000                 = "proton-10.0"
	SteamProtonNext                 = "proton-next"
	SteamProtonHotfix               = "proton-hotfix"
	SteamProtonExperimental         = "proton-experimental"
	SteamProtonBattlEyeRuntime      = "proton-battleye-runtime"
	SteamProtonEasyAntiCheatRuntime = "proton-easy-anti-cheat-runtime"
)

var ProtonRuntimesNames = map[string]string{
	UmuProton:      "umu-proton",
	ProtonGeCustom: "proton-ge",
	ProtonCachyOs:  "proton-cachyos",
	ProtonEm:       "proton-em",
}

var SteamProtonDirectories = map[string]string{
	SteamProton316:                  "Proton 3.16",
	SteamProton370:                  "Proton 3.7",
	SteamProton411:                  "Proton 4.11",
	SteamProton420:                  "Proton 4.2",
	SteamProton500:                  "Proton 5.0",
	SteamProton513:                  "Proton 5.13",
	SteamProton630:                  "Proton 6.3",
	SteamProton700:                  "Proton 7.0",
	SteamProton800:                  "Proton 8.0",
	SteamProton900:                  "Proton 9.0 (Beta)",
	SteamProton1000:                 "Proton 10.0",
	SteamProtonNext:                 "Proton Next",
	SteamProtonHotfix:               "Proton Hotfix",
	SteamProtonExperimental:         "Proton - Experimental",
	SteamProtonBattlEyeRuntime:      "Proton BattlEye Runtime",
	SteamProtonEasyAntiCheatRuntime: "Proton Easy Anti-Cheat Runtime",
}

func AllSteamProtonRuntimes() []string {
	return slices.Collect(maps.Keys(SteamProtonDirectories))
}

func AllProtonRuntimes() []string {
	return slices.Collect(maps.Values(ProtonRuntimesNames))
}

func AllProtonOptions() []string {
	return slices.Collect(maps.Keys(ProtonOptionsEnv))
}
