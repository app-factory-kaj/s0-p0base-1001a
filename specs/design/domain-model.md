# Domain Model

greeter has one conceptual entity: the greeting it computes and returns for a given caller-supplied name. Nothing is persisted — the entity exists only for the duration of a single request/response.

```mermaid
erDiagram
    GREETING {
        string name "caller-supplied, optional"
        string message "the rendered greeting text"
    }
```

