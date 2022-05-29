package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"graduation-project/render-api/model"
	"io/ioutil"
	"net/http"
)

var (
	httpClt  *http.Client
	sessions map[string]*model.User
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/templates/*")
	router.Static("/assets", "./assets")

	sessions = make(map[string]*model.User)

	httpClt = http.DefaultClient

	router.Handle("GET", "/signup", func(context *gin.Context) {
		context.HTML(200, "signup.html", nil)
	})

	router.Handle("GET", "/signup2", func(context *gin.Context) {
		context.HTML(200, "signup2.html", nil)
	})

	router.Handle("GET", "/signin", func(context *gin.Context) {
		context.HTML(200, "signin.html", nil)
	})

	router.Handle("GET", "/index", func(context *gin.Context) {
		challenges := GetChallenges()
		fmt.Println("challenges =", challenges)
		context.HTML(200, "index.html", challenges)
	})

	router.Handle("GET", "/challenges", func(context *gin.Context) {
		token, err := context.Cookie("token")
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		_, ok := sessions[token]
		isOnline := false
		if !ok {
			fmt.Println("user has not exist")
		} else {
			isOnline = true
		}
		challenges := GetChallenges()
		fmt.Println("challenges =", challenges)
		companies := GetCompanies()
		data := map[string]interface{}{
			"challenges": challenges,
			"companies":  companies,
			"isOnline":   isOnline,
		}
		context.HTML(200, "quests.html", data)
	})

	router.Handle("GET", "/challenge/:id", func(context *gin.Context) {
		token, err := context.Cookie("token")
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		user, ok := sessions[token]
		isOnline := false
		if !ok {
			fmt.Println("user has not exist")
		} else {
			isOnline = true
		}
		challengeId := context.Param("id")
		challenge := GetChallengeById(challengeId)
		company := GetCompanyById(challenge.CompanyID)
		comments, err := GetCommentsByChallengeId(challengeId)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		fmt.Println("comments =", comments)
		data := map[string]interface{}{
			"challenge": challenge,
			"company":   company,
			"user":      user,
			"isOnline":  isOnline,
			"comments":  comments,
		}
		context.HTML(200, "detail-quest.html", data)
	})

	router.GET("/profile2", func(context *gin.Context) {
		token, err := context.Cookie("token")
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		user, ok := sessions[token]
		isOnline := false
		if !ok {
			fmt.Println("user has not exist")
			context.Redirect(307, "/signin")
		} else {
			isOnline = true
		}
		challenges := GetChallengesByUserId(user.ID)
		fmt.Println("number of number_of_answers =", len(challenges))
		data := map[string]interface{}{
			"user":              user,
			"number_of_answers": len(challenges),
			"challenges":        challenges,
			"is_online":         isOnline,
		}
		context.HTML(200, "your-progression.html", data)
	})

	router.GET("/achievements", func(context *gin.Context) {
		achievements := GetAchievements()
		data := map[string]interface{}{
			"achievements": achievements,
		}
		context.HTML(200, "achievements.html", data)
	})

	router.Handle("POST", "/registration", func(context *gin.Context) {
		username := context.Request.FormValue("username")
		email := context.Request.FormValue("email")
		telegram := context.Request.FormValue("telegram")
		password := context.Request.FormValue("password")
		user := model.User{
			Username: username,
			Password: password,
			Email:    email,
			Telegram: telegram,
			IsAdmin:  0,
		}
		fmt.Println("password =", password)
		createdUser, err := RegisterUser(user)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		fmt.Println("user =", createdUser)
		context.Redirect(301, "/signin")
	})

	router.Handle("POST", "/login", func(context *gin.Context) {
		username := context.Request.FormValue("username")
		password := context.Request.FormValue("password")
		authInfo := model.AuthInfo{
			Username: username,
			Password: password,
		}
		tokenStr, err := GetToken(authInfo)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		fmt.Println("tokenStr =", tokenStr)
		user, err := GetUserByToken(tokenStr)
		if err != nil {
			fmt.Println("error =", err.Error())
			return
		}
		context.SetCookie("token", tokenStr, 600, "/", "", true, true)
		sessions[tokenStr] = user
		fmt.Println("user =", user)
		context.Redirect(301, "/index")
	})

	router.POST("/answer", func(context *gin.Context) {
		token, err := context.Cookie("token")
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		user, ok := sessions[token]
		isOnline := false
		if !ok {
			fmt.Println("user has not exist")
		} else {
			isOnline = true
		}
		fmt.Println(user, isOnline)
		if !isOnline {
			context.Redirect(301, "/signin")
			return
		}
		answerStr := context.Request.FormValue("answer")
		challengeId := context.Request.FormValue("challenge_id")
		answer := model.Answer{
			UserID:      user.ID,
			ChallengeID: challengeId,
			Answer:      answerStr,
			Status:      1,
		}
		GiveAnswer(token, answer)
		context.Redirect(301, "/challenge/"+challengeId)
	})

	router.POST("/comment", func(context *gin.Context) {
		token, err := context.Cookie("token")
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		user, ok := sessions[token]
		isOnline := false
		if !ok {
			fmt.Println("user has not exist")
		} else {
			isOnline = true
		}
		if !isOnline {
			context.Redirect(301, "/signin")
			return
		}
		description := context.Request.FormValue("description")
		challengeId := context.Request.FormValue("challenge_id")
		fmt.Println("challengeId =", challengeId)
		comment := model.Comment{
			ChallengeID: challengeId,
			UserId:      user.ID,
			Description: description,
		}
		CreateComment(token, comment)
		context.Redirect(301, "/challenge/"+challengeId)
	})

	router.Run(":5000")
}

func RegisterUser(user model.User) (*model.User, error) {
	url := "http://localhost:8080/user/registration"
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	resp, err := httpClt.Do(req)
	createdUser := model.User{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return nil, err
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	err = json.Unmarshal(data, &createdUser)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	return &createdUser, nil
}

func GetUserByToken(tokenStr string) (*model.User, error) {
	url := "http://localhost:8080/user/info"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+tokenStr)
	resp, err := httpClt.Do(req)
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	user := model.User{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	return &user, nil
}

func GetToken(authInfo model.AuthInfo) (string, error) {
	url := "http://localhost:8080/auth/user"
	data, err := json.Marshal(authInfo)
	if err != nil {
		return "", err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	resp, err := httpClt.Do(req)
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return "", err
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return "", err
	}
	tokenStr := ""
	err = json.Unmarshal(data, &tokenStr)
	if err != nil {
		fmt.Println("error =", err.Error())
		return "", err
	}
	return tokenStr, nil
}

func GetChallengeById(challengeId string) model.Challenge {
	url := "http://localhost:8080/challenge/" + challengeId
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	challenge := model.Challenge{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return challenge
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return challenge
	}
	err = json.Unmarshal(data, &challenge)
	if err != nil {
		fmt.Println("error =", err.Error())
		return challenge
	}
	return challenge
}

func GetChallenges() []*model.Challenge {
	url := "http://localhost:8080/challenges"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	challenges := []*model.Challenge{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return challenges
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return challenges
	}
	err = json.Unmarshal(data, &challenges)
	if err != nil {
		fmt.Println("error =", err.Error())
		return challenges
	}
	return challenges
}

func GetChallengesByUserId(userId string) []*model.Challenge {
	url := "http://localhost:8080/challenges/user/" + userId
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	challenges := []*model.Challenge{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return challenges
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return challenges
	}
	err = json.Unmarshal(data, &challenges)
	if err != nil {
		fmt.Println("error =", err.Error())
		return challenges
	}
	return challenges
}

func GetCompanies() []*model.Company {
	url := "http://localhost:8080/companies"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	companies := []*model.Company{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return companies
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return companies
	}
	err = json.Unmarshal(data, &companies)
	if err != nil {
		fmt.Println("error =", err.Error())
		return companies
	}
	return companies
}

func GetCompanyById(companyId string) model.Company {
	url := "http://localhost:8080/company/" + companyId
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	company := model.Company{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return company
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return company
	}
	err = json.Unmarshal(data, &company)
	if err != nil {
		fmt.Println("error =", err.Error())
		return company
	}
	return company
}

func GiveAnswer(tokenStr string, answer model.Answer) (*model.Answer, error) {
	url := "http://localhost:8080/answer"
	data, err := json.Marshal(answer)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+tokenStr)
	resp, err := httpClt.Do(req)
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return nil, err
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	return nil, nil
	//user := model.User{}
	//err = json.Unmarshal(data, &user)
	//if err != nil {
	//	fmt.Println("error =", err.Error())
	//	return nil, err
	//}
	//return &user, nil
}

func GetAchievements() []*model.Achievement {
	url := "http://localhost:8080/achievements"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	achievements := []*model.Achievement{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return achievements
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return achievements
	}
	err = json.Unmarshal(data, &achievements)
	if err != nil {
		fmt.Println("error =", err.Error())
		return achievements
	}
	return achievements
}

func CreateComment(tokenStr string, comment model.Comment) (*model.Comment, error) {
	url := "http://localhost:8080/comment"
	data, err := json.Marshal(comment)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+tokenStr)
	resp, err := httpClt.Do(req)
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return nil, err
	}
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	return nil, nil
}

func GetCommentsByChallengeId(challengeId string) ([]*model.Comment, error) {
	url := "http://localhost:8080/comments/challenge/" + challengeId
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := httpClt.Do(req)
	comments := []*model.Comment{}
	if err != nil {
		fmt.Println("error when do request:", url, " =", err.Error())
		return comments, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println("data from", url, " =", string(data))
	if err != nil {
		fmt.Println("error =", err.Error())
		return comments, err
	}
	err = json.Unmarshal(data, &comments)
	if err != nil {
		fmt.Println("error =", err.Error())
		return comments, err
	}
	return comments, nil
}
