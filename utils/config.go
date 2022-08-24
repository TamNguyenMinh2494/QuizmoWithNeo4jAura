package utils

type Config struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Uri          string `json:"uri"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}
