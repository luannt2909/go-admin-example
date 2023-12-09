package enum

type UserRole int

const (
	RoleAdmin UserRole = iota + 1
	RoleGuest
)

var userRoleMap = map[UserRole]string{
	RoleAdmin: "admin",
	RoleGuest: "guest",
}

func (r UserRole) String() string {
	return userRoleMap[r]
}
