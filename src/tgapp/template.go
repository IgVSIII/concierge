package tgapp

import (
	"concierge-bot/src/common"
	"fmt"
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
)

func getHelp() string {
	return "\\**Список команд:\\**\n" +
		"__Для всех пользователей__\n" +
		"/" + COMMAND_REGISTRATION + " - Регистрация (запустить процесс регистрации)\n" +
		"/" + COMMAND_INFO + " получить информацию по жк\n" +
		"__Для зарегестрированных__\n" +
		"/" + COMMAND_ABOUTME + " - посмотреть информацию о себе\n" +
		"/" + COMMAND_INFO_HOMES + " информация по домам жк\n" +
		"Так же это сообщение можно вызвать командой /" + COMMAND_HELP
}

func getHomes(homes []common.Home) string {
	result := "Список домов:\n"
	for _, value := range homes {
		result += value.Name + " получить информацию о доме /" + COMMAND_INFO + "_" + value.Name + "\n"
		result += "А так можно получить информацию о жильце квартиры (после : указать нужный номер) /" + COMMAND_INFO + "_" + value.Name + ":1" + "\n"
	}
	return result
}

func regHomes(homes []common.Home) string {
	result := "Укажите в сообщение ваш дом:\n"
	for _, value := range homes {
		result += value.Name + "\n"
	}
	return result
}

func aboutMe(resident common.Resident) string {
	return "Данные обо мне:\n" +
		"Telegramm - " + "@" + resident.Name + "\n" +
		"Дом - " + resident.Home + "\n" +
		"Квартира - " + fmt.Sprintf("%d", resident.Apartment) + "\n" +
		"Доп. Информация - " + resident.Description
}

func aboutResidents(residents []common.Resident) string {
	result := ""
	for _, resident := range residents {
		result += "Данные жильца:\n" +
			"Telegramm - " + "@" + resident.Name + "\n" +
			"Дом - " + resident.Home + "\n" +
			"Квартира - " + fmt.Sprintf("%d", resident.Apartment) + "\n" +
			"Доп. Информация - " + resident.Description + "\n"
	}
	return result
}

func aboutCoplex(complex common.ResidentilaComplex) string {
	return "О комплексе:\n" +
		"Название - " + complex.Name + "\n" +
		"Доп. Информация - " + complex.Description
}

func aboutHome(home common.Home) string {
	return "Данные дома:\n" +
		"Жк - " + home.ResidentialComplex + "\n" +
		"Дом - " + home.Name + "\n" +
		"Всего квартир - " + fmt.Sprintf("%d", home.Apartments) + "\n" +
		"Этажей - " + fmt.Sprintf("%d", home.Floors) + "\n" +
		"Первый жилой этаж - " + fmt.Sprintf("%d", home.FirstResidentialFloor) + "\n" +
		"Подъездов - " + fmt.Sprintf("%d", home.Entrances) + "\n" +
		"Доп. Информация - " + home.Description
}

func botAnswer() string {
	return "Общение с ботами не поддерживается"
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
