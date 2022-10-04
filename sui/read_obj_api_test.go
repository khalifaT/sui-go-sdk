package sui

import (
	"context"
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{
			name:    "user 1",
			address: "0xecadb4fa212afbe396ddc8fe0b4ae559a4cc5cc6",
			wantErr: false,
		},
		{
			name:    "unexistant user",
			address: "0xecadb4fa212afbe396ddc8fe0b4ae559a4cc5ccc",
			wantErr: true,
		},
	}
	cli := NewSuiClient("https://gateway.devnet.sui.io:443")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := cli.GetBalance(context.Background(), tt.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("balance error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("\nbalance of the user %v is %v\n", tt.address, b)

		})
	}
}
