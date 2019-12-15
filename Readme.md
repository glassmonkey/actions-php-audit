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
      - uses: glassmonenkey/actions-php-audit@master
        id: checker # id is required if called from other steps
        with:
          WORK_DIR: /path/to/dir # Relative path of the directory where composer.lock exists
      # Todo other actions
      - name: sample message
        run: echo "${{ steps.checker.outputs.message }}"
```

## Locally Run
You can run with A sample `composer.lock`
```bash
docker-composer run app
```

## License
[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)
