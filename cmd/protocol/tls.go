package protocol

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"
)

func tlsCertConfig() (*tls.Config, error) {
	cwd, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current working directory: %v\n", err)
	}

	execPwd := filepath.Dir(cwd)

	cert, err := tls.LoadX509KeyPair(execPwd+"/configs/ca.crt.pem", execPwd+"/configs/ca.key.pem")
	if err != nil {
		return nil, fmt.Errorf("Failed to load certificate %v\n", err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return &config, nil
}

func tlsNonCertConfig() *tls.Config {
	return &tls.Config{InsecureSkipVerify: true}
}
