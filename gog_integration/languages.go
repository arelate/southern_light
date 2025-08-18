package gog_integration

import (
	"iter"
	"maps"
)

var languageNames = map[string]string{
	"ar":     "Arabic",
	"be":     "Belarusian",
	"bl":     "Bulgarian",
	"br":     "Portuguese (Brazilian)",
	"ca":     "Catalan",
	"cn":     "Chinese Simplified",
	"cz":     "Czech",
	"da":     "Danish",
	"de":     "German",
	"en":     "English",
	"es":     "Spanish",
	"es_mx":  "Latin American Spanish",
	"et":     "Estonian",
	"fa":     "Farsi",
	"fi":     "Finnish",
	"fr":     "French",
	"gk":     "Greek",
	"gog_IN": "Inuktitut",
	"he":     "Hebrew",
	"hr":     "Croatian",
	"hu":     "Hungarian",
	"id":     "Indonesian",
	"is":     "Icelandic",
	"it":     "Italian",
	"jp":     "Japanese",
	"ko":     "Korean",
	"la":     "Latin",
	"nl":     "Dutch",
	"no":     "Norwegian",
	"pl":     "Polish",
	"pt":     "Portuguese",
	"ro":     "Romanian",
	"ru":     "Russian",
	"sb":     "Serbian",
	"sk":     "Slovak",
	"sv":     "Swedish",
	"th":     "Thai",
	"tr":     "Turkish",
	"uk":     "Ukrainian",
	"vi":     "Vietnamese",
	"zh":     "Chinese Traditional",
}

var nativeLanguageNames = map[string]string{
	"ar":     "العربية",
	"be":     "беларуская",
	"bl":     "български",
	"br":     "Português do Brasil",
	"ca":     "català",
	"cn":     "中文(简体)",
	"cz":     "český",
	"da":     "Dansk",
	"de":     "Deutsch",
	"en":     "English",
	"es":     "español",
	"es_mx":  "Español (AL)",
	"et":     "eesti",
	"fa":     "فارسی",
	"fi":     "suomi",
	"fr":     "français",
	"gk":     "Ελληνικά",
	"gog_IN": "Inuktitut",
	"he":     "עברית",
	"hr":     "hrvatski",
	"hu":     "magyar",
	"id":     "bahasa Indonesia",
	"is":     "Íslenska",
	"it":     "italiano",
	"jp":     "日本語",
	"ko":     "한국어",
	"la":     "latine",
	"nl":     "nederlands",
	"no":     "norsk",
	"pl":     "polski",
	"pt":     "português",
	"ro":     "română",
	"ru":     "русский",
	"sb":     "Српска",
	"sk":     "slovenský",
	"sv":     "svenska",
	"th":     "ไทย",
	"tr":     "Türkçe",
	"uk":     "yкраїнська",
	"vi":     "Tiếng Việt",
	"zh":     "中文(繁體)",
}

func LanguageCodeByNativeName(name string) string {
	for lc, nln := range nativeLanguageNames {
		if nln == name {
			return lc
		}
	}
	return ""
}

func LanguageNativeName(langCode string) string {
	if nln, ok := nativeLanguageNames[langCode]; ok {
		return nln
	} else {
		return "Unknown"
	}
}

func AllLanguageCodes() iter.Seq[string] {
	return maps.Keys(languageNames)
}

func LanguageCodesCloValues() []string {
	defaultLangCode := "en"
	langCodes := []string{defaultLangCode}
	for lc := range AllLanguageCodes() {
		if lc == defaultLangCode {
			continue
		}
		langCodes = append(langCodes, lc)
	}
	return langCodes
}
