package main

import (
	"net/mail"
	"time"
)

type Member struct {
	name      string
	email     string
	dateAdded time.Time
}

type MemberClubStore struct {
	members []Member
}

func (mcs *MemberClubStore) AddMember(member Member) {
	if !isValidEmail(member.email) {
		return
	}

	if mcs.containsEmail(member.email) {
		return
	}

	mcs.members = append(mcs.members, member)
}

func (mcs *MemberClubStore) GetMembers() []Member {
	return mcs.members
}

func (mcs *MemberClubStore) containsEmail(email string) bool {
	for _, m := range mcs.members {
		if m.email == email {
			return true
		}
	}
	return false
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
