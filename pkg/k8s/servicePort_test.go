package k8s

import "testing"

func TestParseEndpointRole(t *testing.T) {
	tests := []struct {
		role string
		want int
	}{
		{role: "prefill", want: 1},
		{role: "decode", want: 2},
		{role: "normal", want: 0},
		{role: "", want: 0},
	}

	for _, tt := range tests {
		if got := parseEndpointRole(tt.role); got != tt.want {
			t.Fatalf("role %q: expected %d, got %d", tt.role, tt.want, got)
		}
	}
}

func TestParseEndpointUint16(t *testing.T) {
	if got := parseEndpointUint16("5557"); got != 5557 {
		t.Fatalf("expected 5557, got %d", got)
	}

	if got := parseEndpointUint16("70000"); got != 0 {
		t.Fatalf("expected invalid uint16 parse to return 0, got %d", got)
	}
}

func TestParseEndpointWeight(t *testing.T) {
	tests := []struct {
		value string
		want  uint8
	}{
		{value: "", want: 1},
		{value: "5", want: 5},
		{value: "0", want: 1},
		{value: "99", want: 10},
		{value: "bad", want: 1},
	}

	for _, tt := range tests {
		if got := parseEndpointWeight(tt.value); got != tt.want {
			t.Fatalf("value %q: expected %d, got %d", tt.value, tt.want, got)
		}
	}
}

func TestMakeEndpointEntriesDefaults(t *testing.T) {
	entries := []EndpointEntry{
		{IP: "10.0.0.1", Weight: 1},
	}

	if entries[0].Weight != 1 {
		t.Fatalf("expected default weight 1, got %d", entries[0].Weight)
	}
}
