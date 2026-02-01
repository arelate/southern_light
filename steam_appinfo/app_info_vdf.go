package steam_appinfo

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
				appInfo.Common = appInfoCommonVdf(rootValues)
			case "extended":
				appInfo.Extended = appInfoExtendedVdf(rootValues)
			case "config":
				appInfo.Config = appInfoConfigVdf(rootValues)
			case "depots":
				appInfo.Depots = appInfoDepotsVdf(rootValues)
			case "ufs":
				appInfo.Ufs = appInfoUfsVdf(rootValues)
			default:
				return nil, errors.New("unknown root appinfo key: " + rootValues.Key)
			}
		}
	}

	return appInfo, nil
}

func appInfoCommonVdf(commonKeyValues *steam_vdf.KeyValues) *AppInfoCommon {

	aic := &AppInfoCommon{}

	var err error

	for ii, commonKv := range commonKeyValues.Values {

		var strVal string
		if commonKv.Value != nil {
			strVal = *commonKv.Value
		}

		switch commonKv.Key {
		case "name":
			aic.Name = strVal
		case "name_localized":
			aic.NameLocalized = mapFromValues(commonKv)
		case "type":
			aic.Type = strVal
		case "ReleaseState":
			aic.ReleaseState = strVal
		case "logo":
			aic.Logo = strVal
		case "logo_small":
			aic.LogoSmall = strVal
		case "clienticon":
			aic.ClientIcon = strVal
		case "clienttga":
			aic.ClientTga = strVal
		case "icon":
			aic.Icon = strVal
		case "oslist":
			aic.OsList = strVal
		case "osarch":
			aic.OsArch = strVal
		case "osextended":
			aic.OsExtended = strVal
		case "languages":
			aic.Languages = sliceFromKeys(commonKv)
		case "content_descriptors":
			aic.ContentDescriptors = sliceFromValues(commonKv)
		case "content_descriptors_including_dlc":
			aic.ContentDescriptorsIncludingDlc = sliceFromValues(commonKv)
		case "steam_deck_compatibility":
			aic.SteamDeckCompatibility = steamDeckCompatibilityVdf(commonKv)
		case "steam_deck_blog_url":
			aic.SteamDeckBlogUrl = strVal
		case "controllertagwizard":
			aic.ControllerTagWizard = strVal
		case "metacritic_name":
			aic.MetacriticName = strVal
		case "controller_support":
			aic.ControllerSupport = strVal
		case "small_capsule":
			aic.SmallCapsule = mapFromValues(commonKv)
		case "header_image":
			aic.HeaderImage = mapFromValues(commonKv)
		case "library_assets":
			aic.LibraryAssets = libraryAssetsVdf(commonKv)
		case "library_assets_full":
			aic.LibraryAssetsFull = libraryAssetsFullVdf(commonKv)
		case "store_asset_mtime":
			aic.StoreAssetMtime, err = strconv.ParseInt(strVal, 10, 64)
		case "associations":
			aic.Associations = commonAssociationsVdf(commonKv)
		case "primary_genre":
			aic.PrimaryGenre = strVal
		case "genres":
			aic.Genres = sliceFromValues(commonKv)
		case "category":
			aic.Category = sliceFromKeys(commonKv)
		case "supported_languages":
			aic.SupportedLanguages = supportedLanguagesVdf(commonKv)
		case "steam_release_date":
			aic.SteamReleaseDate, err = strconv.ParseInt(strVal, 10, 64)
		case "metacritic_score":
			aic.MetacriticScore, err = strconv.Atoi(strVal)
		case "metacritic_fullurl":
			aic.MetacriticFullUrl = strVal
		case "community_visible_stats":
			aic.CommunityVisibleStats, err = strconv.Atoi(strVal)
		case "community_hub_visible":
			aic.CommunityHubVisible, err = strconv.Atoi(strVal)
		case "gameid":
			aic.GameId, err = strconv.ParseInt(strVal, 10, 32)
		case "store_tags":
			aic.StoreTags = sliceFromValues(commonKv)
		case "review_score":
			aic.ReviewScore, err = strconv.Atoi(strVal)
		case "review_percentage":
			aic.ReviewPercentage, err = strconv.Atoi(strVal)
		default:
			panic(errors.New("unknown appinfo common key: " + commonKv.Key + " at " + strconv.Itoa(ii)))
		}

		if err != nil {
			panic(err)
		}
	}
	return aic
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
		case "library_hero_blur":
			la.LibraryHeroBlur = strVal
		case "library_logo":
			la.LibraryLogo = strVal
		case "library_header":
			la.LibraryHeader = strVal
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
		case "library_hero_blur":
			laf.LibraryHeroBlur = image2xAssetsVdf(lafKv)
		case "library_header":
			laf.LibraryHeader = image2xAssetsVdf(lafKv)
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

func appInfoExtendedVdf(extendedKeyValues *steam_vdf.KeyValues) *AppInfoExtended {

	aie := new(AppInfoExtended)

	for ii, ekv := range extendedKeyValues.Values {

		var strVal string
		if ekv.Value != nil {
			strVal = *ekv.Value
		}

		switch ekv.Key {
		case "deckresolutionoverride":
			aie.DeckResolutionOverride = strVal
		case "developer":
			aie.Developer = strVal
		case "publisher":
			aie.Publisher = strVal
		case "homepage":
			aie.Homepage = strVal
		case "listofdlc":
			aie.ListOfDlc = strVal
		case "DLCAvailableOnStore":
			aie.DlcAvailableOnStore = strVal
		default:
			panic("unknown extended key: " + ekv.Key + " at " + strconv.Itoa(ii))
		}
	}

	return aie
}

func appInfoConfigVdf(configKeyValues *steam_vdf.KeyValues) *AppInfoConfig {

	aic := new(AppInfoConfig)

	for ii, ckv := range configKeyValues.Values {

		var strVal string
		if ckv.Value != nil {
			strVal = *ckv.Value
		}

		var err error

		switch ckv.Key {
		case "installdir":
			aic.InstallDir = strVal
		case "launch":
			aic.Launch = launchVdf(ckv)
		case "steamcontrollertouchtemplateindex":
			aic.SteamControllerTouchTemplateIndex, err = strconv.Atoi(strVal)
		case "steamcontrollertouchconfigdetails":
			// TODO
		case "steamcontrollertemplateindex":
			aic.SteamControllerTemplateIndex, err = strconv.Atoi(strVal)
		case "steamdecktouchscreen":
			aic.SteamDeckTouchScreen = strVal
		case "steamconfigurator3rdpartynative":
			aic.SteamConfigurator3rdPartyNative = strVal
		default:
			panic("unknown config key: " + ckv.Key + " at " + strconv.Itoa(ii))
		}

		if err != nil {
			panic(err)
		}
	}

	return aic
}

func launchVdf(kv *steam_vdf.KeyValues) []LaunchOption {

	los := make([]LaunchOption, 0)

	for _, launchKv := range kv.Values {

		lo := LaunchOption{
			DescriptionLoc: make(map[string]string),
		}

		for _, lkv := range launchKv.Values {

			var strVal string
			if lkv.Value != nil {
				strVal = *lkv.Value
			}

			switch lkv.Key {
			case "executable":
				lo.Executable = strVal
			case "arguments":
				lo.Arguments = strVal
			case "workingdir":
				lo.WorkingDir = strVal
			case "type":
				lo.Type = strVal
			case "config":
				lo.Config = LaunchOptionConfig{}
				for _, ckv := range lkv.Values {
					switch ckv.Key {
					case "oslist":
						if ckv.Value != nil {
							lo.Config.OsList = *ckv.Value
						}
					case "osarch":
						if ckv.Value != nil {
							lo.Config.OsArch = *ckv.Value
						}
					default:
						panic("unknown launch config key: " + ckv.Key)
					}
				}
			case "description_loc":
				for _, dlkv := range lkv.Values {
					if dlkv.Value != nil {
						lo.DescriptionLoc[dlkv.Key] = *dlkv.Value
					}
				}
			case "description":
				lo.Description = strVal
			default:
				panic("unknown launch key: " + lkv.Key)
			}
		}

		los = append(los, lo)
	}

	return los
}

func appInfoDepotsVdf(depotsKeyValues *steam_vdf.KeyValues) *AppInfoDepots {
	return nil
}

func appInfoUfsVdf(ufsKeyValues *steam_vdf.KeyValues) *AppInfoUfs {
	return nil
}
