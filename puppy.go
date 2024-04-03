package puppy

import (
	"fmt"

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
func Version11() {
	fmt.Println("Hi I am from version v1.1.0")
}
