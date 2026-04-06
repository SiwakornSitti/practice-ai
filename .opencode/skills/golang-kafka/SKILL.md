# Golang Kafka Skill

This skill outlines the standards and best practices for integrating Apache Kafka into the Go backend (`app/`) within this repository. 

Load this skill whenever you are tasked with producing or consuming Kafka messages.

## 1. Library Selection
**Mandatory Library:** Use a pure Go library such as `github.com/segmentio/kafka-go` or `github.com/IBM/sarama`.
*   **Reason:** The repository mandates `CGO_ENABLED=0` for minimal Alpine/Scratch Docker containers. Libraries like `confluent-kafka-go` require CGO and `librdkafka`, which breaks this constraint. `segmentio/kafka-go` is generally recommended for its modern context-aware API.

## 2. Clean Architecture Integration

### Consumers (Entrypoints)
Consumers are an entrypoint into the application, exactly like HTTP handlers.
*   **Location:** Place consumers in the `delivery/` folder (e.g., `internal/feature/delivery/kafka_consumer.go`) or a dedicated `worker/` folder if the feature is purely background-processed.
*   **Responsibility:** The consumer's job is to read the message, deserialize it (JSON/Protobuf), extract tracing contexts, and pass the data to the `usecase`. It should NEVER contain business logic.

### Producers (Outbound)
Producing a message is an outbound data operation.
*   **Location:** Define the publisher interface in the `domain/` layer. Implement it in the `repository/` layer (or a specialized `infrastructure/` layer).
*   **Injection:** Inject the publisher interface into the `usecase` so the business logic can emit events without knowing about Kafka.

**Example Domain Interface:**
```go
package domain

import "context"

type UserEventPublisher interface {
    PublishUserCreated(ctx context.Context, user *User) error
}
```

## 3. Best Practices & Strict Rules

### Context & Tracing
*   Always pass `context.Context` down to the Kafka read/write functions.
*   **Headers:** Extract OpenTelemetry/Trace IDs from Kafka headers in the Consumer and inject them into the Go `context`. When producing, inject the Trace ID from the `context` into the Kafka message headers to maintain distributed tracing.

### Graceful Shutdown
*   Kafka consumers **must** be tied into the application's `SIGINT`/`SIGTERM` graceful shutdown mechanism.
*   Upon receiving a shutdown signal, the consumer must stop reading new messages and finish processing the current message before the application exits.
*   Always call `Reader.Close()` or `Client.Close()` to prevent resource leaks and ensure consumer group rebalancing happens quickly.

### Offset Management & Exactly-Once Semantics
*   **Auto-Commit:** Disable auto-commit if message processing is critical. Commit offsets manually *only after* the `usecase` successfully processes the message.
*   **Idempotency:** Because Kafka guarantees "at-least-once" delivery (and exactly-once is hard to configure end-to-end), your `usecase` MUST be idempotent. Check if a message has already been processed (e.g., via a unique message ID or database state) before acting on it.

### Error Handling & Retries
*   If a transient error occurs (e.g., database down), do not commit the offset. Allow the message to be retried, or implement a retry loop with exponential backoff.
*   If a poisonous message (bad JSON, invalid format) is received, log the error, send it to a Dead Letter Queue (DLQ) topic if configured, and **commit the offset** so it does not block the entire partition.