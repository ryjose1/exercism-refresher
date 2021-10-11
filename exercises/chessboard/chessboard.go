package chessboard

//  Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	for _, square := range cb[rank] {
		if square {
			ret++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	for _, rank := range cb {
		if file > 0 && file <= len(rank) && rank[file-1] {
			ret++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	//return len(cb) * len(cb[1])
	for _, rank := range cb {
		ret += len(rank)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	for k, _ := range cb {
		ret += CountInRank(cb, k)
	}
	return
}
