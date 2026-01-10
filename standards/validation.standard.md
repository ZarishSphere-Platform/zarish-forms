---

# 📁 `standards/validation.standard.md`

```markdown
# Zarish Validation Standard

## Purpose

Defines validation grammar for fields.

---

## Supported Rules

| Rule | Example |
|----|--------|
| min | min:0 |
| max | max:120 |
| regex | regex:^\\d+$ |
| length | length:5,20 |
| required | required |
| unique | unique |

---

## Validation Syntax

```text
min:0,max:120