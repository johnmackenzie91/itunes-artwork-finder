package itunes

import "fmt"

// errNon2XX is returned when the itunes api returns a response that is not in the 200 range
type errNon2XX int

func (e errNon2XX) Error() string {
	return fmt.Errorf("itunes returned a non 200 status: %d", e).Error()
}
