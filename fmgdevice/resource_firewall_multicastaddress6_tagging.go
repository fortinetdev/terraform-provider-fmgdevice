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

func resourceFirewallMulticastAddress6Tagging() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallMulticastAddress6TaggingCreate,
		Read:   resourceFirewallMulticastAddress6TaggingRead,
		Update: resourceFirewallMulticastAddress6TaggingUpdate,
		Delete: resourceFirewallMulticastAddress6TaggingDelete,

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
			"multicast_address6": &schema.Schema{
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

func resourceFirewallMulticastAddress6TaggingCreate(d *schema.ResourceData, m interface{}) error {
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
	multicast_address6 := d.Get("multicast_address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address6"] = multicast_address6

	obj, err := getObjectFirewallMulticastAddress6Tagging(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallMulticastAddress6Tagging resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallMulticastAddress6Tagging(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallMulticastAddress6Tagging(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallMulticastAddress6Tagging resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallMulticastAddress6Tagging(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallMulticastAddress6Tagging resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallMulticastAddress6TaggingRead(d, m)
}

func resourceFirewallMulticastAddress6TaggingUpdate(d *schema.ResourceData, m interface{}) error {
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
	multicast_address6 := d.Get("multicast_address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address6"] = multicast_address6

	obj, err := getObjectFirewallMulticastAddress6Tagging(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastAddress6Tagging resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallMulticastAddress6Tagging(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallMulticastAddress6Tagging resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallMulticastAddress6TaggingRead(d, m)
}

func resourceFirewallMulticastAddress6TaggingDelete(d *schema.ResourceData, m interface{}) error {
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
	multicast_address6 := d.Get("multicast_address6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address6"] = multicast_address6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallMulticastAddress6Tagging(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallMulticastAddress6Tagging resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallMulticastAddress6TaggingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	multicast_address6 := d.Get("multicast_address6").(string)
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
	if multicast_address6 == "" {
		multicast_address6 = importOptionChecking(m.(*FortiClient).Cfg, "multicast_address6")
		if multicast_address6 == "" {
			return fmt.Errorf("Parameter multicast_address6 is missing")
		}
		if err = d.Set("multicast_address6", multicast_address6); err != nil {
			return fmt.Errorf("Error set params multicast_address6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["multicast_address6"] = multicast_address6

	o, err := c.ReadFirewallMulticastAddress6Tagging(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallMulticastAddress6Tagging resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallMulticastAddress6Tagging(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallMulticastAddress6Tagging resource from API: %v", err)
	}
	return nil
}

func flattenFirewallMulticastAddress6TaggingCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallMulticastAddress6TaggingName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallMulticastAddress6TaggingTags2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallMulticastAddress6Tagging(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("category", flattenFirewallMulticastAddress6TaggingCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "FirewallMulticastAddress6Tagging-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallMulticastAddress6TaggingName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallMulticastAddress6Tagging-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("tags", flattenFirewallMulticastAddress6TaggingTags2edl(o["tags"], d, "tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["tags"], "FirewallMulticastAddress6Tagging-Tags"); ok {
			if err = d.Set("tags", vv); err != nil {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tags: %v", err)
		}
	}

	return nil
}

func flattenFirewallMulticastAddress6TaggingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallMulticastAddress6TaggingCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallMulticastAddress6TaggingName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallMulticastAddress6TaggingTags2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallMulticastAddress6Tagging(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallMulticastAddress6TaggingCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallMulticastAddress6TaggingName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		t, err := expandFirewallMulticastAddress6TaggingTags2edl(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	return &obj, nil
}
