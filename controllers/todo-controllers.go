package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/database"
	"todo/models"

	"github.com/gin-gonic/gin"
)

// GET REQUEST
func GetAllTodo(c *gin.Context) {
	fmt.Println("Get All todo list")
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var todo []models.Todo
	// todo = append(todo, models.Todo{Title: "Coding", Body: "Basic Code practice", Id: 10})
	db := database.SetupDB()

	//id get from login email
	var sessionEmail = Get()
	var userid int
	db.QueryRow("SELECT user_id from users WHERE user_email = $1", sessionEmail).Scan(&userid)
	fmt.Printf("Id is %d", userid)

	rows, _ := db.Query("SELECT title,body,id FROM todo WHERE user_id = $1", userid)

	for rows.Next() {
		var title string
		var body string
		var id int
		err := rows.Scan(&title, &body, &id)
		if err != nil {
			log.Fatal(err)
		}
		todo = append(todo, models.Todo{Title: title, Body: body, Id: id})
	}

	fmt.Println(todo)
	c.JSON(200, todo)
}

// POST REQUEST
func AddOneTodo(c *gin.Context) {
	fmt.Println("Add one todo")
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var todo models.Todo
	json.NewDecoder(c.Request.Body).Decode(&todo)

	//validation for empty title
	if todo.Title == "" {
		// c.JSON(http.StatusNoContent"No data inside json")  //204
		c.JSON(http.StatusNotAcceptable, "No data inside json body")
		return
	}
	//validation of tuplicate title
	db := database.SetupDB()
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err)
	}
	var todosFromDb []models.Todo
	for rows.Next() {
		var id int
		var title string
		var body string
		var userid int
		err := rows.Scan(&id, &title, &body, &userid)
		if err != nil {
			log.Fatal(err)
		}
		todosFromDb = append(todosFromDb, models.Todo{Id: id, Title: title, Body: body, User: &models.User{UserId: userid}})
		fmt.Println(todosFromDb)
	}
	fmt.Println(todosFromDb)
	for i := 0; i < len(todosFromDb); i++ {
		if todosFromDb[i].Title == todo.Title {
			c.JSON(http.StatusConflict, "This todo already exist")
			return
		}
	}

	var sessionEmail = Get()
	var userid int
	db.QueryRow("SELECT user_id from users WHERE user_email = $1", sessionEmail).Scan(&userid)
	// fmt.Println(userid)
	fmt.Println(todo.Title, todo.Body, userid)
	db.Query("INSERT INTO todo(title,body,user_id) VALUES($1,$2,$3)", todo.Title, todo.Body, userid)
	c.JSON(200, "Todo added...")
}

// DELETE REQUEST
func DeleteOneTodo(c *gin.Context) {
	fmt.Println("Delete One todo")
	c.Writer.Header().Set("Content-type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var todo models.Todo
	json.NewDecoder(c.Request.Body).Decode(&todo)
	// fmt.Println(todo)

	db := database.SetupDB()
	id := c.Param("id")
	db.Query("DELETE FROM todo WHERE id = $1", id)
	c.JSON(http.StatusOK, "Todo removed of given id")
}

//test API
func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to todo Backend API"})
}

// just for testing
func GetAllList(c *gin.Context) {
	fmt.Println("Get All list")
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var todo []models.Todo
	// todo = append(todo, models.Todo{Title: "Coding", Body: "Basic Code practice", Id: 10})
	log.Print("Before db")
	db := database.SetupDB()
	log.Print("After db")
	//id get from login email
	rows, err := db.Query("SELECT * FROM todo")
	log.Print("After query executed")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var title string
		var body string
		var userid int
		err := rows.Scan(&id, &title, &body, &userid)
		if err != nil {
			log.Fatal(err)
		}
		todo = append(todo, models.Todo{Title: title, Body: body, Id: id, User: &models.User{UserId: userid}})
	}
	fmt.Println(todo)
	c.JSON(200, todo)
}
