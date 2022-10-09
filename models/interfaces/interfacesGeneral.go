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
}
