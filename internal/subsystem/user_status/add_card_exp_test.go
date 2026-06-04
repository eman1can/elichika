package user_status

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/item"
	"elichika/internal/userdata"

	"github.com/stretchr/testify/assert"
)

func TestAddExp(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			CardExp: 0,
		},
	}
	addCardExp(&session, new(item.EXP.Amount(100)))
	assert.Equal(t, int32(100), session.UserStatus.CardExp)
}
