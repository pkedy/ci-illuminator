package illuminator

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTConfig struct {
	ClientID    string `json:"clientId"`
	Broker      string `json:"broker"`
	RootCA      string `json:"rootCA"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"privateKey"`
}

func ConnectMQTT(config *MQTTConfig) MQTT.Client {
	tlsConfig := getTLSConfig(config)

	opts := MQTT.NewClientOptions()
	opts.SetClientID(config.ClientID)
	if tlsConfig != nil {
		opts.SetTLSConfig(tlsConfig)
	}
	opts.AddBroker(config.Broker)
	opts.SetAutoReconnect(true)

	return MQTT.NewClient(opts)
}

func getTLSConfig(config *MQTTConfig) *tls.Config {
	roots, err := getCertPool(config.RootCA)
	if err != nil {
		log.Fatal(err)
	}

	clientCAs, err := getCertPool(config.Certificate)
	if err != nil {
		log.Fatal(err)
	}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(config.Certificate, config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		RootCAs:      roots,
		ClientCAs:    clientCAs,
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig
}

func getCertPool(pemPath string) (*x509.CertPool, error) {
	certs := x509.NewCertPool()

	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}
	certs.AppendCertsFromPEM(pemData)
	return certs, nil
}
