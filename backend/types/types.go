package types

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Permission struct {
	ID         string `json:"ID"`
	Permission string `json:"permission"`
}
type Ticket struct {
	ID       string `json:"ID"`
	Content  string `json:"content"`
	Owner    string `json:"owner_id"`
	Solution string `json:"solution_id"`
	Title    string `json:"title"`
}

type Solution struct {
	ID       string `json:"ID"`
	Owner    string `json:"owner_id"`
	Ticket   string `json:"ticket_id"`
	Solution string `json:"solution"`
}

type Role struct {
	ID   string `json:"ID"`
	Role string `json:"role"`
}
