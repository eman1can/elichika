package serverdata

import (
	"encoding/json"

	"elichika/internal/config"
	utils2 "elichika/internal/utils"

	"xorm.io/xorm"
)

type NgWord struct {
	Word string `xorm:"pk"`
}

func InitializeNgWord(session *xorm.Session) {
	files := []string{config.ServerInitJsons + "wordlist_gl.json", config.ServerInitJsons + "wordlist_jp.json"}
	for _, file := range files {
		wordsJson := utils2.ReadAllText(file)
		words := []string{}
		err := json.Unmarshal([]byte(wordsJson), &words)
		utils2.CheckErr(err)
		for _, word := range words {
			ngWord := NgWord{
				Word: word,
			}
			exist, err := session.Table("s_ng_word").Exist(&ngWord)
			utils2.CheckErr(err)
			if !exist {
				_, err = session.Table("s_ng_word").Insert(&ngWord)
				utils2.CheckErr(err)
			}
		}
	}
}

func init() {
	addTable("s_ng_word", NgWord{}, InitializeNgWord)
}
