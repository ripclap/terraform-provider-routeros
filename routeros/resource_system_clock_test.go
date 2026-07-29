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

const testSystemClockTask = "routeros_system_clock.set"

// testSystemClockTimeLayout is the format of the /system/clock `time` field.
const testSystemClockTimeLayout = "15:04:05"

// testSystemClockNudge is how far the clock is moved to prove the write reached the device: far larger than a
// request round trip so the read-back is not the clock's natural progression, yet small enough not to disturb
// a device left in this state by a failing step.
const testSystemClockNudge = 5 * time.Minute

// testSystemClockSkew is the tolerance of the read-back checks: the wall clock keeps running while
// the provider writes the value and reads it back again.
const testSystemClockSkew = time.Minute

func TestAccSystemClockTest_basic(t *testing.T) {
	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			steps := makeSteps(t, name)
			if steps == nil {
				return
			}

			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps:             steps,
			})

		})
	}
}

// makeSteps builds one step per writable attribute of /system/clock. The date and time are derived from the
// device, not hardcoded: RouterOS refuses a timestamp older than the package build time ("cannot set time
// before package build time"). The step order is fixed (time zone, date, time) because changing the zone
// shifts the wall clock.
func makeSteps(t *testing.T, name string) []resource.TestStep {
	t.Helper()

	// UTC is used for both transports so the step exercises the time-zone-name write and read-back without moving the clock.
	const timeZone = "UTC"

	clock, ok := testSystemClockFetch(t)
	if !ok {
		t.Skip("cannot read /system/clock from the device under test")
		return nil
	}

	date, baseline := clock["date"], clock["time"]
	if date == "" || baseline == "" {
		t.Skipf("/system/clock did not report a date and a time: %v", clock)
		return nil
	}

	baselineTime, err := time.Parse(testSystemClockTimeLayout, baseline)
	if err != nil {
		t.Skipf("the device reported the time as %q, which is not %v: %v",
			baseline, testSystemClockTimeLayout, err)
		return nil
	}
	nudged := baselineTime.Add(testSystemClockNudge).Format(testSystemClockTimeLayout)

	step := func(k, v string, extra ...resource.TestCheckFunc) resource.TestStep {
		checks := []resource.TestCheckFunc{
			testResourcePrimaryInstanceId(testSystemClockTask),
		}
		checks = append(checks, extra...)

		return resource.TestStep{
			Config: fmt.Sprintf(`%v
			resource "routeros_system_clock" "set" {
				%v = "%v"
			}`, providerConfig, k, v),
			Check: resource.ComposeTestCheckFunc(checks...),
		}
	}

	return []resource.TestStep{
		step("time_zone_name", timeZone,
			resource.TestCheckResourceAttr(testSystemClockTask, "time_zone_name", timeZone)),
		step("date", date,
			resource.TestCheckResourceAttr(testSystemClockTask, "date", date)),
		// A running clock cannot be asserted by equality, and setting it to its current time would pass even if
		// nothing was written; it is moved forward by a fixed amount so the read-back proves the write landed, then restored.
		step("time", nudged,
			testCheckSystemClockTime(testSystemClockTask, nudged)),
		step("time", baseline,
			testCheckSystemClockTime(testSystemClockTask, baseline)),
	}
}

// testCheckSystemClockTime asserts the clock in the state is the applied timestamp, give or take
// the time the round trip to the device takes.
func testCheckSystemClockTime(address, applied string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[address]
		if !ok {
			return fmt.Errorf("resource %v not found in the state", address)
		}

		got := rs.Primary.Attributes["time"]
		gotTime, err := time.Parse(testSystemClockTimeLayout, got)
		if err != nil {
			return fmt.Errorf("the time reported by the device (%q) is not %v: %v",
				got, testSystemClockTimeLayout, err)
		}
		appliedTime, err := time.Parse(testSystemClockTimeLayout, applied)
		if err != nil {
			return fmt.Errorf("the applied time (%q) is not %v: %v",
				applied, testSystemClockTimeLayout, err)
		}

		delta := gotTime.Sub(appliedTime)
		if delta < 0 {
			// Midnight rollover between the write and the read-back.
			delta += 24 * time.Hour
		}
		if delta > testSystemClockSkew {
			return fmt.Errorf("the device clock reads %v, %v away from the applied time %v: "+
				"the value does not appear to have been written", got, delta, applied)
		}

		return nil
	}
}

// testSystemClockFetch reads /system/clock over REST and returns it as a string map.
func testSystemClockFetch(t *testing.T) (map[string]string, bool) {
	t.Helper()

	host := reHost.FindStringSubmatch(origHostURL)
	if host == nil {
		return nil, false
	}
	port := ""
	if m := rePort.FindStringSubmatch(origHostURL); m != nil {
		port = ":" + m[1]
	}
	req, err := http.NewRequest("GET", "https://"+host[1]+port+"/rest/system/clock", nil)
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
	var clock map[string]string
	if json.Unmarshal(body, &clock) != nil {
		return nil, false
	}
	return clock, true
}
