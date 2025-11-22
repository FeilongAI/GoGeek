package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
}
type EditUserInfo struct {
	Nickname    string
	Birthday    time.Time
	Description string
}
