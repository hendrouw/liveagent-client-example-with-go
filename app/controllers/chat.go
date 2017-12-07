package controllers

import (
	"fmt"
	"github.com/revel/revel"
	//"liveagentchat/app/controllers"
	//"encoding/json"
	"liveagentchat/app/models"
)

type Chat struct {
	*revel.Controller
}

func (c Chat) GetSession() revel.Result {
	visitor := c.Params.Query.Get("visitor")
	var res models.ChatSession

	//item := models.ChatSession{act, "xxxxx", "xxxxx", "ccccccc"}
	//res = item
	res = getSession(visitor)
	//fmt.Println(res)
	return c.RenderJSON(res)
}

func (c Chat) GetMessages() revel.Result {
	key := c.Params.Query.Get("key")
	at := c.Params.Query.Get("at")
	var res string

	res = getSfMessages(key, at)
	return c.RenderJSON(res)
}

func (c Chat) SendMessage() revel.Result {
	var res string
	req := models.ReqMessage{}
	c.Params.BindJSON(&req)
	fmt.Println("req:", req)

	res = sendSfMessage(req.Key, req.AffinityToken, req.Message, req.Sequence)

	return c.RenderJSON(res)
}
