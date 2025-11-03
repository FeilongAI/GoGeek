package domain

import "time"

type User struct {
	Email    string
	Password string
}
type EditUserInfo struct {
	Nickname    string
	Birthday    time.Time
	Description string
}
