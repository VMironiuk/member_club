package main

import (
	"reflect"
	"testing"
	"time"
)

func TestAddMemberStoresNewMember(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMember := Member{
		name:      "member",
		email:     "member@example.com",
		dateAdded: time.Now(),
	}

	mcs.AddMember(expectedMember)

	if mcs.members[0] != expectedMember {
		t.Errorf("Expected %v, got %v instead\n", expectedMember, mcs.members[0])
	}
}

func TestAddMemberStoresTwoMembersOnAddingTwoMembers(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := []Member{
		Member{
			name:      "member1",
			email:     "member1@example.com",
			dateAdded: time.Now(),
		},
		Member{
			name:      "member2",
			email:     "member2@example.com",
			dateAdded: time.Now(),
		},
	}

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	if !reflect.DeepEqual(mcs.members, expectedMembers) {
		t.Errorf("Expected %v, got %v instead\n", expectedMembers, mcs.members)
	}
}
