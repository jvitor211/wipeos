# Contributing to WipeOs

We love your input! We want to make contributing to WipeOs as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

### Pull Requests

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

### Development Setup

1. **Prerequisites**
   - Go 1.22 or higher
   - Git
   - Make (optional, but recommended)

2. **Clone and setup**
   ```bash
   git clone https://github.com/joao-rrondon/wipeOs.git
   cd wipeOs
   go mod download
   ```

3. **Install development tools**
   ```bash
   make install-tools
   ```

4. **Run tests**
   ```bash
   make test
   ```

5. **Run linter**
   ```bash
   make lint
   ```

### Code Style

- We use `gofmt` and `goimports` for code formatting
- Follow Go best practices and idioms
- Write clear, concise comments
- Use meaningful variable and function names

### Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting PR
- Aim for good test coverage
- Use table-driven tests where appropriate

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

### Security Considerations

Since WipeOs deals with secure file deletion:

- Be extremely careful with file operations
- Always validate file paths and permissions
- Test thoroughly on different operating systems
- Consider edge cases (symlinks, special files, etc.)
- Never commit test files with sensitive content

## Bug Reports

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/joao-rrondon/wipeOs/issues).

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

## Feature Requests

We welcome feature requests! Please:

1. Check if the feature has already been requested
2. Clearly describe the feature and its use case
3. Consider if it aligns with the project's goals
4. Be willing to help implement it

## Commit Message Format

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

Examples:
```
feat: add browser data cleaning for Safari
fix: prevent panic when file doesn't exist
docs: update installation instructions
```

## Code of Conduct

### Our Pledge

We pledge to make participation in our project a harassment-free experience for everyone, regardless of age, body size, disability, ethnicity, gender identity and expression, level of experience, nationality, personal appearance, race, religion, or sexual identity and orientation.

### Our Standards

Examples of behavior that contributes to creating a positive environment include:

- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

### Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported by contacting the project team. All complaints will be reviewed and investigated and will result in a response that is deemed necessary and appropriate to the circumstances.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## Questions?

Don't hesitate to ask questions by:
- Opening an issue
- Starting a discussion
- Reaching out to maintainers

Thank you for contributing to WipeOs! 🎉 