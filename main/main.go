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

//func Sha256(reader io.Reader) (string, error) {
//	h := sha256.New()
//	var err error
//	if _, err = io.Copy(h, reader); err == nil {
//		return fmt.Sprintf("%x", h.Sum(nil)), nil
//	}
//	return "", err
//}
//
//func compareFileHashes(path1, path2 string) error {
//
//	f1, err := os.Open(path1)
//	if err != nil {
//		return err
//	}
//	defer f1.Close()
//
//	fhash1, err := Sha256(f1)
//	if err != nil {
//		return err
//	}
//
//	f2, err := os.Open(path2)
//	if err != nil {
//		return err
//	}
//	defer f2.Close()
//
//	fhash2, err := Sha256(f2)
//	if err != nil {
//		return err
//	}
//
//	if fhash1 != fhash2 {
//		return errors.New("files produced different hashes")
//	}
//
//	return nil
//}
