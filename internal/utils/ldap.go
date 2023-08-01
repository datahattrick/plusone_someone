package utils

import (
	"crypto/tls"
	"fmt"
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

	err = connectionBind(l, cfg)
	if err != nil {
		return err
	}

	err = searchRequest(l, cfg)
	if err != nil {
		return err
	}

	return nil
}

func connectionBind(l *ldap.Conn, cfg *Config) error {
	err := l.Bind(cfg.LDAP.Bind, cfg.LDAP.Password)
	if err != nil {
		log.Println("Unable to Bind to provided account")
		return err
	}
	return nil
}

func searchRequest(l *ldap.Conn, cfg *Config) error {
	searchRequest := ldap.NewSearchRequest(
		cfg.LDAP.UserGroup,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",
		[]string{"dn", "cn"}, nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Printf("Unable to find users at the location %s", cfg.LDAP.UserGroup)
		return err
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
	return nil
}
