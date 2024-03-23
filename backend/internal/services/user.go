package service

type UserType int

const (
	NewUser       UserType = iota
	NonMemberUser UserType = iota
	ManagerUser   UserType = iota
	ArtistUser    UserType = iota
)

type User struct {
	UserID   uint64
	Name     string
	Email    string
	Password string
	Type     UserType
}
