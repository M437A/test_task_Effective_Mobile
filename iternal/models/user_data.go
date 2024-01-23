package models

type UserData struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int64  `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type UserDataRequest struct {
	Page       int        `json:"page"`
	UserFilter UserFilter `json:"user_filter"`
}

type UserFilter struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int64  `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}
