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
	if mcs.contains(member) {
		return
	}
	mcs.members = append(mcs.members, member)
}

func (mcs *MemberClubStore) contains(member Member) bool {
	for _, m := range mcs.members {
		if m == member {
			return true
		}
	}
	return false
}
