package color

import "fmt"

func LogColorErr(s string) error {
	return fmt.Errorf(s)
}
