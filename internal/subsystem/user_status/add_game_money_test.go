package user_status

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/item"
	"elichika/internal/userdata"

	"github.com/stretchr/testify/assert"
)

// TestAddMoney - Should add 100 Money
func TestAddMoney(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GameMoney: 0,
		},
	}
	addGameMoney(&session, new(item.Gold.Amount(100)))
	assert.Equal(t, int32(100), session.UserStatus.GameMoney)
}

// TestAddMoneyFull - Should add 5 Money, taking us to max
func TestAddMoneyFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GameMoney: (1 << 31) - 5,
		},
	}
	addGameMoney(&session, new(item.Gold.Amount(100)))
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.GameMoney)
}

// TestShouldOnlyAddMoney - Should only add EXP content
func TestShouldOnlyAddMoney(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GameMoney: 123,
		},
	}
	addCardExp(&session, new(item.EXP.Amount(100)))
	assert.Equal(t, int32(123), session.UserStatus.GameMoney)
}
