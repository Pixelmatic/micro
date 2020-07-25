package custom

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/internal/helper"
	"github.com/micro/micro/v2/internal/namespace"
)

var (
	// ErrUnauthorized is returned by Authorize when a context without a blank account tries to access
	// a restricted namespace
	ErrUnauthorized = errors.New("An account is required")
	// ErrForbidden is returned by Authorize when a context is trying to access a namespace it doesn't
	// have access to
	ErrForbidden = errors.New("Access denied to namespace")
)

const (
	// DefaultNamespace used by the server
	DefaultNamespace = "micro"
	// NamespaceKey is used to set/get the namespace from the context
	NamespaceKey = "Micro-Namespace"
)

func TLSConfig(ctx *cli.Context) (*tls.Config, error) {
	return helper.TLSConfig(ctx)
}

func ACMEHosts(ctx *cli.Context) []string {
	return helper.ACMEHosts(ctx)
}

func Authorize(ctx context.Context, ns string, opts ...namespace.AuthorizeOption) error {
	return namespace.Authorize(ctx, ns, opts...)
}

func Public(ns string) namespace.AuthorizeOption {
	return namespace.Public(ns)
}
