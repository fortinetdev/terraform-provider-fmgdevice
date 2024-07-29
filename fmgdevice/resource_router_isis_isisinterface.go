// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IS-IS interface configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterIsisIsisInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterIsisIsisInterfaceCreate,
		Read:   resourceRouterIsisIsisInterfaceRead,
		Update: resourceRouterIsisIsisInterfaceUpdate,
		Delete: resourceRouterIsisIsisInterfaceDelete,

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
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
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
	}
}

func resourceRouterIsisIsisInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterIsisIsisInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterIsisIsisInterface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterIsisIsisInterface(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating RouterIsisIsisInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterIsisIsisInterfaceRead(d, m)
}

func resourceRouterIsisIsisInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterIsisIsisInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisIsisInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterIsisIsisInterface(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating RouterIsisIsisInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterIsisIsisInterfaceRead(d, m)
}

func resourceRouterIsisIsisInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterIsisIsisInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting RouterIsisIsisInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterIsisIsisInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterIsisIsisInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisIsisInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterIsisIsisInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterIsisIsisInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterIsisIsisInterfaceAuthKeychainL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisIsisInterfaceAuthKeychainL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterIsisIsisInterfaceAuthModeL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceAuthModeL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceAuthSendOnlyL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceAuthSendOnlyL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceCircuitType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceCsnpIntervalL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceCsnpIntervalL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloIntervalL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloIntervalL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloMultiplierL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloMultiplierL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceHelloPadding2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceLspInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceLspRetransmitInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMeshGroup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMeshGroupId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMetricL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceMetricL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceNetworkType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfacePriorityL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfacePriorityL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceStatus62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceWideMetricL12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterIsisIsisInterfaceWideMetricL22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterIsisIsisInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_keychain_l1", flattenRouterIsisIsisInterfaceAuthKeychainL12edl(o["auth-keychain-l1"], d, "auth_keychain_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-keychain-l1"], "RouterIsisIsisInterface-AuthKeychainL1"); ok {
			if err = d.Set("auth_keychain_l1", vv); err != nil {
				return fmt.Errorf("Error reading auth_keychain_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_keychain_l1: %v", err)
		}
	}

	if err = d.Set("auth_keychain_l2", flattenRouterIsisIsisInterfaceAuthKeychainL22edl(o["auth-keychain-l2"], d, "auth_keychain_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-keychain-l2"], "RouterIsisIsisInterface-AuthKeychainL2"); ok {
			if err = d.Set("auth_keychain_l2", vv); err != nil {
				return fmt.Errorf("Error reading auth_keychain_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_keychain_l2: %v", err)
		}
	}

	if err = d.Set("auth_mode_l1", flattenRouterIsisIsisInterfaceAuthModeL12edl(o["auth-mode-l1"], d, "auth_mode_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-mode-l1"], "RouterIsisIsisInterface-AuthModeL1"); ok {
			if err = d.Set("auth_mode_l1", vv); err != nil {
				return fmt.Errorf("Error reading auth_mode_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_mode_l1: %v", err)
		}
	}

	if err = d.Set("auth_mode_l2", flattenRouterIsisIsisInterfaceAuthModeL22edl(o["auth-mode-l2"], d, "auth_mode_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-mode-l2"], "RouterIsisIsisInterface-AuthModeL2"); ok {
			if err = d.Set("auth_mode_l2", vv); err != nil {
				return fmt.Errorf("Error reading auth_mode_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_mode_l2: %v", err)
		}
	}

	if err = d.Set("auth_send_only_l1", flattenRouterIsisIsisInterfaceAuthSendOnlyL12edl(o["auth-send-only-l1"], d, "auth_send_only_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-send-only-l1"], "RouterIsisIsisInterface-AuthSendOnlyL1"); ok {
			if err = d.Set("auth_send_only_l1", vv); err != nil {
				return fmt.Errorf("Error reading auth_send_only_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_send_only_l1: %v", err)
		}
	}

	if err = d.Set("auth_send_only_l2", flattenRouterIsisIsisInterfaceAuthSendOnlyL22edl(o["auth-send-only-l2"], d, "auth_send_only_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-send-only-l2"], "RouterIsisIsisInterface-AuthSendOnlyL2"); ok {
			if err = d.Set("auth_send_only_l2", vv); err != nil {
				return fmt.Errorf("Error reading auth_send_only_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_send_only_l2: %v", err)
		}
	}

	if err = d.Set("circuit_type", flattenRouterIsisIsisInterfaceCircuitType2edl(o["circuit-type"], d, "circuit_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["circuit-type"], "RouterIsisIsisInterface-CircuitType"); ok {
			if err = d.Set("circuit_type", vv); err != nil {
				return fmt.Errorf("Error reading circuit_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading circuit_type: %v", err)
		}
	}

	if err = d.Set("csnp_interval_l1", flattenRouterIsisIsisInterfaceCsnpIntervalL12edl(o["csnp-interval-l1"], d, "csnp_interval_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["csnp-interval-l1"], "RouterIsisIsisInterface-CsnpIntervalL1"); ok {
			if err = d.Set("csnp_interval_l1", vv); err != nil {
				return fmt.Errorf("Error reading csnp_interval_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csnp_interval_l1: %v", err)
		}
	}

	if err = d.Set("csnp_interval_l2", flattenRouterIsisIsisInterfaceCsnpIntervalL22edl(o["csnp-interval-l2"], d, "csnp_interval_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["csnp-interval-l2"], "RouterIsisIsisInterface-CsnpIntervalL2"); ok {
			if err = d.Set("csnp_interval_l2", vv); err != nil {
				return fmt.Errorf("Error reading csnp_interval_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading csnp_interval_l2: %v", err)
		}
	}

	if err = d.Set("hello_interval_l1", flattenRouterIsisIsisInterfaceHelloIntervalL12edl(o["hello-interval-l1"], d, "hello_interval_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval-l1"], "RouterIsisIsisInterface-HelloIntervalL1"); ok {
			if err = d.Set("hello_interval_l1", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval_l1: %v", err)
		}
	}

	if err = d.Set("hello_interval_l2", flattenRouterIsisIsisInterfaceHelloIntervalL22edl(o["hello-interval-l2"], d, "hello_interval_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval-l2"], "RouterIsisIsisInterface-HelloIntervalL2"); ok {
			if err = d.Set("hello_interval_l2", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval_l2: %v", err)
		}
	}

	if err = d.Set("hello_multiplier_l1", flattenRouterIsisIsisInterfaceHelloMultiplierL12edl(o["hello-multiplier-l1"], d, "hello_multiplier_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-multiplier-l1"], "RouterIsisIsisInterface-HelloMultiplierL1"); ok {
			if err = d.Set("hello_multiplier_l1", vv); err != nil {
				return fmt.Errorf("Error reading hello_multiplier_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_multiplier_l1: %v", err)
		}
	}

	if err = d.Set("hello_multiplier_l2", flattenRouterIsisIsisInterfaceHelloMultiplierL22edl(o["hello-multiplier-l2"], d, "hello_multiplier_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-multiplier-l2"], "RouterIsisIsisInterface-HelloMultiplierL2"); ok {
			if err = d.Set("hello_multiplier_l2", vv); err != nil {
				return fmt.Errorf("Error reading hello_multiplier_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_multiplier_l2: %v", err)
		}
	}

	if err = d.Set("hello_padding", flattenRouterIsisIsisInterfaceHelloPadding2edl(o["hello-padding"], d, "hello_padding")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-padding"], "RouterIsisIsisInterface-HelloPadding"); ok {
			if err = d.Set("hello_padding", vv); err != nil {
				return fmt.Errorf("Error reading hello_padding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_padding: %v", err)
		}
	}

	if err = d.Set("lsp_interval", flattenRouterIsisIsisInterfaceLspInterval2edl(o["lsp-interval"], d, "lsp_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lsp-interval"], "RouterIsisIsisInterface-LspInterval"); ok {
			if err = d.Set("lsp_interval", vv); err != nil {
				return fmt.Errorf("Error reading lsp_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lsp_interval: %v", err)
		}
	}

	if err = d.Set("lsp_retransmit_interval", flattenRouterIsisIsisInterfaceLspRetransmitInterval2edl(o["lsp-retransmit-interval"], d, "lsp_retransmit_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["lsp-retransmit-interval"], "RouterIsisIsisInterface-LspRetransmitInterval"); ok {
			if err = d.Set("lsp_retransmit_interval", vv); err != nil {
				return fmt.Errorf("Error reading lsp_retransmit_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lsp_retransmit_interval: %v", err)
		}
	}

	if err = d.Set("mesh_group", flattenRouterIsisIsisInterfaceMeshGroup2edl(o["mesh-group"], d, "mesh_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-group"], "RouterIsisIsisInterface-MeshGroup"); ok {
			if err = d.Set("mesh_group", vv); err != nil {
				return fmt.Errorf("Error reading mesh_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_group: %v", err)
		}
	}

	if err = d.Set("mesh_group_id", flattenRouterIsisIsisInterfaceMeshGroupId2edl(o["mesh-group-id"], d, "mesh_group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["mesh-group-id"], "RouterIsisIsisInterface-MeshGroupId"); ok {
			if err = d.Set("mesh_group_id", vv); err != nil {
				return fmt.Errorf("Error reading mesh_group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mesh_group_id: %v", err)
		}
	}

	if err = d.Set("metric_l1", flattenRouterIsisIsisInterfaceMetricL12edl(o["metric-l1"], d, "metric_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric-l1"], "RouterIsisIsisInterface-MetricL1"); ok {
			if err = d.Set("metric_l1", vv); err != nil {
				return fmt.Errorf("Error reading metric_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric_l1: %v", err)
		}
	}

	if err = d.Set("metric_l2", flattenRouterIsisIsisInterfaceMetricL22edl(o["metric-l2"], d, "metric_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["metric-l2"], "RouterIsisIsisInterface-MetricL2"); ok {
			if err = d.Set("metric_l2", vv); err != nil {
				return fmt.Errorf("Error reading metric_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metric_l2: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterIsisIsisInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterIsisIsisInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("network_type", flattenRouterIsisIsisInterfaceNetworkType2edl(o["network-type"], d, "network_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-type"], "RouterIsisIsisInterface-NetworkType"); ok {
			if err = d.Set("network_type", vv); err != nil {
				return fmt.Errorf("Error reading network_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("priority_l1", flattenRouterIsisIsisInterfacePriorityL12edl(o["priority-l1"], d, "priority_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-l1"], "RouterIsisIsisInterface-PriorityL1"); ok {
			if err = d.Set("priority_l1", vv); err != nil {
				return fmt.Errorf("Error reading priority_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_l1: %v", err)
		}
	}

	if err = d.Set("priority_l2", flattenRouterIsisIsisInterfacePriorityL22edl(o["priority-l2"], d, "priority_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-l2"], "RouterIsisIsisInterface-PriorityL2"); ok {
			if err = d.Set("priority_l2", vv); err != nil {
				return fmt.Errorf("Error reading priority_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_l2: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterIsisIsisInterfaceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterIsisIsisInterface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("status6", flattenRouterIsisIsisInterfaceStatus62edl(o["status6"], d, "status6")); err != nil {
		if vv, ok := fortiAPIPatch(o["status6"], "RouterIsisIsisInterface-Status6"); ok {
			if err = d.Set("status6", vv); err != nil {
				return fmt.Errorf("Error reading status6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status6: %v", err)
		}
	}

	if err = d.Set("wide_metric_l1", flattenRouterIsisIsisInterfaceWideMetricL12edl(o["wide-metric-l1"], d, "wide_metric_l1")); err != nil {
		if vv, ok := fortiAPIPatch(o["wide-metric-l1"], "RouterIsisIsisInterface-WideMetricL1"); ok {
			if err = d.Set("wide_metric_l1", vv); err != nil {
				return fmt.Errorf("Error reading wide_metric_l1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wide_metric_l1: %v", err)
		}
	}

	if err = d.Set("wide_metric_l2", flattenRouterIsisIsisInterfaceWideMetricL22edl(o["wide-metric-l2"], d, "wide_metric_l2")); err != nil {
		if vv, ok := fortiAPIPatch(o["wide-metric-l2"], "RouterIsisIsisInterface-WideMetricL2"); ok {
			if err = d.Set("wide_metric_l2", vv); err != nil {
				return fmt.Errorf("Error reading wide_metric_l2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wide_metric_l2: %v", err)
		}
	}

	return nil
}

func flattenRouterIsisIsisInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterIsisIsisInterfaceAuthKeychainL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthKeychainL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthModeL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceAuthModeL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceAuthPasswordL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthPasswordL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterIsisIsisInterfaceAuthSendOnlyL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceAuthSendOnlyL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceCircuitType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceCsnpIntervalL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceCsnpIntervalL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloIntervalL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloIntervalL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloMultiplierL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloMultiplierL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceHelloPadding2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceLspInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceLspRetransmitInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMeshGroup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMeshGroupId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMetricL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceMetricL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceNetworkType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfacePriorityL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfacePriorityL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceStatus62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceWideMetricL12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterIsisIsisInterfaceWideMetricL22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterIsisIsisInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_keychain_l1"); ok || d.HasChange("auth_keychain_l1") {
		t, err := expandRouterIsisIsisInterfaceAuthKeychainL12edl(d, v, "auth_keychain_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keychain-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_keychain_l2"); ok || d.HasChange("auth_keychain_l2") {
		t, err := expandRouterIsisIsisInterfaceAuthKeychainL22edl(d, v, "auth_keychain_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-keychain-l2"] = t
		}
	}

	if v, ok := d.GetOk("auth_mode_l1"); ok || d.HasChange("auth_mode_l1") {
		t, err := expandRouterIsisIsisInterfaceAuthModeL12edl(d, v, "auth_mode_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-mode-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_mode_l2"); ok || d.HasChange("auth_mode_l2") {
		t, err := expandRouterIsisIsisInterfaceAuthModeL22edl(d, v, "auth_mode_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-mode-l2"] = t
		}
	}

	if v, ok := d.GetOk("auth_password_l1"); ok || d.HasChange("auth_password_l1") {
		t, err := expandRouterIsisIsisInterfaceAuthPasswordL12edl(d, v, "auth_password_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-password-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_password_l2"); ok || d.HasChange("auth_password_l2") {
		t, err := expandRouterIsisIsisInterfaceAuthPasswordL22edl(d, v, "auth_password_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-password-l2"] = t
		}
	}

	if v, ok := d.GetOk("auth_send_only_l1"); ok || d.HasChange("auth_send_only_l1") {
		t, err := expandRouterIsisIsisInterfaceAuthSendOnlyL12edl(d, v, "auth_send_only_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-send-only-l1"] = t
		}
	}

	if v, ok := d.GetOk("auth_send_only_l2"); ok || d.HasChange("auth_send_only_l2") {
		t, err := expandRouterIsisIsisInterfaceAuthSendOnlyL22edl(d, v, "auth_send_only_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-send-only-l2"] = t
		}
	}

	if v, ok := d.GetOk("circuit_type"); ok || d.HasChange("circuit_type") {
		t, err := expandRouterIsisIsisInterfaceCircuitType2edl(d, v, "circuit_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["circuit-type"] = t
		}
	}

	if v, ok := d.GetOk("csnp_interval_l1"); ok || d.HasChange("csnp_interval_l1") {
		t, err := expandRouterIsisIsisInterfaceCsnpIntervalL12edl(d, v, "csnp_interval_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csnp-interval-l1"] = t
		}
	}

	if v, ok := d.GetOk("csnp_interval_l2"); ok || d.HasChange("csnp_interval_l2") {
		t, err := expandRouterIsisIsisInterfaceCsnpIntervalL22edl(d, v, "csnp_interval_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csnp-interval-l2"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval_l1"); ok || d.HasChange("hello_interval_l1") {
		t, err := expandRouterIsisIsisInterfaceHelloIntervalL12edl(d, v, "hello_interval_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval-l1"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval_l2"); ok || d.HasChange("hello_interval_l2") {
		t, err := expandRouterIsisIsisInterfaceHelloIntervalL22edl(d, v, "hello_interval_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval-l2"] = t
		}
	}

	if v, ok := d.GetOk("hello_multiplier_l1"); ok || d.HasChange("hello_multiplier_l1") {
		t, err := expandRouterIsisIsisInterfaceHelloMultiplierL12edl(d, v, "hello_multiplier_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-multiplier-l1"] = t
		}
	}

	if v, ok := d.GetOk("hello_multiplier_l2"); ok || d.HasChange("hello_multiplier_l2") {
		t, err := expandRouterIsisIsisInterfaceHelloMultiplierL22edl(d, v, "hello_multiplier_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-multiplier-l2"] = t
		}
	}

	if v, ok := d.GetOk("hello_padding"); ok || d.HasChange("hello_padding") {
		t, err := expandRouterIsisIsisInterfaceHelloPadding2edl(d, v, "hello_padding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-padding"] = t
		}
	}

	if v, ok := d.GetOk("lsp_interval"); ok || d.HasChange("lsp_interval") {
		t, err := expandRouterIsisIsisInterfaceLspInterval2edl(d, v, "lsp_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lsp-interval"] = t
		}
	}

	if v, ok := d.GetOk("lsp_retransmit_interval"); ok || d.HasChange("lsp_retransmit_interval") {
		t, err := expandRouterIsisIsisInterfaceLspRetransmitInterval2edl(d, v, "lsp_retransmit_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lsp-retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("mesh_group"); ok || d.HasChange("mesh_group") {
		t, err := expandRouterIsisIsisInterfaceMeshGroup2edl(d, v, "mesh_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-group"] = t
		}
	}

	if v, ok := d.GetOk("mesh_group_id"); ok || d.HasChange("mesh_group_id") {
		t, err := expandRouterIsisIsisInterfaceMeshGroupId2edl(d, v, "mesh_group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-group-id"] = t
		}
	}

	if v, ok := d.GetOk("metric_l1"); ok || d.HasChange("metric_l1") {
		t, err := expandRouterIsisIsisInterfaceMetricL12edl(d, v, "metric_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric-l1"] = t
		}
	}

	if v, ok := d.GetOk("metric_l2"); ok || d.HasChange("metric_l2") {
		t, err := expandRouterIsisIsisInterfaceMetricL22edl(d, v, "metric_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metric-l2"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterIsisIsisInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok || d.HasChange("network_type") {
		t, err := expandRouterIsisIsisInterfaceNetworkType2edl(d, v, "network_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}

	if v, ok := d.GetOk("priority_l1"); ok || d.HasChange("priority_l1") {
		t, err := expandRouterIsisIsisInterfacePriorityL12edl(d, v, "priority_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-l1"] = t
		}
	}

	if v, ok := d.GetOk("priority_l2"); ok || d.HasChange("priority_l2") {
		t, err := expandRouterIsisIsisInterfacePriorityL22edl(d, v, "priority_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-l2"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterIsisIsisInterfaceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("status6"); ok || d.HasChange("status6") {
		t, err := expandRouterIsisIsisInterfaceStatus62edl(d, v, "status6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status6"] = t
		}
	}

	if v, ok := d.GetOk("wide_metric_l1"); ok || d.HasChange("wide_metric_l1") {
		t, err := expandRouterIsisIsisInterfaceWideMetricL12edl(d, v, "wide_metric_l1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wide-metric-l1"] = t
		}
	}

	if v, ok := d.GetOk("wide_metric_l2"); ok || d.HasChange("wide_metric_l2") {
		t, err := expandRouterIsisIsisInterfaceWideMetricL22edl(d, v, "wide_metric_l2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wide-metric-l2"] = t
		}
	}

	return &obj, nil
}
