package business

import (
	"log"

	"github.com/teris-io/shortid"
)

func generateID() string {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Println("Couldn't create new shortid instance")
	}

	id, err := sid.Generate()
	if err != nil {
		log.Println("Couldn't generate UUId")
	}

	return id
}
