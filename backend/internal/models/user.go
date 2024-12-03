package models

type UserBase struct {
	Name        string
	Surname     string
	Nickname    string
	Email       string
	Phone       string
	Photo       string
	Description string
}

type UserCreate struct {
	UserBase
	Password string
}

type User struct {
	ID int
	UserBase
}

type UserChange struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`
	Photo    string `json:"photo"`
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

type FriendList struct {
	IDUser    int
	IDFriends []int
}

type ChatList struct {
	IDUser  int
	IDChats []int
}

type Chat struct {
	ID          int
	Name        string
	Description string
	IDCreator   int
	IDUsers     []int
}
