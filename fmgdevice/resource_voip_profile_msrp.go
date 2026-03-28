// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> MSRP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVoipProfileMsrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVoipProfileMsrpUpdate,
		Read:   resourceVoipProfileMsrpRead,
		Update: resourceVoipProfileMsrpUpdate,
		Delete: resourceVoipProfileMsrpDelete,

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
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_msg_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_msg_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVoipProfileMsrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVoipProfileMsrp(d)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfileMsrp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateVoipProfileMsrp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VoipProfileMsrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("VoipProfileMsrp")

	return resourceVoipProfileMsrpRead(d, m)
}

func resourceVoipProfileMsrpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVoipProfileMsrp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VoipProfileMsrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVoipProfileMsrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVoipProfileMsrp(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading VoipProfileMsrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVoipProfileMsrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VoipProfileMsrp resource from API: %v", err)
	}
	return nil
}

func flattenVoipProfileMsrpLogViolations2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileMsrpMaxMsgSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileMsrpMaxMsgSizeAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVoipProfileMsrpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVoipProfileMsrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("log_violations", flattenVoipProfileMsrpLogViolations2edl(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "VoipProfileMsrp-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("max_msg_size", flattenVoipProfileMsrpMaxMsgSize2edl(o["max-msg-size"], d, "max_msg_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-msg-size"], "VoipProfileMsrp-MaxMsgSize"); ok {
			if err = d.Set("max_msg_size", vv); err != nil {
				return fmt.Errorf("Error reading max_msg_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_msg_size: %v", err)
		}
	}

	if err = d.Set("max_msg_size_action", flattenVoipProfileMsrpMaxMsgSizeAction2edl(o["max-msg-size-action"], d, "max_msg_size_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-msg-size-action"], "VoipProfileMsrp-MaxMsgSizeAction"); ok {
			if err = d.Set("max_msg_size_action", vv); err != nil {
				return fmt.Errorf("Error reading max_msg_size_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_msg_size_action: %v", err)
		}
	}

	if err = d.Set("status", flattenVoipProfileMsrpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VoipProfileMsrp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenVoipProfileMsrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVoipProfileMsrpLogViolations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrpMaxMsgSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrpMaxMsgSizeAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVoipProfileMsrpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVoipProfileMsrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandVoipProfileMsrpLogViolations2edl(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("max_msg_size"); ok || d.HasChange("max_msg_size") {
		t, err := expandVoipProfileMsrpMaxMsgSize2edl(d, v, "max_msg_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-msg-size"] = t
		}
	}

	if v, ok := d.GetOk("max_msg_size_action"); ok || d.HasChange("max_msg_size_action") {
		t, err := expandVoipProfileMsrpMaxMsgSizeAction2edl(d, v, "max_msg_size_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-msg-size-action"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVoipProfileMsrpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
