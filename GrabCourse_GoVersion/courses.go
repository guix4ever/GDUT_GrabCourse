// courses.go

package main

const (
	sessionID = "JSESSIONID=1BCC91CA39E9FBFA7CB3B921E8AF0F32" // 请替换为你的Cookie
)

var Courses = []struct {
	courseName string
	courseCode string
}{
	{
		courseName: "体育(4)",
		courseCode: "1274385",
	},
	// 添加更多课程信息(kcmc不变，主要是找kcrwdm)
}
