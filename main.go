package main

import (
    "github.com/graphql-go/handler"
    "net/http"
    "graphql/schema"
)

func main() {
    h := handler.New(&handler.Config{
        Schema: &schema.Schema,
        Pretty: true,
        GraphiQL: true,
    })

    http.Handle("/graphql", h)
    http.ListenAndServe(":8080", nil)
}