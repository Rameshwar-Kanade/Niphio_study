# packages

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] packages`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree packages`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init packages
kpt live apply packages --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
