// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SNMP Remote Network Monitoring (RMON) Ethernet statistics configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpRmonStat() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpRmonStatCreate,
		Read:   resourceSystemSnmpRmonStatRead,
		Update: resourceSystemSnmpRmonStatUpdate,
		Delete: resourceSystemSnmpRmonStatDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSnmpRmonStatCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSnmpRmonStat(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpRmonStat resource while getting object: %v", err)
	}

	v, err := c.CreateSystemSnmpRmonStat(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpRmonStat resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceSystemSnmpRmonStatRead(d, m)
		} else {
			return fmt.Errorf("Error creating SystemSnmpRmonStat resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpRmonStatRead(d, m)
}

func resourceSystemSnmpRmonStatUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSnmpRmonStat(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpRmonStat resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpRmonStat(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpRmonStat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemSnmpRmonStatRead(d, m)
}

func resourceSystemSnmpRmonStatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemSnmpRmonStat(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpRmonStat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpRmonStatRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSnmpRmonStat(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpRmonStat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpRmonStat(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpRmonStat resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpRmonStatId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpRmonStatOwner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpRmonStatSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSnmpRmonStat(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemSnmpRmonStatId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemSnmpRmonStat-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("owner", flattenSystemSnmpRmonStatOwner(o["owner"], d, "owner")); err != nil {
		if vv, ok := fortiAPIPatch(o["owner"], "SystemSnmpRmonStat-Owner"); ok {
			if err = d.Set("owner", vv); err != nil {
				return fmt.Errorf("Error reading owner: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading owner: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemSnmpRmonStatSource(o["source"], d, "source")); err != nil {
		if vv, ok := fortiAPIPatch(o["source"], "SystemSnmpRmonStat-Source"); ok {
			if err = d.Set("source", vv); err != nil {
				return fmt.Errorf("Error reading source: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpRmonStatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpRmonStatId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpRmonStatOwner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpRmonStatSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSnmpRmonStat(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemSnmpRmonStatId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("owner"); ok || d.HasChange("owner") {
		t, err := expandSystemSnmpRmonStatOwner(d, v, "owner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owner"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok || d.HasChange("source") {
		t, err := expandSystemSnmpRmonStatSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
