// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: MD5 key.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfAreaVirtualLinkMd5Keys() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfAreaVirtualLinkMd5KeysCreate,
		Read:   resourceRouterOspfAreaVirtualLinkMd5KeysRead,
		Update: resourceRouterOspfAreaVirtualLinkMd5KeysUpdate,
		Delete: resourceRouterOspfAreaVirtualLinkMd5KeysDelete,

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
			"area": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_link": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"key_string": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func resourceRouterOspfAreaVirtualLinkMd5KeysCreate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	obj, err := getObjectRouterOspfAreaVirtualLinkMd5Keys(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfAreaVirtualLinkMd5Keys resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfAreaVirtualLinkMd5Keys(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterOspfAreaVirtualLinkMd5Keys resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfAreaVirtualLinkMd5KeysRead(d, m)
}

func resourceRouterOspfAreaVirtualLinkMd5KeysUpdate(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	obj, err := getObjectRouterOspfAreaVirtualLinkMd5Keys(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfAreaVirtualLinkMd5Keys resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfAreaVirtualLinkMd5Keys(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfAreaVirtualLinkMd5Keys resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceRouterOspfAreaVirtualLinkMd5KeysRead(d, m)
}

func resourceRouterOspfAreaVirtualLinkMd5KeysDelete(d *schema.ResourceData, m interface{}) error {
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
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	err = c.DeleteRouterOspfAreaVirtualLinkMd5Keys(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfAreaVirtualLinkMd5Keys resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfAreaVirtualLinkMd5KeysRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	area := d.Get("area").(string)
	virtual_link := d.Get("virtual_link").(string)
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
	if area == "" {
		area = importOptionChecking(m.(*FortiClient).Cfg, "area")
		if area == "" {
			return fmt.Errorf("Parameter area is missing")
		}
		if err = d.Set("area", area); err != nil {
			return fmt.Errorf("Error set params area: %v", err)
		}
	}
	if virtual_link == "" {
		virtual_link = importOptionChecking(m.(*FortiClient).Cfg, "virtual_link")
		if virtual_link == "" {
			return fmt.Errorf("Parameter virtual_link is missing")
		}
		if err = d.Set("virtual_link", virtual_link); err != nil {
			return fmt.Errorf("Error set params virtual_link: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["area"] = area
	paradict["virtual_link"] = virtual_link

	o, err := c.ReadRouterOspfAreaVirtualLinkMd5Keys(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfAreaVirtualLinkMd5Keys resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfAreaVirtualLinkMd5Keys(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfAreaVirtualLinkMd5Keys resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAreaVirtualLinkMd5KeysId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfAreaVirtualLinkMd5Keys(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenRouterOspfAreaVirtualLinkMd5KeysId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "RouterOspfAreaVirtualLinkMd5Keys-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenRouterOspfAreaVirtualLinkMd5KeysFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfAreaVirtualLinkMd5KeysId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysKeyString4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectRouterOspfAreaVirtualLinkMd5Keys(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandRouterOspfAreaVirtualLinkMd5KeysId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("key_string"); ok || d.HasChange("key_string") {
		t, err := expandRouterOspfAreaVirtualLinkMd5KeysKeyString4thl(d, v, "key_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-string"] = t
		}
	}

	return &obj, nil
}
