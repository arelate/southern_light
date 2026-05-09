package egs_integration

import (
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const hashLen = 20

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
		Filename         string `json:"Filename"`
		FileHash         string `json:"FileHash"`
		IsReadOnly       bool   `json:"bIsReadOnly"`
		IsCompressed     bool   `json:"bIsCompressed"`
		IsUnixExecutable bool   `json:"bIsUnixExecutable"`
		FileChunkParts   []struct {
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

func parseJsonInt(inStr string) (*big.Int, error) {
	num := new(big.Int)
	shift := 0

	for i := 0; i < len(inStr); i += 3 {
		end := i + 3
		if end > len(inStr) {
			end = len(inStr)
		}
		val, err := strconv.ParseInt(inStr[i:end], 10, 16)
		if err != nil {
			return nil, err
		}
		num.Add(num, new(big.Int).Lsh(big.NewInt(val), uint(shift)))
		shift += 8
	}
	return num, nil
}

func parseJsonHash(hash string) ([]byte, error) {

	num, err := parseJsonInt(hash)
	if err != nil {
		return nil, err
	}

	result := make([]byte, hashLen)
	numBytes := num.Bytes()

	for i := 0; i < len(numBytes); i++ {
		result[hashLen-1-i] = numBytes[i]
	}

	return result, nil
}

func (jm *JsonManifest) Manifest() (*Manifest, error) {

	manifest := new(Manifest)

	var err error

	if manifest.Metadata, err = jm.metadata(); err != nil {
		return nil, err
	}

	if manifest.FileList, err = jm.filelist(); err != nil {
		return nil, err
	}

	if manifest.ChunkList, err = jm.chunklist(manifest.FileList); err != nil {
		return nil, err
	}

	manifest.CustomFields = jm.customFields()

	return manifest, nil
}

func (jm *JsonManifest) metadata() (*Metadata, error) {

	var err error

	var fml *big.Int
	if fml, err = parseJsonInt(jm.ManifestFileVersion); err != nil {
		return nil, err
	}

	var appId *big.Int
	if appId, err = parseJsonInt(jm.AppId); err != nil {
		return nil, err
	}

	return new(Metadata{
		FeatureLevel:  uint32(fml.Uint64()),
		IsFileData:    jm.IsFileData,
		AppId:         uint32(appId.Uint64()),
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

func (jm *JsonManifest) filelist() (*FileList, error) {
	filelist := new(FileList)

	filelist.Count = uint32(len(jm.FileManifestList))

	for _, fml := range jm.FileManifestList {

		shaHash, err := parseJsonHash(fml.FileHash)
		if err != nil {
			return nil, err
		}

		var flags uint8
		if fml.IsReadOnly {
			flags |= 1
		}
		if fml.IsCompressed {
			flags |= 1 << 1
		}
		if fml.IsUnixExecutable {
			flags |= 1 << 2
		}

		file := File{
			Filename: fml.Filename,
			ShaHash:  shaHash,
		}

		for _, fcp := range fml.FileChunkParts {
			var partUuid uuid.UUID
			partUuid, err = uuid.Parse(fcp.Guid)
			if err != nil {
				return nil, err
			}

			var partOffset *big.Int
			partOffset, err = parseJsonInt(fcp.Offset)
			if err != nil {
				return nil, err
			}

			var partSize *big.Int
			partSize, err = parseJsonInt(fcp.Size)
			if err != nil {
				return nil, err
			}

			part := ChunkPart{
				ParentUuid: partUuid,
				Offset:     uint32(partOffset.Uint64()),
				Size:       uint32(partSize.Uint64()),
				Chunk: new(Chunk{
					Uuid: partUuid,
				}),
			}

			file.Parts = append(file.Parts, part)
		}

		filelist.List = append(filelist.List, file)
	}

	return filelist, nil
}

func (jm *JsonManifest) chunklist(filelist *FileList) (*ChunkList, error) {
	chunklist := new(ChunkList{
		Lookup: make(map[uuid.UUID]uint32),
		Count:  uint32(len(jm.ChunkFilesizeList)),
	})

	var index uint32

	for ii, file := range filelist.List {
		for _, part := range file.Parts {

			chunkUuid := uuidString(part.Chunk.Uuid)

			if groupStr, ok := jm.DataGroupList[chunkUuid]; ok {
				if groupInt, err := parseJsonInt(groupStr); err == nil {
					part.Chunk.Group = uint8(groupInt.Uint64())
				} else {
					return nil, err
				}
			}

			if hashStr, ok := jm.ChunkHashList[chunkUuid]; ok {
				if hashInt, err := parseJsonInt(hashStr); err == nil {
					part.Chunk.Hash = hashInt.Uint64()
				} else {
					return nil, err
				}
			}

			if shaHashStr, ok := jm.ChunkShaList[chunkUuid]; ok {

				if shaHash, err := hex.DecodeString(shaHashStr); err == nil {
					part.Chunk.ShaHash = shaHash
				} else {
					return nil, err
				}
			}

			if fileSizeStr, ok := jm.ChunkFilesizeList[chunkUuid]; ok {
				if fileSize, err := parseJsonInt(fileSizeStr); err == nil {
					part.Chunk.FileSize = fileSize.Uint64()
				} else {
					return nil, err
				}
			}

			chunklist.Chunks = append(chunklist.Chunks, part.Chunk)
			chunklist.Lookup[part.ParentUuid] = index
			index++

			filelist.List[ii].Size += part.Chunk.FileSize
		}
	}

	return chunklist, nil
}

func uuidString(u uuid.UUID) string {

	us := u.String()
	us = strings.ToUpper(us)
	us = strings.Replace(us, "-", "", -1)
	return us
}

func (jm *JsonManifest) customFields() *CustomFields {
	return new(CustomFields{
		Count:  uint32(len(jm.CustomFields)),
		Fields: jm.CustomFields,
	})
}
