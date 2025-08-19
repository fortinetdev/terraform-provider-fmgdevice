// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: FortiGate connector profile configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerFortigateProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerFortigateProfileCreate,
		Read:   resourceExtensionControllerFortigateProfileRead,
		Update: resourceExtensionControllerFortigateProfileUpdate,
		Delete: resourceExtensionControllerFortigateProfileDelete,

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
			"extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backhaul_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"backhaul_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipsec_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceExtensionControllerFortigateProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectExtensionControllerFortigateProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerFortigateProfile resource while getting object: %v", err)
	}

	_, err = c.CreateExtensionControllerFortigateProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerFortigateProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerFortigateProfileRead(d, m)
}

func resourceExtensionControllerFortigateProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectExtensionControllerFortigateProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigateProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateExtensionControllerFortigateProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigateProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerFortigateProfileRead(d, m)
}

func resourceExtensionControllerFortigateProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerFortigateProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerFortigateProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerFortigateProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerFortigateProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigateProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerFortigateProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigateProfile resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerFortigateProfileExtension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileLanExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {
		result["backhaul_interface"] = flattenExtensionControllerFortigateProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {
		result["backhaul_ip"] = flattenExtensionControllerFortigateProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {
		result["ipsec_tunnel"] = flattenExtensionControllerFortigateProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerFortigateProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerFortigateProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtensionControllerFortigateProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("extension", flattenExtensionControllerFortigateProfileExtension(o["extension"], d, "extension")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension"], "ExtensionControllerFortigateProfile-Extension"); ok {
			if err = d.Set("extension", vv); err != nil {
				return fmt.Errorf("Error reading extension: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerFortigateProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ExtensionControllerFortigateProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan_extension", flattenExtensionControllerFortigateProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["lan-extension"], "ExtensionControllerFortigateProfile-LanExtension"); ok {
				if err = d.Set("lan_extension", vv); err != nil {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenExtensionControllerFortigateProfileLanExtension(o["lan-extension"], d, "lan_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["lan-extension"], "ExtensionControllerFortigateProfile-LanExtension"); ok {
					if err = d.Set("lan_extension", vv); err != nil {
						return fmt.Errorf("Error reading lan_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenExtensionControllerFortigateProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerFortigateProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerFortigateProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerFortigateProfileExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileLanExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-interface"], _ = expandExtensionControllerFortigateProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["backhaul-ip"], _ = expandExtensionControllerFortigateProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append)
	}
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ipsec-tunnel"], _ = expandExtensionControllerFortigateProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerFortigateProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerFortigateProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerFortigateProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("extension"); ok || d.HasChange("extension") {
		t, err := expandExtensionControllerFortigateProfileExtension(d, v, "extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandExtensionControllerFortigateProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok || d.HasChange("lan_extension") {
		t, err := expandExtensionControllerFortigateProfileLanExtension(d, v, "lan_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerFortigateProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
