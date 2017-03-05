package server

type Bill struct {
	Id       int64
	CreditId int64
	Year     int
	Month    int
	Day      int
	Amount   float64
	Balance  float64 `show total balance`
	Credit   float64 `This month of credit`
}

func (b *Bill) Save() {
	CommonSave(b,"bill")

}
