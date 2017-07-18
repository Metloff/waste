package app

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/satori/go.uuid"
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
	err1                    = "Данные введены не корректно. Попробуйте пожалуйста еще раз"
	err2                    = "Что то пошло не так. Попробуйте пожалуйста еще раз"
	succ                    = "Информация успешно сохранена"
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

	re, _ = regexp.Compile("(?P<amount>\\d{0,})(?P<title>.*)#(?P<category>.*)?")
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

	match := re.FindStringSubmatch(msg.Text)

	switch {
	case msg.Text == "/start":
		m.tg.SendMessage(msg.Chat.ID, "Хаю хай", "Markdown", &startKeyboard)
	case match != nil:
		message := m.addWaste(match, msg.User)
		m.tg.SendMessage(msg.Chat.ID, message, "Markdown", &startKeyboard)
	case msg.Text == graph:
		message := m.generateLink(msg.User.ID)
		m.tg.SendMessage(msg.Chat.ID, message, "Markdown", &startKeyboard)
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

	return nil
}

func (m *manager) giveCurrentMonthStatisticList(telegramID uint64) (message string) {
	results := m.dbs.CurrentMonthStatistic(telegramID)
	if len(results) == 0 {
		return "У вас нет трат за этот период"
	}

	for _, res := range results {
		message += fmt.Sprintf("%s | Кол-во трат: %v | ₽: %v", res.Category, res.Count, res.Amount)
		message += "\n ---------------\n\n"
	}

	return message
}

func (m *manager) givePreviousMonthStatisticList(telegramID uint64) (message string) {
	results := m.dbs.PreviousMonthStatistic(telegramID)
	if len(results) == 0 {
		return "У вас нет трат за этот период"
	}

	for _, res := range results {
		message += fmt.Sprintf("%s | Кол-во трат: %v | ₽: %v", res.Category, res.Count, res.Amount)
		message += "\n ---------------\n\n"
	}

	return message
}

func (m *manager) addWaste(match []string, tgUser *tg.User) string {
	// Добавить валидацию полей
	// Добавить приведение категории и названия к общему виду (с большой буквы)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}
	fmt.Printf("Amount: %s  Tilte: %s Category: %s \n", result["amount"], result["title"], result["category"])

	user, err := m.dbs.FindOrCreateUser(tgUser.ID, tgUser.FirstName, tgUser.LastName, tgUser.LanguageCode)
	if err != nil {
		log.Println(err)
		return err2
	}

	amount, err := strconv.ParseUint(result["amount"], 10, 64)
	if err != nil {
		log.Println(err)
		return err1
	}
	_, err = m.dbs.CreateTransaction(user.ID, user.TelegramID, amount, result["title"], result["category"])
	if err != nil {
		// отправить ошибку (something wrong)
		log.Println(err)
		return err2
	}

	return succ
}

func (m *manager) generateLink(tid uint64) (message string) {
	// Creating UUID Version 4
	uuid := uuid.NewV4().String()

	_, err := m.dbs.SetUUID(tid, uuid)
	if err != nil {
		log.Println(err)
	}

	// message = fmt.Sprintf("http://127.0.0.1:3000/%s", uuid)
	message = fmt.Sprintf("http://192.168.14.195:3000/%s", uuid)
	return message
}
