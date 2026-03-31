# Contributing

Thank you for your interest in improving this project.

---

## How to Contribute

### Report a Bug

If a method returns an unexpected result, crashes, or behaves incorrectly:

1. Open an [issue](../../issues/new?template=bug_report.md)
2. Include:
   - The **method name** and SDK language (e.g. `spot.createLimitOrder` in TypeScript)
   - A **description of the expected vs. actual behavior**
   - The **error message or response** you received (redact credentials)
   - Your **environment**: OS, runtime version, SDK version
   - Steps to reproduce

### Request a Feature or New Endpoint

If you'd like a new method, endpoint, or behavior:

1. Open an [issue](../../issues/new?template=feature_request.md)
2. Describe:
   - What you want the method or feature to do
   - Which API endpoint it corresponds to (link to the API docs if applicable)
   - Your use case

### Report a Documentation Issue

If something in the README or docs is wrong, outdated, or unclear:

1. Open an [issue](../../issues/new?template=docs.md) describing what needs to be corrected
2. Include the section and a suggested correction if possible

### Ask a Question

For usage questions or integration help, open an [issue](../../issues/new) with the `question` label rather than sending email or a direct message. This keeps answers visible to others with the same question.

---

## What We Do Not Accept

We do **not** accept pull requests that:

- Manually add, remove, or rename individual methods without a corresponding API change
- Hardcode parameters that should be driven by the specification
- Modify method signatures or descriptions without an upstream change

---

## Issue Labels

| Label | Meaning |
|-------|---------|
| `bug` | Something is broken |
| `feature-request` | New method or capability |
| `docs` | Documentation fix or improvement |
| `question` | Usage or integration question |
| `upstream` | Root cause is in the upstream specification |
| `wontfix` | Out of scope for this repository |
