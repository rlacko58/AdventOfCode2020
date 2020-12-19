package convert

func StrArrToInt(strArr *[]string) []int {
	intArr := make([]int, 0)
	for _, v := range(*strArr) {
		intArr = append(intArr, ToInt(v))
	}
	return intArr
}