package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"web/model"
)

type WordsClient interface {
	GetWords() (data []*model.Data, err error)
	UpdateWord(word string) (err error)
}

type httpClient struct {
	apiURL string
}

func NewHttpClient(apiURL string) httpClient {
	return httpClient{apiURL}
}

func (h httpClient) GetWords() (data []*model.Data, err error) {
	u, _ := url.Parse(h.apiURL)
	u.Path = path.Join(u.Path, "words")
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Request to API failed")
	}
	defer res.Body.Close()
	byteArray, _ := ioutil.ReadAll(res.Body)
	jsonBytes := ([]byte)(byteArray)

	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	return data, nil
}

func (h httpClient) UpdateWord(word string) (err error) {
	u, _ := url.Parse(h.apiURL)
	u.Path = path.Join(u.Path, "words")
	reqVal := model.Request{Word: word}
	reqJSON, _ := json.Marshal(reqVal)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(reqJSON))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("Request to API failed")
	}
	return nil
}
