// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the PPPoE interfaces.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPppoeInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPppoeInterfaceCreate,
		Read:   resourceSystemPppoeInterfaceRead,
		Update: resourceSystemPppoeInterfaceUpdate,
		Delete: resourceSystemPppoeInterfaceDelete,

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
			"ac_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dial_on_demand": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ipunnumbered": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"padt_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"pppoe_unnumbered_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemPppoeInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSystemPppoeInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemPppoeInterface resource while getting object: %v", err)
	}

	_, err = c.CreateSystemPppoeInterface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemPppoeInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemPppoeInterfaceRead(d, m)
}

func resourceSystemPppoeInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSystemPppoeInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPppoeInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPppoeInterface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemPppoeInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemPppoeInterfaceRead(d, m)
}

func resourceSystemPppoeInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSystemPppoeInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPppoeInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPppoeInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPppoeInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemPppoeInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPppoeInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPppoeInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemPppoeInterfaceAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemPppoeInterfaceDialOnDemand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceIpunnumbered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfacePadtRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfacePppoeUnnumberedNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPppoeInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ac_name", flattenSystemPppoeInterfaceAcName(o["ac-name"], d, "ac_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-name"], "SystemPppoeInterface-AcName"); ok {
			if err = d.Set("ac_name", vv); err != nil {
				return fmt.Errorf("Error reading ac_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_name: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenSystemPppoeInterfaceAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "SystemPppoeInterface-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemPppoeInterfaceDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemPppoeInterface-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("dial_on_demand", flattenSystemPppoeInterfaceDialOnDemand(o["dial-on-demand"], d, "dial_on_demand")); err != nil {
		if vv, ok := fortiAPIPatch(o["dial-on-demand"], "SystemPppoeInterface-DialOnDemand"); ok {
			if err = d.Set("dial_on_demand", vv); err != nil {
				return fmt.Errorf("Error reading dial_on_demand: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dial_on_demand: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", flattenSystemPppoeInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["disc-retry-timeout"], "SystemPppoeInterface-DiscRetryTimeout"); ok {
			if err = d.Set("disc_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenSystemPppoeInterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["idle-timeout"], "SystemPppoeInterface-IdleTimeout"); ok {
			if err = d.Set("idle_timeout", vv); err != nil {
				return fmt.Errorf("Error reading idle_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("ipunnumbered", flattenSystemPppoeInterfaceIpunnumbered(o["ipunnumbered"], d, "ipunnumbered")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipunnumbered"], "SystemPppoeInterface-Ipunnumbered"); ok {
			if err = d.Set("ipunnumbered", vv); err != nil {
				return fmt.Errorf("Error reading ipunnumbered: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipunnumbered: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystemPppoeInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6"], "SystemPppoeInterface-Ipv6"); ok {
			if err = d.Set("ipv6", vv); err != nil {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", flattenSystemPppoeInterfaceLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lcp-echo-interval"], "SystemPppoeInterface-LcpEchoInterval"); ok {
			if err = d.Set("lcp_echo_interval", vv); err != nil {
				return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", flattenSystemPppoeInterfaceLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if vv, ok := fortiAPIPatch(o["lcp-max-echo-fails"], "SystemPppoeInterface-LcpMaxEchoFails"); ok {
			if err = d.Set("lcp_max_echo_fails", vv); err != nil {
				return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemPppoeInterfaceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemPppoeInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("padt_retry_timeout", flattenSystemPppoeInterfacePadtRetryTimeout(o["padt-retry-timeout"], d, "padt_retry_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["padt-retry-timeout"], "SystemPppoeInterface-PadtRetryTimeout"); ok {
			if err = d.Set("padt_retry_timeout", vv); err != nil {
				return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
		}
	}

	if err = d.Set("pppoe_unnumbered_negotiate", flattenSystemPppoeInterfacePppoeUnnumberedNegotiate(o["pppoe-unnumbered-negotiate"], d, "pppoe_unnumbered_negotiate")); err != nil {
		if vv, ok := fortiAPIPatch(o["pppoe-unnumbered-negotiate"], "SystemPppoeInterface-PppoeUnnumberedNegotiate"); ok {
			if err = d.Set("pppoe_unnumbered_negotiate", vv); err != nil {
				return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if err = d.Set("service_name", flattenSystemPppoeInterfaceServiceName(o["service-name"], d, "service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-name"], "SystemPppoeInterface-ServiceName"); ok {
			if err = d.Set("service_name", vv); err != nil {
				return fmt.Errorf("Error reading service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemPppoeInterfaceUsername(o["username"], d, "username")); err != nil {
		if vv, ok := fortiAPIPatch(o["username"], "SystemPppoeInterface-Username"); ok {
			if err = d.Set("username", vv); err != nil {
				return fmt.Errorf("Error reading username: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	return nil
}

func flattenSystemPppoeInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPppoeInterfaceAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPppoeInterfaceDialOnDemand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceDiscRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceIpunnumbered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceLcpEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceLcpMaxEchoFails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfacePadtRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfacePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemPppoeInterfacePppoeUnnumberedNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPppoeInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ac_name"); ok || d.HasChange("ac_name") {
		t, err := expandSystemPppoeInterfaceAcName(d, v, "ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-name"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandSystemPppoeInterfaceAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemPppoeInterfaceDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("dial_on_demand"); ok || d.HasChange("dial_on_demand") {
		t, err := expandSystemPppoeInterfaceDialOnDemand(d, v, "dial_on_demand")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-on-demand"] = t
		}
	}

	if v, ok := d.GetOk("disc_retry_timeout"); ok || d.HasChange("disc_retry_timeout") {
		t, err := expandSystemPppoeInterfaceDiscRetryTimeout(d, v, "disc_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disc-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok || d.HasChange("idle_timeout") {
		t, err := expandSystemPppoeInterfaceIdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ipunnumbered"); ok || d.HasChange("ipunnumbered") {
		t, err := expandSystemPppoeInterfaceIpunnumbered(d, v, "ipunnumbered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipunnumbered"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok || d.HasChange("ipv6") {
		t, err := expandSystemPppoeInterfaceIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("lcp_echo_interval"); ok || d.HasChange("lcp_echo_interval") {
		t, err := expandSystemPppoeInterfaceLcpEchoInterval(d, v, "lcp_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("lcp_max_echo_fails"); ok || d.HasChange("lcp_max_echo_fails") {
		t, err := expandSystemPppoeInterfaceLcpMaxEchoFails(d, v, "lcp_max_echo_fails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-max-echo-fails"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemPppoeInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("padt_retry_timeout"); ok || d.HasChange("padt_retry_timeout") {
		t, err := expandSystemPppoeInterfacePadtRetryTimeout(d, v, "padt_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padt-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemPppoeInterfacePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("pppoe_unnumbered_negotiate"); ok || d.HasChange("pppoe_unnumbered_negotiate") {
		t, err := expandSystemPppoeInterfacePppoeUnnumberedNegotiate(d, v, "pppoe_unnumbered_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pppoe-unnumbered-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("service_name"); ok || d.HasChange("service_name") {
		t, err := expandSystemPppoeInterfaceServiceName(d, v, "service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-name"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok || d.HasChange("username") {
		t, err := expandSystemPppoeInterfaceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	return &obj, nil
}
