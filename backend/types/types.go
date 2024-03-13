package types

type User struct {
	ID       string `form:"ID"`
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Role     string `form:"role"`
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
