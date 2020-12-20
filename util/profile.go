package util

import (
	"encoding/json"
	"errors"
	"expert-back/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goslib/utils/utlHttp"
	"net/http"
)

const KeyAuthCookie = "iadmin"

func GinGetAccountProfile(ctx *gin.Context) (*model.AccountProfile, error) {
	cookie, err := ctx.Cookie(KeyAuthCookie)
	if err != nil {
		return nil, err
	}
	return GetAccountProfile(cookie)
}

func GetAccountProfile(token string) (*model.AccountProfile, error) {
	req, err := http.NewRequest(http.MethodGet, "https://asc.shusim.com/iadmin/user/info", nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: KeyAuthCookie, Value: token})
	err, code, res := utlHttp.DoHttpCall(req)
	if err != nil {
		fmt.Println("err", err, code, string(res))
		return nil, err
	}

	if code != http.StatusOK {
		return nil, errors.New(string(res))
	}

	profile := new(model.ResAccountProfile)
	err = json.Unmarshal(res, profile)
	if err != nil {
		return nil, err
	}
	if profile.Data == nil {
		return nil, errors.New(profile.GetErrMsg())
	}
	return profile.Data, nil
}
