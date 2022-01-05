package chessboard

type Rank []bool

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	r, found := cb[rank]
	if !found {
		return 0
	}

	sum := 0
	for _, b := range r {
		if b {
			sum++
		}
	}

	return sum
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}
	sum := 0
	for _, rank := range cb {
		if rank[file-1] {
			sum++
		}
	}

	return sum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	sum := 0

	for _, r := range cb {
		sum += len(r)
	}

	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	sum := 0
	for _, r := range cb {
		for _, b := range r {
			if b {
				sum++
			}
		}
	}

	return sum
}
