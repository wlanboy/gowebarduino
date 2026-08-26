package model

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		command Command
		wantOK  bool
	}{
		{"empty call", Command{Call: ""}, false},
		{"valid call", Command{Call: "on"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := tt.command.Validate()
			if ok != tt.wantOK {
				t.Errorf("Validate() ok = %v, want %v", ok, tt.wantOK)
			}
		})
	}
}
