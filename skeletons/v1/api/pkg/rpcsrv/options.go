package rpcsrv

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

// TLSInfo contains configuration items related to certificate.
type TLSInfo struct {
	// CertFile is a file containing a PEM-encoded certificate
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

type Options struct {
	BindAddress           string        `json:"bind-address"          mapstructure:"bind-address"`
	BindPort              int           `json:"bind-port"             mapstructure:"bind-port"`
	ServerCert            TLSInfo       `json:"tls"                   mapstructure:"tls"`
	Keepalive             time.Duration `json:"keepalive"             mapstructure:"keepalive"`
	Timeout               time.Duration `json:"timeout"               mapstructure:"timeout"`
	MaxConnectionAge      time.Duration `json:"max-conn-age"          mapstructure:"max-conn-age"`
	MaxConnectionAgeGrace time.Duration `json:"max-conn-age-grace"    mapstructure:"max-conn-age-grace"`
	MaxMsgSize            int           `json:"max-msg-size"          mapstructure:"max-msg-size"`
	MaxConcurrentStreams  int           `json:"max-con-streams"       mapstructure:"max-con-streams"`
}

func NewOptions() *Options {
	return &Options{
		BindAddress: "0.0.0.0",
		BindPort:    8081,
		MaxMsgSize:  4 * 1024 * 1024,
	}
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (s *Options) Validate() []error {
	var errors []error

	if s.BindPort < 0 || s.BindPort > 65535 {
		errors = append(
			errors,
			fmt.Errorf(
				"--insecure-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port",
				s.BindPort,
			),
		)
	}

	return errors
}

// AddFlags adds flags related to features for a specific api server to the
// specified FlagSet.
func (s *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "grpc.bind-address", s.BindAddress,
		"The IP address on which to serve the --grpc.bind-port(set to 0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	fs.IntVar(&s.BindPort, "grpc.bind-port", s.BindPort, "The port on which to serve grpc access. Set to zero to disable.")

	fs.IntVar(&s.MaxMsgSize, "grpc.max-msg-size", s.MaxMsgSize, "gRPC max message size.")
	fs.IntVar(&s.MaxConcurrentStreams, "grpc.max-con-streams", s.MaxConcurrentStreams, "gRPC number of concurrent streams.")
	fs.DurationVar(&s.Timeout, "grpc.timeout", s.Timeout, "gRPC connection timeout period.")
	fs.DurationVar(&s.Keepalive, "grpc.keepalive", s.Keepalive, "gRPC connection keepalive period.")
	fs.DurationVar(&s.MaxConnectionAge, "grpc.max-conn-age", s.MaxConnectionAge, "gRPC maximum time for connection.")
	fs.DurationVar(&s.MaxConnectionAgeGrace, "grpc.max-conn-age-grace", s.MaxConnectionAgeGrace,
		"An additive period after grpc.max-conn-age after which the connection will be forcibly closed.")

	fs.StringVar(&s.ServerCert.CertFile, "grpc.tls.cert-file", s.ServerCert.CertFile,
		"File containing the default x509 Certificate for gRPC. (CA cert, if any, concatenated "+
			"after server cert).")
	fs.StringVar(&s.ServerCert.KeyFile, "grpc.tls.private-key-file",
		s.ServerCert.KeyFile, ""+
			"File containing the default x509 private key matching --grpc.tls.cert-file.")
}
