package checkip

import (
	"log"
	"testing"
)

func TestRedactSecrets(test *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{" ", " "},
		{"key=a", "key=REDACTED"},
		{"key=", "key="},
		{"abckey=1234abcd", "abckey=REDACTED"},
		{`Get "https://api.shodan.io/shodan/host/209.141.33.65?key=iGaABCDEFGAtiZuH4ghsdAGH4T8LE9GW": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`,
			`Get "https://api.shodan.io/shodan/host/209.141.33.65?key=REDACTED": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`},
	}
	for _, t := range tests {
		got := redactSecrets(t.in)
		if got != t.out {
			log.Fatalf("got %s, wanted %s", got, t.out)
		}
	}
}
