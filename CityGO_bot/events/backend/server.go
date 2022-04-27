package backend

import (
	"encoding/json"
	"graduation-project/CityGO_bot/models"
	"io/ioutil"
	"log"
	"net/http"
)

//client

const (
	serverUrl = "http://localhost:8080/"
)

// Надо вынести в отдельную функцию запросы.
func GetFromBackend(url string, body interface{}) (response interface{}, err error) {
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
	resp, err := http.Get("http://localhost:8080/user/telegram/" + login)
	if err != nil {
		log.Printf("getUserByLogin | Error: %s", login)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &user)
	log.Printf("getUserByLogin | Info: user %s %v found", login, user)
	return
}

func GetUsersChallenges(user models.User) (usersChallenges []models.Challenge, err error) {
	// body, err := GetFromBackend("challenges/user/"+user.ID, nil)
	// json.Unmarshal([]byte(body.(string)), &usersChallenges)
	// return
	resp, err := http.Get(serverUrl + "challenges/user/" + user.ID)
	if err != nil {
		log.Printf("GetUsersChallenges | Error: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &usersChallenges)
	log.Printf("GetUsersChallenges | Info: user %v found", usersChallenges)
	return
}
