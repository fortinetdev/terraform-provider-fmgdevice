// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Settings for TACACS+ accounting events filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogTacacsAccounting2Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogTacacsAccounting2FilterUpdate,
		Read:   resourceLogTacacsAccounting2FilterRead,
		Update: resourceLogTacacsAccounting2FilterUpdate,
		Delete: resourceLogTacacsAccounting2FilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"cli_cmd_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_change_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogTacacsAccounting2FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectLogTacacsAccounting2Filter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting2Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogTacacsAccounting2Filter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting2Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogTacacsAccounting2Filter")

	return resourceLogTacacsAccounting2FilterRead(d, m)
}

func resourceLogTacacsAccounting2FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteLogTacacsAccounting2Filter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogTacacsAccounting2Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsAccounting2FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadLogTacacsAccounting2Filter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccounting2Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogTacacsAccounting2Filter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccounting2Filter resource from API: %v", err)
	}
	return nil
}

func flattenLogTacacsAccounting2FilterCliCmdAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogTacacsAccounting2FilterConfigChangeAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogTacacsAccounting2FilterLoginAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogTacacsAccounting2Filter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cli_cmd_audit", flattenLogTacacsAccounting2FilterCliCmdAudit(o["cli-cmd-audit"], d, "cli_cmd_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-cmd-audit"], "LogTacacsAccounting2Filter-CliCmdAudit"); ok {
			if err = d.Set("cli_cmd_audit", vv); err != nil {
				return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
		}
	}

	if err = d.Set("config_change_audit", flattenLogTacacsAccounting2FilterConfigChangeAudit(o["config-change-audit"], d, "config_change_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-change-audit"], "LogTacacsAccounting2Filter-ConfigChangeAudit"); ok {
			if err = d.Set("config_change_audit", vv); err != nil {
				return fmt.Errorf("Error reading config_change_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_change_audit: %v", err)
		}
	}

	if err = d.Set("login_audit", flattenLogTacacsAccounting2FilterLoginAudit(o["login-audit"], d, "login_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-audit"], "LogTacacsAccounting2Filter-LoginAudit"); ok {
			if err = d.Set("login_audit", vv); err != nil {
				return fmt.Errorf("Error reading login_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_audit: %v", err)
		}
	}

	return nil
}

func flattenLogTacacsAccounting2FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogTacacsAccounting2FilterCliCmdAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccounting2FilterConfigChangeAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccounting2FilterLoginAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogTacacsAccounting2Filter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cli_cmd_audit"); ok || d.HasChange("cli_cmd_audit") {
		t, err := expandLogTacacsAccounting2FilterCliCmdAudit(d, v, "cli_cmd_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-cmd-audit"] = t
		}
	}

	if v, ok := d.GetOk("config_change_audit"); ok || d.HasChange("config_change_audit") {
		t, err := expandLogTacacsAccounting2FilterConfigChangeAudit(d, v, "config_change_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-change-audit"] = t
		}
	}

	if v, ok := d.GetOk("login_audit"); ok || d.HasChange("login_audit") {
		t, err := expandLogTacacsAccounting2FilterLoginAudit(d, v, "login_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-audit"] = t
		}
	}

	return &obj, nil
}
