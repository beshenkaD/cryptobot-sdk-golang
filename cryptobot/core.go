package cryptobot

const (
	InvoicePaidStatus    = "paid"
	InvoiceActiveStatus  = "active"
	InvoiceExpiredStatus = "expired"

	BTC  = "BTC"
	TON  = "TON"
	ETH  = "ETH"
	USDT = "USDT"
	USDC = "USDC"
	BUSD = "BUSD"
	LTC  = "LTC"
	BNB  = "BNB"
	TRX  = "TRX"

	InvoiceViewItemPaidBtnName    = "viewItem"
	InvoiceOpenChannelPaidBtnName = "openChannel"
	InvoiceCallbackPaidBtnName    = "callback"

	TransferCompleteStatus = "completed"
)

type response struct {
	Ok     bool          `json:"ok"`
	Result any           `json:"result"`
	Error  responseError `json:"error"`
}

type responseError struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

type Invoice struct {
	// Unique ID for this invoice.
	ID int64 `json:"invoice_id"`

	// Status of the invoice, can be either “active”, “paid” or “expired”.
	Status string `json:"status"`

	// Hash of the invoice.
	Hash string `json:"hash"`

	// Currency code. Currently, can be “BTC”, “TON”, “ETH”, “USDT”, “USDC” or “BUSD”.
	Asset string `json:"asset"`

	// Amount of the invoice.
	Amount string `json:"amount"`

	// Optional. Amount of charged service fees. Returned only if the invoice has the paid status.
	Fee string `json:"fee"`

	// URL should be presented to the user to pay the invoice.
	PayURL string `json:"pay_url"`

	// Optional. Description for this invoice.
	Description string `json:"description"`

	// Date the invoice was created in ISO 8601 format.
	CreatedAt string `json:"created_at"`

	// Optional. Price of the asset in USD. Returned only if the invoice has the paid status.
	UsdRate string `json:"usd_rate"`

	// True, if the user can add comment to the payment.
	AllowComments bool `json:"allow_comments"`

	// True, if the user can pay the invoice anonymously.
	AllowAnonymous bool `json:"allow_anonymous"`

	// Optional. Date the invoice expires in Unix time.
	ExpirationDate string `json:"expiration_date"`

	// Optional. Date the invoice was paid in Unix time.
	PaidAt string `json:"paid_at"`

	// True, if the invoice was paid anonymously.
	PaidAnonymously bool `json:"paid_anonymously"`

	// Optional. Comment to the payment from the user.
	Comment string `json:"comment"`

	// Optional. Text of the hidden message for this invoice.
	HiddenMessage string `json:"hidden_message"`

	// Optional. Previously provided data for this invoice.
	Payload string `json:"payload"`

	// Optional. Name of the button, can be “viewItem”, “openChannel” or “callback”.
	PaidBtnName string `json:"paid_btn_name"`

	// Optional. URL of the button.
	PaidBtnURL string `json:"paid_btn_url"`
}

type Transfer struct {
	// Unique ID for this transfer.
	ID int64 `json:"transfer_id"`

	// Telegram user ID the transfer was sent to.
	UserID string `json:"user_id"`

	// Currency code. Currently, can be “BTC”, “TON”, “ETH”, “USDT”, “USDC” or “BUSD”.
	Asset string `json:"asset"`

	// Amount of the transfer.
	Amount string `json:"amount"`

	// Status of the transfer, can be “completed”.
	Status string `json:"status"`

	// Date the transfer was completed in ISO 8601 format.
	CompletedAt string `json:"completed_at"`

	// Optional. Comment for this transfer.
	Comment string `json:"comment"`
}

type Check struct {
	// Unique ID for this check.
	ID int64 `json:"check_id"`

	// Hash of the check.
	Hash string `json:"hash"`

	// Cryptocurrency alphabetic code. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Asset string `json:"asset"`

	// Amount of the check in float.
	Amount string `json:"amount"`

	// URL should be provided to the user to activate the check.
	BotCheckURL string `json:"bot_check_url"`

	// Status of the check, can be “active” or “activated”.
	Status string `json:"status"`

	// Date the check was created in ISO 8601 format.
	CreatedAt string `json:"created_at"`

	// Date the check was activated in ISO 8601 format.
	ActivatedAt string `json:"activated_at"`
}
