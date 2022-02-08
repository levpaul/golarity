package magnet

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/levpaul/golarity/src/common"
)

const (
	paramDisplayName    = "dn"
	paramAddressTracker = "tr"
	paramExactTopic     = "xt"
)

type Magnet struct {
	DisplayName string // dn
	// TODO: For backwards compatibility with existing links, clients should also support the Base32 encoded version of the hash
	ExactTopic     common.Hash // xt - only accepting btih atm
	AddressTracker []string    // tr
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
			if magnet.ExactTopic != nil {
				return nil, fmt.Errorf("multiple 'Exact Topic' params found")
			}
			// Expect value like: 'urn:btih:E3D418E6B176F3E9FCF9193A78B6AD4DF1D656E4'
			if len(splitParam[1]) != 49 {
				return nil, fmt.Errorf("invalid exact topic supplied - accepting only urn:btih values")
			}
			magnet.ExactTopic = []byte(splitParam[1][9:])
		case paramAddressTracker:
			// This param supports multiple values
			decoded, err := url.QueryUnescape(splitParam[1])
			if err != nil {
				return nil, fmt.Errorf("could not decode address tracker '%s' - error '%w'", splitParam[1], err)
			}
			magnet.AddressTracker = append(magnet.AddressTracker, decoded)
		default:
			return nil, fmt.Errorf("unexpected param found, '%s'", splitParam[0])
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
