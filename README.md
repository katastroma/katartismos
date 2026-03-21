# Katartismos

Provisioner interface for GitOps on Kubernetes. Defines the gRPC service
contract for applying manifests to a cluster and pruning resources no longer in
the rendered output.

A provisioner takes manifests and applies them to the cluster using server-side
apply with Kubernetes impersonation. It prunes resources labeled with the
tenant's identity that are no longer present in the current render.

## Why It Exists

Provisioning is one of three fundamental GitOps operations (fetch, render,
provision). Katartismos extracts the provisioning contract so that:

- Provisioner implementations are independently testable and deployable
- The orchestrator ([pharos](https://github.com/katastroma/pharos)) can call
  any provisioner that satisfies the contract
- The ecosystem can converge on a shared contract instead of each project
  coupling provisioning into a monolith

## Ecosystem

Katartismos is one of three GitOps service interfaces defined by
[katastroma](https://github.com/katastroma):

- [naukleros](https://github.com/katastroma/naukleros) — retriever interface
- [keleustēs](https://github.com/katastroma/keleustes) — renderer interface
- **katartismos** (this) — provisioner interface

[Histia](https://github.com/katastroma/histia) is katastroma's provisioner
implementation.
