# Katartismos

Provisioner interface for GitOps on Kubernetes. Defines the contract for
answering: **how do we make them exist?**

## What This Is

A Go module containing an interface definition and its associated types. Not a
controller, not a server, not a CLI. It is the contract that provisioners
implement.

A provisioner takes a resource inventory — the set of Kubernetes resources that
should exist — and applies, deletes, or diffs them against a cluster.

## Why It Exists

Every GitOps system provisions resources, but they all do it internally, tightly
coupled to their own reconcile and resolution logic. Katartismos extracts the
operation into a standalone interface so that:

- Provisioner implementations are independently testable
- Orchestrators can swap provisioners without changing their reconcile loop
- The ecosystem can converge on a shared contract — a native Kubernetes
  implementation, a Flux-based adapter, or an ArgoCD-backed service can all
  satisfy the same interface

## Ecosystem

Katartismos is one of two GitOps primitive interfaces defined by
[katastroma](https://github.com/katastroma). The other is
[keleustēs](https://github.com/katastroma/keleustes) (the resolver interface).

[Histia](https://github.com/katastroma/histia) is katastroma's reference
provisioner implementation.
