package user_status

import (
	"testing"

	"elichika/internal/client"
	"elichika/internal/enum"
	"elichika/internal/item"
	"elichika/internal/userdata"

	"github.com/stretchr/testify/assert"
)

// TestAddSubscriptionCoin - Should add 100 Subscription Coins
func TestAddSubscriptionCoin(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			SubscriptionCoin: 0,
		},
	}
	coin := client.Content{ContentType: enum.ContentTypeSubscriptionCoin, ContentAmount: 100}
	addSubscriptionCoin(&session, &coin)
	assert.Equal(t, int32(100), session.UserStatus.SubscriptionCoin)
}

// TestAddSubscriptionCoinFull - Should add 5 Subscription Coins, taking us to max
func TestAddSubscriptionCoinFull(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			SubscriptionCoin: (1 << 31) - 5,
		},
	}
	coin := client.Content{ContentType: enum.ContentTypeSubscriptionCoin, ContentAmount: 100}
	addSubscriptionCoin(&session, &coin)
	assert.Equal(t, int32((1<<31)-1), session.UserStatus.SubscriptionCoin)
}

// TestShouldOnlyAddSubscriptionCoin - Should only add Subscription Coins content
func TestShouldOnlyAddSubscriptionCoin(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			SubscriptionCoin: 123,
		},
	}
	addSubscriptionCoin(&session, new(item.Gold.Amount(100)))
	assert.Equal(t, int32(123), session.UserStatus.SubscriptionCoin)
}

// TestAddSubscriptionCoinNegative - Should remove 10 Subscription Coins
func TestAddSubscriptionCoinNegative(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			SubscriptionCoin: 123,
		},
	}
	coin := client.Content{ContentType: enum.ContentTypeSubscriptionCoin, ContentAmount: -10}
	addSubscriptionCoin(&session, &coin)
	assert.Equal(t, int32(113), session.UserStatus.SubscriptionCoin)
}

// TestAddSubscriptionCoinNegativeBelowZero - Should remove 5 Subscription Coins, taking us to 0
func TestAddSubscriptionCoinNegativeBelowZero(t *testing.T) {
	session := userdata.Session{
		UserStatus: &client.UserStatus{
			SubscriptionCoin: 5,
		},
	}
	coin := client.Content{ContentType: enum.ContentTypeSubscriptionCoin, ContentAmount: -10}
	addSubscriptionCoin(&session, &coin)
	assert.Equal(t, int32(0), session.UserStatus.SubscriptionCoin)
}
