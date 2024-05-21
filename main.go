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

func fetchSportLine(sport string) (*LineResponse, error) {
	resp, err := http.Get(baseURL + sport)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch sport line: %w", err)
	}

	// 1. Why do we need to close the response body?
	// 2. How to handle the error if the response body fails to close?
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Printf("error closing response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result LineResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

func main() {
	resp, err := http.Get(baseURL + "soccer")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

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
