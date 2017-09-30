package httpapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	r := chi.NewRouter()
	r.HandleFunc("/", Index)

	s := httptest.NewServer(r)
	defer s.Close()

	testCases := map[string]struct {
		code int
		body string
	}{
		"basic get": {
			code: http.StatusOK,
			body: "Watchtower",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/", s.URL))
			assert.Nil(t, err)
			assert.Equal(t, testCase.code, res.StatusCode)

			defer res.Body.Close()
			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			assert.Contains(t, buf.String(), testCase.body)
		})
	}
}
