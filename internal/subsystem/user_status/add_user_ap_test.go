package user_status

import (
	"testing"
	"time"

	"elichika/internal/client"
	gamedata2 "elichika/internal/gamedata"
	"elichika/internal/userdata"
	"elichika/internal/utils"

	"github.com/stretchr/testify/assert"
)

// TestAddUserActivityPoints - Should add 2 AP with no issue
func TestAddUserActivityPoints(t *testing.T) {
	gamedata := gamedata2.Gamedata{
		ConstantInt: ,
	}
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			ActivityPointCount:   1,
			ActivityPointResetAt: utils.BeginOfNextHalfDay(time.Now()).Unix(),
		},
	}
	AddUserActivityPoints(&session, 2)
	assert.Equal(t, int32(2), session.UserStatus.ActivityPointCount)
}

// TestAddActivityPointsFull
// TestAddActivityPointsAfterReset
