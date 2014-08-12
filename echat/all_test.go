package main

import "testing"

var (
	testuser User
)

func SetupTest() {
	userlist = make(map[int]User)
	testuser = User{id: 1, nick: "test"}
}
func TestSetupTest(t *testing.T) {
	SetupTest()
	if userlist == nil {
		t.Errorf("userlist is nil")
	}
}

func TestAddUserToList(t *testing.T) {
	AddUserToList(testuser)
	if len(userlist) != 1 {
		t.Errorf("Error, adding user to list did not increse length")
	}
}

func TestUserQuit(t *testing.T) {
	if len(userlist) != 1 {
		t.Errorf("Error, user was not in list to begin with")
	}
	testuser.Quit()
	if len(userlist) != 0 {
		t.Errorf("Error, user was apparently not removed from list")
	}
	if testuser.dead != true {
		t.Errorf("Error, user was not set to dead")
	}
}

func TestStrCat(t *testing.T) {
	if strcat("Hello", " World") != "Hello World" {
		t.Errorf("String concatenation test failed")
	}
}

func TestCheckNickCollision(t *testing.T) {
	SetupTest()
	AddUserToList(testuser)
	if CheckNickCollision("test") != true {
		t.Errorf("Nick collision test failed")
	}
}

func TestCheckNickCollisionCase(t *testing.T) {
	SetupTest()
	AddUserToList(testuser)
	if CheckNickCollision("TEST") != true {
		t.Errorf("Nick collision check is not case sensitive")
	}
}

func TestCheckNickCollision3(t *testing.T) {
	testuser.Quit()
	if CheckNickCollision("test") != false {
		t.Errorf("Nick collision triggered even though there is no user")
	}
}
