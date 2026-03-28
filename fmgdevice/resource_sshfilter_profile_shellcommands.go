// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SSH command filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSshFilterProfileShellCommands() *schema.Resource {
	return &schema.Resource{
		Create: resourceSshFilterProfileShellCommandsCreate,
		Read:   resourceSshFilterProfileShellCommandsRead,
		Update: resourceSshFilterProfileShellCommandsUpdate,
		Delete: resourceSshFilterProfileShellCommandsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSshFilterProfileShellCommandsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSshFilterProfileShellCommands(d)
	if err != nil {
		return fmt.Errorf("Error creating SshFilterProfileShellCommands resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSshFilterProfileShellCommands(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSshFilterProfileShellCommands(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SshFilterProfileShellCommands resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSshFilterProfileShellCommands(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SshFilterProfileShellCommands resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSshFilterProfileShellCommandsRead(d, m)
			} else {
				return fmt.Errorf("Error creating SshFilterProfileShellCommands resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSshFilterProfileShellCommandsRead(d, m)
}

func resourceSshFilterProfileShellCommandsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSshFilterProfileShellCommands(d)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfileShellCommands resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSshFilterProfileShellCommands(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfileShellCommands resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSshFilterProfileShellCommandsRead(d, m)
}

func resourceSshFilterProfileShellCommandsDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSshFilterProfileShellCommands(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SshFilterProfileShellCommands resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSshFilterProfileShellCommandsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSshFilterProfileShellCommands(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SshFilterProfileShellCommands resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSshFilterProfileShellCommands(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SshFilterProfileShellCommands resource from API: %v", err)
	}
	return nil
}

func flattenSshFilterProfileShellCommandsAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsAlert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSshFilterProfileShellCommands(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSshFilterProfileShellCommandsAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SshFilterProfileShellCommands-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("alert", flattenSshFilterProfileShellCommandsAlert2edl(o["alert"], d, "alert")); err != nil {
		if vv, ok := fortiAPIPatch(o["alert"], "SshFilterProfileShellCommands-Alert"); ok {
			if err = d.Set("alert", vv); err != nil {
				return fmt.Errorf("Error reading alert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alert: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSshFilterProfileShellCommandsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SshFilterProfileShellCommands-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenSshFilterProfileShellCommandsLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "SshFilterProfileShellCommands-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("pattern", flattenSshFilterProfileShellCommandsPattern2edl(o["pattern"], d, "pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["pattern"], "SshFilterProfileShellCommands-Pattern"); ok {
			if err = d.Set("pattern", vv); err != nil {
				return fmt.Errorf("Error reading pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("severity", flattenSshFilterProfileShellCommandsSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "SshFilterProfileShellCommands-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("type", flattenSshFilterProfileShellCommandsType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SshFilterProfileShellCommands-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSshFilterProfileShellCommandsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSshFilterProfileShellCommandsAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsAlert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSshFilterProfileShellCommands(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSshFilterProfileShellCommandsAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("alert"); ok || d.HasChange("alert") {
		t, err := expandSshFilterProfileShellCommandsAlert2edl(d, v, "alert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alert"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSshFilterProfileShellCommandsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandSshFilterProfileShellCommandsLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok || d.HasChange("pattern") {
		t, err := expandSshFilterProfileShellCommandsPattern2edl(d, v, "pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandSshFilterProfileShellCommandsSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSshFilterProfileShellCommandsType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
