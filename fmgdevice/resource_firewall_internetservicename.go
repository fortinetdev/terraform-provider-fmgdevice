// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Define internet service names.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceName() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceNameCreate,
		Read:   resourceFirewallInternetServiceNameRead,
		Update: resourceFirewallInternetServiceNameUpdate,
		Delete: resourceFirewallInternetServiceNameDelete,

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
			"city_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"country_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallInternetServiceNameCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceName(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceName resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallInternetServiceName(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallInternetServiceName(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallInternetServiceName resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallInternetServiceName(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallInternetServiceName resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallInternetServiceNameRead(d, m)
}

func resourceFirewallInternetServiceNameUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceName(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceName resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInternetServiceName(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallInternetServiceNameRead(d, m)
}

func resourceFirewallInternetServiceNameDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallInternetServiceName(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetServiceName(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInternetServiceName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceName(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceName resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceNameCityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceNameCountryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceNameInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceNameRegionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceNameType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceName(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("city_id", flattenFirewallInternetServiceNameCityId(o["city-id"], d, "city_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["city-id"], "FirewallInternetServiceName-CityId"); ok {
			if err = d.Set("city_id", vv); err != nil {
				return fmt.Errorf("Error reading city_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading city_id: %v", err)
		}
	}

	if err = d.Set("country_id", flattenFirewallInternetServiceNameCountryId(o["country-id"], d, "country_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["country-id"], "FirewallInternetServiceName-CountryId"); ok {
			if err = d.Set("country_id", vv); err != nil {
				return fmt.Errorf("Error reading country_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading country_id: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenFirewallInternetServiceNameInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["internet-service-id"], "FirewallInternetServiceName-InternetServiceId"); ok {
			if err = d.Set("internet_service_id", vv); err != nil {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallInternetServiceNameName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallInternetServiceName-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("region_id", flattenFirewallInternetServiceNameRegionId(o["region-id"], d, "region_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["region-id"], "FirewallInternetServiceName-RegionId"); ok {
			if err = d.Set("region_id", vv); err != nil {
				return fmt.Errorf("Error reading region_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading region_id: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallInternetServiceNameType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallInternetServiceName-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceNameCityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameCountryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameRegionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceName(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("city_id"); ok || d.HasChange("city_id") {
		t, err := expandFirewallInternetServiceNameCityId(d, v, "city_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city-id"] = t
		}
	}

	if v, ok := d.GetOk("country_id"); ok || d.HasChange("country_id") {
		t, err := expandFirewallInternetServiceNameCountryId(d, v, "country_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandFirewallInternetServiceNameInternetServiceId(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallInternetServiceNameName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("region_id"); ok || d.HasChange("region_id") {
		t, err := expandFirewallInternetServiceNameRegionId(d, v, "region_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallInternetServiceNameType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
