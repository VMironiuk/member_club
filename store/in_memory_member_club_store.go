package store

import (
	"net/mail"
	"regexp"
	"time"
)

type InMemoryMemberClubStore struct {
	members []Member
}

func (store *InMemoryMemberClubStore) AddMember(member Member) error {
	if !isNameValid(member.Name) {
		return &InvalidMemberNameError{}
	}

	if !isValidEmail(member.Email) {
		return &InvalidMemberEmailError{}
	}

	if store.containsEmail(member.Email) {
		return &MemberWithSameEmailError{}
	}

	member.DateAdded = time.Now()
	store.members = append(store.members, member)

	return nil
}

func (store *InMemoryMemberClubStore) GetMembers() []Member {
	return store.members
}

// Helpers

func (store *InMemoryMemberClubStore) containsEmail(email string) bool {
	for _, m := range store.members {
		if m.Email == email {
			return true
		}
	}
	return false
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isNameValid(name string) bool {
	r := regexp.MustCompile(`^[a-zA-Z \.]+$`)
	return r.MatchString(name)
}
