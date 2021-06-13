package server

import (
	"log"

	"github.com/kwong21/graphql-go-cardkeeper/resolver"
	"github.com/kwong21/graphql-go-cardkeeper/service"
)

func Init(s service.DataService) {

	rootResolver, err := resolver.NewRoot(s)

	if err != nil {
		log.Fatal("Cannot instantiate the root resolver", err)
	}

	r := NewRouter(rootResolver)
	r.Run()
}
