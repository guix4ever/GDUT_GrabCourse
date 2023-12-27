// main.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// SelectCourse 封装课程选课功能
// SelectCourse 封装课程选课功能
func SelectCourse(sessionID string, courseName string, courseCode string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		startTime := time.Now() // 记录选课开始时间

		data := url.Values{
			"kcrwdm": {courseCode},
			"kcmc":   {courseName},
		}

		client := &http.Client{}

		// 创建 HTTP 请求
		req, err := http.NewRequest("POST", "https://jxfw.gdut.edu.cn/xsxklist!getAdd.action", strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Printf("Error creating request for %s: %v\n", courseName, err)
			return
		}

		// 设置请求头
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		//"Accept-Language": "en-US,en;q=0.5"
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Content-Type", "text/html;charset=UTF-8")
		req.Header.Set("Origin", "https://jxfw.gdut.edu.cn")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Referer", "https://jxfw.gdut.edu.cn/xskjcjxx!kjcjList.action")
		req.Header.Set("Cookie", sessionID) // 请确保 SESSIONID 是已定义的变量
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("DNT", "1")
		req.Header.Set("Sec-GPC", "1")
		req.Header.Set("Host", "jxfw.gdut.edu.cn")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Content-Length", "36")
		req.Header.Set("X-Requested-With", "XMLHttpRequest")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for %s: %v\n", courseName, err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response for %s: %v\n", courseName, err)
			return
		}

		// 输出选课结果和选课耗时
		fmt.Printf("Time: %s, Course: %s\n", startTime.Format(time.DateTime), courseName)

		if strings.Contains(string(body), "成功") {
			fmt.Printf("Successfully selected %s.\n", courseName)
			break
		} else {
			fmt.Printf("Failed to select %s. Retrying...\n", courseName)
			// 等待一段时间再次尝试选课，可以自行调整等待时间
			time.Sleep(time.Second * 500)
		}
	}
}

func main() { // 创建一个等待组，用于等待所有goroutines完成
	var wg sync.WaitGroup

	// 从courses.go文件导入课程信息
	courses := Courses

	// 启动并发选课(如果不止想选一门)
	for _, course := range courses {
		wg.Add(1)
		go SelectCourse(sessionID, course.courseName, course.courseCode, &wg)
	}

	// 等待所有goroutines完成
	wg.Wait()
}
