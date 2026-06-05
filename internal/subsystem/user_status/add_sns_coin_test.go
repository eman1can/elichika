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

// TestAddSnsCoinNegative - Should remove 10 Free SNS Coins
func TestAddSnsCoinNegative(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			FreeSnsCoin: 123,
		},
	}
	addSnsCoin(&session, new(item.StarGem.Amount(-10)))
	assert.Equal(t, int32(113), session.UserStatus.FreeSnsCoin)
}

// TestAddSnsCoinNegativeBelowZero - Should remove 5 SNS Coins, taking us to 0
func TestAddSnsCoinNegativeBelowZero(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			FreeSnsCoin: 5,
		},
	}
	addSnsCoin(&session, new(item.StarGem.Amount(-10)))
	assert.Equal(t, int32(0), session.UserStatus.FreeSnsCoin)
}

// TestAddSnsCoinAppleNegative - Should remove 10 Free SNS Coins
func TestAddSnsCoinAppleNegative(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AppleSnsCoin: 123,
		},
	}
	addSnsCoin(&session, new(item.StarGemApple.Amount(-10)))
	assert.Equal(t, int32(113), session.UserStatus.AppleSnsCoin)
}

// TestAddSnsCoinAppleNegativeBelowZero - Should remove 5 SNS Coins, taking us to 0
func TestAddSnsCoinAppleNegativeBelowZero(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AppleSnsCoin: 5,
		},
	}
	addSnsCoin(&session, new(item.StarGemApple.Amount(-10)))
	assert.Equal(t, int32(0), session.UserStatus.AppleSnsCoin)
}

// TestAddSnsCoinGoogleNegative - Should remove 10 Free SNS Coins
func TestAddSnsCoinGoogleNegative(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GoogleSnsCoin: 123,
		},
	}
	addSnsCoin(&session, new(item.StarGemGoogle.Amount(-10)))
	assert.Equal(t, int32(113), session.UserStatus.GoogleSnsCoin)
}

// TestAddSnsCoinGoogleNegativeBelowZero - Should remove 5 SNS Coins, taking us to 0
func TestAddSnsCoinGoogleNegativeBelowZero(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GoogleSnsCoin: 5,
		},
	}
	addSnsCoin(&session, new(item.StarGemGoogle.Amount(-10)))
	assert.Equal(t, int32(0), session.UserStatus.GoogleSnsCoin)
}
