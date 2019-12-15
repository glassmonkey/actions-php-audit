#  Github Action for PHP Audit
Using [security-checker](https://github.com/sensiolabs/security-checker) on Github Actions

## Usage

.github/workflows/checker.yml
```yaml
name: Alert Composer Audit
on:
  push:
    branches:
      - master
jobs:
  check-validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v1
      - uses: ./
        name: checker
        id: checker
      # Todo other actions
      - name: sample message
        run: echo "${{ steps.checker.outputs.message }}"
```