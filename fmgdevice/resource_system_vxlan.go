// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure VXLAN devices.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVxlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVxlanCreate,
		Read:   resourceSystemVxlanRead,
		Update: resourceSystemVxlanUpdate,
		Delete: resourceSystemVxlanDelete,

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
			"dstport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"evpn_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_from_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"remote_ip6": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemVxlanCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVxlan(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVxlan resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVxlan(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemVxlan resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVxlanRead(d, m)
}

func resourceSystemVxlanUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVxlan(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVxlan resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVxlan(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemVxlan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemVxlanRead(d, m)
}

func resourceSystemVxlanDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemVxlan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVxlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVxlanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVxlan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemVxlan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVxlan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVxlan resource from API: %v", err)
	}
	return nil
}

func flattenSystemVxlanDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVxlanEvpnId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVxlanInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVxlanIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVxlanLearnFromTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVxlanMulticastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVxlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVxlanRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVxlanRemoteIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemVxlanVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVxlan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dstport", flattenSystemVxlanDstport(o["dstport"], d, "dstport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport"], "SystemVxlan-Dstport"); ok {
			if err = d.Set("dstport", vv); err != nil {
				return fmt.Errorf("Error reading dstport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("evpn_id", flattenSystemVxlanEvpnId(o["evpn-id"], d, "evpn_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["evpn-id"], "SystemVxlan-EvpnId"); ok {
			if err = d.Set("evpn_id", vv); err != nil {
				return fmt.Errorf("Error reading evpn_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading evpn_id: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVxlanInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemVxlan-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemVxlanIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "SystemVxlan-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("learn_from_traffic", flattenSystemVxlanLearnFromTraffic(o["learn-from-traffic"], d, "learn_from_traffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["learn-from-traffic"], "SystemVxlan-LearnFromTraffic"); ok {
			if err = d.Set("learn_from_traffic", vv); err != nil {
				return fmt.Errorf("Error reading learn_from_traffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading learn_from_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", flattenSystemVxlanMulticastTtl(o["multicast-ttl"], d, "multicast_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["multicast-ttl"], "SystemVxlan-MulticastTtl"); ok {
			if err = d.Set("multicast_ttl", vv); err != nil {
				return fmt.Errorf("Error reading multicast_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemVxlanName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemVxlan-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("remote_ip", flattenSystemVxlanRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-ip"], "SystemVxlan-RemoteIp"); ok {
			if err = d.Set("remote_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("remote_ip6", flattenSystemVxlanRemoteIp6(o["remote-ip6"], d, "remote_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-ip6"], "SystemVxlan-RemoteIp6"); ok {
			if err = d.Set("remote_ip6", vv); err != nil {
				return fmt.Errorf("Error reading remote_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_ip6: %v", err)
		}
	}

	if err = d.Set("vni", flattenSystemVxlanVni(o["vni"], d, "vni")); err != nil {
		if vv, ok := fortiAPIPatch(o["vni"], "SystemVxlan-Vni"); ok {
			if err = d.Set("vni", vv); err != nil {
				return fmt.Errorf("Error reading vni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	return nil
}

func flattenSystemVxlanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVxlanDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanEvpnId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVxlanInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVxlanIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanLearnFromTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanMulticastTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanRemoteIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVxlanRemoteIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemVxlanVni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVxlan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstport"); ok || d.HasChange("dstport") {
		t, err := expandSystemVxlanDstport(d, v, "dstport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	if v, ok := d.GetOk("evpn_id"); ok || d.HasChange("evpn_id") {
		t, err := expandSystemVxlanEvpnId(d, v, "evpn_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["evpn-id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemVxlanInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandSystemVxlanIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("learn_from_traffic"); ok || d.HasChange("learn_from_traffic") {
		t, err := expandSystemVxlanLearnFromTraffic(d, v, "learn_from_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-from-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_ttl"); ok || d.HasChange("multicast_ttl") {
		t, err := expandSystemVxlanMulticastTtl(d, v, "multicast_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-ttl"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemVxlanName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok || d.HasChange("remote_ip") {
		t, err := expandSystemVxlanRemoteIp(d, v, "remote_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip6"); ok || d.HasChange("remote_ip6") {
		t, err := expandSystemVxlanRemoteIp6(d, v, "remote_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip6"] = t
		}
	}

	if v, ok := d.GetOk("vni"); ok || d.HasChange("vni") {
		t, err := expandSystemVxlanVni(d, v, "vni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	return &obj, nil
}
