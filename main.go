package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Status Status `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Receiver  string    `json:"receivedBy,omitempty"`
	Histories []History `json:"histories"`
}

type History struct {
	Description string    `json:"description"`
	CreatedAt   string    `json:"createdAt"`
	Formatted   Formatted `json:"formatted"`
}

type Status struct {
	Code              string            `json:"code"`
	Message           string            `json:"message"`
	ReceivementStatus ReceivementStatus `json:"-"`
}

type ReceivementStatus struct {
	IsReceived bool   `json:"is_received"`
	Receiver   string `json:"receiver"`
}

type Formatted struct {
	CreatedAt string `json:"createdAt"`
}

func main() {
	var url = "https://gist.githubusercontent.com/nubors/eecf5b8dc838d4e6cc9de9f7b5db236f/raw/d34e1823906d3ab36ccc2e687fcafedf3eacfac9/jne-awb.html"
	var result Response

	html, err := requestAndGetHTMLData(url)
	if err != nil {
		panic(err)
	}

	histories, currentStatus, err := parsingHtmlToHistory(html)
	if err != nil {
		panic(err)
	}

	histories = revereseArray(histories)
	deliveryStatus := checkStatus(currentStatus)

	result.Status = deliveryStatus
	result.Data.Receiver = deliveryStatus.ReceivementStatus.Receiver
	result.Data.Histories = histories

	response, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(response))
}

// requestAndGetHTMLData get html from url and
// convert to string
func requestAndGetHTMLData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

// parsingHtmlToHistory parsing html and
// convert to []History and get currentStatus
func parsingHtmlToHistory(html string) ([]History, string, error) {
	var histories []History
	var status string
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, "", err
	}

	doc.Find(".tracking").Each(func(i int, sel *goquery.Selection) {
		if strings.Contains(sel.Text(), "History") {
			history := History{}
			sel.Find("tbody td").Each(func(j int, sel2 *goquery.Selection) {
				if j%2 == 0 {
					t, _ := time.ParseInLocation(layout, sel2.Text(), jakartaTime)
					history.CreatedAt = t.Format(timeLayout)
					history.Formatted.CreatedAt = r.Replace(t.Format(formattedLayout))
				} else {
					history.Description = sel2.Text()
					histories = append(histories, history)
					history = History{}
					status = sel2.Text()
				}
			})
		}
	})
	return histories, status, nil
}

// revereseArray reverse array of history
func revereseArray(arr []History) []History {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// checkStatus get latest status of tracking
func checkStatus(currentStatus string) Status {
	var history Status
	for stat, value := range status {
		if strings.Contains(currentStatus, stat) {
			history = value
			if stat == "DELIVERED TO" {
				status1 := strings.Split(currentStatus, "[")
				if len(status) > 1 {
					receiver := strings.Split(status1[1], "|")
					history.ReceivementStatus.IsReceived = true
					history.ReceivementStatus.Receiver = receiver[0]
				}
			}
			break
		}
	}
	return history
}
