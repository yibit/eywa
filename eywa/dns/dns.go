package dns

import (
	"time"

	"github.com/miekg/dns"
)

func Dns(server, domain string, timeout int) (string, []string, error) {
	cli := dns.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	message := dns.Msg{}
	message.SetQuestion(domain+".", dns.TypeA)
	r, _, err := cli.Exchange(&message, server)
	if err != nil {
		return "", nil, err
	}

	var addrs []string
	var cname string
	for _, it := range r.Answer {
		addr, ok := it.(*dns.A)
		if ok {
			addrs = append(addrs, addr.A.String())
			continue
		}

		cname_, ok := it.(*dns.CNAME)
		if ok {
			cname = cname_.Target
		}
	}

	return cname, addrs, nil
}

func Fqdn(server, fqdn string, timeout int) ([]string, error) {
	cli := dns.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	message := dns.Msg{}
	message.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	r, _, err := cli.Exchange(&message, server)
	if err != nil {
		return nil, err
	}

	var addrs []string
	for _, it := range r.Answer {
		addr, ok := it.(*dns.A)
		if ok {
			addrs = append(addrs, addr.A.String())
		}
	}

	return addrs, nil
}
