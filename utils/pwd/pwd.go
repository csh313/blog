package pwd

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"server/global"
	"time"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	Username string `json:"user_name"` // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
}

// HashPwd加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// CheckPwd验证密码 hashedPwd hash之后的密码  plainPwd输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func GenerateJWT(user JwtPayLoad) (string, error) {

	// 创建一个 JWT Token，指定使用 HS256 签名方法和声明内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    user.UserID,
		"role":      user.Role,
		"NickName":  user.NickName,
		"Username":  user.Username,
		"Issuer":    global.Config.Jwt.Issuer,
		"ExpiresAt": time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires)).Unix(), //
	})
	// 使用密钥对 token 进行签名
	SignedToken, err := token.SignedString([]byte(global.Config.Jwt.Secret))
	//"Bearer " 是标准的身份验证方式前缀，通常与 JWT 一起使用。
	//在 HTTP 请求的 Authorization 头部中，Bearer token 用来指示传递的是一个 token。
	return "Bearer " + SignedToken, err
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	//如果token包含了"Bearer "就先去除掉在解析验证JWT
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//类型断言，判断是否是正确的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("签名方法有误")
		}
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	//解析
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
