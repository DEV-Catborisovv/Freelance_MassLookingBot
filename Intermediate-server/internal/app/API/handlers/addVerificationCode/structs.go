package addverificationcode

type request struct {
	Phone_number string `json:"phone_number"`
	Code         string `json:"code"`
}
