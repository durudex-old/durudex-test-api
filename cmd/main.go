/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"log"
	"os"

	"github.com/durudex/durudex-test-api/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	handler := http.NewHandler()

	handler.InitRoutes(app)

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	log.Printf("Server is runned it '%s:%s'", host, port)

	if err := app.Listen(host + ":" + port); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
