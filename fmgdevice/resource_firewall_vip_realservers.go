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

func resourceFirewallVipRealservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipRealserversCreate,
		Read:   resourceFirewallVipRealserversRead,
		Update: resourceFirewallVipRealserversUpdate,
		Delete: resourceFirewallVipRealserversDelete,

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
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"holddown_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			},
			"translate_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verify_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"health_check_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"seq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallVipRealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	obj, err := getObjectFirewallVipRealservers(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVipRealservers resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallVipRealservers(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallVipRealservers(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallVipRealservers resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallVipRealservers(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallVipRealservers resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallVipRealserversRead(d, m)
}

func resourceFirewallVipRealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	obj, err := getObjectFirewallVipRealservers(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipRealservers resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallVipRealservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipRealservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceFirewallVipRealserversRead(d, m)
}

func resourceFirewallVipRealserversDelete(d *schema.ResourceData, m interface{}) error {
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
	vip := d.Get("vip").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	wsParams["adom"] = adomv

	err = c.DeleteFirewallVipRealservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVipRealservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipRealserversRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	vip := d.Get("vip").(string)
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
	if vip == "" {
		vip = importOptionChecking(m.(*FortiClient).Cfg, "vip")
		if vip == "" {
			return fmt.Errorf("Parameter vip is missing")
		}
		if err = d.Set("vip", vip); err != nil {
			return fmt.Errorf("Error set params vip: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["vip"] = vip

	o, err := c.ReadFirewallVipRealservers(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallVipRealservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVipRealservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVipRealservers resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVipRealserversAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipRealserversClientIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipRealserversHealthcheck2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversHolddownInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversHttpHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversMaxConnections2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversMonitor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallVipRealserversPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversTranslateHost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversVerifyCert2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversHealthCheckProto2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipRealserversSeq2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVipRealservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("address", flattenFirewallVipRealserversAddress2edl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "FirewallVipRealservers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("client_ip", flattenFirewallVipRealserversClientIp2edl(o["client-ip"], d, "client_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-ip"], "FirewallVipRealservers-ClientIp"); ok {
			if err = d.Set("client_ip", vv); err != nil {
				return fmt.Errorf("Error reading client_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_ip: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenFirewallVipRealserversHealthcheck2edl(o["healthcheck"], d, "healthcheck")); err != nil {
		if vv, ok := fortiAPIPatch(o["healthcheck"], "FirewallVipRealservers-Healthcheck"); ok {
			if err = d.Set("healthcheck", vv); err != nil {
				return fmt.Errorf("Error reading healthcheck: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenFirewallVipRealserversHolddownInterval2edl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "FirewallVipRealservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenFirewallVipRealserversHttpHost2edl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "FirewallVipRealservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVipRealserversId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FirewallVipRealservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallVipRealserversIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FirewallVipRealservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenFirewallVipRealserversMaxConnections2edl(o["max-connections"], d, "max_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-connections"], "FirewallVipRealservers-MaxConnections"); ok {
			if err = d.Set("max_connections", vv); err != nil {
				return fmt.Errorf("Error reading max_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}

	if err = d.Set("monitor", flattenFirewallVipRealserversMonitor2edl(o["monitor"], d, "monitor")); err != nil {
		if vv, ok := fortiAPIPatch(o["monitor"], "FirewallVipRealservers-Monitor"); ok {
			if err = d.Set("monitor", vv); err != nil {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallVipRealserversPort2edl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FirewallVipRealservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallVipRealserversStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallVipRealservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenFirewallVipRealserversTranslateHost2edl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "FirewallVipRealservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallVipRealserversType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "FirewallVipRealservers-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenFirewallVipRealserversVerifyCert2edl(o["verify-cert"], d, "verify_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-cert"], "FirewallVipRealservers-VerifyCert"); ok {
			if err = d.Set("verify_cert", vv); err != nil {
				return fmt.Errorf("Error reading verify_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("weight", flattenFirewallVipRealserversWeight2edl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "FirewallVipRealservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("health_check_proto", flattenFirewallVipRealserversHealthCheckProto2edl(o["health-check-proto"], d, "health_check_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-proto"], "FirewallVipRealservers-HealthCheckProto"); ok {
			if err = d.Set("health_check_proto", vv); err != nil {
				return fmt.Errorf("Error reading health_check_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_proto: %v", err)
		}
	}

	if err = d.Set("seq", flattenFirewallVipRealserversSeq2edl(o["seq"], d, "seq")); err != nil {
		if vv, ok := fortiAPIPatch(o["seq"], "FirewallVipRealservers-Seq"); ok {
			if err = d.Set("seq", vv); err != nil {
				return fmt.Errorf("Error reading seq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading seq: %v", err)
		}
	}

	return nil
}

func flattenFirewallVipRealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVipRealserversAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipRealserversClientIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipRealserversHealthcheck2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHolddownInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHttpHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversMaxConnections2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversMonitor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallVipRealserversPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversTranslateHost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversVerifyCert2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHealthCheckProto2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversSeq2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVipRealservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandFirewallVipRealserversAddress2edl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("client_ip"); ok || d.HasChange("client_ip") {
		t, err := expandFirewallVipRealserversClientIp2edl(d, v, "client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-ip"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok || d.HasChange("healthcheck") {
		t, err := expandFirewallVipRealserversHealthcheck2edl(d, v, "healthcheck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandFirewallVipRealserversHolddownInterval2edl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandFirewallVipRealserversHttpHost2edl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandFirewallVipRealserversId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFirewallVipRealserversIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok || d.HasChange("max_connections") {
		t, err := expandFirewallVipRealserversMaxConnections2edl(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok || d.HasChange("monitor") {
		t, err := expandFirewallVipRealserversMonitor2edl(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFirewallVipRealserversPort2edl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallVipRealserversStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandFirewallVipRealserversTranslateHost2edl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandFirewallVipRealserversType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("verify_cert"); ok || d.HasChange("verify_cert") {
		t, err := expandFirewallVipRealserversVerifyCert2edl(d, v, "verify_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandFirewallVipRealserversWeight2edl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("health_check_proto"); ok || d.HasChange("health_check_proto") {
		t, err := expandFirewallVipRealserversHealthCheckProto2edl(d, v, "health_check_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-proto"] = t
		}
	}

	if v, ok := d.GetOk("seq"); ok || d.HasChange("seq") {
		t, err := expandFirewallVipRealserversSeq2edl(d, v, "seq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq"] = t
		}
	}

	return &obj, nil
}
