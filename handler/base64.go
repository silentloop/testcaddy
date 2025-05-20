package handler

import (
	"encoding/json"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/silentloop/testcaddy/crypto"
)

func init() {
	caddy.RegisterModule(Base64Handler{})
}

type Base64Handler struct {
	Mode string `json:"mode"` // "encode" or "decode"
}

func (Base64Handler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.base64",
		New: func() caddy.Module { return new(Base64Handler) },
	}
}

func (h *Base64Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	var body struct {
		Data string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return nil
	}

	switch h.Mode {
	case "encode":
		encoded := crypto.EncodeBase64URLString(body.Data)
		json.NewEncoder(w).Encode(map[string]string{"data": encoded})
	case "decode":
		decoded, err := crypto.DecodeBase64URLToString(body.Data)
		if err != nil {
			http.Error(w, "decode error: "+err.Error(), http.StatusBadRequest)
			return nil
		}
		json.NewEncoder(w).Encode(map[string]string{"data": decoded})
	default:
		http.Error(w, "invalid mode", http.StatusBadRequest)
	}

	return nil
}
