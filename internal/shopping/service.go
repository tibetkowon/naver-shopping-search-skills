package shopping

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var htmlTagRe = regexp.MustCompile(`<[^>]+>`)

// stripTags는 HTML 태그를 제거합니다.
func stripTags(s string) string {
	return htmlTagRe.ReplaceAllString(s, "")
}

// formatPrice는 숫자 문자열을 ₩1,234,567 형식으로 변환합니다.
func formatPrice(s string) string {
	if s == "" || s == "0" {
		return "가격 정보 없음"
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return s
	}

	// 천 단위 콤마 삽입
	formatted := fmt.Sprintf("%d", n)
	var result []byte
	for i, c := range formatted {
		if i > 0 && (len(formatted)-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte(c))
	}

	return "₩" + string(result)
}

// PrintCheck는 가격 확인 결과를 출력합니다 (정확도순).
func PrintCheck(items []Product) {
	if len(items) == 0 {
		fmt.Println("검색 결과가 없습니다.")
		return
	}
	fmt.Println("=== 가격 확인 결과 ===")
	for i, item := range items {
		title := stripTags(item.Title)
		price := formatPrice(item.Lprice)
		fmt.Printf("%d. %s | %s | %s\n", i+1, title, price, item.Link)
	}
}

// PrintCompare는 가격 비교 결과를 출력합니다 (가격 오름차순).
func PrintCompare(items []Product) {
	if len(items) == 0 {
		fmt.Println("검색 결과가 없습니다.")
		return
	}
	fmt.Println("=== 가격 비교 결과 ===")
	maxTitleLen := 0
	for _, item := range items {
		l := len([]rune(stripTags(item.Title)))
		if l > maxTitleLen {
			maxTitleLen = l
		}
	}
	for i, item := range items {
		title := stripTags(item.Title)
		price := formatPrice(item.Lprice)
		mall := item.MallName
		if mall == "" {
			mall = "-"
		}
		padding := strings.Repeat(" ", maxTitleLen-len([]rune(title)))
		fmt.Printf("%d. %s%s | %s | %s\n", i+1, title, padding, price, mall)
	}
}

// PrintLink는 구매 링크를 출력합니다.
func PrintLink(items []Product) {
	if len(items) == 0 {
		fmt.Println("검색 결과가 없습니다.")
		return
	}
	fmt.Println("=== 구매 링크 ===")
	for i, item := range items {
		title := stripTags(item.Title)
		fmt.Printf("%d. %s | %s\n", i+1, title, item.Link)
	}
}
