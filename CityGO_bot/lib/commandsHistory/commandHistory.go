package commandshistory

import (
	"bufio"
	"graduation-project/CityGO_bot/models"
	"strings"
)

type CommandHistoryItem struct {
	ChatId int
	User   models.User
	Text   string
}

func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func StringToDateFormat(s string) (date string) {
	//2020-07-15T00:00:00Z
	date = s + "T00:00:00Z"
	return
}

//"1,2,3,4"
func StringToIdList(s string) (ids []string) {
	ids = strings.Split(s, ",")
	return
}
