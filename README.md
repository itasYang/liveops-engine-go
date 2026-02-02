# liveops-engine-go
A rule-driven LiveOps engine built with Go. It enables configurable activites and quests with a state-machine core, idempotent rewards, auditability, and basic observability.

## What this is
This project is a backend system designed for game LiveOps operations (events, quests, rewards).
It prioritizes correctness, maintainability, and operational readiness over gameplay presentation.

## Core features (MVP)
- Config-driven activities/quests (JSON/YAML)
- Quest state machine (start / active / claimable / claimed / expired)
- Idempotent reward settlement (duplicate requests do not double-grant)
- Audit log for key actions and settlements
- Redis cache for hot player quest state (optional in early phase)
- Basic observability: structured logs + metrics placeholders

## Non-goals
- Game client / Unity / rendering
- Real-time combat, physics, or matchmaking
- Monetization systems (shop, gacha, payments)
- Complex distributed transactions across many services

## High-level architecture
- **API service**: receives player events and exposes query/claim endpoints
- **Core engine**: evaluates rules and drives state transitions
- **Storage**: MySQL for authoritative data, Redis for caching
- **Observability**: logs + metrics

## Project status
- Week 0: repository + docs scaffolding (in progress)
- Next: implement minimal runnable API and core state machine

## Quick start
> Coming soon.


