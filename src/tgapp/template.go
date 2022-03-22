package tgapp

import (
	"concierge-bot/src/common"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	REG_STATUS_START     = "start"
	REG_STATUS_HOME      = "home"
	REG_STATUS_APARTMENT = "apartment"
	REG_STATUS_END       = "end"
)

const (
	COMMAND_INFO         = "info"
	COMMAND_INFO_HOMES   = "info_homes"
	COMMAND_HELP         = "help"
	COMMAND_START        = "start"
	COMMAND_REGISTRATION = "registration"
	COMMAND_ABOUTME      = "aboutme"
	COMMAND_APARTMENTS   = "apartments"
)

func getHelp() string {
	return "<b>Список команд:</b> \n" +
		"<i>Для всех пользователей</i> \n" +
		"/" + COMMAND_REGISTRATION + " - Регистрация (запустить процесс регистрации)\n" +
		"/" + COMMAND_INFO + " получить информацию по жк\n" +
		"<i>Для зарегестрированных</i> \n" +
		"/" + COMMAND_ABOUTME + " - посмотреть информацию о себе\n" +
		"/" + COMMAND_INFO_HOMES + " информация по домам жк\n" +
		"Так же это сообщение можно вызвать командой /" + COMMAND_HELP
}

func getHomes(homes []common.Home) string {
	result := "<b>Список домов: </b>\n"
	for _, value := range homes {
		result += value.Name + " получить информацию о доме /" + COMMAND_INFO + "_" + value.Name + "\n"
		result += "А так можно получить информацию о жильце квартиры (после : указать нужный номер) /" + COMMAND_INFO + "_" + value.Name + ":1" + "\n"
	}
	return result
}

/*
func regHomes(homes []common.Home) string {
	result := "Укажите в сообщение ваш дом:\n"
	for _, value := range homes {
		result += value.Name + "\n"
	}
	return result
}
*/
func regHomes() string {
	result := "Укажите в сообщение ваш дом:\n"
	return result
}

func regHomesButton(homes []common.Home) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, home := range homes {
		var row []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(home.Name, home.Name)
		row = append(row, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}
	return keyboard
	//msg.ReplyMarkup = keyboard
}

func aboutMe(resident common.Resident) string {
	return "<b>Данные обо мне:</b> \n" +
		"Telegramm - " + "@" + resident.Name + "\n" +
		"Дом - " + resident.Home + "\n" +
		"Квартира - " + fmt.Sprintf("%d", resident.Apartment) + "\n" +
		"Доп. Информация - " + resident.Description
}

func aboutResidents(residents []common.Resident) string {
	result := ""
	for _, resident := range residents {
		result += "<b>Данные жильца:</b> \n" +
			"Telegramm - " + "@" + resident.Name + "\n" +
			"Дом - " + resident.Home + "\n" +
			"Квартира - " + fmt.Sprintf("%d", resident.Apartment) + "\n" +
			"Доп. Информация - " + resident.Description + "\n"
	}
	return result
}

func aboutCoplex(complex common.ResidentilaComplex) string {
	return "<b>О комплексе:</b> \n" +
		"Название - " + complex.Name + "\n" +
		"Доп. Информация - " + complex.Description
}

func aboutHome(home common.Home) string {
	return "<b>Данные дома:</b> \n" +
		"Жк - " + home.ResidentialComplex + "\n" +
		"Дом - " + home.Name + "\n" +
		"Всего квартир - " + fmt.Sprintf("%d", home.Apartments) + "\n" +
		"Этажей - " + fmt.Sprintf("%d", home.Floors) + "\n" +
		"Первый жилой этаж - " + fmt.Sprintf("%d", home.FirstResidentialFloor) + "\n" +
		"Подъездов - " + fmt.Sprintf("%d", home.Entrances) + "\n" +
		"Доп. Информация - " + home.Description + "\n" +
		"Информация по квартирам - /" + COMMAND_APARTMENTS + "_" + home.Name
}

func apartmentsMap(home common.Home, apartments []int) string {
	result := "<b>Данные дома:</b> \n\n"

	for enrtance := 1; enrtance <= home.Entrances; enrtance++ {
		result += fmt.Sprintf("Подъезд - %d \n", enrtance)
		for floor := home.FirstResidentialFloor; floor <= home.Floors; floor++ {
			apartStart := (floor-home.FirstResidentialFloor)*home.Apartments + 1
			apartEnd := apartStart + home.Apartments
			apartList := fmt.Sprintf("Этаж %d: \n", floor)
			for apart := apartStart; apart < apartEnd; apart++ {
				if apartments[apart] == 0 {
					apartList += fmt.Sprintf(" %d", apart)
				} else {
					apartList += fmt.Sprintf(" <b>|%d|</b>", apart)
				}
			}
			apartList += "\n"
			result += apartList
		}
	}

	return result
}

func botAnswer() string {
	return fmt.Sprintf("Общение с ботами не поддерживается")
}

func notSupport() string {
	return "Не понимаю, уточнить список команд можно /" + COMMAND_HELP
}

func errorMessage() string {
	return "Что, то пошло не так"
}

func notFound() string {
	return "Переданное значение не найдено"
}

func notRegistred() string {
	return "Для доступа к этой команде нужно пройти регистрацию /" + COMMAND_REGISTRATION
}
