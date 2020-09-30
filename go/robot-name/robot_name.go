// Package robotname names robots uniquely
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var allNames = make(map[string]bool)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Robot represents a robot with a name and very little else.
type Robot struct {
	name string
}

func randomName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

// Name returns the name of the robot.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(allNames) == 26*26*1000 {
		return "", errors.New("all names used")
	}

	r.name = randomName()
	for allNames[r.name] {
		r.name = randomName()
	}
	allNames[r.name] = true

	return r.name, nil
}

// Reset resets the name of the robot and gives it a new one.
func (r *Robot) Reset() {
	r.name = ""
}
