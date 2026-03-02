package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kowon/naver-shopping-search-skills/internal/client"
	"github.com/kowon/naver-shopping-search-skills/internal/shopping"
)

const display = 5

func usage() {
	fmt.Fprintf(os.Stderr, `사용법: shop <command> <query>

Commands:
  check   상품 가격을 확인합니다 (정확도순)
  compare 가격을 비교합니다 (가격 오름차순)
  link    구매 링크를 제공합니다

예시:
  shop check 아이폰16
  shop compare 맥북프로
  shop link 에어팟프로
`)
	os.Exit(1)
}

func main() {
	// .env 파일 로드 (OS 환경변수가 우선)
	_ = godotenv.Load()

	if len(os.Args) < 3 {
		usage()
	}

	command := os.Args[1]
	query := os.Args[2]

	c, err := client.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "오류:", err)
		os.Exit(1)
	}

	var sortParam string
	switch command {
	case "check":
		sortParam = "sim"
	case "compare":
		sortParam = "asc"
	case "link":
		sortParam = "sim"
	default:
		fmt.Fprintf(os.Stderr, "알 수 없는 명령어: %s\n\n", command)
		usage()
	}

	resp, err := c.Search(query, sortParam, display)
	if err != nil {
		fmt.Fprintln(os.Stderr, "검색 실패:", err)
		os.Exit(1)
	}

	switch command {
	case "check":
		shopping.PrintCheck(resp.Items)
	case "compare":
		shopping.PrintCompare(resp.Items)
	case "link":
		shopping.PrintLink(resp.Items)
	}
}
