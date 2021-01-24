package telegram

import (
	"github.com/yanzay/tbot"
)

var today = tbot.KeyboardButton{
	Text:            "📗 Стих на cегодня",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var tomorrow = tbot.KeyboardButton{
	Text:            "📘 Стих на завтра",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var hostToday = tbot.KeyboardButton{
	Text:            "🙋 Сегодня ведет",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

var hostTomorrow = tbot.KeyboardButton{
	Text:            "🙋 ‍Завтра ведет",
	RequestContact:  false,
	RequestLocation: false,
	RequestPoll:     nil,
}

func MenuButtons() *tbot.ReplyKeyboardMarkup {
	return &tbot.ReplyKeyboardMarkup{
		Keyboard: [][]tbot.KeyboardButton{
			{today, tomorrow, hostToday, hostTomorrow},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}
