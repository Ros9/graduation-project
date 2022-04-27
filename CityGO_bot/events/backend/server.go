package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graduation-project/CityGO_bot/models"
	"io/ioutil"
	"log"
	"net/http"
)

//client

const (
	serverUrl = "http://localhost:8080/"
)

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
	resp, err := http.Get(serverUrl + "challenges/telegram/user/" + user.ID)
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

func GetAvailableChallenges() (challenges []models.Challenge, err error) {
	resp, err := http.Get(serverUrl + "challenges/telegram")
	if err != nil {
		log.Printf("GetAvailableChallenges | Error: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &challenges)
	log.Printf("GetAvailableChallenges | Info: user %v found", challenges)

	for _, x := range challenges {
		fmt.Println(x)
	}

	return
}

func PostAnswerCode(code string) (resultMessage string, err error) {
	answer := models.Answer{Answer: code}
	body, _ := json.Marshal(answer)

	resp, err := http.Post(serverUrl+"answer/telegram", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("PostAnswerCode | Error: %v", err)
		return
	}
	defer resp.Body.Close()
	//Сделать реализацию клиент - сервис.
	result := models.Result{}
	respBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &result)
	log.Printf("PostAnswerCode | Info: result - %v", result)

	resultMessage = fmt.Sprintf("%v", result.Status)
	return
}
