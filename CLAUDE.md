# Project: Naver Shopping Local Skill (Go)

## Overview
This project is a high-performance local tool designed as an **OpenClaw Skill**. It allows an AI agent to search products, compare prices, and provide shopping links using the Naver Shopping API directly on the user's MacBook.

## Tech Stack & Environment
- **Language:** Go (Golang)
- **Runtime:** Local binary execution (No external server)
- **OS:** macOS (MacBook)
- **API Spec:** `./docs/apis/Naver_Shopping_API_Spec.md`

## Available Skills
### Global Skills (Shared across all projects)
- **Location:** `~/.claude/skills/`
- **manage_skills.md**: For evolving project rules and learning patterns.
- **generate_openclaw_spec.md**: Universal tool to update `SKILL.md` with YAML Frontmatter.

### Project-Specific Skills
- **Location:** `.claude/skills/`
- **implement_naver_shopping_feature.md**: Core logic for Naver API integration.
- **plan_shopping_logic.md**: Planning mandates in `docs/plans/`.
- **verify_implementation.md**: Go build/format verification (Project specific setup).
- **write_code_tutor.md**: Educational reviews in Korean.

## Core Instructions
- **Phase 1 (Setup):** Initialize the Go module and create the API client structure based on the spec.
- **Phase 2 (Logic):** Implement "Price Check", "Price Comparison", and "Purchase Link" commands.
- **Phase 3 (Skill Integration):** Use `generate_openclaw_spec.md` to create the final `SKILL.md` manifest.
- **Security:** Always load `NAVER_CLIENT_ID` and `NAVER_CLIENT_SECRET` from environment variables.

## Guidelines
- Follow idiomatic Go patterns (Standard library or `resty`).
- Keep the output concise and optimized for AI agent responses (Name | Price | Link).
- Use Korean for all user-facing reports and documentation as requested.
