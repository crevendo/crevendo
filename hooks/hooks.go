package hooks

import "github.com/crevendo/libraries/crevendo-go/crevendo/hook"

type CartHooks struct {
	hook.UnimplementedCartHooksServer
}

type CategoryHook struct {
	hook.UnimplementedCategoryHooksServer
}

type GatewayHooks struct {
	hook.UnimplementedGatewayHooksServer
}

type OrderHooks struct {
	hook.UnimplementedOrderHooksServer
}

type PaymentHooks struct {
	hook.UnimplementedPaymentHooksServer
}

type ProductHook struct {
	hook.UnimplementedProductHooksServer
}
