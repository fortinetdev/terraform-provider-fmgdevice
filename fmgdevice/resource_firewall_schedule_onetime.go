// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Onetime schedule configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallScheduleOnetime() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallScheduleOnetimeCreate,
		Read:   resourceFirewallScheduleOnetimeRead,
		Update: resourceFirewallScheduleOnetimeUpdate,
		Delete: resourceFirewallScheduleOnetimeDelete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_utc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expiration_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"global_object": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_utc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallScheduleOnetimeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallScheduleOnetime(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleOnetime resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallScheduleOnetime(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallScheduleOnetime(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallScheduleOnetime resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallScheduleOnetime(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallScheduleOnetime resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallScheduleOnetimeRead(d, m)
}

func resourceFirewallScheduleOnetimeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallScheduleOnetime(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleOnetime resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallScheduleOnetime(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleOnetime resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallScheduleOnetimeRead(d, m)
}

func resourceFirewallScheduleOnetimeDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallScheduleOnetime(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallScheduleOnetime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallScheduleOnetimeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallScheduleOnetime(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallScheduleOnetime resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallScheduleOnetime(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleOnetime resource from API: %v", err)
	}
	return nil
}

func flattenFirewallScheduleOnetimeColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeEnd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenFirewallScheduleOnetimeEndUtc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeExpirationDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeGlobalObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenFirewallScheduleOnetimeStartUtc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallScheduleOnetime(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("color", flattenFirewallScheduleOnetimeColor(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "FirewallScheduleOnetime-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("end", flattenFirewallScheduleOnetimeEnd(o["end"], d, "end")); err != nil {
		if vv, ok := fortiAPIPatch(o["end"], "FirewallScheduleOnetime-End"); ok {
			if err = d.Set("end", vv); err != nil {
				return fmt.Errorf("Error reading end: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("end_utc", flattenFirewallScheduleOnetimeEndUtc(o["end-utc"], d, "end_utc")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-utc"], "FirewallScheduleOnetime-EndUtc"); ok {
			if err = d.Set("end_utc", vv); err != nil {
				return fmt.Errorf("Error reading end_utc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_utc: %v", err)
		}
	}

	if err = d.Set("expiration_days", flattenFirewallScheduleOnetimeExpirationDays(o["expiration-days"], d, "expiration_days")); err != nil {
		if vv, ok := fortiAPIPatch(o["expiration-days"], "FirewallScheduleOnetime-ExpirationDays"); ok {
			if err = d.Set("expiration_days", vv); err != nil {
				return fmt.Errorf("Error reading expiration_days: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expiration_days: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallScheduleOnetimeFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-object"], "FirewallScheduleOnetime-FabricObject"); ok {
			if err = d.Set("fabric_object", vv); err != nil {
				return fmt.Errorf("Error reading fabric_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("global_object", flattenFirewallScheduleOnetimeGlobalObject(o["global-object"], d, "global_object")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-object"], "FirewallScheduleOnetime-GlobalObject"); ok {
			if err = d.Set("global_object", vv); err != nil {
				return fmt.Errorf("Error reading global_object: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_object: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallScheduleOnetimeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallScheduleOnetime-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenFirewallScheduleOnetimeStart(o["start"], d, "start")); err != nil {
		if vv, ok := fortiAPIPatch(o["start"], "FirewallScheduleOnetime-Start"); ok {
			if err = d.Set("start", vv); err != nil {
				return fmt.Errorf("Error reading start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("start_utc", flattenFirewallScheduleOnetimeStartUtc(o["start-utc"], d, "start_utc")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-utc"], "FirewallScheduleOnetime-StartUtc"); ok {
			if err = d.Set("start_utc", vv); err != nil {
				return fmt.Errorf("Error reading start_utc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_utc: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallScheduleOnetimeUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "FirewallScheduleOnetime-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenFirewallScheduleOnetimeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallScheduleOnetimeColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeEnd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeEndUtc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeExpirationDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeFabricObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeGlobalObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeStartUtc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallScheduleOnetime(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok || d.HasChange("color") {
		t, err := expandFirewallScheduleOnetimeColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok || d.HasChange("end") {
		t, err := expandFirewallScheduleOnetimeEnd(d, v, "end")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOk("end_utc"); ok || d.HasChange("end_utc") {
		t, err := expandFirewallScheduleOnetimeEndUtc(d, v, "end_utc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-utc"] = t
		}
	}

	if v, ok := d.GetOk("expiration_days"); ok || d.HasChange("expiration_days") {
		t, err := expandFirewallScheduleOnetimeExpirationDays(d, v, "expiration_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiration-days"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok || d.HasChange("fabric_object") {
		t, err := expandFirewallScheduleOnetimeFabricObject(d, v, "fabric_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("global_object"); ok || d.HasChange("global_object") {
		t, err := expandFirewallScheduleOnetimeGlobalObject(d, v, "global_object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-object"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallScheduleOnetimeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok || d.HasChange("start") {
		t, err := expandFirewallScheduleOnetimeStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("start_utc"); ok || d.HasChange("start_utc") {
		t, err := expandFirewallScheduleOnetimeStartUtc(d, v, "start_utc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-utc"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandFirewallScheduleOnetimeUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
