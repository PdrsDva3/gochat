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

type FriendList struct {
	IDUser    int   `json:"id_user"`
	IDFriends []int `json:"id_friends"`
}

type FriendInfo struct {
	IDFriend int    `json:"id_friend"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
}

type FriendListInfo struct {
	IDUser  int          `json:"id_user"`
	Friends []FriendInfo `json:"friends"`
}

type Friend struct {
	IDUser   int
	IDFriend int
}

type ChatBase struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ChatList struct {
	IDUser int        `json:"id"`
	Chats  []ChatBase `json:"chats"`
}

type Chat struct {
	ChatBase
	Description string `json:"description"`
	IDCreator   int    `json:"id_creator"`
	IDUsers     []int  `json:"id_users"`
}

type ChatCreate struct {
	IDCreator   int    `json:"id_creator"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IDUsers     []int  `json:"id_users"`
}
