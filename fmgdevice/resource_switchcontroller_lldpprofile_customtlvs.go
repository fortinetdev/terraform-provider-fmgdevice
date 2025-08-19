// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configuration method to edit custom TLV entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLldpProfileCustomTlvs() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLldpProfileCustomTlvsCreate,
		Read:   resourceSwitchControllerLldpProfileCustomTlvsRead,
		Update: resourceSwitchControllerLldpProfileCustomTlvsUpdate,
		Delete: resourceSwitchControllerLldpProfileCustomTlvsDelete,

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
			"lldp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"information_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"oui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subtype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerLldpProfileCustomTlvsCreate(d *schema.ResourceData, m interface{}) error {
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
	lldp_profile := d.Get("lldp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["lldp_profile"] = lldp_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerLldpProfileCustomTlvs(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfileCustomTlvs resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerLldpProfileCustomTlvs(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerLldpProfileCustomTlvsRead(d, m)
}

func resourceSwitchControllerLldpProfileCustomTlvsUpdate(d *schema.ResourceData, m interface{}) error {
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
	lldp_profile := d.Get("lldp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["lldp_profile"] = lldp_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerLldpProfileCustomTlvs(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfileCustomTlvs resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerLldpProfileCustomTlvs(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerLldpProfileCustomTlvsRead(d, m)
}

func resourceSwitchControllerLldpProfileCustomTlvsDelete(d *schema.ResourceData, m interface{}) error {
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
	lldp_profile := d.Get("lldp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["lldp_profile"] = lldp_profile

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerLldpProfileCustomTlvs(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpProfileCustomTlvsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	lldp_profile := d.Get("lldp_profile").(string)
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
	if lldp_profile == "" {
		lldp_profile = importOptionChecking(m.(*FortiClient).Cfg, "lldp_profile")
		if lldp_profile == "" {
			return fmt.Errorf("Parameter lldp_profile is missing")
		}
		if err = d.Set("lldp_profile", lldp_profile); err != nil {
			return fmt.Errorf("Error set params lldp_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["lldp_profile"] = lldp_profile

	o, err := c.ReadSwitchControllerLldpProfileCustomTlvs(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfileCustomTlvs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpProfileCustomTlvs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfileCustomTlvs resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpProfileCustomTlvsInformationString2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsOui2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsSubtype2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("information_string", flattenSwitchControllerLldpProfileCustomTlvsInformationString2edl(o["information-string"], d, "information_string")); err != nil {
		if vv, ok := fortiAPIPatch(o["information-string"], "SwitchControllerLldpProfileCustomTlvs-InformationString"); ok {
			if err = d.Set("information_string", vv); err != nil {
				return fmt.Errorf("Error reading information_string: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading information_string: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerLldpProfileCustomTlvsName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerLldpProfileCustomTlvs-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("oui", flattenSwitchControllerLldpProfileCustomTlvsOui2edl(o["oui"], d, "oui")); err != nil {
		if vv, ok := fortiAPIPatch(o["oui"], "SwitchControllerLldpProfileCustomTlvs-Oui"); ok {
			if err = d.Set("oui", vv); err != nil {
				return fmt.Errorf("Error reading oui: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oui: %v", err)
		}
	}

	if err = d.Set("subtype", flattenSwitchControllerLldpProfileCustomTlvsSubtype2edl(o["subtype"], d, "subtype")); err != nil {
		if vv, ok := fortiAPIPatch(o["subtype"], "SwitchControllerLldpProfileCustomTlvs-Subtype"); ok {
			if err = d.Set("subtype", vv); err != nil {
				return fmt.Errorf("Error reading subtype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subtype: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerLldpProfileCustomTlvsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLldpProfileCustomTlvsInformationString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsOui2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsSubtype2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("information_string"); ok || d.HasChange("information_string") {
		t, err := expandSwitchControllerLldpProfileCustomTlvsInformationString2edl(d, v, "information_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["information-string"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerLldpProfileCustomTlvsName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oui"); ok || d.HasChange("oui") {
		t, err := expandSwitchControllerLldpProfileCustomTlvsOui2edl(d, v, "oui")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oui"] = t
		}
	}

	if v, ok := d.GetOk("subtype"); ok || d.HasChange("subtype") {
		t, err := expandSwitchControllerLldpProfileCustomTlvsSubtype2edl(d, v, "subtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subtype"] = t
		}
	}

	return &obj, nil
}
