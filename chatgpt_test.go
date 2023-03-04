package chatgpt

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_Chat(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		mock    func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request))
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success_new_client",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewClient("api", GPT35Turbo)
				require.NoError(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{
						"id": "chatcmpl-6pt58qjQYXr3kwBnaArWbyj9bKxar",
						"object": "chat.completion",
						"created": 1677824482,
						"model": "gpt-3.5-turbo-0301",
						"usage": {
							"prompt_tokens": 66,
							"completion_tokens": 337,
							"total_tokens": 403
						},
						"choices": [
							{
								"message": {
									"role": "assistant",
									"content": "Yes sir!"
								},
								"finish_reason": "stop",
								"index": 0
							}
						]
					}`))
				}
			},
			want: "Yes sir!",
		},
		{
			name: "success_new_with_constraint",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewClient("api", "", "You're a travel assistant that only know anything about travel and nothing else")
				require.NoError(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{
						"id": "chatcmpl-6pt58qjQYXr3kwBnaArWbyj9bKxar",
						"object": "chat.completion",
						"created": 1677824482,
						"model": "gpt-3.5-turbo-0301",
						"usage": {
							"prompt_tokens": 66,
							"completion_tokens": 337,
							"total_tokens": 403
						},
						"choices": [
							{
								"message": {
									"role": "assistant",
									"content": "Yes sir!"
								},
								"finish_reason": "stop",
								"index": 0
							}
						]
					}`))
				}
			},
			want: "Yes sir!",
		},
		{
			name: "error_success_but_no_answer",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewClient("api", "")
				require.NoError(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(`{
						"choices": []
					}`))
				}
			},
			wantErr: true,
		},
		{
			name: "error_http_status_not_ok",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewClient("api", "")
				require.NoError(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{}`))
				}
			},
			wantErr: true,
		},
		{
			name: "error_cant_decode_json",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewClient("api", "")
				require.NoError(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{`))
				}
			},
			wantErr: true,
		},
		{
			name: "error_timeout",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewCustom(&http.Client{
					Timeout: time.Millisecond * 200,
				}, "api", "")
				require.NoError(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					time.Sleep(time.Millisecond * 500)
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{}`))
				}
			},
			wantErr: true,
		},
		{
			name: "error_new_with_constraint",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewClient("", "")
				require.Error(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{}`))
				}
			},
			wantErr: true,
		},
		{
			name: "error_new_custom",
			mock: func(t *testing.T) (*Client, func(http.ResponseWriter, *http.Request)) {
				c, err := NewCustom(&http.Client{}, "", "")
				require.Error(t, err)
				return c, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(`{}`))
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, handler := tt.mock(t)
			srv := httptest.NewServer(http.HandlerFunc(handler))
			defer srv.Close()
			if c != nil {
				c.host = srv.URL
				got, err := c.Chat(tt.args.message)
				require.Equal(t, tt.wantErr, err != nil)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
