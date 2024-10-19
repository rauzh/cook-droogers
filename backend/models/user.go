package models

type UserType int

const (
	NonMemberUser UserType = iota
	ManagerUser   UserType = iota
	ArtistUser    UserType = iota
	AdminUser     UserType = iota
)

const (
	NonMemberUserStr string = "nonmember"
	ManagerUserStr   string = "manager"
	ArtistUserStr    string = "artist"
	AdminUserStr     string = "admin"
)

type User struct {
	UserID   uint64
	Name     string
	Email    string
	Password string
	Type     UserType
}

func UserTypeStrToEnum(t string) UserType {
	switch t {
	case "nonmember":
		return 0
	case "manager":
		return 1
	case "artist":
		return 2
	case "admin":
		return 3
	default:
		return 0
	}
}
