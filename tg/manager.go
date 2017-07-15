package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Manager - менеджер общения с Telegram Bot API.
type Manager interface {
	GetMessagesChan() (<-chan *Message, error)
	SendMessage(chatID int, text, parse string,
		keyboard *ReplyKeyboardMarkup) (*Response, error)
}

type manager struct {
	token     string
	lastUpdID int
}

// NewManager - ...
func NewManager(token string) Manager {
	return &manager{token: token}
}

func (m *manager) makeGETRequest(uri string) (*Response, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetMessagesChan возвращает канал с сообщениями.
func (m *manager) GetMessagesChan() (<-chan *Message, error) {
	msgChan := make(chan *Message)

	go func() {
		for {
			uri := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=100&offset=%d",
				m.token, m.lastUpdID+1)
			resp, _ := m.makeGETRequest(uri)

			var updates []*Update
			err := json.Unmarshal(resp.Result, &updates)
			// log.Println(updates[0])

			if err != nil {
				panic(err)
			}

			for _, upd := range updates {
				m.lastUpdID = upd.UpdateID
				msgChan <- upd.Message
			}
		}
	}()

	return msgChan, nil
}

func (m *manager) SendMessage(chatID int, text, parse string,
	keyboard *ReplyKeyboardMarkup) (*Response, error) {
	uri := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", m.token)
	var msg SendingMessage

	msg = SendingMessage{
		ChatID:      chatID,
		Text:        text,
		ParseMode:   parse,
		ReplyMarkup: keyboard,
	}

	resp, err := m.makePOSTRequest(uri, msg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(resp)
	if !resp.Ok {
		err = fmt.Errorf("Tg error %d", resp.ErrorCode)
		return nil, err
	}

	return resp, nil
}

func (m *manager) makePOSTRequest(uri string, fields SendingMessage) (*Response, error) {

	client := http.Client{}

	jsonStr, err := json.Marshal(fields)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(jsonStr))
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))
	result := &Response{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
