# Platform Foundation

## Purpose

Define architectural constraints for Sauris as a self-hosted personal cloud for people and AI agents. This is a design baseline; runtime features described here are not implemented by the bootstrap scaffold.

## Requirements

### Requirement: Runtime-independent control plane

The Sauris core control plane SHALL be able to run without Kubernetes and SHALL treat its deployment environment separately from the infrastructure it manages.

#### Scenario: Managing mixed targets
- **GIVEN** Sauris is deployed outside Kubernetes
- **WHEN** an authorized user connects a Docker host and a Kubernetes cluster
- **THEN** both targets can be represented without migrating the control plane into Kubernetes

### Requirement: Shared interface semantics

The platform SHALL apply the same resource semantics, authorization rules, and operation tracking to requests through HTTP API, CLI, web console, and MCP interfaces.

#### Scenario: Equivalent requests through different interfaces
- **GIVEN** two callers have equivalent identity, delegated scope, and requested changes
- **WHEN** one requests a change through the CLI and the other through MCP
- **THEN** the same core authorization and planning rules apply

### Requirement: Explicit backend capabilities

Providers SHALL advertise supported capabilities, remain separate from connection transports, and reject unsupported requested behavior explicitly.

#### Scenario: Unsupported storage behavior
- **GIVEN** a Docker target supports local volumes but not Kubernetes PVCs
- **WHEN** a request requires PVC provisioning on that target
- **THEN** planning reports the unsupported capability before execution

### Requirement: Separate intent and execution

The platform SHALL distinguish durable desired resources, plans computed against observed state, and operations recording execution. Execution SHALL validate relevant preconditions and expose partial failures rather than claim cross-backend atomicity.

#### Scenario: State changes after planning
- **GIVEN** a plan was computed against a particular resource observation
- **WHEN** a conflicting change invalidates that plan's preconditions before execution
- **THEN** execution reports the conflict and requires replanning instead of blindly applying stale changes

### Requirement: Independent compute and network lifecycles

Compute and network resources SHALL have independent lifecycles and SHALL connect through explicit service endpoints with a defined reachability context.

#### Scenario: Publishing an existing host service
- **GIVEN** a service already runs outside Sauris-managed compute
- **WHEN** an authorized user registers its endpoint and configures a supported route
- **THEN** network integration does not require converting the service into a Sauris-managed workload

### Requirement: Explicit ownership and delegated execution

The platform SHALL preserve external resource ownership, distinguish initiating and execution identities, and SHALL NOT treat runtime observations, repository contents, or application logs as authorization to mutate resources.

#### Scenario: Discovering externally managed resources
- **GIVEN** a connected target contains resources managed by another tool
- **WHEN** Sauris discovers them
- **THEN** discovery alone does not adopt them or authorize reconciliation against their existing manager

### Requirement: Model-independent operation

Core platform operations SHALL remain usable without a model provider. AI clients SHALL use the same scoped platform capabilities as other callers.

#### Scenario: Operating without AI credentials
- **GIVEN** no model provider or AI client is configured
- **WHEN** an authorized user uses an implemented platform capability through the CLI or API
- **THEN** absence of a model provider does not prevent that operation
