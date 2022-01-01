package dto

type Query struct {
	Account  string
	Password string
	JWT      *string
}

type GetInfoByJWT struct {
	ID      int32
	Account string
}
