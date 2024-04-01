package main

import (
	"log"
	"testing"

	"github.com/kwhitaker/go-htmx-templ/models"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var todoName = "Test Todo"

// This giant test function is not ideal,
// but I didn't want to do more work because this isn't a real app
func TestMain(t *testing.T) {
	t.Setenv("MODE", "test")

	db, err := gorm.Open(sqlite.Open("test.db"), nil)

	require.NoError(t, err)

	db.AutoMigrate(&models.Todo{})

	t.Cleanup(func() {
		db.Exec("DROP TABLE todos")
		db.AutoMigrate(&models.Todo{})
	})

	expect := playwright.NewPlaywrightAssertions(5000)
	err = playwright.Install()

	if err != nil {
		log.Fatalf("Failed to install Playwright: %v", err)
	}

	pw, err := playwright.Run()

	require.NoError(t, err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})

	require.NoError(t, err)

	context, err := browser.NewContext()

	require.NoError(t, err)

	page, err := context.NewPage()

	require.NoError(t, err)

	_, err = page.Goto("http://127.0.0.1:3000")

	require.NoError(t, err)

	require.NoError(t, expect.Page(page).ToHaveTitle("Sample To-Dos"))

	err = expect.Locator(page.GetByPlaceholder("Enter new to-do...")).ToBeFocused()

	require.NoError(t, err)

	// CREATE
	err = page.GetByPlaceholder("Enter new to-do...").Fill(todoName)

	require.NoError(t, err)

	err = page.Keyboard().Press("Enter")

	require.NoError(t, err)

	err = expect.Locator(page.GetByTestId("todo-1")).ToBeVisible()

	require.NoError(t, err)

	// TOGGLE
	checkbox := page.GetByTestId("todo-1-toggle")

	err = checkbox.Click()

	require.NoError(t, err)

	err = expect.Locator(checkbox).ToBeChecked()

	require.NoError(t, err)

	// UPDATE
	input := page.GetByTestId("todo-1-input")

	err = input.Fill("Updated Test Todo")

	require.NoError(t, err)

	err = input.Blur()

	require.NoError(t, err)

	err = expect.Locator(input).ToHaveValue("Updated Test Todo")

	require.NoError(t, err)

	toast := page.GetByTestId("todo-1-toast")

	err = expect.Locator(toast).ToBeVisible()

	require.NoError(t, err)

	// DELETE
	delBtn := page.GetByTestId("todo-1-delete")

	err = delBtn.Click()

	require.NoError(t, err)

	err = expect.Locator(page.GetByTestId("todo-1")).ToBeHidden()

	require.NoError(t, err)
}
