package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Customer struct {
	gorm.Model
	AgentID int    `json:"agent_id"`
	RoomID  string `json:"room_id"`
	Status  string `json:"status"`
}

type InitiateReq struct {
	AppID  string `json:"app_id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

type InitiateResp struct {
	Data struct {
		IdentityToken string `json:"identity_token"`
		RoomID        string `json:"room_id"`
		SdkUser       struct {
			ID          int    `json:"id"`
			Token       string `json:"token"`
			Email       string `json:"email"`
			Password    string `json:"password"`
			DisplayName string `json:"display_name"`
			AvatarURL   string `json:"avatar_url"`
			Extras      struct {
				AdditionalExtras map[string]string `json:"additional_extras"`
				IsCustomer       bool              `json:"is_customer"`
				Type             string            `json:"customer"`
			} `json:"extras"`
		} `json:"sdk_user"`
		IsSessional bool `json:"is_sessional"`
	} `json:"data"`
}

type CustomerDataFull struct {
	AppID         string `json:"app_id"`
	Source        string `json:"source"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	AvatarURL     string `json:"avatar_url"`
	Extras        string `json:"extras"`
	IsResolved    bool   `json:"is_resolved"`
	LatestService struct {
		ID                    int    `json:"id"`
		UserID                int    `json:"user_id"`
		RoomLogID             int    `json:"room_log_id"`
		AppID                 int    `json:"app_id"`
		RoomID                string `json:"room_id"`
		Notes                 string `json:"notes"`
		ResolvedAt            string `json:"resolved_at"`
		IsResolved            bool   `json:"is_resolved"`
		CreatedAt             string `json:"created_at"`
		UpdatedAt             string `json:"updated_at"`
		FirstCommentID        string `json:"first_comment_id"`
		LastCommentID         string `json:"last_comment_id"`
		RetrievedAt           string `json:"retrieved_at"`
		FirstCommentTimestamp string `json:"first_comment_timestamp"`
	} `json:"latest_service"`
	RoomID         string `json:"room_id"`
	CandidateAgent struct {
		ID                  int      `json:"id"`
		Name                string   `json:"name"`
		Email               string   `json:"email"`
		AuthenticationToken string   `json:"authentication_token"`
		CreatedAt           string   `json:"created_at"`
		UpdatedAt           string   `json:"updated_at"`
		SdkEmail            string   `json:"sdk_email"`
		SdkKey              string   `json:"sdk_key"`
		IsAvailable         bool     `json:"is_available"`
		Type                int      `json:"type"`
		AvatarURL           string   `json:"avatar_url"`
		AppID               int      `json:"app_id"`
		IsVerified          bool     `json:"is_verified"`
		NotificationsRoomID string   `json:"notifications_room_id"`
		BubbleColor         string   `json:"bubble_color"`
		QismoKey            string   `json:"qismo_key"`
		DirectLoginToken    string   `json:"direct_login_token"`
		TypeAsString        string   `json:"type_as_string"`
		AssignedRules       []string `json:"assigned_rules"`
	} `json:"candidate_agent"`
}

type CustomerData struct {
	AppID          string `json:"app_id"`
	CandidateAgent struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		IsAvailable bool   `json:"is_available"`
	} `json:"candidate_agent"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	RoomId       string `json:"room_id"`
	IsNewSession bool   `json:"is_new_session"`
	IsResolved   bool   `json:"is_resolved"`
}

/*
type QiscusInitiated struct {
	{
		"app_id": "oni-bgo2lummmhvzqxbt5",
		"source": "qiscus",
		"name": "sudah",
		"email": "sudah@gmail.com",
		"avatar_url": "https:\/\/somewhereonlyweknow.net\/prod\/image\/upload\/wMWsDZP6ta\/1516689726-ic_qiscus_client.png",
		"extras": "{\"timezone_offset\":7}",
		"is_resolved": true,
		"latest_service": {
			"id": 244,
			"user_id": 1,
			"room_log_id": 53,
			"app_id": 1,
			"room_id": "1905692",
			"notes": null,
			"resolved_at": "2019-02-04 04:49:47",
			"is_resolved": true,
			"created_at": "2019-02-04 04:49:47",
			"updated_at": "2019-02-04 04:49:47",
			"first_comment_id": "15167003",
			"last_comment_id": "15167015",
			"retrieved_at": "2019-02-04 04:49:47",
			"first_comment_timestamp": null
		},
		"room_id": "1905692",
		"candidate_agent": {
			"id": 22,
			"name": "dewi",
			"email": "dewi@mail.com",
			"authentication_token": "NlASwSIUnAqoTcFjYNBR",
			"created_at": "2019-01-17 06:50:20",
			"updated_at": "2019-01-18 10:12:59",
			"sdk_email": "vsC6x_dewi@mail.com",
			"sdk_key": "NZTGb",
			"is_available": true,
			"type": 2,
			"avatar_url": "https:\/\/somewhereonlyweknow.net\/prod\/image\/upload\/D1se5xo40I\/1516941944-Screen_Shot_2018-01-26_at_11.45.20.png",
			"app_id": 1,
			"is_verified": false,
			"notifications_room_id": "1692312",
			"bubble_color": "#666666",
			"qismo_key": "43Ondc",
			"direct_login_token": null,
			"type_as_string": "agent",
			"assigned_rules": [
				"fb_messaging",
				"line_messaging",
				"qiscus_messaging",
				"wa_messaging",
				"telegram_messaging"
			]
	}
}
*/

type QiscusResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Queue struct {
	AgentID int `json:"agent_id"`
	RoomID  int `json:"room_id"`
}
