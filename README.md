# Sink

`Sink` is a high-performance, multi-tenant form backend engine built from first principles in Go. It acts as a centralized data ingestion layer, allowing multiple independent tenants to safely submit, process, and store unstructured or semi-structured form data into a relational PostgreSQL database.

Rather than relying on rigid, high-level frameworks, Sink handles raw data processing, robust tenant isolation, and concurrent database persistence with low-level precision and minimal dependencies.
