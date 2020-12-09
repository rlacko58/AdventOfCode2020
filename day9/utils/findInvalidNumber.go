package utils

func FindInvalidNumber(preAmbleLen int, numbers *[]int) int {
	for i := preAmbleLen; i < len(*numbers); i++ {
		if !FindSumInRange((*numbers)[i], i-preAmbleLen, i, numbers) {
			return (*numbers)[i]
		}
	}
	return -1
}