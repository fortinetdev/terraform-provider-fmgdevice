// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure FTP AntiVirus options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusProfileFtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusProfileFtpUpdate,
		Read:   resourceAntivirusProfileFtpRead,
		Update: resourceAntivirusProfileFtpUpdate,
		Delete: resourceAntivirusProfileFtpDelete,

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
		},
	}
}

func resourceAntivirusProfileFtpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAntivirusProfileFtp(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileFtp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAntivirusProfileFtp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileFtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AntivirusProfileFtp")

	return resourceAntivirusProfileFtpRead(d, m)
}

func resourceAntivirusProfileFtpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAntivirusProfileFtp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusProfileFtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusProfileFtpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAntivirusProfileFtp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AntivirusProfileFtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusProfileFtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusProfileFtp resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusProfileFtpArchiveBlock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtpArchiveLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtpAvScan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpEmulator2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtpExternalBlocklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpFortiai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpFortindr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpFortisandbox2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpMalwareStream2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpOutbreakPrevention2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusProfileFtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("archive_block", flattenAntivirusProfileFtpArchiveBlock2edl(o["archive-block"], d, "archive_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive-block"], "AntivirusProfileFtp-ArchiveBlock"); ok {
			if err = d.Set("archive_block", vv); err != nil {
				return fmt.Errorf("Error reading archive_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive_block: %v", err)
		}
	}

	if err = d.Set("archive_log", flattenAntivirusProfileFtpArchiveLog2edl(o["archive-log"], d, "archive_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive-log"], "AntivirusProfileFtp-ArchiveLog"); ok {
			if err = d.Set("archive_log", vv); err != nil {
				return fmt.Errorf("Error reading archive_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive_log: %v", err)
		}
	}

	if err = d.Set("av_scan", flattenAntivirusProfileFtpAvScan2edl(o["av-scan"], d, "av_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-scan"], "AntivirusProfileFtp-AvScan"); ok {
			if err = d.Set("av_scan", vv); err != nil {
				return fmt.Errorf("Error reading av_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_scan: %v", err)
		}
	}

	if err = d.Set("emulator", flattenAntivirusProfileFtpEmulator2edl(o["emulator"], d, "emulator")); err != nil {
		if vv, ok := fortiAPIPatch(o["emulator"], "AntivirusProfileFtp-Emulator"); ok {
			if err = d.Set("emulator", vv); err != nil {
				return fmt.Errorf("Error reading emulator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emulator: %v", err)
		}
	}

	if err = d.Set("options", flattenAntivirusProfileFtpOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "AntivirusProfileFtp-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("external_blocklist", flattenAntivirusProfileFtpExternalBlocklist2edl(o["external-blocklist"], d, "external_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist"], "AntivirusProfileFtp-ExternalBlocklist"); ok {
			if err = d.Set("external_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist: %v", err)
		}
	}

	if err = d.Set("fortiai", flattenAntivirusProfileFtpFortiai2edl(o["fortiai"], d, "fortiai")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai"], "AntivirusProfileFtp-Fortiai"); ok {
			if err = d.Set("fortiai", vv); err != nil {
				return fmt.Errorf("Error reading fortiai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai: %v", err)
		}
	}

	if err = d.Set("fortindr", flattenAntivirusProfileFtpFortindr2edl(o["fortindr"], d, "fortindr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr"], "AntivirusProfileFtp-Fortindr"); ok {
			if err = d.Set("fortindr", vv); err != nil {
				return fmt.Errorf("Error reading fortindr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr: %v", err)
		}
	}

	if err = d.Set("fortisandbox", flattenAntivirusProfileFtpFortisandbox2edl(o["fortisandbox"], d, "fortisandbox")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox"], "AntivirusProfileFtp-Fortisandbox"); ok {
			if err = d.Set("fortisandbox", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox: %v", err)
		}
	}

	if err = d.Set("malware_stream", flattenAntivirusProfileFtpMalwareStream2edl(o["malware-stream"], d, "malware_stream")); err != nil {
		if vv, ok := fortiAPIPatch(o["malware-stream"], "AntivirusProfileFtp-MalwareStream"); ok {
			if err = d.Set("malware_stream", vv); err != nil {
				return fmt.Errorf("Error reading malware_stream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malware_stream: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention", flattenAntivirusProfileFtpOutbreakPrevention2edl(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "AntivirusProfileFtp-OutbreakPrevention"); ok {
			if err = d.Set("outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenAntivirusProfileFtpQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "AntivirusProfileFtp-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	return nil
}

func flattenAntivirusProfileFtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusProfileFtpArchiveBlock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtpArchiveLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtpAvScan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpEmulator2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtpExternalBlocklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpFortiai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpFortindr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpFortisandbox2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpMalwareStream2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpOutbreakPrevention2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusProfileFtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("archive_block"); ok || d.HasChange("archive_block") {
		t, err := expandAntivirusProfileFtpArchiveBlock2edl(d, v, "archive_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive-block"] = t
		}
	}

	if v, ok := d.GetOk("archive_log"); ok || d.HasChange("archive_log") {
		t, err := expandAntivirusProfileFtpArchiveLog2edl(d, v, "archive_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive-log"] = t
		}
	}

	if v, ok := d.GetOk("av_scan"); ok || d.HasChange("av_scan") {
		t, err := expandAntivirusProfileFtpAvScan2edl(d, v, "av_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-scan"] = t
		}
	}

	if v, ok := d.GetOk("emulator"); ok || d.HasChange("emulator") {
		t, err := expandAntivirusProfileFtpEmulator2edl(d, v, "emulator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emulator"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandAntivirusProfileFtpOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist"); ok || d.HasChange("external_blocklist") {
		t, err := expandAntivirusProfileFtpExternalBlocklist2edl(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("fortiai"); ok || d.HasChange("fortiai") {
		t, err := expandAntivirusProfileFtpFortiai2edl(d, v, "fortiai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai"] = t
		}
	}

	if v, ok := d.GetOk("fortindr"); ok || d.HasChange("fortindr") {
		t, err := expandAntivirusProfileFtpFortindr2edl(d, v, "fortindr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox"); ok || d.HasChange("fortisandbox") {
		t, err := expandAntivirusProfileFtpFortisandbox2edl(d, v, "fortisandbox")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox"] = t
		}
	}

	if v, ok := d.GetOk("malware_stream"); ok || d.HasChange("malware_stream") {
		t, err := expandAntivirusProfileFtpMalwareStream2edl(d, v, "malware_stream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malware-stream"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok || d.HasChange("outbreak_prevention") {
		t, err := expandAntivirusProfileFtpOutbreakPrevention2edl(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandAntivirusProfileFtpQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	return &obj, nil
}
