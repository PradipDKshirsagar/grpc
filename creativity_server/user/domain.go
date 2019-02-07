package user

type User struct {
	ID        int    `json: "id"`
	Age       int    `json:"age"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Response struct {
	Output string `json:"Output"`
}
