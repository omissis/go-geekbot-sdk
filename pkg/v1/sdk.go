package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const DefaultBaseURL = "https://api.geekbot.com/v1"

var (
	ErrCannotCreateNewRequest = errors.New("cannot create new request")
	ErrCannotListReports      = errors.New("cannot list reports")
	ErrCannotListStandups     = errors.New("cannot list standups")
	ErrCannotListTeams        = errors.New("cannot list teams")
)

func NewSDK(httpClient *http.Client, baseURL string, apiKey string) *SDK {
	return &SDK{
		httpClient: httpClient,
		baseURL:    baseURL,
		apiKey:     apiKey,
	}
}

type SDK struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// ListTeams will return an object describing the team the API key was generated for.
//
// There is a little inconsistency in the API as the method is called List Teams but it returns a single team.
func (s *SDK) ListTeams() (*Team, error) {
	var team Team

	if err := s.do("/teams/", &team); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotListTeams, err)
	}

	return &team, nil
}

// func (s *SDK) CreateReport(report any) ([]map[string]any, error) {
// 	return nil, nil
// }

func (s *SDK) ListReports(filters ListReportsFilters) ([]Report, error) {
	var reports []Report

	if err := s.do("/reports/?"+filters.QueryString(), &reports); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotListReports, err)
	}

	return reports, nil
}

func (s *SDK) ListStandups() ([]Standup, error) {
	var standups []Standup

	if err := s.do("/standups/", &standups); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotListStandups, err)
	}

	return standups, nil
}

// func (s *SDK) ReadStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

// func (s *SDK) CreateStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

// func (s *SDK) UpdateStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

// func (s *SDK) ReplaceStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

// func (s *SDK) DuplicateStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

// func (s *SDK) DeleteStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

// func (s *SDK) StartStandup() ([]map[string]any, error) {
// 	return nil, nil
// }

func (s *SDK) do(path string, result any) error {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, s.baseURL+path, nil)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotCreateNewRequest, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", s.apiKey)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotListTeams, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: API returned code %d", ErrCannotListTeams, res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return fmt.Errorf("%w: %w", ErrCannotListTeams, err)
	}

	return nil
}

type ListReportsFilters struct {
	Limit       *uint
	StandupID   *uint
	UserID      *uint
	After       *int64
	Before      *int64
	QuestionIDs []uint
	HTML        *bool
}

func (f ListReportsFilters) QueryString() string {
	qs := &strings.Builder{}

	if f.Limit != nil {
		qs.WriteString(fmt.Sprintf("limit=%d&", *f.Limit))
	}

	if f.StandupID != nil {
		qs.WriteString(fmt.Sprintf("standup_id=%d&", *f.StandupID))
	}

	if f.UserID != nil {
		qs.WriteString(fmt.Sprintf("user_id=%d&", *f.UserID))
	}

	if f.After != nil {
		qs.WriteString(fmt.Sprintf("after=%d&", *f.After))
	}

	if f.Before != nil {
		qs.WriteString(fmt.Sprintf("before=%d&", *f.Before))
	}

	if len(f.QuestionIDs) > 0 {
		for _, id := range f.QuestionIDs {
			qs.WriteString(fmt.Sprintf("question_ids[]=%d&", id))
		}
	}

	if f.HTML != nil {
		qs.WriteString(fmt.Sprintf("html=%t", *f.HTML))
	}

	return strings.TrimRight(qs.String(), "&")
}
