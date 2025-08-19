// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Select the real servers that this Access Proxy will distribute traffic to.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebProxyApiGateway6Realservers() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyApiGateway6RealserversCreate,
		Read:   resourceZtnaWebProxyApiGateway6RealserversRead,
		Update: resourceZtnaWebProxyApiGateway6RealserversUpdate,
		Delete: resourceZtnaWebProxyApiGateway6RealserversDelete,

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
			"web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"api_gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_check_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"holddown_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceZtnaWebProxyApiGateway6RealserversCreate(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGateway6Realservers(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGateway6Realservers resource while getting object: %v", err)
	}

	_, err = c.CreateZtnaWebProxyApiGateway6Realservers(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxyApiGateway6Realservers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceZtnaWebProxyApiGateway6RealserversRead(d, m)
}

func resourceZtnaWebProxyApiGateway6RealserversUpdate(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectZtnaWebProxyApiGateway6Realservers(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGateway6Realservers resource while getting object: %v", err)
	}

	_, err = c.UpdateZtnaWebProxyApiGateway6Realservers(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxyApiGateway6Realservers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceZtnaWebProxyApiGateway6RealserversRead(d, m)
}

func resourceZtnaWebProxyApiGateway6RealserversDelete(d *schema.ResourceData, m interface{}) error {
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
	web_proxy := d.Get("web_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteZtnaWebProxyApiGateway6Realservers(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxyApiGateway6Realservers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyApiGateway6RealserversRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	web_proxy := d.Get("web_proxy").(string)
	api_gateway6 := d.Get("api_gateway6").(string)
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
	if web_proxy == "" {
		web_proxy = importOptionChecking(m.(*FortiClient).Cfg, "web_proxy")
		if web_proxy == "" {
			return fmt.Errorf("Parameter web_proxy is missing")
		}
		if err = d.Set("web_proxy", web_proxy); err != nil {
			return fmt.Errorf("Error set params web_proxy: %v", err)
		}
	}
	if api_gateway6 == "" {
		api_gateway6 = importOptionChecking(m.(*FortiClient).Cfg, "api_gateway6")
		if api_gateway6 == "" {
			return fmt.Errorf("Parameter api_gateway6 is missing")
		}
		if err = d.Set("api_gateway6", api_gateway6); err != nil {
			return fmt.Errorf("Error set params api_gateway6: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["web_proxy"] = web_proxy
	paradict["api_gateway6"] = api_gateway6

	o, err := c.ReadZtnaWebProxyApiGateway6Realservers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGateway6Realservers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxyApiGateway6Realservers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxyApiGateway6Realservers resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyApiGateway6RealserversAddrType3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversAddress3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHolddownInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHttpHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversIp3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversPort3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversStatus3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversTranslateHost3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversWeight3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaWebProxyApiGateway6Realservers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("addr_type", flattenZtnaWebProxyApiGateway6RealserversAddrType3rdl(o["addr-type"], d, "addr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-type"], "ZtnaWebProxyApiGateway6Realservers-AddrType"); ok {
			if err = d.Set("addr_type", vv); err != nil {
				return fmt.Errorf("Error reading addr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("address", flattenZtnaWebProxyApiGateway6RealserversAddress3rdl(o["address"], d, "address")); err != nil {
		if vv, ok := fortiAPIPatch(o["address"], "ZtnaWebProxyApiGateway6Realservers-Address"); ok {
			if err = d.Set("address", vv); err != nil {
				return fmt.Errorf("Error reading address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("health_check", flattenZtnaWebProxyApiGateway6RealserversHealthCheck3rdl(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "ZtnaWebProxyApiGateway6Realservers-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("health_check_proto", flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto3rdl(o["health-check-proto"], d, "health_check_proto")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check-proto"], "ZtnaWebProxyApiGateway6Realservers-HealthCheckProto"); ok {
			if err = d.Set("health_check_proto", vv); err != nil {
				return fmt.Errorf("Error reading health_check_proto: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check_proto: %v", err)
		}
	}

	if err = d.Set("holddown_interval", flattenZtnaWebProxyApiGateway6RealserversHolddownInterval3rdl(o["holddown-interval"], d, "holddown_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["holddown-interval"], "ZtnaWebProxyApiGateway6Realservers-HolddownInterval"); ok {
			if err = d.Set("holddown_interval", vv); err != nil {
				return fmt.Errorf("Error reading holddown_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading holddown_interval: %v", err)
		}
	}

	if err = d.Set("http_host", flattenZtnaWebProxyApiGateway6RealserversHttpHost3rdl(o["http-host"], d, "http_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-host"], "ZtnaWebProxyApiGateway6Realservers-HttpHost"); ok {
			if err = d.Set("http_host", vv); err != nil {
				return fmt.Errorf("Error reading http_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_host: %v", err)
		}
	}

	if err = d.Set("fosid", flattenZtnaWebProxyApiGateway6RealserversId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ZtnaWebProxyApiGateway6Realservers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenZtnaWebProxyApiGateway6RealserversIp3rdl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "ZtnaWebProxyApiGateway6Realservers-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaWebProxyApiGateway6RealserversPort3rdl(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ZtnaWebProxyApiGateway6Realservers-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaWebProxyApiGateway6RealserversStatus3rdl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ZtnaWebProxyApiGateway6Realservers-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("translate_host", flattenZtnaWebProxyApiGateway6RealserversTranslateHost3rdl(o["translate-host"], d, "translate_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["translate-host"], "ZtnaWebProxyApiGateway6Realservers-TranslateHost"); ok {
			if err = d.Set("translate_host", vv); err != nil {
				return fmt.Errorf("Error reading translate_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading translate_host: %v", err)
		}
	}

	if err = d.Set("weight", flattenZtnaWebProxyApiGateway6RealserversWeight3rdl(o["weight"], d, "weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["weight"], "ZtnaWebProxyApiGateway6Realservers-Weight"); ok {
			if err = d.Set("weight", vv); err != nil {
				return fmt.Errorf("Error reading weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebProxyApiGateway6RealserversFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaWebProxyApiGateway6RealserversAddrType3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddress3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheckProto3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHolddownInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHttpHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversIp3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversPort3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversStatus3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversTranslateHost3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversWeight3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebProxyApiGateway6Realservers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_type"); ok || d.HasChange("addr_type") {
		t, err := expandZtnaWebProxyApiGateway6RealserversAddrType3rdl(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok || d.HasChange("address") {
		t, err := expandZtnaWebProxyApiGateway6RealserversAddress3rdl(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandZtnaWebProxyApiGateway6RealserversHealthCheck3rdl(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_proto"); ok || d.HasChange("health_check_proto") {
		t, err := expandZtnaWebProxyApiGateway6RealserversHealthCheckProto3rdl(d, v, "health_check_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-proto"] = t
		}
	}

	if v, ok := d.GetOk("holddown_interval"); ok || d.HasChange("holddown_interval") {
		t, err := expandZtnaWebProxyApiGateway6RealserversHolddownInterval3rdl(d, v, "holddown_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["holddown-interval"] = t
		}
	}

	if v, ok := d.GetOk("http_host"); ok || d.HasChange("http_host") {
		t, err := expandZtnaWebProxyApiGateway6RealserversHttpHost3rdl(d, v, "http_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-host"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandZtnaWebProxyApiGateway6RealserversId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandZtnaWebProxyApiGateway6RealserversIp3rdl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandZtnaWebProxyApiGateway6RealserversPort3rdl(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandZtnaWebProxyApiGateway6RealserversStatus3rdl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("translate_host"); ok || d.HasChange("translate_host") {
		t, err := expandZtnaWebProxyApiGateway6RealserversTranslateHost3rdl(d, v, "translate_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["translate-host"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok || d.HasChange("weight") {
		t, err := expandZtnaWebProxyApiGateway6RealserversWeight3rdl(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	return &obj, nil
}
