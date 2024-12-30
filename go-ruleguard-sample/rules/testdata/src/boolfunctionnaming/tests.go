package boolfuncnametest

func IsPositive(x int) bool {
	return x > 0
}

func Zero(x int) bool { // want "bool function name should start with 'Is'"
	return x == 0
}
