Zarish Forms Platform (ZFP)

Spreadsheet‑Driven, No‑Code Form & App Generation System

Version: 1.0.0
Status: Ready‑to‑Use (Copy–Paste Friendly)
Audience: NGO Operators, Non‑Coders, Architects, Contributors


---

1️⃣ Treat zarish-forms as a PLATFORM (Not a Repo)

Platform Purpose

zarish-forms is the canonical form, schema, and app‑definition platform for all ZarishSphere modules.

It is the single source of truth for:

Forms

Validations

Data models

UI generation

Workflow triggers


Canonical Repository Structure

zarish-forms/
├─ README.md
├─ spreadsheets/          # Raw NGO-friendly inputs
│  ├─ emr_patient.xlsx
│  ├─ hr_staff.xlsx
│  └─ finance_expense.xlsx
│
├─ schemas/               # Generated canonical JSON schemas
│  ├─ emr/
│  │  └─ patient_registration.schema.json
│  └─ hr/
│     └─ staff_profile.schema.json
│
├─ previews/              # Rendered UI previews (auto)
│  └─ emr_patient.html
│
├─ generators/            # Code generators (later)
│  ├─ frontend/
│  ├─ backend/
│  └─ database/
│
├─ workflows/             # Business & approval flows
│  └─ patient_onboarding.workflow.json
│
└─ .github/workflows/
   └─ validate-forms.yml


---

2️⃣ Spreadsheet Column Standard (MASTER TEMPLATE)

🔹 Zarish Universal Form Spreadsheet

This is the MOST IMPORTANT asset

Sheet Name: form_definition

Column Name	Required	Description	Example

module	✅	Platform module name	EMR
form_id	✅	Unique form identifier	patient_registration
form_title	✅	Human readable name	Patient Registration
field_name	✅	Database-safe name	date_of_birth
field_label	✅	UI label	Date of Birth
field_type	✅	Data type	string, number, date, boolean, enum
required	✅	Is mandatory	yes / no
default	❌	Default value	today
validation	❌	Rules	min:0,max:120
enum_values	❌	Pipe-separated	male
ui_component	❌	Preferred UI	text, select, datepicker
section	❌	UI grouping	Demographics
order	❌	Display order	1
role_visibility	❌	RBAC	doctor,nurse
workflow_event	❌	Trigger	on_submit


✅ NGO staff only fill rows — NO CODE


---

3️⃣ Canonical Zarish Form Schema (JSON)

🔹 Schema Philosophy

Spreadsheet → Canonical JSON → Everything else

Strict, versioned, validated


🔹 Zarish Form Schema v1 (Copy‑Paste)

{
  "$schema": "https://zarishsphere.org/schemas/form.schema.json",
  "schema_version": "1.0.0",
  "module": "EMR",
  "form_id": "patient_registration",
  "form_title": "Patient Registration",
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

This schema feeds:

UI rendering

API validation

Database generation

Workflow engines



---

4️⃣ Preview‑First Philosophy (MANDATORY)

🔹 Rule

> ❌ Nothing is deployed unless it is previewed and approved



Preview Lifecycle

1. Spreadsheet uploaded


2. Schema generated


3. Preview rendered (HTML / React)


4. Validation report shown


5. Human approval


6. Commit allowed



Preview Output Example

Form: Patient Registration
---------------------------------
[ Date of Birth  ]  (Required)

✔ Validation: OK
✔ Required fields: Complete
✔ Workflow trigger: on_submit

This protects:

Patients

Data quality

NGO reputation



---

5️⃣ GitHub Actions – Validation First (Simple & Safe)

.github/workflows/validate-forms.yml

name: Validate Zarish Forms

on:
  pull_request:
    paths:
      - "schemas/**"

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Validate JSON
        run: |
          for file in schemas/**/*.json; do
            jq . "$file" > /dev/null
          done

Later you can add:

Schema linting

Preview build

Auto‑commit generators



---

6️⃣ Golden Rules of Zarish Forms

✅ Spreadsheet is king
✅ JSON is truth
✅ Preview before deploy
✅ Modules stay isolated
✅ No NGO staff writes code


---

7️⃣ What You Now Have (Very Important)

You now possess:

A platform definition

A universal spreadsheet standard

A canonical JSON schema

A safe preview workflow

A future‑proof automation path


This is production‑grade architecture, not theory.


---

If you want next, I can:

Create the exact Google Sheet template

Define workflow schema v1

Design RBAC model

Align this with OpenMRS & FHIR


You are building something rare and meaningful. 🌍