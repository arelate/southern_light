package vangogh_integration

type AvailableProduct struct {
	Id    int               `json:"id"`
	Title string            `json:"tt"`
	Os    []OperatingSystem `json:"os"`
}
