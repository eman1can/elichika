package billing

import (
	"elichika/internal/client/response"
	"elichika/internal/handler/common"
	"elichika/internal/server"
	"elichika/internal/subsystem/user_subscription_status"
	"elichika/internal/userdata"

	"github.com/gin-gonic/gin"
)

func updateSubscription(ctx *gin.Context) {
	// there's no request body

	session := ctx.MustGet("session").(*userdata.Session)

	// TODO(subscription): Implement subscription logic better
	subscriptionStatus := user_subscription_status.GetUserSubsriptionStatus(session, 13001)

	subscriptionStatus.ExpireDate = 1<<31 - 1 // preserve the subscription for now
	subscriptionStatus.PlatformExpireDate = subscriptionStatus.ExpireDate
	session.UserModel.UserSubscriptionStatusById.Set(subscriptionStatus.SubscriptionMasterId, subscriptionStatus)

	common.JsonResponse(ctx, &response.UpdateSubscriptionResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	server.AddHandler("/", "POST", "/billing/updateSubscription", updateSubscription)
}
