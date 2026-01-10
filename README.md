# Zarish Forms Platform (ZFP)


**Zarish Forms** is the canonical, spreadsheet-first, schema-driven, no-code Form & App Definition Platform for the ZarishSphere – NGO Operations Platform.

It enables NGOs, governments, and humanitarian programs to define forms, data models, validation rules, UI behavior, workflows, and interoperability mappings using Google Sheets / Excel, without writing code.

Zarish Forms is a governance layer, not a data system.

---

## 🌍 Why Zarish Forms Exists

Humanitarian software often fails because:

- Forms are hard-coded
- Changes require developers
- Validation is inconsistent
- Audits are difficult
- Interoperability is an afterthought
- Zarish Forms solves this by making forms a standard, not an implementation.

---


## 🎯 Core Principles (Non-Negotiable)

1. **Spreadsheet First**
NGOs define everything in tabular formats they already know.

2. **Schema as Truth**
Canonical JSON schemas are the single source of technical truth.

3. **Preview Before Deploy**
Nothing goes live without human validation.

4. **GitHub Native Governance**
Versioned, auditable, reviewable.

5. **Domain Agnostic**
Health, HR, Finance, M&E, Education, Relief.

6. **FHIR Aligned, Not Locked**
Interoperable by design.


---


## 🧠 What Zarish Forms IS (and IS NOT)

### ✅ Zarish Forms IS**
- A form definition platform
- A schema factory
- A UI & API contract generator
- A validation and governance layer

### ❌ Zarish Forms is NOT**
- A database
- A production backend
- A workflow execution engine
- An EMR / HR / Finance system

> Zarish Forms **defines**.
> ZarishSphere services **execute**.

---


## 🧱 High-Level Architecture


```text

Spreadsheet (Google Sheets / Excel)
↓
Zarish Forms (Go 1.25)
↓
Canonical JSON Schemas (Truth)
↓
 ┌─────────────┬─────────────┬─────────────┐ │ Frontend UI │ Backend     │ Interop     │ │ (React)     │ Artifacts   │ (FHIR)      │ └─────────────┴─────────────┴─────────────┘

```


---


## 🧰 Technology Stack (Authoritative)

**Core Platform**

- Language: Go 1.25

- Schema Standard: JSON Schema (2020-12)

- Spreadsheet Parsing: excelize

- Schema Validation: santhosh-tekuri/jsonschema

- Workflow Model: JSON state machines

- Frontend (Preview & UX): React 18,  Next.js 14

- Schema-driven rendering

- Interoperability

- HL7 FHIR R4 (aligned)

- OpenMRS compatible concepts

- DHIS2-friendly exports


---


## 📦 Repository Structure


text```

zarish-forms/
├─ README.md
├─ VERSION
│
├─ spreadsheets/        # NGO inputs (SOURCE)
├─ schemas/             # Canonical JSON (TRUTH)
│  ├─ v1/
│  └─ archive/
│
├─ previews/            # Human validation (SAFETY)
│  ├─ html/
│  └─ react/
│
├─ workflows/           # Business state models
├─ standards/           # Rules & contracts
│
├─ cmd/                 # Go CLI entrypoints
├─ internal/            # Go core logic
│
└─ .github/workflows/   # CI & governance

```


---


## 📘 Standards (MANDATORY READING)

All contributors MUST follow these standards:


| Standard | Description |
| :--- | :--- |

| spreadsheet.standard.md | How NGOs define forms |
| schema.standard.md | Canonical JSON schema rules |
| ui.standard.md | UI rendering rules |
| validation.standard.md | Validation grammar |
| Workflow standard | State & approval rules |
| FHIR mapping | Interoperability alignment |



*Any change violating standards will be rejected.*


---


## 🧾 Spreadsheet-First Model

- One spreadsheet = one form

- One row = one field

- No formulas

- No macros

- Values only


**This makes Zarish Forms usable by:**

- NGO field staff

- Program officers

- Health workers

- Administrators

---


## 🔐 Governance & Safety


### Zarish Forms enforces:

- Mandatory schema validation

- Mandatory preview generation

- Mandatory human approval

- Mandatory GitHub review

- Immutable schema history

**No silent changes**.

**No production surprises**.


---


## 🔁 Lifecycle (End-to-End)

- NGO fills spreadsheet

- Spreadsheet committed to GitHub

- Go compiler generates schema

- Schema validated (CI)

- Preview generated

- Human approval

- Downstream services consume schema

---


## 🧬 FHIR & Interoperability

### Zarish Forms aligns with:

**FHIR R4**

- Patient, Encounter, Observation, Practitioner

- Questionnaire / QuestionnaireResponse


*FHIR is supported, not enforced.*

---


## 🗺️ Module Rollout Strategy

**ZarishSphere modules adopt Zarish Forms in phases:**

- EMR (core health)

- HR

- Finance

- Programs & M&E

- Education & Relief

**Each module is:**

- Independent

- Versioned

- Auditable

---

## 👥 Who This Is For

- NGOs & INGOs

- Government health systems

- Donors & auditors

- Platform engineers

- Long-term humanitarian programs

---

## 📜 Versioning Policy

- **MAJOR** → breaking schema changes

- **MINOR** → backward-compatible additions

- **PATCH** → metadata / documentation

---

## 🤝 Contribution Rules

- No direct schema editing

- No skipping previews

- No bypassing standards

- No undocumented changes


**Zarish Forms is a shared public good.**

---


## 🏁 Final Statement

**Zarish Forms** is not a form builder.
It is a governance standard for humanitarian software, designed to outlive technologies, teams, and funding cycles.

**If it is not defined in Zarish Forms, it does not exist in ZarishSphere.**

---

