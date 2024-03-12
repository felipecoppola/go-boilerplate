package httpjson_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipecoppola/go-boilerplate/cmd/app/handler/httpjson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"gopkg.in/h2non/gentleman.v2"
)

func Test_healthRequest(t *testing.T) {
	app, _ := httpjson.New(zap.L())
	server := httptest.NewServer(app)
	t.Cleanup(server.Close)

	c := gentleman.New()
	c.BaseURL(server.URL)

	t.Run("Check health check", func(t *testing.T) {
		r, err := c.
			Get().
			AddPath("/health").
			Do()

		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
	})
}
