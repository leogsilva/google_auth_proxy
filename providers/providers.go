package providers

import (
	"github.com/bitly/go-simplejson"
)

type Provider interface {
	Data() *ProviderData
	GetEmailAddress(auth_response *simplejson.Json,
		access_token string) (string, error)
}

func New(provider string, p *ProviderData) Provider {
	switch provider {
	case "myusa":
		return NewMyUsaProvider(p)
	case "linkedin":
		return NewLinkedInProvider(p)
        case "cloudfoundry":
                return NewCloudfoundryProvider(p)
	default:
		return NewGoogleProvider(p)
	}
}
