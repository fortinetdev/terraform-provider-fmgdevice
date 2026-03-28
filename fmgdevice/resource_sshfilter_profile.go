// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure SSH filter profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSshFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSshFilterProfileCreate,
		Read:   resourceSshFilterProfileRead,
		Update: resourceSshFilterProfileUpdate,
		Delete: resourceSshFilterProfileDelete,

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
			"block": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_command_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"shell_commands": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSshFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSshFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SshFilterProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSshFilterProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSshFilterProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SshFilterProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSshFilterProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SshFilterProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSshFilterProfileRead(d, m)
}

func resourceSshFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSshFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSshFilterProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSshFilterProfileRead(d, m)
}

func resourceSshFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteSshFilterProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SshFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSshFilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSshFilterProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SshFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSshFilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SshFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenSshFilterProfileBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSshFilterProfileDefaultCommandLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSshFilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommands(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenSshFilterProfileShellCommandsAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			v := flattenSshFilterProfileShellCommandsAlert(i["alert"], d, pre_append)
			tmp["alert"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Alert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSshFilterProfileShellCommandsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenSshFilterProfileShellCommandsLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			v := flattenSshFilterProfileShellCommandsPattern(i["pattern"], d, pre_append)
			tmp["pattern"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Pattern")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			v := flattenSshFilterProfileShellCommandsSeverity(i["severity"], d, pre_append)
			tmp["severity"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Severity")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSshFilterProfileShellCommandsType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SshFilterProfile-ShellCommands-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSshFilterProfileShellCommandsAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsAlert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSshFilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("block", flattenSshFilterProfileBlock(o["block"], d, "block")); err != nil {
		if vv, ok := fortiAPIPatch(o["block"], "SshFilterProfile-Block"); ok {
			if err = d.Set("block", vv); err != nil {
				return fmt.Errorf("Error reading block: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block: %v", err)
		}
	}

	if err = d.Set("default_command_log", flattenSshFilterProfileDefaultCommandLog(o["default-command-log"], d, "default_command_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-command-log"], "SshFilterProfile-DefaultCommandLog"); ok {
			if err = d.Set("default_command_log", vv); err != nil {
				return fmt.Errorf("Error reading default_command_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_command_log: %v", err)
		}
	}

	if err = d.Set("log", flattenSshFilterProfileLog(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "SshFilterProfile-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenSshFilterProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SshFilterProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shell_commands", flattenSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands")); err != nil {
			if vv, ok := fortiAPIPatch(o["shell-commands"], "SshFilterProfile-ShellCommands"); ok {
				if err = d.Set("shell_commands", vv); err != nil {
					return fmt.Errorf("Error reading shell_commands: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading shell_commands: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shell_commands"); ok {
			if err = d.Set("shell_commands", flattenSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands")); err != nil {
				if vv, ok := fortiAPIPatch(o["shell-commands"], "SshFilterProfile-ShellCommands"); ok {
					if err = d.Set("shell_commands", vv); err != nil {
						return fmt.Errorf("Error reading shell_commands: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading shell_commands: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSshFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSshFilterProfileBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSshFilterProfileDefaultCommandLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSshFilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommands(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandSshFilterProfileShellCommandsAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["alert"], _ = expandSshFilterProfileShellCommandsAlert(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSshFilterProfileShellCommandsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandSshFilterProfileShellCommandsLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandSshFilterProfileShellCommandsPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["severity"], _ = expandSshFilterProfileShellCommandsSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSshFilterProfileShellCommandsType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSshFilterProfileShellCommandsAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSshFilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block"); ok || d.HasChange("block") {
		t, err := expandSshFilterProfileBlock(d, v, "block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block"] = t
		}
	}

	if v, ok := d.GetOk("default_command_log"); ok || d.HasChange("default_command_log") {
		t, err := expandSshFilterProfileDefaultCommandLog(d, v, "default_command_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-command-log"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandSshFilterProfileLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSshFilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("shell_commands"); ok || d.HasChange("shell_commands") {
		t, err := expandSshFilterProfileShellCommands(d, v, "shell_commands")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shell-commands"] = t
		}
	}

	return &obj, nil
}
