package models

type RequestBody struct {
	MessageInfo Message `json:"message"`
}

type Message struct {
	Text        string  `json:"text"`
	ChatInfo    Chat    `json:"chat"`
	StickerInfo Sticker `json:"sticker"`
	PhotosInfo  []Photo `json:"photo"`
}

type Chat struct {
	ChatId int64 `json:"id"`
}

type Sticker struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	StickerType  string `json:"type"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	IsAnimated   bool   `json:"is_animated"`
	IsVideo      bool   `json"is_video"`
}

type Photo struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type ResponseBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
