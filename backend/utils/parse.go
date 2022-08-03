package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// title `# `,summary `> `,
func ParseMd(src string, reg string) (dist string, err error) {
	reg_join := fmt.Sprintf("(?m)^%s(?s:(.*?))\\n", reg)
	reg1 := regexp.MustCompile(reg_join)
	if reg1 == nil {
		err = errors.New("regexp err")
		return
	}
	//根据规则提取关键信息
	// result1 := reg1.FindAllStringSubmatch(str1, -1)
	result := reg1.FindAllString(src, 1)
	var str1 string
	if len(result) == 0 {
		err = errors.New("regexp no match")
		return
	}

	string1 := strings.Split(result[0], reg)

	for _, v := range string1 {
		if len(v) > 5 {
			str1 = v
			break
		}
	}

	dist = strings.TrimSpace(str1)
	if len([]rune(dist)) < 3 {
		err = errors.New("match is less than 3 unicode")
		dist = ""
		return
	}
	if len([]rune(dist)) > 60 {
		dist = string([]rune(dist)[:60])
	}
	return
}

func GetUid(srcId interface{}) uint {
	suid := fmt.Sprintf("%v", srcId)
	dduid, err := strconv.Atoi(suid)
	if err != nil {
		return 0
	}
	return uint(dduid)
}

// 生成随机验证码
func MakeEmailCode() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	randn := n.Int64()
	rn := randn + 100000
	res := strconv.FormatInt(rn, 10)
	return res
}
