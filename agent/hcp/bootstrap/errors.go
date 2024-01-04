package bootstrap

import "errors"

var (
	ErrFetchBootstrapUnauthorized = errors.New("unauthorized error when fetching bootstrap config")
	ErrFetchBootstrapForbidden    = errors.New("forbidden error when fetching bootstrap config")
)
