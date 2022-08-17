package helper

/* Tidak bisa diimport */
func sayHello(guest, name string) string {
	return "Hello" + guest + ", I'm" + name
}

/* Bisa diimport */
func SayHello(guest, name string) string {
	return "Hello " + guest + ", I'm " + name
}
