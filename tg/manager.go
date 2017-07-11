package tg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Manager - менеджер общения с Telegram Bot API.
type Manager interface {
	GetMessagesChan() (<-chan *Message, error)
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
