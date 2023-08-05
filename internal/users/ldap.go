package users

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/datahattrick/plusone_someone/internal/database"
	"github.com/datahattrick/plusone_someone/internal/utils"
	"github.com/go-ldap/ldap/v3"
	"github.com/google/uuid"
)

func CreateUsersFromLDAP(sr *ldap.SearchResult, cfg *utils.Config) {
	var ur []*User
	for _, entry := range sr.Entries {
		user, ok := MapLDAPUsers(entry, cfg)
		if !ok {
			continue
		}
		ur = append(ur, user)
	}
	log.Printf("Synced %v users to DB", len(ur))
}

func MapLDAPUsers(entry *ldap.Entry, cfg *utils.Config) (*User, bool) {
	params := Userparams{
		FirstName: entry.GetAttributeValue(cfg.LDAP.FirstNameAttr),
		LastName:  entry.GetAttributeValue(cfg.LDAP.LastNameAttr),
		Email:     entry.GetAttributeValue(cfg.LDAP.EmailAttr),
		Username:  entry.GetAttributeValue(cfg.LDAP.UsernameAttr),
	}

	err := utils.Validate.Struct(params)
	if err != nil {
		// for _, err := range err.(validator.ValidationErrors) {

		// 	fmt.Println(err.Field())
		// 	fmt.Println(err.Tag())
		// 	fmt.Println(err.Kind())
		// }
		return nil, false
	}

	c := context.Background()

	user, err := utils.Database.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Username:  params.Username,
		Email:     params.Email,
	})
	if strings.HasPrefix(err.Error(), "UNIQUE") {
		return nil, false
	}
	if err != nil {
		log.Println("Failed to add user to database")
	}
	ur := DatabaseUserToUser(user)
	return &ur, true
}
