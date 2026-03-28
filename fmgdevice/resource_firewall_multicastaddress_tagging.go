// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Config object tagging.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallMulticastAddressTagging() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallMulticastAddressTaggingCreate,
		Read:   resourceFirewallMulticastAddressTaggingRead,
		Update: resourceFirewallMulticastAddressTaggingUpdate,
		Delete: resourceFirewallMulticastAddressTaggingDelete,

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
			"multicast_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallMulticastAddressTaggingCreate(d *schema.ResourceData, m interface{}) error {
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
	multicast_address := d.Get("multicast_address").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address"] = multicast_address

	obj, err := getObjectFirewallMulticastAddressTagging(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallMulticastAddressTagging resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallMulticastAddressTagging(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallMulticastAddressTagging(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallMulticastAddressTagging resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallMulticastAddressTagging(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallMulticastAddressTagging resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallMulticastAddressTaggingRead(d, m)
}

func resourceFirewallMulticastAddressTaggingUpdate(d *schema.ResourceData, m interface{}) error {
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
	multicast_address := d.Get("multicast_address").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address"] = multicast_address

	obj, err := getObjectFirewallMulticastAddressTagging(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastAddressTagging resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallMulticastAddressTagging(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastAddressTagging resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallMulticastAddressTaggingRead(d, m)
}

func resourceFirewallMulticastAddressTaggingDelete(d *schema.ResourceData, m interface{}) error {
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
	multicast_address := d.Get("multicast_address").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address"] = multicast_address

	wsParams["adom"] = adomv

	err = c.DeleteFirewallMulticastAddressTagging(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallMulticastAddressTagging resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallMulticastAddressTaggingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	multicast_address := d.Get("multicast_address").(string)
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
	if multicast_address == "" {
		multicast_address = importOptionChecking(m.(*FortiClient).Cfg, "multicast_address")
		if multicast_address == "" {
			return fmt.Errorf("Parameter multicast_address is missing")
		}
		if err = d.Set("multicast_address", multicast_address); err != nil {
			return fmt.Errorf("Error set params multicast_address: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address"] = multicast_address

	o, err := c.ReadFirewallMulticastAddressTagging(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallMulticastAddressTagging resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallMulticastAddressTagging(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallMulticastAddressTagging resource from API: %v", err)
	}
	return nil
}

func flattenFirewallMulticastAddressTaggingCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastAddressTaggingName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddressTaggingTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallMulticastAddressTagging(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenFirewallMulticastAddressTaggingCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "FirewallMulticastAddressTagging-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallMulticastAddressTaggingName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallMulticastAddressTagging-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("tags", flattenFirewallMulticastAddressTaggingTags2edl(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "FirewallMulticastAddressTagging-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	return nil
}

func flattenFirewallMulticastAddressTaggingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallMulticastAddressTaggingCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastAddressTaggingName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddressTaggingTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallMulticastAddressTagging(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallMulticastAddressTaggingCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallMulticastAddressTaggingName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandFirewallMulticastAddressTaggingTags2edl(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	return &obj, nil
}
