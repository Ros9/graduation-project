package backend

import (
	"CityGO_bot/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	serverUrl = "http://localhost:8080/"
)

//client

// Надо вынести в отдельную функцию запросы.
func getFromBackend(url string, body interface{}) (response interface{}, err error) {
	resp, err := http.Get(serverUrl + url)
	if err != nil {
		log.Printf("getFromBackend | Error: %s", err)
		return
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	return
}

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

func getUsersChallenges(user models.User) (usersChallenges []models.Challenge, err error) {
	body, err := getFromBackend("user/id/"+user.ID, nil)
	json.Unmarshal([]byte(body.(string)), usersChallenges)
	return
}
