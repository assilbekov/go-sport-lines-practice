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

func fetchSportLine(sport string) (*LinesResponse, error) {
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

	var result LinesResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &result, nil
}

func processSportLine(sport string, response *LinesResponse) (*Line, error) {
	switch sport {
	case "soccer":
		soccer, err := strconv.ParseFloat(response.Lines.Soccer, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse soccer line: %w", err)
		}
		return &Line{Soccer: soccer}, nil
	case "football":
		football, err := strconv.ParseFloat(response.Lines.Football, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse football line: %w", err)
		}
		return &Line{Football: football}, nil
	case "baseball":
		baseball, err := strconv.ParseFloat(response.Lines.Baseball, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse baseball line: %w", err)
		}
		return &Line{Baseball: baseball}, nil
	default:
		return nil, fmt.Errorf("unknown sport: %s", sport)
	}
}

func main() {
	sports := []string{"soccer", "football", "baseball"}

	for _, sport := range sports {
		response, err := fetchSportLine(sport)
		if err != nil {
			fmt.Printf("failed to fetch sport line: %v\n", err)
			continue
		}

		fmt.Printf("response: %+v\n", response)

		line, err := processSportLine(sport, response)
		if err != nil {
			fmt.Printf("failed to process sport line: %v\n", err)
			continue
		}

		fmt.Printf("sport: %s, line: %+v\n", sport, line)
	}
}
