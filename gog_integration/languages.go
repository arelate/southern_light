package gog_integration

import (
	"iter"
	"maps"
)

// Note: to update the list - navigate to gog.com/account, open page source, Console
// run `document.querySelector("[ng-controller='accountLanguageFiltersCtrl as language']")`
// Copy HTML and extract those pairs

var nativeLanguageNames = map[string]string{
	"en":     "English",
	"id":     "bahasa Indonesia",
	"gog_BO": "Botulo",
	"ca":     "català",
	"cz":     "český",
	"da":     "Dansk",
	"gog_DE": "Derpish",
	"de":     "Deutsch",
	"et":     "eesti",
	"es":     "español",
	"es_mx":  "Español (AL)",
	"eu":     "euskara",
	"fr":     "français",
	"hr":     "hrvatski",
	"gog_IN": "Inuktitut",
	"is":     "Íslenska",
	"it":     "italiano",
	"la":     "latine",
	"lt":     "lietuvių kalba",
	"hu":     "magyar",
	"nl":     "nederlands",
	"no":     "norsk",
	"pl":     "polski",
	"pt":     "português",
	"br":     "Português do Brasil",
	"ro":     "română",
	"sk":     "slovenský",
	"fi":     "suomi",
	"sv":     "svenska",
	"vi":     "Tiếng Việt",
	"tr":     "Türkçe",
	"uk":     "yкраїнська",
	"gk":     "Ελληνικά",
	"be":     "беларуская",
	"bl":     "български",
	"ru":     "русский",
	"sb":     "Српска",
	"he":     "עברית",
	"ar":     "العربية",
	"ms":     "بهاس ملايو",
	"fa":     "فارسی",
	"th":     "ไทย",
	"ko":     "한국어",
	"zh":     "中文",
	"cn":     "中文(简体)",
	"jp":     "日本語",
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
	return maps.Keys(nativeLanguageNames)
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
