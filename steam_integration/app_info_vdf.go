package steam_integration

import (
	"errors"
	"strconv"

	"github.com/arelate/southern_light/steam_vdf"
)

func AppInfoVdf(keyValues []*steam_vdf.KeyValues) (*AppInfo, error) {

	appInfo := &AppInfo{}

	if len(keyValues) == 0 {
		return nil, errors.New("cannot create appinfo from empty key values")
	}

	if len(keyValues) > 1 {
		return nil, errors.New("appinfo can only be created from a single key values")
	}

	for _, rootKv := range keyValues {

		appInfo.AppId = rootKv.Key

		for _, rootValues := range rootKv.Values {
			switch rootValues.Key {
			case "common":
				if aiCommon, err := appInfoCommonVdf(rootValues); err == nil {
					appInfo.Common = aiCommon
				} else {
					return nil, err
				}
			case "extended":
				if aiExtended, err := appInfoExtendedVdf(rootValues); err == nil {
					appInfo.Extended = aiExtended
				} else {
					return nil, err
				}
			case "config":
				if aiConfig, err := appInfoConfigVdf(rootValues); err == nil {
					appInfo.Config = aiConfig
				} else {
					return nil, err
				}
			case "depots":
				if aiDepots, err := appInfoDepotsVdf(rootValues); err == nil {
					appInfo.Depots = aiDepots
				} else {
					return nil, err
				}
			case "ufs":
				if aiUfs, err := appInfoUfsVdf(rootValues); err == nil {
					appInfo.Ufs = aiUfs
				} else {
					return nil, err
				}
			default:
				return nil, errors.New("unknown root appinfo key: " + rootValues.Key)
			}
		}
	}

	return appInfo, nil
}

func appInfoCommonVdf(commonKeyValues *steam_vdf.KeyValues) (*AppInfoCommon, error) {

	appInfoCommon := &AppInfoCommon{}

	var err error

	for _, commonKv := range commonKeyValues.Values {

		var strVal string
		if commonKv.Value != nil {
			strVal = *commonKv.Value
		}

		switch commonKv.Key {
		case "name":
			appInfoCommon.Name = strVal
		case "type":
			appInfoCommon.Type = strVal
		case "ReleaseState":
			appInfoCommon.ReleaseState = strVal
		case "logo":
			appInfoCommon.Logo = strVal
		case "logo_small":
			appInfoCommon.LogoSmall = strVal
		case "clienticon":
			appInfoCommon.ClientIcon = strVal
		case "clienttga":
			appInfoCommon.ClientTga = strVal
		case "icon":
			appInfoCommon.Icon = strVal
		case "oslist":
			appInfoCommon.OsList = strVal
		case "osarch":
			appInfoCommon.OsArch = strVal
		case "osextended":
			appInfoCommon.OsExtended = strVal
		case "languages":
			appInfoCommon.Languages = sliceFromKeys(commonKv)
		case "steam_deck_compatibility":
			appInfoCommon.SteamDeckCompatibility = steamDeckCompatibilityVdf(commonKv)
		case "metacritic_name":
			appInfoCommon.MetacriticName = strVal
		case "controller_support":
			appInfoCommon.ControllerSupport = strVal
		case "small_capsule":
			appInfoCommon.SmallCapsule = mapFromValues(commonKv)
		case "header_image":
			appInfoCommon.HeaderImage = mapFromValues(commonKv)
		case "library_assets":
			appInfoCommon.LibraryAssets = libraryAssetsVdf(commonKv)
		case "library_assets_full":
			appInfoCommon.LibraryAssetsFull = libraryAssetsFullVdf(commonKv)
		case "store_asset_mtime":
			appInfoCommon.StoreAssetMtime, err = strconv.ParseInt(strVal, 10, 64)
		case "associations":
			appInfoCommon.Associations = commonAssociationsVdf(commonKv)
		case "primary_genre":
			appInfoCommon.PrimaryGenre = strVal
		case "genres":
			appInfoCommon.Genres = sliceFromValues(commonKv)
		case "category":
			appInfoCommon.Category = sliceFromKeys(commonKv)
		case "supported_languages":
			appInfoCommon.SupportedLanguages = supportedLanguagesVdf(commonKv)
		case "steam_release_date":
			appInfoCommon.SteamReleaseDate, err = strconv.ParseInt(strVal, 10, 64)
		case "metacritic_score":
			appInfoCommon.MetacriticScore, err = strconv.Atoi(strVal)
		case "metacritic_fullurl":
			appInfoCommon.MetacriticFullUrl = strVal
		case "community_visible_stats":
			appInfoCommon.CommunityVisibleStats, err = strconv.Atoi(strVal)
		case "community_hub_visible":
			appInfoCommon.CommunityHubVisible, err = strconv.Atoi(strVal)
		case "gameid":
			appInfoCommon.GameId, err = strconv.ParseInt(strVal, 10, 32)
		case "store_tags":
			appInfoCommon.StoreTags = sliceFromValues(commonKv)
		case "review_score":
			appInfoCommon.ReviewScore, err = strconv.Atoi(strVal)
		case "review_percentage":
			appInfoCommon.ReviewPercentage, err = strconv.Atoi(strVal)
		default:
			return nil, errors.New("unknown appinfo common key: " + commonKv.Key)
		}

		if err != nil {
			return nil, err
		}
	}
	return appInfoCommon, nil
}

