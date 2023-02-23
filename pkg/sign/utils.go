package sign

import "math/rand"

// 个性签名
var signatures = []string{
	"你未必光芒万丈，但始终温暖有光。",
	"生活就是在寻中找，在找寻；",
	"青春看起来如此完美，没空闲去浪费时间。",
	"做个温暖的人，心底有暖意 脸上有笑容，眼里有欢喜",
	"你若三心二意，我会逢场作戏。",
	"修的是心，行的是活，心若自在，人才乐活。",
	"放轻松，就当漫游地球。",
	"生活大概就是生下来活下去。"}

// RandomSignature 生成随机的个性简介
func RandomSignature() string {
	return signatures[rand.Intn(len(signatures))]
}
