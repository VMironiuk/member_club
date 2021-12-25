package main

import (
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
