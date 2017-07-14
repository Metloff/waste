package app

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/wastebot/dbs"
	"github.com/wastebot/tg"
)

const (
	graph                   = "📈 График трат"
	list                    = "🗒 Список трат"
	listThisMonthString     = "Список за текущий месяц"
	listPreviousMonthString = "Список за прошлый месяц"
	info                    = "📍 Инфо"
	goBackString            = "🔙 Назад"
)

var (
	nilKeyboard = tg.ReplyKeyboardMarkup{
		Keyboard: [][]tg.KeyboardButton{},
	}

	startKeyboard = tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: graph}, {Text: list}},
			{{Text: info}},
		},
	}

	listKeyboard = tg.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]tg.KeyboardButton{
			{{Text: listThisMonthString}, {Text: listPreviousMonthString}},
			{{Text: info}, {Text: goBackString}},
		},
	}
)

type manager struct {
	dbs dbs.Manager
	tg  tg.Manager
}

type Manager interface {
	Listen()
}

// NewManager - ...
func NewManager(dbs dbs.Manager, tg tg.Manager) Manager {
	return &manager{
		dbs: dbs,
		tg:  tg,
	}
}

// Listen ...
func (m *manager) Listen() {
	log.Println("Listening...")

	msgChan, _ := m.tg.GetMessagesChan()
	for msg := range msgChan {
		go m.handleMsg(msg)
	}
}

func (m *manager) handleMsg(msg *tg.Message) error {
	if msg == nil {
		return errors.New("Nil pointer on msg")
	}

	switch {
	case msg.Text == "/start":
		m.tg.SendMessage(msg.Chat.ID, "Хаю хай", "Markdown", &startKeyboard)
	case msg.Text == list:
		m.tg.SendMessage(msg.Chat.ID, "Выберите промежуток", "Markdown", &listKeyboard)
	case msg.Text == listThisMonthString:
		message := m.giveCurrentMonthStatisticList(msg.User.ID)
		m.tg.SendMessage(msg.Chat.ID, message, "Markdown", &listKeyboard)
	case msg.Text == listPreviousMonthString:
		message := m.givePreviousMonthStatisticList(msg.User.ID)
		m.tg.SendMessage(msg.Chat.ID, message, "Markdown", &listKeyboard)
	case msg.Text == goBackString:
		message := fmt.Sprintf("Чего изволите, %s %s?", msg.User.FirstName, msg.User.LastName)
		m.tg.SendMessage(msg.Chat.ID, message, "Markdown", &startKeyboard)
	default:
		message := fmt.Sprintf("Чего изволите, %s %s?", msg.User.FirstName, msg.User.LastName)
		m.tg.SendMessage(msg.Chat.ID, message, "Markdown", &startKeyboard)
	}

	re, _ := regexp.Compile("(?P<amount>\\d{0,})(?P<title>.*)#(?P<category>.*)?")
	match := re.FindStringSubmatch(msg.Text)
	fmt.Println(match)
	if match == nil {
		log.Println("sdf")
		return errors.New("olol")
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	fmt.Printf("Amount: %s  Tilte: %s Category: %s \n", result["amount"], result["title"], result["category"])
	log.Println(msg.User)

	user, err := m.dbs.FindOrCreateUser(msg.User.ID, msg.User.FirstName, msg.User.LastName, msg.User.LanguageCode)
	if err != nil {
		// попросить юзера отправить данные еще раз
		log.Println(err)
	}

	amount, err := strconv.ParseUint(result["amount"], 10, 64)
	if err != nil {
		// попросить юзера отправить данные еще раз
		log.Println(err)
	}
	_, err = m.dbs.CreateTransaction(user.ID, user.TelegramID, amount, result["title"], result["category"])
	if err != nil {
		// отправить ошибку (something wrong)
		log.Println(err)
	}

	log.Println(msg)
	return nil
}

func (m *manager) giveCurrentMonthStatisticList(telegramID uint64) string {
	results := m.dbs.CurrentMonthStatistic(telegramID)
	if len(results) == 0 {
		return "У вас нет трат за этот период"
	}

	var message string
	for _, res := range results {
		message += fmt.Sprintf("%s | Кол-во трат: %v | ₽: %v", res.Category, res.Count, res.Amount)
		message += "\n ---------------\n\n"
	}

	return message
}

func (m *manager) givePreviousMonthStatisticList(telegramID uint64) string {
	results := m.dbs.PreviousMonthStatistic(telegramID)
	if len(results) == 0 {
		return "У вас нет трат за этот период"
	}

	var message string
	for _, res := range results {
		message += fmt.Sprintf("%s | Кол-во трат: %v | ₽: %v", res.Category, res.Count, res.Amount)
		message += "\n ---------------\n\n"
	}

	return message
}
