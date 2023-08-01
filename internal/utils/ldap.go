package utils

import (
	"crypto/tls"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func LDAPStartTLS(cfg *Config) error {
	l, err := ldap.DialURL("ldaps://ldap.example.com:636")
	if err != nil {
		log.Println("Attempted to dial LDAP server, Failed")
		return err
	}
	defer l.Close()
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Println("Attempted to start TLS connection, failed")
		return err
	}
	return nil
}
