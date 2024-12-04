package jwtx

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/gin-gonic/gin"
)

// jwt生成
const (
	Code    = "code"
	Uid     = "uid"
	Role    = "role"
	IsSuper = "super"
	Kind    = "kind"
)

const (
	REFRESH = 2 //刷新用
	ACCESS  = 1 //业务用
)

type Token struct {
	SecretKey string //密钥
	Exp       int64  //过期时间
	UserId    int64  //用户id
	IsSuper   bool   //是否是超管
	LoginCode string //登陆code
	Kind      int    //类型
	Role      string //角色
}

func GenToken(t Token) (string, error) {
	claims := make(jwt.MapClaims) //数据仓声明
	now := time.Now().Unix()
	claims["iat"] = now
	claims["exp"] = t.Exp
	claims[Uid] = t.UserId
	claims[Code] = t.LoginCode
	claims[Role] = t.Role
	claims[IsSuper] = t.IsSuper
	claims[Kind] = t.Kind
	token := jwt.New(jwt.SigningMethodHS256) //token对象
	token.Claims = claims                    //token添加数据仓
	return token.SignedString([]byte(t.SecretKey))
}

// jwt获取
func GetUserid(ctx *gin.Context) (int64, error) {
	if _id, has := ctx.Get(Uid); has {
		if id, ok := _id.(float64); ok {
			return int64(id), nil
		}
	}
	return 0, errors.New("no uid from ctx")
}

// code 获取
func GetLoginCode(ctx *gin.Context) (string, error) {
	if _tmp, has := ctx.Get(Code); has {
		if __tmp, ok := _tmp.(string); ok {
			return __tmp, nil
		}
	}
	return "", errors.New("no code from ctx")
}

func GetRole(ctx *gin.Context) (string, error) {
	if _tmp, has := ctx.Get(Role); has {
		if __tmp, ok := _tmp.(string); ok {
			return __tmp, nil
		}
	}
	return "", errors.New("no code from ctx")
}

func GetIsSuper(ctx *gin.Context) (bool, error) {
	if _tmp, has := ctx.Get(IsSuper); has {
		if __tmp, ok := _tmp.(string); ok {
			return __tmp == "1", nil
		}
	}
	return false, errors.New("no code from ctx")
}

func GetKind(ctx *gin.Context) (int, error) {
	if _tmp, has := ctx.Get(Kind); has {
		if __tmp, ok := _tmp.(float64); ok {
			return int(__tmp), nil
		}
	}
	return 0, errors.New("no code from ctx")
}
