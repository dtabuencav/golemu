package main

// Check if error
func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}
