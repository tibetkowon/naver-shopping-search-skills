# naver-shopping-search-skills

네이버 쇼핑 API를 사용하여 상품 가격 확인, 가격 비교, 구매 링크 조회를 수행하는 **OpenClaw 로컬 스킬**입니다. Go 바이너리로 실행되며 AI 에이전트가 직접 호출합니다.

## 기능

| 명령어 | 설명 | 정렬 |
|--------|------|------|
| `check` | 상품 최저가 확인 | 정확도순 |
| `compare` | 쇼핑몰별 가격 비교 | 가격 오름차순 |
| `link` | 상품 구매 링크 제공 | 정확도순 |

## 시작하기

### 1. 사전 요구사항

- Go 1.21 이상
- [네이버 개발자 센터](https://developers.naver.com)에서 발급한 Client ID / Secret

### 2. 빌드

```bash
git clone https://github.com/tibetkowon/naver-shopping-search-skills.git
cd naver-shopping-search-skills
go build -o shop ./cmd/main.go
```

### 3. 환경변수 설정

```bash
cp .env.example .env
```

`.env` 파일을 열어 실제 키 값을 입력합니다:

```
NAVER_CLIENT_ID=your_client_id_here
NAVER_CLIENT_SECRET=your_client_secret_here
```

> OS 환경변수가 `.env` 파일보다 우선합니다.

## 사용법

```bash
./shop <command> <검색어>
```

### 예시

```bash
# 가격 확인
./shop check 아이폰16

# 가격 비교
./shop compare 맥북프로

# 구매 링크
./shop link 에어팟프로
```

### 출력 예시

```
=== 가격 확인 결과 ===
1. Apple 아이폰 16 256GB | ₩1,250,000 | https://shopping.naver.com/...
2. 아이폰16 케이스 투명 | ₩12,900 | https://shopping.naver.com/...

=== 가격 비교 결과 ===
1. Apple MacBook Pro 14 M4 | ₩2,490,000 | 애플코리아공식몰
2. Apple MacBook Pro 14 M4 | ₩2,510,000 | 쿠팡

=== 구매 링크 ===
1. Apple AirPods Pro 2세대 | https://shopping.naver.com/...
2. 에어팟프로2 실리콘케이스 | https://shopping.naver.com/...
```

## 프로젝트 구조

```
naver-shopping-search-skills/
├── cmd/
│   └── main.go                 # CLI 진입점
├── internal/
│   ├── client/
│   │   └── naver.go            # Naver API HTTP 클라이언트
│   └── shopping/
│       ├── models.go           # 응답 구조체
│       └── service.go          # 출력 포맷 로직
├── docs/
│   ├── apis/                   # Naver Shopping API 명세
│   ├── plans/                  # 구현 계획 문서
│   └── reviews/                # Go 패턴 학습 리뷰
├── .env.example                # 환경변수 예시
├── SKILL.md                    # OpenClaw 스킬 매니페스트
└── go.mod
```

## 문서

- [Naver Shopping API 명세](docs/apis/Naver_Shopping_API_Spec.md)
- [구현 계획](docs/plans/initial_implementation_plan.md)
- [Go 패턴 학습 리뷰](docs/reviews/go_patterns_review.md)
