// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 SNMP managers (hosts).

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSnmpCommunityHosts() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpCommunityHostsCreate,
		Read:   resourceSwitchControllerSnmpCommunityHostsRead,
		Update: resourceSwitchControllerSnmpCommunityHostsUpdate,
		Delete: resourceSwitchControllerSnmpCommunityHostsDelete,

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
			"snmp_community": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSnmpCommunityHostsCreate(d *schema.ResourceData, m interface{}) error {
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
	snmp_community := d.Get("snmp_community").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["snmp_community"] = snmp_community

	obj, err := getObjectSwitchControllerSnmpCommunityHosts(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSnmpCommunityHosts resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerSnmpCommunityHosts(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSnmpCommunityHosts resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerSnmpCommunityHostsRead(d, m)
}

func resourceSwitchControllerSnmpCommunityHostsUpdate(d *schema.ResourceData, m interface{}) error {
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
	snmp_community := d.Get("snmp_community").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["snmp_community"] = snmp_community

	obj, err := getObjectSwitchControllerSnmpCommunityHosts(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpCommunityHosts resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpCommunityHosts(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpCommunityHosts resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerSnmpCommunityHostsRead(d, m)
}

func resourceSwitchControllerSnmpCommunityHostsDelete(d *schema.ResourceData, m interface{}) error {
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
	snmp_community := d.Get("snmp_community").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["snmp_community"] = snmp_community

	err = c.DeleteSwitchControllerSnmpCommunityHosts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSnmpCommunityHosts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpCommunityHostsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	snmp_community := d.Get("snmp_community").(string)
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
	if snmp_community == "" {
		snmp_community = importOptionChecking(m.(*FortiClient).Cfg, "snmp_community")
		if snmp_community == "" {
			return fmt.Errorf("Parameter snmp_community is missing")
		}
		if err = d.Set("snmp_community", snmp_community); err != nil {
			return fmt.Errorf("Error set params snmp_community: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["snmp_community"] = snmp_community

	o, err := c.ReadSwitchControllerSnmpCommunityHosts(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpCommunityHosts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpCommunityHosts(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpCommunityHosts resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpCommunityHostsId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSnmpCommunityHostsIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSwitchControllerSnmpCommunityHosts(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSwitchControllerSnmpCommunityHostsId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerSnmpCommunityHosts-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerSnmpCommunityHostsIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "SwitchControllerSnmpCommunityHosts-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpCommunityHostsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSnmpCommunityHostsId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpCommunityHostsIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func getObjectSwitchControllerSnmpCommunityHosts(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerSnmpCommunityHostsId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandSwitchControllerSnmpCommunityHostsIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
