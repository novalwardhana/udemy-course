package helper

func CategoryPanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
