package store

import (
	"testing"
)

func TestAddMemberStoresNewMember(t *testing.T) {
	sut := makeSUT()
	expectedMember := makeTestedMembers()[0]

	err := sut.AddMember(expectedMember)

	expectEqualErrors(err, nil, t)
	expectEqualMembers(sut.GetMembers(), []Member{expectedMember}, t)
}

func TestAddMemberStoresAllAddedMembersOnAddingValidMembers(t *testing.T) {
	sut := makeSUT()
	expectedMembers := makeTestedMembers()

	err := addMembers(&sut, expectedMembers)

	expectEqualErrors(err, nil, t)
	expectEqualMembers(sut.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithExistedEmail(t *testing.T) {
	sut := makeSUT()
	expectedMembers := makeTestedMembers()

	addMembers(&sut, expectedMembers)
	err := sut.AddMember(expectedMembers[0])

	expectEqualErrors(err, &MemberWithSameEmailError{}, t)
	expectEqualMembers(sut.GetMembers(), expectedMembers, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidEmail(t *testing.T) {
	sut := makeSUT()
	invalidMembers := makeTestedMembersWithInvalidEmail()

	err := addMembers(&sut, invalidMembers)

	expectEqualErrors(err, &InvalidMemberEmailError{}, t)
	expectEqualMembers(sut.GetMembers(), []Member{}, t)
}

func TestAddMemberDoesNotAddMemberWithInvalidName(t *testing.T) {
	sut := makeSUT()
	invalidMembers := makeTestedMembersWithInvalidName()

	err := addMembers(&sut, invalidMembers)

	expectEqualErrors(err, &InvalidMemberNameError{}, t)
	expectEqualMembers(sut.GetMembers(), []Member{}, t)
}

func TestGetMembersReturnsAllValidMembersAdded(t *testing.T) {
	sut := makeSUT()
	expectedMembers := makeTestedMembers()

	addMembers(&sut, expectedMembers)

	fetchedMembers := sut.GetMembers()
	expectEqualMembers(fetchedMembers, expectedMembers, t)
}

// Helpers

func makeSUT() MemberClubStore {
	return &InMemoryMemberClubStore{}
}

func makeTestedMembers() []Member {
	return []Member{
		{
			Name:  "John Smith",
			Email: "member1@example.com",
		},
		{
			Name:  "John H. Doe",
			Email: "member2@example.com",
		},
	}
}

func makeTestedMembersWithInvalidEmail() []Member {
	return []Member{
		{
			Name:  "Mr. Brown",
			Email: "bad-email",
		},
		{
			Name:  "Dan Ben",
			Email: "@mail",
		},
	}
}

func makeTestedMembersWithInvalidName() []Member {
	return []Member{
		{
			Name:  "$member1",
			Email: "member1@example.com",
		},
		{
			Name:  "member2",
			Email: "member2@example.com",
		},
	}
}

func addMembers(store *MemberClubStore, members []Member) error {
	for _, m := range members {
		err := (*store).AddMember(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func expectEqualMembers(givenMembers []Member, expectedMembers []Member, t *testing.T) {
	if len(givenMembers) == 0 && len(expectedMembers) == 0 {
		return
	}

	if len(givenMembers) != len(expectedMembers) {
		t.Errorf("Expected %v members, got %v instead\n", expectedMembers, givenMembers)
		return
	}

	for i := range givenMembers {
		if givenMembers[i].Name != expectedMembers[i].Name || givenMembers[i].Email != expectedMembers[i].Email {
			t.Errorf("Expected %v members, got %v instead\n", expectedMembers, givenMembers)
			return
		}
	}
}

func expectEqualErrors(givenError error, expectedError error, t *testing.T) {
	givenErrorText := errorText(givenError)
	expectedErrorText := errorText(expectedError)

	if givenErrorText != expectedErrorText {
		t.Errorf("Expected %v error, got %v instead\n", expectedErrorText, givenErrorText)
	}
}

func errorText(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
