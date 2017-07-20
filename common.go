package randomnames

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomAdjectiveAnimal() string {
	return fmt.Sprintf("%s %s", RandomAdjective(), RandomAnimal())
}

func SafeRandomAdjectiveAnimal() string {
	return fmt.Sprintf("%s %s", SafeRandomAdjective(), SafeRandomAnimal())
}
