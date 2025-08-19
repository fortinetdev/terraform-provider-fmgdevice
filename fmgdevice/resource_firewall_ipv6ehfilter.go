// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 extension header filter.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIpv6EhFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpv6EhFilterUpdate,
		Read:   resourceFirewallIpv6EhFilterRead,
		Update: resourceFirewallIpv6EhFilterUpdate,
		Delete: resourceFirewallIpv6EhFilterDelete,

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
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dest_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hdopt_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"hop_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"no_next": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIpv6EhFilterUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectFirewallIpv6EhFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpv6EhFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallIpv6EhFilter(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpv6EhFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallIpv6EhFilter")

	return resourceFirewallIpv6EhFilterRead(d, m)
}

func resourceFirewallIpv6EhFilterDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteFirewallIpv6EhFilter(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIpv6EhFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpv6EhFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadFirewallIpv6EhFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpv6EhFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpv6EhFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpv6EhFilter resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpv6EhFilterAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterDestOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterHdoptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallIpv6EhFilterHopOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterNoNext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterRoutingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func refreshObjectFirewallIpv6EhFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth", flattenFirewallIpv6EhFilterAuth(o["auth"], d, "auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth"], "FirewallIpv6EhFilter-Auth"); ok {
			if err = d.Set("auth", vv); err != nil {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("dest_opt", flattenFirewallIpv6EhFilterDestOpt(o["dest-opt"], d, "dest_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["dest-opt"], "FirewallIpv6EhFilter-DestOpt"); ok {
			if err = d.Set("dest_opt", vv); err != nil {
				return fmt.Errorf("Error reading dest_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dest_opt: %v", err)
		}
	}

	if err = d.Set("fragment", flattenFirewallIpv6EhFilterFragment(o["fragment"], d, "fragment")); err != nil {
		if vv, ok := fortiAPIPatch(o["fragment"], "FirewallIpv6EhFilter-Fragment"); ok {
			if err = d.Set("fragment", vv); err != nil {
				return fmt.Errorf("Error reading fragment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fragment: %v", err)
		}
	}

	if err = d.Set("hdopt_type", flattenFirewallIpv6EhFilterHdoptType(o["hdopt-type"], d, "hdopt_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["hdopt-type"], "FirewallIpv6EhFilter-HdoptType"); ok {
			if err = d.Set("hdopt_type", vv); err != nil {
				return fmt.Errorf("Error reading hdopt_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hdopt_type: %v", err)
		}
	}

	if err = d.Set("hop_opt", flattenFirewallIpv6EhFilterHopOpt(o["hop-opt"], d, "hop_opt")); err != nil {
		if vv, ok := fortiAPIPatch(o["hop-opt"], "FirewallIpv6EhFilter-HopOpt"); ok {
			if err = d.Set("hop_opt", vv); err != nil {
				return fmt.Errorf("Error reading hop_opt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hop_opt: %v", err)
		}
	}

	if err = d.Set("no_next", flattenFirewallIpv6EhFilterNoNext(o["no-next"], d, "no_next")); err != nil {
		if vv, ok := fortiAPIPatch(o["no-next"], "FirewallIpv6EhFilter-NoNext"); ok {
			if err = d.Set("no_next", vv); err != nil {
				return fmt.Errorf("Error reading no_next: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading no_next: %v", err)
		}
	}

	if err = d.Set("routing", flattenFirewallIpv6EhFilterRouting(o["routing"], d, "routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["routing"], "FirewallIpv6EhFilter-Routing"); ok {
			if err = d.Set("routing", vv); err != nil {
				return fmt.Errorf("Error reading routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading routing: %v", err)
		}
	}

	if err = d.Set("routing_type", flattenFirewallIpv6EhFilterRoutingType(o["routing-type"], d, "routing_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["routing-type"], "FirewallIpv6EhFilter-RoutingType"); ok {
			if err = d.Set("routing_type", vv); err != nil {
				return fmt.Errorf("Error reading routing_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading routing_type: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpv6EhFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIpv6EhFilterAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterDestOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterHdoptType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallIpv6EhFilterHopOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterNoNext(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterRoutingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func getObjectFirewallIpv6EhFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandFirewallIpv6EhFilterAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("dest_opt"); ok || d.HasChange("dest_opt") {
		t, err := expandFirewallIpv6EhFilterDestOpt(d, v, "dest_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest-opt"] = t
		}
	}

	if v, ok := d.GetOk("fragment"); ok || d.HasChange("fragment") {
		t, err := expandFirewallIpv6EhFilterFragment(d, v, "fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragment"] = t
		}
	}

	if v, ok := d.GetOk("hdopt_type"); ok || d.HasChange("hdopt_type") {
		t, err := expandFirewallIpv6EhFilterHdoptType(d, v, "hdopt_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hdopt-type"] = t
		}
	}

	if v, ok := d.GetOk("hop_opt"); ok || d.HasChange("hop_opt") {
		t, err := expandFirewallIpv6EhFilterHopOpt(d, v, "hop_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hop-opt"] = t
		}
	}

	if v, ok := d.GetOk("no_next"); ok || d.HasChange("no_next") {
		t, err := expandFirewallIpv6EhFilterNoNext(d, v, "no_next")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no-next"] = t
		}
	}

	if v, ok := d.GetOk("routing"); ok || d.HasChange("routing") {
		t, err := expandFirewallIpv6EhFilterRouting(d, v, "routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routing"] = t
		}
	}

	if v, ok := d.GetOk("routing_type"); ok || d.HasChange("routing_type") {
		t, err := expandFirewallIpv6EhFilterRoutingType(d, v, "routing_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routing-type"] = t
		}
	}

	return &obj, nil
}
