package egs_integration

type JsonManifest struct {
	ManifestFileVersion string `json:"ManifestFileVersion"`
	IsFileData          bool   `json:"bIsFileData"`
	AppId               string `json:"AppID"`
	AppName             string `json:"AppNameString"`
	BuildVersion        string `json:"BuildVersionString"`
	LaunchExe           string `json:"LaunchExeString"`
	LaunchCommand       string `json:"LaunchCommand"`
	PrereqIds           []any  `json:"PrereqIds"`
	PrereqName          string `json:"PrereqName"`
	PrereqPath          string `json:"PrereqPath"`
	PrereqArgs          string `json:"PrereqArgs"`
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
