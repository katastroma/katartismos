# Katartismos

Provisioner interface for [katastroma](https://github.com/katastroma). Defines
the client-facing API for provisioner implementations.

A provisioner takes manifests, labels, and a service account identity. It
applies the manifests, stamps them with the labels, and removes any resources in
the cluster with those labels that are not in the manifests (pruning).
