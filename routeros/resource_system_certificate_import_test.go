package routeros

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testSystemCertificatesImportAddress = "routeros_system_certificate.external"

func TestAccSystemCertificatesTest_import(t *testing.T) {
	// ROS 7.12.x does not return IDs for created files.
	if !testCheckMinVersion(t, testFileMinVersion) {
		t.Logf("Test skipped, the minimum required version is %v", testFileMinVersion)
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			var externalCrt MikrotikItem

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				CheckDestroy:      testCheckResourceDestroy("/certificate", "routeros_system_certificate"),
				Steps: []resource.TestStep{
					{
						Config: testAccSystemCertificatesImportConfig(),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemCertificatesImportAddress),
							testCheckResourceExists("routeros_system_certificate.external", "/certificate", &externalCrt),
							testCheckMikrotikItemAttr("routeros_system_certificate.external", &externalCrt, "name", "external.crt"),
							resource.TestCheckResourceAttr("routeros_system_certificate.external", "name", "external.crt"),
							resource.TestCheckResourceAttr("routeros_system_certificate.external", "common_name", "External Certificate"),
							resource.TestCheckResourceAttr("routeros_system_certificate.external", "private_key", "true"),
							// The import consumes the files it reads, see below.
							testCheckImportSourceFilesConsumed("external.crt", "external.key"),
						),
						// /certificate/import moves the certificate and key out of /file into the certificate store, so the
						// two routeros_file resources no longer exist and Terraform expectedly plans them for re-creation.
						ExpectNonEmptyPlan: true,
					},
				},
			})

		})
	}
}

// testCheckImportSourceFilesConsumed asserts that RouterOS removed the given files from /file,
// which is what it does with the source files of a successful /certificate/import.
func testCheckImportSourceFilesConsumed(names ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		present, ok := testCertificateFileNames()
		if !ok {
			return fmt.Errorf("cannot read /file from the device under test")
		}

		for _, name := range names {
			if present[name] {
				return fmt.Errorf("%v is still present in /file, the certificate import did not "+
					"consume it", name)
			}
		}

		return nil
	}
}

// testCertificateFileNames returns the set of names in /file. Only the name field is requested:
// the contents of a binary file would not survive being decoded as a JSON string.
func testCertificateFileNames() (map[string]bool, bool) {
	host := reHost.FindStringSubmatch(origHostURL)
	if host == nil {
		return nil, false
	}
	port := ""
	if m := rePort.FindStringSubmatch(origHostURL); m != nil {
		port = ":" + m[1]
	}
	req, err := http.NewRequest("GET", "https://"+host[1]+port+"/rest/file?.proplist=name", nil)
	if err != nil {
		return nil, false
	}
	req.SetBasicAuth(os.Getenv("ROS_USERNAME"), os.Getenv("ROS_PASSWORD"))
	cl := &http.Client{
		Timeout:   15 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	res, err := cl.Do(req)
	if err != nil {
		return nil, false
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, false
	}
	var items []map[string]string
	if json.Unmarshal(body, &items) != nil {
		return nil, false
	}

	names := make(map[string]bool, len(items))
	for _, item := range items {
		names[item["name"]] = true
	}
	return names, true
}

func testAccSystemCertificatesImportConfig() string {
	return providerConfig + `
data "routeros_x509" "cert" {
	data = <<EOT
	-----BEGIN CERTIFICATE-----
	MIIBlTCCATugAwIBAgIINLsws71B5zIwCgYIKoZIzj0EAwIwHzEdMBsGA1UEAwwU
	RXh0ZXJuYWwgQ2VydGlmaWNhdGUwHhcNMjQwNTE3MjEyOTUzWhcNMjUwNTE3MjEy
	OTUzWjAfMR0wGwYDVQQDDBRFeHRlcm5hbCBDZXJ0aWZpY2F0ZTBZMBMGByqGSM49
	AgEGCCqGSM49AwEHA0IABKE1g0Qj4ujIold9tklu2z4BUu/K7xDFF5YmedtOfJyM
	1/80APNboqn71y4m4XNE1JNtQuR2bSZPHVrzODkR16ujYTBfMA8GA1UdEwEB/wQF
	MAMBAf8wDgYDVR0PAQH/BAQDAgG2MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEF
	BQcDAjAdBgNVHQ4EFgQUNXd5bvluIV9YAhGc5yMHc6OzXpMwCgYIKoZIzj0EAwID
	SAAwRQIhAODte/qS6CE30cvnQpxP/ObWBPIPZnHtkFHIIC1AOSXwAiBGCGQE+aJY
	W72Rw0Y1ckvlt6sU0urkzGuj5wxVF/gSYA==
	-----END CERTIFICATE-----
EOT
}

resource "routeros_file" "key" {
	name     = "external.key"
	contents = <<EOT
-----BEGIN ENCRYPTED PRIVATE KEY-----
MIHeMEkGCSqGSIb3DQEFDTA8MBsGCSqGSIb3DQEFDDAOBAiy/wEW6/MglgICCAAw
HQYJYIZIAWUDBAEqBBD6v8dLA2FjPn62Xz57pcu9BIGQhclivPw1eC2b14ea58Tw
nzDdbYN6/yUiMqapW2xZaT7ZFnbEai4n9/utgtEDnfKHlZvZj2kRhvYoWrvTkt/W
1mkd5d/runsn+B5GO+CMHFHh4t41WMpZysmg+iP8FiiehOQEsWyEZFaedxfYYtSL
Sk+abxJ+NMQoh+S5d73niu1CO8uqQjOd8BoSOurURsOh
-----END ENCRYPTED PRIVATE KEY-----
EOT
}

resource "routeros_file" "cert" {
	name     = "external.crt"
	contents = data.routeros_x509.cert.pem
}

resource "routeros_system_certificate" "external" {
	name        = "external.crt"
	common_name = data.routeros_x509.cert.common_name
	import {
		cert_file_name  = routeros_file.cert.name
		key_file_name   = routeros_file.key.name
		passphrase      = "11111111"
	}
	depends_on = [routeros_file.key, routeros_file.cert]
}
`
}
