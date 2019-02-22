/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tea-go/tea-go-web-boilerplate/pkg/http/rest"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/listing"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/storage/mysql"
)

func main() {
	/*
		// error handling omitted for simplicity
		s, _ := json.NewStorage()
		lister := listing.NewService(s)
		// set up the HTTP server
		router := rest.Handler(lister)

		fmt.Println("The beer server is on tap now: http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	*/
	// error handling omitted for simplicity
	s, _ := mysql.NewStorage()
	// automatically migrate schemas
	s.AutoMigrate()
	// create a listeing service
	lister := listing.NewService(s)
	// set up the HTTP server
	router := rest.Handler(lister)
	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
