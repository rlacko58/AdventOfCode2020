package utils

func getColumn(arr *[][2]int, index int) []int {
	col := make([]int, 0)
	for _, v := range(*arr) {
		col = append(col, v[index])
	}
	return col
}

func calcM(remainders *[]int) int {
	remainder := 1
	for _, v := range(*remainders) {
		if v == 0 {
			continue
		}
		remainder *= v
	}
	return remainder
}

func calcMi(mods *[]int, index int) int {
	Mi := 1
	for i, v := range(*mods) {
		if i == index || v == 0 {
			continue
		}
		Mi *= v
	} 
	return Mi
}

func findY(x, mod int) int {
	simpX := x % mod
	for y:= 0; true; y++ {
		if (simpX*y)%mod == 1 {
			return y
		} 
	}
	return -1
}

func convertRemainders(diffs, mods *[]int) []int {
	remainders := make([]int, 0)
	for i, v := range(*diffs) {
		remainders = append(remainders, (*mods)[i]-v)
	}
	return remainders
}

func FindNumber(nums *[][2]int) int {
	modArr := getColumn(nums, 0)
	diffs := getColumn(nums, 1)
	remainders := convertRemainders(&diffs, &modArr)

	M := calcM(&modArr)

	MiArr := make([]int, 0)
	for i := range(remainders) {
		MiArr = append(MiArr, calcMi(&modArr, i))
	}

	YiArr := make([]int, 0)
	for i, Mi := range(MiArr) {
		YiArr = append(YiArr, findY(Mi, modArr[i]))
	} 
	
	x := 0
	for i := range(YiArr) {
		x += remainders[i] * MiArr[i] * YiArr[i]
	}
	x %= M

	return x 
}