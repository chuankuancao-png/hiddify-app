package config

import "github.com/sagernet/sing-box/option"

func (r *Rule) MakeRule() option.DefaultRule {
	rule := option.DefaultRule{}

	rule.DomainSuffix = append(rule.DomainSuffix, r.DomainSuffixes...)
	rule.Domain = append(rule.Domain, r.Domains...)
	rule.DomainKeyword = append(rule.DomainKeyword, r.DomainKeywords...)
	rule.DomainRegex = append(rule.DomainRegex, r.DomainRegexes...)

	for _, ip := range r.IpCidrs {
		if len(ip) > 7 && ip[:7] == "geoip:" {
			rule.GeoIP = append(rule.GeoIP, ip[7:])
		} else {
			rule.IPCIDR = append(rule.IPCIDR, ip)
		}
	}

	rule.PortRange = append(rule.PortRange, r.PortRanges...)

	switch r.Network {
	case Network_TCP:
		rule.Network = []string{"tcp"}
	case Network_UDP:
		rule.Network = []string{"udp"}
	}

	for _, p := range r.Protocols {
		switch p {
		case Protocol_TLS:
			rule.Protocol = append(rule.Protocol, "tls")
		case Protocol_HTTP:
			rule.Protocol = append(rule.Protocol, "http")
		case Protocol_QUIC:
			rule.Protocol = append(rule.Protocol, "quic")
		case Protocol_STUN:
			rule.Protocol = append(rule.Protocol, "stun")
		case Protocol_DNS:
			rule.Protocol = append(rule.Protocol, "dns")
		case Protocol_Bittorrent:
			rule.Protocol = append(rule.Protocol, "bittorrent")
		}
	}

	rule.ProcessName = append(rule.ProcessName, r.ProcessNames...)
	rule.ProcessPath = append(rule.ProcessPath, r.ProcessPaths...)

	return rule
}
