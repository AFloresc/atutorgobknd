package rbac

//Role models the roles inside app
type Role string

const (
	// RoleLogin privileges, granted after account creation
	RoleLogin Role = "login"
	// RoleAdmin Administrative user, has access lesson statistics.
	RoleAdmin Role = "admin"
)
