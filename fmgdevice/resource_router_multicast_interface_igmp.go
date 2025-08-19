// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IGMP configuration options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticastInterfaceIgmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastInterfaceIgmpUpdate,
		Read:   resourceRouterMulticastInterfaceIgmpRead,
		Update: resourceRouterMulticastInterfaceIgmpUpdate,
		Delete: resourceRouterMulticastInterfaceIgmpDelete,

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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"immediate_leave_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"last_member_query_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_member_query_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"query_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query_max_response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"router_alert_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterMulticastInterfaceIgmpUpdate(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectRouterMulticastInterfaceIgmp(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastInterfaceIgmp resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticastInterfaceIgmp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastInterfaceIgmp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterMulticastInterfaceIgmp")

	return resourceRouterMulticastInterfaceIgmpRead(d, m)
}

func resourceRouterMulticastInterfaceIgmpDelete(d *schema.ResourceData, m interface{}) error {
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
	var_interface := d.Get("interface").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteRouterMulticastInterfaceIgmp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastInterfaceIgmp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastInterfaceIgmpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	var_interface := d.Get("interface").(string)
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
	if var_interface == "" {
		var_interface = importOptionChecking(m.(*FortiClient).Cfg, "interface")
		if var_interface == "" {
			return fmt.Errorf("Parameter interface is missing")
		}
		if err = d.Set("interface", var_interface); err != nil {
			return fmt.Errorf("Error set params interface: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["interface"] = var_interface

	o, err := c.ReadRouterMulticastInterfaceIgmp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastInterfaceIgmp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastInterfaceIgmp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastInterfaceIgmp resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastInterfaceIgmpAccessGroup3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryCount3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryInterval3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpQueryTimeout3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpRouterAlertCheck3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpVersion3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticastInterfaceIgmp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("access_group", flattenRouterMulticastInterfaceIgmpAccessGroup3rdl(o["access-group"], d, "access_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["access-group"], "RouterMulticastInterfaceIgmp-AccessGroup"); ok {
			if err = d.Set("access_group", vv); err != nil {
				return fmt.Errorf("Error reading access_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading access_group: %v", err)
		}
	}

	if err = d.Set("immediate_leave_group", flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup3rdl(o["immediate-leave-group"], d, "immediate_leave_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["immediate-leave-group"], "RouterMulticastInterfaceIgmp-ImmediateLeaveGroup"); ok {
			if err = d.Set("immediate_leave_group", vv); err != nil {
				return fmt.Errorf("Error reading immediate_leave_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading immediate_leave_group: %v", err)
		}
	}

	if err = d.Set("last_member_query_count", flattenRouterMulticastInterfaceIgmpLastMemberQueryCount3rdl(o["last-member-query-count"], d, "last_member_query_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-member-query-count"], "RouterMulticastInterfaceIgmp-LastMemberQueryCount"); ok {
			if err = d.Set("last_member_query_count", vv); err != nil {
				return fmt.Errorf("Error reading last_member_query_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_member_query_count: %v", err)
		}
	}

	if err = d.Set("last_member_query_interval", flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval3rdl(o["last-member-query-interval"], d, "last_member_query_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["last-member-query-interval"], "RouterMulticastInterfaceIgmp-LastMemberQueryInterval"); ok {
			if err = d.Set("last_member_query_interval", vv); err != nil {
				return fmt.Errorf("Error reading last_member_query_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_member_query_interval: %v", err)
		}
	}

	if err = d.Set("query_interval", flattenRouterMulticastInterfaceIgmpQueryInterval3rdl(o["query-interval"], d, "query_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-interval"], "RouterMulticastInterfaceIgmp-QueryInterval"); ok {
			if err = d.Set("query_interval", vv); err != nil {
				return fmt.Errorf("Error reading query_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_interval: %v", err)
		}
	}

	if err = d.Set("query_max_response_time", flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime3rdl(o["query-max-response-time"], d, "query_max_response_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-max-response-time"], "RouterMulticastInterfaceIgmp-QueryMaxResponseTime"); ok {
			if err = d.Set("query_max_response_time", vv); err != nil {
				return fmt.Errorf("Error reading query_max_response_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_max_response_time: %v", err)
		}
	}

	if err = d.Set("query_timeout", flattenRouterMulticastInterfaceIgmpQueryTimeout3rdl(o["query-timeout"], d, "query_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["query-timeout"], "RouterMulticastInterfaceIgmp-QueryTimeout"); ok {
			if err = d.Set("query_timeout", vv); err != nil {
				return fmt.Errorf("Error reading query_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading query_timeout: %v", err)
		}
	}

	if err = d.Set("router_alert_check", flattenRouterMulticastInterfaceIgmpRouterAlertCheck3rdl(o["router-alert-check"], d, "router_alert_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-alert-check"], "RouterMulticastInterfaceIgmp-RouterAlertCheck"); ok {
			if err = d.Set("router_alert_check", vv); err != nil {
				return fmt.Errorf("Error reading router_alert_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_alert_check: %v", err)
		}
	}

	if err = d.Set("version", flattenRouterMulticastInterfaceIgmpVersion3rdl(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "RouterMulticastInterfaceIgmp-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenRouterMulticastInterfaceIgmpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticastInterfaceIgmpAccessGroup3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceIgmpImmediateLeaveGroup3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryCount3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryInterval3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryMaxResponseTime3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryTimeout3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpRouterAlertCheck3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpVersion3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastInterfaceIgmp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_group"); ok || d.HasChange("access_group") {
		t, err := expandRouterMulticastInterfaceIgmpAccessGroup3rdl(d, v, "access_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-group"] = t
		}
	}

	if v, ok := d.GetOk("immediate_leave_group"); ok || d.HasChange("immediate_leave_group") {
		t, err := expandRouterMulticastInterfaceIgmpImmediateLeaveGroup3rdl(d, v, "immediate_leave_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["immediate-leave-group"] = t
		}
	}

	if v, ok := d.GetOk("last_member_query_count"); ok || d.HasChange("last_member_query_count") {
		t, err := expandRouterMulticastInterfaceIgmpLastMemberQueryCount3rdl(d, v, "last_member_query_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-member-query-count"] = t
		}
	}

	if v, ok := d.GetOk("last_member_query_interval"); ok || d.HasChange("last_member_query_interval") {
		t, err := expandRouterMulticastInterfaceIgmpLastMemberQueryInterval3rdl(d, v, "last_member_query_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-member-query-interval"] = t
		}
	}

	if v, ok := d.GetOk("query_interval"); ok || d.HasChange("query_interval") {
		t, err := expandRouterMulticastInterfaceIgmpQueryInterval3rdl(d, v, "query_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-interval"] = t
		}
	}

	if v, ok := d.GetOk("query_max_response_time"); ok || d.HasChange("query_max_response_time") {
		t, err := expandRouterMulticastInterfaceIgmpQueryMaxResponseTime3rdl(d, v, "query_max_response_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-max-response-time"] = t
		}
	}

	if v, ok := d.GetOk("query_timeout"); ok || d.HasChange("query_timeout") {
		t, err := expandRouterMulticastInterfaceIgmpQueryTimeout3rdl(d, v, "query_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-timeout"] = t
		}
	}

	if v, ok := d.GetOk("router_alert_check"); ok || d.HasChange("router_alert_check") {
		t, err := expandRouterMulticastInterfaceIgmpRouterAlertCheck3rdl(d, v, "router_alert_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-alert-check"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok || d.HasChange("version") {
		t, err := expandRouterMulticastInterfaceIgmpVersion3rdl(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	return &obj, nil
}
