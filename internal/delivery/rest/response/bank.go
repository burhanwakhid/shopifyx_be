package response

type BankResponse struct {
	Id            string `json:"bankAccountId"`
	AccountName   string `json:"bankAccountName"`
	AccountNumber string `json:"bankAccountNumber"`
	Name          string `json:"bankName"`
}
