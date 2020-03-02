package config

import "encoding/json"

// WebConfig Config
type WebConfig struct {
	Router string   `json:"router"`
	Type   string   `json:"type"`
	Resp   *WebResp `json:"resp"`
}

// WebResp WebResp
type WebResp struct {
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	File    string            `json:"file"`
	Body    json.RawMessage   `json:"body"`
}
