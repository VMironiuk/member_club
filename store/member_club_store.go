package store

import (
	"time"
)

type InvalidMemberNameError struct{}

func (err *InvalidMemberNameError) Error() string {
	return "Invalid member name. Can contain only letters, spaces and dots."
}

type InvalidMemberEmailError struct{}

func (err *InvalidMemberEmailError) Error() string {
	return "Invalid member email."
}

type MemberWithSameEmailError struct{}

func (err *MemberWithSameEmailError) Error() string {
	return "Member with same email is already registered."
}

type Member struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	DateAdded time.Time `json:"date_added"`
}

type MemberClubStore interface {
	AddMember(mmember Member) error
	GetMembers() []Member
}
