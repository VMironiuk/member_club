package main

import (
	"reflect"
	"testing"
	"time"
)

func TestAddMemberStoresNewMember(t *testing.T) {
	store := InMemoryMemberClubStore{}
	expectedMember := makeTestedMembers()[0]

	store.AddMember(expectedMember)

	expectEqual(store.GetMembers(), []Member{expectedMember}, t)
}

func TestAddMemberStoresAllAddedMembersOnAddingValidMembers(t *testing.T) {
	store := InMemoryMemberClubStore{}
	expectedMembers := makeTestedMembers()

	addMembers(&store, expectedMembers)

	expectEqual(store.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	store := InMemoryMemberClubStore{}
	expectedMembers := makeTestedMembers()

	addMembers(&store, expectedMembers)
	store.AddMember(expectedMembers[0])

	expectEqual(store.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidEmail(t *testing.T) {
	store := InMemoryMemberClubStore{}
	invalidMembers := makeTestedMembersWithInvalidEmail()

	addMembers(&store, invalidMembers)

	expectEqual(store.GetMembers(), []Member{}, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidName(t *testing.T) {
	store := InMemoryMemberClubStore{}
	invalidMembers := makeTestedMembersWithInvalidName()

	addMembers(&store, invalidMembers)

	expectEqual(store.GetMembers(), []Member{}, t)
}

func TestGetMembersReturnsAllMembersAdded(t *testing.T) {
	store := InMemoryMemberClubStore{}
	expectedMembers := makeTestedMembers()

	addMembers(&store, expectedMembers)

	fetchedMembers := store.GetMembers()
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

func addMembers(store *InMemoryMemberClubStore, members []Member) {
	for _, m := range members {
		store.AddMember(m)
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
