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

// RandomNameAnimal returns a pseudo-random name and animal from the lists
func RandomNameAnimal() string {
	return fmt.Sprintf("%s %s", RandomName(), RandomAnimal())
}

// SafeRandomNameAnimal returns a pseudo-random name and animal from the lists goro-safely
func SafeRandomNameAnimal() string {
	return fmt.Sprintf("%s %s", SafeRandomName(), SafeRandomAnimal())
}

// RandomAdjectiveName returns a pseudo-random adjective and name from the lists
func RandomAdjectiveName() string {
	return fmt.Sprintf("%s %s", RandomAdjective(), RandomName())
}

// SafeRandomAdjectiveName returns a pseudo-random adjective and animal from the lists goro-safely
func SafeRandomAdjectiveName() string {
	return fmt.Sprintf("%s %s", SafeRandomAdjective(), SafeRandomName())
}
