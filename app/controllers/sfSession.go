package controllers

import (
	"bytes"
	"fmt"
	//"github.com/revel/revel"
	"encoding/json"
	"io/ioutil"
	"liveagentchat/app/models"
	"net/http"
)

/*type SessionSF struct {
	*revel.Controller
}*/

func getSession(visitor string) models.ChatSession {
	var sessionObject models.ChatSession
	var resObject models.ChatSession

	url := BaseUrl + "chat/rest/System/SessionId"
	//fmt.Println("URL:>", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-LIVEAGENT-API-VERSION", "40")
	req.Header.Set("X-LIVEAGENT-AFFINITY", "null")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	json.Unmarshal(body, &sessionObject)
	fmt.Println("sessionObject", sessionObject)
	//item := models.ChatSession{visitor, "xxxxx", "xxxxx", "ccccccc"}
	chasitor := getChasitorInit(visitor, sessionObject)
	resObject = models.ChatSession{sessionObject.Key, sessionObject.Id, sessionObject.AffinityToken, sessionObject.ClientPollTimeout, chasitor}
	return resObject
}

func getChasitorInit(visitor string, oSession models.ChatSession) string {
	//var resObject models.ChatSession

	url := BaseUrl + "chat/rest/Chasitor/ChasitorInit"
	//fmt.Println("URL:>", url)
	var sxBody string = `{"organizationId": "` + OrganizationId + `","deploymentId": "` + DeploymentId + `", "buttonId": "` + ButtonId + `",`
	sxBody += `"sessionId":"` + oSession.Id + `",`
	sxBody += `"userAgent": "","language": "en-US","screenResolution": "1900x1080","visitorName": "` + visitor + `",`
	sxBody += `"prechatDetails": [],"prechatEntities": [],"receiveQueueUpdates": true,"isPost": true}`
	fmt.Println("sxBody:", sxBody)
	var jsonStr = []byte(sxBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-LIVEAGENT-API-VERSION", "40")
	req.Header.Set("X-LIVEAGENT-AFFINITY", oSession.AffinityToken)
	req.Header.Set("X-LIVEAGENT-SESSION-KEY", oSession.Key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return string(body)
}
