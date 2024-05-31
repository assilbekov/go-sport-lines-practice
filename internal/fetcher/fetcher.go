package fetcher

import (
	"encoding/json"
	"fmt"
	"go-sport-lines-practice/internal/storage"
	"io"
	"net/http"
	"strconv"
)

type Fetcher struct {
	BaseURL string
}

type LineResponse struct {
	Soccer   string `json:"SOCCER"`
	Football string `json:"FOOTBALL"`
	Baseball string `json:"BASEBALL"`
}

type LinesResponse struct {
	Lines LineResponse `json:"lines"`
}

func NewFetcher(baseURL string) *Fetcher {
	return &Fetcher{
		BaseURL: baseURL,
	}
}

func (f *Fetcher) fetchAndParseSportLines(sport string) (*LinesResponse, error) {
	resp, err := http.Get(f.BaseURL + sport)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch sport line: %w", err)
	}
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

func (f *Fetcher) convertLinesResponseToLine(response *LinesResponse) (*storage.Line, error) {
	var line storage.Line
	var err error

	line.Soccer, err = strconv.ParseFloat(response.Lines.Soccer, 64)
	if err != nil {
		return &line, fmt.Errorf("failed to parse soccer line: %w", err)
	}

	line.Football, err = strconv.ParseFloat(response.Lines.Football, 64)
	if err != nil {
		return &line, fmt.Errorf("failed to parse football line: %w", err)
	}

	line.Baseball, err = strconv.ParseFloat(response.Lines.Baseball, 64)
	if err != nil {
		return &line, fmt.Errorf("failed to parse baseball line: %w", err)
	}

	return &line, nil
}

func (f *Fetcher) FetchSportLines(sport string) (*storage.Line, error) {
	resp, err := f.fetchAndParseSportLines(sport)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch and parse sport lines: %w", err)
	}

	line, err := f.convertLinesResponseToLine(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert lines response to line: %w", err)
	}

	return line, nil
}
