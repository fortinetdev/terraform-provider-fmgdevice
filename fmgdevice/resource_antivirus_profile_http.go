// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure HTTP AntiVirus options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusProfileHttp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusProfileHttpUpdate,
		Read:   resourceAntivirusProfileHttpRead,
		Update: resourceAntivirusProfileHttpUpdate,
		Delete: resourceAntivirusProfileHttpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"archive_block": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"archive_log": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"av_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content_disarm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"emulator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"external_blocklist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiai": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortindr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortisandbox": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"malware_stream": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbreak_prevention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_content_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAntivirusProfileHttpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	obj, err := getObjectAntivirusProfileHttp(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileHttp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAntivirusProfileHttp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileHttp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AntivirusProfileHttp")

	return resourceAntivirusProfileHttpRead(d, m)
}

func resourceAntivirusProfileHttpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	profile := d.Get("profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteAntivirusProfileHttp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusProfileHttp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusProfileHttpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile := d.Get("profile").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if device_vdom == "" {
		device_vdom = importOptionChecking(m.(*FortiClient).Cfg, "device_vdom")
		if device_vdom == "" {
			return fmt.Errorf("Parameter device_vdom is missing")
		}
		if err = d.Set("device_vdom", device_vdom); err != nil {
			return fmt.Errorf("Error set params device_vdom: %v", err)
		}
	}
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile"] = profile

	o, err := c.ReadAntivirusProfileHttp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AntivirusProfileHttp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusProfileHttp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusProfileHttp resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusProfileHttpArchiveBlock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileHttpArchiveLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileHttpAvScan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpContentDisarm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpEmulator2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileHttpExternalBlocklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpFortiai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpFortindr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpFortisandbox2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpMalwareStream2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpOutbreakPrevention2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpUnknownContentEncoding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusProfileHttp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("archive_block", flattenAntivirusProfileHttpArchiveBlock2edl(o["archive-block"], d, "archive_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive-block"], "AntivirusProfileHttp-ArchiveBlock"); ok {
			if err = d.Set("archive_block", vv); err != nil {
				return fmt.Errorf("Error reading archive_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive_block: %v", err)
		}
	}

	if err = d.Set("archive_log", flattenAntivirusProfileHttpArchiveLog2edl(o["archive-log"], d, "archive_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive-log"], "AntivirusProfileHttp-ArchiveLog"); ok {
			if err = d.Set("archive_log", vv); err != nil {
				return fmt.Errorf("Error reading archive_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive_log: %v", err)
		}
	}

	if err = d.Set("av_scan", flattenAntivirusProfileHttpAvScan2edl(o["av-scan"], d, "av_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-scan"], "AntivirusProfileHttp-AvScan"); ok {
			if err = d.Set("av_scan", vv); err != nil {
				return fmt.Errorf("Error reading av_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_scan: %v", err)
		}
	}

	if err = d.Set("content_disarm", flattenAntivirusProfileHttpContentDisarm2edl(o["content-disarm"], d, "content_disarm")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-disarm"], "AntivirusProfileHttp-ContentDisarm"); ok {
			if err = d.Set("content_disarm", vv); err != nil {
				return fmt.Errorf("Error reading content_disarm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_disarm: %v", err)
		}
	}

	if err = d.Set("emulator", flattenAntivirusProfileHttpEmulator2edl(o["emulator"], d, "emulator")); err != nil {
		if vv, ok := fortiAPIPatch(o["emulator"], "AntivirusProfileHttp-Emulator"); ok {
			if err = d.Set("emulator", vv); err != nil {
				return fmt.Errorf("Error reading emulator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emulator: %v", err)
		}
	}

	if err = d.Set("options", flattenAntivirusProfileHttpOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "AntivirusProfileHttp-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("external_blocklist", flattenAntivirusProfileHttpExternalBlocklist2edl(o["external-blocklist"], d, "external_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist"], "AntivirusProfileHttp-ExternalBlocklist"); ok {
			if err = d.Set("external_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist: %v", err)
		}
	}

	if err = d.Set("fortiai", flattenAntivirusProfileHttpFortiai2edl(o["fortiai"], d, "fortiai")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai"], "AntivirusProfileHttp-Fortiai"); ok {
			if err = d.Set("fortiai", vv); err != nil {
				return fmt.Errorf("Error reading fortiai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai: %v", err)
		}
	}

	if err = d.Set("fortindr", flattenAntivirusProfileHttpFortindr2edl(o["fortindr"], d, "fortindr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr"], "AntivirusProfileHttp-Fortindr"); ok {
			if err = d.Set("fortindr", vv); err != nil {
				return fmt.Errorf("Error reading fortindr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr: %v", err)
		}
	}

	if err = d.Set("fortisandbox", flattenAntivirusProfileHttpFortisandbox2edl(o["fortisandbox"], d, "fortisandbox")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox"], "AntivirusProfileHttp-Fortisandbox"); ok {
			if err = d.Set("fortisandbox", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox: %v", err)
		}
	}

	if err = d.Set("malware_stream", flattenAntivirusProfileHttpMalwareStream2edl(o["malware-stream"], d, "malware_stream")); err != nil {
		if vv, ok := fortiAPIPatch(o["malware-stream"], "AntivirusProfileHttp-MalwareStream"); ok {
			if err = d.Set("malware_stream", vv); err != nil {
				return fmt.Errorf("Error reading malware_stream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malware_stream: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention", flattenAntivirusProfileHttpOutbreakPrevention2edl(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "AntivirusProfileHttp-OutbreakPrevention"); ok {
			if err = d.Set("outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenAntivirusProfileHttpQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "AntivirusProfileHttp-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("unknown_content_encoding", flattenAntivirusProfileHttpUnknownContentEncoding2edl(o["unknown-content-encoding"], d, "unknown_content_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["unknown-content-encoding"], "AntivirusProfileHttp-UnknownContentEncoding"); ok {
			if err = d.Set("unknown_content_encoding", vv); err != nil {
				return fmt.Errorf("Error reading unknown_content_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading unknown_content_encoding: %v", err)
		}
	}

	return nil
}

func flattenAntivirusProfileHttpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusProfileHttpArchiveBlock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileHttpArchiveLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileHttpAvScan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpContentDisarm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpEmulator2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileHttpExternalBlocklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpFortiai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpFortindr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpFortisandbox2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpMalwareStream2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpOutbreakPrevention2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpUnknownContentEncoding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusProfileHttp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("archive_block"); ok || d.HasChange("archive_block") {
		t, err := expandAntivirusProfileHttpArchiveBlock2edl(d, v, "archive_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive-block"] = t
		}
	}

	if v, ok := d.GetOk("archive_log"); ok || d.HasChange("archive_log") {
		t, err := expandAntivirusProfileHttpArchiveLog2edl(d, v, "archive_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive-log"] = t
		}
	}

	if v, ok := d.GetOk("av_scan"); ok || d.HasChange("av_scan") {
		t, err := expandAntivirusProfileHttpAvScan2edl(d, v, "av_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-scan"] = t
		}
	}

	if v, ok := d.GetOk("content_disarm"); ok || d.HasChange("content_disarm") {
		t, err := expandAntivirusProfileHttpContentDisarm2edl(d, v, "content_disarm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-disarm"] = t
		}
	}

	if v, ok := d.GetOk("emulator"); ok || d.HasChange("emulator") {
		t, err := expandAntivirusProfileHttpEmulator2edl(d, v, "emulator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emulator"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandAntivirusProfileHttpOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist"); ok || d.HasChange("external_blocklist") {
		t, err := expandAntivirusProfileHttpExternalBlocklist2edl(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("fortiai"); ok || d.HasChange("fortiai") {
		t, err := expandAntivirusProfileHttpFortiai2edl(d, v, "fortiai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai"] = t
		}
	}

	if v, ok := d.GetOk("fortindr"); ok || d.HasChange("fortindr") {
		t, err := expandAntivirusProfileHttpFortindr2edl(d, v, "fortindr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox"); ok || d.HasChange("fortisandbox") {
		t, err := expandAntivirusProfileHttpFortisandbox2edl(d, v, "fortisandbox")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox"] = t
		}
	}

	if v, ok := d.GetOk("malware_stream"); ok || d.HasChange("malware_stream") {
		t, err := expandAntivirusProfileHttpMalwareStream2edl(d, v, "malware_stream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malware-stream"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok || d.HasChange("outbreak_prevention") {
		t, err := expandAntivirusProfileHttpOutbreakPrevention2edl(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandAntivirusProfileHttpQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("unknown_content_encoding"); ok || d.HasChange("unknown_content_encoding") {
		t, err := expandAntivirusProfileHttpUnknownContentEncoding2edl(d, v, "unknown_content_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-content-encoding"] = t
		}
	}

	return &obj, nil
}
