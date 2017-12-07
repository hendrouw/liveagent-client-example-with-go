package models

type ChatSession struct {
	Key               string `json:"key"`
	Id                string `json:"id"`
	AffinityToken     string `json:"affinityToken"`
	ClientPollTimeout int    `json:"clientPollTimeout"`
	ChasitorInit      string `json:"chasitorInit"`
}

type ReqMessage struct {
	Key           string `json:"key"`
	Message       string `json:"message"`
	AffinityToken string `json:"affinityToken"`
	Sequence      int    `json:"sequence"`
}
