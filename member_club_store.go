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
	if mcs.containsEmail(member) {
		return
	}
	mcs.members = append(mcs.members, member)
}

func (mcs *MemberClubStore) GetMembers() []Member {
	return mcs.members
}

func (mcs *MemberClubStore) containsEmail(member Member) bool {
	for _, m := range mcs.members {
		if m.email == member.email {
			return true
		}
	}
	return false
}
