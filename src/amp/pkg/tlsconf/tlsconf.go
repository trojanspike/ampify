// Public Domain (-) 2010-2011 The Ampify Authors.
// See the Ampify UNLICENSE file for details.

// The tls package provides utilities to support TLS connections. When the
// package is initialised via ``tlsconf.Init()``, it generates a default
// configuration from the TLS Certificate Authorities data included with Ampify.
package tlsconf

import (
	"amp/runtime"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"time"
)

var Config *tls.Config

func GenConfig(file string) (config *tls.Config, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(data)
	config = &tls.Config{
		Rand:    rand.Reader,
		Time:    time.Now,
		RootCAs: roots,
	}
	return config, nil
}

// Set the ``tlsconf.Config`` variable.
func Init() {
	path := runtime.AmpifyRoot + "/environ/local/share/cacerts/ca.cert"
	var err error
	Config, err = GenConfig(path)
	if err != nil {
		runtime.Error("ERROR: Couldn't load %s: %s\n", path, err)
	}
}
