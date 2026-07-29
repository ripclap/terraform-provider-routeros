package routeros

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testResourceSystemResourceIrq = "routeros_system_resource_irq.test_system_resource_irq"

// IRQ entries cannot be created or removed, so this test uses the singleton pattern without
// CheckDestroy and discovers a writable interrupt at runtime (IRQ numbers are hardware specific).
func TestAccSystemResourceIrqTest_basic(t *testing.T) {
	testCheckMenu(t, "/system/resource/irq")

	irq, ok := testFindWritableIrq(t)
	if !ok {
		t.Skip("no interrupt with read-only=false on this device, the CPU assignment cannot be changed")
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccSystemResourceIrqConfig(irq, "0"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceSystemResourceIrq),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrq, "irq", irq),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrq, "cpu", "0"),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrq, "read_only", "false"),
							resource.TestCheckResourceAttrSet(testResourceSystemResourceIrq, "users"),
							resource.TestCheckResourceAttrSet(testResourceSystemResourceIrq, "irq_count"),
							resource.TestCheckResourceAttrSet(testResourceSystemResourceIrq, "per_cpu_count"),
						),
					},
					{
						Config: testAccSystemResourceIrqConfig(irq, "auto"),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testResourceSystemResourceIrq),
							resource.TestCheckResourceAttr(testResourceSystemResourceIrq, "cpu", "auto"),
						),
					},
				},
			})
		})
	}
}

// testFindWritableIrq returns the number of the least busy interrupt whose CPU assignment the
// device allows to be changed (read-only=false).
func testFindWritableIrq(t *testing.T) (string, bool) {
	t.Helper()

	items, ok := testIrqRestItems(t, "/system/resource/irq")
	if !ok {
		return "", false
	}

	best := ""
	bestCount := -1
	for _, item := range items {
		if item["read-only"] != "false" {
			continue
		}
		irq, err := strconv.Atoi(item["irq"])
		if err != nil {
			continue
		}
		count, err := strconv.Atoi(item["count"])
		if err != nil {
			count = 0
		}
		if bestCount == -1 || count < bestCount {
			best, bestCount = strconv.Itoa(irq), count
		}
	}

	return best, best != ""
}

// testIrqRestItems fetches a menu over REST and returns its entries as string maps.
func testIrqRestItems(t *testing.T, path string) ([]map[string]string, bool) {
	t.Helper()

	host := reHost.FindStringSubmatch(origHostURL)
	if host == nil {
		return nil, false
	}
	port := ""
	if m := rePort.FindStringSubmatch(origHostURL); m != nil {
		port = ":" + m[1]
	}
	req, err := http.NewRequest("GET", "https://"+host[1]+port+"/rest"+path, nil)
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
	return items, true
}

func testAccSystemResourceIrqConfig(irq, cpu string) string {
	return providerConfig + `
resource "routeros_system_resource_irq" "test_system_resource_irq" {
	irq = ` + irq + `
	cpu = "` + cpu + `"
}
`
}
