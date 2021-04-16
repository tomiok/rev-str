package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	s := parseFlags()
	if shouldRunCLI(s) {
		reversed := runCLI(s)
		fmt.Println(reversed)
	}
	runWebserver()
}

func shouldRunCLI(s string) bool {
	return s != ""
}

func runCLI(s string) string {
	return reverseStringArray(s)
}

func parseFlags() string {
	s := flag.String("s", "", "add a string value to reverse in CLI")
	flag.Parse()
	return *s
}

func runWebserver() {
	app := fiber.New()

	app.Post("/reverse", reverseHandler)
	must(app.Listen(":9595"))
}

func must(err error) {
	panic(err)
}
