package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//ImageKeywords : struct to create an array of ImageData
type ImageKeywords struct {
	ImageDataInfo []ImageData `json:"imageKeywords"`
}

//ImageData : Struct to bind image metadata
type ImageData struct {
	Text  string `json:"text"`
	Score string `json:"score"`
}

func main() {
	var uri string
	fmt.Println("Enter image url")
	fmt.Scanf("%s", &uri)
	url := []string{"http://gateway-a.watsonplatform.net/calls/url/URLGetRankedImageKeywords?apikey=Your-Alchemy-API-KEY&outputMode=json&knowledgeGraph=0&url=", uri}
	finalUrl := strings.Join(url, "")
	// fmt.Println(finalUrl)
	resp, err := http.Get(finalUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result ImageKeywords
	json.Unmarshal(body, &result)
	// fmt.Println(result)
	fmt.Println("Keywords related to image")
	fmt.Println(result.ImageDataInfo[0].Text)

}
