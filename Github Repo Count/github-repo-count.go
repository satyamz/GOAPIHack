package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Repo : Struct repo
type Repo struct {
	Count int64 `json:"public_repos"`
}

func main() {
	var id string
	var result Repo
	fmt.Println("Enter Github user ID:")
	fmt.Scanf("%s", &id)
	uri := []string{"https://api.github.com/user/", id}
	final := strings.Join(uri, "")

	resp, err := http.Get(final)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	fmt.Println("Public Repositories:")
	fmt.Println(result.Count)
}
