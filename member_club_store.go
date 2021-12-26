package main

import (
	"time"
)

type Member struct {
	name      string
	email     string
	dateAdded time.Time
}

type MemberClubStore interface {
	AddMember(mmember Member)
	GetMembers() []Member
}
