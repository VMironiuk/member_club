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

	expectEqual(mcs.members, []Member{expectedMember}, t)
}

func TestAddMemberStoresTwoMembersOnAddingTwoMembers(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestdMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	expectEqual(mcs.members, expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestdMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])
	mcs.AddMember(expectedMembers[0])

	expectEqual(mcs.members, expectedMembers, t)
}

func TestGetMembersReturnsAllMembersAdded(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestdMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	fetchedMembers := mcs.GetMembers()
	expectEqual(fetchedMembers, expectedMembers, t)
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

func expectEqual(givenMembers []Member, expectedMembers []Member, t *testing.T) {
	if !reflect.DeepEqual(givenMembers, expectedMembers) {
		t.Errorf("Expected %v, got %v instead\n", expectedMembers, givenMembers)
	}
}