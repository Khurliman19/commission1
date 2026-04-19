package commission

func Calculate(amount int64, isAlif bool) int64 {
	if isAlif {
		return 0
	}
	return amount * 29 / 10000 // 0.29%
}

func Validate(amount int64) bool {
	return amount >= 500 && amount <= 15_000_000
}
