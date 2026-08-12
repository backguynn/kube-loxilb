package k8s

import (
	"reflect"
	"testing"
)

func TestUnmarshalNetworkList(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []networkList
		wantErr bool
	}{
		{
			name:  "json array format",
			input: `[{"name":"multus-net","namespace":"default"}]`,
			want:  []networkList{{Name: "multus-net", Namespace: "default"}},
		},
		{
			name:  "plain single network",
			input: "loxilb/loxilb-nad2",
			want:  []networkList{{Name: "loxilb/loxilb-nad2"}},
		},
		{
			name:  "plain comma separated duplicates",
			input: "loxilb/loxilb-nad2,loxilb/loxilb-nad2",
			want: []networkList{
				{Name: "loxilb/loxilb-nad2"},
				{Name: "loxilb/loxilb-nad2"},
			},
		},
		{
			name:  "plain comma separated with spaces",
			input: " net-a , net-b , ",
			want: []networkList{
				{Name: "net-a"},
				{Name: "net-b"},
			},
		},
		{
			name:    "malformed json annotation",
			input:   `[{"name":"multus-net"}`,
			wantErr: true,
		},
		{
			name:    "empty annotation",
			input:   "   ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalNetworkList(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("expected %#v, got %#v", tt.want, got)
			}
		})
	}
}
