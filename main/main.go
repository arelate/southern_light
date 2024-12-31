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

	scDir := filepath.Join(ucDir, "Steam", "config")
	fn := "loginusers.vdf"
	fp := filepath.Join(scDir, fn)

	keyValues, err := steam_vdf.Parse(fp)
	if err != nil {
		panic(err)
	}

	personaName := steam_vdf.GetKevValuesByKey(keyValues, "PersonaName")

	if personaName != nil {
		fmt.Println(personaName.Key, *personaName.Value)
	} else {
		fmt.Println("PersonaName not found")
	}

	users := steam_vdf.GetKevValuesByKey(keyValues, "users")
	if users != nil {
		for _, kv := range users.Values {
			steamId, err := steam_integration.SteamIdFromUserId(kv.Key)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Current user Steam Id: %d\n", steamId)

		}
	} else {
		fmt.Println("users not found")
	}

	//uhDir, err := os.UserHomeDir()
	//if err != nil {
	//	panic(err)
	//}
	//
	//downloadsDir := filepath.Join(uhDir, "Downloads")
	//ofn := filepath.Join(downloadsDir, fn)
	//
	//if err := steam_vdf.Write(ofn, keyValues); err != nil {
	//	panic(err)
	//}

}
