package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user/repos", nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer Your_Token")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s", result)
}
