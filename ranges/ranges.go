package ranges

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"time"
)

// Ranges represents an https://www.gstatic.com/ipranges/cloud.json or
// https://www.gstatic.com/ipranges/goog.json document
type Ranges struct {
	CreationTime    time.Time `json:"-"`
	CreationTimeRaw string    `json:"creationTime"`
	Prefixes        []Prefix  `json:"prefixes"`
	SyncToken       time.Time `json:"-"`
	SyncTokenRaw    string    `json:"syncToken"`
}

// LookupIPv4 returns the Prefix structs that contain a range that includes the
// passed IPv4 address
func (r *Ranges) LookupIPv4(ip net.IP) ([]Prefix, error) {
	var results []Prefix

	for _, p := range r.Prefixes {
		if p.IPV4Prefix != "" {
			_, pIPNet, err := net.ParseCIDR(p.IPV4Prefix)
			if err != nil {
				return nil, err
			}

			if pIPNet.Contains(ip) {
				results = append(results, p)
			}
		}
	}

	return results, nil
}

// LookupIPv6 returns the Prefix structs that contain a range that includes the
// passed IPv6 address
func (r *Ranges) LookupIPv6(ip net.IP) ([]Prefix, error) {
	var results []Prefix

	for _, p := range r.Prefixes {
		if p.IPV6Prefix != "" {
			_, pIPNet, err := net.ParseCIDR(p.IPV6Prefix)
			if err != nil {
				return nil, err
			}

			if pIPNet.Contains(ip) {
				results = append(results, p)
			}
		}
	}

	return results, nil
}

//New is a constructor for Ranges
func New(r io.Reader) (*Ranges, error) {
	var ranges Ranges

	doc, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(doc, &ranges)
	if err != nil {
		return nil, err
	}

	d, err := parseCreationTime(&ranges.CreationTimeRaw)
	if err != nil {
		return nil, err
	}
	ranges.CreationTime = d

	s, err := strconv.ParseInt(ranges.SyncTokenRaw[0:10], 10, 64)
	if err != nil {
		return nil, err
	}

	m, err := strconv.ParseInt(ranges.SyncTokenRaw[10:13], 10, 64)
	if err != nil {
		return nil, err
	}

	ranges.SyncToken = time.Unix(s, m*1000000).UTC()

	return &ranges, nil
}

func parseCreationTime(s *string) (time.Time, error) {
	const creationTimeFormat = "2006-01-02T15:04:05.000"
	t, err := time.Parse(creationTimeFormat, *s)
	return t, err
}
