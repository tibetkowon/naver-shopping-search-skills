# 네이버 검색 > 쇼핑 API 명세서

네이버 쇼핑 검색 결과 조회를 위한 RESTful API입니다. 검색어와 검색 조건을 쿼리 스트링으로 전달하여 XML 또는 JSON 형식의 결과를 반환받을 수 있습니다.

## 1. 개요
* **특징**: 비로그인 방식 오픈 API (HTTP 헤더에 클라이언트 아이디와 시크릿 전송)
* **호출 한도**: 검색 API 전체 통합 일일 25,000회

## 2. 사전 준비 사항
1. **네이버 개발자 센터**에서 애플리케이션을 등록합니다.
2. **Client ID**와 **Client Secret**을 발급받습니다.
3. 애플리케이션 설정에서 **검색 API** 권한이 추가되어 있는지 확인합니다.

## 3. API 레퍼런스

### 기본 정보
| 항목 | 내용 |
| :--- | :--- |
| **HTTP 메서드** | GET |
| **인증 방식** | HTTP Header에 Client ID, Secret 포함 |
| **결과 형식** | XML, JSON |

### 요청 URL
* **JSON**: `https://openapi.naver.com/v1/search/shop.json`
* **XML**: `https://openapi.naver.com/v1/search/shop.xml`

### 요청 파라미터
| 파라미터 | 타입 | 필수 | 설명 |
| :--- | :--- | :--- | :--- |
| `query` | String | Y | 검색어 (UTF-8 인코딩 필수) |
| `display` | Integer | N | 한 번에 표시할 검색 결과 개수 (기본 10, 최대 100) |
| `start` | Integer | N | 검색 시작 위치 (기본 1, 최대 1000) |
| `sort` | String | N | 정렬 방식: `sim`(정확도), `date`(날짜순), `asc`(가격 오름차순), `dsc`(가격 내림차순) |
| `filter` | String | N | 상품 유형 필터: `naverpay`(네이버페이 연동 상품) |
| `exclude` | String | N | 제외 상품 유형: `used`(중고), `rental`(렌탈), `cbshop`(해외직구) |

### 요청 헤더
| 헤더 이름 | 값 |
| :--- | :--- |
| `X-Naver-Client-Id` | {발급받은 클라이언트 아이디} |
| `X-Naver-Client-Secret` | {발급받은 클라이언트 시크릿} |

---

## 4. 응답 요소 (Response)

| 요소명 | 타입 | 설명 |
| :--- | :--- | :--- |
| `lastBuildDate` | dateTime | 검색 결과를 생성한 시간 |
| `total` | Integer | 총 검색 결과 개수 |
| `start` | Integer | 검색 시작 위치 |
| `display` | Integer | 한 번에 표시할 결과 개수 |
| `items` | Array | 개별 검색 결과 목록 |
| `items/title` | String | 상품 이름 (검색어 일치 부분은 `<b>` 태그 포함) |
| `items/link` | String | 상품 정보 URL |
| `items/image` | String | 섬네일 이미지 URL |
| `items/lprice` | Integer | 최저가 (정보 없으면 0) |
| `items/hprice` | Integer | 최고가 (정보 없으면 0) |
| `items/mallName` | String | 판매 쇼핑몰 이름 |
| `items/productId` | Integer | 네이버 쇼핑 상품 ID |
| `items/productType` | Integer | 상품 타입 (일반, 중고, 단종 등) |
| `items/brand` | String | 브랜드 이름 |
| `items/maker` | String | 제조사 이름 |
| `items/category1~4` | String | 카테고리 대/중/소/세분류 |

---

## 5. 주요 오류 코드

| 코드 | 상태값 | 메시지 | 해결 방법 |
| :--- | :--- | :--- | :--- |
| `SE01` | 400 | Incorrect query request | URL, 파라미터 오타 확인 |
| `SE02` | 400 | Invalid display value | display 값이 1~100 범위인지 확인 |
| `SE03` | 400 | Invalid start value | start 값이 1~1000 범위인지 확인 |
| `SE06` | 400 | Malformed encoding | 검색어가 UTF-8로 인코딩되었는지 확인 |
| `403` | 403 | Forbidden | API 권한 설정(검색) 확인 |
| `SE99` | 500 | System Error | 시스템 내부 에러 (개발자 포럼 문의) |
