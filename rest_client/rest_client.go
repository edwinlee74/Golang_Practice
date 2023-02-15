package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io" //ioutil 已在1.19版本中捨棄, 改以os或io替代
	"log"
	"net/http"
)

var client = &http.Client{}
var api_url = "https://api.github.com/repos/edwinlee74/Golang_Practice/issues"
var token = "Bearer {YOUR_TOKEN}"

func get_method() {
	req, err := http.NewRequest(http.MethodGet, api_url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)
}

func post_method() {
	type Payload struct {
		Title     string   `json:"title"`
		Body      string   `json:"body"`
		Assignees []string `json:"assignees"`
		Milestone int      `json:"milestone"`
		Labels    []string `json:"labels"`
	}

	payload := Payload{
		Title:     "Found a bug",
		Body:      "I'\\''m having a problem with this.",
		Assignees: []string{"edwinlee74"},
		Milestone: 2,
		Labels:    []string{"bug"},
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", api_url, bytes.NewBuffer(payloadJson))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)
}

func patch_method() {
	type Payload struct {
		Title     string   `json:"title"`
		Body      string   `json:"body"`
		Assignees []string `json:"assignees"`
		Milestone int      `json:"milestone"`
		State     string   `json:"state"`
		Labels    []string `json:"labels"`
	}

	payload := Payload{
		Title:     "Found a bug",
		Body:      "I'\\''m having a problem with this.",
		Assignees: []string{"edwinlee74"},
		Milestone: 2,
		State:     "open",
		Labels:    []string{"bug"},
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", api_url+"/3", bytes.NewBuffer(payloadJson))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)
}

func put_method() {
	type Payload struct {
		Lock_Reason string `json:"lock_reason"`
	}

	payload := Payload{
		Lock_Reason: "off-topic",
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest(http.MethodPut, api_url+"/3/lock", bytes.NewBuffer(payloadJson))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)
}

func delete_method() {
	req, err := http.NewRequest(http.MethodDelete, api_url+"/3/lock", nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", token)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)
}

func main() {
	post_method()

}
