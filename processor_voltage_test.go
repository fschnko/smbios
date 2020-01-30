package smbios

import "testing"

func TestProcessorVoltage_Value(t *testing.T) {
	tests := []struct {
		name string
		v    ProcessorVoltage
		want float32
	}{
		{
			name: "Empty value",
			v:    0x00,
			want: 0.0,
		}, {
			name: "Legacy mode 5V",
			v:    0x01,
			want: 5.0,
		}, {
			name: "Legacy mode 3.3V",
			v:    0x02,
			want: 3.3,
		}, {
			name: "Legacy mode 2.9V",
			v:    0x04,
			want: 2.9,
		}, {
			name: "Normal mode 0",
			v:    0x80,
			want: 0.0,
		}, {
			name: "Normal mode 5V",
			v:    0xb2,
			want: 5.0,
		}, {
			name: "Normal mode 3.3V",
			v:    0xa1,
			want: 3.3,
		}, {
			name: "Normal mode 2.9V",
			v:    0x9d,
			want: 2.9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Value(); got != tt.want {
				t.Errorf("ProcessorVoltage.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
