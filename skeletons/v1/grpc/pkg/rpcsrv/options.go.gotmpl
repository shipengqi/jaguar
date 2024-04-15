package rpcsrv

import (
	"fmt"
	"math"
	"time"

	"github.com/spf13/pflag"
)

const (
	infinity                           = time.Duration(math.MaxInt64)
	defaultMaxMsgSize                  = 4 << 20 // 4 * 1024 * 1024
	defaultMaxConcurrentStreams        = 100000
	defaultKeepAliveTime               = 30 * time.Second
	defaultConnectionIdleTime          = 10 * time.Second
	defaultMaxServerConnectionAgeGrace = 10 * time.Second
	defaultMiniKeepAliveTimeRate       = 2
	defaultMetricsPath                 = "/debug/metrics"
)

// TLSInfo contains configuration items related to certificate.
type TLSInfo struct {
	// CertFile is a file containing a PEM-encoded certificate
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

type Options struct {
	BindAddress           string          `json:"bind-address"          mapstructure:"bind-address"`
	BindPort              int             `json:"bind-port"             mapstructure:"bind-port"`
	MaxMsgSize            int             `json:"max-msg-size"          mapstructure:"max-msg-size"`
	MaxConcurrentStreams  int             `json:"max-con-streams"       mapstructure:"max-con-streams"`
	ServerCert            TLSInfo         `json:"tls"                   mapstructure:"tls"`
	Keepalive             time.Duration   `json:"keepalive"             mapstructure:"keepalive"`
	Timeout               time.Duration   `json:"timeout"               mapstructure:"timeout"`
	MaxConnectionAge      time.Duration   `json:"max-conn-age"          mapstructure:"max-conn-age"`
	MaxConnectionAgeGrace time.Duration   `json:"max-conn-age-grace"    mapstructure:"max-conn-age-grace"`
	UnaryInterceptors     []string        `json:"unary-interceptors"    mapstructure:"unary-interceptors"`
	StreamInterceptors    []string        `json:"stream-interceptors"   mapstructure:"stream-interceptors"`
	MetricsOptions        *MetricsOptions `json:"metrics"               mapstructure:"metrics"`
}

type MetricsOptions struct {
	Path        string `json:"path"            mapstructure:"path"`
	BindAddress string `json:"bind-address"    mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"       mapstructure:"bind-port"`
}

// NewOptions creates an Options object with default parameters.
func NewOptions() *Options {
	return &Options{
		BindAddress:           "0.0.0.0",
		BindPort:              8081,
		Keepalive:             defaultKeepAliveTime,
		Timeout:               infinity,
		MaxConnectionAge:      defaultMaxServerConnectionAgeGrace,
		MaxConnectionAgeGrace: defaultMaxServerConnectionAgeGrace,
		MaxMsgSize:            defaultMaxMsgSize,
		MaxConcurrentStreams:  defaultMaxConcurrentStreams,
		MetricsOptions: &MetricsOptions{
			BindAddress: "127.0.0.1",
			Path:        defaultMetricsPath,
		},
	}
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *Options) Validate() []error {
	var errs []error

	if o.BindPort < 0 || o.BindPort > 65535 {
		errs = append(
			errs,
			fmt.Errorf(
				"--bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
				o.BindPort,
			),
		)
	}
	errs = append(errs, o.MetricsOptions.Validate()...)

	return errs
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.BindAddress, "grpc.bind-address", o.BindAddress,
		"The IP address on which to serve the --grpc.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fs.IntVar(&o.BindPort, "grpc.bind-port", o.BindPort, "The port on which to serve grpc access. Set to zero to disable.")

	fs.IntVar(&o.MaxMsgSize, "grpc.max-msg-size", o.MaxMsgSize, "gRPC max message size.")
	fs.IntVar(&o.MaxConcurrentStreams, "grpc.max-con-streams", o.MaxConcurrentStreams, "gRPC number of concurrent streams.")
	fs.DurationVar(&o.Timeout, "grpc.timeout", o.Timeout, "gRPC connection timeout period.")
	fs.DurationVar(&o.Keepalive, "grpc.keepalive", o.Keepalive, "gRPC connection keepalive period.")
	fs.DurationVar(&o.MaxConnectionAge, "grpc.max-conn-age", o.MaxConnectionAge, "gRPC maximum time for connection.")
	fs.DurationVar(&o.MaxConnectionAgeGrace, "grpc.max-conn-age-grace", o.MaxConnectionAgeGrace,
		"An additive period after grpc.max-conn-age after which the connection will be forcibly closed.")

	fs.StringVar(&o.ServerCert.CertFile, "grpc.tls.cert-file", o.ServerCert.CertFile,
		"File containing the default x509 Certificate for gRPC. (CA cert, if any, concatenated "+
			"after server cert).")
	fs.StringVar(&o.ServerCert.KeyFile, "grpc.tls.private-key-file",
		o.ServerCert.KeyFile, ""+
			"File containing the default x509 private key matching --grpc.tls.cert-file.")
	// See:
	//    https://github.com/grpc-ecosystem/go-grpc-middleware
	//    https://github.com/open-telemetry/opentelemetry-go
	fs.StringSliceVar(&o.UnaryInterceptors, "grpc.unary-interceptors", o.UnaryInterceptors,
		"List of allowed unary interceptors for gRPC server, comma separated. If this list is empty default unary interceptors will be used.")
	fs.StringSliceVar(&o.StreamInterceptors, "grpc.stream-interceptors", o.StreamInterceptors,
		"List of allowed stream interceptors for gRPC server, comma separated. If this list is empty default stream interceptors will be used.")

	o.MetricsOptions.AddFlags(fs)
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (o *MetricsOptions) Validate() []error {
	var errors []error

	if o.BindPort < 0 || o.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--metrics.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
				o.BindPort,
			),
		)
	}

	return errors
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (o *MetricsOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.BindAddress, "metrics.bind-address", o.BindAddress,
		"The IP address on which to serve the --metrics.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fs.IntVar(&o.BindPort, "metrics.bind-port", o.BindPort,
		"The port on which to serve unsecured, unauthenticated access. Set an available port "+
			"to enable metrics for the gRPC server. Set to zero to disable.")
	fs.StringVar(&o.Path, "metrics.path", o.Path,
		"Set metrics path for the gRPC server.")
}
