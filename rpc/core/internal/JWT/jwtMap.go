package JWT

import (
	"context"
	"douyin/rpc/core/internal/svc"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Username string
	Password string
}

func JWTAuth(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		fmt.Println("进入 jwt JWTAuth")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Unexpected signing method")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("douyin"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["Username"], claims["Password"])
		return claims, nil
	} else {
		fmt.Println("获取claims出错了")
		return nil, err
	}
}

func JWTAuthToken(svcCtx *svc.ServiceContext, ctx context.Context, tokenString string) error {
	claims, err := JWTAuth(tokenString)
	if err != nil {
		return err
	}
	username := claims["Username"].(string)
	password := claims["Password"].(string)
	_, err = svcCtx.UserModel.FindOneByToken(ctx, username, password)
	if err != nil {
		fmt.Println("查找token里的用户出错了")
		return err
	}
	return nil
}
