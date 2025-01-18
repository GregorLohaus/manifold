package types

import (
	"errors"
	"time"
)

const TIME_LAYOUT = "2006-01-02T15:04:05-0700"

type Session struct {
	Id          *string    `json:"id,omitempty"`
	User        *string    `json:"user,omitempty"`
	Token       *string    `json:"session_token,omitempty"`
	ExpiresAt   *time.Time `json:"-"`
	ExpiresAtDB *string    `json:"expires_at,omitempty" validate:"omitnil,datetime=2006-01-02T15:04:05Z0700"`
}

func (s *Session) Query(session *string) (query string, args map[string]interface{}) {
	query = `select *, user.* from $user_session;`
	if session != nil {
		args = map[string]interface{}{"user_session": session}
		return
	}
	args = map[string]interface{}{"user_session": s.Id}
	return
}

func (s *Session) UpdateToken(token string, session *string) (query string, args map[string]interface{}) {
	query = `update $user_session set session_token = $user_session_token;`
	if session != nil {
		s.Token = &token
		args = map[string]interface{}{"user_session": session, "user_session_token": token}
		return
	}
	args = map[string]interface{}{"user_session": s.Id, "user_session_token": token}
	return
}

func (s *Session) UpdateExpiresAt(time time.Time, session *string) (query string, args map[string]interface{}) {
	query = `update $user_session set expires_at = $time;`
	if session != nil {
		s.ExpiresAt = &time
		args = map[string]interface{}{"user_session": session, "time": time.Format(TIME_LAYOUT)}
		return
	}
	args = map[string]interface{}{"user_session": s.Id, "time": time.Format(TIME_LAYOUT)}
	return
}

func (s *Session) QueryByUser(user string) (query string, args map[string]interface{}) {
	query = `select * from session where user == $user;`
	args = map[string]interface{}{"user": user}
	return
}

func (s *Session) QueryByToken(session_token string) (query string, args map[string]interface{}) {
	query = `select * from session where session_token == $session_token;`
	args = map[string]interface{}{"session_token": session_token}
	return
}

func (s *Session) CreateForUser() (query string, args map[string]interface{}) {
	query = `create session set user=$user, session_token=$user_session_token, expires_at=$expires_at;`
	args = map[string]interface{}{"user": s.User, "user_session_token": s.Token, "expires_at": s.ExpiresAt.Format(TIME_LAYOUT)}
	return
}

func (s *Session) DBExpieryToExpiery() error {
	if s.ExpiresAtDB == nil {
		return errors.New("EpiresAtDB not set")
	}
	t, err := time.Parse(TIME_LAYOUT, *s.ExpiresAtDB)
	if err != nil {
		return err
	}
	s.ExpiresAt = &t
	return nil
}
