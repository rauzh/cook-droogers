package models

type UserType int

const (
	NonMemberUser UserType = iota
	ManagerUser   UserType = iota
	ArtistUser    UserType = iota
)

type User struct {
	UserID   uint64   `json:"user_id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Type     UserType `json:"type"`
}
