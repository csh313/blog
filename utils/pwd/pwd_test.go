package pwd

import (
	"fmt"
	"server/core"
	"server/global"
	"testing"
)

// 测试文件
func TestCs(t *testing.T) {
	core.InitConf()

	jwtPayLoad := JwtPayLoad{
		"csh",
		"chen",
		1,
		2,
	}
	fmt.Println("---", global.Config.Jwt.Issuer, "-----")
	fmt.Println(jwtPayLoad)
	jwt, err := GenerateJWT(jwtPayLoad)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jwt)
	fmt.Println("===============")
	parseJWT, err := ParseJWT(jwt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parseJWT)
	fmt.Println("------")
	fmt.Println(parseJWT["role"])

}
