package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"go-test-example/utils"
)

func TestMain(m *testing.M) {

	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println(".env file not found, using system environment")
	}

	err := playwright.Install()
	if err != nil {
		log.Fatalf("could not install playwright: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	utils.PlaywrightBrowser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Args:     []string{"--disable-blink-features=AutomationControlled"},
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	status := m.Run()

	utils.PlaywrightBrowser.Close()
	pw.Stop()
	os.Exit(status)
}
