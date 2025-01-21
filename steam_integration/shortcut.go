package steam_integration

import (
	"errors"
	"github.com/arelate/southern_light/steam_vdf"
	"hash/crc32"
	"slices"
	"strconv"
	"strings"
)

type Shortcut struct {
	AppId               uint32
	AppName             string
	Exe                 string
	StartDir            string
	Icon                string
	ShortcutPath        string
	LaunchOptions       string
	IsHidden            uint32
	AllowDesktopConfig  uint32
	AllowOverlay        uint32
	OpenVR              uint32
	Devkit              uint32
	DevkitGameId        string
	DevkitOverrideAppId string
	LastPlayTime        uint32
	FlatpakAppId        string
	Tags                []string
}

func (s *Shortcut) KeyValues(index string) *steam_vdf.KeyValues {
	kv := &steam_vdf.KeyValues{
		Key:    index,
		Type:   steam_vdf.BinaryTypeNone,
		Values: make([]*steam_vdf.KeyValues, 0),
	}

	kv.Values = append(kv.Values, s.AppIdKeyValue(), s.AppNameKeyValue(), s.ExeKeyValue())

	kv.Values = append(kv.Values, s.StartDirKeyValue())
	kv.Values = append(kv.Values, s.IconKeyValue())
	kv.Values = append(kv.Values, s.ShortcutPathKeyValue())
	kv.Values = append(kv.Values, s.LaunchOptionsKeyValues())
	kv.Values = append(kv.Values, s.IsHiddenKeyValues())
	kv.Values = append(kv.Values, s.AllowDesktopConfigKeyValues())
	kv.Values = append(kv.Values, s.AllowOverlayKeyValues())
	kv.Values = append(kv.Values, s.OpenVRKeyValues())
	kv.Values = append(kv.Values, s.DevkitKeyValues())
	kv.Values = append(kv.Values, s.DevkitGameIdKeyValues())
	kv.Values = append(kv.Values, s.DevkitOverrideAppIdKeyValues())
	kv.Values = append(kv.Values, s.LastPlayTimeKeyValues())
	kv.Values = append(kv.Values, s.FlatpakAppIdKeyValues())
	kv.Values = append(kv.Values, s.TagsKeyValues())

	return kv
}

func (s *Shortcut) AppIdKeyValue() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "appid",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.AppId,
	}
}

func (s *Shortcut) AppNameKeyValue() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "appname",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.AppName,
	}
}

func (s *Shortcut) ExeKeyValue() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "Exe",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: strings.Join([]string{"\"", s.Exe, "\""}, ""),
	}
}

func (s *Shortcut) StartDirKeyValue() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "StartDir",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: strings.Join([]string{"\"", s.StartDir, "\""}, "")}
}

func (s *Shortcut) IconKeyValue() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "icon",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.Icon,
	}
}

func (s *Shortcut) ShortcutPathKeyValue() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "ShortcutPath",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.ShortcutPath,
	}
}

func (s *Shortcut) LaunchOptionsKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "LaunchOptions",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.LaunchOptions,
	}
}

func (s *Shortcut) IsHiddenKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "IsHidden",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.IsHidden,
	}
}

func (s *Shortcut) AllowDesktopConfigKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "AllowDesktopConfig",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.AllowDesktopConfig,
	}
}

func (s *Shortcut) AllowOverlayKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "AllowOverlay",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.AllowOverlay,
	}
}

func (s *Shortcut) OpenVRKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "OpenVR",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.OpenVR,
	}
}

func (s *Shortcut) DevkitKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "Devkit",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.Devkit,
	}
}

func (s *Shortcut) DevkitGameIdKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "DevkitGameId",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.DevkitGameId,
	}
}

func (s *Shortcut) DevkitOverrideAppIdKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "DevkitOverrideAppId",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.DevkitOverrideAppId,
	}
}

