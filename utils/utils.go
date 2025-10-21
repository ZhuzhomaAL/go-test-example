package utils

import (
	"github.com/playwright-community/playwright-go"
	"log"
	"net/url"
	"os"
)

var PlaywrightBrowser playwright.Browser

func CreateContext() playwright.BrowserContext {
	context, err := PlaywrightBrowser.NewContext(playwright.BrowserNewContextOptions{
		Viewport: &playwright.Size{Width: 1920, Height: 1080},
	})
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	return context
}

func CreatePageWithURL(context playwright.BrowserContext, url string) playwright.Page {
	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("Failed to create new page: %v", err)
		return nil
	}

	if _, err = page.Goto(url); err != nil {
		page.Close()
		log.Fatalf("Failed to go to page: %v", err)
		return nil
	}

	return page
}

func GetURL(route string) (string, error) {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		log.Fatalf("BASE_URL variable is not set")
	}
	fullURL, err := url.JoinPath(baseURL, route)
	if err != nil {
		return "", err
	}

	return fullURL, nil
}
