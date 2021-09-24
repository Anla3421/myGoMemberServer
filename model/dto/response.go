package dto

type Response struct {
	Account  string  `json:"account,omitempty"`
	Password string  `json:"password,omitempty"`
	JWT      *string `json:"jwt,omitempty"`
	Log      int     `json:"log,omitempty"`
}
