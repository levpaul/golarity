package magnet

import (
	"fmt"
	"strings"
)

const (
	paramDisplayName    = "dn"
	paramAddressTracker = "tr"
	paramExactTopic     = "xt"
)

type Magnet struct {
	DisplayName    string   // dn
	ExactTopic     string   // xt
	AddressTracker []string // tr
}

func ParseMagnet(magnetURL string) (*Magnet, error) {
	if len(magnetURL) >= 8 && magnetURL[0:8] == "magnet:?" {
		magnetURL = magnetURL[8:]
	}

	magnet := new(Magnet)

	params := strings.Split(magnetURL, "&")
	for _, p := range params {
		splitParam := strings.Split(p, "=")
		if len(splitParam) != 2 {
			return nil, fmt.Errorf("bad param found, expected a '='... param: '%p'")
		}
		switch splitParam[0] {
		case paramDisplayName:
			if magnet.DisplayName != "" {
				return nil, fmt.Errorf("multiple 'Display Name' params found")
			}
			magnet.DisplayName = splitParam[1]
		case paramExactTopic:
			if magnet.ExactTopic != "" {
				return nil, fmt.Errorf("multiple 'Exact Topic' params found")
			}
			magnet.ExactTopic = splitParam[1]
		case paramAddressTracker:
			// This param supports multiple values
			magnet.AddressTracker = append(magnet.AddressTracker, splitParam[1])
		}
	}

	return magnet, nil
}

func (m *Magnet) PrintInfo() {
	fmt.Println("=== Magnet info ===")
	fmt.Printf("Display Name: %s\n", m.DisplayName)
	fmt.Printf("Exact Topic: %s\n", m.ExactTopic)
	for _, tr := range m.AddressTracker {
		fmt.Printf("Tracker: %s\n", tr)
	}
}
