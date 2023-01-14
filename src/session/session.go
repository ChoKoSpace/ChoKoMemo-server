package session

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ChoKoSpace/ChoKoMemo-server/src/config"
	"github.com/ChoKoSpace/ChoKoMemo-server/src/model"
)

func GenetareToken() string {
	randValue := make([]byte, 16)
	_, err := rand.Read(randValue)
	if err != nil {
		panic("Failed to generate Toekn")
	}

	sha := sha256.New()
	sha.Write(randValue)
	hash := sha.Sum(nil)

	return hex.EncodeToString(hash)
}

func IsValidToken(userId string, token string) bool {
	db := model.GetDB()
	session := model.Session{}
	if err := db.Where("user_id = ?", userId).First(&session).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return false
		}
	}

	if session.Token == token {
		if time.Now().Before(session.ExpiresAt) {
			return true
		}
	}
	return false
}

/* Return token */
func CreateSession(userId string) (string, error) {
	_userId, err := strconv.Atoi(userId)
	if err != nil {
		return "", fmt.Errorf("[Session]-> userId is not integer (%v)", userId)
	}

	newToken := GenetareToken()
	newExpiresAt := time.Now().Add(time.Second * config.TOKEN_LIFETIME)

	var session model.Session
	db := model.GetDB()
	if err := db.Where("user_id = ?", _userId).First(&session).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			//신구 Session Row 추가
			newSession := model.Session{}
			newSession.UserID = _userId
			newSession.Token = newToken
			newSession.ExpiresAt = newExpiresAt

			if err := db.Create(&newSession).Error; err != nil {
				return "", fmt.Errorf("Failed to create session [New](%v)", err)
			}
		}
	} else {
		//이미 user_Id로 만들어진 세션이 있다. 이 경우 새 토큰으로 row갱신
		updates := map[string]interface{}{
			"token":      newToken,
			"expires_at": newExpiresAt,
		}
		if err := db.Model(&model.Session{}).Where("user_id = ?", _userId).Updates(updates).Error; err != nil {
			return "", fmt.Errorf("Failed to create session [Update](%v)", err)
		}
	}
	return newToken, nil
}

func RefreshSession(userId string) error {
	//lifetime 연장
	db := model.GetDB()

	newExpiresAt := time.Now().Add(time.Second * config.TOKEN_LIFETIME)

	if err := db.Model(&model.Session{}).Where("user_id = ?", userId).Update("expires_at", newExpiresAt).Error; err != nil {
		return fmt.Errorf("Refresh Session Error (%v)", err)
	}
	return nil
}
