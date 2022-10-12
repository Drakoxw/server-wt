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

type Login struct {
	Email string `bson:"email,omitempty"`
	Pass  string `bson:"pass,omitempty"`
}
