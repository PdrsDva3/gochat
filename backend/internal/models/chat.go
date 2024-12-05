package models

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
