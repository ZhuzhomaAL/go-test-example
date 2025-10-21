package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
	"go-test-example/utils"
)

func TestClientServerDataSync(t *testing.T) {
	const expStr = "Apple Mac Studio 2023 M2 Max - 30 Core GPU / 12 / 64 ГБ / 512 ГБ Серебро (Silver)"

	URL, err := utils.GetURL("/product/3335")
	require.NoError(t, err, "Failed to get URL: %w", err)
	require.NotEmpty(t, URL, "URL is empty")

	ctx := utils.CreateContext()
	defer ctx.Close()

	var requests []string
	ctx.On("request", func(req playwright.Request) {
		requests = append(requests, req.URL())
	})

	page := utils.CreatePageWithURL(ctx, URL)
	htmlContent, err := page.Content()
	require.NoError(t, err, "Failed to get page content: %w", err)
	require.NotEmpty(t, htmlContent)
	require.Contains(t, htmlContent, expStr)

	for _, url := range requests {
		require.False(t, strings.Contains(url, "/v1/products/"), "Request to: %s detected", url)
	}

	resp, err := http.Get(URL)
	require.NoError(t, err, t, err, "Failed to perform GET request to %s: %w", URL, err)
	require.Equal(t, http.StatusOK, resp.StatusCode,
		"GET request status: %d, expected: %d", resp.StatusCode, http.StatusOK)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Failed to read response body: %w", err)

	bodyString := string(body)
	require.Contains(t, bodyString, expStr, "Page does not contain expected string: %s", expStr)
}
