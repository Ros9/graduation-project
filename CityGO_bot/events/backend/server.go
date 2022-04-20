package backend

import (
	"CityGO_bot/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//client

func GetUserByLogin(login string) (user models.User, err error) {
	resp, err := http.Get("http://localhost:8080/user/login/" + login)
	if err != nil {
		log.Printf("getUserByLogin | Error: %s", login)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, user)
	log.Printf("getUserByLogin | Info: user %s found", login)
	return
}
