package tgapp

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgApp struct {
	DB    DBConnector
	Token string
}

func (t TgApp) Run() {
	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		log.Fatalln("TG:Run %w ", err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			if update.Message.From.IsBot {
				msg.Text = botAnswer()
			} else if update.Message.IsCommand() {
				msg = t.commandController(msg, update.Message)
			} else {
				msg = t.otherController(msg, update.Message)
			}
			msg.ParseMode = "html"
			if _, err := bot.Send(msg); err != nil {
				log.Println("TG:RUN Send message %w ", err)
			}
		} else if update.CallbackQuery != nil {
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			msg = t.collbackController(msg, update.CallbackQuery)
			msg.ParseMode = "html"
			if _, err := bot.Send(msg); err != nil {
				log.Println("TG:RUN Send message %w ", err)
			}
		}
	}
}

func (t TgApp) collbackController(msg tgbotapi.MessageConfig, update *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	err := t.DB.UpdateResidentAddHome(
		fmt.Sprintf("%d", update.From.ID),
		update.Data,
		REG_STATUS_HOME,
	)
	if err != nil {
		log.Println("TG:otherController Failed update home %w ", err)
		msg.Text = errorMessage()
		return msg
	}
	msg.Text = "Теперь укажите номер квартиры"
	return msg
}

func (t TgApp) commandController(msg tgbotapi.MessageConfig, update *tgbotapi.Message) tgbotapi.MessageConfig {
	fmt.Println(update.Command())
	if update.Command() == COMMAND_START || update.Command() == COMMAND_HELP {
		msg.Text = getHelp()
		return msg
	}

	if update.Command() == COMMAND_REGISTRATION {
		err := t.DB.CreateResidentBase(fmt.Sprintf("%d", update.From.ID),
			update.From.UserName,
			REG_STATUS_START,
		)
		if err != nil {
			log.Println("TG:otherController Failed create/update resident %w ", err)
			msg.Text = errorMessage()
			return msg
		}

		homes, err := t.DB.GetHomes()
		if err != nil {
			log.Println("TG:otherController Failed get homes %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		//msg.Text = regHomes(homes)
		msg.Text = regHomes()
		msg.ReplyMarkup = regHomesButton(homes)
		return msg
	}

	if update.Command() == COMMAND_INFO {
		complex, err := t.DB.GetResidentilaComplex()
		if err != nil {
			log.Println("TG:commandController Failed get complex %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		msg.Text = aboutCoplex(complex)
		return msg
	}

	if !t.checkState(fmt.Sprintf("%d", update.From.ID)) {
		msg.Text = notRegistred()
		return msg
	}

	if update.Command() == COMMAND_INFO_HOMES {
		homes, err := t.DB.GetHomes()
		if err != nil {
			log.Println("TG:commandController Failed get homes %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		msg.Text = getHomes(homes)
		return msg
	}

	if update.Command() == COMMAND_ABOUTME {
		resident, err := t.DB.GetResident(fmt.Sprintf("%d", update.From.ID))
		if err != nil {
			log.Println("TG:commandController Failed get info %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		msg.Text = aboutMe(resident)
		return msg
	}

	homes, err := t.DB.GetHomes()

	if err != nil {
		log.Println("TG:commandController Failed get homes %w ", err)
		msg.Text = errorMessage()
		return msg
	}

	for _, home := range homes {
		if update.Command() == COMMAND_INFO+"_"+home.Name {
			splitText := strings.Split(update.Text, ":")
			if len(splitText) > 1 {
				converApartment, err := strconv.Atoi(splitText[1])
				if err != nil {
					log.Println(fmt.Sprintf("TG:commandController Apartment not convert %s ", splitText[1]))
					msg.Text = errorMessage()
					return msg
				}
				residents, err := t.DB.GetResidentsByHomeAndApartment(home.Name, converApartment)

				if err != nil {
					log.Println("TG:commandController Resident not found")
					msg.Text = notFound()
					return msg
				}

				msg.Text = aboutResidents(residents)
				return msg
			} else {
				msg.Text = aboutHome(home)
				return msg
			}
		}
	}

	msg.Text = notSupport()
	return msg
}

func (t TgApp) otherController(msg tgbotapi.MessageConfig, update *tgbotapi.Message) tgbotapi.MessageConfig {
	resident, err := t.DB.GetResident(fmt.Sprintf("%d", update.From.ID))
	if err != nil {
		log.Println("TG:otherController Failed check %w ", err)
		return msg
	}

	if resident.Status == REG_STATUS_START {
		home, err := t.DB.GetHome(update.Text)
		if err != nil {
			log.Println("TG:otherController Failed get home %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		if home.Name == "" {
			log.Println(fmt.Sprintf("TG:otherController Home not found %s ", update.Text))
			msg.Text = notFound()
			return msg
		}
		err = t.DB.UpdateResidentAddHome(
			fmt.Sprintf("%d", update.From.ID),
			update.Text,
			REG_STATUS_HOME,
		)
		if err != nil {
			log.Println("TG:otherController Failed update home %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		msg.Text = "Теперь укажите номер квартиры"
		return msg
	} else if resident.Status == REG_STATUS_HOME {
		converApartment, err := strconv.Atoi(update.Text)
		if err != nil {
			log.Println(fmt.Sprintf("TG:otherController Apartment not convert %s ", update.Text))
			msg.Text = notFound()
			return msg
		}

		err = t.DB.UpdateResidentAddApartment(fmt.Sprintf("%d", update.From.ID),
			converApartment,
			REG_STATUS_APARTMENT,
		)
		if err != nil {
			log.Println("TG:otherController Failed update apartment %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		msg.Text = "Осталось внести информацию о себе. Впишите, что считаете нужным."
		return msg
	} else if resident.Status == REG_STATUS_APARTMENT {

		err = t.DB.UpdateResidentAddDescription(fmt.Sprintf("%d", update.From.ID),
			update.Text,
			REG_STATUS_END,
		)
		if err != nil {
			log.Println("TG:otherController Failed update description %w ", err)
			msg.Text = errorMessage()
			return msg
		}
		msg.Text = "Регистрация завершена, если хотите обновить данные запустить регистрацию можно повторно /" + COMMAND_REGISTRATION
		return msg

	}
	msg.Text = notSupport()
	return msg
}

func (t TgApp) checkState(id string) bool {
	resident, err := t.DB.GetResident(id)
	if err != nil {
		log.Println("TG:checkState Failed check %w ", err)
		return false
	}
	if resident.Status == REG_STATUS_END {
		return true
	} else {
		return false
	}
}
