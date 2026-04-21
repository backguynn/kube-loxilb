package api

import "testing"

func TestEpSelectEnterpriseAlignment(t *testing.T) {
	if LbSelN3 != 6 {
		t.Fatalf("expected LbSelN3 to be 6, got %d", LbSelN3)
	}

	if LbSelCHWBL != 8 {
		t.Fatalf("expected LbSelCHWBL to be 8, got %d", LbSelCHWBL)
	}

	if LbSelGPUAware != 9 {
		t.Fatalf("expected LbSelGPUAware to be 9, got %d", LbSelGPUAware)
	}

	if LbSelWRRHash != 10 {
		t.Fatalf("expected LbSelWRRHash to be 10, got %d", LbSelWRRHash)
	}
}

func TestBackendProtocolTypeValidation(t *testing.T) {
	tests := []struct {
		name     string
		protocol BackendProtocolType
		want     bool
	}{
		{name: "empty", protocol: "", want: true},
		{name: "http1", protocol: BackendProtocolHTTP1, want: true},
		{name: "http2", protocol: BackendProtocolHTTP2, want: true},
		{name: "both", protocol: BackendProtocolBoth, want: true},
		{name: "invalid", protocol: "grpc", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.protocol.IsValid(); got != tt.want {
				t.Fatalf("expected IsValid()=%v, got %v", tt.want, got)
			}
		})
	}
}
