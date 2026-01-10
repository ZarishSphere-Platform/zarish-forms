# Zarish Forms Security

Zarish Forms is designed to **define forms safely** and **avoid execution risks**.

## Principles
- No direct data storage
- No running workflows
- Schema validation before every commit
- Immutable audit history
- GitHub Actions CI ensures validation

## Security Practices
- Use Go's strong typing and compile-time checks
- Validate all spreadsheet inputs
- Generate UI previews for human inspection
- Version all schemas
- Restrict branch permissions

## Incident Reporting
- Report security issues via GitHub Security Advisories
- Follow ZarishSphere security response process
