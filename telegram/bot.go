package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "7788821483:AAGVLaYZzDggtBpIv5iy2Osz2KwD0E5PS9c"

func StartBot() {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Auth on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Task Manager Bot! Используйте команды:\n/create_task, /view_tasks, /update_task, /delete_task")
			bot.Send(msg)
			// case "create_task":
		//     handleCreateTask(bot, update, taskHandlers)
		// case "view_tasks":
		//     handleViewTasks(bot, update, taskHandlers)
		// case "update_task":
		//     handleUpdateTask(bot, update, taskHandlers)
		// case "delete_task":
		//     handleDeleteTask(bot, update, taskHandlers)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /start для списка команд.")
			bot.Send(msg)
		}
	}
}

func handleCreateTask() {

}

func handleViewTasks() {

}

func handleUpdateTask() {

}

func handleDeleteTask() {

}