func (s *Shortcut) LastPlayTimeKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "LastPlayTime",
		Type:       steam_vdf.BinaryTypeInt,
		TypedValue: s.LastPlayTime,
	}
}

func (s *Shortcut) TagsKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "tags",
		Type:       steam_vdf.BinaryTypeNone,
		TypedValue: nil,
	}
}

func (s *Shortcut) FlatpakAppIdKeyValues() *steam_vdf.KeyValues {
	return &steam_vdf.KeyValues{
		Key:        "FlatpakAppId",
		Type:       steam_vdf.BinaryTypeString,
		TypedValue: s.FlatpakAppId,
	}
}

func NewShortcut() *Shortcut {
	return &Shortcut{
		AppId:               0,
		AppName:             "",
		Exe:                 "",
		StartDir:            "",
		Icon:                "",
		ShortcutPath:        "",
		LaunchOptions:       "",
		IsHidden:            0,
		AllowDesktopConfig:  1,
		AllowOverlay:        1,
		OpenVR:              0,
		Devkit:              0,
		DevkitGameId:        "",
		DevkitOverrideAppId: "",
		LastPlayTime:        0,
		FlatpakAppId:        "",
		Tags:                nil,
	}
}

func GetShortcutByAppId(kvShortcuts *steam_vdf.KeyValues, appId uint32) *steam_vdf.KeyValues {
	for _, shortcut := range kvShortcuts.Values {
		for _, kv := range shortcut.Values {
			if kv.Key == "appid" && kv.Type == steam_vdf.BinaryTypeInt {
				if kv.TypedValue.(uint32) == appId {
					return shortcut
				}
			}
		}
	}
	return nil
}

func AppendShortcut(kvShortcuts *steam_vdf.KeyValues, shortcut *Shortcut) error {

	index := "0"
	if len(kvShortcuts.Values) > 0 {
		indexStr := kvShortcuts.Values[len(kvShortcuts.Values)-1].Key
		if indexInt, err := strconv.ParseInt(indexStr, 10, 32); err == nil {
			index = strconv.FormatInt(indexInt+1, 10)
		} else {
			return err
		}
	}

	kvShortcuts.Values = append(kvShortcuts.Values, shortcut.KeyValues(index))

	return nil
}

func UpdateShortcut(index string, kvShortcuts *steam_vdf.KeyValues, shortcut *Shortcut) error {

	for ii, kv := range kvShortcuts.Values {
		if kv.Key == index {
			kvShortcuts.Values[ii] = shortcut.KeyValues(index)
			return nil
		}
	}

	return errors.New("shortcuts.vdf does not have shortcut index " + index)

}

func ShortcutAppId(appName string) uint32 {
	key := appName
	crc := uint64(crc32.ChecksumIEEE([]byte(key)))
	crc = crc | 0x80000000
	crc = (crc << 32) | 0x02000000
	result := (crc >> 32) | 0x100000000
	return uint32(result)
}

func RemoveShortcuts(kvShortcuts *steam_vdf.KeyValues, appIds ...uint32) error {

	if len(kvShortcuts.Values) == 0 {
		return nil
	}

	expCap := len(kvShortcuts.Values)
	if len(kvShortcuts.Values) > len(appIds) {
		expCap = len(kvShortcuts.Values) - len(appIds)
	}

	filteredKeyValues := make([]*steam_vdf.KeyValues, 0, expCap)
	for _, shortcut := range kvShortcuts.Values {
		remove := false
		for _, kv := range shortcut.Values {
			if kv.Key == "appid" && kv.Type == steam_vdf.BinaryTypeInt {
				shortcutAppId := kv.TypedValue.(uint32)
				if slices.Contains(appIds, shortcutAppId) {
					remove = true
					break
				}
			}
		}
		if remove {
			continue
		}
		filteredKeyValues = append(filteredKeyValues, shortcut)
	}

	kvShortcuts.Values = filteredKeyValues

	return nil
}
