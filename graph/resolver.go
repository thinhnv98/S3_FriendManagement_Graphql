package graph

import (
	"S3_FriendManagement_Graphql/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	IUserService         services.IUserService
	IFriendService       services.IFriendService
	ISubscriptionService services.ISubscriptionService
	IBlockingService     services.IBlockingService
}
