package gog_integration

type OrderProduct struct {
	//Status interface{} `json:"status"`
	//RelatedAccount interface{} `json:"relatedAccount"`
	Price struct {
		BaseAmount   string `json:"baseAmount"`
		Amount       string `json:"amount"`
		IsFree       bool   `json:"isFree"`
		IsDiscounted bool   `json:"isDiscounted"`
		Symbol       string `json:"symbol"`
	} `json:"price"`
	Image                      string `json:"image"`
	Title                      string `json:"title"`
	Id                         string `json:"id"`
	IsRefunded                 bool   `json:"isRefunded"`
	CashValue                  value  `json:"cashValue"`
	WalletValue                value  `json:"walletValue"`
	IsPreorder                 bool   `json:"isPreorder"`
	DisplayAutomaticRefundLink bool   `json:"displayAutomaticRefundLink"`
	//RefundDate interface{} `json:"refundDate"`
}

func (op *OrderProduct) GetImage() string {
	return op.Image
}
