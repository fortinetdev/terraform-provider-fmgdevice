// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure IMAP AntiVirus options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusProfileImap() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusProfileImapUpdate,
		Read:   resourceAntivirusProfileImapRead,
		Update: resourceAntivirusProfileImapUpdate,
		Delete: resourceAntivirusProfileImapDelete,

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
			"executables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceAntivirusProfileImapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAntivirusProfileImap(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileImap resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAntivirusProfileImap(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfileImap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("AntivirusProfileImap")

	return resourceAntivirusProfileImapRead(d, m)
}

func resourceAntivirusProfileImapDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAntivirusProfileImap(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusProfileImap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusProfileImapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAntivirusProfileImap(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AntivirusProfileImap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusProfileImap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusProfileImap resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusProfileImapArchiveBlock2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileImapArchiveLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileImapAvScan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapContentDisarm2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapEmulator2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapExecutables2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileImapExternalBlocklist2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapFortiai2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapFortindr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapFortisandbox2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapMalwareStream2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapOutbreakPrevention2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusProfileImap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("archive_block", flattenAntivirusProfileImapArchiveBlock2edl(o["archive-block"], d, "archive_block")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive-block"], "AntivirusProfileImap-ArchiveBlock"); ok {
			if err = d.Set("archive_block", vv); err != nil {
				return fmt.Errorf("Error reading archive_block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive_block: %v", err)
		}
	}

	if err = d.Set("archive_log", flattenAntivirusProfileImapArchiveLog2edl(o["archive-log"], d, "archive_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["archive-log"], "AntivirusProfileImap-ArchiveLog"); ok {
			if err = d.Set("archive_log", vv); err != nil {
				return fmt.Errorf("Error reading archive_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading archive_log: %v", err)
		}
	}

	if err = d.Set("av_scan", flattenAntivirusProfileImapAvScan2edl(o["av-scan"], d, "av_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-scan"], "AntivirusProfileImap-AvScan"); ok {
			if err = d.Set("av_scan", vv); err != nil {
				return fmt.Errorf("Error reading av_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_scan: %v", err)
		}
	}

	if err = d.Set("content_disarm", flattenAntivirusProfileImapContentDisarm2edl(o["content-disarm"], d, "content_disarm")); err != nil {
		if vv, ok := fortiAPIPatch(o["content-disarm"], "AntivirusProfileImap-ContentDisarm"); ok {
			if err = d.Set("content_disarm", vv); err != nil {
				return fmt.Errorf("Error reading content_disarm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading content_disarm: %v", err)
		}
	}

	if err = d.Set("emulator", flattenAntivirusProfileImapEmulator2edl(o["emulator"], d, "emulator")); err != nil {
		if vv, ok := fortiAPIPatch(o["emulator"], "AntivirusProfileImap-Emulator"); ok {
			if err = d.Set("emulator", vv); err != nil {
				return fmt.Errorf("Error reading emulator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emulator: %v", err)
		}
	}

	if err = d.Set("executables", flattenAntivirusProfileImapExecutables2edl(o["executables"], d, "executables")); err != nil {
		if vv, ok := fortiAPIPatch(o["executables"], "AntivirusProfileImap-Executables"); ok {
			if err = d.Set("executables", vv); err != nil {
				return fmt.Errorf("Error reading executables: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading executables: %v", err)
		}
	}

	if err = d.Set("options", flattenAntivirusProfileImapOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "AntivirusProfileImap-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("external_blocklist", flattenAntivirusProfileImapExternalBlocklist2edl(o["external-blocklist"], d, "external_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist"], "AntivirusProfileImap-ExternalBlocklist"); ok {
			if err = d.Set("external_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist: %v", err)
		}
	}

	if err = d.Set("fortiai", flattenAntivirusProfileImapFortiai2edl(o["fortiai"], d, "fortiai")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai"], "AntivirusProfileImap-Fortiai"); ok {
			if err = d.Set("fortiai", vv); err != nil {
				return fmt.Errorf("Error reading fortiai: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai: %v", err)
		}
	}

	if err = d.Set("fortindr", flattenAntivirusProfileImapFortindr2edl(o["fortindr"], d, "fortindr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr"], "AntivirusProfileImap-Fortindr"); ok {
			if err = d.Set("fortindr", vv); err != nil {
				return fmt.Errorf("Error reading fortindr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr: %v", err)
		}
	}

	if err = d.Set("fortisandbox", flattenAntivirusProfileImapFortisandbox2edl(o["fortisandbox"], d, "fortisandbox")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox"], "AntivirusProfileImap-Fortisandbox"); ok {
			if err = d.Set("fortisandbox", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox: %v", err)
		}
	}

	if err = d.Set("malware_stream", flattenAntivirusProfileImapMalwareStream2edl(o["malware-stream"], d, "malware_stream")); err != nil {
		if vv, ok := fortiAPIPatch(o["malware-stream"], "AntivirusProfileImap-MalwareStream"); ok {
			if err = d.Set("malware_stream", vv); err != nil {
				return fmt.Errorf("Error reading malware_stream: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading malware_stream: %v", err)
		}
	}

	if err = d.Set("outbreak_prevention", flattenAntivirusProfileImapOutbreakPrevention2edl(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "AntivirusProfileImap-OutbreakPrevention"); ok {
			if err = d.Set("outbreak_prevention", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenAntivirusProfileImapQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "AntivirusProfileImap-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	return nil
}

func flattenAntivirusProfileImapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusProfileImapArchiveBlock2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileImapArchiveLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileImapAvScan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapContentDisarm2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapEmulator2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapExecutables2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileImapExternalBlocklist2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapFortiai2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapFortindr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapFortisandbox2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapMalwareStream2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapOutbreakPrevention2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusProfileImap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("archive_block"); ok || d.HasChange("archive_block") {
		t, err := expandAntivirusProfileImapArchiveBlock2edl(d, v, "archive_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive-block"] = t
		}
	}

	if v, ok := d.GetOk("archive_log"); ok || d.HasChange("archive_log") {
		t, err := expandAntivirusProfileImapArchiveLog2edl(d, v, "archive_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["archive-log"] = t
		}
	}

	if v, ok := d.GetOk("av_scan"); ok || d.HasChange("av_scan") {
		t, err := expandAntivirusProfileImapAvScan2edl(d, v, "av_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-scan"] = t
		}
	}

	if v, ok := d.GetOk("content_disarm"); ok || d.HasChange("content_disarm") {
		t, err := expandAntivirusProfileImapContentDisarm2edl(d, v, "content_disarm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-disarm"] = t
		}
	}

	if v, ok := d.GetOk("emulator"); ok || d.HasChange("emulator") {
		t, err := expandAntivirusProfileImapEmulator2edl(d, v, "emulator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emulator"] = t
		}
	}

	if v, ok := d.GetOk("executables"); ok || d.HasChange("executables") {
		t, err := expandAntivirusProfileImapExecutables2edl(d, v, "executables")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["executables"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandAntivirusProfileImapOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist"); ok || d.HasChange("external_blocklist") {
		t, err := expandAntivirusProfileImapExternalBlocklist2edl(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("fortiai"); ok || d.HasChange("fortiai") {
		t, err := expandAntivirusProfileImapFortiai2edl(d, v, "fortiai")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai"] = t
		}
	}

	if v, ok := d.GetOk("fortindr"); ok || d.HasChange("fortindr") {
		t, err := expandAntivirusProfileImapFortindr2edl(d, v, "fortindr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox"); ok || d.HasChange("fortisandbox") {
		t, err := expandAntivirusProfileImapFortisandbox2edl(d, v, "fortisandbox")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox"] = t
		}
	}

	if v, ok := d.GetOk("malware_stream"); ok || d.HasChange("malware_stream") {
		t, err := expandAntivirusProfileImapMalwareStream2edl(d, v, "malware_stream")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malware-stream"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok || d.HasChange("outbreak_prevention") {
		t, err := expandAntivirusProfileImapOutbreakPrevention2edl(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandAntivirusProfileImapQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	return &obj, nil
}
