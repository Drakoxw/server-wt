package interfaces

type ResponseGen struct {
	Status  string
	Message string
	Data    []any
}

type IClienteInsert struct {
	Email    string `bson:"email,omitempty"`
	UserName string `bson:"userName,omitempty"`
	Pass     string `bson:"pass,omitempty"`
	Role     string `bson:"role,omitempty"`
	DateReg  string `bson:"dateReg,omitempty"`
}

type IProviderInsert struct {
	Email          string `bson:"email"`
	UserName       string `bson:"userName"`
	Pass           string `bson:"pass"`
	City           string `bson:"city"`
	NameEnterprise string `bson:"nameEnterprise"`
	Role           string `bson:"role"`
	PrivilegeLevel int16  `bson:"privilegeLevel"`
	DateReg        string `bson:"dateReg,omitempty"`
}

type Login struct {
	Email string `bson:"email,omitempty"`
	Pass  string `bson:"pass,omitempty"`
}
