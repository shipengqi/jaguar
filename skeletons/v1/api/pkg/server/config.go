package server

import (
	"net"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// RecommendedHomeDir defines the default directory used to place all idm service configurations.
	RecommendedHomeDir = ".idm"

	// RecommendedEnvPrefix defines the ENV prefix used by all idm service.
	RecommendedEnvPrefix = "IDM"

	// RecommendedTokenTimeout defines the duration that a jwt token is valid.
	RecommendedTokenTimeout = 30 * time.Minute

	// RecommendedMaxRefresh defines the duration that a jwt token can be refreshed.
	RecommendedMaxRefresh = 12 * time.Hour
)

// Config is a structure used to configure a GenericAPIServer.
type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   false,
		Jwt: &JwtInfo{
			Realm:      "idm jwt",
			Timeout:    RecommendedTokenTimeout,
			MaxRefresh: RecommendedMaxRefresh,
		},
	}
}

// SecureServingInfo holds configuration of the TLS server.
type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	TLSInfo     TLSInfo
}

// Address join host IP address and host port number into an address string, like: 0.0.0.0:8443.
func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

// TLSInfo contains configuration items related to certificate.
type TLSInfo struct {
	// CertFile is a file containing a PEM-encoded certificate
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

// InsecureServingInfo holds configuration of the insecure http server.
type InsecureServingInfo struct {
	Address string
}

// JwtInfo defines jwt fields used to create jwt authentication middleware.
type JwtInfo struct {
	// defaults to "idm jwt"
	Realm string
	// defaults to empty
	Key string
	// defaults to one hour
	Timeout time.Duration
	// defaults to zero
	MaxRefresh time.Duration
}
