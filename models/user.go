package models

type Contact struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"` //Пользователь, которому принадлежит этот контакт
}

func GetContact(id uint) *Contact {
	contact := &Contact{"Ivan", "340987534", 777}
	return contact
}
