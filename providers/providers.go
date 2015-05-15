package providers

import (
        "log"
	"github.com/bitly/go-simplejson"
)

type Provider interface {
	Data() *ProviderData
	GetEmailAddress(auth_response *simplejson.Json,
		access_token string) (string, error)
}

func New(provider string, p *ProviderData) Provider {
        log.Printf("Selecting provider %s",provider)
	switch provider {
	case "myusa":
		return NewMyUsaProvider(p)
	case "linkedin":
		return NewLinkedInProvider(p)
        case "cloudfoundry":
                return NewCloudfoundryProvider(p)
	default:
                log.Printf("Found %s",provider)
		return NewGoogleProvider(p)
	}
}
