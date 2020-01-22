package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/card"
)

func main() {
	xendit.Opt.SecretKey = "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab"

	createAuthorizationData := card.CreateAuthorizationParams{
		TokenID:          "5e280aeecb812150cac94743",
		AuthenticationID: "5e280aeecb812150cac94744",
		ExternalID:       "cardAuth-" + time.Now().String(),
		Amount:           200000,
	}

	createAuthorizationResp, err := card.CreateAuthorization(&createAuthorizationData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created authorization: %+v\n", createAuthorizationResp)

	reverseAuthorizationData := card.ReverseAuthorizationParams{
		ChargeID:   createAuthorizationResp.ID,
		ExternalID: "reverseAuth-" + time.Now().String(),
	}

	reverseAuthorizationResp, err := card.ReverseAuthorization(&reverseAuthorizationData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reversed authorization: %+v\n", reverseAuthorizationResp)
}
