package gog_integration

// TotalPagesProxy allow unmarshalling
// account, orders, store, wishlist pages (.TotalPages) and
// catalog pages (.Pages) structs to query total pages
// with GetTotalPages method
type TotalPagesProxy struct {
	Pages      int              `json:"pages"`
	TotalPages int              `json:"totalPages"`
	Pagination *TotalPagesProxy `json:"pagination"`
}

func (tpp *TotalPagesProxy) GetTotalPages() int {
	if tpp.Pagination != nil {
		return tpp.Pagination.GetTotalPages()
	}
	if tpp.TotalPages > 0 {
		return tpp.TotalPages
	}
	return tpp.Pages
}
