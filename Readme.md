#  Github Action for PHP Audit
Using [security-checker](https://github.com/sensiolabs/security-checker) on Github Actions  
This Action report `PHP Security infomation` as a issue.   
e.g. https://github.com/glassmonkey/actions-php-audit/issues/4

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
      - uses: glassmonkey/actions-php-audit@v1
        with:
          WORK_DIR: path/to/dir # Relative path of the directory where composer.lock exists
          ISSUE_LABELS: bugs # Apply labels to a issue. 
```

## Locally Run
You can run with A sample `composer.lock`
```bash
docker-composer run app
```

## License
[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)
