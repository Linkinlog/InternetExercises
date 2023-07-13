package chessboard

// File Declare a type named File which stores if a piece occupies a square - this will be a slice of booleans
type File []bool

// Chessboard Declares a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	counter := 0
	for _, el := range cb[file] {
		if el {
			counter = counter + 1
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > len(cb) {
		return 0 // rank must be within the size of the chessboard
	}
	counter := 0
	for _, el := range cb {
		if el[rank-1] {
			counter = counter + 1
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for _, el := range cb {
		counter = counter + len(el)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, el := range cb {
		for _, innerEl := range el {
			if innerEl {
				counter = counter + 1
			}
		}
	}
	return counter
}
