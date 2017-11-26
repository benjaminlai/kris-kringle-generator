package generator

import (
	"github.com/BenjaminLai/kris-kringle-generator/app/models"
)

type assigner struct {
	participants map[int]*models.Participant
	receivers    []*models.Receiver
}

func NewAssigner(participants map[int]*models.Participant, receivers []*models.Receiver) *assigner {
	return &assigner{
		participants: participants,
		receivers:    receivers,
	}
}

func (a *assigner) AssignReceiverToParticipant() []*models.Participant {
	var participantsWithReceivers []*models.Participant

	for _, participant := range a.participants {
		receiver := a.findSuitableReceiver(participant.GetBlacklistNames())
		participantsWithReceivers = append(participantsWithReceivers, &models.Participant{
			Name: participant.Name,
			Receiver: &models.Receiver{
				Name:  receiver.Name,
				Phone: receiver.Phone,
			},
		})
	}

	return participantsWithReceivers
}

func (a *assigner) findSuitableReceiver(blacklist []string) *models.Receiver {
	var receiver *models.Receiver

	receiverCounter := 0
	found := false

	for !found {
		if !contains(blacklist, a.receivers[receiverCounter].Name) {
			found = true

			receiver = &models.Receiver{
				Name:  a.receivers[receiverCounter].Name,
				Phone: a.receivers[receiverCounter].Phone,
			}
			a.receivers = append(a.receivers[:receiverCounter], a.receivers[receiverCounter+1:]...)
		}
		receiverCounter++
	}

	return receiver
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
