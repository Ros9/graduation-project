package telegram

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text  string      `json:"text"`
	From  From        `json:"from"`
	Chat  Chat        `json:"chat"`
	Photo []PhotoItem `json:"photo"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type PhotoItem struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FileSize     float32 `json:"file_size"`
	Width        float32 `json:"width"`
	Height       float32 `json:"height"`
}

type GetFileReq struct {
	FileId string `json:"file_id"`
}

type GetFileResp struct {
	Ok     bool            `json:"ok"`
	Result GetFileRespInfo `json:"result"`
}

type GetFileRespInfo struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FileSize     float32 `json:"file_size"`
	FilePath     string  `json:"file_path"`
}

// 		   "file_id": "AgACAgIAAxkBAAIDAWKTrRNPJYbzmf0IdBU2oipO0xZVAAIdvjEbcjOZSMWmhJfce_kAAQEAAwIAA3kAAyQE",
//         "file_unique_id": "AQADHb4xG3IzmUh-",
//         "file_size": 74635,
//         "file_path": "photos/file_1.jpg"
