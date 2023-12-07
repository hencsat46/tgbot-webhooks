package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"tgbot/internal/messages"
	"tgbot/internal/models"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	Download([]models.Photo) error
	Upload(models.Photo) error
	Commands(string) (int, error)
}

func New(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) RecieveMessage(ctx echo.Context) error {
	var body models.RequestBody

	if err := ctx.Bind(&body); err != nil {
		log.Println(err)
		return err
	}
	log.Println("всё ок")
	log.Println(body.MessageInfo.Text)
	//h.usecase.Download(body.MessageInfo.PhotosInfo)
	//h.SendMessage(body.MessageInfo.ChatInfo.ChatId, body.MessageInfo.Text)
	return nil
}

func (h *handler) Commands(chatId int64, text string) {
	action, err := h.usecase.Commands(text)
	if err != nil {
		log.Println(err)
		return
	}

	switch action {
	case 1:

		message := messages.Profile(chatId)
		h.SendMessage(chatId)
	case 2:

	case 3:

	}
}

func (h *handler) SendMessage(chatId int64, text string) {
	requestBody := &models.ResponseBody{ChatID: chatId, Text: text}
	log.Println(requestBody)
	encodedBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("Marshalling error: ", err)

		return
	}
	log.Println(os.Getenv("TOKEN") + "/sendMessage")
	response, err := http.Post(os.Getenv("TOKEN")+"/sendMessage", "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}

func (h *handler) Routes(e *echo.Echo) {
	e.POST("/", h.RecieveMessage)
}

func (h *handler) IdentifyMessageType(message models.Message) {
	if message.Text != "" {
		if message.Text[0] == '/' {

		}

	} else if message.StickerInfo.FileId != "" {

	} else if len(message.PhotosInfo) != 0 {
		h.usecase.Download(message.PhotosInfo)
	}
}
