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

func resourceLogTacacsAccountingFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogTacacsAccountingFilterUpdate,
		Read:   resourceLogTacacsAccountingFilterRead,
		Update: resourceLogTacacsAccountingFilterUpdate,
		Delete: resourceLogTacacsAccountingFilterDelete,

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

func resourceLogTacacsAccountingFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogTacacsAccountingFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccountingFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogTacacsAccountingFilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccountingFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogTacacsAccountingFilter")

	return resourceLogTacacsAccountingFilterRead(d, m)
}

func resourceLogTacacsAccountingFilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogTacacsAccountingFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting LogTacacsAccountingFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsAccountingFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogTacacsAccountingFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccountingFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogTacacsAccountingFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccountingFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogTacacsAccountingFilterCliCmdAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogTacacsAccountingFilterConfigChangeAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogTacacsAccountingFilterLoginAudit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogTacacsAccountingFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cli_cmd_audit", flattenLogTacacsAccountingFilterCliCmdAudit(o["cli-cmd-audit"], d, "cli_cmd_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-cmd-audit"], "LogTacacsAccountingFilter-CliCmdAudit"); ok {
			if err = d.Set("cli_cmd_audit", vv); err != nil {
				return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
		}
	}

	if err = d.Set("config_change_audit", flattenLogTacacsAccountingFilterConfigChangeAudit(o["config-change-audit"], d, "config_change_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["config-change-audit"], "LogTacacsAccountingFilter-ConfigChangeAudit"); ok {
			if err = d.Set("config_change_audit", vv); err != nil {
				return fmt.Errorf("Error reading config_change_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading config_change_audit: %v", err)
		}
	}

	if err = d.Set("login_audit", flattenLogTacacsAccountingFilterLoginAudit(o["login-audit"], d, "login_audit")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-audit"], "LogTacacsAccountingFilter-LoginAudit"); ok {
			if err = d.Set("login_audit", vv); err != nil {
				return fmt.Errorf("Error reading login_audit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_audit: %v", err)
		}
	}

	return nil
}

func flattenLogTacacsAccountingFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogTacacsAccountingFilterCliCmdAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccountingFilterConfigChangeAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccountingFilterLoginAudit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogTacacsAccountingFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cli_cmd_audit"); ok || d.HasChange("cli_cmd_audit") {
		t, err := expandLogTacacsAccountingFilterCliCmdAudit(d, v, "cli_cmd_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-cmd-audit"] = t
		}
	}

	if v, ok := d.GetOk("config_change_audit"); ok || d.HasChange("config_change_audit") {
		t, err := expandLogTacacsAccountingFilterConfigChangeAudit(d, v, "config_change_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-change-audit"] = t
		}
	}

	if v, ok := d.GetOk("login_audit"); ok || d.HasChange("login_audit") {
		t, err := expandLogTacacsAccountingFilterLoginAudit(d, v, "login_audit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-audit"] = t
		}
	}

	return &obj, nil
}
