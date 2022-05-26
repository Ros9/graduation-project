package telegram

import (
	"encoding/json"
	"fmt"
	"graduation-project/CityGO_bot/events/backend"
	"graduation-project/CityGO_bot/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

type TestMessage struct {
	Message string //`json: "message"`
}

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	//Проверка юзера
	currentUser, err := backend.GetUserByLogin(username)
	if err != nil {
		log.Printf("doCmd | Error: %s", err.Error())
		return p.tg.SendMessage(chatID, msgUserNotFound+fmt.Sprintf(" \"%s\"", username))
	} else if currentUser == (models.User{}) {
		log.Printf("doCmd | Error: User object is empty")
		return p.tg.SendMessage(chatID, msgUserNotFound+fmt.Sprintf(" \"%s\"", username))
	}

	switch text {
	case "/mychallenges":
		//Поиск челленджей юзера
		usersChallenges, err := backend.GetUsersChallenges(currentUser)
		if err != nil {
			log.Printf("doCmd | Error: %s", err.Error())
			return p.tg.SendMessage(chatID, msgUserNotFound)
		} else if len(usersChallenges) == 0 {
			log.Printf("doCmd | Error: %s", "user doesn't have challlenges")
			return p.tg.SendMessage(chatID, "У тебя нет активных челленджей!")
		}
		var usersChallengesMessage string
		for i, challenge := range usersChallenges {
			usersChallengesMessage += fmt.Sprintf("%v. %s\n", i+1, challenge.Title)
		}
		return p.tg.SendMessage(chatID, usersChallengesMessage)
	case "/mychallengesfullinfo":
		usersChallenges, err := backend.GetUsersChallenges(currentUser)
		if err != nil {
			log.Printf("doCmd | Error: %s", err.Error())
			return p.tg.SendMessage(chatID, msgUserNotFound)
		} else if len(usersChallenges) == 0 {
			log.Printf("doCmd | Error: %s", "user doesn't have challlenges")
			return p.tg.SendMessage(chatID, "У тебя нет активных челленджей!")
		}
		var usersChallengesMessage string
		for i, challenge := range usersChallenges {
			usersChallengesMessage += fmt.Sprintf("%v. %s\n\nПериод: %s\n\n%s\n\n\n", i+1, challenge.Title, (challenge.StartDate[:10] + " - " + challenge.EndDate[:10]), challenge.Description)
		}
		return p.tg.SendMessage(chatID, usersChallengesMessage)
	case "/challenges":
		challenges, err := backend.GetAvailableChallenges()
		if err != nil {
			log.Printf("doCmd | Error: %s", err.Error())
			return p.tg.SendMessage(chatID, msgUserNotFound)
		}
		usersChallengesMessage := "АКТИВНЫЕ ЧЕЛЛЕНДЖИ\n\n"
		for i, challenge := range challenges {
			usersChallengesMessage += fmt.Sprintf("%v. %s\n\nПериод: %s\n%s\n\n\n", i+1, challenge.Title, (challenge.StartDate[:10] + " - " + challenge.EndDate[:10]), challenge.Description)
		}
		return p.tg.SendMessage(chatID, usersChallengesMessage)
	case "/test":
		resp, err := http.Get("http://localhost:8080/test")
		if err != nil {
			return p.tg.SendMessage(chatID, msgUnknownCommand)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		testMessage := &TestMessage{}
		json.Unmarshal(body, testMessage)
		return p.tg.SendMessage(chatID, testMessage.Message)
	case "/privet":
		if username == "Alibi1" {
			return p.tg.SendMessage(chatID, "Ты крассавчик")
		} else if username == "Abdrassul" {
			return p.tg.SendMessage(chatID, "Хохлам здесь не место")
		} else if username == "Inuroboros" {
			return p.tg.SendMessage(chatID, "Привет Айнура Борос")
		} else if username == "mazyukaaa" {
			return p.tg.SendMessage(chatID, "Музыка?")
		}
		return p.tg.SendMessage(chatID, "I don't know you")
	case HelpCmd:
		return p.sendHelp(chatID, username)
	case StartCmd:
		return p.sendHello(chatID)
	case "/commands":
		commands := "Список всех комманд\n\n1. /challenges - Активные челленджи\n2. /mychallenges - Мои челленджи\n3. /mychallengesfullinfo - Полная инфа о моих челленджах\n4. /help - На помощь!"

		if currentUser.IsAdmin == 1 {
			commands += "\n\n Admin Commands\n1. /createchallenge"
		}

		return p.tg.SendMessage(chatID, commands)
	default:
		resultMessage, err := backend.PostAnswerCode(currentUser.ID, text)
		if err != nil {
			log.Printf("doCmd | Error: %s", err.Error())
			return p.tg.SendMessage(chatID, msgCodeNotFound)
		}
		return p.tg.SendMessage(chatID, resultMessage)
	}
}

func (p *Processor) sendHelp(chatID int, login string) error {
	fmt.Println(login)
	return p.tg.SendMessage(chatID, fmt.Sprintf("%s \"%s\"", msgHelp, login))
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
