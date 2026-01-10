# Zarish Forms – FHIR Mapping Tables & Module-by-Module Rollout Plan

**Platform:** ZarishSphere – Zarish Forms
**Version:** 1.0.0
**Status:**  Implementation-Grade

---

## PART A — FHIR MAPPING TABLES (AUTHORITATIVE)

### A.1 Purpose

These tables define how **Zarish Forms schemas map to HL7 FHIR R4 resources**.

Goals:

* Interoperability with OpenMRS, DHIS2, national HIS
* Donor & regulator compliance
* Future data exchange without refactoring

Zarish Forms remains **FHIR-compatible but not FHIR-dependent**.

---

## A.2 General Mapping Rules

| Zarish Concept | FHIR Concept               |
| -------------- | -------------------------- |
| module         | FHIR resource type         |
| form_id        | profile / questionnaire id |
| field_name     | element path               |
| field_type     | FHIR data type             |
| enum           | CodeableConcept            |
| validation     | FHIR constraints           |

---

## A.3 EMR Module → FHIR Mapping

### Form: Patient Registration

**FHIR Resource:** Patient

| Zarish Field  | Type   | FHIR Path                | FHIR Type | Notes              |
| ------------- | ------ | ------------------------ | --------- | ------------------ |
| patient_id    | string | Patient.identifier.value | string    | Primary identifier |
| full_name     | string | Patient.name.text        | HumanName | Display name       |
| gender        | enum   | Patient.gender           | code      | male/female/other  |
| date_of_birth | date   | Patient.birthDate        | date      | ISO-8601           |
| phone         | string | Patient.telecom.value    | string    | Optional           |
| address       | text   | Patient.address.text     | Address   | Free text          |

---

### Form: Clinical Encounter

**FHIR Resource:** Encounter

| Zarish Field   | Type     | FHIR Path                        | Notes           |
| -------------- | -------- | -------------------------------- | --------------- |
| encounter_id   | string   | Encounter.identifier.value       |                 |
| patient_id     | string   | Encounter.subject.reference      | Patient/{id}    |
| encounter_date | datetime | Encounter.period.start           |                 |
| encounter_type | enum     | Encounter.type                   | CodeableConcept |
| provider       | string   | Encounter.participant.individual |                 |

---

## A.4 Observation / Vitals Mapping

**FHIR Resource:** Observation

| Zarish Field     | Type     | FHIR Path                       | Example  |
| ---------------- | -------- | ------------------------------- | -------- |
| observation_code | enum     | Observation.code                | BP, TEMP |
| value            | number   | Observation.valueQuantity.value | 120      |
| unit             | string   | Observation.valueQuantity.unit  | mmHg     |
| recorded_at      | datetime | Observation.effectiveDateTime   |          |

---

## A.5 HR Module → FHIR Mapping

**FHIR Resource:** Practitioner

| Zarish Field | Type   | FHIR Path                       | Notes |
| ------------ | ------ | ------------------------------- | ----- |
| staff_id     | string | Practitioner.identifier.value   |       |
| full_name    | string | Practitioner.name.text          |       |
| role         | enum   | Practitioner.qualification.code |       |
| phone        | string | Practitioner.telecom.value      |       |

---

## A.6 Finance Module (FHIR-Compatible, Not Native)

FHIR has no core finance resource for NGO expenses.
Mapping is **FHIR-compatible extension model**.

**Base Resource:** Basic

| Zarish Field | Type   | FHIR Path                | Notes            |
| ------------ | ------ | ------------------------ | ---------------- |
| claim_id     | string | Basic.identifier.value   |                  |
| amount       | number | Basic.extension.amount   | Custom extension |
| currency     | string | Basic.extension.currency |                  |
| reason       | text   | Basic.code.text          |                  |

---

## A.7 FHIR Questionnaire Compatibility

Zarish Forms can be auto-exported to:

* `Questionnaire`
* `QuestionnaireResponse`

| Zarish Element | FHIR Questionnaire              |
| -------------- | ------------------------------- |
| form_id        | Questionnaire.id                |
| form_title     | Questionnaire.title             |
| fields         | Questionnaire.item[]            |
| required       | Questionnaire.item.required     |
| enum           | Questionnaire.item.answerOption |

---

## PART B — MODULE-BY-MODULE ROLLOUT PLAN

### B.1 Rollout Philosophy

Principles:

* Start small
* Validate early
* Expand horizontally
* Never break production

Each module is **independently deployable**.

---

## B.2 Phase 0 — Foundation (Month 0)

**Objective:** Platform readiness

* zarish-forms repository live
* Spreadsheet standard approved
* Schema standard locked
* GitHub validation enabled

✅ Output: Governance-ready platform

---

## B.3 Phase 1 — EMR Core (Months 1–2)

**Modules:** EMR

**Forms:**

* Patient Registration
* Encounter
* Vitals

**Deliverables:**

* Spreadsheet templates
* JSON schemas
* HTML previews
* FHIR Patient + Observation mapping

**Success Criteria:**

* Non-coder creates patient form
* Preview approved
* Schema consumed by EMR service

---

## B.4 Phase 2 — HR Module (Month 3)

**Forms:**

* Staff Profile
* Role Assignment

**FHIR:** Practitioner

**Why HR next:**

* Low clinical risk
* High operational value

---

## B.5 Phase 3 — Finance Module (Month 4)

**Forms:**

* Expense Claim
* Budget Request

**Notes:**

* Uses FHIR-compatible extensions
* Strong audit & approval workflows

---

## B.6 Phase 4 — Programs & M&E (Month 5)

**Forms:**

* Program Enrollment
* Indicator Tracking

**Integration:**

* DHIS2-compatible exports

---

## B.7 Phase 5 — Education & Relief (Month 6+)

**Education:**

* Student Registration
* Attendance

**Relief:**

* Beneficiary Registration
* Distribution Tracking

---

## B.8 Rollout Dependency Matrix

| Module  | Depends On              |
| ------- | ----------------------- |
| EMR     | Zarish Forms core       |
| HR      | Forms + RBAC            |
| Finance | HR + Approval workflows |
| M&E     | EMR + Finance           |
| Relief  | EMR + M&E               |

---

## FINAL DECLARATION

With these mappings and rollout plan:

* ZarishSphere is **FHIR-aligned**
* Vendor lock-in is avoided
* Donor & government compliance is achievable
* Platform growth is controlled and auditable

**This document is implementation-ready.**

---

