package smbios

import (
	"reflect"
	"testing"
	"time"
)

func Test_biosTime(t *testing.T) {

	tests := []struct {
		str  string
		want time.Time
	}{
		{
			str:  "10/24/2013",
			want: time.Date(2013, time.October, 24, 0, 0, 0, 0, time.UTC),
		}, {
			str:  "10/24/99",
			want: time.Date(1999, time.October, 24, 0, 0, 0, 0, time.UTC),
		}, {
			str:  "10/24/74",
			want: time.Date(1974, time.October, 24, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := biosTime(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("biosTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
