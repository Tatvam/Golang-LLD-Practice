package utils

func GetRowAndColumn(size, number int) (int, int) {
	row := (number - 1) / size
	col := (number - 1) % size
	if row%2 == 1 {
		col = size - col - 1
	}
	return row, col

}
