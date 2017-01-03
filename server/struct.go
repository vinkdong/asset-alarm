package server

type Credit struct {
	name           string
	icon           string
	amount         float64
	debit          float64
	balance        float64
	account_date   int8
	repayment_date int8
}
