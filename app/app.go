package app

import (
	"fmt"
	"os"

	"github.com/BenjaminLai/kris-kringle-generator/app/generator"
	"github.com/BenjaminLai/kris-kringle-generator/app/messenger"
	"github.com/BenjaminLai/kris-kringle-generator/app/models"
	"github.com/BenjaminLai/kris-kringle-generator/app/utils"
)

type Configuration struct {
	TwilioBaseURL    string
	TwilioAccountSid string
	TwilioAuthToken  string
	TwilioNumberFrom string
}

func Generate() {

	cfg := loadConfigs()

	fileLocation := "resources/participants.json"

	participants, err := generator.ReadFromJSON(fileLocation)
	if err != nil {
		fmt.Println(err)
	}

	receivers := buildReceiversNameList(participants)

	utils.Shuffle(receivers)

	assigner := generator.NewAssigner(participants, receivers)
	fullList := assigner.AssignReceiverToParticipant()

	messenger := messenger.NewTwilioMessenger(cfg.TwilioBaseURL,
		cfg.TwilioNumberFrom,
		cfg.TwilioAccountSid,
		cfg.TwilioAuthToken,
	)

	if err := messenger.Message(fullList); err != nil {
		fmt.Println(err)
	}

	for _, participant := range fullList {
		fmt.Println(participant.Name)
		fmt.Println(participant.Receiver.Name)
		fmt.Println("-------------------------")
	}
}

func loadConfigs() *Configuration {
	return &Configuration{
		TwilioBaseURL:    os.Getenv("TWILIO_SMS_BASEURL"),
		TwilioAccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioNumberFrom: os.Getenv("TWILIO_NUMBER_FROM"),
	}
}

func buildReceiversNameList(participants map[int]*models.Participant) []*models.Receiver {
	var receivers []*models.Receiver

	for _, participant := range participants {
		receivers = append(receivers, &models.Receiver{
			Name:  participant.Name,
			Phone: participant.Phone,
		})
	}

	return receivers
}
