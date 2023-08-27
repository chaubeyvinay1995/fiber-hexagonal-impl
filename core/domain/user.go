package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

//func NewPerson(id int, email string, password string) *User {
//	return &User{
//		Email:    email,
//		Password: password,
//	}
//}

func (u *User) GetEmail() string {
	return u.Email
}
