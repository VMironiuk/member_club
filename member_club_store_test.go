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

	addMembers(&mcs, expectedMembers)

	expectEqual(mcs.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestedMembers()

	addMembers(&mcs, expectedMembers)
	mcs.AddMember(expectedMembers[0])

	expectEqual(mcs.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidEmail(t *testing.T) {
	mcs := MemberClubStore{}
	invalidMembers := makeTestedMembersWithInvalidEmail()

	addMembers(&mcs, invalidMembers)

	expectEqual(mcs.GetMembers(), []Member{}, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidName(t *testing.T) {
	mcs := MemberClubStore{}
	invalidMembers := makeTestedMembersWithInvalidName()

	addMembers(&mcs, invalidMembers)

	expectEqual(mcs.GetMembers(), []Member{}, t)
}

func TestGetMembersReturnsAllMembersAdded(t *testing.T) {
	mcs := MemberClubStore{}
	expectedMembers := makeTestedMembers()

	addMembers(&mcs, expectedMembers)

	fetchedMembers := mcs.GetMembers()
	expectEqual(fetchedMembers, expectedMembers, t)
}

// Helpers

func makeTestedMembers() []Member {
	return []Member{
		{
			name:      "John Smith",
			email:     "member1@example.com",
			dateAdded: time.Now(),
		},
		{
			name:      "John H. Doe",
			email:     "member2@example.com",
			dateAdded: time.Now(),
		},
	}
}

func makeTestedMembersWithInvalidEmail() []Member {
	return []Member{
		{
			name:      "Mr. Brown",
			email:     "bad-email",
			dateAdded: time.Now(),
		},
		{
			name:      "Dan Ben",
			email:     "@mail",
			dateAdded: time.Now(),
		},
	}
}

func makeTestedMembersWithInvalidName() []Member {
	return []Member{
		{
			name:      "$member1",
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

func addMembers(mcs *MemberClubStore, members []Member) {
	for _, m := range members {
		mcs.AddMember(m)
	}
}

func expectEqual(givenMembers []Member, expectedMembers []Member, t *testing.T) {
	if len(givenMembers) == 0 && len(expectedMembers) == 0 {
		return
	}

	if !reflect.DeepEqual(givenMembers, expectedMembers) {
		t.Errorf("Expected %v, got %v instead\n", expectedMembers, givenMembers)
	}
}
