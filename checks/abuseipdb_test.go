package checks

import (
	"fmt"
	"github.com/jreisinger/checkip/check"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"testing"
)

func TestCheckAbuseIPDB(t *testing.T) {
	t.Run("given valid response then result and no error is returned", func(t *testing.T) {
		handlerFn := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			rw.Write(LoadResponse(t, "abuseipdb_response.json"))
		})
		setAbuseIPDBConfig(t)
		testUrl := SetMockHttpClient(t, handlerFn)
		setAbuseIPDBUrl(t, testUrl)

		result, err := CheckAbuseIPDB(net.ParseIP("118.25.6.39"))
		require.NoError(t, err)
		assert.Equal(t, "abuseipdb.com", result.Name)
		assert.Equal(t, check.TypeInfoSec, result.Type)
		assert.Equal(t, true, result.Malicious)
		assert.Equal(t, "domain: tencent.com, usage type: Data Center/Web Hosting/Transit", result.Info.Summary())
	})

	t.Run("given non 2xx response then error is returned", func(t *testing.T) {
		handlerFn := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
		})
		setAbuseIPDBConfig(t)
		testUrl := SetMockHttpClient(t, handlerFn)
		setAbuseIPDBUrl(t, testUrl)

		_, err := CheckAbuseIPDB(net.ParseIP("118.25.6.39"))
		require.Error(t, err)
	})
}

// --- test helpers ---

func setAbuseIPDBUrl(t *testing.T, testUrl string) {
	url := abuseIPDBUrl
	abuseIPDBUrl = testUrl
	t.Cleanup(func() {
		abuseIPDBUrl = url
	})
}

func setAbuseIPDBConfig(t *testing.T) {
	SetMockConfig(t, func(key string) (string, error) {
		if key == "ABUSEIPDB_API_KEY" {
			return "123-secret-789", nil
		}
		return "", fmt.Errorf("unexpected key %s received", key)
	})
}
