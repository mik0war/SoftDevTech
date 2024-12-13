package types

type ProductCost struct {
	CostId         uint `gorm:"primaryKey;autoIncrement"`
	Value          float32
	ProductId      uint
	StartTimeStamp string
	Currency       string
}
