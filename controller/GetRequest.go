package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/anggunpermata/integration-webhook/config"
	"github.com/anggunpermata/integration-webhook/models"
)

func AssignAgent(cus models.CustomerData, agentId string) (*models.QiscusResponse, error) {
	data := url.Values{}
	data.Set("room_id", cus.RoomId)
	data.Set("agent_id", agentId)
	data.Set("app_id", cus.AppID)
	data.Set("name", cus.Name)

	var w http.ResponseWriter
	var r *http.Request
	resp, err := InitiateChat(w, r, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	newResp := new(models.QiscusResponse)
	if err2 := json.Unmarshal(resp, &newResp.Data); err2 != nil {
		return nil, err
	} else {
		newResp.Message = "Assign Agent Success"
	}
	return newResp, nil
}

func InitiateChat(w http.ResponseWriter, r *http.Request, payload *strings.Reader) ([]byte, error) {

	uri := "https://multichannel.qiscus.com/api/v1/admin/service/assign_agent"
	method := "POST"

	fmt.Println(payload)
	//---------------------
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, payload)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error from structuring requests")
	}

	AdminToken := config.GoDotEnvVariable("AdminToken")
	AppCode := config.GoDotEnvVariable("AppId")

	req.Header.Add("Authorization", AdminToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Qiscus-App-Id", AppCode)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error from getting response")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error from read response body")
	}
	return body, nil
	//-------------------------------
}
