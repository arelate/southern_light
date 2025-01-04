package main

import (
	"fmt"
	"github.com/arelate/southern_light/steam_vdf"
	"os"
	"path/filepath"
)

func main() {
	uhDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	downloadsDir := filepath.Join(uhDir, "Downloads")
	fn := "shortcuts.vdf"
	ifp := filepath.Join(downloadsDir, fn)

	inputFile, err := os.Open(ifp)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	kv, err := steam_vdf.ParseBinary(ifp)
	if err != nil {
		panic(err)
	}

	fmt.Println(kv)

	// This sample code only makes sense for loginusers.vdf

	//personaName := steam_vdf.GetKevValuesByKey(keyValues, "PersonaName")
	//if personaName != nil {
	//	fmt.Println(personaName.Key, *personaName.Value)
	//} else {
	//	fmt.Println("PersonaName not found")
	//}
	//
	//users := steam_vdf.GetKevValuesByKey(keyValues, "users")
	//if users != nil {
	//	for _, kv := range users.Values {
	//		steamId, err := steam_integration.SteamIdFromUserId(kv.Key)
	//		if err != nil {
	//			panic(err)
	//		}
	//		fmt.Printf("Current user Steam Id: %d\n", steamId)
	//
	//	}
	//	//if err := users.WriteString(os.Stdout, 0); err != nil {
	//	//	panic(err)
	//	//}
	//} else {
	//	fmt.Println("users not found")
	//}

	//ofp := filepath.Join(os.TempDir(), fn)
	//
	//if err := steam_vdf.WriteText(ofp, keyValues); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("file://" + ofp)
	//
	//if err := compareFileHashes(ifp, ofp); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("input and output files are identical")
	//}

}
