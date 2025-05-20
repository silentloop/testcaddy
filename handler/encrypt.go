package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"

	"github.com/silentloop/testcaddy/crypto"
)

func init() {
	caddy.RegisterModule(EncryptHandler{})
}

type EncryptHandler struct {
	Key string `json:"key"`
	IV  string `json:"iv"`
}

func (EncryptHandler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.encrypt",
		New: func() caddy.Module { return &EncryptHandler{} }, // 用指標
	}
}

func (h *EncryptHandler) Provision(ctx caddy.Context) error {
	return nil
}

func (h EncryptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return nil
	}
	plaintext := data["data"]
	ciphertext, iv := crypto.EncryptAES([]byte(h.Key), []byte(plaintext))
	resp := map[string]string{
		"cipher": ciphertext,
		"iv":     string(iv),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	return nil
}

func (h *EncryptHandler) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for d.NextBlock(0) {
			switch d.Val() {
			case "key":
				if !d.Args(&h.Key) {
					return d.ArgErr()
				}
			case "iv":
				if !d.Args(&h.IV) {
					return d.ArgErr()
				}
			default:
				return d.Errf("unexpected token: %s", d.Val())
			}
		}
	}
	fmt.Printf("🔐 encrypt config parsed: key=%s iv=%s\n", h.Key, h.IV)
	return nil
}
