// Package robotname names robots uniquely
package robotname

import (
	"errors"
	"fmt"
)

var allNames = make([]int, 0)
var nameIndex = 0

// Generate all possible names once then just randomly put them in a list of names.
// Keep track of which names has been used by moving an index forward when names are used.
// Returns new names at constant price as opposed to generate a new random name until one
// that hasn't been seen before is found, which would require generating on average
// 338.000 random names until the last one is found.
func init() {
	tempMap := make(map[int]struct{})
	for i := 0; i < 26*26*1000; i++ {
		tempMap[i] = struct{}{}
	}

	for k := range tempMap {
		allNames = append(allNames, k)
	}
}

// Robot represents a robot with a name and very little else.
type Robot struct {
	name string
}

func (r *Robot) generateName() error {
	if nameIndex == 26*26*1000 {
		return errors.New("all names used")
	}

	r.name = convertToName(allNames[nameIndex])
	nameIndex++
	// fmt.Printf("%d\n", nameIndex)
	return nil
}

func convertToName(i int) string {
	a := i % 26
	i /= 26
	b := i % 26
	i /= 26

	return fmt.Sprintf("%c%c%03d", a+'A', b+'A', i)
}

// Name returns the name of the robot.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		err := r.generateName()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

// Reset resets the name of the robot and gives it a new one.
func (r *Robot) Reset() {
	r.name = ""
}
