// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IS-IS.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterIsisUpdate,
		Read:   resourceRouterIsisRead,
		Update: resourceRouterIsisUpdate,
		Delete: resourceRouterIsisDelete,

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
			"adjacency_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"adjacency_check6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_passive_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_passive_only6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_keychain_l1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_keychain_l2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"auth_mode_l1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_mode_l2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_password_l1": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"auth_password_l2": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"auth_sendonly_l1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_sendonly_l2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_originate6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_lsp_errors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isis_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_keychain_l1": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"auth_keychain_l2": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"auth_mode_l1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_mode_l2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_password_l1": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"auth_password_l2": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"auth_send_only_l1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_send_only_l2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"circuit_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"csnp_interval_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"csnp_interval_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hello_interval_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hello_interval_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hello_multiplier_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hello_multiplier_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hello_padding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lsp_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"lsp_retransmit_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mesh_group": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mesh_group_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"metric_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"metric_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wide_metric_l1": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"wide_metric_l2": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"isis_net": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"net": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"lsp_gen_interval_l1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lsp_gen_interval_l2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lsp_refresh_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_lsp_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metric_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overload_bit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overload_bit_on_startup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"overload_bit_suppress": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"redistribute_l1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute_l1_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redistribute_l2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute_l2_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"redistribute6_l1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute6_l1_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redistribute6_l2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redistribute6_l2_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"spf_interval_exp_l1": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"spf_interval_exp_l2": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"summary_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"summary_address6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterIsisUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterIsis(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsis resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterIsis(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsis resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("RouterIsis")

	return resourceRouterIsisRead(d, m)
}

func resourceRouterIsisDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterIsis(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterIsis resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterIsis(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsis resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterIsis(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsis resource from API: %v", err)
	}
	return nil
}

func flattenRouterIsisAdjacencyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAdjacencyCheck6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAdvPassiveOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAdvPassiveOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAuthKeychainL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisAuthKeychainL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisAuthModeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAuthModeL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAuthSendonlyL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisAuthSendonlyL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisDynamicHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIgnoreLspErrors(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_l1"
		if _, ok := i["auth-keychain-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceAuthKeychainL1(i["auth-keychain-l1"], d, pre_append)
			tmp["auth_keychain_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-AuthKeychainL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_l2"
		if _, ok := i["auth-keychain-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceAuthKeychainL2(i["auth-keychain-l2"], d, pre_append)
			tmp["auth_keychain_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-AuthKeychainL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_l1"
		if _, ok := i["auth-mode-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceAuthModeL1(i["auth-mode-l1"], d, pre_append)
			tmp["auth_mode_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-AuthModeL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_l2"
		if _, ok := i["auth-mode-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceAuthModeL2(i["auth-mode-l2"], d, pre_append)
			tmp["auth_mode_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-AuthModeL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_send_only_l1"
		if _, ok := i["auth-send-only-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceAuthSendOnlyL1(i["auth-send-only-l1"], d, pre_append)
			tmp["auth_send_only_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-AuthSendOnlyL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_send_only_l2"
		if _, ok := i["auth-send-only-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceAuthSendOnlyL2(i["auth-send-only-l2"], d, pre_append)
			tmp["auth_send_only_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-AuthSendOnlyL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_type"
		if _, ok := i["circuit-type"]; ok {
			v := flattenRouterIsisIsisInterfaceCircuitType(i["circuit-type"], d, pre_append)
			tmp["circuit_type"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-CircuitType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l1"
		if _, ok := i["csnp-interval-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceCsnpIntervalL1(i["csnp-interval-l1"], d, pre_append)
			tmp["csnp_interval_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-CsnpIntervalL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l2"
		if _, ok := i["csnp-interval-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceCsnpIntervalL2(i["csnp-interval-l2"], d, pre_append)
			tmp["csnp_interval_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-CsnpIntervalL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l1"
		if _, ok := i["hello-interval-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceHelloIntervalL1(i["hello-interval-l1"], d, pre_append)
			tmp["hello_interval_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-HelloIntervalL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l2"
		if _, ok := i["hello-interval-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceHelloIntervalL2(i["hello-interval-l2"], d, pre_append)
			tmp["hello_interval_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-HelloIntervalL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l1"
		if _, ok := i["hello-multiplier-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceHelloMultiplierL1(i["hello-multiplier-l1"], d, pre_append)
			tmp["hello_multiplier_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-HelloMultiplierL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l2"
		if _, ok := i["hello-multiplier-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceHelloMultiplierL2(i["hello-multiplier-l2"], d, pre_append)
			tmp["hello_multiplier_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-HelloMultiplierL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_padding"
		if _, ok := i["hello-padding"]; ok {
			v := flattenRouterIsisIsisInterfaceHelloPadding(i["hello-padding"], d, pre_append)
			tmp["hello_padding"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-HelloPadding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lsp_interval"
		if _, ok := i["lsp-interval"]; ok {
			v := flattenRouterIsisIsisInterfaceLspInterval(i["lsp-interval"], d, pre_append)
			tmp["lsp_interval"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-LspInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lsp_retransmit_interval"
		if _, ok := i["lsp-retransmit-interval"]; ok {
			v := flattenRouterIsisIsisInterfaceLspRetransmitInterval(i["lsp-retransmit-interval"], d, pre_append)
			tmp["lsp_retransmit_interval"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-LspRetransmitInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_group"
		if _, ok := i["mesh-group"]; ok {
			v := flattenRouterIsisIsisInterfaceMeshGroup(i["mesh-group"], d, pre_append)
			tmp["mesh_group"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-MeshGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_group_id"
		if _, ok := i["mesh-group-id"]; ok {
			v := flattenRouterIsisIsisInterfaceMeshGroupId(i["mesh-group-id"], d, pre_append)
			tmp["mesh_group_id"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-MeshGroupId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l1"
		if _, ok := i["metric-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceMetricL1(i["metric-l1"], d, pre_append)
			tmp["metric_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-MetricL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l2"
		if _, ok := i["metric-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceMetricL2(i["metric-l2"], d, pre_append)
			tmp["metric_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-MetricL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenRouterIsisIsisInterfaceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			v := flattenRouterIsisIsisInterfaceNetworkType(i["network-type"], d, pre_append)
			tmp["network_type"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-NetworkType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l1"
		if _, ok := i["priority-l1"]; ok {
			v := flattenRouterIsisIsisInterfacePriorityL1(i["priority-l1"], d, pre_append)
			tmp["priority_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-PriorityL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l2"
		if _, ok := i["priority-l2"]; ok {
			v := flattenRouterIsisIsisInterfacePriorityL2(i["priority-l2"], d, pre_append)
			tmp["priority_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-PriorityL2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenRouterIsisIsisInterfaceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status6"
		if _, ok := i["status6"]; ok {
			v := flattenRouterIsisIsisInterfaceStatus6(i["status6"], d, pre_append)
			tmp["status6"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-Status6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l1"
		if _, ok := i["wide-metric-l1"]; ok {
			v := flattenRouterIsisIsisInterfaceWideMetricL1(i["wide-metric-l1"], d, pre_append)
			tmp["wide_metric_l1"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-WideMetricL1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l2"
		if _, ok := i["wide-metric-l2"]; ok {
			v := flattenRouterIsisIsisInterfaceWideMetricL2(i["wide-metric-l2"], d, pre_append)
			tmp["wide_metric_l2"] = fortiAPISubPartPatch(v, "RouterIsis-IsisInterface-WideMetricL2")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterIsisIsisInterfaceAuthKeychainL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisIsisInterfaceAuthKeychainL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisIsisInterfaceAuthModeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceAuthModeL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceAuthSendOnlyL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceAuthSendOnlyL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceCircuitType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceCsnpIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceCsnpIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloMultiplierL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloMultiplierL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloPadding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceLspInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceLspRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMeshGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMeshGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMetricL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMetricL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisIsisInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfacePriorityL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfacePriorityL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceStatus6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceWideMetricL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceWideMetricL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisNet(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterIsisIsisNetId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterIsis-IsisNet-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net"
		if _, ok := i["net"]; ok {
			v := flattenRouterIsisIsisNetNet(i["net"], d, pre_append)
			tmp["net"] = fortiAPISubPartPatch(v, "RouterIsis-IsisNet-Net")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterIsisIsisNetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisNetNet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisLspGenIntervalL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisLspGenIntervalL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisLspRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisMaxLspLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisMetricStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisOverloadBit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisOverloadBitOnStartup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisOverloadBitSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "level"
	if _, ok := i["level"]; ok {
		result["level"] = flattenRouterIsisRedistributeLevel(i["level"], d, pre_append)
	}

	pre_append = pre + ".0." + "metric"
	if _, ok := i["metric"]; ok {
		result["metric"] = flattenRouterIsisRedistributeMetric(i["metric"], d, pre_append)
	}

	pre_append = pre + ".0." + "metric_type"
	if _, ok := i["metric-type"]; ok {
		result["metric_type"] = flattenRouterIsisRedistributeMetricType(i["metric-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenRouterIsisRedistributeProtocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "routemap"
	if _, ok := i["routemap"]; ok {
		result["routemap"] = flattenRouterIsisRedistributeRoutemap(i["routemap"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenRouterIsisRedistributeStatus(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterIsisRedistributeLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeL1List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistributeL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistributeL2List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistribute6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "level"
	if _, ok := i["level"]; ok {
		result["level"] = flattenRouterIsisRedistribute6Level(i["level"], d, pre_append)
	}

	pre_append = pre + ".0." + "metric"
	if _, ok := i["metric"]; ok {
		result["metric"] = flattenRouterIsisRedistribute6Metric(i["metric"], d, pre_append)
	}

	pre_append = pre + ".0." + "metric_type"
	if _, ok := i["metric-type"]; ok {
		result["metric_type"] = flattenRouterIsisRedistribute6MetricType(i["metric-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "protocol"
	if _, ok := i["protocol"]; ok {
		result["protocol"] = flattenRouterIsisRedistribute6Protocol(i["protocol"], d, pre_append)
	}

	pre_append = pre + ".0." + "routemap"
	if _, ok := i["routemap"]; ok {
		result["routemap"] = flattenRouterIsisRedistribute6Routemap(i["routemap"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenRouterIsisRedistribute6Status(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterIsisRedistribute6Level(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Metric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6MetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6Routemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistribute6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6L1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6L1List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisRedistribute6L2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisRedistribute6L2List(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisSpfIntervalExpL1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenRouterIsisSpfIntervalExpL2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenRouterIsisSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterIsisSummaryAddressId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterIsis-SummaryAddress-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			v := flattenRouterIsisSummaryAddressLevel(i["level"], d, pre_append)
			tmp["level"] = fortiAPISubPartPatch(v, "RouterIsis-SummaryAddress-Level")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenRouterIsisSummaryAddressPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "RouterIsis-SummaryAddress-Prefix")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterIsisSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddressLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisSummaryAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenRouterIsisSummaryAddress6Id(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterIsis-SummaryAddress6-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			v := flattenRouterIsisSummaryAddress6Level(i["level"], d, pre_append)
			tmp["level"] = fortiAPISubPartPatch(v, "RouterIsis-SummaryAddress6-Level")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			v := flattenRouterIsisSummaryAddress6Prefix6(i["prefix6"], d, pre_append)
			tmp["prefix6"] = fortiAPISubPartPatch(v, "RouterIsis-SummaryAddress6-Prefix6")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterIsisSummaryAddress6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6Level(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisSummaryAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterIsis(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("adjacency_check", flattenRouterIsisAdjacencyCheck(o["adjacency-check"], d, "adjacency_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["adjacency-check"], "RouterIsis-AdjacencyCheck"); ok {
			if err = d.Set("adjacency_check", vv); err != nil {
				return fmt.Errorf("Error reading adjacency_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adjacency_check: %v", err)
		}
	}

	if err = d.Set("adjacency_check6", flattenRouterIsisAdjacencyCheck6(o["adjacency-check6"], d, "adjacency_check6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adjacency-check6"], "RouterIsis-AdjacencyCheck6"); ok {
			if err = d.Set("adjacency_check6", vv); err != nil {
				return fmt.Errorf("Error reading adjacency_check6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adjacency_check6: %v", err)
		}
	}

	if err = d.Set("adv_passive_only", flattenRouterIsisAdvPassiveOnly(o["adv-passive-only"], d, "adv_passive_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-passive-only"], "RouterIsis-AdvPassiveOnly"); ok {
			if err = d.Set("adv_passive_only", vv); err != nil {
				return fmt.Errorf("Error reading adv_passive_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_passive_only: %v", err)
		}
	}

	if err = d.Set("adv_passive_only6", flattenRouterIsisAdvPassiveOnly6(o["adv-passive-only6"], d, "adv_passive_only6")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-passive-only6"], "RouterIsis-AdvPassiveOnly6"); ok {
			if err = d.Set("adv_passive_only6", vv); err != nil {
				return fmt.Errorf("Error reading adv_passive_only6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_passive_only6: %v", err)
		}
	}

	if err = d.Set("auth_keychain_l1", flattenRouterIsisAuthKeychainL1(o["auth-keychain-l1"], d, "auth_keychain_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-keychain-l1"], "RouterIsis-AuthKeychainL1"); ok {
			if err = d.Set("auth_keychain_l1", vv); err != nil {
				return fmt.Errorf("Error reading auth_keychain_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_keychain_l1: %v", err)
		}
	}

	if err = d.Set("auth_keychain_l2", flattenRouterIsisAuthKeychainL2(o["auth-keychain-l2"], d, "auth_keychain_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-keychain-l2"], "RouterIsis-AuthKeychainL2"); ok {
			if err = d.Set("auth_keychain_l2", vv); err != nil {
				return fmt.Errorf("Error reading auth_keychain_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_keychain_l2: %v", err)
		}
	}

	if err = d.Set("auth_mode_l1", flattenRouterIsisAuthModeL1(o["auth-mode-l1"], d, "auth_mode_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-mode-l1"], "RouterIsis-AuthModeL1"); ok {
			if err = d.Set("auth_mode_l1", vv); err != nil {
				return fmt.Errorf("Error reading auth_mode_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_mode_l1: %v", err)
		}
	}

	if err = d.Set("auth_mode_l2", flattenRouterIsisAuthModeL2(o["auth-mode-l2"], d, "auth_mode_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-mode-l2"], "RouterIsis-AuthModeL2"); ok {
			if err = d.Set("auth_mode_l2", vv); err != nil {
				return fmt.Errorf("Error reading auth_mode_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_mode_l2: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_l1", flattenRouterIsisAuthSendonlyL1(o["auth-sendonly-l1"], d, "auth_sendonly_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-sendonly-l1"], "RouterIsis-AuthSendonlyL1"); ok {
			if err = d.Set("auth_sendonly_l1", vv); err != nil {
				return fmt.Errorf("Error reading auth_sendonly_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_sendonly_l1: %v", err)
		}
	}

	if err = d.Set("auth_sendonly_l2", flattenRouterIsisAuthSendonlyL2(o["auth-sendonly-l2"], d, "auth_sendonly_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-sendonly-l2"], "RouterIsis-AuthSendonlyL2"); ok {
			if err = d.Set("auth_sendonly_l2", vv); err != nil {
				return fmt.Errorf("Error reading auth_sendonly_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_sendonly_l2: %v", err)
		}
	}

	if err = d.Set("default_originate", flattenRouterIsisDefaultOriginate(o["default-originate"], d, "default_originate")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-originate"], "RouterIsis-DefaultOriginate"); ok {
			if err = d.Set("default_originate", vv); err != nil {
				return fmt.Errorf("Error reading default_originate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_originate: %v", err)
		}
	}

	if err = d.Set("default_originate6", flattenRouterIsisDefaultOriginate6(o["default-originate6"], d, "default_originate6")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-originate6"], "RouterIsis-DefaultOriginate6"); ok {
			if err = d.Set("default_originate6", vv); err != nil {
				return fmt.Errorf("Error reading default_originate6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_originate6: %v", err)
		}
	}

	if err = d.Set("dynamic_hostname", flattenRouterIsisDynamicHostname(o["dynamic-hostname"], d, "dynamic_hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["dynamic-hostname"], "RouterIsis-DynamicHostname"); ok {
			if err = d.Set("dynamic_hostname", vv); err != nil {
				return fmt.Errorf("Error reading dynamic_hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dynamic_hostname: %v", err)
		}
	}

	if err = d.Set("ignore_lsp_errors", flattenRouterIsisIgnoreLspErrors(o["ignore-lsp-errors"], d, "ignore_lsp_errors")); err != nil {
		if vv, ok := fortiAPIPatch(o["ignore-lsp-errors"], "RouterIsis-IgnoreLspErrors"); ok {
			if err = d.Set("ignore_lsp_errors", vv); err != nil {
				return fmt.Errorf("Error reading ignore_lsp_errors: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ignore_lsp_errors: %v", err)
		}
	}

	if err = d.Set("is_type", flattenRouterIsisIsType(o["is-type"], d, "is_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["is-type"], "RouterIsis-IsType"); ok {
			if err = d.Set("is_type", vv); err != nil {
				return fmt.Errorf("Error reading is_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading is_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("isis_interface", flattenRouterIsisIsisInterface(o["isis-interface"], d, "isis_interface")); err != nil {
			if vv, ok := fortiAPIPatch(o["isis-interface"], "RouterIsis-IsisInterface"); ok {
				if err = d.Set("isis_interface", vv); err != nil {
					return fmt.Errorf("Error reading isis_interface: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading isis_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("isis_interface"); ok {
			if err = d.Set("isis_interface", flattenRouterIsisIsisInterface(o["isis-interface"], d, "isis_interface")); err != nil {
				if vv, ok := fortiAPIPatch(o["isis-interface"], "RouterIsis-IsisInterface"); ok {
					if err = d.Set("isis_interface", vv); err != nil {
						return fmt.Errorf("Error reading isis_interface: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading isis_interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("isis_net", flattenRouterIsisIsisNet(o["isis-net"], d, "isis_net")); err != nil {
			if vv, ok := fortiAPIPatch(o["isis-net"], "RouterIsis-IsisNet"); ok {
				if err = d.Set("isis_net", vv); err != nil {
					return fmt.Errorf("Error reading isis_net: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading isis_net: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("isis_net"); ok {
			if err = d.Set("isis_net", flattenRouterIsisIsisNet(o["isis-net"], d, "isis_net")); err != nil {
				if vv, ok := fortiAPIPatch(o["isis-net"], "RouterIsis-IsisNet"); ok {
					if err = d.Set("isis_net", vv); err != nil {
						return fmt.Errorf("Error reading isis_net: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading isis_net: %v", err)
				}
			}
		}
	}

	if err = d.Set("lsp_gen_interval_l1", flattenRouterIsisLspGenIntervalL1(o["lsp-gen-interval-l1"], d, "lsp_gen_interval_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["lsp-gen-interval-l1"], "RouterIsis-LspGenIntervalL1"); ok {
			if err = d.Set("lsp_gen_interval_l1", vv); err != nil {
				return fmt.Errorf("Error reading lsp_gen_interval_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lsp_gen_interval_l1: %v", err)
		}
	}

	if err = d.Set("lsp_gen_interval_l2", flattenRouterIsisLspGenIntervalL2(o["lsp-gen-interval-l2"], d, "lsp_gen_interval_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["lsp-gen-interval-l2"], "RouterIsis-LspGenIntervalL2"); ok {
			if err = d.Set("lsp_gen_interval_l2", vv); err != nil {
				return fmt.Errorf("Error reading lsp_gen_interval_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lsp_gen_interval_l2: %v", err)
		}
	}

	if err = d.Set("lsp_refresh_interval", flattenRouterIsisLspRefreshInterval(o["lsp-refresh-interval"], d, "lsp_refresh_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lsp-refresh-interval"], "RouterIsis-LspRefreshInterval"); ok {
			if err = d.Set("lsp_refresh_interval", vv); err != nil {
				return fmt.Errorf("Error reading lsp_refresh_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lsp_refresh_interval: %v", err)
		}
	}

	if err = d.Set("max_lsp_lifetime", flattenRouterIsisMaxLspLifetime(o["max-lsp-lifetime"], d, "max_lsp_lifetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-lsp-lifetime"], "RouterIsis-MaxLspLifetime"); ok {
			if err = d.Set("max_lsp_lifetime", vv); err != nil {
				return fmt.Errorf("Error reading max_lsp_lifetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_lsp_lifetime: %v", err)
		}
	}

	if err = d.Set("metric_style", flattenRouterIsisMetricStyle(o["metric-style"], d, "metric_style")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric-style"], "RouterIsis-MetricStyle"); ok {
			if err = d.Set("metric_style", vv); err != nil {
				return fmt.Errorf("Error reading metric_style: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric_style: %v", err)
		}
	}

	if err = d.Set("overload_bit", flattenRouterIsisOverloadBit(o["overload-bit"], d, "overload_bit")); err != nil {
		if vv, ok := fortiAPIPatch(o["overload-bit"], "RouterIsis-OverloadBit"); ok {
			if err = d.Set("overload_bit", vv); err != nil {
				return fmt.Errorf("Error reading overload_bit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overload_bit: %v", err)
		}
	}

	if err = d.Set("overload_bit_on_startup", flattenRouterIsisOverloadBitOnStartup(o["overload-bit-on-startup"], d, "overload_bit_on_startup")); err != nil {
		if vv, ok := fortiAPIPatch(o["overload-bit-on-startup"], "RouterIsis-OverloadBitOnStartup"); ok {
			if err = d.Set("overload_bit_on_startup", vv); err != nil {
				return fmt.Errorf("Error reading overload_bit_on_startup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overload_bit_on_startup: %v", err)
		}
	}

	if err = d.Set("overload_bit_suppress", flattenRouterIsisOverloadBitSuppress(o["overload-bit-suppress"], d, "overload_bit_suppress")); err != nil {
		if vv, ok := fortiAPIPatch(o["overload-bit-suppress"], "RouterIsis-OverloadBitSuppress"); ok {
			if err = d.Set("overload_bit_suppress", vv); err != nil {
				return fmt.Errorf("Error reading overload_bit_suppress: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading overload_bit_suppress: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterIsisRedistribute(o["redistribute"], d, "redistribute")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute"], "RouterIsis-Redistribute"); ok {
				if err = d.Set("redistribute", vv); err != nil {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterIsisRedistribute(o["redistribute"], d, "redistribute")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute"], "RouterIsis-Redistribute"); ok {
					if err = d.Set("redistribute", vv); err != nil {
						return fmt.Errorf("Error reading redistribute: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("redistribute_l1", flattenRouterIsisRedistributeL1(o["redistribute-l1"], d, "redistribute_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute-l1"], "RouterIsis-RedistributeL1"); ok {
			if err = d.Set("redistribute_l1", vv); err != nil {
				return fmt.Errorf("Error reading redistribute_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute_l1: %v", err)
		}
	}

	if err = d.Set("redistribute_l1_list", flattenRouterIsisRedistributeL1List(o["redistribute-l1-list"], d, "redistribute_l1_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute-l1-list"], "RouterIsis-RedistributeL1List"); ok {
			if err = d.Set("redistribute_l1_list", vv); err != nil {
				return fmt.Errorf("Error reading redistribute_l1_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute_l1_list: %v", err)
		}
	}

	if err = d.Set("redistribute_l2", flattenRouterIsisRedistributeL2(o["redistribute-l2"], d, "redistribute_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute-l2"], "RouterIsis-RedistributeL2"); ok {
			if err = d.Set("redistribute_l2", vv); err != nil {
				return fmt.Errorf("Error reading redistribute_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute_l2: %v", err)
		}
	}

	if err = d.Set("redistribute_l2_list", flattenRouterIsisRedistributeL2List(o["redistribute-l2-list"], d, "redistribute_l2_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute-l2-list"], "RouterIsis-RedistributeL2List"); ok {
			if err = d.Set("redistribute_l2_list", vv); err != nil {
				return fmt.Errorf("Error reading redistribute_l2_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute_l2_list: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute6", flattenRouterIsisRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
			if vv, ok := fortiAPIPatch(o["redistribute6"], "RouterIsis-Redistribute6"); ok {
				if err = d.Set("redistribute6", vv); err != nil {
					return fmt.Errorf("Error reading redistribute6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading redistribute6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute6"); ok {
			if err = d.Set("redistribute6", flattenRouterIsisRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
				if vv, ok := fortiAPIPatch(o["redistribute6"], "RouterIsis-Redistribute6"); ok {
					if err = d.Set("redistribute6", vv); err != nil {
						return fmt.Errorf("Error reading redistribute6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading redistribute6: %v", err)
				}
			}
		}
	}

	if err = d.Set("redistribute6_l1", flattenRouterIsisRedistribute6L1(o["redistribute6-l1"], d, "redistribute6_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute6-l1"], "RouterIsis-Redistribute6L1"); ok {
			if err = d.Set("redistribute6_l1", vv); err != nil {
				return fmt.Errorf("Error reading redistribute6_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute6_l1: %v", err)
		}
	}

	if err = d.Set("redistribute6_l1_list", flattenRouterIsisRedistribute6L1List(o["redistribute6-l1-list"], d, "redistribute6_l1_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute6-l1-list"], "RouterIsis-Redistribute6L1List"); ok {
			if err = d.Set("redistribute6_l1_list", vv); err != nil {
				return fmt.Errorf("Error reading redistribute6_l1_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute6_l1_list: %v", err)
		}
	}

	if err = d.Set("redistribute6_l2", flattenRouterIsisRedistribute6L2(o["redistribute6-l2"], d, "redistribute6_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute6-l2"], "RouterIsis-Redistribute6L2"); ok {
			if err = d.Set("redistribute6_l2", vv); err != nil {
				return fmt.Errorf("Error reading redistribute6_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute6_l2: %v", err)
		}
	}

	if err = d.Set("redistribute6_l2_list", flattenRouterIsisRedistribute6L2List(o["redistribute6-l2-list"], d, "redistribute6_l2_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["redistribute6-l2-list"], "RouterIsis-Redistribute6L2List"); ok {
			if err = d.Set("redistribute6_l2_list", vv); err != nil {
				return fmt.Errorf("Error reading redistribute6_l2_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redistribute6_l2_list: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l1", flattenRouterIsisSpfIntervalExpL1(o["spf-interval-exp-l1"], d, "spf_interval_exp_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["spf-interval-exp-l1"], "RouterIsis-SpfIntervalExpL1"); ok {
			if err = d.Set("spf_interval_exp_l1", vv); err != nil {
				return fmt.Errorf("Error reading spf_interval_exp_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spf_interval_exp_l1: %v", err)
		}
	}

	if err = d.Set("spf_interval_exp_l2", flattenRouterIsisSpfIntervalExpL2(o["spf-interval-exp-l2"], d, "spf_interval_exp_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["spf-interval-exp-l2"], "RouterIsis-SpfIntervalExpL2"); ok {
			if err = d.Set("spf_interval_exp_l2", vv); err != nil {
				return fmt.Errorf("Error reading spf_interval_exp_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading spf_interval_exp_l2: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address", flattenRouterIsisSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
			if vv, ok := fortiAPIPatch(o["summary-address"], "RouterIsis-SummaryAddress"); ok {
				if err = d.Set("summary_address", vv); err != nil {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterIsisSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
				if vv, ok := fortiAPIPatch(o["summary-address"], "RouterIsis-SummaryAddress"); ok {
					if err = d.Set("summary_address", vv); err != nil {
						return fmt.Errorf("Error reading summary_address: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address6", flattenRouterIsisSummaryAddress6(o["summary-address6"], d, "summary_address6")); err != nil {
			if vv, ok := fortiAPIPatch(o["summary-address6"], "RouterIsis-SummaryAddress6"); ok {
				if err = d.Set("summary_address6", vv); err != nil {
					return fmt.Errorf("Error reading summary_address6: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading summary_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address6"); ok {
			if err = d.Set("summary_address6", flattenRouterIsisSummaryAddress6(o["summary-address6"], d, "summary_address6")); err != nil {
				if vv, ok := fortiAPIPatch(o["summary-address6"], "RouterIsis-SummaryAddress6"); ok {
					if err = d.Set("summary_address6", vv); err != nil {
						return fmt.Errorf("Error reading summary_address6: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading summary_address6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterIsisFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterIsisAdjacencyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAdjacencyCheck6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAdvPassiveOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAdvPassiveOnly6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthKeychainL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisAuthKeychainL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisAuthModeL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthModeL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthPasswordL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisAuthPasswordL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisAuthSendonlyL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisAuthSendonlyL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisDynamicHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIgnoreLspErrors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-keychain-l1"], _ = expandRouterIsisIsisInterfaceAuthKeychainL1(d, i["auth_keychain_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-keychain-l2"], _ = expandRouterIsisIsisInterfaceAuthKeychainL2(d, i["auth_keychain_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-mode-l1"], _ = expandRouterIsisIsisInterfaceAuthModeL1(d, i["auth_mode_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-mode-l2"], _ = expandRouterIsisIsisInterfaceAuthModeL2(d, i["auth_mode_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_password_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-password-l1"], _ = expandRouterIsisIsisInterfaceAuthPasswordL1(d, i["auth_password_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_password_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-password-l2"], _ = expandRouterIsisIsisInterfaceAuthPasswordL2(d, i["auth_password_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_send_only_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-send-only-l1"], _ = expandRouterIsisIsisInterfaceAuthSendOnlyL1(d, i["auth_send_only_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_send_only_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-send-only-l2"], _ = expandRouterIsisIsisInterfaceAuthSendOnlyL2(d, i["auth_send_only_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["circuit-type"], _ = expandRouterIsisIsisInterfaceCircuitType(d, i["circuit_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["csnp-interval-l1"], _ = expandRouterIsisIsisInterfaceCsnpIntervalL1(d, i["csnp_interval_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csnp_interval_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["csnp-interval-l2"], _ = expandRouterIsisIsisInterfaceCsnpIntervalL2(d, i["csnp_interval_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval-l1"], _ = expandRouterIsisIsisInterfaceHelloIntervalL1(d, i["hello_interval_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-interval-l2"], _ = expandRouterIsisIsisInterfaceHelloIntervalL2(d, i["hello_interval_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-multiplier-l1"], _ = expandRouterIsisIsisInterfaceHelloMultiplierL1(d, i["hello_multiplier_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-multiplier-l2"], _ = expandRouterIsisIsisInterfaceHelloMultiplierL2(d, i["hello_multiplier_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_padding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hello-padding"], _ = expandRouterIsisIsisInterfaceHelloPadding(d, i["hello_padding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lsp_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lsp-interval"], _ = expandRouterIsisIsisInterfaceLspInterval(d, i["lsp_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lsp_retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["lsp-retransmit-interval"], _ = expandRouterIsisIsisInterfaceLspRetransmitInterval(d, i["lsp_retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mesh-group"], _ = expandRouterIsisIsisInterfaceMeshGroup(d, i["mesh_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mesh_group_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mesh-group-id"], _ = expandRouterIsisIsisInterfaceMeshGroupId(d, i["mesh_group_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["metric-l1"], _ = expandRouterIsisIsisInterfaceMetricL1(d, i["metric_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["metric-l2"], _ = expandRouterIsisIsisInterfaceMetricL2(d, i["metric_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandRouterIsisIsisInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["network-type"], _ = expandRouterIsisIsisInterfaceNetworkType(d, i["network_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-l1"], _ = expandRouterIsisIsisInterfacePriorityL1(d, i["priority_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-l2"], _ = expandRouterIsisIsisInterfacePriorityL2(d, i["priority_l2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandRouterIsisIsisInterfaceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status6"], _ = expandRouterIsisIsisInterfaceStatus6(d, i["status6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wide-metric-l1"], _ = expandRouterIsisIsisInterfaceWideMetricL1(d, i["wide_metric_l1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wide_metric_l2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["wide-metric-l2"], _ = expandRouterIsisIsisInterfaceWideMetricL2(d, i["wide_metric_l2"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterIsisIsisInterfaceAuthKeychainL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthKeychainL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthModeL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceAuthModeL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceAuthPasswordL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthPasswordL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthSendOnlyL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceAuthSendOnlyL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceCircuitType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceCsnpIntervalL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceCsnpIntervalL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloIntervalL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloIntervalL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloMultiplierL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloMultiplierL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloPadding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceLspInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceLspRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMeshGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMeshGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMetricL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMetricL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfacePriorityL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfacePriorityL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceStatus6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceWideMetricL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceWideMetricL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisNet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterIsisIsisNetId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "net"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["net"], _ = expandRouterIsisIsisNetNet(d, i["net"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterIsisIsisNetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisNetNet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLspGenIntervalL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLspGenIntervalL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisLspRefreshInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisMaxLspLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisMetricStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisOverloadBit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisOverloadBitOnStartup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisOverloadBitSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["level"], _ = expandRouterIsisRedistributeLevel(d, i["level"], pre_append)
	}
	pre_append = pre + ".0." + "metric"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric"], _ = expandRouterIsisRedistributeMetric(d, i["metric"], pre_append)
	}
	pre_append = pre + ".0." + "metric_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric-type"], _ = expandRouterIsisRedistributeMetricType(d, i["metric_type"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandRouterIsisRedistributeProtocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "routemap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["routemap"], _ = expandRouterIsisRedistributeRoutemap(d, i["routemap"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandRouterIsisRedistributeStatus(d, i["status"], pre_append)
	}

	return result, nil
}

func expandRouterIsisRedistributeLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeL1List(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistributeL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistributeL2List(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistribute6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "level"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["level"], _ = expandRouterIsisRedistribute6Level(d, i["level"], pre_append)
	}
	pre_append = pre + ".0." + "metric"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric"], _ = expandRouterIsisRedistribute6Metric(d, i["metric"], pre_append)
	}
	pre_append = pre + ".0." + "metric_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["metric-type"], _ = expandRouterIsisRedistribute6MetricType(d, i["metric_type"], pre_append)
	}
	pre_append = pre + ".0." + "protocol"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["protocol"], _ = expandRouterIsisRedistribute6Protocol(d, i["protocol"], pre_append)
	}
	pre_append = pre + ".0." + "routemap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["routemap"], _ = expandRouterIsisRedistribute6Routemap(d, i["routemap"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandRouterIsisRedistribute6Status(d, i["status"], pre_append)
	}

	return result, nil
}

func expandRouterIsisRedistribute6Level(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Metric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6MetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6Routemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistribute6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6L1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6L1List(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisRedistribute6L2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisRedistribute6L2List(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisSpfIntervalExpL1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandRouterIsisSpfIntervalExpL2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandRouterIsisSummaryAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterIsisSummaryAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["level"], _ = expandRouterIsisSummaryAddressLevel(d, i["level"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandRouterIsisSummaryAddressPrefix(d, i["prefix"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterIsisSummaryAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddressLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddressPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandRouterIsisSummaryAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandRouterIsisSummaryAddress6Id(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["level"], _ = expandRouterIsisSummaryAddress6Level(d, i["level"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix6"], _ = expandRouterIsisSummaryAddress6Prefix6(d, i["prefix6"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterIsisSummaryAddress6Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6Level(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisSummaryAddress6Prefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterIsis(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adjacency_check"); ok || d.HasChange("adjacency_check") {
		t, err := expandRouterIsisAdjacencyCheck(d, v, "adjacency_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adjacency-check"] = t
		}
	}

	if v, ok := d.GetOk("adjacency_check6"); ok || d.HasChange("adjacency_check6") {
		t, err := expandRouterIsisAdjacencyCheck6(d, v, "adjacency_check6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adjacency-check6"] = t
		}
	}

	if v, ok := d.GetOk("adv_passive_only"); ok || d.HasChange("adv_passive_only") {
		t, err := expandRouterIsisAdvPassiveOnly(d, v, "adv_passive_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-passive-only"] = t
		}
	}

	if v, ok := d.GetOk("adv_passive_only6"); ok || d.HasChange("adv_passive_only6") {
		t, err := expandRouterIsisAdvPassiveOnly6(d, v, "adv_passive_only6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-passive-only6"] = t
		}
	}

	if v, ok := d.GetOk("auth_keychain_l1"); ok || d.HasChange("auth_keychain_l1") {
		t, err := expandRouterIsisAuthKeychainL1(d, v, "auth_keychain_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keychain-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_keychain_l2"); ok || d.HasChange("auth_keychain_l2") {
		t, err := expandRouterIsisAuthKeychainL2(d, v, "auth_keychain_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keychain-l2"] = t
		}
	}

	if v, ok := d.GetOk("auth_mode_l1"); ok || d.HasChange("auth_mode_l1") {
		t, err := expandRouterIsisAuthModeL1(d, v, "auth_mode_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-mode-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_mode_l2"); ok || d.HasChange("auth_mode_l2") {
		t, err := expandRouterIsisAuthModeL2(d, v, "auth_mode_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-mode-l2"] = t
		}
	}

	if v, ok := d.GetOk("auth_password_l1"); ok || d.HasChange("auth_password_l1") {
		t, err := expandRouterIsisAuthPasswordL1(d, v, "auth_password_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-password-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_password_l2"); ok || d.HasChange("auth_password_l2") {
		t, err := expandRouterIsisAuthPasswordL2(d, v, "auth_password_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-password-l2"] = t
		}
	}

	if v, ok := d.GetOk("auth_sendonly_l1"); ok || d.HasChange("auth_sendonly_l1") {
		t, err := expandRouterIsisAuthSendonlyL1(d, v, "auth_sendonly_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-sendonly-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_sendonly_l2"); ok || d.HasChange("auth_sendonly_l2") {
		t, err := expandRouterIsisAuthSendonlyL2(d, v, "auth_sendonly_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-sendonly-l2"] = t
		}
	}

	if v, ok := d.GetOk("default_originate"); ok || d.HasChange("default_originate") {
		t, err := expandRouterIsisDefaultOriginate(d, v, "default_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_originate6"); ok || d.HasChange("default_originate6") {
		t, err := expandRouterIsisDefaultOriginate6(d, v, "default_originate6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate6"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_hostname"); ok || d.HasChange("dynamic_hostname") {
		t, err := expandRouterIsisDynamicHostname(d, v, "dynamic_hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-hostname"] = t
		}
	}

	if v, ok := d.GetOk("ignore_lsp_errors"); ok || d.HasChange("ignore_lsp_errors") {
		t, err := expandRouterIsisIgnoreLspErrors(d, v, "ignore_lsp_errors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ignore-lsp-errors"] = t
		}
	}

	if v, ok := d.GetOk("is_type"); ok || d.HasChange("is_type") {
		t, err := expandRouterIsisIsType(d, v, "is_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["is-type"] = t
		}
	}

	if v, ok := d.GetOk("isis_interface"); ok || d.HasChange("isis_interface") {
		t, err := expandRouterIsisIsisInterface(d, v, "isis_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isis-interface"] = t
		}
	}

	if v, ok := d.GetOk("isis_net"); ok || d.HasChange("isis_net") {
		t, err := expandRouterIsisIsisNet(d, v, "isis_net")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isis-net"] = t
		}
	}

	if v, ok := d.GetOk("lsp_gen_interval_l1"); ok || d.HasChange("lsp_gen_interval_l1") {
		t, err := expandRouterIsisLspGenIntervalL1(d, v, "lsp_gen_interval_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lsp-gen-interval-l1"] = t
		}
	}

	if v, ok := d.GetOk("lsp_gen_interval_l2"); ok || d.HasChange("lsp_gen_interval_l2") {
		t, err := expandRouterIsisLspGenIntervalL2(d, v, "lsp_gen_interval_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lsp-gen-interval-l2"] = t
		}
	}

	if v, ok := d.GetOk("lsp_refresh_interval"); ok || d.HasChange("lsp_refresh_interval") {
		t, err := expandRouterIsisLspRefreshInterval(d, v, "lsp_refresh_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lsp-refresh-interval"] = t
		}
	}

	if v, ok := d.GetOk("max_lsp_lifetime"); ok || d.HasChange("max_lsp_lifetime") {
		t, err := expandRouterIsisMaxLspLifetime(d, v, "max_lsp_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-lsp-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("metric_style"); ok || d.HasChange("metric_style") {
		t, err := expandRouterIsisMetricStyle(d, v, "metric_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric-style"] = t
		}
	}

	if v, ok := d.GetOk("overload_bit"); ok || d.HasChange("overload_bit") {
		t, err := expandRouterIsisOverloadBit(d, v, "overload_bit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overload-bit"] = t
		}
	}

	if v, ok := d.GetOk("overload_bit_on_startup"); ok || d.HasChange("overload_bit_on_startup") {
		t, err := expandRouterIsisOverloadBitOnStartup(d, v, "overload_bit_on_startup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overload-bit-on-startup"] = t
		}
	}

	if v, ok := d.GetOk("overload_bit_suppress"); ok || d.HasChange("overload_bit_suppress") {
		t, err := expandRouterIsisOverloadBitSuppress(d, v, "overload_bit_suppress")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overload-bit-suppress"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterIsisRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("redistribute_l1"); ok || d.HasChange("redistribute_l1") {
		t, err := expandRouterIsisRedistributeL1(d, v, "redistribute_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute-l1"] = t
		}
	}

	if v, ok := d.GetOk("redistribute_l1_list"); ok || d.HasChange("redistribute_l1_list") {
		t, err := expandRouterIsisRedistributeL1List(d, v, "redistribute_l1_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute-l1-list"] = t
		}
	}

	if v, ok := d.GetOk("redistribute_l2"); ok || d.HasChange("redistribute_l2") {
		t, err := expandRouterIsisRedistributeL2(d, v, "redistribute_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute-l2"] = t
		}
	}

	if v, ok := d.GetOk("redistribute_l2_list"); ok || d.HasChange("redistribute_l2_list") {
		t, err := expandRouterIsisRedistributeL2List(d, v, "redistribute_l2_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute-l2-list"] = t
		}
	}

	if v, ok := d.GetOk("redistribute6"); ok || d.HasChange("redistribute6") {
		t, err := expandRouterIsisRedistribute6(d, v, "redistribute6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6"] = t
		}
	}

	if v, ok := d.GetOk("redistribute6_l1"); ok || d.HasChange("redistribute6_l1") {
		t, err := expandRouterIsisRedistribute6L1(d, v, "redistribute6_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6-l1"] = t
		}
	}

	if v, ok := d.GetOk("redistribute6_l1_list"); ok || d.HasChange("redistribute6_l1_list") {
		t, err := expandRouterIsisRedistribute6L1List(d, v, "redistribute6_l1_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6-l1-list"] = t
		}
	}

	if v, ok := d.GetOk("redistribute6_l2"); ok || d.HasChange("redistribute6_l2") {
		t, err := expandRouterIsisRedistribute6L2(d, v, "redistribute6_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6-l2"] = t
		}
	}

	if v, ok := d.GetOk("redistribute6_l2_list"); ok || d.HasChange("redistribute6_l2_list") {
		t, err := expandRouterIsisRedistribute6L2List(d, v, "redistribute6_l2_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute6-l2-list"] = t
		}
	}

	if v, ok := d.GetOk("spf_interval_exp_l1"); ok || d.HasChange("spf_interval_exp_l1") {
		t, err := expandRouterIsisSpfIntervalExpL1(d, v, "spf_interval_exp_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spf-interval-exp-l1"] = t
		}
	}

	if v, ok := d.GetOk("spf_interval_exp_l2"); ok || d.HasChange("spf_interval_exp_l2") {
		t, err := expandRouterIsisSpfIntervalExpL2(d, v, "spf_interval_exp_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spf-interval-exp-l2"] = t
		}
	}

	if v, ok := d.GetOk("summary_address"); ok || d.HasChange("summary_address") {
		t, err := expandRouterIsisSummaryAddress(d, v, "summary_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-address"] = t
		}
	}

	if v, ok := d.GetOk("summary_address6"); ok || d.HasChange("summary_address6") {
		t, err := expandRouterIsisSummaryAddress6(d, v, "summary_address6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-address6"] = t
		}
	}

	return &obj, nil
}
