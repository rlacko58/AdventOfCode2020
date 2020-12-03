package utils

func RightThreeDownOne(rows []string) int {
	treeNum := 0

	column := 0
	for row := 1; row < len(rows); row++ {
		column += 3
		column %= len(rows[row])
		if(string(rows[row][column]) == "#") {
			treeNum ++
		}
	}

	return treeNum
}