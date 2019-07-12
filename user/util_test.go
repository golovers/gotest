package user_test

import (
	"os"
	"testing"

	"github.com/golovers/gotest/user"
)

func TestReadUserFromCSV(t *testing.T) {
	f, err := os.Open("testdata/user.csv")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	users := make(user.Users, 0)
	_, err = users.ReadFrom(f)
	if err != nil {
		t.Error(err)
	}
	expected := 2
	if len(users) != expected {
		t.Errorf("got len(users)=%d, wants len(users)=%d", len(users), expected)
		return
	}
	expectedFirstID := "123"
	if users[0].ID != expectedFirstID {
		t.Errorf("got users[0].ID=%s, wants users[0].ID=%s", users[0].ID, expectedFirstID)
	}
}
