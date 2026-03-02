# Go 코드 패턴 학습 리뷰: Naver Shopping Skill

- **작성일:** 2026-03-02
- **대상 파일:** `internal/client/naver.go`, `internal/shopping/models.go`, `internal/shopping/service.go`, `cmd/main.go`

---

## 1. JSON 언마샬링 (JSON Unmarshaling)

### 코드 예시 (`internal/shopping/models.go`)
```go
type SearchResponse struct {
    Total   int       `json:"total"`
    Items   []Product `json:"items"`
}

type Product struct {
    Title  string `json:"title"`
    Lprice string `json:"lprice"`
}
```

### 설명

Go에서 외부 API의 JSON 응답을 받을 때는 **구조체(struct)** 와 **구조체 태그(struct tag)** 를 사용한다.

- `` `json:"total"` `` — JSON 키 이름을 명시적으로 지정한다. JSON 키가 소문자(`total`)여도 Go 필드명은 대문자(`Total`)로 작성한다 (exported field).
- `[]Product` — JSON의 배열(`[]`)은 Go의 슬라이스 타입과 자동으로 매핑된다.
- `json.NewDecoder(resp.Body).Decode(&result)` — HTTP 응답 바디를 스트리밍으로 파싱한다. 전체를 메모리에 올리지 않아 효율적이다.

> **핵심 원칙:** JSON 필드명과 Go 필드명이 다를 때는 반드시 struct tag를 사용한다.

---

## 2. 에러 래핑 (Error Wrapping)

### 코드 예시 (`internal/client/naver.go`)
```go
req, err := http.NewRequest("GET", apiURL+"?"+params.Encode(), nil)
if err != nil {
    return nil, fmt.Errorf("요청 생성 실패: %w", err)
}
```

### 설명

Go는 예외(exception) 대신 **에러 반환** 방식을 사용한다.

- `fmt.Errorf("메시지: %w", err)` — `%w` 동사는 원본 에러를 새 에러 안에 **감싼다(wrap)**. 호출자가 `errors.Is()` 나 `errors.As()` 로 원본 에러를 추출할 수 있다.
- `if err != nil` 패턴 — Go의 관용적 에러 처리 방식. 함수 호출 직후 에러를 확인한다.
- `return nil, err` — 에러가 있으면 즉시 반환하여 이후 코드 실행을 막는다 (early return).

> **핵심 원칙:** 에러에 컨텍스트를 추가할 때는 `fmt.Errorf("%w", err)` 로 원본 에러를 보존한다.

---

## 3. HTTP 클라이언트 패턴

### 코드 예시 (`internal/client/naver.go`)
```go
type NaverClient struct {
    clientID     string
    clientSecret string
    httpClient   *http.Client
}

func New() (*NaverClient, error) {
    // ...
    return &NaverClient{
        httpClient: &http.Client{},
    }, nil
}
```

### 설명

- **생성자 함수 패턴:** Go에는 `new` 키워드 대신 `New()` 같은 생성자 함수를 관례적으로 사용한다. 초기화 중 에러가 발생할 수 있을 때는 `(*T, error)` 형태로 반환한다.
- **`http.Client` 재사용:** `http.Client`는 내부적으로 커넥션 풀을 관리한다. 요청마다 새로 생성하지 않고 구조체 필드로 보관하여 재사용하는 것이 효율적이다.
- **소문자 필드 (unexported):** `clientID`, `clientSecret` 은 패키지 외부에서 직접 접근 불가. 캡슐화를 통해 인증 정보를 보호한다.

> **핵심 원칙:** 상태를 가진 클라이언트는 구조체로 묶고, 생성자 함수로 초기화한다.

---

## 4. 정규식과 문자열 처리

### 코드 예시 (`internal/shopping/service.go`)
```go
var htmlTagRe = regexp.MustCompile(`<[^>]+>`)

func stripTags(s string) string {
    return htmlTagRe.ReplaceAllString(s, "")
}
```

### 설명

- `regexp.MustCompile()` — 패키지 수준 변수로 정규식을 **딱 한 번** 컴파일한다. 함수 안에서 매번 컴파일하면 성능이 떨어진다. `MustCompile`은 패턴이 잘못되면 패닉을 일으키므로, 런타임이 아닌 **초기화 시점**에 오류를 발견할 수 있다.
- `` `<[^>]+>` `` — 백틱(`` ` ``) 문자열 리터럴은 이스케이프 없이 정규식을 그대로 쓸 수 있다.
- 네이버 API는 검색어 일치 부분에 `<b>태그</b>` 를 삽입하므로, 출력 전 반드시 제거해야 한다.

> **핵심 원칙:** 반복 사용하는 정규식은 패키지 변수로 선언해 한 번만 컴파일한다.

---

## 5. 환경변수와 `.env` 파일 로드

### 코드 예시 (`cmd/main.go`)
```go
_ = godotenv.Load()

clientID := os.Getenv("NAVER_CLIENT_ID")
```

### 설명

- `godotenv.Load()` — `.env` 파일을 읽어 환경변수로 등록한다. 파일이 없어도 에러를 무시(`_`)하여 CI/CD 등 환경변수를 직접 설정하는 환경에서도 동작한다.
- **OS 환경변수 우선:** `Load()`는 이미 설정된 환경변수를 덮어쓰지 않는다. 운영 환경(서버, CI)에서 안전하게 동작한다.
- `os.Getenv()` — 환경변수가 없으면 빈 문자열(`""`)을 반환한다. 빈 값 체크로 미설정 여부를 판단한다.

> **핵심 원칙:** 인증 정보는 절대 코드에 하드코딩하지 않는다. `.env`는 `.gitignore`에 추가한다.

---

## 6. `defer`를 이용한 리소스 해제

### 코드 예시 (`internal/client/naver.go`)
```go
resp, err := c.httpClient.Do(req)
if err != nil {
    return nil, fmt.Errorf("API 호출 실패: %w", err)
}
defer resp.Body.Close()
```

### 설명

- `defer` — 현재 함수가 **종료될 때** (return, panic 포함) 지정한 코드를 실행한다.
- `resp.Body.Close()` — HTTP 응답 바디는 반드시 닫아야 커넥션이 풀로 반환된다. `defer`를 쓰면 에러 반환 경로가 여러 개여도 빠짐없이 닫힌다.
- `err != nil` 체크 **이후**에 `defer`를 선언한다 — `err != nil` 이면 `resp`가 `nil`일 수 있으므로 순서가 중요하다.

> **핵심 원칙:** 파일, HTTP 바디, DB 커넥션 등 외부 리소스는 열자마자 `defer`로 닫아둔다.
