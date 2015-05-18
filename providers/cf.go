package providers

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/bitly/go-simplejson"
	"github.com/leogsilva/google_auth_proxy/api"
)

type CloudfoundryProvider struct {
	*ProviderData
}

func NewCloudfoundryProvider(p *ProviderData) *CloudfoundryProvider {
        log.Printf("Provider found for Cloud Foundry")
	p.ProviderName = "cloudfoundry"
	if p.LoginUrl.String() == "" {
		p.LoginUrl = &url.URL{Scheme: "https",
			Host: "uaa.10.0.0.63.xip.io",
			Path: "/oauth/authorize"}
	}
	if p.RedeemUrl.String() == "" {
		p.RedeemUrl = &url.URL{Scheme: "https",
			Host: "uaa.10.0.0.63.xip.io",
			Path: "/oauth/token"}
	}
	if p.ProfileUrl.String() == "" {
		p.ProfileUrl = &url.URL{Scheme: "https",
			Host: "uaa.10.0.0.63.xip.io",
			Path: "/userinfo?schema=openid"}
	}
	if p.Scope == "" {
		p.Scope = "jenkins_demo.user openid"
	}
	return &CloudfoundryProvider{ProviderData: p}
}

func (p *CloudfoundryProvider) GetEmailAddress(unused_auth_response *simplejson.Json,
	access_token string) (string, error) {
	if access_token == "" {
		return "", errors.New("missing access token")
	}
	params := url.Values{}
	req, err := http.NewRequest("GET", p.ProfileUrl.String(), bytes.NewBufferString(params.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-li-format", "json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed %s making request %s", req, err)
		return "", err
	}
	mapResult, err := json.Map()
        log.Printf("User profile %s %s",mapResult, err)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
        email := mapResult["email"]
	return email, nil
}
