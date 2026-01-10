# Zarish Spreadsheet Standard

## Purpose

Defines the **single, universal spreadsheet format** used by NGOs
to create forms without writing code.

This spreadsheet is the **SOURCE OF TRUTH**.

---

## Sheet Name (Mandatory)

`form_definition`

---

## Column Definitions

| Column | Required | Description |
|------|----------|-------------|
| module | ✅ | ZarishSphere module name (EMR, HR, Finance) |
| domain | ✅ | Business domain (clinical, admin, finance) |
| form_id | ✅ | Unique form identifier |
| form_title | ✅ | Human-readable form name |
| field_name | ✅ | Database-safe field name |
| field_label | ✅ | UI label |
| field_type | ✅ | string, number, date, boolean, enum |
| required | ✅ | yes / no |
| default | ❌ | Default value |
| validation | ❌ | Validation rules |
| enum_values | ❌ | Pipe-separated values |
| ui_component | ❌ | input, select, textarea, datepicker |
| section | ❌ | UI grouping |
| order | ❌ | Display order |
| role_visibility | ❌ | Comma-separated roles |
| workflow_event | ❌ | Trigger name |
| audit | ❌ | audit level (strict / normal) |

---

## Rules

- One row = one field
- One spreadsheet = one form
- No macros
- No merged cells
- No formulas (values only)

---

## Example

```text
module: EMR
form_id: patient_registration
field_name: date_of_birth
field_type: date
required: yes