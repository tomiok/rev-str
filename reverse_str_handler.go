package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

func reverseHandler(ctx *fiber.Ctx) error {
	s := ctx.Query("s")
	reversed := reverseStringArray(s)
	b, err := json.Marshal(res{Reverse: reversed})

	if err != nil {
		log.Println("cannot marshall structure" + err.Error())
		return err
	}
	ctx.Set("Content-type", "application/json")
	return ctx.Send(b)
}

type res struct {
	Reverse string `json:"reverse_string"`
}
