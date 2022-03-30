package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/lumos/datasource"
	"github.com/margostino/lumos/helpers"
	"github.com/margostino/lumos/loader"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

var countryMapping = loader.Load()

func Reply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//countryMapping := loader.Load()
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error updating →", err)
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	reply := lookupReply(update)

	data := Response{
		Msg:    reply,
		Method: "sendMessage",
		ChatID: update.Message.Chat.ID,
	}

	message, _ := json.Marshal(data)
	log.Printf("Response %s", string(message))
	fmt.Fprintf(w, string(message))

}

func lookupReply(update tgbotapi.Update) string {
	input := strings.ToLower(update.Message.Text)

	if input != "sweden" { // strategy logic tbd
		return helpers.Greeting()
	}
	return datasource.GetData(countryMapping["sweden"])
}
