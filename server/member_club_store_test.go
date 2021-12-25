package main

import (
	"reflect"
	"testing"
	"time"
)

func TestAddMemberStoresNewMember(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMember := makeTestdMembers()[0]

	mcs.AddMember(expectedMember)

	if mcs.members[0] != expectedMember {
		t.Errorf("Expected %v, got %v instead\n", expectedMember, mcs.members[0])
	}
}

func TestAddMemberStoresTwoMembersOnAddingTwoMembers(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestdMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	if !reflect.DeepEqual(mcs.members, expectedMembers) {
		t.Errorf("Expected %v, got %v instead\n", expectedMembers, mcs.members)
	}
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestdMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])
	mcs.AddMember(expectedMembers[0])

	if !reflect.DeepEqual(mcs.members, expectedMembers) {
		t.Errorf("Expected %v, got %v instead\n", expectedMembers, mcs.members)
	}
}

func TestGetMembersReturnsAllMembersAdded(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestdMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	fetchedMembers := mcs.GetMembers()
	if !reflect.DeepEqual(fetchedMembers, expectedMembers) {
		t.Errorf("Expected %v, got %v instead\n", fetchedMembers, expectedMembers)
	}
}

// Helpers

func makeTestdMembers() []Member {
	return []Member{
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
}
