package greetings

import "fmt"
import "errors"
import "math/rand"

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with an message
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// Return a random selected message format by specifying
	// a random inde for the slice of formats
	return formats[rand.Intn(len(formats))]
}
