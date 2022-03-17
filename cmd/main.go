/*
 * Copyright © 2022 Durudex

 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"log"
	"net/http"

	"github.com/durudex/durudex-test-api/internal/delivery/graphql"
)

func main() {
	handler := graphql.NewHandler()

	http.Handle("/", handler.PlaygroundHandler())
	http.Handle("/query", handler.GraphqlHandler())

	log.Println("Server is runned it 'localhost:8000'")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
