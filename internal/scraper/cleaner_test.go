package scraper

import "testing"

func TestParsePrice(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		// Standard German format
		{"German Standard", "1200,50 €", 1200.50, false},
		{"German Simple", "1200,00", 1200.00, false},

		// Messy whitespace/symbols
		{"Messy Whitespace", "\n  € 45,90 \t", 45.90, false},
		{"Just Number", "89", 89.00, false},

		// The "Remove Dots" rule
		{"Thousands Separator", "1.500", 1500.00, false},

		// Error cases
		{"Empty", "", 0, true},
		{"Text Only", "Sold Out", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePrice(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("ParsePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
