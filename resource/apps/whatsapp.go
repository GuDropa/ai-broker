package modules

import (
	"fmt"
	"io"
	"net/http"
)

type Channel struct {
	Id      string
	Name    string
	BaseUrl string
}

func (c *Channel) GetChats(opts any) (any, error) {

	// url := "https://gate.whapi.cloud/chats?count=100"
	url := c.BaseUrl + "/chats?count=100"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	return nil, nil
}
