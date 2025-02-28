# Contributing to MIDAI

First off, thanks for taking the time to contribute! ❤️

MIDAI is a Golang CLI tool for managing todos with iCloud integration. All types of contributions are encouraged and valued. See the [Table of Contents](#table-of-contents) for different ways to help and details about how this project handles them. Please make sure to read the relevant section before making your contribution. It will make it a lot easier for us maintainers and smooth out the experience for all involved. The community looks forward to your contributions. 🎉

> And if you like the project, but just don't have time to contribute, that's fine. There are other easy ways to support the project and show your appreciation, which we would also be very happy about:
>
> - Star the project
> - Tweet about it
> - Refer this project in your project's readme
> - Mention the project at local meetups and tell your friends/colleagues

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [I Have a Question](#i-have-a-question)
- [I Want To Contribute](#i-want-to-contribute)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)
- [Your First Code Contribution](#your-first-code-contribution)
- [Improving The Documentation](#improving-the-documentation)
- [Styleguides](#styleguides)
- [Commit Messages](#commit-messages)
- [Join The Project Team](#join-the-project-team)

## Code of Conduct

This project and everyone participating in it is governed by the
[MIDAI Code of Conduct](blob/master/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code. Please report unacceptable behavior
to the project maintainers.

## I Have a Question

> If you want to ask a question, we assume that you have read the available documentation.

Before you ask a question, it is best to search for existing [Issues](/issues) that might help you. In case you have found a suitable issue and still need clarification, you can write your question in this issue. It is also advisable to search the internet for answers first.

If you then still feel the need to ask a question and need clarification, we recommend the following:

- Open an [Issue](/issues/new).
- Provide as much context as you can about what you're running into.
- Provide project and platform versions (Go version, macOS/iOS version, etc), depending on what seems relevant.

We will then take care of the issue as soon as possible.

## I Want To Contribute

> ### Legal Notice
>
> When contributing to this project, you must agree that you have authored 100% of the content, that you have the necessary rights to the content and that the content you contribute may be provided under the project license.

### Reporting Bugs

#### Before Submitting a Bug Report

A good bug report shouldn't leave others needing to chase you up for more information. Therefore, we ask you to investigate carefully, collect information and describe the issue in detail in your report. Please complete the following steps in advance to help us fix any potential bug as fast as possible.

- Make sure that you are using the latest version.
- Determine if your bug is really a bug and not an error on your side e.g. using incompatible Go version or iCloud credentials issues.
- To see if other users have experienced (and potentially already solved) the same issue you are having, check if there is not already a bug report existing for your bug or error in the [bug tracker](issues?q=label%3Abug).
- Also make sure to search the internet (including Stack Overflow) to see if users outside of the GitHub community have discussed the issue.
- Collect information about the bug:
  - Stack trace (Traceback)
  - OS, Platform and Version (Windows, Linux, macOS, x86, ARM)
  - Version of Go, and any relevant packages
  - Your iCloud setup (without sharing sensitive credentials)
  - Possibly your input and the output
  - Can you reliably reproduce the issue? And can you also reproduce it with older versions?

#### How Do I Submit a Good Bug Report?

> You must never report security related issues, vulnerabilities or bugs including sensitive information to the issue tracker, or elsewhere in public. Instead sensitive bugs must be sent by email to the project maintainers.

We use GitHub issues to track bugs and errors. If you run into an issue with the project:

- Open an [Issue](/issues/new).
- Explain the behavior you would expect and the actual behavior.
- Please provide as much context as possible and describe the _reproduction steps_ that someone else can follow to recreate the issue on their own. This usually includes your code. For good bug reports you should isolate the problem and create a reduced test case.
- Provide the information you collected in the previous section.

Once it's filed:

- The project team will label the issue accordingly.
- A team member will try to reproduce the issue with your provided steps. If there are no reproduction steps or no obvious way to reproduce the issue, the team will ask you for those steps and mark the issue as `needs-repro`. Bugs with the `needs-repro` tag will not be addressed until they are reproduced.
- If the team is able to reproduce the issue, it will be marked `needs-fix`, as well as possibly other tags (such as `critical`), and the issue will be left to be [implemented by someone](#your-first-code-contribution).

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion for MIDAI, **including completely new features and minor improvements to existing functionality**. Following these guidelines will help maintainers and the community to understand your suggestion and find related suggestions.

#### Before Submitting an Enhancement

- Make sure that you are using the latest version.
- Read the documentation carefully and find out if the functionality is already covered, maybe by an individual configuration.
- Perform a [search](/issues) to see if the enhancement has already been suggested. If it has, add a comment to the existing issue instead of opening a new one.
- Find out whether your idea fits with the scope and aims of the project. It's up to you to make a strong case to convince the project's developers of the merits of this feature. Keep in mind that we want features that will be useful to the majority of our users and not just a small subset.

#### How Do I Submit a Good Enhancement Suggestion?

Enhancement suggestions are tracked as [GitHub issues](/issues).

- Use a **clear and descriptive title** for the issue to identify the suggestion.
- Provide a **step-by-step description of the suggested enhancement** in as many details as possible.
- **Describe the current behavior** and **explain which behavior you expected to see instead** and why. At this point you can also tell which alternatives do not work for you.
- You may want to **include screenshots and animated GIFs** which help you demonstrate the steps or point out the part which the suggestion is related to.
- **Explain why this enhancement would be useful** to most MIDAI users. You may also want to point out the other projects that solved it better and which could serve as inspiration.

### Your First Code Contribution

Unsure where to begin contributing to MIDAI? You can start by looking through these `beginner` and `help-wanted` issues:

- [Beginner issues](/issues?q=label%3Abeginner) - issues which should only require a few lines of code, and a test or two.
- [Help wanted issues](/issues?q=label%3A"help+wanted") - issues which should be a bit more involved than `beginner` issues.

#### Local Development

To set up MIDAI for local development:

1. Fork the repository on GitHub
2. Clone your fork locally
3. Install Go (version 1.16 or higher recommended)
4. Run `go mod download` to install dependencies
5. Create a branch for your changes
6. Make your changes
7. Run tests with `go test ./...`
8. Push your branch and submit a pull request

### Improving The Documentation

MIDAI's documentation is as important as its code. Good documentation helps users understand how to use the tool effectively and helps new contributors get started. Documentation contributions can include:

- Improving the README.md
- Adding or updating CLI command documentation
- Creating tutorials or guides for features like iCloud integration
- Adding code comments that make the codebase more accessible

## Styleguides

### Go Code Styleguide

- Follow standard Go formatting with `gofmt`
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Document all exported functions, types, and methods according to [godoc](https://blog.golang.org/godoc-documenting-go-code) conventions
- Use meaningful variable names
- Group related declarations
- Organize imports alphabetically

### Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line
- Consider starting the commit message with an applicable emoji:
  - ✨ `:sparkles:` when adding a new feature
  - 🐛 `:bug:` when fixing a bug
  - 📚 `:books:` when adding or updating documentation
  - 🧹 `:broom:` when refactoring code
  - 🧪 `:test_tube:` when adding tests
