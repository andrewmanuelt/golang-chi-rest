package helper

func HandleError(errors error) {
	if errors != nil {
		panic(errors)
	}
}
