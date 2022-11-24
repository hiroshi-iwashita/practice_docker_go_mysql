package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package.
// It returns the type Model, which embeds all of the types we want to
// be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User:  User{},
		Token: Token{},
	}
}

// Models is the type for this package. Note that any model that is
// included as a member in this type is available to us throughout the
// application, anywhere that the app variable is used, provided that the
// model is also added in the New function.
type Models struct {
	User  User
	Token Token
}

// User is the stucture which holds one user from the database. Note
// that it embeds a token type.
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
}

// Token is the data structure for any token in the database. Note that
// we do not send the TokenHash (a slice of bytes) in any exported JSON.
type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}
