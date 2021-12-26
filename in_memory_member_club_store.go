package main

import (
	"net/mail"
	"regexp"
)

type InMemoryMemberClubStore struct {
	members []Member
}

func (store *InMemoryMemberClubStore) AddMember(member Member) {
	if !isNameValid(member.name) {
		return
	}

	if !isValidEmail(member.email) {
		return
	}

	if store.containsEmail(member.email) {
		return
	}

	store.members = append(store.members, member)
}

func (store *InMemoryMemberClubStore) GetMembers() []Member {
	return store.members
}

// Helpers

func (store *InMemoryMemberClubStore) containsEmail(email string) bool {
	for _, m := range store.members {
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

func isNameValid(name string) bool {
	r := regexp.MustCompile(`^[a-zA-Z \.]+$`)
	return r.MatchString(name)
}
