package topinterview145

/**
判断最后进位大于1，说明要扩容，并且首位设置为1即可
*/

func plusOne(digits []int) []int {

	top := 1
	for i := len(digits) - 1; i >= 0; i-- {
		res := digits[i] + top
		digits[i] = res % 10
		top = res / 10
	}
	if top == 0 {
		return digits
	}

	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
