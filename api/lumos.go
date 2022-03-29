package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func Reply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error updating →", err)
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	filename := pwd + "/_files/sweden.json"
	fmt.Println("DATA PATH:  " + filename)

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}
	var fileData interface{}
	json.Unmarshal(byteValue, &fileData)

	text := "🪄   Happiness can be found, even in the darkest of times, if one only remembers to turn on the light.\n" +
		"🌎   We do not need magic to transform our world.\n" + filename
	data := Response{Msg: text,
		Method: "sendMessage",
		ChatID: update.Message.Chat.ID}

	msg, _ := json.Marshal(data)
	log.Printf("Response %s", string(msg))
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(msg))

}
