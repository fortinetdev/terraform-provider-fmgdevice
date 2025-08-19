// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure DNS servers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDnsServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsServerCreate,
		Read:   resourceSystemDnsServerRead,
		Update: resourceSystemDnsServerUpdate,
		Delete: resourceSystemDnsServerDelete,

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
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"doh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"doh3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"doq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemDnsServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemDnsServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDnsServer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemDnsServer(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemDnsServer resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemDnsServerRead(d, m)
}

func resourceSystemDnsServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemDnsServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDnsServer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDnsServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemDnsServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemDnsServerRead(d, m)
}

func resourceSystemDnsServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemDnsServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDnsServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadSystemDnsServer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemDnsServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDnsServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDnsServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsServerDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemDnsServerDoh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsServerDoh3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsServerDoq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsServerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func refreshObjectSystemDnsServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dnsfilter_profile", flattenSystemDnsServerDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "SystemDnsServer-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("doh", flattenSystemDnsServerDoh(o["doh"], d, "doh")); err != nil {
		if vv, ok := fortiAPIPatch(o["doh"], "SystemDnsServer-Doh"); ok {
			if err = d.Set("doh", vv); err != nil {
				return fmt.Errorf("Error reading doh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading doh: %v", err)
		}
	}

	if err = d.Set("doh3", flattenSystemDnsServerDoh3(o["doh3"], d, "doh3")); err != nil {
		if vv, ok := fortiAPIPatch(o["doh3"], "SystemDnsServer-Doh3"); ok {
			if err = d.Set("doh3", vv); err != nil {
				return fmt.Errorf("Error reading doh3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading doh3: %v", err)
		}
	}

	if err = d.Set("doq", flattenSystemDnsServerDoq(o["doq"], d, "doq")); err != nil {
		if vv, ok := fortiAPIPatch(o["doq"], "SystemDnsServer-Doq"); ok {
			if err = d.Set("doq", vv); err != nil {
				return fmt.Errorf("Error reading doq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading doq: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemDnsServerMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "SystemDnsServer-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemDnsServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemDnsServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenSystemDnsServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDnsServerDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemDnsServerDoh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerDoh3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerDoq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func getObjectSystemDnsServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandSystemDnsServerDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("doh"); ok || d.HasChange("doh") {
		t, err := expandSystemDnsServerDoh(d, v, "doh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["doh"] = t
		}
	}

	if v, ok := d.GetOk("doh3"); ok || d.HasChange("doh3") {
		t, err := expandSystemDnsServerDoh3(d, v, "doh3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["doh3"] = t
		}
	}

	if v, ok := d.GetOk("doq"); ok || d.HasChange("doq") {
		t, err := expandSystemDnsServerDoq(d, v, "doq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["doq"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandSystemDnsServerMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemDnsServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
