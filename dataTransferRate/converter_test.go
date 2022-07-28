package dataTransferRate

import (
	"fmt"
	"testing"
)

func TestNewConverter_Convert_To(t *testing.T) {
	testCases := []struct {
		testCase     string
		transferRate string
		toUnit       string
		expected     float64
	}{
		{"convert from 1 Gbps to Mbps", "1Gbps", Mbps, 1000},
		{"convert from 1 Gbps to Mbps", "1 Gbps", Mbps, 1000},
		{"convert from 1 Gbps to Mbps", "1	Gbps", Mbps, 1000},
		{"convert from 1 Gbps to Mbps", " 1 Gbps", Mbps, 1000},
		{"convert from 1 Gbps to Mbps", "1 Gbps ", Mbps, 1000},

		{"convert from 1 Kbps to Kbps", "1 Kbps", Kbps, 1},
		{"convert from 2 Kbps to Kbps", "2 Kbps ", Kbps, 2},
		{"convert from 1000 Kbps to Mbps", "1000 Kbps", Mbps, 1},
		{"convert from 2000 Kbps to Mbps", "2000 Kbps ", Mbps, 2},
		{"convert from 100 Kbps to Mbps", "100 Kbps ", Mbps, 0.1},
		{"convert from 1000000 Kbps to Gbps", "1000000 Kbps", Gbps, 1},
		{"convert from 2000000 Kbps to Gbps", "2000000 Kbps ", Gbps, 2},
		{"convert from 100 Kbps to Gbps", "100 Kbps ", Gbps, 0.0001},
		{"convert from 0.0001 Kbps to Gbps", "0.0001 Kbps ", Gbps, 1e-10},
		{"convert from 1000000000 Kbps to Tbps", "1000000000 Kbps", Tbps, 1},
		{"convert from 2000000000 Kbps to Tbps", "2000000000 Kbps ", Tbps, 2},
		{"convert from 100 Kbps to Tbps", "100 Kbps ", Tbps, 1e-07},

		{"convert from 1 Mbps to Kbps", "1 Mbps", Kbps, 1000},
		{"convert from 2 Mbps to Kbps", "2 Mbps ", Kbps, 2000},
		{"convert from 0.6 Mbps to Kbps", "0.6 Mbps ", Kbps, 600},
		{"convert from 1 Mbps to Mbps", "1 Mbps", Mbps, 1},
		{"convert from 2 Mbps to Mbps", "2 Mbps ", Mbps, 2},
		{"convert from 1000 Mbps to Gbps", "1000 Mbps", Gbps, 1},
		{"convert from 2000 Mbps to Gbps", "2000 Mbps ", Gbps, 2},
		{"convert from 500 Mbps to Gbps", "500 Mbps", Gbps, 0.5},
		{"convert from 1000000 Mbps to Tbps", "1000000 Mbps", Tbps, 1},
		{"convert from 2000000 Mbps to Tbps", "2000000 Mbps ", Tbps, 2},
		{"convert from 100 Mbps to Tbps", "100 Mbps ", Tbps, 0.0001},

		{"convert from 1 Gbps to Kbps", "1 Gbps", Kbps, 1000000},
		{"convert from 2 Gbps to Kbps", "2 Gbps ", Kbps, 2000000},
		{"convert from 1 Gbps to Mbps", "1 Gbps", Mbps, 1000},
		{"convert from 2 Gbps to Mbps", "2 Gbps ", Mbps, 2000},
		{"convert from 0.1 Gbps to Mbps", "0.1 Gbps ", Mbps, 100},
		{"convert from 1 Gbps to Gbps", "1 Gbps", Gbps, 1},
		{"convert from 2 Gbps to Gbps", "2 Gbps ", Gbps, 2},
		{"convert from 1000 Gbps to Tbps", "1000 Gbps", Tbps, 1},
		{"convert from 2000 Gbps to Tbps", "2000 Gbps ", Tbps, 2},
		{"convert from 50 Gbps to Tbps", "50 Gbps ", Tbps, 0.05},

		{"convert from 1 Tbps to Kbps", "1 Tbps", Kbps, 1000000000},
		{"convert from 2 Tbps to Kbps", "2 Tbps ", Kbps, 2000000000},
		{"convert from 0.3 Tbps to Kbps", "0.3 Tbps ", Kbps, 3e+8},
		{"convert from 1 Tbps to Mbps", "1 Tbps", Mbps, 1000000},
		{"convert from 2 Tbps to Mbps", "2 Tbps ", Mbps, 2000000},
		{"convert from 1 Tbps to Gbps", "1 Tbps", Gbps, 1000},
		{"convert from 2 Tbps to Gbps", "2 Tbps ", Gbps, 2000},
		{"convert from 1 Tbps to Tbps", "1 Tbps", Tbps, 1},
		{"convert from 2 Tbps to Tbps", "2 Tbps ", Tbps, 2},
	}

	for _, tc := range testCases {
		got, err := NewConverter().Convert(tc.transferRate).To(tc.toUnit)

		if err != nil {
			t.Run(tc.testCase, func(t *testing.T) {
				t.Error(err)
			})
		}

		if got != tc.expected {
			t.Run(tc.testCase, func(t *testing.T) {
				t.Errorf("expected: %.v, got: %.v", tc.expected, got)
			})
		}
	}
}

func Example_kpbsToMbps() {
	inMbps, _ := NewConverter().Convert("1000 Kbps").To(Mbps)
	fmt.Println(inMbps)
	// Output:
	// 1
}

// Converts 1 Gbps in Mbps
func Example_gbpsToMbps() {
	inMbps, _ := NewConverter().Convert("1 Gbps").To(Mbps)
	fmt.Println(inMbps)
	// Output:
	// 1000
}
