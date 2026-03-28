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

func resourceZtnaTrafficForwardProxyUrlRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxyUrlRouteCreate,
		Read:   resourceZtnaTrafficForwardProxyUrlRouteRead,
		Update: resourceZtnaTrafficForwardProxyUrlRouteUpdate,
		Delete: resourceZtnaTrafficForwardProxyUrlRouteDelete,

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
			"traffic_forward_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"service_connector": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"url_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceZtnaTrafficForwardProxyUrlRouteCreate(d *schema.ResourceData, m interface{}) error {
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
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	obj, err := getObjectZtnaTrafficForwardProxyUrlRoute(d)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxyUrlRoute resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadZtnaTrafficForwardProxyUrlRoute(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateZtnaTrafficForwardProxyUrlRoute(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ZtnaTrafficForwardProxyUrlRoute resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateZtnaTrafficForwardProxyUrlRoute(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ZtnaTrafficForwardProxyUrlRoute resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaTrafficForwardProxyUrlRouteRead(d, m)
}

func resourceZtnaTrafficForwardProxyUrlRouteUpdate(d *schema.ResourceData, m interface{}) error {
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
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	obj, err := getObjectZtnaTrafficForwardProxyUrlRoute(d)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyUrlRoute resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateZtnaTrafficForwardProxyUrlRoute(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxyUrlRoute resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceZtnaTrafficForwardProxyUrlRouteRead(d, m)
}

func resourceZtnaTrafficForwardProxyUrlRouteDelete(d *schema.ResourceData, m interface{}) error {
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
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	wsParams["adom"] = adomv

	err = c.DeleteZtnaTrafficForwardProxyUrlRoute(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxyUrlRoute resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxyUrlRouteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	traffic_forward_proxy := d.Get("traffic_forward_proxy").(string)
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
	if traffic_forward_proxy == "" {
		traffic_forward_proxy = importOptionChecking(m.(*FortiClient).Cfg, "traffic_forward_proxy")
		if traffic_forward_proxy == "" {
			return fmt.Errorf("Parameter traffic_forward_proxy is missing")
		}
		if err = d.Set("traffic_forward_proxy", traffic_forward_proxy); err != nil {
			return fmt.Errorf("Error set params traffic_forward_proxy: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["traffic_forward_proxy"] = traffic_forward_proxy

	o, err := c.ReadZtnaTrafficForwardProxyUrlRoute(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyUrlRoute resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxyUrlRoute(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxyUrlRoute resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxyUrlRouteName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectZtnaTrafficForwardProxyUrlRoute(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenZtnaTrafficForwardProxyUrlRouteName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ZtnaTrafficForwardProxyUrlRoute-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("service_connector", flattenZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(o["service-connector"], d, "service_connector")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-connector"], "ZtnaTrafficForwardProxyUrlRoute-ServiceConnector"); ok {
			if err = d.Set("service_connector", vv); err != nil {
				return fmt.Errorf("Error reading service_connector: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_connector: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(o["url-pattern"], d, "url_pattern")); err != nil {
		if vv, ok := fortiAPIPatch(o["url-pattern"], "ZtnaTrafficForwardProxyUrlRoute-UrlPattern"); ok {
			if err = d.Set("url_pattern", vv); err != nil {
				return fmt.Errorf("Error reading url_pattern: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxyUrlRouteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandZtnaTrafficForwardProxyUrlRouteName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaTrafficForwardProxyUrlRoute(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandZtnaTrafficForwardProxyUrlRouteName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service_connector"); ok || d.HasChange("service_connector") {
		t, err := expandZtnaTrafficForwardProxyUrlRouteServiceConnector2edl(d, v, "service_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-connector"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok || d.HasChange("url_pattern") {
		t, err := expandZtnaTrafficForwardProxyUrlRouteUrlPattern2edl(d, v, "url_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	return &obj, nil
}
