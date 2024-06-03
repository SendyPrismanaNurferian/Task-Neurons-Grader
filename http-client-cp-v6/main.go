package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	// Hit API https://animechan.xyz/api/quotes/anime?title=naruto with method GET:
	resp, err := http.Get("https://animechan.xyz/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotes []Animechan
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data   `json:"data"`
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	requestBody := bytes.NewBuffer(postBody)

	// Hit API https://postman-echo.com/post with method POST:
	resp, err := http.Post("https://postman-echo.com/post", "application/json", requestBody)
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	var postmanResponse Postman
	err = json.Unmarshal(body, &postmanResponse)
	if err != nil {
		return Postman{}, err
	}

	return postmanResponse, nil
}

func main() {
	get, err := ClientGet()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Anime quotes:", get)
	}

	post, err := ClientPost()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Postman response:", post)
	}
}
