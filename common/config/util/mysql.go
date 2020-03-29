package util

type Mysql struct {
	Address      string `json:"address"`
	Port         int    `json:"port"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	DbName       string `json:"db_name"`
}
