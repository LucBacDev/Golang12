package entity

type Product struct {
	ID   int  `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
}