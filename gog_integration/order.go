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
	Date               int64  `json:"date"`
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
	Products      []OrderProduct `json:"products"`
	GiftCode      any            `json:"giftCode"`
	IsResendable  bool           `json:"isResendable"`
	StatusPageUrl string         `json:"statusPageUrl"`
}
