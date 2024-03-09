package main

import (
	"InternshipProject/backend/api"
	"InternshipProject/backend/components"
	"InternshipProject/backend/types"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

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

	router.GET("/user", func(c *gin.Context) {

		user, _ := loadUsers(db)
		// []types.User{
		// 	{
		// 		ID:       "1",
		// 		Username: "ehdyr",
		// 		Email:    "eDD@setif-univ.com",
		// 		Password: "somepassword",
		// 		Role:     "1",
		// 	},
		// 	{
		// 		ID:       "2",
		// 		Username: "ehdyr",
		// 		Email:    "eDD@setif-univ.com",
		// 		Password: "somepassword",
		// 		Role:     "1",
		// 	},
		// 	{
		// 		ID:       "3",
		// 		Username: "ehdyr",
		// 		Email:    "eDD@setif-univ.com",
		// 		Password: "somepassword",
		// 		Role:     "1",
		// 	},
		// 	{
		// 		ID:       "4",
		// 		Username: "ehdyr",
		// 		Email:    "eDD@setif-univ.com",
		// 		Password: "somepassword",
		// 		Role:     "1",
		// 	},
		// 	{
		// 		ID:       "5",
		// 		Username: "ehdyr",
		// 		Email:    "eDD@setif-univ.com",
		// 		Password: "somepassword",
		// 		Role:     "1",
		// 	},
		// }
		ticket := []types.Ticket{
			{
				ID:       "1",
				Content:  "hello there, this is the content of the ticket",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the first ticket",
			},
			{
				ID:       "2",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 2nd ticket",
			},
			{
				ID:       "3",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 3rd ticket",
			},
			{
				ID:       "4",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 4th ticket",
			},
			{
				ID:       "5",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 2nd ticket",
			},
			{
				ID:       "6",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 3rd ticket",
			},
			{
				ID:       "7",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 4th ticket",
			},
			{
				ID:       "2",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 2nd ticket",
			},
			{
				ID:       "3",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 3rd ticket",
			},
			{
				ID:       "4",
				Content:  "wassup",
				Owner:    "edd",
				Solution: "some solution",
				Title:    "the 4th ticket",
			},
		}
		api.Render(c, components.UserPage(user, ticket))
	})

	router.POST("/tickets", AddTicket)
	router.GET("/tickets", GetTickets)

	router.GET("/permissions", GetPermissions)
	router.GET("/roles", GetRoles)
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
	var permissions []types.Permission
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

func GetRoles(c *gin.Context) {
	roles, err := loadRoles(db)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, roles)
}

func AddUser(c *gin.Context) {
	var user types.User

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

func loadRoles(db *sql.DB) ([]types.Role, error) {
	roles := []types.Role{}

	rows, err := db.Query("SELECT * FROM Roles ;")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer db.Close()
	for rows.Next() {
		var ROLE types.Role
		if err := rows.Scan(&ROLE.ID, &ROLE.Role); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		roles = append(roles, ROLE)
	}
	return roles, nil
}

func AddTicket(c *gin.Context) {
	var ticket types.Ticket

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
	var solution types.Solution
	if err := c.BindJSON(&solution); err != nil {
		return
	}

	if _, err := insertSolution(solution); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, solution)
}

func AddRole(c *gin.Context) {
	var role types.Role
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
	var permission types.Permission
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

func insertPermission(permission types.Permission) (int64, error) {
	result, err := db.Exec("INSERT INTO Permissions (permission) VALUES (?)", permission.Permission)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}

func insertRole(role types.Role) (int64, error) {
	result, err := db.Exec("INSERT INTO Roles (role) VALUES (?)", role.Role)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}
func insertSolution(solution types.Solution) (int64, error) {
	result, err := db.Exec("INSERT INTO Solutions (owner_id, ticket_id, solution) VALUES(?,?,?)", solution.Owner, solution.Ticket, solution.Solution)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}
func insertTicket(ticket types.Ticket) (int64, error) {
	result, err := db.Exec("INSERT INTO Tickets (content, owner_id, title) VALUES (?,?,?)", ticket.Content, ticket.Owner, ticket.Title)
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	return id, nil
}

func insertUser(user types.User) (int64, error) {
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
func loadUsers(db *sql.DB) ([]types.User, error) {
	Users := []types.User{}
	rows, err := db.Query("SELECT * FROM Users ;")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user types.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		Users = append(Users, user)
	}
	return Users, nil
}

func loadPermissions(db *sql.DB) ([]types.Permission, error) {
	permissions := []types.Permission{}
	rows, err := db.Query("SELECT * FROM Permissions :")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var perm types.Permission
		if err := rows.Scan(&perm.ID, &perm.Permission); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		permissions = append(permissions, perm)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	return permissions, nil
}

func loadTickets(db *sql.DB) ([]types.Ticket, error) {
	tickets := []types.Ticket{}
	rows, err := db.Query("SELECT * FROM Tickets ;")
	if err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var ticket types.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Content, &ticket.Owner, &ticket.Solution); err != nil {
			return nil, fmt.Errorf("error : %v", err)
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error : %v", err)
	}
	return tickets, nil
}

func getRole(user types.User, db *sql.DB) (types.Role, error) {
	var user_role types.Role
	row := db.QueryRow("SELECT Roles.role FROM Roles INNER JOIN Users ON Users.role = Roles.ID WHERE Users.ID = ? ;", user.ID)

	if err := row.Scan(&user_role.ID, &user_role.Role); err != nil {
		if err == sql.ErrNoRows {
			return user_role, fmt.Errorf("error : %v", err)
		}
		return user_role, fmt.Errorf("error : %v", err)
	}
	return user_role, nil
}
