// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Show vendor and the MAC address they have.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVendorMac() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVendorMacUpdate,
		Read:   resourceFirewallVendorMacRead,
		Update: resourceFirewallVendorMacUpdate,
		Delete: resourceFirewallVendorMacDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"adom": &schema.Schema{
				Type:	 schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
						"device_name": &schema.Schema{
				Type:	 schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
					"fosid": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
			},
			"mac_number": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
			},
			"obsolete": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}


func resourceFirewallVendorMacUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallVendorMac(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVendorMac resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVendorMac(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVendorMac resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallVendorMac")

	return resourceFirewallVendorMacRead(d, m)
}

func resourceFirewallVendorMacDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallVendorMac(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVendorMac resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVendorMacRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallVendorMac(mkey, paradict)
	if err != nil {
	d.SetId("")
		return fmt.Errorf("Error reading FirewallVendorMac resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVendorMac(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVendorMac resource from API: %v", err)
	}
	return nil
}


func flattenFirewallVendorMacId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVendorMacMacNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVendorMacName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVendorMacObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}



func refreshObjectFirewallVendorMac(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("fosid", flattenFirewallVendorMacId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallVendorMac-Id"); ok {
            if err = d.Set("fosid", vv); err != nil {
                return fmt.Errorf("Error reading fosid: %v", err)
            }
        } else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("mac_number", flattenFirewallVendorMacMacNumber(o["mac-number"], d, "mac_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-number"], "FirewallVendorMac-MacNumber"); ok {
            if err = d.Set("mac_number", vv); err != nil {
                return fmt.Errorf("Error reading mac_number: %v", err)
            }
        } else {
			return fmt.Errorf("Error reading mac_number: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallVendorMacName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallVendorMac-Name"); ok {
            if err = d.Set("name", vv); err != nil {
                return fmt.Errorf("Error reading name: %v", err)
            }
        } else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenFirewallVendorMacObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if vv, ok := fortiAPIPatch(o["obsolete"], "FirewallVendorMac-Obsolete"); ok {
            if err = d.Set("obsolete", vv); err != nil {
                return fmt.Errorf("Error reading obsolete: %v", err)
            }
        } else {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}


	return nil
}

func flattenFirewallVendorMacFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandFirewallVendorMacId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVendorMacMacNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVendorMacName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVendorMacObsolete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallVendorMac(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})
	

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallVendorMacId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("mac_number"); ok || d.HasChange("mac_number") {
		t, err := expandFirewallVendorMacMacNumber(d, v, "mac_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-number"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallVendorMacName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok || d.HasChange("obsolete") {
		t, err := expandFirewallVendorMacObsolete(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}


	return &obj, nil
}

