package main

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
	name      string
	email     string
	dateAdded time.Time
}

type MemberClubStore interface {
	AddMember(mmember Member) error
	GetMembers() []Member
}
