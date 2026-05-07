package egs_integration

import "strconv"

type JsonManifest struct {
	ManifestFileVersion string   `json:"ManifestFileVersion"`
	IsFileData          bool     `json:"bIsFileData"`
	AppId               string   `json:"AppID"`
	AppName             string   `json:"AppNameString"`
	BuildVersion        string   `json:"BuildVersionString"`
	LaunchExe           string   `json:"LaunchExeString"`
	LaunchCommand       string   `json:"LaunchCommand"`
	PrereqIds           []string `json:"PrereqIds"`
	PrereqName          string   `json:"PrereqName"`
	PrereqPath          string   `json:"PrereqPath"`
	PrereqArgs          string   `json:"PrereqArgs"`
	FileManifestList    []struct {
		Filename       string `json:"Filename"`
		FileHash       string `json:"FileHash"`
		FileChunkParts []struct {
			Guid   string `json:"Guid"`
			Offset string `json:"Offset"`
			Size   string `json:"Size"`
		} `json:"FileChunkParts"`
	} `json:"FileManifestList"`
	ChunkHashList     map[string]string `json:"ChunkHashList"`
	ChunkShaList      map[string]string `json:"ChunkShaList"`
	DataGroupList     map[string]string `json:"DataGroupList"`
	ChunkFilesizeList map[string]string `json:"ChunkFilesizeList"`
	CustomFields      map[string]string `json:"CustomFields"`
}

func parseJsonInt(ns string) (int, error) {
	var ni, sh int

	for i := 0; i < len(ns); i += 3 {
		part := ns[i : i+3]

		if val, err := strconv.ParseInt(part, 10, 32); err == nil {
			ni += int(val) << sh
			sh += 8
		} else {
			return 0, err
		}
	}

	return ni, nil
}

func (jm *JsonManifest) Manifest() (*Manifest, error) {

	manifest := new(Manifest)

	var err error

	if manifest.Metadata, err = jm.metadata(); err != nil {
		return nil, err
	}

	manifest.CustomFields = jm.customFields()

	return manifest, nil
}

func (jm *JsonManifest) metadata() (*Metadata, error) {

	var err error

	var fml int
	if fml, err = parseJsonInt(jm.ManifestFileVersion); err != nil {
		return nil, err
	}

	var appId int
	if appId, err = parseJsonInt(jm.AppId); err != nil {
		return nil, err
	}

	return new(Metadata{
		FeatureLevel:  uint32(fml),
		IsFileData:    jm.IsFileData,
		AppId:         uint32(appId),
		AppName:       jm.AppName,
		BuildVersion:  jm.BuildVersion,
		LaunchExe:     jm.LaunchExe,
		LaunchCommand: jm.LaunchCommand,
		PrereqIds:     jm.PrereqIds,
		PrereqName:    jm.PrereqName,
		PrereqPath:    jm.PrereqPath,
		PrereqArgs:    jm.PrereqArgs,
	}), nil
}

func (jm *JsonManifest) customFields() *CustomFields {
	return new(CustomFields{
		Count:  uint32(len(jm.CustomFields)),
		Fields: jm.CustomFields,
	})
}
