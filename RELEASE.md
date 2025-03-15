# Release Process

This document outlines the process for creating and publishing new releases of Commit-AI.

## Versioning

Commit-AI follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html). Given a version number MAJOR.MINOR.PATCH:

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes

## Release Checklist

Before creating a new release, ensure the following:

1. All tests pass successfully
2. Documentation is updated to reflect any changes
3. CHANGELOG.md is updated with all notable changes
4. Version numbers are updated in relevant files:
   - main.go (the `version` constant)
   - DOCUMENTATION.md

## Creating a Release

### 1. Update Changelog

Update the CHANGELOG.md file with the new version and all changes, following the Keep a Changelog format.

### 2. Update Version Number

Update the version number in the relevant files mentioned in the checklist.

### 3. Commit Changes

Commit these changes with a message like:

```bash
commit-ai --language=en
# This should generate something like: "chore: prepare for vX.Y.Z release"
```

### 4. Create a Tag

Create a Git tag for the new version:

```bash
git tag -a vX.Y.Z -m "Release vX.Y.Z"
```

### 5. Push Changes and Tags

Push the changes and tags to the repository:

```bash
git push origin main
git push origin vX.Y.Z
```

### 6. Build Release Binaries

Build binaries for all supported platforms:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/commit-ai-linux-amd64

# macOS
GOOS=darwin GOARCH=amd64 go build -o bin/commit-ai-darwin-amd64

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/commit-ai-windows-amd64.exe
```

### 7. Generate Checksums

Generate SHA-256 checksums for all binaries:

```bash
cd bin && sha256sum commit-ai-* > checksums.txt
```

### 8. Create a GitHub Release

1. Go to the repository on GitHub
2. Click on "Releases"
3. Click on "Create a new release"
4. Select the tag version
5. Add a title and description (you can copy from the CHANGELOG)
6. Upload the binaries and checksums
7. Publish the release

## Hotfix Releases

For urgent fixes to the latest release:

1. Create a branch from the latest tag
2. Apply the fix
3. Follow the regular release process, incrementing only the PATCH version 