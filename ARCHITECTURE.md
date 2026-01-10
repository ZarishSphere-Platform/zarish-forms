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
