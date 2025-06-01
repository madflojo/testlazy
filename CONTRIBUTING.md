# Contributing to go-monorepo

Thank you for considering contributing to this project! Your time, skills, and perspective are valued.

## Issues and Proposals

Bugs, proposals, and feature requests are all welcome. To get started, please open an issue on GitHub and provide as much detail as possible.

## Commit Messages

This project follows the [Conventional Commits](https://www.conventionalcommits.org/) specification. All commit messages should be structured as follows:

    <type>[optional scope]: <description>

    [optional body]

    [optional footer(s)]

Types include:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (white-space, formatting, etc.)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `build`: Changes that affect the build system or external dependencies
- `ci`: Changes to CI configuration and scripts
- `chore`: Other changes that do not modify src or test files

Examples:
    feat(modulea): add support for feature X
    fix(moduleb): resolve nil pointer in handler
    docs: update README with setup steps

Commits following this convention will be automatically included in the changelog when a release is created.