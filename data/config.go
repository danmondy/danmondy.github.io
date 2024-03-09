package data

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host         string       `json:"host"`
	Port         string       `json:"port"`
	Db           DbConnection `json:"db"`
	TemplatePath string       `json:"templateFolder"`
	PublicDir    string       `json:"publicFolder"`
}

type DbConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ReadJSON(file string, c interface{}) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, c)
}

func WriteJSON(file string, c interface{}) error {
	bytes, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return os.WriteFile(file, bytes, os.ModePerm)
}
