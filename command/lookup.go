package command

import (
	"github.com/sampointer/digg/ranges"

	"net"
	"net/http"
	"strings"
)

//Lookup returns Prefixes the ranges of which include the passed IP address
func Lookup(q string) ([]ranges.Prefix, error) {
	var r []ranges.Prefix

	urls := []string{
		"https://www.gstatic.com/ipranges/goog.json",
		"https://www.gstatic.com/ipranges/cloud.json",
	}

	for _, url := range urls {
		l, err := lookup(q, url)
		if err != nil {
			return r, err
		}

		r = append(r, l...)

	}

	return r, nil
}

func lookup(q string, url string) ([]ranges.Prefix, error) {
	var client http.Client
	var prefixes []ranges.Prefix

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	r, err := ranges.New(resp.Body)
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(q)

	if isIPv4(ip) {
		p, err := r.LookupIPv4(ip)
		if err != nil {
			return nil, err
		}
		prefixes = append(prefixes, p...)
	} else {
		p, err := r.LookupIPv6(ip)
		if err != nil {
			return nil, err
		}
		prefixes = append(prefixes, p...,
		)
	}

	return prefixes, nil
}

func isIPv4(ip net.IP) bool {
	return strings.Contains(ip.String(), ".")
}
