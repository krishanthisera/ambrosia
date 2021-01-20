package models

type Topic struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creation    string `json:"creation"`
	LastEdit    string `json:"lastedit"`
	Stories     string `json:"stories"`
}

type RedisConnect struct {
	ConString string
	Channel   string
}
