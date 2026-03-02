# 구현 계획: Naver Shopping Local Skill (Go)

- **작성일:** 2026-03-02
- **상태:** 완료 (Implemented)

---

## 1. 개요 및 목표

AI 에이전트가 사용자의 쇼핑 관련 질문에 자율적으로 답변할 수 있도록, 네이버 쇼핑 검색 API를 직접 호출하는 로컬 Go 바이너리 스킬을 구현한다.

**핵심 목표:**
- 환경변수 기반 안전한 인증 정보 관리
- AI 응답에 최적화된 간결한 텍스트 출력 포맷
- `check`, `compare`, `link` 세 가지 명령어로 모든 쇼핑 조회 요구 충족

---

## 2. 아키텍처 결정 사항

### 2-1. 디렉토리 구조
```
naver-shopping-search-skills/
├── cmd/
│   └── main.go              # CLI 진입점 (명령어 파싱 및 라우팅)
├── internal/
│   ├── client/
│   │   └── naver.go         # Naver API HTTP 클라이언트
│   └── shopping/
│       ├── models.go        # API 응답 구조체 (SearchResponse, Product)
│       └── service.go       # 비즈니스 로직 (HTML 제거, 가격 포맷, 출력)
├── go.mod                   # 모듈 정의 + godotenv 의존성
├── go.sum
├── .env.example             # 인증 정보 예시 (커밋 포함)
├── .gitignore               # .env, shop 바이너리 제외
└── SKILL.md                 # OpenClaw 매니페스트
```

### 2-2. 의존성 선택
| 패키지 | 용도 | 선택 이유 |
|--------|------|-----------|
| `github.com/joho/godotenv` | `.env` 파일 로드 | 표준적이고 가볍다. OS 환경변수 우선 보장 |
| 표준 라이브러리 (`net/http`, `encoding/json`) | API 호출 및 파싱 | 외부 의존성 최소화 |

### 2-3. 환경변수 우선순위
```
OS 환경변수 > .env 파일
```
`godotenv.Load()` (override 없음) 사용 → 이미 설정된 OS 환경변수를 덮어쓰지 않음.

---

## 3. 명령어 설계

| 명령어 | sort 파라미터 | 출력 형식 | 사용 시나리오 |
|--------|--------------|-----------|--------------|
| `check`   | `sim` (정확도순) | `번호. 상품명 \| 최저가 \| 링크` | "이 상품 얼마야?" |
| `compare` | `asc` (가격 오름차순) | `번호. 상품명 \| 최저가 \| 쇼핑몰` | "어디가 제일 싸?" |
| `link`    | `sim` (정확도순) | `번호. 상품명 \| 링크` | "구매 링크 알려줘" |

**display 기본값: 5** — AI 응답 컨텍스트 절약 + 핵심 결과만 제공

---

## 4. 출력 포맷 설계

```
=== 가격 확인 결과 ===
1. 상품명 | ₩1,234,567 | https://...
2. 상품명 | ₩2,000,000 | https://...
```

- `<b>` 태그 제거: `regexp` 기반 `stripTags()` 함수
- 가격 포맷: `strconv.ParseInt` → 천 단위 콤마 → `₩` 접두사
- 가격 정보 없음(0): "가격 정보 없음" 출력

---

## 5. 보안 고려사항

- `NAVER_CLIENT_ID`, `NAVER_CLIENT_SECRET` 을 코드에 하드코딩하지 않음
- `.env` 파일은 `.gitignore`에 포함하여 저장소에 노출되지 않도록 처리
- 인증 정보 미설정 시 명확한 에러 메시지와 함께 즉시 종료

---

## 6. 향후 확장 포인트

- `--display N` 플래그로 결과 개수 조절
- `exclude=used:cbshop` 필터로 중고/해외직구 제외 옵션
- `-o json` 플래그로 구조화된 JSON 출력 지원 (에이전트 파싱 최적화)
- 네이버페이 연동 상품만 필터링 (`filter=naverpay`)
