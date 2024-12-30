package main

import (
	"fmt"
	"github.com/arelate/southern_light/steam_integration"
	"github.com/arelate/southern_light/steam_vdf"
	"os"
	"path/filepath"
)

func main() {
	ucDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	scDir := filepath.Join(ucDir, "Steam")
	fn := "config/loginusers.vdf"
	fp := filepath.Join(scDir, fn)

	keyValues, err := steam_vdf.Parse(fp)
	if err != nil {
		panic(err)
	}

	if val := steam_vdf.GetValByRelKey(keyValues, "PersonaName"); val != nil {
		fmt.Println(*val)
	}

	if userIds := steam_vdf.GetDirectChildren(keyValues, "users"); len(userIds) > 0 {
		for _, userId := range userIds {
			fmt.Println(steam_integration.SteamIdFromUserId(userId))
		}
	}
}
