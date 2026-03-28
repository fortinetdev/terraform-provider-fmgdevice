// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Select the real servers that this server load balancing VIP will distribute traffic to.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallVip6Realservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVip6RealserversCreate,
		Read:   resourceFirewallVip6RealserversRead,
		Update: resourceFirewallVip6RealserversUpdate,
		Delete: resourceFirewallVip6RealserversDelete,

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
			"vip6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"holddown_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_connections": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"translate_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallVip6RealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	obj, err := getObjectFirewallVip6Realservers(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip6Realservers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVip6Realservers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVip6Realservers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVip6Realservers resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallVip6Realservers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVip6Realservers resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallVip6RealserversRead(d, m)
}

func resourceFirewallVip6RealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	obj, err := getObjectFirewallVip6Realservers(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6Realservers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVip6Realservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6Realservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallVip6RealserversRead(d, m)
}

func resourceFirewallVip6RealserversDelete(d *schema.ResourceData, m interface{}) error {
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
	vip6 := d.Get("vip6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	wsParams["adom"] = adomv

	err = c.DeleteFirewallVip6Realservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip6Realservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip6RealserversRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vip6 := d.Get("vip6").(string)
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
	if vip6 == "" {
		vip6 = importOptionChecking(m.(*FortiClient).Cfg, "vip6")
		if vip6 == "" {
			return fmt.Errorf("Parameter vip6 is missing")
		}
		if err = d.Set("vip6", vip6); err != nil {
			return fmt.Errorf("Error set params vip6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip6"] = vip6

	o, err := c.ReadFirewallVip6Realservers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVip6Realservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip6Realservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip6Realservers resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVip6RealserversClientIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHealthcheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHolddownInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHttpHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversMaxConnections2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVip6RealserversPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversTranslateHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversVerifyCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip6RealserversWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVip6Realservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("client_ip", flattenFirewallVip6RealserversClientIp2edl(o["client-ip"], d, "client_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-ip"], "FirewallVip6Realservers-ClientIp"); ok {
			if err = d.Set("client_ip", vv); err != nil {
				return fmt.Errorf("Error reading client_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_ip: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenFirewallVip6RealserversHealthcheck2edl(o["healthcheck"], d, "healthcheck")); err != nil {
		if vv, ok := fortiAPIPatch(o["healthcheck"], "FirewallVip6Realservers-Healthcheck"); ok {
			if err = d.Set("healthcheck", vv); err != nil {
				return fmt.Errorf("Error reading healthcheck: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenFirewallVip6RealserversHolddownInterval2edl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "FirewallVip6Realservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenFirewallVip6RealserversHttpHost2edl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "FirewallVip6Realservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVip6RealserversId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallVip6Realservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallVip6RealserversIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallVip6Realservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenFirewallVip6RealserversMaxConnections2edl(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "FirewallVip6Realservers-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("monitor", flattenFirewallVip6RealserversMonitor2edl(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "FirewallVip6Realservers-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallVip6RealserversPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallVip6Realservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallVip6RealserversStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallVip6Realservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenFirewallVip6RealserversTranslateHost2edl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "FirewallVip6Realservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenFirewallVip6RealserversVerifyCert2edl(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "FirewallVip6Realservers-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("weight", flattenFirewallVip6RealserversWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "FirewallVip6Realservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenFirewallVip6RealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVip6RealserversClientIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHealthcheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHolddownInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHttpHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversMaxConnections2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVip6RealserversPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversTranslateHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversVerifyCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip6Realservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("client_ip"); ok || d.HasChange("client_ip") {
		t, err := expandFirewallVip6RealserversClientIp2edl(d, v, "client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-ip"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok || d.HasChange("healthcheck") {
		t, err := expandFirewallVip6RealserversHealthcheck2edl(d, v, "healthcheck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandFirewallVip6RealserversHolddownInterval2edl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandFirewallVip6RealserversHttpHost2edl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallVip6RealserversId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallVip6RealserversIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandFirewallVip6RealserversMaxConnections2edl(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandFirewallVip6RealserversMonitor2edl(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallVip6RealserversPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallVip6RealserversStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandFirewallVip6RealserversTranslateHost2edl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandFirewallVip6RealserversVerifyCert2edl(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandFirewallVip6RealserversWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
