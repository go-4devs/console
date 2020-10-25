package console_test

import (
	"testing"

	"gitoa.ru/go-4devs/console"
)

func TestFind(t *testing.T) {
	cases := map[string]string{
		"fdevs:console:test": "fdevs:console:test",
		"fd:c:t":             "fdevs:console:test",
		"fd::t":              "fdevs:console:test",
		"f:c:t":              "fdevs:console:test",
		"f:c:a":              "fdevs:console:arg",
	}

	for name, ex := range cases {
		res, err := console.Find(name)
		if err != nil {
			t.Errorf("expect <nil> err, got:%s", err)

			continue
		}

		if res.Name != ex {
			t.Errorf("expect: %s, got: %s", ex, res)
		}
	}
}
