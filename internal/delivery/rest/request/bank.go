package request

type BankRequest struct {
	BankName          string `json:"bankName" validate:"required,string,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" validate:"required,string,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,string,min=5,max=15"`
}

// {
// 	"bankName":"", // not null, minLength 5, maxLength 15
// 	"bankAccountName":"", // not null, minLength 5, maxLength 15
// 	"bankAccountNumber":"" // not null, minLength 5, maxLength 15
// }
