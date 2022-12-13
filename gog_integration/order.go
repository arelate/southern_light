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
	// TODO: find non-nil data to infer type
	Distributor        interface{} `json:"distributor"`
	Date               int         `json:"date"`
	MoneybackGuarantee bool        `json:"moneybackGuarantee"`
	Status             string      `json:"status"`
	PaymentMethod      string      `json:"paymentMethod"`
	// TODO: find non-nil data to infer type
	ValidUntil      interface{} `json:"validUntil"`
	CheckoutLink    string      `json:"checkoutLink"`
	ReceiptLink     string      `json:"receiptLink"`
	Total           value       `json:"total"`
	StoreCreditUsed value       `json:"storeCreditUsed"`
	// TODO: find non-nil data to infer type
	GiftRecipient interface{} `json:"giftRecipient"`
	// TODO: find non-nil data to infer type
	GiftSender interface{} `json:"giftSender"`
	Products   []struct {
		// TODO: find non-nil data to infer type
		Status interface{} `json:"status"`
		// TODO: find non-nil data to infer type
		RelatedAccount interface{} `json:"relatedAccount"`
		Price          struct {
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
		// TODO: find non-nil data to infer type
		RefundDate interface{} `json:"refundDate"`
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
