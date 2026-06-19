# Sink

`Sink` is a high-performance, multi-tenant form backend engine built from first principles in Go. It acts as a centralized data ingestion layer, allowing multiple independent tenants to safely submit, process, and store unstructured or semi-structured form data into a relational PostgreSQL database.

Rather than relying on rigid, high-level frameworks, Sink handles raw data processing, robust tenant isolation, and concurrent database persistence with low-level precision and minimal dependencies.

## Features

* **Strict Multi-Tenancy:** Complete data separation ensures that tenants can only configure, access, and manage their own forms and submissions.
* **Dynamic Form Ingestion:** Schema-flexible handling of payload data, mapping incoming multi-tenant JSON fields seamlessly to relational PostgreSQL storage.
* **Lightweight & Fast:** Built in Go using standard library primitives and direct driver connections for minimal overhead and maximum throughput.
* **Context-Driven Architecture:** Leverages Go’s `context` package to propagate tenant metadata safely across API and database boundaries.

## Architecture

Sink operates on a decoupled design where the ingestion layer is strictly separated from the core persistence layer: