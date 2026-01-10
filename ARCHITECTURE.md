# Zarish Forms Architecture

## Overview
Zarish Forms is a **schema factory and governance system** within ZarishSphere.
It separates **definition** (forms, workflows) from **execution** (EMR, HR, Finance).

### High-Level Flow


The end-to-end flow of Zarish Forms:

```text
Spreadsheet (Google Sheets / Excel)
        ↓
   Zarish Forms (Go 1.25 Compiler)
        ↓
Canonical JSON Schemas (Truth)
        ↓
 ┌───────────────────┬───────────────────┬───────────────────┐
 │ Frontend UI        │ Backend Artifacts │ Interoperability  │
 │ (React / Next.js)  │ (Go 1.25)        │ (FHIR R4)        │
 └───────────────────┴───────────────────┴───────────────────┘
```

### Components
- **CLI (`cmd/`)** – processes spreadsheets, generates schemas and previews
- **Internal (`internal/`)** – spreadsheet parsing, validation, workflow engine, generator
- **Schemas (`schemas/`)** – canonical JSON schemas
- **Previews (`previews/`)** – HTML/React previews
- **Workflows (`workflows/`)** – approval/lifecycle definitions
- **Standards (`standards/`)** – documentation of rules

### FHIR Mapping
- Patient → `Patient`
- Encounter → `Encounter`
- Vitals → `Observation`
- Staff → `Practitioner`
- Finance → `Basic` extensions
- Questionnaire export supported

### Technology Stack
- Go 1.25, JSON Schema, excelize, fsm (workflow)
- React 18 + Next.js 14 (previews)
- PostgreSQL 15+ downstream


---
