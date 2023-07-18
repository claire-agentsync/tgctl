package scaffold

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

const colorRed string = "\033[91m"
const colorBlue string = "\033[34m"
const colorGreen string = "\033[92m"
const colorCyan string = "\033[36m"
const colorReset string = "\033[0m"

type Scaff struct {
	Account  string
	Region   string
	Service  string
	Stack    string
	Kind     string
	Resource string
	Target   string
	DryRun   bool
}

// Builds the scaffold directory
func (s *Scaff) Build() (string, error) {
	a := []string{s.Account, s.Region, s.Service, s.Stack, s.Kind, s.Resource}
	d := _Form(a)
	vd, err := _Valid(d)
	if err != nil {
		return d, err
	}
	if s.DryRun == true {
		return vd, nil
	}
	err = _Make(vd)
	if err != nil {
		return vd, err
	}
	return vd, nil
}

// Join directory paths and trim
func _Form(a []string) string {
	dir := strings.Join(a, "/")
	dir = strings.Trim(dir, "/")
	return dir
}

// Check if directory is valid
func _Valid(d string) (string, error) {
	c := strings.Contains(d, "//")
	if c == true {
		err := fmt.Sprintf("%v: is not a valid path", d)
		return d, errors.New(err)
	}
	return d, nil
}

// Make an absolute directory
/*
func _Full(d string) (string, error) {
	absdir, err := filepath.Abs(d)
        fmt.Println("From _Full:", absdir)
	if err != nil {
		return absdir, err
	}
	return absdir, err
}
*/
func _Make(d string) error {
	err := os.MkdirAll(d, 0750)
	return err
}

// // TESTS
func TestCorrectValid(t *testing.T) {
	a := [][]string{
		{"example-account"},
		{"example-account", "example-region"},
		{"example-account", "example-region", "example-service"},
		{"example-account", "example-region", "example-service", "example-stack"},
		{"example-account", "example-region", "example-service", "example-stack", "example-kind"},
	}
	for _, i := range a {
		d := _Form(i)
		dir, err := _Valid(d)
		if err != nil {
                    t.Errorf("%v invalid, want: valid", dir)
		}
	}
}

/*
func TestIncorrectForm(t *testing) {

}
*/
