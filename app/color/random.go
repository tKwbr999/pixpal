package color

import (
	"math/rand"
	"time"
)

func Random(amount int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(amount)
}
