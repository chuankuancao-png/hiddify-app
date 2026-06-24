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
	case Network_tcp:
		rule.Network = []string{"tcp"}
	case Network_udp:
		rule.Network = []string{"udp"}
	}

	for _, p := range r.Protocols {
		switch p {
		case Protocol_tls:
			rule.Protocol = append(rule.Protocol, "tls")
		case Protocol_http:
			rule.Protocol = append(rule.Protocol, "http")
		case Protocol_quic:
			rule.Protocol = append(rule.Protocol, "quic")
		case Protocol_stun:
			rule.Protocol = append(rule.Protocol, "stun")
		case Protocol_dns:
			rule.Protocol = append(rule.Protocol, "dns")
		case Protocol_bittorrent:
			rule.Protocol = append(rule.Protocol, "bittorrent")
		}
	}

	rule.ProcessName = append(rule.ProcessName, r.ProcessNames...)
	rule.ProcessPath = append(rule.ProcessPath, r.ProcessPaths...)

	return rule
}
