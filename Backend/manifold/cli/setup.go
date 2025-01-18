package cli

import (
	"fmt"
	"gitlab.com/manifold555112/manifold/lib"
)

func setup() error {
	db, err := lib.GetDb(nil)
	if err != nil {
		return err
	}
	result, err := db.Query(
		`BEGIN TRANSACTION;
		DEFINE NAMESPACE manifold;
		USE NS manifold;

		DEFINE DATABASE system;
		USE DB system;

		DEFINE TABLE user SCHEMALESS;
		DEFINE FIELD first_name ON user TYPE string;
		DEFINE FIELD last_name ON user TYPE string;
		DEFINE FIELD company ON user TYPE option<string>;
		DEFINE FIELD phone ON user TYPE option<string>;
		DEFINE FIELD email ON user TYPE string ASSERT string::is::email($value);
		DEFINE FIELD password ON user TYPE string VALUE IF $input {crypto::argon2::generate($input)} ELSE {$before};
		DEFINE FIELD roles ON user TYPE array<number>;
		DEFINE FIELD child_users ON user TYPE option<array<record<user>>>;
		DEFINE FIELD parent_user ON user TYPE option<record<user>>;
		DEFINE FIELD registration_key ON user TYPE option<string>;
		DEFINE FIELD verified ON user TYPE bool DEFAULT false;
		DEFINE FIELD plan ON user TYPE number DEFAULT 1;
		DEFINE FIELD plan_expiery ON user TYPE option<datetime>;

		DEFINE TABLE query_history SCHEMALESS;
		DEFINE FIELD query ON query_history TYPE string;
		DEFINE FIELD run_at ON query_history TYPE datetime;

		DEFINE TABLE user_session SCHEMALESS;
		DEFINE FIELD session_token ON user_session TYPE string;
		DEFINE FIELD user ON user_session TYPE record<user> ASSERT (SELECT count(id) == 1 as count FROM $value).count ;
		DEFINE FIELD expires_at ON user_session TYPE datetime;
		DEFINE INDEX unique_user ON TABLE user_session FIELDS user UNIQUE;
		COMMIT`,
		nil)
	if err != nil {
		return err
	}
	fmt.Printf("%v", result)
	return nil
}
