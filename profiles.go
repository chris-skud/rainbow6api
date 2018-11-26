package rainbow6api

import (
	"encoding/json"
	"strings"
	"time"
)

type Profile struct {
	ProfileID    string `json:"profileId"`
	Units        int    `json:"units"`
	UnitsEarned  int    `json:"unitsEarned"`
	UnitsSpent   int    `json:"unitsSpent"`
	Xp           int    `json:"xp"`
	IsClubMember bool   `json:"isClubMember"`
	CurrentLevel struct {
		LevelNumber int    `json:"levelNumber"`
		TierName    string `json:"tierName"`
		XpLowBound  int    `json:"xpLowBound"`
		XpHighBound int    `json:"xpHighBound"`
	} `json:"currentLevel"`
	CreationDate time.Time `json:"creationDate"`
}

const profileURL = "https://public-ubiservices.ubi.com/v1/profiles/club"

func (a *API) Profiles(profileIDs []string) ([]Profile, error) {
	if err := a.checkSession(); err != nil {
		return nil, err
	}

	profileIDsValue := strings.Join(profileIDs, ",")
	var url = profileURL + "?profileIds=" + profileIDsValue

	b, err := a.request(url)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	err = json.Unmarshal(b, &profiles)

	return profiles, err
}
