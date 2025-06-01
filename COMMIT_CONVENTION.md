# Commit Convention

This project follows the [Conventional Commits](https://www.conventionalcommits.org/) specification.

## Format

Each commit message consists of a **header**, an optional **body**, and an optional **footer**.

```text
<type>(<optional scope>): <subject>

<optional body>

<optional footer>
```

### Commit Message Header

```text
<type>(<optional scope>): <subject>
```

The `type` and `subject` fields are mandatory; the `scope` field is optional.

### Type

Must be one of the following:

- **feat**: A new feature
- **fix**: A bug fix
- **docs**: Documentation only changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, etc.)
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **perf**: A code change that improves performance
- **test**: Adding missing tests or correcting existing tests
- **build**: Changes that affect the build system or external dependencies
- **ci**: Changes to CI configuration and scripts
- **chore**: Other changes that do not modify src or test files

### Scope

The scope specifies what area of the codebase your commit touches, and is usually the module or package name. For example:

- **modulea**: Changes to the modulea package
- **moduleb**: Changes to the moduleb package

### Subject

The subject is a succinct description of the change:

- Use the imperative, present tense: "change" not "changed" nor "changes"
- Do not capitalize the first letter
- Do not end with a period (.)

### Body

The body should include the motivation for the change and contrast this with previous behavior.

### Footer

The footer should contain any information about **Breaking Changes** and is also the place to reference GitHub issues closed by the change.

**Breaking Changes** should start with the word `BREAKING CHANGE:` followed by a description.

## Examples

```text
feat(modulea): add support for feature X

Implements the new X feature with backward compatibility.

Closes #123
```

```text
fix(moduleb): resolve nil pointer in handler

Ensure the handler checks for nil before use to avoid panics.

BREAKING CHANGE: The handler now returns an error when input is nil.
```

```text
docs: update contributing guidelines
```

```text
refactor: simplify configuration parsing logic
```