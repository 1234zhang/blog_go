package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, passward string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: passward}).Find(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}