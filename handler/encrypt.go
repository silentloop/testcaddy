package handler

import (
	"encoding/json"
	"net/http"

	"github.com/caddyserver/caddy/v2"
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
		New: func() caddy.Module { return new(EncryptHandler) },
	}
}

func (h *EncryptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	var body struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return nil
	}

	encrypted, err := crypto.EncryptAES([]byte(h.Key), []byte(h.IV), body.Data)
	if err != nil {
		http.Error(w, "encryption error: "+err.Error(), http.StatusInternalServerError)
		return nil
	}

	json.NewEncoder(w).Encode(map[string]string{"data": encrypted})
	return nil
}
