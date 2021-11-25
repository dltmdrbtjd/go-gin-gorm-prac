package models

type Book struct {
	ID       uint   `json:"id" gorm:"primary+key"`
	Title    string `json:"title"`
	Location string `json:"location"`
	Explain  string `json:"explain"`
}
