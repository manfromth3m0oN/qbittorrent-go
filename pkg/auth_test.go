package pkg

import "testing"

func TestAuth(t *testing.T) {
	client := NewClient("http://localhost:8080")

	err := client.Login("admin", "adminadmin")
	if err != nil {
		t.Fatal("Failed to login")
	}

	if client.token == "" {
		t.Errorf("No token returned")
	}

	t.Logf("Recived token %s", client.token)

	err = client.Logout()
	if err != nil {
		t.Fatal("Failed to logout")
	}
}
