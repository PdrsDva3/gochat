package models

type UserBase struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type User struct {
	ID int `json:"id"`
	UserBase
}

type UserChange struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Nickname    string `json:"nickname"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

type UserChangePWD struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type UserChangeEmail struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserChangePhone struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
