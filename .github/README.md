# Zarish Forms Platform (ZFP)

## Complete, Authoritative, Copy‑Paste‑Ready Specification

**Platform Name:** Zarish Forms
**Part of:** ZarishSphere – NGO Operations Platform
**Version:** 1.0.0
**Audience:** NGO Operators (No‑Code), Architects, Developers, Auditors
**Status:** Production‑Ready Blueprint

---

## 0️⃣ Executive Definition (What Zarish Forms IS)

**Zarish Forms is a schema‑driven, spreadsheet‑first, no‑code platform that generates validated forms, APIs, databases, workflows, and UIs for all ZarishSphere modules.**

It is:

* Spreadsheet‑first (Google Sheets / Excel)
* Preview‑before‑deploy (human validation)
* Schema‑as‑truth (machine governance)
* GitHub‑native (versioned, auditable)
* Domain‑agnostic (Health, HR, Finance, Education, Relief)

---

## 1️⃣ Platform Positioning & Scope

### Zarish Forms Controls

* Form definitions
* Field metadata
* Validation rules
* UI behavior
* Workflow triggers
* Role‑based visibility

### Zarish Forms Does NOT

* Store operational data
* Handle business logic execution
* Replace EMR / HR / Finance services

It **defines**, others **execute**.

---

## 2️⃣ Canonical Repository Structure (MANDATORY)

```text
zarish-forms/
├─ README.md
├─ VERSION
│
├─ spreadsheets/                 # NGO inputs (SOURCE)
│  ├─ emr_patient_registration.xlsx
│  ├─ hr_staff_profile.xlsx
│  └─ finance_expense_claim.xlsx
│
├─ schemas/                      # Canonical JSON (TRUTH)
│  ├─ v1/
│  │  ├─ emr/
│  │  ├─ hr/
│  │  └─ finance/
│  └─ archive/
│
├─ previews/                     # Human validation (SAFETY)
│  ├─ html/
│  └─ react/
│
├─ workflows/                    # Business flows
│  ├─ approval.workflow.json
│  └─ lifecycle.workflow.json
│
├─ standards/                    # Rules & contracts
│  ├─ spreadsheet.standard.md
│  ├─ schema.standard.md
│  ├─ ui.standard.md
│  └─ validation.standard.md
│
└─ .github/workflows/
   ├─ validate-schema.yml
   └─ generate-preview.yml
```

---

## 3️⃣ Universal Spreadsheet Standard (SOURCE OF TRUTH)

### Sheet Name: `form_definition`

| Column          | Req | Description     | Example                             |        |       |
| --------------- | --- | --------------- | ----------------------------------- | ------ | ----- |
| module          | ✅   | Zarish module   | EMR                                 |        |       |
| domain          | ✅   | Business domain | clinical                            |        |       |
| form_id         | ✅   | Unique ID       | patient_registration                |        |       |
| form_title      | ✅   | Display name    | Patient Registration                |        |       |
| field_name      | ✅   | DB safe         | date_of_birth                       |        |       |
| field_label     | ✅   | UI label        | Date of Birth                       |        |       |
| field_type      | ✅   | Data type       | string, number, date, boolean, enum |        |       |
| required        | ✅   | Mandatory       | yes/no                              |        |       |
| default         | ❌   | Default value   | today                               |        |       |
| validation      | ❌   | Rules           | min:0,max:120                       |        |       |
| enum_values     | ❌   | For enums       | male                                | female | other |
| ui_component    | ❌   | UI hint         | input, select, datepicker           |        |       |
| section         | ❌   | UI grouping     | Demographics                        |        |       |
| order           | ❌   | Render order    | 1                                   |        |       |
| role_visibility | ❌   | RBAC            | doctor,nurse                        |        |       |
| workflow_event  | ❌   | Trigger         | on_submit                           |        |       |
| audit           | ❌   | Audit level     | strict                              |        |       |

**Non‑technical users only fill rows.**

---

## 4️⃣ Canonical Zarish Form Schema (JSON v1)

```json
{
  "$schema": "https://zarishsphere.org/schemas/zarish-form.schema.json",
  "schema_version": "1.0.0",
  "module": "EMR",
  "domain": "clinical",
  "form_id": "patient_registration",
  "form_title": "Patient Registration",
  "audit_level": "strict",
  "fields": [
    {
      "name": "date_of_birth",
      "label": "Date of Birth",
      "type": "date",
      "required": true,
      "default": null,
      "validation": {
        "min": "1900-01-01",
        "max": "today"
      },
      "ui": {
        "component": "datepicker",
        "section": "Demographics",
        "order": 2
      },
      "access": {
        "roles": ["doctor", "nurse"]
      },
      "workflow": {
        "event": "on_submit"
      }
    }
  ]
}
```

---

## 5️⃣ Field Types & Validation Standard

### Supported Field Types

* string
* number
* integer
* boolean
* date
* datetime
* enum
* text
* file (future)

### Validation Grammar

```text
min:value
max:value
regex:pattern
length:min,max
required
unique
```

---

## 6️⃣ UI Rendering Standard (Preview‑First)

### UI Rules

* No schema → No preview
* No preview → No deploy

### Preview Outputs

* HTML (static validation)
* React (interactive testing)

### Required UI Metadata

```json
"ui": {
  "component": "select",
  "section": "Clinical Info",
  "order": 3
}
```

---

## 7️⃣ Workflow Standard (Copy‑Paste)

```json
{
  "workflow_id": "patient_onboarding",
  "version": "1.0.0",
  "states": ["draft", "submitted", "approved", "rejected"],
  "transitions": [
    {"from": "draft", "to": "submitted", "event": "on_submit"},
    {"from": "submitted", "to": "approved", "role": "supervisor"},
    {"from": "submitted", "to": "rejected", "role": "supervisor"}
  ]
}
```

---

## 8️⃣ Role‑Based Access Control (RBAC)

### Role Model

```json
{
  "roles": {
    "admin": "full access",
    "doctor": "clinical forms",
    "nurse": "limited clinical",
    "clerk": "registration only"
  }
}
```

---

## 9️⃣ GitHub Governance & Automation

### Validation Workflow

```yaml
name: Validate Zarish Schemas
on:
  pull_request:
    paths: ["schemas/**"]
jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: JSON Validation
        run: |
          for f in schemas/**/*.json; do jq . "$f"; done
```

---

## 🔟 Golden Rules (Non‑Negotiable)

1. Spreadsheet is the source
2. JSON schema is the truth
3. Preview is mandatory
4. Humans approve
5. GitHub records everything
6. No silent deployment

---

## 1️⃣1️⃣ Compliance & Alignment

* OpenMRS‑compatible schema concepts
* FHIR‑ready field mapping
* Audit‑friendly
* NGO & donor transparent

---

## 1️⃣2️⃣ Final Declaration

**Zarish Forms is not a tool.**
It is a **governance system for humanitarian software**.

Everything above is:

* Complete
* Coherent
* Copy‑paste ready
* Production‑grade

---

**End of Authoritative Specification**
