// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure MAPI protocol options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptionsMapi() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsMapiUpdate,
		Read:   resourceFirewallProfileProtocolOptionsMapiRead,
		Update: resourceFirewallProfileProtocolOptionsMapiUpdate,
		Delete: resourceFirewallProfileProtocolOptionsMapiDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uncompressed_nest_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uncompressed_oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallProfileProtocolOptionsMapiUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectFirewallProfileProtocolOptionsMapi(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsMapi resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptionsMapi(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsMapi resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallProfileProtocolOptionsMapi")

	return resourceFirewallProfileProtocolOptionsMapiRead(d, m)
}

func resourceFirewallProfileProtocolOptionsMapiDelete(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	wsParams["adom"] = adomv

	err = c.DeleteFirewallProfileProtocolOptionsMapi(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptionsMapi resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsMapiRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile_protocol_options := d.Get("profile_protocol_options").(string)
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
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadFirewallProfileProtocolOptionsMapi(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsMapi resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptionsMapi(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsMapi resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsMapiOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsMapiOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsMapiScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptionsMapi(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("options", flattenFirewallProfileProtocolOptionsMapiOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "FirewallProfileProtocolOptionsMapi-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenFirewallProfileProtocolOptionsMapiOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "FirewallProfileProtocolOptionsMapi-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("ports", flattenFirewallProfileProtocolOptionsMapiPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "FirewallProfileProtocolOptionsMapi-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenFirewallProfileProtocolOptionsMapiScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "FirewallProfileProtocolOptionsMapi-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallProfileProtocolOptionsMapiStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallProfileProtocolOptionsMapi-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenFirewallProfileProtocolOptionsMapiUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "FirewallProfileProtocolOptionsMapi-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "FirewallProfileProtocolOptionsMapi-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsMapiFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsMapiOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsMapiOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsMapiScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptionsMapi(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandFirewallProfileProtocolOptionsMapiOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsMapiOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandFirewallProfileProtocolOptionsMapiPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandFirewallProfileProtocolOptionsMapiScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallProfileProtocolOptionsMapiStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandFirewallProfileProtocolOptionsMapiUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsMapiUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	return &obj, nil
}
