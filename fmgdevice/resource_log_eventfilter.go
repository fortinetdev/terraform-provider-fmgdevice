// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure log event filters.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogEventfilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogEventfilterUpdate,
		Read:   resourceLogEventfilterRead,
		Update: resourceLogEventfilterUpdate,
		Delete: resourceLogEventfilterDelete,

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
			"cifs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"connector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiextender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webproxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_activity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogEventfilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLogEventfilter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogEventfilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LogEventfilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("LogEventfilter")

	return resourceLogEventfilterRead(d, m)
}

func resourceLogEventfilterDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLogEventfilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LogEventfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogEventfilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogEventfilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LogEventfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogEventfilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogEventfilter resource from API: %v", err)
	}
	return nil
}

func flattenLogEventfilterCifs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterConnector(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterEndpoint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterFortiextender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterRestApi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterRouter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterSecurityRating(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterSwitchController(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterVpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterWanOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterWebproxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogEventfilterWirelessActivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogEventfilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("cifs", flattenLogEventfilterCifs(o["cifs"], d, "cifs")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs"], "LogEventfilter-Cifs"); ok {
			if err = d.Set("cifs", vv); err != nil {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs: %v", err)
		}
	}

	if err = d.Set("connector", flattenLogEventfilterConnector(o["connector"], d, "connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["connector"], "LogEventfilter-Connector"); ok {
			if err = d.Set("connector", vv); err != nil {
				return fmt.Errorf("Error reading connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading connector: %v", err)
		}
	}

	if err = d.Set("endpoint", flattenLogEventfilterEndpoint(o["endpoint"], d, "endpoint")); err != nil {
		if vv, ok := fortiAPIPatch(o["endpoint"], "LogEventfilter-Endpoint"); ok {
			if err = d.Set("endpoint", vv); err != nil {
				return fmt.Errorf("Error reading endpoint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading endpoint: %v", err)
		}
	}

	if err = d.Set("event", flattenLogEventfilterEvent(o["event"], d, "event")); err != nil {
		if vv, ok := fortiAPIPatch(o["event"], "LogEventfilter-Event"); ok {
			if err = d.Set("event", vv); err != nil {
				return fmt.Errorf("Error reading event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("fortiextender", flattenLogEventfilterFortiextender(o["fortiextender"], d, "fortiextender")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiextender"], "LogEventfilter-Fortiextender"); ok {
			if err = d.Set("fortiextender", vv); err != nil {
				return fmt.Errorf("Error reading fortiextender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiextender: %v", err)
		}
	}

	if err = d.Set("ha", flattenLogEventfilterHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "LogEventfilter-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("rest_api", flattenLogEventfilterRestApi(o["rest-api"], d, "rest_api")); err != nil {
		if vv, ok := fortiAPIPatch(o["rest-api"], "LogEventfilter-RestApi"); ok {
			if err = d.Set("rest_api", vv); err != nil {
				return fmt.Errorf("Error reading rest_api: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rest_api: %v", err)
		}
	}

	if err = d.Set("router", flattenLogEventfilterRouter(o["router"], d, "router")); err != nil {
		if vv, ok := fortiAPIPatch(o["router"], "LogEventfilter-Router"); ok {
			if err = d.Set("router", vv); err != nil {
				return fmt.Errorf("Error reading router: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenLogEventfilterSdwan(o["sdwan"], d, "sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdwan"], "LogEventfilter-Sdwan"); ok {
			if err = d.Set("sdwan", vv); err != nil {
				return fmt.Errorf("Error reading sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdwan: %v", err)
		}
	}

	if err = d.Set("security_rating", flattenLogEventfilterSecurityRating(o["security-rating"], d, "security_rating")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-rating"], "LogEventfilter-SecurityRating"); ok {
			if err = d.Set("security_rating", vv); err != nil {
				return fmt.Errorf("Error reading security_rating: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_rating: %v", err)
		}
	}

	if err = d.Set("switch_controller", flattenLogEventfilterSwitchController(o["switch-controller"], d, "switch_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller"], "LogEventfilter-SwitchController"); ok {
			if err = d.Set("switch_controller", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller: %v", err)
		}
	}

	if err = d.Set("system", flattenLogEventfilterSystem(o["system"], d, "system")); err != nil {
		if vv, ok := fortiAPIPatch(o["system"], "LogEventfilter-System"); ok {
			if err = d.Set("system", vv); err != nil {
				return fmt.Errorf("Error reading system: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("user", flattenLogEventfilterUser(o["user"], d, "user")); err != nil {
		if vv, ok := fortiAPIPatch(o["user"], "LogEventfilter-User"); ok {
			if err = d.Set("user", vv); err != nil {
				return fmt.Errorf("Error reading user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("vpn", flattenLogEventfilterVpn(o["vpn"], d, "vpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn"], "LogEventfilter-Vpn"); ok {
			if err = d.Set("vpn", vv); err != nil {
				return fmt.Errorf("Error reading vpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn: %v", err)
		}
	}

	if err = d.Set("wan_opt", flattenLogEventfilterWanOpt(o["wan-opt"], d, "wan_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["wan-opt"], "LogEventfilter-WanOpt"); ok {
			if err = d.Set("wan_opt", vv); err != nil {
				return fmt.Errorf("Error reading wan_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wan_opt: %v", err)
		}
	}

	if err = d.Set("webproxy", flattenLogEventfilterWebproxy(o["webproxy"], d, "webproxy")); err != nil {
		if vv, ok := fortiAPIPatch(o["webproxy"], "LogEventfilter-Webproxy"); ok {
			if err = d.Set("webproxy", vv); err != nil {
				return fmt.Errorf("Error reading webproxy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webproxy: %v", err)
		}
	}

	if err = d.Set("wireless_activity", flattenLogEventfilterWirelessActivity(o["wireless-activity"], d, "wireless_activity")); err != nil {
		if vv, ok := fortiAPIPatch(o["wireless-activity"], "LogEventfilter-WirelessActivity"); ok {
			if err = d.Set("wireless_activity", vv); err != nil {
				return fmt.Errorf("Error reading wireless_activity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wireless_activity: %v", err)
		}
	}

	return nil
}

func flattenLogEventfilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogEventfilterCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterEndpoint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterFortiextender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterRestApi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterRouter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSecurityRating(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSwitchController(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterVpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterWanOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterWebproxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogEventfilterWirelessActivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogEventfilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("cifs"); ok || d.HasChange("cifs") {
		t, err := expandLogEventfilterCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("connector"); ok || d.HasChange("connector") {
		t, err := expandLogEventfilterConnector(d, v, "connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector"] = t
		}
	}

	if v, ok := d.GetOk("endpoint"); ok || d.HasChange("endpoint") {
		t, err := expandLogEventfilterEndpoint(d, v, "endpoint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok || d.HasChange("event") {
		t, err := expandLogEventfilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("fortiextender"); ok || d.HasChange("fortiextender") {
		t, err := expandLogEventfilterFortiextender(d, v, "fortiextender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiextender"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok || d.HasChange("ha") {
		t, err := expandLogEventfilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("rest_api"); ok || d.HasChange("rest_api") {
		t, err := expandLogEventfilterRestApi(d, v, "rest_api")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api"] = t
		}
	}

	if v, ok := d.GetOk("router"); ok || d.HasChange("router") {
		t, err := expandLogEventfilterRouter(d, v, "router")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok || d.HasChange("sdwan") {
		t, err := expandLogEventfilterSdwan(d, v, "sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("security_rating"); ok || d.HasChange("security_rating") {
		t, err := expandLogEventfilterSecurityRating(d, v, "security_rating")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-rating"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller"); ok || d.HasChange("switch_controller") {
		t, err := expandLogEventfilterSwitchController(d, v, "switch_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok || d.HasChange("system") {
		t, err := expandLogEventfilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok || d.HasChange("user") {
		t, err := expandLogEventfilterUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("vpn"); ok || d.HasChange("vpn") {
		t, err := expandLogEventfilterVpn(d, v, "vpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn"] = t
		}
	}

	if v, ok := d.GetOk("wan_opt"); ok || d.HasChange("wan_opt") {
		t, err := expandLogEventfilterWanOpt(d, v, "wan_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-opt"] = t
		}
	}

	if v, ok := d.GetOk("webproxy"); ok || d.HasChange("webproxy") {
		t, err := expandLogEventfilterWebproxy(d, v, "webproxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy"] = t
		}
	}

	if v, ok := d.GetOk("wireless_activity"); ok || d.HasChange("wireless_activity") {
		t, err := expandLogEventfilterWirelessActivity(d, v, "wireless_activity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-activity"] = t
		}
	}

	return &obj, nil
}
