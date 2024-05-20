package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Line struct {
	Soccer   float64
	Football float64
	Baseball float64
}

func fetchLine(c chan Line) {
	for {
		c <- Line{
			Soccer:   1.5,
			Football: 2.5,
			Baseball: 3.5,
		}
	}
}

type LineJSON struct {
	Soccer   float64 `json:"SOCCER"`
	Football float64 `json:"FOOTBALL"`
	Baseball float64 `json:"BASEBALL"`
}

type LineProviderJSON struct {
	Lines LineJSON `json:"lines"`
}

type LinesInfo struct {
	Soccer   float64
	Football float64
	Baseball float64
}

func NewLinesInfo() *LinesInfo {
	return &LinesInfo{}
}

func (li *LinesInfo) fetchLinesInfo() {
	soccerTicker := time.NewTicker(1 * time.Second)
	footballTicker := time.NewTicker(3 * time.Second)
	baseballTicker := time.NewTicker(4 * time.Second)

	go func() {
		select {
		case <-soccerTicker.C:
			resp, err := http.Get("http://localhost:8000/api/v1/lines/soccer")
			if err != nil {
				fmt.Errorf("failed to fetch soccer line: %v", err)
			}
			defer resp.Body.Close()

			// Parse JSON response
			var lineProvider LineProviderJSON
			if err := json.NewDecoder(resp.Body).Decode(&lineProvider); err != nil {
				fmt.Errorf("failed to decode soccer line: %v", err)
			}

			li.Soccer = lineProvider.Lines.Soccer
		case <-footballTicker.C:
			resp, err := http.Get("http://localhost:8000/api/v1/lines/football")
			if err != nil {
				fmt.Errorf("failed to fetch football line: %v", err)
			}
			defer resp.Body.Close()

			// Parse JSON response
			var lineProvider LineProviderJSON
			if err := json.NewDecoder(resp.Body).Decode(&lineProvider); err != nil {
				fmt.Errorf("failed to decode football line: %v", err)
			}

			li.Football = lineProvider.Lines.Football

		case <-baseballTicker.C:
			resp, err := http.Get("http://localhost:8000/api/v1/lines/baseball")
			if err != nil {
				fmt.Errorf("failed to fetch baseball line: %v", err)
			}
			defer resp.Body.Close()

			// Parse JSON response
			var lineProvider LineProviderJSON
			if err := json.NewDecoder(resp.Body).Decode(&lineProvider); err != nil {
				fmt.Errorf("failed to decode baseball line: %v", err)
			}

			li.Baseball = lineProvider.Lines.Baseball
		}
	}()
}

func main() {
	c := make(chan Line)
	go fetchLine(c)

	linesInfo := NewLinesInfo()
	linesInfo.fetchLinesInfo()

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(linesInfo)
	}
}