func mapFromValues(kv *steam_vdf.KeyValues) map[string]string {
	kvm := make(map[string]string)
	for _, vkv := range kv.Values {
		if vkv.Value != nil {
			kvm[vkv.Key] = *vkv.Value
		}
	}
	return kvm
}

func sliceFromValues(kv *steam_vdf.KeyValues) []string {
	kvs := make([]string, 0, len(kv.Values))
	for _, vkv := range kv.Values {
		if vkv.Value != nil {
			kvs = append(kvs, *vkv.Value)
		}
	}
	return kvs
}

func sliceFromKeys(kv *steam_vdf.KeyValues) []string {
	kvs := make([]string, 0, len(kv.Values))
	for _, vkv := range kv.Values {
		kvs = append(kvs, vkv.Key)
	}
	return kvs
}

func commonAssociationsVdf(kv *steam_vdf.KeyValues) map[string]string {
	kvm := make(map[string]string)

	for _, association := range kv.Values {

		var at, an string

		for _, atn := range association.Values {
			switch atn.Key {
			case "type":
				at = *atn.Value
			case "name":
				an = *atn.Value
			default:
				panic("unknown associations key: " + atn.Key)
			}
		}

		kvm[at] = an
	}

	return kvm
}

func libraryAssetsVdf(kv *steam_vdf.KeyValues) *LibraryAssets {
	la := new(LibraryAssets)

	for _, libraryKv := range kv.Values {

		var strVal string
		if libraryKv.Value != nil {
			strVal = *libraryKv.Value
		}

		switch libraryKv.Key {
		case "library_capsule":
			la.LibraryCapsule = strVal
		case "library_hero":
			la.LibraryHero = strVal
		case "library_logo":
			la.LibraryLogo = strVal
		case "logo_position":
			la.LogoPosition = logoPositionVdf(libraryKv)
		default:
			panic("unknown library_assets key: " + libraryKv.Key)
		}
	}

	return la
}

func libraryAssetsFullVdf(kv *steam_vdf.KeyValues) *LibraryAssetsFull {
	laf := new(LibraryAssetsFull)

	for _, lafKv := range kv.Values {

		switch lafKv.Key {
		case "library_capsule":
			laf.LibraryCapsule = image2xAssetsVdf(lafKv)
		case "library_hero":
			laf.LibraryHero = image2xAssetsVdf(lafKv)
		case "library_logo":
			laf.LibraryLogo = libraryLogoAssetsFullVdf(lafKv)
		default:
			panic("unknown library_assets_full key: " + lafKv.Key)
		}
	}

	return laf
}

func imageVdf(kv *steam_vdf.KeyValues) map[string]string {
	im := make(map[string]string)
	for _, imageKv := range kv.Values {
		if imageKv.Value != nil {
			im[imageKv.Key] = *imageKv.Value
		}
	}
	return im
}

func image2xAssetsVdf(kv *steam_vdf.KeyValues) *Image2xAssets {
	i2xa := new(Image2xAssets)
	i2xa.Image = make(map[string]string)
	i2xa.Image2x = make(map[string]string)

	for _, ia := range kv.Values {

		switch ia.Key {
		case "image":
			i2xa.Image = imageVdf(ia)
		case "image2x":
			i2xa.Image2x = imageVdf(ia)
		case "logo_position": // ignore
		default:
			panic("unknown image2x_assets key: " + ia.Key)
		}
	}

	return i2xa
}

