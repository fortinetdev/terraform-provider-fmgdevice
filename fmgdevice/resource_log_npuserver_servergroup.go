// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> create server group.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogNpuServerServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogNpuServerServerGroupCreate,
		Read:   resourceLogNpuServerServerGroupRead,
		Update: resourceLogNpuServerServerGroupUpdate,
		Delete: resourceLogNpuServerServerGroupDelete,

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
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"log_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_gen_event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_tx_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_start_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sw_log_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceLogNpuServerServerGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectLogNpuServerServerGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating LogNpuServerServerGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("group_name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLogNpuServerServerGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLogNpuServerServerGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating LogNpuServerServerGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateLogNpuServerServerGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating LogNpuServerServerGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "group_name"))

	return resourceLogNpuServerServerGroupRead(d, m)
}

func resourceLogNpuServerServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectLogNpuServerServerGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating LogNpuServerServerGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateLogNpuServerServerGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogNpuServerServerGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "group_name"))

	return resourceLogNpuServerServerGroupRead(d, m)
}

func resourceLogNpuServerServerGroupDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteLogNpuServerServerGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogNpuServerServerGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogNpuServerServerGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadLogNpuServerServerGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading LogNpuServerServerGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogNpuServerServerGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogNpuServerServerGroup resource from API: %v", err)
	}
	return nil
}

func flattenLogNpuServerServerGroupGroupName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupLogFormat2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupLogGenEvent2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupLogMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupLogTxMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupLogUserInfo2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupServerNumber2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupServerStartId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogNpuServerServerGroupSwLogFlags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogNpuServerServerGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group_name", flattenLogNpuServerServerGroupGroupName2edl(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "LogNpuServerServerGroup-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("log_format", flattenLogNpuServerServerGroupLogFormat2edl(o["log-format"], d, "log_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-format"], "LogNpuServerServerGroup-LogFormat"); ok {
			if err = d.Set("log_format", vv); err != nil {
				return fmt.Errorf("Error reading log_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_format: %v", err)
		}
	}

	if err = d.Set("log_gen_event", flattenLogNpuServerServerGroupLogGenEvent2edl(o["log-gen-event"], d, "log_gen_event")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-gen-event"], "LogNpuServerServerGroup-LogGenEvent"); ok {
			if err = d.Set("log_gen_event", vv); err != nil {
				return fmt.Errorf("Error reading log_gen_event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_gen_event: %v", err)
		}
	}

	if err = d.Set("log_mode", flattenLogNpuServerServerGroupLogMode2edl(o["log-mode"], d, "log_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-mode"], "LogNpuServerServerGroup-LogMode"); ok {
			if err = d.Set("log_mode", vv); err != nil {
				return fmt.Errorf("Error reading log_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_mode: %v", err)
		}
	}

	if err = d.Set("log_tx_mode", flattenLogNpuServerServerGroupLogTxMode2edl(o["log-tx-mode"], d, "log_tx_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-tx-mode"], "LogNpuServerServerGroup-LogTxMode"); ok {
			if err = d.Set("log_tx_mode", vv); err != nil {
				return fmt.Errorf("Error reading log_tx_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_tx_mode: %v", err)
		}
	}

	if err = d.Set("log_user_info", flattenLogNpuServerServerGroupLogUserInfo2edl(o["log-user-info"], d, "log_user_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-user-info"], "LogNpuServerServerGroup-LogUserInfo"); ok {
			if err = d.Set("log_user_info", vv); err != nil {
				return fmt.Errorf("Error reading log_user_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_user_info: %v", err)
		}
	}

	if err = d.Set("server_number", flattenLogNpuServerServerGroupServerNumber2edl(o["server-number"], d, "server_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-number"], "LogNpuServerServerGroup-ServerNumber"); ok {
			if err = d.Set("server_number", vv); err != nil {
				return fmt.Errorf("Error reading server_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_number: %v", err)
		}
	}

	if err = d.Set("server_start_id", flattenLogNpuServerServerGroupServerStartId2edl(o["server-start-id"], d, "server_start_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-start-id"], "LogNpuServerServerGroup-ServerStartId"); ok {
			if err = d.Set("server_start_id", vv); err != nil {
				return fmt.Errorf("Error reading server_start_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_start_id: %v", err)
		}
	}

	if err = d.Set("sw_log_flags", flattenLogNpuServerServerGroupSwLogFlags2edl(o["sw-log-flags"], d, "sw_log_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["sw-log-flags"], "LogNpuServerServerGroup-SwLogFlags"); ok {
			if err = d.Set("sw_log_flags", vv); err != nil {
				return fmt.Errorf("Error reading sw_log_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sw_log_flags: %v", err)
		}
	}

	return nil
}

func flattenLogNpuServerServerGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogNpuServerServerGroupGroupName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupLogFormat2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupLogGenEvent2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupLogMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupLogTxMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupLogUserInfo2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupServerNumber2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupServerStartId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogNpuServerServerGroupSwLogFlags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogNpuServerServerGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandLogNpuServerServerGroupGroupName2edl(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("log_format"); ok || d.HasChange("log_format") {
		t, err := expandLogNpuServerServerGroupLogFormat2edl(d, v, "log_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-format"] = t
		}
	}

	if v, ok := d.GetOk("log_gen_event"); ok || d.HasChange("log_gen_event") {
		t, err := expandLogNpuServerServerGroupLogGenEvent2edl(d, v, "log_gen_event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-gen-event"] = t
		}
	}

	if v, ok := d.GetOk("log_mode"); ok || d.HasChange("log_mode") {
		t, err := expandLogNpuServerServerGroupLogMode2edl(d, v, "log_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-mode"] = t
		}
	}

	if v, ok := d.GetOk("log_tx_mode"); ok || d.HasChange("log_tx_mode") {
		t, err := expandLogNpuServerServerGroupLogTxMode2edl(d, v, "log_tx_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-tx-mode"] = t
		}
	}

	if v, ok := d.GetOk("log_user_info"); ok || d.HasChange("log_user_info") {
		t, err := expandLogNpuServerServerGroupLogUserInfo2edl(d, v, "log_user_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-user-info"] = t
		}
	}

	if v, ok := d.GetOk("server_number"); ok || d.HasChange("server_number") {
		t, err := expandLogNpuServerServerGroupServerNumber2edl(d, v, "server_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-number"] = t
		}
	}

	if v, ok := d.GetOk("server_start_id"); ok || d.HasChange("server_start_id") {
		t, err := expandLogNpuServerServerGroupServerStartId2edl(d, v, "server_start_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-start-id"] = t
		}
	}

	if v, ok := d.GetOk("sw_log_flags"); ok || d.HasChange("sw_log_flags") {
		t, err := expandLogNpuServerServerGroupSwLogFlags2edl(d, v, "sw_log_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-log-flags"] = t
		}
	}

	return &obj, nil
}
