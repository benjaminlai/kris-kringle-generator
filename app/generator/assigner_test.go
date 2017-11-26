package generator

import (
	"fmt"
	"testing"

	"github.com/BenjaminLai/kris-kringle-generator/app/models"
)

func TestCanAssignReceiverToParticipant(t *testing.T) {
	// asserter := assert.New(t)

	mockParticipantsList := map[int]*models.Participant{
		0: &models.Participant{
			Name:      "Ben",
			Blacklist: []string{"Joyce", "Gavin", "Ben"},
		},
		1: &models.Participant{
			Name:      "Joyce",
			Blacklist: []string{"Ben", "Gavin", "Joyce"},
		},
		2: &models.Participant{
			Name:      "Tegan",
			Blacklist: []string{"Derrick", "Nick", "Tegan"},
		},
		3: &models.Participant{
			Name:      "Steven",
			Blacklist: []string{"Debbie", "Steven"},
		},
	}

	mockReceivers := []*models.Receiver{
		&models.Receiver{
			Name: "Joyce",
		},
		&models.Receiver{
			Name: "Ben",
		},
		&models.Receiver{
			Name: "Tegan",
		},
		&models.Receiver{
			Name: "Steven",
		},
	}

	// expectedReceivers := []string{mockReceivers[1].Name, mockReceivers[0].Name}
	assigner := NewAssigner(mockParticipantsList, mockReceivers)
	participants := assigner.AssignReceiverToParticipant()

	// asserter.Equal(0, len(mockReceivers))
	for _, participant := range participants {

		fmt.Println("name", participant.Name)
		fmt.Println("recevier", participant.Receiver.Name)
		// asserter.Equal(participant.Receiver.Name, expectedReceivers[i])
	}
}
