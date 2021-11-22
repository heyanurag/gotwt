package twitter

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Meta struct {
	Result_count int `json:"result_count"`
}

type Response struct {
	Data []User `json:"data"`
	Meta Meta   `json:"meta"`
}
