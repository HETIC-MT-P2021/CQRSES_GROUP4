package user

const (
	OPERATOR     Role = 0x1
	ADMIN        Role = 0x1 << 1
	SYSTEM_ADMIN Role = 0x1 << 2
)

func (r Role) IsOperator() bool {
	return r&OPERATOR != 0
}

func (r Role) IsAdmin() bool {
	return r&ADMIN != 0
}

func (r Role) IsSystemAdmin() bool {
	return r&SYSTEM_ADMIN != 0
}