package controllers

import (
	"bytes"
	"fmt"
	//"github.com/revel/revel"
	//"encoding/json"
	"io/ioutil"
	//"liveagentchat/app/models"
	"net/http"
)

/*type SessionSF struct {
	*revel.Controller
}*/

func getSfMessages(sessionKey string, AffinityToken string) string {
	var sxResult string

	url := BaseUrl + "chat/rest/System/Messages"
	//fmt.Println("URL:>", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-LIVEAGENT-API-VERSION", "40")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LIVEAGENT-AFFINITY", AffinityToken)
	req.Header.Set("X-LIVEAGENT-SESSION-KEY", sessionKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	sxResult = string(body)
	return sxResult
}

func sendSfMessage(sessionKey string, AffinityToken string, msg string, sequence int) string {
	fmt.Println("sessionKey:", sessionKey)
	var sxResult string

	url := BaseUrl + "chat/rest/Chasitor/ChatMessage"
	//fmt.Println("URL:>", url)
	var sxBody string = `{"text":"` + msg + `"}`

	fmt.Println("sxBody:", sxBody)
	var jsonStr = []byte(sxBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-LIVEAGENT-API-VERSION", "40")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LIVEAGENT-AFFINITY", AffinityToken)
	req.Header.Set("X-LIVEAGENT-SESSION-KEY", sessionKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	sxResult = string(body)
	return sxResult
}
