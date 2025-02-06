package gog_integration

type OrderPage struct {
	TotalPages int     `json:"totalPages"`
	Orders     []Order `json:"orders"`
}
