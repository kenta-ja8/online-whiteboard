/*
 * HelloApiSchema
 *
 * Practice api schema
 *
 * API version: 1.0.0
 * Contact: doriven@example.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/app"
	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	// UserApiService := openapi.NewUserApiService()
	UserApiService := app.NewUserApiService()
	UserApiController := openapi.NewUserApiController(UserApiService)

	router := openapi.NewRouter(UserApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
