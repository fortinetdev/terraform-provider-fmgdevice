// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Show Internet Service application.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceUpdate,
		Read:   resourceFirewallInternetServiceRead,
		Update: resourceFirewallInternetServiceUpdate,
		Delete: resourceFirewallInternetServiceDelete,

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
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extra_ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"extra_ip6_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"icon_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip6_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"singularity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallInternetServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetService(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetService resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallInternetService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallInternetService")

	return resourceFirewallInternetServiceRead(d, m)
}

func resourceFirewallInternetServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallInternetService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetService(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallInternetService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetService resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtraIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtraIp6RangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIconId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIp6RangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceSingularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("database", flattenFirewallInternetServiceDatabase(o["database"], d, "database")); err != nil {
		if vv, ok := fortiAPIPatch(o["database"], "FirewallInternetService-Database"); ok {
			if err = d.Set("database", vv); err != nil {
				return fmt.Errorf("Error reading database: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("direction", flattenFirewallInternetServiceDirection(o["direction"], d, "direction")); err != nil {
		if vv, ok := fortiAPIPatch(o["direction"], "FirewallInternetService-Direction"); ok {
			if err = d.Set("direction", vv); err != nil {
				return fmt.Errorf("Error reading direction: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("extra_ip_range_number", flattenFirewallInternetServiceExtraIpRangeNumber(o["extra-ip-range-number"], d, "extra_ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["extra-ip-range-number"], "FirewallInternetService-ExtraIpRangeNumber"); ok {
			if err = d.Set("extra_ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
		}
	}

	if err = d.Set("extra_ip6_range_number", flattenFirewallInternetServiceExtraIp6RangeNumber(o["extra-ip6-range-number"], d, "extra_ip6_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["extra-ip6-range-number"], "FirewallInternetService-ExtraIp6RangeNumber"); ok {
			if err = d.Set("extra_ip6_range_number", vv); err != nil {
				return fmt.Errorf("Error reading extra_ip6_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extra_ip6_range_number: %v", err)
		}
	}

	if err = d.Set("icon_id", flattenFirewallInternetServiceIconId(o["icon-id"], d, "icon_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["icon-id"], "FirewallInternetService-IconId"); ok {
			if err = d.Set("icon_id", vv); err != nil {
				return fmt.Errorf("Error reading icon_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallInternetService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip_number", flattenFirewallInternetServiceIpNumber(o["ip-number"], d, "ip_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-number"], "FirewallInternetService-IpNumber"); ok {
			if err = d.Set("ip_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("ip_range_number", flattenFirewallInternetServiceIpRangeNumber(o["ip-range-number"], d, "ip_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-range-number"], "FirewallInternetService-IpRangeNumber"); ok {
			if err = d.Set("ip_range_number", vv); err != nil {
				return fmt.Errorf("Error reading ip_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("ip6_range_number", flattenFirewallInternetServiceIp6RangeNumber(o["ip6-range-number"], d, "ip6_range_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-range-number"], "FirewallInternetService-Ip6RangeNumber"); ok {
			if err = d.Set("ip6_range_number", vv); err != nil {
				return fmt.Errorf("Error reading ip6_range_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_range_number: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallInternetServiceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallInternetService-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenFirewallInternetServiceObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if vv, ok := fortiAPIPatch(o["obsolete"], "FirewallInternetService-Obsolete"); ok {
			if err = d.Set("obsolete", vv); err != nil {
				return fmt.Errorf("Error reading obsolete: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("singularity", flattenFirewallInternetServiceSingularity(o["singularity"], d, "singularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["singularity"], "FirewallInternetService-Singularity"); ok {
			if err = d.Set("singularity", vv); err != nil {
				return fmt.Errorf("Error reading singularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading singularity: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtraIpRangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtraIp6RangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIconId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpRangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIp6RangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceObsolete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceSingularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("database"); ok || d.HasChange("database") {
		t, err := expandFirewallInternetServiceDatabase(d, v, "database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok || d.HasChange("direction") {
		t, err := expandFirewallInternetServiceDirection(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("extra_ip_range_number"); ok || d.HasChange("extra_ip_range_number") {
		t, err := expandFirewallInternetServiceExtraIpRangeNumber(d, v, "extra_ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("extra_ip6_range_number"); ok || d.HasChange("extra_ip6_range_number") {
		t, err := expandFirewallInternetServiceExtraIp6RangeNumber(d, v, "extra_ip6_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-ip6-range-number"] = t
		}
	}

	if v, ok := d.GetOk("icon_id"); ok || d.HasChange("icon_id") {
		t, err := expandFirewallInternetServiceIconId(d, v, "icon_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-id"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallInternetServiceId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip_number"); ok || d.HasChange("ip_number") {
		t, err := expandFirewallInternetServiceIpNumber(d, v, "ip_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-number"] = t
		}
	}

	if v, ok := d.GetOk("ip_range_number"); ok || d.HasChange("ip_range_number") {
		t, err := expandFirewallInternetServiceIpRangeNumber(d, v, "ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range-number"] = t
		}
	}

	if v, ok := d.GetOk("ip6_range_number"); ok || d.HasChange("ip6_range_number") {
		t, err := expandFirewallInternetServiceIp6RangeNumber(d, v, "ip6_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-range-number"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallInternetServiceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok || d.HasChange("obsolete") {
		t, err := expandFirewallInternetServiceObsolete(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	if v, ok := d.GetOk("singularity"); ok || d.HasChange("singularity") {
		t, err := expandFirewallInternetServiceSingularity(d, v, "singularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["singularity"] = t
		}
	}

	return &obj, nil
}
