package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

//ImageKeywords : struct to store image keywords
type ImageKeywords struct {
	ImageDataInfo []ImageData `json:"imageKeywords"`
}

//ImageData : Struct to store image metadata
type ImageData struct {
	Text  string `json:"text"`
	Score string `json:"score"`
}

func main() {
	var str string
	fmt.Println("Enter Youtube video URL")
	fmt.Scanf("%s", &str)
	uri1 := strings.Split(str, "=")
	youtubeID := uri1[1]
	uri := path.Join("img.youtube.com/vi", youtubeID, "default.jpg")
	url := []string{"http://gateway-a.watsonplatform.net/calls/url/URLGetRankedImageKeywords?apikey=Your-API-Key&outputMode=json&knowledgeGraph=0&url=http://", uri}
	finalUrl := strings.Join(url, "")
	resp, err := http.Get(finalUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result ImageKeywords
	json.Unmarshal(body, &result)
	fmt.Println(result.ImageDataInfo[0].Text)

}
