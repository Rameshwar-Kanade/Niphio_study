# my-kpt-package

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] my-kpt-package`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree my-kpt-package`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init my-kpt-package
kpt live apply my-kpt-package --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
