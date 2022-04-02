package processor

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/lumos/processor"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

var bot, _ = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
var _, _ = Notify()

func Notify() (tgbotapi.Message, error) {
	var msg = tgbotapi.NewMessage(1929798658, "testing testing")
	return bot.Send(msg)
}

func Reply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error updating →", err)
	}

	log.Printf("[%s@%d] %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)

	reply := processor.LookupReply(update)

	data := Response{
		Msg:    reply,
		Method: "sendMessage",
		ChatID: update.Message.Chat.ID,
	}

	message, _ := json.Marshal(data)
	log.Printf("Response %s", string(message))
	fmt.Fprintf(w, string(message))

}
