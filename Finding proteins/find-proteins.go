package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type FoodName struct {
	Q      string `json:"q"`
	Max    string `json:"max"`
	Format string `json:"offset"`
}
type Number struct {
	No   string `json:"ndbno"`
	Type string `json:"type"`
}

type List struct {
	List Item `json:"list"`
}
type Item struct {
	Item []NDBN `json:"item"`
}
type NDBN struct {
	Ndbno string `json:"ndbno"`
}
type Report struct {
	ReportFinal Food `json:"report"`
}
type Food struct {
	FoodItem Neutrients `json:"food"`
}
type Neutrients struct {
	AllNeutrients []Element `json:"nutrients"`
}

type Element struct {
	NID   int64   `json:"nutrient_id"`
	Value float64 `json:"value"`
}

func main() {
	var str string
	in := bufio.NewReader(os.Stdin)

	str, err := in.ReadString('\n')

	url1 := "http://api.nal.usda.gov/ndb/search?api_key=DEMO_KEY"

	foodName := &FoodName{str, "25", "0"}
	buf, _ := json.Marshal(foodName)
	b := bytes.NewBuffer(buf)

	req, err := http.NewRequest("POST", url1, b)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var l List
	json.Unmarshal(body, &l)
	ndbno := l.List.Item[0].Ndbno
	url2 := "http://api.nal.usda.gov/ndb/reports?api_key=DEMO_KEY"
	no := &Number{ndbno, "f"}
	buf1, _ := json.Marshal(no)
	b1 := bytes.NewBuffer(buf1)
	req1, err := http.NewRequest("POST", url2, b1)
	if err != nil {
		panic(err)
	}
	req1.Header.Set("Content-Type", "application/json")

	client1 := &http.Client{}
	resp1, err := client1.Do(req1)
	if err != nil {
		panic(err)
	}

	defer resp1.Body.Close()

	var FinalReport Report
	body1, _ := ioutil.ReadAll(resp1.Body)
	json.Unmarshal(body1, &FinalReport)
	AnsArray := FinalReport.ReportFinal.FoodItem.AllNeutrients

	var Val float64
	for _, element := range AnsArray {
		if element.NID == 203 {
			Val = element.Value
			break
		}
	}

	strAns := strconv.FormatFloat(Val, 'f', -1, 64)
	fmt.Println(strAns + "g")

}
