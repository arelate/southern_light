package vangogh_integration

type AvailableProduct struct {
	Id               string            `json:"id"`
	Title            string            `json:"tt"`
	OperatingSystems []OperatingSystem `json:"os"`
	Dlc              map[string]string `json:"dlc,omitempty"`
}
