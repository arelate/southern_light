package egs_integration

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"uuid"

	"github.com/arelate/southern_light/vangogh_integration"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/nod"
)

const manifestMagic uint32 = 0x44BEC00C

const ManifestExt = ".manifest"

const (
	StorageUncompressed uint8 = 0x00
	StorageCompressed   uint8 = 0x01
	StorageEncrypted    uint8 = 0x02
)

type Manifest struct {
	Header       *Header
	Metadata     *Metadata
	ChunkList    *ChunkList
	FileList     *FileList
	CustomFields *CustomFields
}

type Header struct {
	Offset           uint32
	SizeUncompressed uint32
	SizeCompressed   uint32
	ShaHash          []byte
	Storage          uint8
	FeatureLevel     uint32
}

type Metadata struct {
	Offset        uint32
	Version       uint8
	FeatureLevel  uint32
	IsFileData    bool
	AppId         uint32
	AppName       string
	BuildVersion  string
	LaunchExe     string
	LaunchCommand string
	PrereqIds     []string
	PrereqName    string
	PrereqPath    string
	PrereqArgs    string
	BuildId       string
}

type ChunkList struct {
	Offset  uint32
	Version uint8
	Count   uint32
	Chunks  []*Chunk
	Lookup  map[uuid.UUID]uint32
}

type Chunk struct {
	Uuid       uuid.UUID
	Hash       uint64
	ShaHash    []byte
	Group      uint8
	WindowSize uint32
	FileSize   uint64
}

func chunkDir(featureLevel uint32) string {
	if featureLevel < 3 {
		return "Chunks"
	} else if featureLevel < 6 {
		return "ChunksV2"
	} else if featureLevel < 15 {
		return "ChunksV3"
	}
	return "ChunksV4"
}

func (chk *Chunk) Path(featureLevel uint32) string {
	base := fmt.Sprintf("%02d/%016X_%X.chunk", chk.Group, chk.Hash, chk.Uuid[:])
	return filepath.Join(chunkDir(featureLevel), base)
}

type FileList struct {
	Offset  uint32
	Version uint8
	Count   uint32
	List    []File
}

type File struct {
	Filename      string
	SymlinkTarget string
	ShaHash       []byte
	Flags         uint8
	InstallTags   []string
	Size          uint64
	Parts         []ChunkPart
}

type ChunkPart struct {
	DataSize   uint32
	ParentUuid uuid.UUID
	Offset     uint32
	Size       uint32
	Chunk      *Chunk
}

type CustomFields struct {
	Offset  uint32
	Version uint8
	Count   uint32
	Fields  map[string]string
}

func GetManifest(appName string, gameManifest *GameManifest, operatingSystem vangogh_integration.OperatingSystem, force bool) (*Manifest, error) {

	egma := nod.Begin("getting EGS manifest...")
	defer egma.Done()

	manifestsDir := vangogh_integration.AbsProductTypeDir(vangogh_integration.EgsManifests)

	kvManifests, err := kevlar.New(manifestsDir, ManifestExt)
	if err != nil {
		return nil, err
	}

	osAppNameKey := fmt.Sprintf("%s-%s", appName, operatingSystem)

	if !kvManifests.Has(osAppNameKey) || force {
		if err = FetchManifests(osAppNameKey, gameManifest, kvManifests); err != nil {
			return nil, err
		}
	}

	absManifestFilename := filepath.Join(manifestsDir, osAppNameKey+ManifestExt)

	manifestFile, err := os.Open(absManifestFilename)
	if err != nil {
		return nil, err
	}
	defer manifestFile.Close()

	return ReadManifest(manifestFile)
}

func FetchManifests(key string, gameManifest *GameManifest, kvManifests kevlar.KeyValues) error {

	manifestUrls, err := gameManifest.Urls()
	if err != nil {
		return err
	}

	client, err := GetClient()
	if err != nil {
		return err
	}

	var downloaded bool

	for _, manifestUrl := range manifestUrls {
		if err = FetchManifest(key, manifestUrl, client, kvManifests); err == nil {
			downloaded = true
			break
		}
	}

	if !downloaded {
		return errors.New("unable to successfully download at least one manifest")
	}

	return nil
}

func FetchManifest(key string, manifestUrl *url.URL, client *http.Client, kvManifests kevlar.KeyValues) error {

	req, err := http.NewRequest(http.MethodGet, manifestUrl.String(), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status)
	}

	return kvManifests.Set(key, resp.Body)
}
