// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ACL actions.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAclIngressAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAclIngressActionUpdate,
		Read:   resourceSwitchControllerAclIngressActionRead,
		Update: resourceSwitchControllerAclIngressActionUpdate,
		Delete: resourceSwitchControllerAclIngressActionDelete,

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
			"ingress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fmgcount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerAclIngressActionUpdate(d *schema.ResourceData, m interface{}) error {
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
	ingress := d.Get("ingress").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ingress"] = ingress

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSwitchControllerAclIngressAction(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngressAction resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAclIngressAction(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngressAction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerAclIngressAction")

	return resourceSwitchControllerAclIngressActionRead(d, m)
}

func resourceSwitchControllerAclIngressActionDelete(d *schema.ResourceData, m interface{}) error {
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
	ingress := d.Get("ingress").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ingress"] = ingress

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSwitchControllerAclIngressAction(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAclIngressAction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAclIngressActionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ingress := d.Get("ingress").(string)
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
	if ingress == "" {
		ingress = importOptionChecking(m.(*FortiClient).Cfg, "ingress")
		if ingress == "" {
			return fmt.Errorf("Parameter ingress is missing")
		}
		if err = d.Set("ingress", ingress); err != nil {
			return fmt.Errorf("Error set params ingress: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ingress"] = ingress

	o, err := c.ReadSwitchControllerAclIngressAction(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngressAction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAclIngressAction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngressAction resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAclIngressActionCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressActionDrop2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAclIngressAction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgcount", flattenSwitchControllerAclIngressActionCount2edl(o["count"], d, "fmgcount")); err != nil {
		if vv, ok := fortiAPIPatch(o["count"], "SwitchControllerAclIngressAction-Count"); ok {
			if err = d.Set("fmgcount", vv); err != nil {
				return fmt.Errorf("Error reading fmgcount: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgcount: %v", err)
		}
	}

	if err = d.Set("drop", flattenSwitchControllerAclIngressActionDrop2edl(o["drop"], d, "drop")); err != nil {
		if vv, ok := fortiAPIPatch(o["drop"], "SwitchControllerAclIngressAction-Drop"); ok {
			if err = d.Set("drop", vv); err != nil {
				return fmt.Errorf("Error reading drop: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drop: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAclIngressActionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAclIngressActionCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressActionDrop2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAclIngressAction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgcount"); ok || d.HasChange("fmgcount") {
		t, err := expandSwitchControllerAclIngressActionCount2edl(d, v, "fmgcount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("drop"); ok || d.HasChange("drop") {
		t, err := expandSwitchControllerAclIngressActionDrop2edl(d, v, "drop")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop"] = t
		}
	}

	return &obj, nil
}
