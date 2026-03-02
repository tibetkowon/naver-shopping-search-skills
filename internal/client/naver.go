package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/kowon/naver-shopping-search-skills/internal/shopping"
)

const apiURL = "https://openapi.naver.com/v1/search/shop.json"

// NaverClient는 네이버 쇼핑 API 클라이언트입니다.
type NaverClient struct {
	clientID     string
	clientSecret string
	httpClient   *http.Client
}

// New는 환경변수에서 인증 정보를 로드하여 NaverClient를 생성합니다.
func New() (*NaverClient, error) {
	clientID := os.Getenv("NAVER_CLIENT_ID")
	clientSecret := os.Getenv("NAVER_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("환경변수가 설정되지 않았습니다: NAVER_CLIENT_ID, NAVER_CLIENT_SECRET\n.env.example을 참고하여 .env 파일을 생성하세요")
	}

	return &NaverClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   &http.Client{},
	}, nil
}

// Search는 쿼리와 정렬 방식으로 상품을 검색합니다.
func (c *NaverClient) Search(query, sort string, display int) (*shopping.SearchResponse, error) {
	params := url.Values{}
	params.Set("query", query)
	params.Set("sort", sort)
	params.Set("display", fmt.Sprintf("%d", display))

	req, err := http.NewRequest("GET", apiURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("요청 생성 실패: %w", err)
	}

	req.Header.Set("X-Naver-Client-Id", c.clientID)
	req.Header.Set("X-Naver-Client-Secret", c.clientSecret)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API 호출 실패: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API 오류 (상태코드: %d)", resp.StatusCode)
	}

	var result shopping.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	return &result, nil
}
