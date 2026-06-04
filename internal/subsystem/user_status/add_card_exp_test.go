package user_status

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/item"
	"elichika/internal/userdata"

	"github.com/stretchr/testify/assert"
)

// TestAddExp - Should add 100 EXP
func TestAddExp(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			CardExp: 0,
		},
	}
	addCardExp(&session, new(item.EXP.Amount(100)))
	assert.Equal(t, int32(100), session.UserStatus.CardExp)
}

// TestAddExpFull - Should add 5 EXP, taking us to max
func TestAddExpFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			CardExp: (1 << 31) - 5,
		},
	}
	addCardExp(&session, new(item.EXP.Amount(100)))
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.CardExp)
}

// TestShouldOnlyAddExp - Should only add EXP content
func TestShouldOnlyAddExp(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			CardExp: 123,
		},
	}
	addCardExp(&session, new(item.Gold.Amount(100)))
	assert.Equal(t, int32(123), session.UserStatus.CardExp)
}
