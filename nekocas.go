package NekoCAS

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type cas struct {
	domain string
	secret string
}

// New: returns a cas struct.
func New(domain string, secret string) *cas {
	return &cas{
		domain: domain,
		secret: secret,
	}
}

// SetDomain: set the cas domain.
func (c *cas) SetDomain(domain string) {
	c.domain = domain
}

// SetSecret: set the cas secret.
func (c *cas) SetSecret(secret string) {
	c.secret = secret
}

// Validate: check the ticket.
func (c *cas) Validate(ticket string) (*user, error) {
	params := map[string]interface{}{
		"service": c.secret,
		"ticket":  ticket,
	}
	targetURL := c.domain + "/validate"
	resp, body, errs := gorequest.New().Get(targetURL).Query(params).End()
	if len(errs) > 0 {
		return nil, errors.Wrapf(errs[0], "request %q", targetURL)
	} else if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("unexpected status code %d for %q", resp.StatusCode, targetURL)
	}

	var response casResponse
	err := json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, errors.Errorf("failed to unmarshal cas response: %v", err)
	}

	return newUser(response), nil
}

func newUser(resp casResponse) *user {
	return &user{
		Name:    resp.Data.Name,
		Email:   resp.Data.Email,
		Token:   resp.Data.Token,
		Message: resp.Message,
	}
}
