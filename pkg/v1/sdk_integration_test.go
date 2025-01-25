//go:build integration

package v1_test

import (
	"net/http"
	"os"
	"testing"

	v1 "github.com/omissis/go-geekbot-sdk/pkg/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSDK_ListTeams(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.geekbot.com/v1"
	apiKey, ok := os.LookupEnv("GEEKBOT_API_KEY")

	require.True(t, ok, "GEEKBOT_API_KEY environment variable is required")

	sdk := v1.NewSDK(httpClient, baseURL, apiKey)

	team, err := sdk.ListTeams()

	assert.NoError(t, err)
	assert.NotEmpty(t, team)
	assert.NotEmpty(t, team.ID)
	assert.NotEmpty(t, team.Name)
	assert.NotEmpty(t, team.Users)
}

func TestSDK_ListReports(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.geekbot.com/v1"
	apiKey, ok := os.LookupEnv("GEEKBOT_API_KEY")

	require.True(t, ok, "GEEKBOT_API_KEY environment variable is required")

	sdk := v1.NewSDK(httpClient, baseURL, apiKey)

	reports, err := sdk.ListReports(v1.ListReportsFilters{})

	assert.NoError(t, err)
	assert.NotEmpty(t, reports)
	assert.NotEmpty(t, reports[0].ID)
	assert.NotEmpty(t, reports[0].Questions)
}

func TestSDK_ListStandups(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.geekbot.com/v1"
	apiKey, ok := os.LookupEnv("GEEKBOT_API_KEY")

	require.True(t, ok, "GEEKBOT_API_KEY environment variable is required")

	sdk := v1.NewSDK(httpClient, baseURL, apiKey)

	standups, err := sdk.ListStandups()

	assert.NoError(t, err)
	assert.NotEmpty(t, standups)
	assert.NotEmpty(t, standups[0].ID)
	assert.NotEmpty(t, standups[0].Name)
	assert.NotEmpty(t, standups[0].Timezone)
	assert.NotEmpty(t, standups[0].Questions)
}
