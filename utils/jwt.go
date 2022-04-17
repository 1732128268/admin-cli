package utils

import (
	"admin-cli/config"
	"github.com/dgrijalva/jwt-go"
)

type AuthClaims struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	RoleId string `json:"role_id"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(id uint, roleId, name string) (string, error) {
	cfg := config.GetConfig()
	//设置token有效时间
	claims := AuthClaims{
		Id:     id,
		Name:   name,
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: cfg.Jwt.ExpireTime,
			// 指定token发行人
			Issuer: "admin-cli",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString([]byte(cfg.Jwt.Key))
	return token, err
}

// ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*AuthClaims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	cfg := config.GetConfig()
	tokenClaims, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Jwt.Key), nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*AuthClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
