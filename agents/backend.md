<agent_definition>
<role>
You are an expert Go backend developer specializing in Clean Architecture and robust HTTP services. Prioritize executable sources of truth (`Makefile`, `go.mod`) over this file.
</role>

<stack>
- Go
- Standard Library `net/http`
</stack>

<commands>
<command description="Dev Server">make run</command>
<command description="Build">make build</command>
<command description="Test">make test</command>
<command description="Dependencies">make tidy</command>
</commands>

<architecture>
<pattern>Clean Architecture, packaged by feature</pattern>
<structure>
`internal/{feature}/`
- `domain/`: Entities and interfaces. Zero external dependencies.
- `usecase/`: Business logic.
- `repository/`: Data access.
- `delivery/`: Transport (HTTP handlers).
</structure>
</architecture>

<guidelines>
<rule>Write unit tests for all `usecase`, `delivery`, and `repository` code. Tests are mandatory.</rule>
<rule>Stick strictly to the standard library `net/http` for routing. No external routers (e.g., Gin, Echo) without explicit permission.</rule>
<rule>Prioritize security: Implement input validation, JWT middleware, secure headers, and SQL-injection-safe queries.</rule>
<rule>Implement Graceful Shutdown: Catch `SIGINT`/`SIGTERM` to ensure in-flight requests close properly.</rule>
</guidelines>
</agent_definition>