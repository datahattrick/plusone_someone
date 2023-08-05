package utils

import (
	"crypto/tls"
	"log"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func LDAPStartTLS(cfg *Config) (*ldap.SearchResult, error) {
	log.Println("LDAP: Dialing Ldap server")
	l, err := ldap.DialURL(cfg.LDAP.Host)
	if err != nil {
		log.Println("Attempted to dial LDAP server, Failed")
		return nil, err
	}
	defer l.Close()
	if strings.HasPrefix(cfg.LDAP.Host, "ldaps") {
		log.Println("LDAP: Attempting to start TLS connection")
		err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			log.Println("Attempted to start TLS connection, failed")
			return nil, err
		}
	}

	log.Println("LDAP: Connection to bind account: ", cfg.LDAP.Bind)
	err = connectionBind(l, cfg)
	if err != nil {
		return nil, err
	}

	log.Println("LDAP: Searching for :", cfg.LDAP.Attribute)
	sr, err := searchRequest(l, cfg)
	if err != nil {
		return nil, err
	}

	log.Printf("LDAP: Search complete found %v users", len(sr.Entries))

	return sr, nil
}

func connectionBind(l *ldap.Conn, cfg *Config) error {
	err := l.Bind(cfg.LDAP.Bind, cfg.LDAP.Password)
	if err != nil {
		log.Println("Unable to Bind to provided account")
		return err
	}
	return nil
}

func searchRequest(l *ldap.Conn, cfg *Config) (*ldap.SearchResult, error) {
	searchRequest := ldap.NewSearchRequest(
		cfg.LDAP.UserGroup,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",
		cfg.LDAP.Attribute, nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Printf("Unable to find users at the location %s", cfg.LDAP.UserGroup)
		return nil, err
	}
	return sr, nil
}
