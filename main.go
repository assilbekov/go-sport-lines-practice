package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const baseURL = "http://localhost:8000/api/v1/lines/"

type Line struct {
	Soccer   float64
	Football float64
	Baseball float64
}

type LineResponse struct {
	Soccer   string `json:"SOCCER"`
	Football string `json:"FOOTBALL"`
	Baseball string `json:"BASEBALL"`
}

type LinesResponse struct {
	Lines LineResponse `json:"lines"`
}

func main() {
	resp, err := http.Get(baseURL + "soccer")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("unexpected status code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result LinesResponse
	if err = json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	fmt.Println(result)

	soccer, err := strconv.ParseFloat(result.Lines.Soccer, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(Line{Soccer: soccer})
}
