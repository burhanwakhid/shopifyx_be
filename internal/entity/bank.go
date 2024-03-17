package entity

type Bank struct {
	Id            string `db:"id"`
	AccountName   string `db:"account_name"`
	AccountNumber string `db:"account_number"`
	Name          string `db:"name"`
	IdUser        string `db:"id_user"`
}
