/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/durudex/durudex-test-api/internal/delivery/graphql"
)

func main() {
	handler := graphql.NewHandler()

	http.Handle("/", handler.PlaygroundHandler())
	http.Handle("/query", handler.GraphqlHandler())

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	log.Printf("Server is runned it '%s:%s'", host, port)

	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
