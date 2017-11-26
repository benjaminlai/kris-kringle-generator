package utils

import (
	"math/rand"

	"github.com/BenjaminLai/kris-kringle-generator/app/models"
)

func Shuffle(src []*models.Receiver) {
	for i := 1; i < len(src); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			src[r], src[i] = src[i], src[r]
		}
	}
}