func libraryLogoAssetsFullVdf(kv *steam_vdf.KeyValues) *LibraryLogoAssetsFull {

	llaf := new(LibraryLogoAssetsFull)
	llaf.Image2xAssets = image2xAssetsVdf(kv)

	for _, llafKv := range kv.Values {

		switch llafKv.Key {
		case "image": // ignore
		case "image2x": // ignore
		case "logo_position":
			llaf.LogoPosition = logoPositionVdf(llafKv)
		default:
			panic("unknown library_logo_assets_full key: " + llafKv.Key)
		}
	}

	return llaf
}

func logoPositionVdf(kv *steam_vdf.KeyValues) *LogoPosition {
	lp := new(LogoPosition)

	var err error

	for _, lpKv := range kv.Values {

		var strVal string
		if lpKv.Value != nil {
			strVal = *lpKv.Value
		}

		switch lpKv.Key {
		case "pinned_position":
			lp.PinnedPosition = strVal
		case "width_pct":
			lp.WidthPct, err = strconv.ParseFloat(strVal, 64)
		case "height_pct":
			lp.HeighPct, err = strconv.ParseFloat(strVal, 64)
		default:
			panic("unknown logo_position key: " + lpKv.Key)
		}

		if err != nil {
			panic(err)
		}
	}

	return lp
}

func supportedLanguagesVdf(kv *steam_vdf.KeyValues) map[string]*LanguageSupport {

	slm := make(map[string]*LanguageSupport)

	for _, langKv := range kv.Values {

		ls := new(LanguageSupport)

		for _, lsKv := range langKv.Values {
			switch lsKv.Key {
			case "supported":
				ls.Supported = true
			case "full_audio":
				ls.FullAudio = true
			case "subtitles":
				ls.Subtitles = true
			default:
				panic("unknown supported_languages key: " + lsKv.Key)
			}
		}

		slm[langKv.Key] = ls
	}

	return slm
}

func steamDeckCompatibilityVdf(kv *steam_vdf.KeyValues) *SteamDeckCompatibility {
	sdc := new(SteamDeckCompatibility)

	var err error

	for _, sdcKv := range kv.Values {

		var strVal string
		if sdcKv.Value != nil {
			strVal = *sdcKv.Value
		}

		switch sdcKv.Key {
		case "category":
			sdc.Category, err = strconv.Atoi(strVal)
		case "steamos_compatibility":
			sdc.SteamOsCompatibility, err = strconv.Atoi(strVal)
		case "test_timestamp":
			sdc.TestTimestamp, err = strconv.ParseInt(strVal, 10, 64)
		case "tested_build_id":
			sdc.TestBuildId, err = strconv.ParseInt(strVal, 10, 64)
		case "tests":
			sdc.Tests = testResultsVdf(sdcKv)
		case "steamos_tests":
			sdc.SteamOsTests = testResultsVdf(sdcKv)
		case "configuration":
			sdc.Configuration = mapFromValues(sdcKv)
		default:
			panic("unknown steam_deck_compatibility key: " + sdcKv.Key)
		}

		if err != nil {
			panic(err)
		}
	}

	return sdc
}

func testResultsVdf(kv *steam_vdf.KeyValues) []SteamTestResult {
	strs := make([]SteamTestResult, 0, len(kv.Values))

	for _, testKv := range kv.Values {

		tr := SteamTestResult{}

		for _, testResultsKv := range testKv.Values {

			var strVal string
			if testResultsKv.Value != nil {
				strVal = *testResultsKv.Value
			}

			var err error

			switch testResultsKv.Key {
			case "display":
				tr.Display, err = strconv.Atoi(strVal)
			case "token":
				tr.Token = strVal
			default:
				panic("unknown test result key: " + testResultsKv.Key)
			}

			if err != nil {
				panic(err)
			}

		}

		strs = append(strs, tr)
	}

	return strs
}

func appInfoExtendedVdf(extendedKeyValues *steam_vdf.KeyValues) (*AppInfoExtended, error) {
	return nil, nil
}

func appInfoConfigVdf(configKeyValues *steam_vdf.KeyValues) (*AppInfoConfig, error) {
	return nil, nil
}

func appInfoDepotsVdf(depotsKeyValues *steam_vdf.KeyValues) (*AppInfoDepots, error) {
	return nil, nil
}

func appInfoUfsVdf(ufsKeyValues *steam_vdf.KeyValues) (*AppInfoUfs, error) {
	return nil, nil
}
