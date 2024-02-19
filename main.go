package main

import (
	"github.com/Zeusina/khlebnik-bot/telegram"
	"github.com/Zeusina/khlebnik-bot/utils"
)

func main() {
	utils.GetEnv()
	utils.ConfigureLogs()
	telegram.StartBot()
}
