package randomnames

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomAdjectiveAnimal returns a pseudo-random adjective and animal from the lists
func RandomAdjectiveAnimal() string {
	return fmt.Sprintf("%s %s", RandomAdjective(), RandomAnimal())
}

// SafeRandomAdjectiveAnimal returns a pseudo-random adjective and animal from the lists goro-safely
func SafeRandomAdjectiveAnimal() string {
	return fmt.Sprintf("%s %s", SafeRandomAdjective(), SafeRandomAnimal())
}
