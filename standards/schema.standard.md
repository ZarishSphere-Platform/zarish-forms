---

# 📁 `standards/schema.standard.md`

```markdown
# Zarish Form Schema Standard

## Purpose

Defines the **canonical JSON schema** generated from spreadsheets.

This schema is the **TRUTH**.

---

## Required Top-Level Fields

- schema_version
- module
- domain
- form_id
- form_title
- fields[]

---

## Field Object Structure

```json
{
  "name": "string",
  "label": "string",
  "type": "string",
  "required": true,
  "default": null,
  "validation": {},
  "ui": {},
  "access": {},
  "workflow": {}
}