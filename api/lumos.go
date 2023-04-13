package processor

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/lumos/common"
	"github.com/margostino/lumos/db"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const VarPrefix = "var_"

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

var bot, _ = newBot()

var variable *db.Variable

func Reply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var update tgbotapi.Update
	var reply string

	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error updating →", err)
	}

	input := update.Message.Text

	log.Printf("[%s@%d] %s", update.Message.From.UserName, update.Message.Chat.ID, input)

	if input == "/start" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to Lumos!")
		btn := tgbotapi.KeyboardButton{
			RequestLocation: true,
			Text:            "Track location",
		}
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
		bot.Send(msg)
	} else if strings.HasPrefix(input, VarPrefix) {
		variable.Datetime = time.Now().UTC().String()
		variable.Name = strings.ReplaceAll(input, VarPrefix, "")

		reply = fmt.Sprintf("📍  Latitude: %f\n"+
			"📍  Longitude: %f\n"+
			"⚡️  Variable: %s\n"+
			"🔎  Send the value and observation (optional).\n"+
			"➡️  Format: {value};{observation}\n",
			variable.Latitude,
			variable.Longitude,
			variable.Name)

	} else if update.Message.Location != nil {
		variableNames := db.GetVariableNames()

		latitude := update.Message.Location.Latitude
		longitude := update.Message.Location.Longitude
		variable = &db.Variable{
			Latitude:  latitude,
			Longitude: longitude,
		}

		message := fmt.Sprintf("📍  Latitude: %f\n"+
			"📍  Longitude: %f\n"+
			"🔎  Pick the variable or send a new one.\n"+
			"➡️  Format (new): {variable_name};{value};{observation}\n",
			latitude,
			longitude)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		buttons := make([]tgbotapi.KeyboardButton, 0)

		for _, variableName := range variableNames {
			button := tgbotapi.KeyboardButton{
				Text: fmt.Sprintf("%s%s", VarPrefix, variableName),
			}
			buttons = append(buttons, button)
		}

		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)
		bot.Send(msg)

	} else if variable != nil {
		normalizedInputList := common.NewString(input).
			ToLower().
			Split(";").
			Values()

		if len(normalizedInputList) == 3 && variable.Name == "" {
			variable.Datetime = time.Now().UTC().String()
			variable.Name = normalizedInputList[0]
			variable.Value = normalizedInputList[1]
			variable.Observation = normalizedInputList[2]
		} else if len(normalizedInputList) == 2 && variable.Name != "" {
			variable.Value = normalizedInputList[0]
			variable.Observation = normalizedInputList[1]
		} else {
			reply = "🚫  Invalid input"
		}

		if isFull(variable) {
			err := db.Append(variable)
			variable = nil

			if err != nil {
				reply = fmt.Sprintf("🛑  Unable to save data: %s", err.Error())
			} else {
				reply = "✅  Data recorded successfully"
			}
		}

	}

	data := Response{
		Msg:    reply,
		Method: "sendMessage",
		ChatID: update.Message.Chat.ID,
	}

	message, _ := json.Marshal(data)
	log.Printf("Response %s", string(message))
	fmt.Fprintf(w, string(message))

	//reply := processor.LookupReply(update)
	//
	//data := Response{
	//	Msg:    reply,
	//	Method: "sendMessage",
	//	ChatID: update.Message.Chat.ID,
	//}
	//
	//message, _ := json.Marshal(data)
	//log.Printf("Response %s", string(message))
	//fmt.Fprintf(w, string(message))

}

func newBot() (*tgbotapi.BotAPI, error) {
	client, error := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	//bot.Debug = true
	common.SilentCheck(error, "when creating a new BotAPI instance")
	//log.Printf("Authorized on account %s\n", client.Self.UserName)
	return client, error
}

func isFull(variable *db.Variable) bool {
	return variable.Name != "" && variable.Value != "" && variable.Datetime != "" && variable.Latitude > 0 && variable.Longitude > 0 && variable.Observation != ""
}
