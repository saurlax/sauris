# Architecture direction

This document records design direction. The repository currently implements only bootstrap executables; the components below are planned.

## Core and interfaces

The control plane is a modular Go service. An HTTP API is the shared service boundary for a console, CLI, and MCP server. Interface adapters translate requests and responses; domain services own resource rules, authorization, plans, operations, and audit behavior.

Keep pure domain and planning packages independent of backend SDKs. Introduce small consumer-owned interfaces as actual implementations need them; do not build a universal provider framework before a working vertical slice.

## Intent, plans, and operations

The control flow is Intent → Observe → Plan → Execute → Observe.

- Resources record durable intent and ownership.
- Observations report actual state, provenance, and freshness; unknown is distinct from healthy.
- Plans describe changes, preconditions, impact, and supported recovery behavior.
- Operations record execution steps, partial outcomes, cancellation, and completion.
- Controllers observe convergence without confusing request acceptance with application health.

A plan can be applied automatically within delegated policy. Interactive approval is required only when policy requires it. An approved plan must not turn into an unrelated write after state changes.

## Backends and connections

Docker and Kubernetes are distinct providers. A host agent operates Docker on its host. A cluster connector, when needed, reaches Kubernetes through a scoped identity; it does not replace Kubernetes scheduling or require an agent on every cluster node.

Direct API access and connector-mediated access are connection choices. They must not change the domain contract. In-cluster discovery is an optional enrollment convenience, not automatic permission to manage the hosting cluster.

Backend capabilities remain explicit. Docker may use replacement containers; Kubernetes may use Deployment rollouts. Local volumes, PVCs, replica scheduling, and concurrency controls are not interchangeable guarantees.

## Compute and network boundaries

A service endpoint identifies a service and its reachability context, independently of whether it comes from a container, Kubernetes Service, host process, or external device. Loopback addresses and private names are interpreted in an explicit node, connector, or network context.

Private networks, reverse tunnels, HTTP gateways, identity-aware access, and forward proxies are separate capabilities. Select a small number of integrations first. Reuse established networking implementations rather than implementing tunnel or VPN protocols in the control plane.

## State and ownership

PostgreSQL is the intended control-plane store, not a dependency of the bootstrap scaffold. It will persist Sauris-owned intent, enrollment, relationships, policy, and operation history. Providers remain authoritative for observed backend state.

Discovered resources are initially observed. Adoption is explicit. Detect external managers such as Compose projects, Helm, Argo CD, and Flux where possible, and avoid reconciling against them without a deliberate ownership decision.

## Deployment direction

Start with a binary or Compose-based control plane and systemd-managed host agents. Offer Kubernetes packaging later without making it a requirement. Agents should initiate authenticated connections so private hosts do not need public management ports.

Initially colocate control-plane modules. Gateways, relays, and build workers have separate runtime roles and can move to dedicated hosts as needed. Application traffic bypasses the control-plane API.

Existing workloads should continue through a temporary control-plane outage where their backend permits it. Gateways and agents need explicit last-applied configuration, credential expiry, and revocation behavior. Retain a recovery path that does not depend solely on the overlay being managed.

## Delivery sequence

1. Define core resource, provider, plan, and operation contracts through OpenSpec.
2. Build one Docker deployment slice: enroll a host, deploy an image, observe status, and read logs through API and CLI.
3. Add a small MCP tool surface using the same services and authorization.
4. Connect service endpoints to one gateway/tunnel implementation and access policy.
5. Add Git build and delivery, followed by an existing-cluster Kubernetes provider.
6. Expand the console, network providers, and backup/recovery workflows around verified behavior.

This is sequencing guidance, not a promise that every planned feature will be implemented in the first release.

## References

- [Zeabur OSS architecture](https://github.com/zeabur-oss/zeabur/blob/main/docs/reference/architecture.mdx): resource projections, ownership, and planned changes. Its Kubernetes-backed persistence and user-token execution model are not adopted wholesale.
- [Zeabur OSS module boundaries](https://github.com/zeabur-oss/zeabur/blob/main/docs/internal/crates.md): narrow interfaces and pure planning logic.
- [Agent Foundation README](https://github.com/converge-ai-labs/agent-foundation/blob/main/README.md): product introduction, component choices, and executable source-development instructions.
