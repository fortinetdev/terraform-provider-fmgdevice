// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure how log messages are displayed on the GUI.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogGuiDisplay() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogGuiDisplayUpdate,
		Read:   resourceLogGuiDisplayRead,
		Update: resourceLogGuiDisplayUpdate,
		Delete: resourceLogGuiDisplayDelete,

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
			"fortiview_unscanned_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resolve_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogGuiDisplayUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogGuiDisplay(d)
	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource while getting object: %v", err)
	}

	_, err = c.UpdateLogGuiDisplay(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogGuiDisplay")

	return resourceLogGuiDisplayRead(d, m)
}

func resourceLogGuiDisplayDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogGuiDisplay(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogGuiDisplay resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogGuiDisplayRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogGuiDisplay(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogGuiDisplay resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogGuiDisplay(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogGuiDisplay resource from API: %v", err)
	}
	return nil
}

func flattenLogGuiDisplayFortiviewUnscannedApps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogGuiDisplayResolveApps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogGuiDisplayResolveHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogGuiDisplay(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fortiview_unscanned_apps", flattenLogGuiDisplayFortiviewUnscannedApps(o["fortiview-unscanned-apps"], d, "fortiview_unscanned_apps")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview-unscanned-apps"], "LogGuiDisplay-FortiviewUnscannedApps"); ok {
			if err = d.Set("fortiview_unscanned_apps", vv); err != nil {
				return fmt.Errorf("Error reading fortiview_unscanned_apps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview_unscanned_apps: %v", err)
		}
	}

	if err = d.Set("resolve_apps", flattenLogGuiDisplayResolveApps(o["resolve-apps"], d, "resolve_apps")); err != nil {
		if vv, ok := fortiAPIPatch(o["resolve-apps"], "LogGuiDisplay-ResolveApps"); ok {
			if err = d.Set("resolve_apps", vv); err != nil {
				return fmt.Errorf("Error reading resolve_apps: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resolve_apps: %v", err)
		}
	}

	if err = d.Set("resolve_hosts", flattenLogGuiDisplayResolveHosts(o["resolve-hosts"], d, "resolve_hosts")); err != nil {
		if vv, ok := fortiAPIPatch(o["resolve-hosts"], "LogGuiDisplay-ResolveHosts"); ok {
			if err = d.Set("resolve_hosts", vv); err != nil {
				return fmt.Errorf("Error reading resolve_hosts: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resolve_hosts: %v", err)
		}
	}

	return nil
}

func flattenLogGuiDisplayFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogGuiDisplayFortiviewUnscannedApps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogGuiDisplayResolveApps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogGuiDisplayResolveHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogGuiDisplay(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fortiview_unscanned_apps"); ok || d.HasChange("fortiview_unscanned_apps") {
		t, err := expandLogGuiDisplayFortiviewUnscannedApps(d, v, "fortiview_unscanned_apps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview-unscanned-apps"] = t
		}
	}

	if v, ok := d.GetOk("resolve_apps"); ok || d.HasChange("resolve_apps") {
		t, err := expandLogGuiDisplayResolveApps(d, v, "resolve_apps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-apps"] = t
		}
	}

	if v, ok := d.GetOk("resolve_hosts"); ok || d.HasChange("resolve_hosts") {
		t, err := expandLogGuiDisplayResolveHosts(d, v, "resolve_hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-hosts"] = t
		}
	}

	return &obj, nil
}
