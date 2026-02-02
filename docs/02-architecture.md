# Architecture

## Overview
The system is structured as a modular backend service with a clear separation between
API handling, core business logic, and infrastructure concerns.

The design prioritizes clarity, testability, and extensibility.

## Core components

### API Layer
- Exposes HTTP endpoints for:
  - Player event ingestion
  - Quest/activity state query
  - Reward claiming
- Handles request validation and authentication (simplified)

### Core Engine
- Evaluates quest rules based on incoming events
- Drives quest state transitions
- Enforces idempotency during reward settlement

### Storage
- MySQL:
  - Authoritative quest state
  - Reward settlement records
  - Audit logs
- Redis (optional in early phase):
  - Cache for hot quest state
  - Rate limiting or deduplication helpers

### Observability
- Structured logging for all critical actions
- Metrics hooks for:
  - Event processing latency
  - Reward grant counts
  - Error rates

## Extensibility
The system is designed to support optional modules via well-defined interfaces.
Additional modules may register:
- New event types
- Custom rule evaluators
- Auxiliary APIs

## Determinism and auditability
All quest state transitions and reward settlements must be traceable via logs.
The same sequence of events should lead to the same resulting state.
