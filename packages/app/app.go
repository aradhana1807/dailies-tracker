package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
)

// App struct
type App struct {
	dailiesService DailiesService
	ctx            context.Context
	config         *Config
}

// NewApp creates a new App application struct
func NewApp(config *Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.dailiesService = DailiesService{}
}

func (a *App) LoadCommissions(dateMs int64) []Commission {
	dailies, err := a.dailiesService.LoadCommissionsForDate(dateMs)
	if err != nil {
		fmt.Println("Error loading dailies:", err)
		return nil
	}
	fmt.Println("Loaded dailies:", dailies)
	return dailies
}

func (a *App) CreateTask(description string, realm string, rewardsJson string) (Commission, error) {
	return a.dailiesService.CreateNewCommission(description, realm, rewardsJson)
}

func (a *App) CompleteTask(id int) error {
	return a.dailiesService.CompleteCommission(id)
}

func (a *App) DeleteTask(id int) error {
	return a.dailiesService.DeleteCommission(id)
}

func (a *App) IsDev() bool {
	exeName, err := os.Executable()
	if err == nil && strings.Contains(exeName, "-dev") {
		return true
	}
	return len(os.Getenv("DEBUG")) != 0 || len(os.Getenv("DEV")) != 0
}

func (a *App) GetLocale() string {
	return a.config.Locale
}

func (a *App) Notify(title string, message string) {
	err := beeep.Notify(title, message, "public/logo-universal.png")
	if err != nil {
		fmt.Println("Error notifying:", err)
	}
}
