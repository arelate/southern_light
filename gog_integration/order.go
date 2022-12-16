package gog_integration

const (
	OrderStatusGiftByMe  = "gift_by_me"
	OrderStatusGiftForMe = "gift_for_me"
	OrderStatusPurchase  = "purchase"
	OrderStatusRefund    = "refund"
)

type value struct {
	Amount          string `json:"amount"`
	Symbol          string `json:"symbol"`
	Code            string `json:"code"`
	IsZero          bool   `json:"isZero"`
	RawAmount       int    `json:"rawAmount"`
	FormattedAmount string `json:"formattedAmount"`
	Full            string `json:"full"`
	ForEmail        string `json:"for_email"`
}

type Order struct {
	PublicId string `json:"publicId"`
	//Distributor        interface{} `json:"distributor"`
	Date               int    `json:"date"`
	MoneybackGuarantee bool   `json:"moneybackGuarantee"`
	Status             string `json:"status"`
	PaymentMethod      string `json:"paymentMethod"`
	//ValidUntil      interface{} `json:"validUntil"`
	CheckoutLink    string `json:"checkoutLink"`
	ReceiptLink     string `json:"receiptLink"`
	Total           value  `json:"total"`
	StoreCreditUsed value  `json:"storeCreditUsed"`
	//GiftRecipient interface{} `json:"giftRecipient"`
	//GiftSender interface{} `json:"giftSender"`
	Products []struct {
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
	} `json:"products"`
	GiftCode      interface{} `json:"giftCode"`
	IsResendable  bool        `json:"isResendable"`
	StatusPageUrl string      `json:"statusPageUrl"`
}

func (ord *Order) GetId() int {
	//orders in theory are timestamped and you can't have two orders in the same moment,
	//so this might work great for id and holds more semantic value than PublicId
	return ord.Date
}
