package server

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	settings "github.com/hjcalderon10/bunny-backend/setting"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	settings.Commons.Host = "localhost"
	settings.Commons.Port = "8080"
	go Start()
	time.Sleep(200 * time.Millisecond)
	resp, err := http.Get("http://localhost:8080/api/v1/health-check")
	assert.Nil(t, err, fmt.Sprintf("%s", err))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
