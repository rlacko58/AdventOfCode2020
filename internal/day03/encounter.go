package day03

func CountEncounteredTrees(rows *[]string, right, down int) int {
	treeNum := 0

	column := 0
	for row := down; row < len(*rows); row+=down {
		column += right
		column %= len((*rows)[row])
		if(string((*rows)[row][column]) == "#") {
			treeNum ++
		}
	}

	return treeNum
}