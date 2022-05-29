package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graduation-project/CityGO_bot/lib/e"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) (updates []Update, err error) {
	defer func() { err = e.WrapIfErr("can't get updates", err) }()

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	//fmt.Println("res = ", res)
	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	// fmt.Println("Host:", c.host)
	// fmt.Println("Path:", path.Join(c.basePath, method))
	// fmt.Println("   c.basePath:", c.basePath)
	// fmt.Println("   method:", method)

	// fmt.Println("\n\nURRRRLLLL: ", u.String())
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetFile(id string) (fileInfo GetFileRespInfo, err error) {

	fmt.Println("NEW URL GetFile ", "https://"+c.host+"/"+c.basePath+"/getFile")
	request := GetFileReq{FileId: id}
	body, _ := json.Marshal(request)

	resp, err := http.Post("https://"+c.host+"/"+c.basePath+"/getFile", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("GetFile | Error: %v", err)
		return
	}
	defer resp.Body.Close()
	//Сделать реализацию клиент - сервис.
	result := GetFileResp{}
	respBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &result)
	log.Printf("GetFile | Info: result - %v", result)
	fileInfo = result.Result
	return
}

func (c *Client) DownloadFileByPath(fileInfo GetFileRespInfo) (filePath string) {
	//https://api.telegram.org/file/bot5207175858:AAEAt-fvFASGsEQO0tTUmmhv5J02f9O2p3k/thumbnails/file_0.jpg

	path := fileInfo.FilePath
	fmt.Println("\n\n\n\nhttps://" + c.host + "/file/" + c.basePath + "/" + path + "\n\n\n")
	resp, err := http.Get("https://" + c.host + "/file/" + c.basePath + "/" + path)
	if err != nil {
		log.Printf("DownloadFileByPath | Error: %v", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error =", err.Error())
	}
	filePath = fmt.Sprintf("temp/pics/%s.jpeg", fileInfo.FileId)
	ioutil.WriteFile(filePath, data, 0666)
	//log.Println("I saved your image buddy!")
	return
}
