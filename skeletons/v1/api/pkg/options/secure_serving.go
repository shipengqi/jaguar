package options

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"

	"github.com/jaguar/apiskeleton/pkg/server"
)

// SecureServingOptions contains configuration items related to HTTPS server startup.
type SecureServingOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	// BindPort is ignored when Listener is set, will serve HTTPS even with 0.
	BindPort int `json:"bind-port"    mapstructure:"bind-port"`
	// Required to be set to true means that BindPort cannot be zero.
	Required bool
	// ServerCert is the TLS cert info for serving secure traffic
	ServerCert TLSInfo `json:"tls"          mapstructure:"tls"`
}

// TLSInfo contains configuration items related to certificate.
type TLSInfo struct {
	// CertFile is a file containing a PEM-encoded certificate
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}

// NewSecureServingOptions creates a SecureServingOptions object with default parameters.
func NewSecureServingOptions() *SecureServingOptions {
	return &SecureServingOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8443,
		Required:    false,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *SecureServingOptions) ApplyTo(c *server.Config) error {
	// SecureServing is required to serve https
	c.SecureServing = &server.SecureServingInfo{
		BindAddress: s.BindAddress,
		BindPort:    s.BindPort,
		TLSInfo: server.TLSInfo{
			CertFile: s.ServerCert.CertFile,
			KeyFile:  s.ServerCert.KeyFile,
		},
	}

	return nil
}

// Validate is used to parse and validate the parameters entered by the user at
// the command line when the program starts.
func (s *SecureServingOptions) Validate() []error {
	if s == nil {
		return nil
	}

	var errs []error

	if s.Required {
		if s.BindPort < 1 || s.BindPort > 65535 {
			errs = append(
				errs,
				fmt.Errorf(
					"--secure.bind-port %v must be between 1 and 65535, inclusive. It cannot be turned off with 0",
					s.BindPort,
				),
			)
		}
		if s.ServerCert.CertFile == "" || s.ServerCert.KeyFile == "" {
			errs = append(
				errs,
				errors.New("--secure.tls.cert-file and --secure.tls.private-key-file must be set"),
			)
		}
	} else if s.BindPort < 0 || s.BindPort > 65535 {
		errs = append(
			errs,
			fmt.Errorf(
				"--secure.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off secure port",
				s.BindPort,
			),
		)
	}

	return errs
}

// AddFlags adds flags related to HTTPS server for a specific APIServer to the
// specified FlagSet.
func (s *SecureServingOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "secure.bind-address", s.BindAddress,
		"The IP address on which to listen for the --secure.bind-port port. The "+
			"associated interface(s) must be reachable by the rest of the engine, and by CLI/web "+
			"clients. If blank, all interfaces will be used (0.0.0.0 for all IPv4 interfaces and :: for all IPv6 interfaces).")
	desc := "The port on which to serve HTTPS with authentication and authorization."
	if s.Required {
		desc += " It cannot be switched off with 0."
	} else {
		desc += " If 0, don't serve HTTPS at all."
	}
	fs.IntVar(&s.BindPort, "secure.bind-port", s.BindPort, desc)

	// Todo CA cert
	fs.StringVar(&s.ServerCert.CertFile, "secure.tls.cert-file", s.ServerCert.CertFile,
		"File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated "+
			"after server cert).")

	fs.StringVar(&s.ServerCert.KeyFile, "secure.tls.private-key-file",
		s.ServerCert.KeyFile, ""+
			"File containing the default x509 private key matching --secure.tls.cert-file.")
}
