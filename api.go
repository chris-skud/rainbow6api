package rainbow6api

const (
	ubiLoginURL = ""
	ubiAPIURL   = ""
)

type API struct {
	appID    string
	email    string
	password string
	platform string
	Session  *Session
}

func New(appID, email, password, platform string) (*API, error) {
	if appID == "" {
		appID = "39baebad-39e5-4552-8c25-2c9b919064e2"
	}

	var api = &API{appID: appID, email: email, password: password, platform: platform}
	err := api.login()
	return api, err
}
