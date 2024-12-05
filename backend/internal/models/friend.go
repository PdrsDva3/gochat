package models

type FriendList struct {
	IDUser    int   `json:"id_user"`
	IDFriends []int `json:"id_friends"`
}

type Friend struct {
	IDUser   int `json:"id_user"`
	IDFriend int `json:"id_friend"`
}
