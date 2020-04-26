package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("指定された引数の数が間違っています。")
		os.Exit(1)
	}
	message := os.Args[1]
	result := formatGithubMessage(message)
	if result == "" {
		fmt.Println("セキュリティに関する問題はありません。")
		os.Exit(1)
	}
	fmt.Print(formatGithubMessage(message))
}

func formatGithubMessage(message string) string {
	const title = "# PHP Security Check Report\\n=============================\\n"
	body, err := selectMessageBody(message)
	if err != nil {
		return ""
	}
	return title + body
}

func selectMessageBody(message string) (string, error) {
	start := strings.Index(message, "[CVE")
	if start == -1 {
		return "", errors.New("エラー情報がありません。")
	}
	result := strings.Replace(message[start:], "[CVE", "\\n- [ ] [CVE", -1)
	return result, nil
}
