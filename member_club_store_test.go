package main

import (
	"reflect"
	"testing"
	"time"
)

func TestAddMemberStoresNewMember(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMember := makeTestedMembers()[0]

	mcs.AddMember(expectedMember)

	expectEqual(mcs.GetMembers(), []Member{expectedMember}, t)
}

func TestAddMemberStoresAllAddedMembersOnAddingValidMembers(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestedMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	expectEqual(mcs.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestedMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])
	mcs.AddMember(expectedMembers[0])

	expectEqual(mcs.GetMembers(), expectedMembers, t)
}

func TestGetMembersReturnsAllMembersAdded(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestedMembers()

	mcs.AddMember(expectedMembers[0])
	mcs.AddMember(expectedMembers[1])

	fetchedMembers := mcs.GetMembers()
	expectEqual(fetchedMembers, expectedMembers, t)
}

// Helpers

func makeTestedMembers() []Member {
	return []Member{
		{
			name:      "member1",
			email:     "member1@example.com",
			dateAdded: time.Now(),
		},
		{
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
