package database

import (
	"os"
	"testing"
)

func setup() {
	dal.InitDB()
	dal.DB.Exec("DELETE FROM Users") // 清空db
}

func TestUser(t *testing.T) {
	setup()
	dal.CreateUser(User{"jack", 18})
	jack, _ := dal.GetUser("jack")
	if jack == nil || jack.Age != 18 {
		t.Fatal("Get user error")
	}
}

func MainTest(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
