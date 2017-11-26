package utils

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/BenjaminLai/kris-kringle-generator/app/models"
)

func TestCanRandomizeSlice(t *testing.T) {

	originalList := setupMockRetrieverList(5)
	fmt.Println("Original: \n", originalList)
	for _, receiver := range originalList {
		fmt.Println(receiver.Name)
	}

	Shuffle(originalList)

	fmt.Println("Shuffled: \n", originalList)
	for _, receiver := range originalList {
		fmt.Println(receiver.Name)
	}
}

func setupMockRetrieverList(amount int) []*models.Receiver {
	receivers := []*models.Receiver{}

	for i := 0; i < amount; i++ {
		receivers = append(receivers, &models.Receiver{
			Name: strconv.Itoa(i),
		})
	}

	return receivers
}
