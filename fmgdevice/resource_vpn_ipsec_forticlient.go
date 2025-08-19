// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiClient policy realm.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecForticlient() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecForticlientCreate,
		Read:   resourceVpnIpsecForticlientRead,
		Update: resourceVpnIpsecForticlientUpdate,
		Delete: resourceVpnIpsecForticlientDelete,

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
			"phase2name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usergroupname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecForticlientCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnIpsecForticlient(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecForticlient resource while getting object: %v", err)
	}

	_, err = c.CreateVpnIpsecForticlient(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecForticlient resource: %v", err)
	}

	d.SetId(getStringKey(d, "realm"))

	return resourceVpnIpsecForticlientRead(d, m)
}

func resourceVpnIpsecForticlientUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectVpnIpsecForticlient(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecForticlient resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnIpsecForticlient(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecForticlient resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "realm"))

	return resourceVpnIpsecForticlientRead(d, m)
}

func resourceVpnIpsecForticlientDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteVpnIpsecForticlient(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecForticlient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecForticlientRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecForticlient(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecForticlient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecForticlient(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecForticlient resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecForticlientPhase2Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenVpnIpsecForticlientRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecForticlientStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecForticlientUsergroupname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectVpnIpsecForticlient(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("phase2name", flattenVpnIpsecForticlientPhase2Name(o["phase2name"], d, "phase2name")); err != nil {
		if vv, ok := fortiAPIPatch(o["phase2name"], "VpnIpsecForticlient-Phase2Name"); ok {
			if err = d.Set("phase2name", vv); err != nil {
				return fmt.Errorf("Error reading phase2name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading phase2name: %v", err)
		}
	}

	if err = d.Set("realm", flattenVpnIpsecForticlientRealm(o["realm"], d, "realm")); err != nil {
		if vv, ok := fortiAPIPatch(o["realm"], "VpnIpsecForticlient-Realm"); ok {
			if err = d.Set("realm", vv); err != nil {
				return fmt.Errorf("Error reading realm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnIpsecForticlientStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "VpnIpsecForticlient-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("usergroupname", flattenVpnIpsecForticlientUsergroupname(o["usergroupname"], d, "usergroupname")); err != nil {
		if vv, ok := fortiAPIPatch(o["usergroupname"], "VpnIpsecForticlient-Usergroupname"); ok {
			if err = d.Set("usergroupname", vv); err != nil {
				return fmt.Errorf("Error reading usergroupname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usergroupname: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecForticlientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecForticlientPhase2Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandVpnIpsecForticlientRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecForticlientStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecForticlientUsergroupname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectVpnIpsecForticlient(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("phase2name"); ok || d.HasChange("phase2name") {
		t, err := expandVpnIpsecForticlientPhase2Name(d, v, "phase2name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phase2name"] = t
		}
	}

	if v, ok := d.GetOk("realm"); ok || d.HasChange("realm") {
		t, err := expandVpnIpsecForticlientRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandVpnIpsecForticlientStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("usergroupname"); ok || d.HasChange("usergroupname") {
		t, err := expandVpnIpsecForticlientUsergroupname(d, v, "usergroupname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroupname"] = t
		}
	}

	return &obj, nil
}
