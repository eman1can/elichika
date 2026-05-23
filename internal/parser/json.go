package parser

import (
	"encoding/json"

	utils2 "elichika/internal/utils"
)

func ParseJson(path string, result any) {
	text := utils2.ReadAllText(path)
	err := json.Unmarshal([]byte(text), result)
	utils2.CheckErr(err)
}
