package main

import "time"

type Member struct {
	name      string
	email     string
	dateAdded time.Time
}

type MemberClubStore struct {
	members []Member
}

func (mcs *MemberClubStore) AddMember(member Member) {
	mcs.members = append(mcs.members, member)
}
