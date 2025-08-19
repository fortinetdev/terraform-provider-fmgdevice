// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Extender controller configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtender() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderCreate,
		Read:   resourceExtensionControllerExtenderRead,
		Update: resourceExtensionControllerExtenderUpdate,
		Delete: resourceExtensionControllerExtenderDelete,

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
			"allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"authorized": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"extension_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"override_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"wan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modem1_extension": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem1_pdn1_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem1_pdn2_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem1_pdn3_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem1_pdn4_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem2_extension": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem2_pdn1_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem2_pdn2_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem2_pdn3_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"modem2_pdn4_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceExtensionControllerExtenderCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectExtensionControllerExtender(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtender resource while getting object: %v", err)
	}

	_, err = c.CreateExtensionControllerExtender(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtender resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderRead(d, m)
}

func resourceExtensionControllerExtenderUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectExtensionControllerExtender(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtender resource while getting object: %v", err)
	}

	_, err = c.UpdateExtensionControllerExtender(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtender resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceExtensionControllerExtenderRead(d, m)
}

func resourceExtensionControllerExtenderDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteExtensionControllerExtender(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtender resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtender(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtender resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtender(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtender resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderAuthorized(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderBandwidthLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderExtName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderExtensionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderOverrideEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderOverrideLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtensionControllerExtenderWanExtension(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := i["modem1-extension"]; ok {
		result["modem1_extension"] = flattenExtensionControllerExtenderWanExtensionModem1Extension(i["modem1-extension"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1_pdn1_interface"
	if _, ok := i["modem1-pdn1-interface"]; ok {
		result["modem1_pdn1_interface"] = flattenExtensionControllerExtenderWanExtensionModem1Pdn1Interface(i["modem1-pdn1-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1_pdn2_interface"
	if _, ok := i["modem1-pdn2-interface"]; ok {
		result["modem1_pdn2_interface"] = flattenExtensionControllerExtenderWanExtensionModem1Pdn2Interface(i["modem1-pdn2-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1_pdn3_interface"
	if _, ok := i["modem1-pdn3-interface"]; ok {
		result["modem1_pdn3_interface"] = flattenExtensionControllerExtenderWanExtensionModem1Pdn3Interface(i["modem1-pdn3-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem1_pdn4_interface"
	if _, ok := i["modem1-pdn4-interface"]; ok {
		result["modem1_pdn4_interface"] = flattenExtensionControllerExtenderWanExtensionModem1Pdn4Interface(i["modem1-pdn4-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := i["modem2-extension"]; ok {
		result["modem2_extension"] = flattenExtensionControllerExtenderWanExtensionModem2Extension(i["modem2-extension"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2_pdn1_interface"
	if _, ok := i["modem2-pdn1-interface"]; ok {
		result["modem2_pdn1_interface"] = flattenExtensionControllerExtenderWanExtensionModem2Pdn1Interface(i["modem2-pdn1-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2_pdn2_interface"
	if _, ok := i["modem2-pdn2-interface"]; ok {
		result["modem2_pdn2_interface"] = flattenExtensionControllerExtenderWanExtensionModem2Pdn2Interface(i["modem2-pdn2-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2_pdn3_interface"
	if _, ok := i["modem2-pdn3-interface"]; ok {
		result["modem2_pdn3_interface"] = flattenExtensionControllerExtenderWanExtensionModem2Pdn3Interface(i["modem2-pdn3-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "modem2_pdn4_interface"
	if _, ok := i["modem2-pdn4-interface"]; ok {
		result["modem2_pdn4_interface"] = flattenExtensionControllerExtenderWanExtensionModem2Pdn4Interface(i["modem2-pdn4-interface"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderWanExtensionModem1Extension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem1Pdn1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem1Pdn2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem1Pdn3Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem1Pdn4Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem2Extension(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem2Pdn1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem2Pdn2Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem2Pdn3Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenExtensionControllerExtenderWanExtensionModem2Pdn4Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectExtensionControllerExtender(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("allowaccess", flattenExtensionControllerExtenderAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["allowaccess"], "ExtensionControllerExtender-Allowaccess"); ok {
			if err = d.Set("allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("authorized", flattenExtensionControllerExtenderAuthorized(o["authorized"], d, "authorized")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorized"], "ExtensionControllerExtender-Authorized"); ok {
			if err = d.Set("authorized", vv); err != nil {
				return fmt.Errorf("Error reading authorized: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorized: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtensionControllerExtenderBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-limit"], "ExtensionControllerExtender-BandwidthLimit"); ok {
			if err = d.Set("bandwidth_limit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if err = d.Set("description", flattenExtensionControllerExtenderDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "ExtensionControllerExtender-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("device_id", flattenExtensionControllerExtenderDeviceId(o["device-id"], d, "device_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-id"], "ExtensionControllerExtender-DeviceId"); ok {
			if err = d.Set("device_id", vv); err != nil {
				return fmt.Errorf("Error reading device_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtensionControllerExtenderEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["enforce-bandwidth"], "ExtensionControllerExtender-EnforceBandwidth"); ok {
			if err = d.Set("enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("ext_name", flattenExtensionControllerExtenderExtName(o["ext-name"], d, "ext_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ext-name"], "ExtensionControllerExtender-ExtName"); ok {
			if err = d.Set("ext_name", vv); err != nil {
				return fmt.Errorf("Error reading ext_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ext_name: %v", err)
		}
	}

	if err = d.Set("extension_type", flattenExtensionControllerExtenderExtensionType(o["extension-type"], d, "extension_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["extension-type"], "ExtensionControllerExtender-ExtensionType"); ok {
			if err = d.Set("extension_type", vv); err != nil {
				return fmt.Errorf("Error reading extension_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extension_type: %v", err)
		}
	}

	if err = d.Set("firmware_provision_latest", flattenExtensionControllerExtenderFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest")); err != nil {
		if vv, ok := fortiAPIPatch(o["firmware-provision-latest"], "ExtensionControllerExtender-FirmwareProvisionLatest"); ok {
			if err = d.Set("firmware_provision_latest", vv); err != nil {
				return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerExtenderId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ExtensionControllerExtender-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenExtensionControllerExtenderLoginPasswordChange(o["login-password-change"], d, "login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-password-change"], "ExtensionControllerExtender-LoginPasswordChange"); ok {
			if err = d.Set("login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("name", flattenExtensionControllerExtenderName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ExtensionControllerExtender-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("override_allowaccess", flattenExtensionControllerExtenderOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-allowaccess"], "ExtensionControllerExtender-OverrideAllowaccess"); ok {
			if err = d.Set("override_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading override_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("override_enforce_bandwidth", flattenExtensionControllerExtenderOverrideEnforceBandwidth(o["override-enforce-bandwidth"], d, "override_enforce_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-enforce-bandwidth"], "ExtensionControllerExtender-OverrideEnforceBandwidth"); ok {
			if err = d.Set("override_enforce_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading override_enforce_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("override_login_password_change", flattenExtensionControllerExtenderOverrideLoginPasswordChange(o["override-login-password-change"], d, "override_login_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["override-login-password-change"], "ExtensionControllerExtender-OverrideLoginPasswordChange"); ok {
			if err = d.Set("override_login_password_change", vv); err != nil {
				return fmt.Errorf("Error reading override_login_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading override_login_password_change: %v", err)
		}
	}

	if err = d.Set("profile", flattenExtensionControllerExtenderProfile(o["profile"], d, "profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile"], "ExtensionControllerExtender-Profile"); ok {
			if err = d.Set("profile", vv); err != nil {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtensionControllerExtenderVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "ExtensionControllerExtender-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wan_extension", flattenExtensionControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension")); err != nil {
			if vv, ok := fortiAPIPatch(o["wan-extension"], "ExtensionControllerExtender-WanExtension"); ok {
				if err = d.Set("wan_extension", vv); err != nil {
					return fmt.Errorf("Error reading wan_extension: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading wan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wan_extension"); ok {
			if err = d.Set("wan_extension", flattenExtensionControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension")); err != nil {
				if vv, ok := fortiAPIPatch(o["wan-extension"], "ExtensionControllerExtender-WanExtension"); ok {
					if err = d.Set("wan_extension", vv); err != nil {
						return fmt.Errorf("Error reading wan_extension: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading wan_extension: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtensionControllerExtenderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtensionControllerExtenderAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderAuthorized(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderBandwidthLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderExtName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderExtensionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderLoginPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderOverrideEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderOverrideLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderWanExtension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem1-extension"], _ = expandExtensionControllerExtenderWanExtensionModem1Extension(d, i["modem1_extension"], pre_append)
	}
	pre_append = pre + ".0." + "modem1_pdn1_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem1-pdn1-interface"], _ = expandExtensionControllerExtenderWanExtensionModem1Pdn1Interface(d, i["modem1_pdn1_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem1_pdn2_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem1-pdn2-interface"], _ = expandExtensionControllerExtenderWanExtensionModem1Pdn2Interface(d, i["modem1_pdn2_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem1_pdn3_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem1-pdn3-interface"], _ = expandExtensionControllerExtenderWanExtensionModem1Pdn3Interface(d, i["modem1_pdn3_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem1_pdn4_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem1-pdn4-interface"], _ = expandExtensionControllerExtenderWanExtensionModem1Pdn4Interface(d, i["modem1_pdn4_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem2-extension"], _ = expandExtensionControllerExtenderWanExtensionModem2Extension(d, i["modem2_extension"], pre_append)
	}
	pre_append = pre + ".0." + "modem2_pdn1_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem2-pdn1-interface"], _ = expandExtensionControllerExtenderWanExtensionModem2Pdn1Interface(d, i["modem2_pdn1_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem2_pdn2_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem2-pdn2-interface"], _ = expandExtensionControllerExtenderWanExtensionModem2Pdn2Interface(d, i["modem2_pdn2_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem2_pdn3_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem2-pdn3-interface"], _ = expandExtensionControllerExtenderWanExtensionModem2Pdn3Interface(d, i["modem2_pdn3_interface"], pre_append)
	}
	pre_append = pre + ".0." + "modem2_pdn4_interface"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["modem2-pdn4-interface"], _ = expandExtensionControllerExtenderWanExtensionModem2Pdn4Interface(d, i["modem2_pdn4_interface"], pre_append)
	}

	return result, nil
}

func expandExtensionControllerExtenderWanExtensionModem1Extension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem1Pdn1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem1Pdn2Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem1Pdn3Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem1Pdn4Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem2Extension(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem2Pdn1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem2Pdn2Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem2Pdn3Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandExtensionControllerExtenderWanExtensionModem2Pdn4Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectExtensionControllerExtender(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allowaccess"); ok || d.HasChange("allowaccess") {
		t, err := expandExtensionControllerExtenderAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("authorized"); ok || d.HasChange("authorized") {
		t, err := expandExtensionControllerExtenderAuthorized(d, v, "authorized")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok || d.HasChange("bandwidth_limit") {
		t, err := expandExtensionControllerExtenderBandwidthLimit(d, v, "bandwidth_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandExtensionControllerExtenderDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("device_id"); ok || d.HasChange("device_id") {
		t, err := expandExtensionControllerExtenderDeviceId(d, v, "device_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok || d.HasChange("enforce_bandwidth") {
		t, err := expandExtensionControllerExtenderEnforceBandwidth(d, v, "enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("ext_name"); ok || d.HasChange("ext_name") {
		t, err := expandExtensionControllerExtenderExtName(d, v, "ext_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-name"] = t
		}
	}

	if v, ok := d.GetOk("extension_type"); ok || d.HasChange("extension_type") {
		t, err := expandExtensionControllerExtenderExtensionType(d, v, "extension_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-type"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok || d.HasChange("firmware_provision_latest") {
		t, err := expandExtensionControllerExtenderFirmwareProvisionLatest(d, v, "firmware_provision_latest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandExtensionControllerExtenderId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok || d.HasChange("login_password") {
		t, err := expandExtensionControllerExtenderLoginPassword(d, v, "login_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok || d.HasChange("login_password_change") {
		t, err := expandExtensionControllerExtenderLoginPasswordChange(d, v, "login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandExtensionControllerExtenderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok || d.HasChange("override_allowaccess") {
		t, err := expandExtensionControllerExtenderOverrideAllowaccess(d, v, "override_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_enforce_bandwidth"); ok || d.HasChange("override_enforce_bandwidth") {
		t, err := expandExtensionControllerExtenderOverrideEnforceBandwidth(d, v, "override_enforce_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("override_login_password_change"); ok || d.HasChange("override_login_password_change") {
		t, err := expandExtensionControllerExtenderOverrideLoginPasswordChange(d, v, "override_login_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandExtensionControllerExtenderProfile(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandExtensionControllerExtenderVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("wan_extension"); ok || d.HasChange("wan_extension") {
		t, err := expandExtensionControllerExtenderWanExtension(d, v, "wan_extension")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-extension"] = t
		}
	}

	return &obj, nil
}
