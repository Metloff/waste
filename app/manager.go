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
	_, err = m.dbs.CreateTransaction(user.ID, amount, result["title"], result["category"])
	if err != nil {
		// отправить ошибку (something wrong)
		log.Println(err)
	}

	log.Println(msg)
	return nil
}
