package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func UserPanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
