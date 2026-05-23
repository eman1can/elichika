package parser

import (
	"encoding/json"

	"elichika/internal/utils"
)

func ParseJson(path string, result any) {
	text := utils.ReadAllText(path)
	err := json.Unmarshal([]byte(text), result)
	utils.CheckErr(err)
}
