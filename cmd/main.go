package main

func main() {
	app := newServer()

	app.Listen(":3000")
}
