# Requirements

## Project goal
Build a backend LiveOps engine for games that supports configurable activities and quests,
with a clear state machine, idempotent reward settlement, and operational observability.

The project is intended as an engineering-focused system rather than a gameplay demo.

## In scope
- Activity / quest definition via configuration (JSON/YAML)
- Quest state machine:
  - NotStarted
  - Active
  - Claimable
  - Claimed
  - Expired
- Player event ingestion (e.g. login, progress, completion)
- Idempotent reward claiming
- Audit logging for critical state transitions
- Basic metrics and structured logging

## Out of scope
- Game client or UI
- Real-time combat or gameplay simulation
- Payment, shop, or gacha systems
- Multi-region or cross-datacenter deployment
- Strong consistency across multiple external services

## Non-goals
- Visual polish or presentation
- Feature completeness comparable to a production game
- Full admin tooling or content pipeline

## Success criteria
- A single activity can be configured and run end-to-end
- Duplicate reward requests do not result in duplicate grants
- Quest state transitions are explainable via logs or replay
- The system can run stably for long periods
