package telegram

import (
	"encoding/json"
	"fmt"
	"graduation-project/CityGO_bot/events/backend"
	"graduation-project/CityGO_bot/models"
	"io/ioutil"
	"log"
	"math/rand"
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
		return p.tg.SendMessage(chatID, msgUserNotFound)
	} else if currentUser == (models.User{}) {
		log.Printf("doCmd | Error: User object is empty")
		return p.tg.SendMessage(chatID, msgUserNotFound)
	}

	switch text {
	case "/mychallenges":
		//Поиск челленджей юзера
		usersChallenges, err := backend.GetUsersChallenges(currentUser)
		if err != nil {
			log.Printf("doCmd | Error: %s", err.Error())
			return p.tg.SendMessage(chatID, msgUserNotFound)
		}

		//МОКИ ===============================
		// usersChallenges := make([]models.Challenge, 0)

		// usersChallenges = append(usersChallenges, models.Challenge{
		// 	ID:            "1",
		// 	CompanyID:     "1",
		// 	Title:         "Челлендж от Додо Пиццы!!!",
		// 	Description:   "Разгадай локацию, найди код и получи промокод на большую пиццу",
		// 	AnswerCode:    "4FL2MDCL",
		// 	TagsIds:       nil,
		// 	AttachmentIds: nil,
		// 	Tags:          nil,
		// 	Attachments:   nil,
		// 	StartDate:     "02-04-2022",
		// 	EndDate:       "02-05-2022",
		// })
		// usersChallenges = append(usersChallenges, models.Challenge{
		// 	ID:            "2",
		// 	CompanyID:     "2",
		// 	Title:         "Странное испытание",
		// 	Description:   "Разгадай локацию, найди код и получи 50% скидку на покупку любого товара в Marwin",
		// 	AnswerCode:    "8ID3MD3F",
		// 	TagsIds:       nil,
		// 	AttachmentIds: nil,
		// 	Tags:          nil,
		// 	Attachments:   nil,
		// 	StartDate:     "05-04-2022",
		// 	EndDate:       "25-05-2022",
		// })

		//надо вынести в отдельную функцию логику создания сообщения "Список челленджей"

		var usersChallengesMessage string
		for i, challenge := range usersChallenges {
			usersChallengesMessage += fmt.Sprintf("%v. %s\n", i+1, challenge.Title)
		}
		return p.tg.SendMessage(chatID, usersChallengesMessage)
	case "/mychallengesfullinfo":
		//Поиск челленджей юзера
		//usersChallenges, err := getUsersChallenges(currentUser)

		//МОКИ ===============================
		usersChallenges := make([]models.Challenge, 0)

		usersChallenges = append(usersChallenges, models.Challenge{
			ID:            "1",
			CompanyID:     "1",
			Title:         "Челлендж от Додо Пиццы!!!",
			Description:   "Разгадай локацию, найди код и получи промокод на большую пиццу",
			AnswerCode:    "4FL2MDCL",
			TagsIds:       nil,
			AttachmentIds: nil,
			Tags:          nil,
			Attachments:   nil,
			StartDate:     "02-04-2022",
			EndDate:       "02-05-2022",
		})
		usersChallenges = append(usersChallenges, models.Challenge{
			ID:            "2",
			CompanyID:     "2",
			Title:         "Странное испытание",
			Description:   "Разгадай локацию, найди код и получи 50% скидку на покупку любого товара в Marwin",
			AnswerCode:    "8ID3MD3F",
			TagsIds:       nil,
			AttachmentIds: nil,
			Tags:          nil,
			Attachments:   nil,
			StartDate:     "05-04-2022",
			EndDate:       "25-05-2022",
		})

		//надо вынести в отдельную функцию логику создания сообщения "Список челленджей"

		var usersChallengesMessage string
		for i, challenge := range usersChallenges {
			usersChallengesMessage += fmt.Sprintf("%v. %s\n\nПериод: %s\n\n%s\n\n\n", i+1, challenge.Title, (challenge.StartDate + " - " + challenge.EndDate), challenge.Description)
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
	default:
		if rand.Int()%2 == 0 {
			return p.tg.SendMessage(chatID, msgCodeActivatedSuccesfully+fmt.Sprintf(" \"%s\"", "Челлендж от Додо Пиццы!!!"))
		}
		return p.tg.SendMessage(chatID, msgCodeNotFound)
	}
}

func (p *Processor) sendHelp(chatID int, login string) error {
	return p.tg.SendMessage(chatID, msgHelp+fmt.Sprintf(" \"%s\"", login))
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
