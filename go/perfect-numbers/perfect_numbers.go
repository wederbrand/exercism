package perfect

import "errors"

type Classification int

const (
	ClassificationDeficient = iota
	ClassificationPerfect   = iota
	ClassificationAbundant  = iota
)

var ErrOnlyPositive error = errors.New("only positive values are supported")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}
	factors := factor(n)
	sum := 0
	for _, i := range factors {
		sum += i
	}
	if sum == int(n) {
		return ClassificationPerfect, nil
	} else if sum < int(n) {
		return ClassificationDeficient, nil
	} else {
		return ClassificationAbundant, nil
	}
}

func factor(n int64) (result []int) {
	for i := 1; i < int(n); i++ {
		if int(n)%i == 0 {
			result = append(result, i)
		}
	}

	return
}
