package models

//todo
type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	User  *User  `json:"user"`
}

//register
type User struct {
	UserId         int    `json:"userid"`
	UserName       string `json:"username"`
	UserEmail      string `json:"email"`
	Password       string `json:"password"`
	RetypePassword string `json:"retypepassword"`
}

//login
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
