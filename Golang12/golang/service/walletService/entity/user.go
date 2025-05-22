package entity

type User struct {
	ID   int    `gorm:"primaryKey"`
	Name string 
}
type UserResponse struct {
	UserID	string
	Status   string  
	Name string 
	AccountNumber string 
}