// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure ingress ACL policies to be applied on managed FortiSwitch ports.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAclIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAclIngressCreate,
		Read:   resourceSwitchControllerAclIngressRead,
		Update: resourceSwitchControllerAclIngressUpdate,
		Delete: resourceSwitchControllerAclIngressDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"count": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"classifier": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceSwitchControllerAclIngressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerAclIngress(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAclIngress resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerAclIngress(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAclIngress resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerAclIngressRead(d, m)
}

func resourceSwitchControllerAclIngressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerAclIngress(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngress resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAclIngress(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSwitchControllerAclIngressRead(d, m)
}

func resourceSwitchControllerAclIngressDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerAclIngress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAclIngress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAclIngressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerAclIngress(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAclIngress(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngress resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAclIngressAction(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {
		result["count"] = flattenSwitchControllerAclIngressActionCount(i["count"], d, pre_append)
	}

	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {
		result["drop"] = flattenSwitchControllerAclIngressActionDrop(i["drop"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerAclIngressActionCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressActionDrop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifier(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {
		result["dst_ip_prefix"] = flattenSwitchControllerAclIngressClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {
		result["dst_mac"] = flattenSwitchControllerAclIngressClassifierDstMac(i["dst-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {
		result["src_ip_prefix"] = flattenSwitchControllerAclIngressClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {
		result["src_mac"] = flattenSwitchControllerAclIngressClassifierSrcMac(i["src-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSwitchControllerAclIngressClassifierVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerAclIngressClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAclIngressClassifierDstMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerAclIngressClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifierVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAclIngress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("action", flattenSwitchControllerAclIngressAction(o["action"], d, "action")); err != nil {
			if vv, ok := fortiAPIPatch(o["action"], "SwitchControllerAclIngress-Action"); ok {
				if err = d.Set("action", vv); err != nil {
					return fmt.Errorf("Error reading action: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading action: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("action"); ok {
			if err = d.Set("action", flattenSwitchControllerAclIngressAction(o["action"], d, "action")); err != nil {
				if vv, ok := fortiAPIPatch(o["action"], "SwitchControllerAclIngress-Action"); ok {
					if err = d.Set("action", vv); err != nil {
						return fmt.Errorf("Error reading action: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("classifier", flattenSwitchControllerAclIngressClassifier(o["classifier"], d, "classifier")); err != nil {
			if vv, ok := fortiAPIPatch(o["classifier"], "SwitchControllerAclIngress-Classifier"); ok {
				if err = d.Set("classifier", vv); err != nil {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading classifier: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classifier"); ok {
			if err = d.Set("classifier", flattenSwitchControllerAclIngressClassifier(o["classifier"], d, "classifier")); err != nil {
				if vv, ok := fortiAPIPatch(o["classifier"], "SwitchControllerAclIngress-Classifier"); ok {
					if err = d.Set("classifier", vv); err != nil {
						return fmt.Errorf("Error reading classifier: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenSwitchControllerAclIngressDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SwitchControllerAclIngress-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerAclIngressId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SwitchControllerAclIngress-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAclIngressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAclIngressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["count"], _ = expandSwitchControllerAclIngressActionCount(d, i["count"], pre_append)
	}
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["drop"], _ = expandSwitchControllerAclIngressActionDrop(d, i["drop"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerAclIngressActionCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressActionDrop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dst-ip-prefix"], _ = expandSwitchControllerAclIngressClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append)
	}
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dst-mac"], _ = expandSwitchControllerAclIngressClassifierDstMac(d, i["dst_mac"], pre_append)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-ip-prefix"], _ = expandSwitchControllerAclIngressClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["src-mac"], _ = expandSwitchControllerAclIngressClassifierSrcMac(d, i["src_mac"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSwitchControllerAclIngressClassifierVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSwitchControllerAclIngressClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerAclIngressClassifierDstMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerAclIngressClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAclIngress(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSwitchControllerAclIngressAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("classifier"); ok || d.HasChange("classifier") {
		t, err := expandSwitchControllerAclIngressClassifier(d, v, "classifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classifier"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSwitchControllerAclIngressDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSwitchControllerAclIngressId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
