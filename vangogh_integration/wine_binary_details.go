package vangogh_integration

type WineBinaryDetails struct {
	Title    string          `json:"title"`
	OS       OperatingSystem `json:"os"`
	Version  string          `json:"version"`
	Filename string          `json:"filename"`
}
