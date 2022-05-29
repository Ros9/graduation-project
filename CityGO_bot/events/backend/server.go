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

	// for _, x := range challenges {
	// 	fmt.Println(x)
	// }

	return
}

func PostAnswerCode(userID, code string) (resultMessage string, err error) {
	answer := models.Answer{UserID: userID, Answer: code}
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

	//resultMessage = //fmt.Sprintf("%v", result.Status)
	if result.Status == 1 {
		resultMessage = fmt.Sprintf("Ура!!! Вы прошли челлендж %s!!!\n\nЗаходите на сайт CityGO.kz и проходите новые челленджи!\nА также вы можете воспользоваться командой /challenges чтобы посмотреть список активных челленджей. GO!", result.Challenge.Title)
	} else if result.Status == 0 {
		resultMessage = fmt.Sprintf("К сожалению код не подошел к активным челленджам :(\nПопробуй ввести снова!")
	}

	return
}

func CreateChallenge(challenge models.Challenge) (result string, challengeResponse models.Challenge) {
	body, _ := json.Marshal(challenge)

	resp, err := http.Post(serverUrl+"challenge/telegram", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("CreateChallenge | Error: %v", err)
		return
	}
	defer resp.Body.Close()
	//Сделать реализацию клиент - сервис.
	resultChallenge := models.Challenge{}
	respBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &resultChallenge)
	log.Printf("CreateChallenge | Info: result - %v", resultChallenge)

	if resultChallenge.Title == challenge.Title && resultChallenge.Description == challenge.Description {
		result = "Челлендж " + resultChallenge.Title + " успешно создан!\n(id = \"" + resultChallenge.ID + "\")\n\n Вы можете отправить фото для Челленджа"
	} else {
		result = "Error, try again"
	}
	challengeResponse = resultChallenge
	return
}

// func PostAttachment(objType, challengeId, filePath string) {
// 	values := map[string]io.Reader{
// 		"file":  mustOpen("main.go"), // lets assume its this file
// 		"other": strings.NewReader("hello world!"),
// 	}
// 	err := Upload(serverUrl+"attachment", values)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func Upload(url string, values map[string]io.Reader) (err error) {
// 	// Prepare a form that you will submit to that URL.
// 	var b bytes.Buffer
// 	w := multipart.NewWriter(&b)
// 	for key, r := range values {
// 		var fw io.Writer
// 		if x, ok := r.(io.Closer); ok {
// 			defer x.Close()
// 		}
// 		// Add an image file
// 		if x, ok := r.(*os.File); ok {
// 			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
// 				return
// 			}
// 		} else {
// 			// Add other fields
// 			if fw, err = w.CreateFormField(key); err != nil {
// 				return
// 			}
// 		}
// 		if _, err = io.Copy(fw, r); err != nil {
// 			return err
// 		}

// 	}
// 	// Don't forget to close the multipart writer.
// 	// If you don't close it, your request will be missing the terminating boundary.
// 	w.Close()

// 	// Now that you have a form, you can submit it to your handler.
// 	req, err := http.NewRequest("POST", url, &b)
// 	if err != nil {
// 		return
// 	}
// 	// Don't forget to set the content type, this will contain the boundary.
// 	req.Header.Set("Content-Type", w.FormDataContentType())

// 	client := http.Client{}
// 	// Submit the request
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return
// 	}

// 	// Check the response
// 	if res.StatusCode != http.StatusOK {
// 		err = fmt.Errorf("bad status: %s", res.Status)
// 	}
// 	return
// }
