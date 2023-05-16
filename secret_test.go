package main

import (
	"testing"
)

func TestTable_Secret_Enctypt_Descript(t *testing.T) {
	tests := []struct {
		name string
		s    Secret
	}{
		{
			"short string in secret",
			Secret("test"),
		},
		{
			"normal string",
			Secret("Lorem ipsum dolor sit amet, consetetur sadipscing elitr"),
		},
		{
			"long string with emty line",
			Secret(`Lorem ipsum dolor sit amet, consetetur sadipscing elitr,
            sed diam nonumy eirmod tempor invidunt ut labore et dolore magna

            aliquyam erat, sed diam voluptua. At vero eos et accusam et justo

            duo dolores et ea rebum. Stet clita kasd gubergren, no sea
            takimata sanctus est Lorem ipsum dolor sit amet.`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Enctypt()
			gote := tt.s
			tt.s.Descript()
			gotd := tt.s
			if tt.s != gotd {
				t.Errorf("Secret.Enctypt() = %v, want %v", gote, gotd)
			}
		})
	}
}
