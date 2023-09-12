package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	randomPort := rand.Intn(65535-1024) + 1024
	opts.Listen = fmt.Sprintf(":%d", randomPort)

	ctx, cancel := context.WithCancel(context.Background())
	go run(ctx)
	time.Sleep(time.Millisecond * 500)
	defer cancel()

	testCases := []struct {
		name     string
		method   string
		endpoint string
		msg      string
		headers  map[string]string
	}{
		{"GET root", "GET", "/", "Hello, World!", map[string]string{"X-Test": "true"}},
		{"POST root", "POST", "/", "Post Message", map[string]string{"Authorization": "Bearer token"}},
		{"GET no headers", "GET", "/", "Hello, No Headers!", map[string]string{}},
		{"GET custom path", "GET", "/custom", "Hello, Custom Path!", map[string]string{"X-Custom": "true"}},
		{"GET multiple headers", "GET", "/", "Hello, Multiple Headers!",
			map[string]string{"X-Header1": "value1", "X-Header2": "value2"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opts.Message = tc.msg

			client := &http.Client{}
			req, err := http.NewRequest(tc.method, fmt.Sprintf("http://localhost:%d%s", randomPort, tc.endpoint), http.NoBody)
			assert.NoError(t, err)

			for k, v := range tc.headers {
				req.Header.Set(k, v)
			}

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			var echoResp map[string]interface{}
			err = json.Unmarshal(body, &echoResp)
			require.NoError(t, err)

			assert.Equal(t, tc.msg, echoResp["message"])
			assert.Equal(t, fmt.Sprintf("%s %s", tc.method, tc.endpoint), echoResp["request"])
			assert.Equal(t, fmt.Sprintf("localhost:%d", randomPort), echoResp["host"])

			for k, v := range tc.headers {
				assert.Equal(t, v, echoResp["headers"].(map[string]interface{})[k])
			}

			assert.NotNil(t, echoResp["remote_addr"])
		})
	}
}
