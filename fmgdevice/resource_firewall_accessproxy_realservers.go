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

func resourceFirewallAccessProxyRealservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyRealserversCreate,
		Read:   resourceFirewallAccessProxyRealserversRead,
		Update: resourceFirewallAccessProxyRealserversUpdate,
		Delete: resourceFirewallAccessProxyRealserversDelete,

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
			"access_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallAccessProxyRealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	obj, err := getObjectFirewallAccessProxyRealservers(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxyRealservers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallAccessProxyRealservers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallAccessProxyRealservers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallAccessProxyRealservers resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallAccessProxyRealservers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallAccessProxyRealservers resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxyRealserversRead(d, m)
}

func resourceFirewallAccessProxyRealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	obj, err := getObjectFirewallAccessProxyRealservers(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyRealservers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallAccessProxyRealservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxyRealservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallAccessProxyRealserversRead(d, m)
}

func resourceFirewallAccessProxyRealserversDelete(d *schema.ResourceData, m interface{}) error {
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
	access_proxy := d.Get("access_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	wsParams["adom"] = adomv

	err = c.DeleteFirewallAccessProxyRealservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxyRealservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyRealserversRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	access_proxy := d.Get("access_proxy").(string)
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
	if access_proxy == "" {
		access_proxy = importOptionChecking(m.(*FortiClient).Cfg, "access_proxy")
		if access_proxy == "" {
			return fmt.Errorf("Parameter access_proxy is missing")
		}
		if err = d.Set("access_proxy", access_proxy); err != nil {
			return fmt.Errorf("Error set params access_proxy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["access_proxy"] = access_proxy

	o, err := c.ReadFirewallAccessProxyRealservers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallAccessProxyRealservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxyRealservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxyRealservers resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyRealserversId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAccessProxyRealserversWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxyRealservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFirewallAccessProxyRealserversId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallAccessProxyRealservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallAccessProxyRealserversIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallAccessProxyRealservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallAccessProxyRealserversPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallAccessProxyRealservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallAccessProxyRealserversStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallAccessProxyRealservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("weight", flattenFirewallAccessProxyRealserversWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "FirewallAccessProxyRealservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenFirewallAccessProxyRealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallAccessProxyRealserversId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyRealserversWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxyRealservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallAccessProxyRealserversId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallAccessProxyRealserversIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallAccessProxyRealserversPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallAccessProxyRealserversStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandFirewallAccessProxyRealserversWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
