package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func NewTLSConfig() *tls.Config {
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("configs/mqtt/ca.crt")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair("configs/mqtt/server.crt", "configs/mqtt/server.key")
	if err != nil {
		panic(err)
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}
}
