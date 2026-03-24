package vangogh_integration

type AvailableProduct struct {
	Id               int               `json:"id"`
	Title            string            `json:"tt"`
	OperatingSystems []OperatingSystem `json:"os"`
}
