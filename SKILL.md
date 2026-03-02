---
name: naver-shopping
description: >
  네이버 쇼핑 API를 사용하여 상품 가격 확인, 가격 비교, 구매 링크 조회를 수행하는 로컬 Go 바이너리 스킬.
  AI 에이전트가 사용자의 쇼핑 관련 질문에 답변할 때 활용한다.
metadata:
  binary: ./shop
  env:
    - NAVER_CLIENT_ID
    - NAVER_CLIENT_SECRET
  build: go build -o shop ./cmd/main.go
---

# Naver Shopping Skill

네이버 쇼핑 검색 API 기반의 로컬 스킬입니다. 상품 가격 확인, 가격 비교, 구매 링크 제공의 세 가지 기능을 제공합니다.

## 사전 요구사항

1. 바이너리 빌드:
   ```bash
   go build -o shop ./cmd/main.go
   ```

2. 환경변수 설정 (`.env` 파일 또는 OS 환경변수):
   ```
   NAVER_CLIENT_ID=your_client_id_here
   NAVER_CLIENT_SECRET=your_client_secret_here
   ```

## Capabilities

| 기능 | 설명 |
|------|------|
| 가격 확인 | 검색어로 상품을 조회하고 최저가를 정확도순으로 표시 |
| 가격 비교 | 동일 상품의 여러 쇼핑몰 가격을 오름차순으로 비교 |
| 구매 링크 | 상품의 네이버 쇼핑 직접 링크 제공 |

## Command-line Usage

```bash
# 가격 확인 (정확도순, 최저가 표시)
./shop check <검색어>

# 가격 비교 (가격 오름차순, 쇼핑몰 표시)
./shop compare <검색어>

# 구매 링크 제공
./shop link <검색어>
```

## Output Format

**check:**
```
=== 가격 확인 결과 ===
1. 상품명 | ₩1,234,567 | https://...
```

**compare:**
```
=== 가격 비교 결과 ===
1. 상품명 | ₩1,234,567 | 쇼핑몰이름
```

**link:**
```
=== 구매 링크 ===
1. 상품명 | https://...
```

## Example Prompts

| 사용자 질문 | 실행 명령어 |
|------------|------------|
| "아이폰16 가격이 얼마야?" | `./shop check 아이폰16` |
| "맥북프로 가격 비교해줘" | `./shop compare 맥북프로` |
| "에어팟프로 살 수 있는 링크 알려줘" | `./shop link 에어팟프로` |
| "갤럭시S25 최저가 알려줘" | `./shop check 갤럭시S25` |
| "다이슨 청소기 쇼핑몰별 가격 비교" | `./shop compare 다이슨청소기` |
