package auth

import (
	"strings"

	"github.io/taserbeat/line-family-bot/modules/env"
	"github.io/taserbeat/line-family-bot/modules/utils"
)

type Authenticate struct {
	FamilyUserIds []string
}

type IAuthenticate interface {
	isFamily(userId string) bool
}

/* 認証クライアント */
var AuthenticateClient Authenticate

/* 認証クライアントを生成する */
func New(familyUserIdsString string) Authenticate {
	authenticate := Authenticate{
		FamilyUserIds: strings.Split(familyUserIdsString, ","),
	}

	return authenticate
}

/* 家族であるか判定する */
func (authenticate Authenticate) IsFamily(userId string) bool {
	return utils.Contains(authenticate.FamilyUserIds, userId)
}

func init() {
	AuthenticateClient = New(env.Env.FamilyUserIds)
}
