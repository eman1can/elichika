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

// TestShouldOnlyAddMoney - Should only add Money content
func TestShouldOnlyAddMoney(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GameMoney: 123,
		},
	}
	addCardExp(&session, new(item.EXP.Amount(100)))
	assert.Equal(t, int32(123), session.UserStatus.GameMoney)
}

// TestAddMoneyNegative - Should remove 10 Money
func TestAddMoneyNegative(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GameMoney: 123,
		},
	}
	addGameMoney(&session, new(item.Gold.Amount(-10)))
	assert.Equal(t, int32(113), session.UserStatus.GameMoney)
}

// TestAddMoneyNegativeBelowZero - Should remove 5 Money, taking us to 0
func TestAddMoneyNegativeBelowZero(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			GameMoney: 5,
		},
	}
	addGameMoney(&session, new(item.Gold.Amount(-10)))
	assert.Equal(t, int32(0), session.UserStatus.GameMoney)
}
