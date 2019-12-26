// +build go1.11

package cookie

import (
	"github.com/gin-contrib/sessions"
	gsessions "github.com/gorilla/sessions"
)

type Store interface {
	sessions.Store
}

// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewStore(keyPairs ...[]byte) Store {
	return &store{gsessions.NewCookieStore(keyPairs...)}
}

type store struct {
	*gsessions.CookieStore
}

func (c *store) Options(options sessions.Options) {
	c.CookieStore.Options = &gsessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
		SameSite: options.SameSite,
	}
}
