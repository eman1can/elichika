package enum

type CardRarityType int32

const (
	CardRarityTypeRare  int32 = 0x0000000a
	CardRarityTypeSRare int32 = 0x00000014
	CardRarityTypeURare int32 = 0x0000001e
)

var (
	// TODO: Refactor to use CardRarityType
	AllCardRarities []int32 = []int32{CardRarityTypeRare, CardRarityTypeSRare, CardRarityTypeURare}
)
