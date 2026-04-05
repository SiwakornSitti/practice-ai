<system_instructions>
<role>
You are an expert AI assistant managing a monorepo consisting of two distinct applications: a Go backend and a Next.js frontend.
</role>

<routing_rules>
<rule>CRITICAL: Do not rely on this root file for implementation details.</rule>
<rule>If working on the Go backend service (`be-agent/`), you MUST read `<path>agents/backend.md</path>` before taking action.</rule>
<rule>If working on the Next.js frontend (`fe-agent/`), you MUST read `<path>agents/frontend.md</path>` before taking action.</rule>
</routing_rules>

<global_constraints>
<constraint>Always `cd` into the appropriate directory (`be-agent/` or `fe-agent/`) before running build tools or scripts.</constraint>
<constraint>Follow 12-Factor App principles. Ensure applications are container-ready, use environment variables for config, log to stdout/stderr, and include optimized Dockerfiles.</constraint>
</global_constraints>
</system_instructions>