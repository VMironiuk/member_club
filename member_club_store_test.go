package main

import (
	"reflect"
	"testing"
	"time"
)

func TestAddMemberStoresNewMember(t *testing.T) {
	sut := makeSUT()
	expectedMember := makeTestedMembers()[0]

	sut.AddMember(expectedMember)

	expectEqual(sut.GetMembers(), []Member{expectedMember}, t)
}

func TestAddMemberStoresAllAddedMembersOnAddingValidMembers(t *testing.T) {
	sut := makeSUT()
	expectedMembers := makeTestedMembers()

	addMembers(&sut, expectedMembers)

	expectEqual(sut.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	sut := makeSUT()
	expectedMembers := makeTestedMembers()

	addMembers(&sut, expectedMembers)
	sut.AddMember(expectedMembers[0])

	expectEqual(sut.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidEmail(t *testing.T) {
	sut := makeSUT()
	invalidMembers := makeTestedMembersWithInvalidEmail()

	addMembers(&sut, invalidMembers)

	expectEqual(sut.GetMembers(), []Member{}, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidName(t *testing.T) {
	sut := makeSUT()
	invalidMembers := makeTestedMembersWithInvalidName()

	addMembers(&sut, invalidMembers)

	expectEqual(sut.GetMembers(), []Member{}, t)
}

func TestGetMembersReturnsAllValidMembersAdded(t *testing.T) {
	sut := makeSUT()
	expectedMembers := makeTestedMembers()

	addMembers(&sut, expectedMembers)

	fetchedMembers := sut.GetMembers()
	expectEqual(fetchedMembers, expectedMembers, t)
}

// Helpers

func makeSUT() MemberClubStore {
	return &InMemoryMemberClubStore{}
}

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

func addMembers(store *MemberClubStore, members []Member) {
	for _, m := range members {
		(*store).AddMember(m)
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
