package puppy

import (
	"github.com/hardik-bhasin07/dog"
)

func Bark() string {
	return "whoof!"
}

func Barks() string {
	return "whoof whoof whoof"
}
func GrownBarks() string {
	return dog.WhenGrownUp(Barks())
}
func GrownBark() string {
	return dog.WhenGrownUp(Bark())
}
