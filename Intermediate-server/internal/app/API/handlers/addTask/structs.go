package addtask

type request struct {
	API_ID      string   `json:"api_id"`
	API_HASH    string   `json:"api_hash"`
	PhoneNumber string   `json:"phone_number"`
	Chats       []string `json:"chats"`
}
