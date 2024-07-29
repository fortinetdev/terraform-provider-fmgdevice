// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure pool exclude subnets.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpamPoolsExclude() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpamPoolsExcludeCreate,
		Read:   resourceSystemIpamPoolsExcludeRead,
		Update: resourceSystemIpamPoolsExcludeUpdate,
		Delete: resourceSystemIpamPoolsExcludeDelete,

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
			"pools": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"exclude_subnet": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpamPoolsExcludeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	pools := d.Get("pools").(string)
	paradict["device"] = device_name
	paradict["pools"] = pools

	obj, err := getObjectSystemIpamPoolsExclude(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpamPoolsExclude resource while getting object: %v", err)
	}

	_, err = c.CreateSystemIpamPoolsExclude(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpamPoolsExclude resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemIpamPoolsExcludeRead(d, m)
}

func resourceSystemIpamPoolsExcludeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	pools := d.Get("pools").(string)
	paradict["device"] = device_name
	paradict["pools"] = pools

	obj, err := getObjectSystemIpamPoolsExclude(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpamPoolsExclude resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpamPoolsExclude(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpamPoolsExclude resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemIpamPoolsExcludeRead(d, m)
}

func resourceSystemIpamPoolsExcludeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	pools := d.Get("pools").(string)
	paradict["device"] = device_name
	paradict["pools"] = pools

	err = c.DeleteSystemIpamPoolsExclude(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpamPoolsExclude resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpamPoolsExcludeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	pools := d.Get("pools").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if pools == "" {
		pools = importOptionChecking(m.(*FortiClient).Cfg, "pools")
		if pools == "" {
			return fmt.Errorf("Parameter pools is missing")
		}
		if err = d.Set("pools", pools); err != nil {
			return fmt.Errorf("Error set params pools: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["pools"] = pools

	o, err := c.ReadSystemIpamPoolsExclude(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpamPoolsExclude resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpamPoolsExclude(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpamPoolsExclude resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpamPoolsExcludeId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpamPoolsExcludeExcludeSubnet3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemIpamPoolsExclude(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemIpamPoolsExcludeId3rdl(o["ID"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["ID"], "SystemIpamPoolsExclude-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("exclude_subnet", flattenSystemIpamPoolsExcludeExcludeSubnet3rdl(o["exclude-subnet"], d, "exclude_subnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclude-subnet"], "SystemIpamPoolsExclude-ExcludeSubnet"); ok {
			if err = d.Set("exclude_subnet", vv); err != nil {
				return fmt.Errorf("Error reading exclude_subnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclude_subnet: %v", err)
		}
	}

	return nil
}

func flattenSystemIpamPoolsExcludeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpamPoolsExcludeId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamPoolsExcludeExcludeSubnet3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSystemIpamPoolsExclude(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemIpamPoolsExcludeId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ID"] = t
		}
	}

	if v, ok := d.GetOk("exclude_subnet"); ok || d.HasChange("exclude_subnet") {
		t, err := expandSystemIpamPoolsExcludeExcludeSubnet3rdl(d, v, "exclude_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-subnet"] = t
		}
	}

	return &obj, nil
}
