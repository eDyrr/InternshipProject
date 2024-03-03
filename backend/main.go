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
	router.GET("/users", GetUsers)
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

func loadUsers(db *sql.DB) ([]User, error) {
	Users := []User{}
	rows, err := db.Query("SELECT * FROM Users ;")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
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
