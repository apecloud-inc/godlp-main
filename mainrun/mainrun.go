package main

import (
	"bytes"
	"fmt"
	dlp "github.com/bytedance/godlp"
	"os"
)

func dupString(src string, coefficient int) string {
	var buffer bytes.Buffer
	for i := 0; i < coefficient; i++ {
		buffer.WriteString(src)
	}
	return buffer.String()
}

func dlpDemo(inStr string) {
	caller := "dlp"
	if eng, err := dlp.NewEngine(caller); err == nil {
		eng.ApplyConfigDefault()
		if results, err := eng.Detect(inStr); err == nil {
			eng.ShowResults(results)
		}
		if outStr, _, err := eng.Deidentify(inStr); err == nil {
			fmt.Printf(outStr)
			fmt.Println()
		}
	}
}

func main() {
	// 获取命令行参数
	args := os.Args[1:]
	if len(args) > 0 {
		// 如果有传入参数，将第一个参数作为输入字符串传递给 dlpDemo() 函数
		dlpDemo(args[0])
		// 在这里使用 outStr 和 resultsStr，可以将其打印出来或者进行其他操作
	} else {
		// 如果没有传入参数，报错
		fmt.Printf("Error: No Args")
		fmt.Println()

// 		defaultInput := `我的邮件是abcd@abcd.com,
// 18612341234是我的电话
// 护照123456789
// 密码1pple1atch
// 你家住在哪里啊? 我家住在北京市海淀区北三环西路43号,
// mac地址 06-06-06-aa-bb-cc
// 收件人：张真人  手机号码：18612341234
// 2022年8月8日，公司在github监控捕捉到有员工将LDAP公共账号密码GitHub开源项目GitHub - oneof777/zsh 。上传信息为srv_changyu 账号。
// srv_changyu 此账号是2018年changyu申请的专门用于访问artifactory平台使用。 由于历史原因，大量项目复用了此账号。
// 出于安全考虑和公司安全部门要求，现在通知大家修改项目中使用了此账号的地方。 我们申请了新的共用账号 用户名：hdmap_artifactory 密码：C1z9q7g0
// 我们索引到一些使用了srv_changyu的代码和文件地址，可以参考 srv_changyu 账号切换 hdmap_artifactory
// 希望大家于8月11日下班前完成更新，近期srv_changyu 账号会下线删除，因此原因导致build失败或其他影响的，可以联系地图云平台小助手获取相关咨询
// 以上 非常感谢大家的配合


// 公共账号密码GitHub
// 账号123密码apple1atch
// USER : hdmapadmin@momenta.ai
// 	PASS HDMAP@momenta
// 	os.environ["OPENAI_API_KEY"] = "sk-O4G7212z47eCB3RfXv2RT3B123FJfcmrKwt5xCbsZO1r1234"
// 	os.environ["OPENAI_API_KEY"] = "sk-O4G****z47eCB3RfXv2RT3B123FJfcmrKwt5xCbsZO1r1234"`
// 		dlpDemo(defaultInput)
	}
}

