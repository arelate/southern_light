package main

import (
	"fmt"
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
	fn := "libraryfolders.vdf"
	fp := filepath.Join(scDir, fn)

	keyValues, err := steam_vdf.Parse(fp)
	if err != nil {
		panic(err)
	}

	fmt.Println(keyValues)

	//if val := steam_vdf.GetValByRelKey(keyValues, "PersonaName"); val != nil {
	//	fmt.Println(*val)
	//}
	//
	//if userIds := steam_vdf.GetDirectChildren(keyValues, "users"); len(userIds) > 0 {
	//	for _, userId := range userIds {
	//		fmt.Println(steam_integration.SteamIdFromUserId(userId))
	//	}
	//}

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
