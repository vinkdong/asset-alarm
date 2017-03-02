package server

type Bill struct {
	Id       int64
	CreditId int64
	Year     int
	Month    int
	Amount   float64
	Balance  float64
}

func (b *Bill) Save() {
	CommonSave(b,"bill")

}
