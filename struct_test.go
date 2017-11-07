package aknstruct

import (
	"log"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type User struct {
		Id       int    `map:"id"`
		Username string `map:"username"`
		Password string `map:"-"`
		Verified bool   `map:"verified"`
	}

	user := User{
		Id:       1411,
		Username: "akbarkn",
		Password: "itsasecret",
		Verified: true,
	}

	u := Map(user)
	log.Println("map user", u)
	log.Println("user id", u["id"])
	log.Println("username", u["username"])
	log.Println("password", u["password"])
	log.Println("verified", u["verified"])

	if u["id"] != user.Id {
		t.Errorf("user id is incorrect. should be %v", user.Id)
		t.Fail()
		return
	}
	if u["username"] != user.Username {
		t.Errorf("username is incorrect. should be %v", user.Id)
		t.Fail()
		return
	}
	if u["password"] != nil {
		t.Errorf("password should be nil")
		t.Fail()
		return
	}
	if u["verified"] != user.Verified {
		t.Errorf("verified is incorrect. should be %v", user.Verified)
		t.Fail()
		return
	}

	t.Logf("SUCCESS")
}
