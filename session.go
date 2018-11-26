package rainbow6api

import (
	"fmt"
	"time"
)

type Session struct {
	Token                         string      `json:"token"`
	Ticket                        string      `json:"ticket"`
	TwoFactorAuthenticationTicket interface{} `json:"twoFactorAuthenticationTicket"`
	Expiration                    time.Time   `json:"expiration"`
	PlatformType                  string      `json:"platformType"`
	ProfileID                     string      `json:"profileId"`
	UserID                        string      `json:"userId"`
	Username                      string      `json:"username"`
	NameOnPlatform                string      `json:"nameOnPlatform"`
	InitializeUser                bool        `json:"initializeUser"`
	SpaceID                       string      `json:"spaceId"`
	Environment                   string      `json:"environment"`
	HasAcceptedLegalOptins        bool        `json:"hasAcceptedLegalOptins"`
	AccountIssues                 interface{} `json:"accountIssues"`
	SessionID                     string      `json:"sessionId"`
	ClientIP                      string      `json:"clientIp"`
	ClientIPCountry               string      `json:"clientIpCountry"`
	ServerTime                    time.Time   `json:"serverTime"`
	RememberMeTicket              string      `json:"rememberMeTicket"`
}

type ErrLogin struct {
	Message         string    `json:"message"`
	ErrorCode       int       `json:"errorCode"`
	HTTPCode        int       `json:"httpCode"`
	ErrorContext    string    `json:"errorContext"`
	MoreInfo        string    `json:"moreInfo"`
	TransactionTime time.Time `json:"transactionTime"`
	TransactionID   string    `json:"transactionId"`
	Environment     string    `json:"environment"`
}

func (s *ErrLogin) Error() string {
	return fmt.Sprintf("errorCode: %d, message: %s", s.ErrorCode, s.Message)
}
