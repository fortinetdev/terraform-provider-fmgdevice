// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterDomainListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterDomainListEntriesCreate,
		Read:   resourceWebfilterDomainListEntriesRead,
		Update: resourceWebfilterDomainListEntriesUpdate,
		Delete: resourceWebfilterDomainListEntriesDelete,

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
			"domain_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWebfilterDomainListEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	domain_list := d.Get("domain_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_list"] = domain_list

	obj, err := getObjectWebfilterDomainListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterDomainListEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("domain")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebfilterDomainListEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebfilterDomainListEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebfilterDomainListEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebfilterDomainListEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebfilterDomainListEntries resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "domain"))

	return resourceWebfilterDomainListEntriesRead(d, m)
}

func resourceWebfilterDomainListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	domain_list := d.Get("domain_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_list"] = domain_list

	obj, err := getObjectWebfilterDomainListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterDomainListEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebfilterDomainListEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterDomainListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "domain"))

	return resourceWebfilterDomainListEntriesRead(d, m)
}

func resourceWebfilterDomainListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	domain_list := d.Get("domain_list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_list"] = domain_list

	wsParams["adom"] = adomv

	err = c.DeleteWebfilterDomainListEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterDomainListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterDomainListEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	domain_list := d.Get("domain_list").(string)
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
	if domain_list == "" {
		domain_list = importOptionChecking(m.(*FortiClient).Cfg, "domain_list")
		if domain_list == "" {
			return fmt.Errorf("Parameter domain_list is missing")
		}
		if err = d.Set("domain_list", domain_list); err != nil {
			return fmt.Errorf("Error set params domain_list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["domain_list"] = domain_list

	o, err := c.ReadWebfilterDomainListEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebfilterDomainListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterDomainListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterDomainListEntries resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterDomainListEntriesDomain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterDomainListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("domain", flattenWebfilterDomainListEntriesDomain2edl(o["domain"], d, "domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain"], "WebfilterDomainListEntries-Domain"); ok {
			if err = d.Set("domain", vv); err != nil {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	return nil
}

func flattenWebfilterDomainListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterDomainListEntriesDomain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterDomainListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain"); ok || d.HasChange("domain") {
		t, err := expandWebfilterDomainListEntriesDomain2edl(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	return &obj, nil
}
