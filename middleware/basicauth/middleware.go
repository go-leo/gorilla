package basicauth

import (
	"context"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ctxKey struct{}

func FromContext(ctx context.Context) (string, bool) {
	value, ok := ctx.Value(ctxKey{}).(string)
	return value, ok
}

type Account struct {
	User     string
	Password string
}

type Accounts []Account

type options struct {
	realm string
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(o *options)

func defaultOptions() *options {
	return &options{
		realm: "Authorization Required",
	}
}

func Realm(realm string) Option {
	return func(o *options) {
		o.realm = realm
	}
}

func Middleware(accounts Accounts, opts ...Option) mux.MiddlewareFunc {
	opt := defaultOptions().apply(opts...)
	realm := "Basic realm=" + strconv.Quote(opt.realm)
	pairs := processAccounts(accounts)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, found := pairs.searchCredential(r.Header.Get("Authorization"))
			if !found {
				w.Header().Set("WWW-Authenticate", realm)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), ctxKey{}, user))
			next.ServeHTTP(w, r)
		})
	}
}

func processAccounts(accounts Accounts) authPairs {
	length := len(accounts)
	if length <= 0 {
		panic(errors.New("basicauth: empty list of authorized credentials"))
	}
	pairs := make(authPairs, 0, length)
	for _, account := range accounts {
		if account.User == "" {
			panic(errors.New("basicauth: user can not be empty"))
		}
		base := account.User + ":" + account.Password
		value := "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
		pairs = append(pairs, authPair{value: value, user: account.User})
	}
	return pairs
}

type authPair struct {
	value string
	user  string
}

type authPairs []authPair

func (a authPairs) searchCredential(authValue string) (string, bool) {
	if authValue == "" {
		return "", false
	}
	for _, pair := range a {
		if subtle.ConstantTimeCompare([]byte(pair.value), []byte(authValue)) == 1 {
			return pair.user, true
		}
	}
	return "", false
}
