package util

func Must(error error) {
	if error != nil {
		panic(error)
	}
}