// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 SNMP managers (hosts).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerSnmpCommunityHosts6() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSnmpCommunityHosts6Create,
		Read:   resourceWirelessControllerSnmpCommunityHosts6Read,
		Update: resourceWirelessControllerSnmpCommunityHosts6Update,
		Delete: resourceWirelessControllerSnmpCommunityHosts6Delete,

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
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerSnmpCommunityHosts6Create(d *schema.ResourceData, m interface{}) error {
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
	community := d.Get("community").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community"] = community

	obj, err := getObjectWirelessControllerSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerSnmpCommunityHosts6 resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWirelessControllerSnmpCommunityHosts6(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWirelessControllerSnmpCommunityHosts6(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WirelessControllerSnmpCommunityHosts6 resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateWirelessControllerSnmpCommunityHosts6(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WirelessControllerSnmpCommunityHosts6 resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceWirelessControllerSnmpCommunityHosts6Read(d, m)
			} else {
				return fmt.Errorf("Error creating WirelessControllerSnmpCommunityHosts6 resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerSnmpCommunityHosts6Read(d, m)
}

func resourceWirelessControllerSnmpCommunityHosts6Update(d *schema.ResourceData, m interface{}) error {
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
	community := d.Get("community").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community"] = community

	obj, err := getObjectWirelessControllerSnmpCommunityHosts6(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmpCommunityHosts6 resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWirelessControllerSnmpCommunityHosts6(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmpCommunityHosts6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerSnmpCommunityHosts6Read(d, m)
}

func resourceWirelessControllerSnmpCommunityHosts6Delete(d *schema.ResourceData, m interface{}) error {
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
	community := d.Get("community").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community"] = community

	wsParams["adom"] = adomv

	err = c.DeleteWirelessControllerSnmpCommunityHosts6(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSnmpCommunityHosts6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSnmpCommunityHosts6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	community := d.Get("community").(string)
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
	if community == "" {
		community = importOptionChecking(m.(*FortiClient).Cfg, "community")
		if community == "" {
			return fmt.Errorf("Parameter community is missing")
		}
		if err = d.Set("community", community); err != nil {
			return fmt.Errorf("Error set params community: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["community"] = community

	o, err := c.ReadWirelessControllerSnmpCommunityHosts6(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WirelessControllerSnmpCommunityHosts6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSnmpCommunityHosts6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmpCommunityHosts6 resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSnmpCommunityHosts6Id3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityHosts6Ipv63rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSnmpCommunityHosts6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerSnmpCommunityHosts6Id3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerSnmpCommunityHosts6-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenWirelessControllerSnmpCommunityHosts6Ipv63rdl(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "WirelessControllerSnmpCommunityHosts6-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerSnmpCommunityHosts6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSnmpCommunityHosts6Id3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityHosts6Ipv63rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSnmpCommunityHosts6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerSnmpCommunityHosts6Id3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandWirelessControllerSnmpCommunityHosts6Ipv63rdl(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	return &obj, nil
}
