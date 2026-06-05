package user_status

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/userdata"

	"github.com/stretchr/testify/assert"
)

// TestAddAccessoryLimit - Should add 100 Money
func TestAddAccessoryLimit(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AccessoryBoxLimit: 0,
		},
	}
	AddUserAccessoryLimit(&session, 100)
	assert.Equal(t, int32(100), session.UserStatus.AccessoryBoxLimit)
}

// TestAddAccessoryLimitFull - Should add 5 Subscription Coins, taking us to max
func TestAddAccessoryLimitFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AccessoryBoxLimit: (1 << 31) - 5,
		},
	}
	AddUserAccessoryLimit(&session, 100)
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.AccessoryBoxLimit)
}

// TestAddAccessoryLimitNegative - Should remove 10 AccessoryLimit
func TestAddAccessoryLimitNegative(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AccessoryBoxLimit: 123,
		},
	}

	AddUserAccessoryLimit(&session, -10)
	assert.Equal(t, int32(113), session.UserStatus.AccessoryBoxLimit)
}

// TestAddAccessoryLimitNegativeBelowZero - Should remove 5 AccessoryLimit, taking us to 0
func TestAddAccessoryLimitNegativeBelowZero(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			AccessoryBoxLimit: 5,
		},
	}
	AddUserAccessoryLimit(&session, -10)
	assert.Equal(t, int32(0), session.UserStatus.AccessoryBoxLimit)
}
