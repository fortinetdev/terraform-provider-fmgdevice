// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SCCP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVoipProfileSccp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVoipProfileSccpUpdate,
		Read:   resourceVoipProfileSccpRead,
		Update: resourceVoipProfileSccpUpdate,
		Delete: resourceVoipProfileSccpDelete,

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
			"block_mcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_call_summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_calls": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVoipProfileSccpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVoipProfileSccp(d)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfileSccp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVoipProfileSccp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfileSccp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VoipProfileSccp")

	return resourceVoipProfileSccpRead(d, m)
}

func resourceVoipProfileSccpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVoipProfileSccp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VoipProfileSccp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVoipProfileSccpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVoipProfileSccp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VoipProfileSccp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVoipProfileSccp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VoipProfileSccp resource from API: %v", err)
	}
	return nil
}

func flattenVoipProfileSccpBlockMcast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSccpLogCallSummary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSccpLogViolations2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSccpMaxCalls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSccpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileSccpVerifyHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVoipProfileSccp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("block_mcast", flattenVoipProfileSccpBlockMcast2edl(o["block-mcast"], d, "block_mcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-mcast"], "VoipProfileSccp-BlockMcast"); ok {
			if err = d.Set("block_mcast", vv); err != nil {
				return fmt.Errorf("Error reading block_mcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_mcast: %v", err)
		}
	}

	if err = d.Set("log_call_summary", flattenVoipProfileSccpLogCallSummary2edl(o["log-call-summary"], d, "log_call_summary")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-call-summary"], "VoipProfileSccp-LogCallSummary"); ok {
			if err = d.Set("log_call_summary", vv); err != nil {
				return fmt.Errorf("Error reading log_call_summary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_call_summary: %v", err)
		}
	}

	if err = d.Set("log_violations", flattenVoipProfileSccpLogViolations2edl(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "VoipProfileSccp-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("max_calls", flattenVoipProfileSccpMaxCalls2edl(o["max-calls"], d, "max_calls")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-calls"], "VoipProfileSccp-MaxCalls"); ok {
			if err = d.Set("max_calls", vv); err != nil {
				return fmt.Errorf("Error reading max_calls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_calls: %v", err)
		}
	}

	if err = d.Set("status", flattenVoipProfileSccpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VoipProfileSccp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("verify_header", flattenVoipProfileSccpVerifyHeader2edl(o["verify-header"], d, "verify_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-header"], "VoipProfileSccp-VerifyHeader"); ok {
			if err = d.Set("verify_header", vv); err != nil {
				return fmt.Errorf("Error reading verify_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_header: %v", err)
		}
	}

	return nil
}

func flattenVoipProfileSccpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVoipProfileSccpBlockMcast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpLogCallSummary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpLogViolations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpMaxCalls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileSccpVerifyHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVoipProfileSccp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_mcast"); ok || d.HasChange("block_mcast") {
		t, err := expandVoipProfileSccpBlockMcast2edl(d, v, "block_mcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-mcast"] = t
		}
	}

	if v, ok := d.GetOk("log_call_summary"); ok || d.HasChange("log_call_summary") {
		t, err := expandVoipProfileSccpLogCallSummary2edl(d, v, "log_call_summary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-call-summary"] = t
		}
	}

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandVoipProfileSccpLogViolations2edl(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("max_calls"); ok || d.HasChange("max_calls") {
		t, err := expandVoipProfileSccpMaxCalls2edl(d, v, "max_calls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-calls"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVoipProfileSccpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("verify_header"); ok || d.HasChange("verify_header") {
		t, err := expandVoipProfileSccpVerifyHeader2edl(d, v, "verify_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-header"] = t
		}
	}

	return &obj, nil
}
