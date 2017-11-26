package messenger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/BenjaminLai/kris-kringle-generator/app/models"
)

const (
	message = "Hello %s from Santa! Your KK is %s"
)

type twilioMessenger struct {
	baseURL    string
	numberFrom string
	accountSid string
	authToken  string
	client     *http.Client
}

func NewTwilioMessenger(baseURL, numberFrom, accountSid, authToken string) *twilioMessenger {
	client := &http.Client{}

	return &twilioMessenger{
		baseURL:    baseURL,
		numberFrom: numberFrom,
		accountSid: accountSid,
		authToken:  authToken,
		client:     client,
	}
}

func (messenger *twilioMessenger) Message(participants []*models.Participant) error {
	twilioURL := fmt.Sprintf("%s/%s/%s", messenger.baseURL, messenger.accountSid, "Messages.json")

	for _, participant := range participants {
		msgData := url.Values{}
		msgData.Set("To", participant.Receiver.Phone)
		msgData.Set("From", messenger.numberFrom)
		msgData.Set("Body", fmt.Sprintf(message, participant.Receiver.Name, participant.Name))
		msgDataReader := *strings.NewReader(msgData.Encode())

		req, err := http.NewRequest("POST", twilioURL, &msgDataReader)
		if err != nil {
			log.Fatal(err)
		}
		req.SetBasicAuth(messenger.accountSid, messenger.authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := messenger.client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var data map[string]interface{}
			decoder := json.NewDecoder(resp.Body)
			err := decoder.Decode(&data)
			if err == nil {
				fmt.Println(data["sid"])
			}
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			fmt.Println(resp.Status)

		}
	}

	return nil
}
