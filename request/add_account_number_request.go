package request

type AddAccountNumber struct {
	AccNumber  string `json:"account_number"`
	AccBalance int    `json:"account_balance"`
}
