package generator

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/BenjaminLai/kris-kringle-generator/app/models"
)

func ReadFromJSON(fileLocation string) (map[int]*models.Participant, error) {
	raw, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	var payload []*models.Participant
	participants := make(map[int]*models.Participant, 0)

	if err := json.Unmarshal(raw, &payload); err != nil {
		log.Fatal(err)
	}

	for i, entry := range payload {
		var blacklist []*models.Receiver
		for _, blacklister := range entry.Blacklist {
			blacklist = append(blacklist, blacklister)
		}

		participant := &models.Participant{
			Name:      entry.Name,
			Phone:     entry.Phone,
			Blacklist: blacklist,
		}

		participants[i] = participant
	}

	return participants, nil
}
