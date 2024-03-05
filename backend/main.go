package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Permission struct {
	ID         string `json:"ID"`
	permission string `json:"permission"`
}
type Ticket struct {
	ID       string `json:"ID"`
	content  string `json:"content"`
	owner    string `json:"owner_id"`
	solution string `json:"solution_id"`
	title    string `json:"title"`
}

type Solution struct {
	ID       string `json:"ID"`
	owner    string `json:"owner_id"`
	ticket   string `json:"ticket_id"`
	solution string `json:"solution"`
}

type Role struct {
	ID   string `json:"ID"`
	role string `json:"role"`
}

var db *sql.DB

func main() {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "eDyrr7355608",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "InternShip",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Print("connection success!\n")

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/users", AddUser)
	router.GET("/users", GetUsers)

	router.POST("/tickets", AddTicket)
	router.Run("localhost:8080")

}

func GetUsers(c *gin.Context) {
	users, err := loadUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetPermissions(c *gin.Context) {
	var permissions []Permission
	var err error

	permissions, err = loadPermissions(db)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, permissions)
}

func GetTickets(c *gin.Context) {
	tickets, err := loadTickets(db)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, tickets)
}

func AddUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		return
	}
	id, err := insertUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.ID = fmt.Sprintf("%d", id)
	c.JSON(http.StatusCreated, user)
}

func AddTicket(c *gin.Context) {
	var ticket Ticket

	if err := c.BindJSON(&ticket); err != nil {
		return
	}

	_, err := insertTicket(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, ticket)
}

func AddSolution(c *gin.Context) {
	var solution Solution
	if err := c.BindJSON(&solution); err != nil {
		return
	}

	if _, err := insertSolution(solution); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, solution)
}

func AddRole(c *gin.Context) {
	var role Role
	if err := c.BindJSON(&role); err != nil {
		return
	}
	_, err := insertRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, role)
}

func AddPermission(c *gin.Context) {
	var permission Permission
	if err := c.BindJSON(&permission); err != nil {
		return
	}
	_, err := insertPermission(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, permission)
}

func insertPermission(permission Permission) (int64, error) {
	result, err := db.Exec("INSERT INTO Permissions (permission) VALUES (?)", permission.permission)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}

func insertRole(role Role) (int64, error) {
	result, err := db.Exec("INSERT INTO Roles (role) VALUES (?)", role.role)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}
func insertSolution(solution Solution) (int64, error) {
	result, err := db.Exec("INSERT INTO Solutions (owner_id, ticket_id, solution) VALUES(?,?,?)", solution.owner, solution.ticket, solution.solution)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}
func insertTicket(ticket Ticket) (int64, error) {
	result, err := db.Exec("INSERT INTO Tickets (content, owner_id, title) VALUES (?,?,?)", ticket.content, ticket.owner, ticket.title)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}

func insertUser(user User) (int64, error) {
	result, err := db.Exec("INSERT INTO Users (username, email, password, role) VALUES (?, ?, ?, ?)", user.Username, user.Email, user.Password, user.Role)

	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}
func loadUsers(db *sql.DB) ([]User, error) {
	Users := []User{}
	rows, err := db.Query("SELECT * FROM Users ;")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		Users = append(Users, user)
	}
	return Users, nil
}

func loadPermissions(db *sql.DB) ([]Permission, error) {
	permissions := []Permission{}
	rows, err := db.Query("SELECT * FROM Permissions :")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var perm Permission
		if err := rows.Scan(&perm.ID, &perm.permission); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		permissions = append(permissions, perm)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	return permissions, nil
}

func loadTickets(db *sql.DB) ([]Ticket, error) {
	tickets := []Ticket{}
	rows, err := db.Query("SELECT * FROM Tickets ;")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var ticket Ticket
		if err := rows.Scan(&ticket.ID, &ticket.content, &ticket.owner, &ticket.solution); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	return tickets, nil
}
