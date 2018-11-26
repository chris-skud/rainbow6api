package rainbow6api_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/chris-skud/rainbow6api"
)

func TestAPI(t *testing.T) {
	ubiAPI, err := rainbow6api.New(
		"314d4fef-e568-454a-ae06-43e3bece12a6",
		"CHANGE TO YOUR USER_LOGIN/EMAIL",
		"CHANGE TO YOUR USER_PASSWORD",
		"uplay",
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	profiles, err := ubiAPI.Profiles([]string{ubiAPI.Session.ProfileID})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, profile := range profiles {
		fmt.Printf("\n%+v\n", profile)
	}
}
