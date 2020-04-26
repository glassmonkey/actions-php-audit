package main

import (
	"testing"
)

func TestMessage(t *testing.T) {
	message := "Symfony Security Check Report ============================= 1 packages have known vulnerabilities. laravel/framework (v5.4.36) --------------------------- composer.lock [CVE-NONE-0001][]: Exploit of encryption failure vulnerability composer.lock [CVE-NONE-0002][]: Cookie serialization vulnerability composer.lock [CVE-2017-14775][]: Timing attack vector for remember me token [CVE-NONE-0001]: https://medium.com/@taylorotwell/laravel-security-release-5-6-15-and-5-5-40-56f1257933a0 [CVE-NONE-0002]: https://laravel.com/docs/5.6/upgrade#upgrade-5.6.30 [CVE-2017-14775]: https://github.com/laravel/framework/pull/21320 Note that this checker can only detect vulnerabilities that are referenced in the SensioLabs security advisories database. Execute this command regularly to check the newly discovered vulnerabilities."
	result := formatGithubMessage(message)
	expect := "# PHP Security Check Report\\n=============================\\n\\n- [ ] [CVE-NONE-0001][]: Exploit of encryption failure vulnerability composer.lock \\n- [ ] [CVE-NONE-0002][]: Cookie serialization vulnerability composer.lock \\n- [ ] [CVE-2017-14775][]: Timing attack vector for remember me token \\n- [ ] [CVE-NONE-0001]: https://medium.com/@taylorotwell/laravel-security-release-5-6-15-and-5-5-40-56f1257933a0 \\n- [ ] [CVE-NONE-0002]: https://laravel.com/docs/5.6/upgrade#upgrade-5.6.30 \\n- [ ] [CVE-2017-14775]: https://github.com/laravel/framework/pull/21320 Note that this checker can only detect vulnerabilities that are referenced in the SensioLabs security advisories database. Execute this command regularly to check the newly discovered vulnerabilities."
	if result != expect {
		t.Fatal("failed test")
	}
}

func TestNoMessage(t *testing.T) {
	message := "Symfony Security Check Report ============================= 1 packages have known vulnerabilities. laravel/framework (v5.4.36) --------------------------- composer.lock"
	result := formatGithubMessage(message)
	if result != "" {
		t.Fatal("failed test")
	}
}

func TestParseBody(t *testing.T) {
	message := "Symfony Security Check Report ============================= 1 packages have known vulnerabilities. laravel/framework (v5.4.36) --------------------------- composer.lock [CVE-NONE-0001][]: Exploit of encryption failure vulnerability composer.lock [CVE-NONE-0002][]: Cookie serialization vulnerability composer.lock [CVE-2017-14775][]: Timing attack vector for remember me token [CVE-NONE-0001]: https://medium.com/@taylorotwell/laravel-security-release-5-6-15-and-5-5-40-56f1257933a0 [CVE-NONE-0002]: https://laravel.com/docs/5.6/upgrade#upgrade-5.6.30 [CVE-2017-14775]: https://github.com/laravel/framework/pull/21320 Note that this checker can only detect vulnerabilities that are referenced in the SensioLabs security advisories database. Execute this command regularly to check the newly discovered vulnerabilities."
	expect := "\\n- [ ] [CVE-NONE-0001][]: Exploit of encryption failure vulnerability composer.lock \\n- [ ] [CVE-NONE-0002][]: Cookie serialization vulnerability composer.lock \\n- [ ] [CVE-2017-14775][]: Timing attack vector for remember me token \\n- [ ] [CVE-NONE-0001]: https://medium.com/@taylorotwell/laravel-security-release-5-6-15-and-5-5-40-56f1257933a0 \\n- [ ] [CVE-NONE-0002]: https://laravel.com/docs/5.6/upgrade#upgrade-5.6.30 \\n- [ ] [CVE-2017-14775]: https://github.com/laravel/framework/pull/21320 Note that this checker can only detect vulnerabilities that are referenced in the SensioLabs security advisories database. Execute this command regularly to check the newly discovered vulnerabilities."
	result, _ := selectMessageBody(message)
	if result != expect {
		t.Fatal("failed test")
	}
}
