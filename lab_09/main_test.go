package main

import (
	"strings"
	"testing"
)

func TestToYAML(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name: "Valid Server Struct",
			input: Server{
				Host:       "localhost",
				Port:       8080,
				Debug:      true,
				AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
			},
			want:    "host: localhost",
			wantErr: false,
		},
		{
			name:    "Simple string",
			input:   "just a string",
			want:    "just a string",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToYAML(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ToYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !strings.Contains(got, tt.want) {
				t.Errorf("ToYAML() result doesn't contain expected string.\nGot:\n%v\nWanted string: %v", got, tt.want)
			}
		})
	}
}

func BenchmarkToYAML(b *testing.B) {
	srv := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ToYAML(srv)
	}
}

func BenchmarkToJSON(b *testing.B) {
	srv := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ToJSON(srv)
	}
}
