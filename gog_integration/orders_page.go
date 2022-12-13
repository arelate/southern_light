package gog_integration

type OrderPage struct {
	TotalPages int     `json:"totalPages"`
	Orders     []Order `json:"orders"`
}

func (op *OrderPage) GetTotalPages() int {
	return op.TotalPages
}

func (op *OrderPage) GetProducts() []IdGetter {
	idGetters := make([]IdGetter, 0)
	for _, ord := range op.Orders {
		ord := ord
		idGetters = append(idGetters, &ord)
	}
	return idGetters
}
