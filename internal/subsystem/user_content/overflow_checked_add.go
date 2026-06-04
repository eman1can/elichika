package user_content

func OverflowCheckedAdd(current *int32, added *int32) {
	res := int64(*current)
	res += int64(*added)
	if res >= (1 << 31) {
		*added = int32(res - (1<<31 - 1))
		*current = 1<<31 - 1
	} else if res > 0 {
		*current = int32(res)
		*added = 0
	} else {
		// TODO(resource): This should panic, as negative values aren't allowed. Each content costing piece of code should check that the user has enough resources before executing
		*current = 0
		*added = 0
	}
}
