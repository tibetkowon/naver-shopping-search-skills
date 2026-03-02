# Skill: Verify Go Implementation

## Description
Ensures code quality and build stability for the local Go binary before any commit or task completion.

## Trigger
- After modifying any Go files.
- Before declaring a task finished.

## Instructions
1. **Format:** Run `go fmt ./...`.
2. **Build:** Run `go build -o shop ./cmd/main.go` (or equivalent) to ensure zero compilation errors.
3. **Test:** Run `go test ./...` if unit tests are present.
4. **Report (Korean):** Notify the user: "Go 코드 빌드 및 포맷팅 검증이 완료되었습니다."
