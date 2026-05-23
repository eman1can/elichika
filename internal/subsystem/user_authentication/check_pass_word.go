package user_authentication

import (
	"elichika/internal/userdata"
	"elichika/internal/userdata/database"
	"elichika/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassWord(session *userdata.Session, passWord string) bool {
	userPassWord := database.UserPassWord{}
	exist, err := session.Db.Table("u_pass_word").Where("user_id = ?", session.UserId).Get(&userPassWord)
	utils.CheckErr(err)
	if !exist {
		return true
	}
	err = bcrypt.CompareHashAndPassword(userPassWord.Hash, []byte(passWord))
	return err == nil
}
