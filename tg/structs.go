package tg

import "encoding/json"

type Response struct {
	Ok        bool            `json:"ok"`
	Result    json.RawMessage `json:"result"`
	ErrorCode int             `json:"error_code"`
}

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

// Message описывает сообщение Telegram.
type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      *Chat  `json:"chat"`
	User      *User  `json:"from"`
}

// Chat описывает чат в системе Telegram.
type Chat struct {
	ID int `json:"id"`
}

// User описывает оправителя сообщения в Telegram.
type User struct {
	ID           uint64 `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`   // optional
	OneTimeKeyboard bool               `json:"one_time_keyboard"` // optional
	Selective       bool               `json:"selective"`         // optional
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type SendingMessage struct {
	ChatID      int                  `json:"chat_id"`
	Text        string               `json:"text"`
	ParseMode   string               `json:"parse_mode"`
	ReplyMarkup *ReplyKeyboardMarkup `json:"reply_markup"`
}
