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

var Emailafterlogin string

func RegisterUser(c *gin.Context) {
	fmt.Println("Register user")
	c.Writer.Header().Set("Content-Type", "applicatiom/json")

	var user models.User
	json.NewDecoder(c.Request.Body).Decode(&user)
	// fmt.Println(user.Password)
	// fmt.Println(user.UserEmail)
	//validation for empty user
	if user.UserName == "" && user.UserEmail == "" && user.Password == "" {
		c.JSON(http.StatusNotAcceptable, "empty fields")
		return
	}

	db := database.SetupDB()
	// validation for duplicate user
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	var usersFromDb []models.User
	for rows.Next() {
		// fmt.Println("for loop....")
		var userid int
		var username string
		var email string
		var password string
		var retypepassword string
		err := rows.Scan(&userid, &username, &email, &password, &retypepassword)
		if err != nil {
			log.Fatal(err)
		}

		usersFromDb = append(usersFromDb, models.User{UserId: userid, UserName: username, UserEmail: email, RetypePassword: retypepassword})
		fmt.Println(usersFromDb)
	}
	for i := 0; i < len(usersFromDb); i++ {
		if usersFromDb[i].UserEmail == user.UserEmail {
			c.JSON(http.StatusConflict, "This user already exist")
			return
		}
	}

	// validation for same password
	if user.Password != user.RetypePassword {
		c.JSON(http.StatusBadRequest, "Please enter both password same")
		return
	}
	//inserting user to DB
	db.Query("INSERT INTO public.users(user_name,user_email,user_password,retype_password) VALUES($1,$2,$3,$4);", user.UserName, user.UserEmail, user.Password, user.RetypePassword)
	c.JSON(200, "user registered!")
}

func LoginUserHandler(c *gin.Context) {
	fmt.Println("Login controller...")
	c.Writer.Header().Set("Content-Type", "applicatoin/json")
	var userLogIn models.LoginUser
	json.NewDecoder(c.Request.Body).Decode(&userLogIn)

	var emailAndPwd models.LoginUser
	db := database.SetupDB()
	email := userLogIn.Email
	pswd := userLogIn.Password

	rows, err := db.Query("SELECT user_email,user_password FROM users where user_email=$1 AND user_password=$2", email, pswd)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var email string
		var password string

		err := rows.Scan(&email, &password)
		if err != nil {
			log.Fatal(err)
		}
		emailAndPwd = models.LoginUser{Email: email, Password: password}
	}

	Emailafterlogin = userLogIn.Email

	// fmt.Println("from db")
	// fmt.Println(emailAndPwd)
	// fmt.Println("from user")
	// fmt.Println(userLogIn.Email, userLogIn.Password)
	if userLogIn.Email == emailAndPwd.Email && userLogIn.Password == emailAndPwd.Password {
		c.JSON(200, "Valid user")
	} else {
		c.JSON(404, "User not found")
	}
}

func Get() string {
	var val = Emailafterlogin
	return val
}
