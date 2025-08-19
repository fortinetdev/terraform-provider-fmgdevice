// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure GENEVE devices.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGeneve() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGeneveCreate,
		Read:   resourceSystemGeneveRead,
		Update: resourceSystemGeneveUpdate,
		Delete: resourceSystemGeneveDelete,

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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemGeneveCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemGeneve(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeneve resource while getting object: %v", err)
	}

	_, err = c.CreateSystemGeneve(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeneve resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemGeneveRead(d, m)
}

func resourceSystemGeneveUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemGeneve(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeneve resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGeneve(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeneve resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemGeneveRead(d, m)
}

func resourceSystemGeneveDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemGeneve(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGeneve resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGeneveRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemGeneve(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeneve resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGeneve(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeneve resource from API: %v", err)
	}
	return nil
}

func flattenSystemGeneveDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGeneveIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveRemoteIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGeneve(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dstport", flattenSystemGeneveDstport(o["dstport"], d, "dstport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstport"], "SystemGeneve-Dstport"); ok {
			if err = d.Set("dstport", vv); err != nil {
				return fmt.Errorf("Error reading dstport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemGeneveInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemGeneve-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemGeneveIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "SystemGeneve-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemGeneveName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemGeneve-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("remote_ip", flattenSystemGeneveRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-ip"], "SystemGeneve-RemoteIp"); ok {
			if err = d.Set("remote_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("remote_ip6", flattenSystemGeneveRemoteIp6(o["remote-ip6"], d, "remote_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-ip6"], "SystemGeneve-RemoteIp6"); ok {
			if err = d.Set("remote_ip6", vv); err != nil {
				return fmt.Errorf("Error reading remote_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_ip6: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemGeneveType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemGeneve-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vni", flattenSystemGeneveVni(o["vni"], d, "vni")); err != nil {
		if vv, ok := fortiAPIPatch(o["vni"], "SystemGeneve-Vni"); ok {
			if err = d.Set("vni", vv); err != nil {
				return fmt.Errorf("Error reading vni: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	return nil
}

func flattenSystemGeneveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGeneveDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGeneveIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveRemoteIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveRemoteIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveVni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGeneve(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dstport"); ok || d.HasChange("dstport") {
		t, err := expandSystemGeneveDstport(d, v, "dstport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemGeneveInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandSystemGeneveIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemGeneveName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok || d.HasChange("remote_ip") {
		t, err := expandSystemGeneveRemoteIp(d, v, "remote_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip6"); ok || d.HasChange("remote_ip6") {
		t, err := expandSystemGeneveRemoteIp6(d, v, "remote_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip6"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemGeneveType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vni"); ok || d.HasChange("vni") {
		t, err := expandSystemGeneveVni(d, v, "vni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	return &obj, nil
}
