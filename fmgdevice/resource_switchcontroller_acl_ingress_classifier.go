// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ACL classifiers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAclIngressClassifier() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAclIngressClassifierUpdate,
		Read:   resourceSwitchControllerAclIngressClassifierRead,
		Update: resourceSwitchControllerAclIngressClassifierUpdate,
		Delete: resourceSwitchControllerAclIngressClassifierDelete,

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
			"ingress": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dst_ip_prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_ip_prefix": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerAclIngressClassifierUpdate(d *schema.ResourceData, m interface{}) error {
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
	ingress := d.Get("ingress").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ingress"] = ingress

	obj, err := getObjectSwitchControllerAclIngressClassifier(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngressClassifier resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAclIngressClassifier(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngressClassifier resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerAclIngressClassifier")

	return resourceSwitchControllerAclIngressClassifierRead(d, m)
}

func resourceSwitchControllerAclIngressClassifierDelete(d *schema.ResourceData, m interface{}) error {
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
	ingress := d.Get("ingress").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ingress"] = ingress

	err = c.DeleteSwitchControllerAclIngressClassifier(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAclIngressClassifier resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAclIngressClassifierRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	ingress := d.Get("ingress").(string)
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
	if ingress == "" {
		ingress = importOptionChecking(m.(*FortiClient).Cfg, "ingress")
		if ingress == "" {
			return fmt.Errorf("Parameter ingress is missing")
		}
		if err = d.Set("ingress", ingress); err != nil {
			return fmt.Errorf("Error set params ingress: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["ingress"] = ingress

	o, err := c.ReadSwitchControllerAclIngressClassifier(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngressClassifier resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAclIngressClassifier(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngressClassifier resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAclIngressClassifierDstIpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAclIngressClassifierDstMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifierSrcIpPrefix2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAclIngressClassifierSrcMac2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifierVlan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAclIngressClassifier(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst_ip_prefix", flattenSwitchControllerAclIngressClassifierDstIpPrefix2edl(o["dst-ip-prefix"], d, "dst_ip_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-ip-prefix"], "SwitchControllerAclIngressClassifier-DstIpPrefix"); ok {
			if err = d.Set("dst_ip_prefix", vv); err != nil {
				return fmt.Errorf("Error reading dst_ip_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_ip_prefix: %v", err)
		}
	}

	if err = d.Set("dst_mac", flattenSwitchControllerAclIngressClassifierDstMac2edl(o["dst-mac"], d, "dst_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-mac"], "SwitchControllerAclIngressClassifier-DstMac"); ok {
			if err = d.Set("dst_mac", vv); err != nil {
				return fmt.Errorf("Error reading dst_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_mac: %v", err)
		}
	}

	if err = d.Set("src_ip_prefix", flattenSwitchControllerAclIngressClassifierSrcIpPrefix2edl(o["src-ip-prefix"], d, "src_ip_prefix")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip-prefix"], "SwitchControllerAclIngressClassifier-SrcIpPrefix"); ok {
			if err = d.Set("src_ip_prefix", vv); err != nil {
				return fmt.Errorf("Error reading src_ip_prefix: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip_prefix: %v", err)
		}
	}

	if err = d.Set("src_mac", flattenSwitchControllerAclIngressClassifierSrcMac2edl(o["src-mac"], d, "src_mac")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-mac"], "SwitchControllerAclIngressClassifier-SrcMac"); ok {
			if err = d.Set("src_mac", vv); err != nil {
				return fmt.Errorf("Error reading src_mac: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_mac: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerAclIngressClassifierVlan2edl(o["vlan"], d, "vlan")); err != nil {
		if vv, ok := fortiAPIPatch(o["vlan"], "SwitchControllerAclIngressClassifier-Vlan"); ok {
			if err = d.Set("vlan", vv); err != nil {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAclIngressClassifierFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAclIngressClassifierDstIpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerAclIngressClassifierDstMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierSrcIpPrefix2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerAclIngressClassifierSrcMac2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierVlan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAclIngressClassifier(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_ip_prefix"); ok || d.HasChange("dst_ip_prefix") {
		t, err := expandSwitchControllerAclIngressClassifierDstIpPrefix2edl(d, v, "dst_ip_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-ip-prefix"] = t
		}
	}

	if v, ok := d.GetOk("dst_mac"); ok || d.HasChange("dst_mac") {
		t, err := expandSwitchControllerAclIngressClassifierDstMac2edl(d, v, "dst_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-mac"] = t
		}
	}

	if v, ok := d.GetOk("src_ip_prefix"); ok || d.HasChange("src_ip_prefix") {
		t, err := expandSwitchControllerAclIngressClassifierSrcIpPrefix2edl(d, v, "src_ip_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip-prefix"] = t
		}
	}

	if v, ok := d.GetOk("src_mac"); ok || d.HasChange("src_mac") {
		t, err := expandSwitchControllerAclIngressClassifierSrcMac2edl(d, v, "src_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-mac"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerAclIngressClassifierVlan2edl(d, v, "vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	return &obj, nil
}
