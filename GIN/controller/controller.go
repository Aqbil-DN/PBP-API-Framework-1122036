package controller

import (
	"log"
	"net/http"
	"strconv"

	m "GIN/model"

	"github.com/gin-gonic/gin"
)

func GinGetAllUsers(c *gin.Context) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM Users"
	rows, err := db.Query(query)
	if err != nil {
		SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer rows.Close()

	var users m.User
	//var user []m.User
	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Name, &users.Age, &users.Address, &users.Password, &users.Email); err != nil {
			SendErrorResponse(c, http.StatusInternalServerError, "Internal server error")
			log.Println(err)
			return
		}
	}

	// if len(user) == 0 {
	// 	SendErrorResponse(c, http.StatusNotFound, "User id not found")
	// 	return
	// }

	c.JSON(200, users)
}

func GinInsertNewUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	name := c.Query("name")
	ageStr := c.Query("age")
	address := c.Query("address")
	password := c.Query("password")
	email := c.Query("email")

	if name == "" || ageStr == "" || address == "" || password == "" || email == "" {
		SendErrorResponse(c, 400, "data ga lengkap")
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		SendErrorResponse(c, 400, "usia ga valid")
		return
	}

	query := "INSERT INTO users (name, age, address, password, email) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		SendErrorResponse(c, 500, "error preparing SQL statement")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, age, address, password, email)
	if err != nil {
		SendErrorResponse(c, 500, "error executing SQL statement")
		return
	}

	SendSuccessResponse(c, http.StatusCreated, "User created successfully")
}

func GinDeleteUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	userId := c.Param("id")

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userId).Scan(&count)
	if err != nil {
		SendErrorResponse(c, 500, "error preparing SQL statement")
		return
	}
	if count == 0 {
		SendErrorResponse(c, 404, "user not found")
		return
	}

	query := "DELETE FROM users WHERE id = ?"

	_, err = db.Exec(query, userId)
	if err != nil {
		SendErrorResponse(c, 400, "bad request")
		return
	}
	SendSuccessResponse(c, 200, "berhasil")
}

func GinUpdateUser(c *gin.Context) {
	db := Connect()
	defer db.Close()

	name := c.Query("name")
	address := c.Query("address")
	userId := c.Param("id")

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userId).Scan(&count)
	if err != nil {
		SendErrorResponse(c, 500, "error preparing SQL statement")
		return
	}
	if count == 0 {
		SendErrorResponse(c, 404, "user not found")
		return
	}

	query := "UPDATE users SET name = ?, address = ? WHERE id = ?"

	_, err = db.Exec(query, name, address, userId)
	if err != nil {
		SendErrorResponse(c, 400, "bad request")
		return
	}
	SendSuccessResponse(c, 200, "berhasil")
}
