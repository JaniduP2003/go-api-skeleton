package types

type RegisterUserPayload struct {
	FirstName string  `json:"firstname"`
	LastName string   `jason:"lastname"`
	Email string      `json:"email"`
	Password string   `json:"password"`
}  //this makes a payload 