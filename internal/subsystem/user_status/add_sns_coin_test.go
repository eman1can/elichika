package user_status

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/item"
	"elichika/internal/userdata"

	"github.com/stretchr/testify/assert"
)

// TestAddFreeSnsCoin - Should add 100 SNS Coins
func TestAddFreeSnsCoin(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			FreeSnsCoin: 0,
		},
	}
	addSnsCoin(&session, new(item.StarGem.Amount(100)))
	assert.Equal(t, int32(100), session.UserStatus.FreeSnsCoin)
}

// TestAddAppleSnsCoin - Should add 100 SNS Coins
func TestAddAppleSnsCoin(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AppleSnsCoin: 0,
		},
	}
	addSnsCoin(&session, new(item.StarGemApple.Amount(100)))
	assert.Equal(t, int32(100), session.UserStatus.AppleSnsCoin)
}

// TestAddGoogleSnsCoin - Should add 100 SNS Coins
func TestAddGoogleSnsCoin(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GoogleSnsCoin: 0,
		},
	}
	addSnsCoin(&session, new(item.StarGemGoogle.Amount(100)))
	assert.Equal(t, int32(100), session.UserStatus.GoogleSnsCoin)
}

// TestAddFreeSnsCoinFull - Should add 5 SNS Coins, taking us to max
func TestAddFreeSnsCoinFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			FreeSnsCoin: (1 << 31) - 5,
		},
	}
	addSnsCoin(&session, new(item.StarGem.Amount(100)))
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.FreeSnsCoin)
}

// TestAddAppleSnsCoinFull - Should add 5 SNS Coins, taking us to max
func TestAddAppleSnsCoinFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AppleSnsCoin: (1 << 31) - 5,
		},
	}
	addSnsCoin(&session, new(item.StarGemApple.Amount(100)))
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.AppleSnsCoin)
}

// TestAddGoogleSnsCoinFull - Should add 5 SNS Coins, taking us to max
func TestAddGoogleSnsCoinFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GoogleSnsCoin: (1 << 31) - 5,
		},
	}
	addSnsCoin(&session, new(item.StarGemGoogle.Amount(100)))
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.GoogleSnsCoin)
}

// TestShouldOnlyAddSnsCoin - Should only allow adding if item is SNS Coins
func TestShouldOnlyAddSnsCoin(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			FreeSnsCoin: 123,
		},
	}
	addSnsCoin(&session, new(item.Gold.Amount(100)))
	assert.Equal(t, int32(123), session.UserStatus.FreeSnsCoin)
}
