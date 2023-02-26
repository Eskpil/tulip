package ca

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"time"
)

func store(publicKey []byte, publicKeyPath string, privateKey []byte, privateKeyPath string) error {
	certOut, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: publicKey}); err != nil {
		return err
	}

	if err := certOut.Close(); err != nil {
		return err
	}

	// Private key
	keyOut, err := os.OpenFile(privateKeyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKey}); err != nil {
		return err
	}

	if err := keyOut.Close(); err != nil {
		return err
	}

	return nil
}

func getCA() (*x509.Certificate, crypto.PrivateKey, error) {
	_, err := os.Stat(os.ExpandEnv("$TULIP_PKI_DIRECTORY/ca.crt"))
	if err != nil {
		log.Infof("CA certificates does not exist, creating them")

		ca := &x509.Certificate{
			SerialNumber: big.NewInt(1653),
			Subject: pkix.Name{
				Organization: []string{"Tulip"},
				Country:      []string{"Norway"},
				CommonName:   "Tulip - CA",
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(10, 0, 0),
			IsCA:                  true,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
			BasicConstraintsValid: true,
		}

		privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
		publicKey := &privateKey.PublicKey
		caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, publicKey, privateKey)
		if err != nil {
			return nil, nil, err
		}

		if err := store(caBytes, os.ExpandEnv("$TULIP_PKI_DIRECTORY/ca.crt"), x509.MarshalPKCS1PrivateKey(privateKey), os.ExpandEnv("$TULIP_PKI_DIRECTORY/ca.key")); err != nil {
			return nil, nil, err
		}

		return nil, nil, nil
	}

	// Load CA
	catls, err := tls.LoadX509KeyPair(os.ExpandEnv("$TULIP_PKI_DIRECTORY/ca.crt"), os.ExpandEnv("$TULIP_PKI_DIRECTORY/ca.key"))
	if err != nil {
		panic(err)
	}

	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}

	return ca, catls.PrivateKey, nil
}

func CreatePair(entity string) (string, string, error) {
	pkiDir := os.Getenv("TULIP_PKI_DIRECTORY")
	if pkiDir == "" {
		return "", "", fmt.Errorf("environment variable \"TULIP_PKI_DIRECTORY\" is not set")
	}

	privateKeyPath := fmt.Sprintf("%s/%s/cert.key", pkiDir, entity)
	publicKeyPath := fmt.Sprintf("%s/%s/cert.crt", pkiDir, entity)

	if _, err := os.Stat(fmt.Sprintf("%s/%s/", pkiDir, entity)); err == nil {
		log.Infof("Reusing existing certificates for: (%s)", entity)

		publicKeyBytes, err := os.ReadFile(publicKeyPath)
		if err != nil {
			return "", "", err
		}

		privateKeyBytes, err := os.ReadFile(publicKeyPath)
		if err != nil {
			return "", "", err
		}

		return string(publicKeyBytes), string(privateKeyBytes), err
	}

	ca, caPrivateKey, err := getCA()
	if err != nil {
		return "", "", err
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Tulip"},
			Country:      []string{"Norway"},
			CommonName:   fmt.Sprintf("Tulip - %s", entity),
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, publicKey, caPrivateKey)

	certificateOut := bytes.NewBuffer([]byte{})
	if err := pem.Encode(certificateOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}); err != nil {
		return "", "", err
	}

	privateKeyOut := bytes.NewBuffer([]byte{})
	if err := pem.Encode(privateKeyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}); err != nil {
		return "", "", err
	}

	entityPath := fmt.Sprintf("%s/%s", os.Getenv("TULIP_PKI_DIRECTORY"), entity)
	if _, err := os.Stat(entityPath); err != nil {
		_ = os.Mkdir(entityPath, 0700)
	}

	if err := store(certBytes, publicKeyPath, x509.MarshalPKCS1PrivateKey(privateKey), privateKeyPath); err != nil {
		return "", "", err
	}

	return certificateOut.String(), privateKeyOut.String(), nil
}
