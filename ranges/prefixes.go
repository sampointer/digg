package ranges

import (
	"encoding/json"
	"fmt"
)

// Prefix holds the detail of a given Google prefix
type Prefix struct {
	IPV4Prefix string `json:"ipv4Prefix,omitempty"`
	IPV6Prefix string `json:"ipv6Prefix,omitempty"`
	Scope      string `json:"scope"`
	Service    string `json:"service"`
}

//String returns a column-format representation of the Prefix
func (p Prefix) String() string {
	var prefix string
	if p.IPV4Prefix != "" {
		prefix = concatPrefix(prefix, p.IPV4Prefix)
	}

	if p.IPV6Prefix != "" {
		prefix = concatPrefix(prefix, p.IPV6Prefix)
	}

	return fmt.Sprintf(
		"prefix: %s scope: %s service: %s",
		prefix,
		p.Scope,
		p.Service,
	)
}

//JSON returns a string of JSON representing the Prefix
func (p Prefix) JSON() (string, error) {
	out, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func concatPrefix(original string, add string) string {
	if original != "" {
		original = fmt.Sprintf("%s,", original)
	}

	return fmt.Sprintf("%s%s", original, add)
}
