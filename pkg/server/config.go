// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package server

import (
	"time"

	"storj.io/gateway-mt/pkg/authclient"
	"storj.io/gateway/miniogw"
)

// AddrConfig honestly only exists to preserve legacy CLI parameter naming.
type AddrConfig struct {
	Address string `help:"Address to serve gateway on" default:"127.0.0.1:7777"`
}

// Config determines how server listens for requests.
type Config struct {
	Server               AddrConfig
	CertDir              string   `help:"directory path to search for TLS certificates" default:"$CONFDIR/certs"`
	InsecureDisableTLS   bool     `help:"listen using insecure connections" releaseDefault:"false" devDefault:"true"`
	DomainName           string   `help:"comma-separated domain suffixes to serve on" releaseDefault:"" devDefault:"localhost"`
	CorsOrigins          string   `help:"list of domains (comma separated) other than the gateway's domain, from which a browser should permit loading resources requested from the gateway" default:"*"`
	EncodeInMemory       bool     `help:"tells libuplink to perform in-memory encoding on file upload" releaseDefault:"true" devDefault:"true"`
	ClientTrustedIPSList []string `help:"list of clients IPs (without port and comma separated) which are trusted; usually used when the service run behinds gateways, load balancers, etc."`
	UseClientIPHeaders   bool     `help:"use the headers sent by the client to identify its IP. When true the list of IPs set by --client-trusted-ips-list, when not empty, is used" default:"true"`
	InsecureLogAll       bool     `help:"insecurely log all errors, paths, and headers" default:"false"`

	Auth            authclient.Config
	S3Compatibility miniogw.S3CompatibilityConfig
	Client          ClientConfig
	ConnectionPool  ConnectionPoolConfig
}

// ConnectionPoolConfig is a config struct for configuring RPC connection pool
// options.
type ConnectionPoolConfig struct {
	Capacity       int           `help:"RPC connection pool capacity" default:"100"`
	KeyCapacity    int           `help:"RPC connection pool key capacity" default:"5"`
	IdleExpiration time.Duration `help:"RPC connection pool idle expiration" default:"2m0s"`
}

// ClientConfig is a configuration struct for the uplink that controls how
// to talk to the rest of the network.
type ClientConfig struct {
	DialTimeout time.Duration `help:"timeout for dials" default:"0h2m00s"`
	UseQosAndCC bool          `help:"use congestion control and QOS settings" default:"true"`
}
