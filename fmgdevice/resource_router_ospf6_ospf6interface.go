// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF6 interface configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspf6Ospf6Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6Ospf6InterfaceCreate,
		Read:   resourceRouterOspf6Ospf6InterfaceRead,
		Update: resourceRouterOspf6Ospf6InterfaceUpdate,
		Delete: resourceRouterOspf6Ospf6InterfaceDelete,

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
			"area_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipsec_auth_alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_enc_alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"enc_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"spi": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"key_rollover_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mtu_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poll_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retransmit_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transmit_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterOspf6Ospf6InterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspf6Ospf6Interface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Ospf6Interface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspf6Ospf6Interface(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspf6Ospf6Interface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspf6Ospf6InterfaceRead(d, m)
}

func resourceRouterOspf6Ospf6InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspf6Ospf6Interface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Ospf6Interface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6Ospf6Interface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6Ospf6Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspf6Ospf6InterfaceRead(d, m)
}

func resourceRouterOspf6Ospf6InterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspf6Ospf6Interface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6Ospf6Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Ospf6InterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspf6Ospf6Interface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Ospf6Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6Ospf6Interface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6Ospf6Interface resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6Ospf6InterfaceAreaId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceDeadInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceHelloInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspf6Ospf6InterfaceIpsecAuthAlg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceIpsecEncAlg2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeys2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi2edl(i["spi"], d, pre_append)
			tmp["spi"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-IpsecKeys-Spi")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceKeyRolloverInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtuIgnore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighbor2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborCost2edl(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborIp62edl(i["ip6"], d, pre_append)
			tmp["ip6"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-Ip6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborPollInterval2edl(i["poll-interval"], d, pre_append)
			tmp["poll_interval"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-PollInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenRouterOspf6Ospf6InterfaceNeighborPriority2edl(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "RouterOspf6Ospf6Interface-Neighbor-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceNeighborCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborIp62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPollInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNetworkType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfacePriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceRetransmitInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceTransmitDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf6Ospf6Interface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("area_id", flattenRouterOspf6Ospf6InterfaceAreaId2edl(o["area-id"], d, "area_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["area-id"], "RouterOspf6Ospf6Interface-AreaId"); ok {
			if err = d.Set("area_id", vv); err != nil {
				return fmt.Errorf("Error reading area_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading area_id: %v", err)
		}
	}

	if err = d.Set("authentication", flattenRouterOspf6Ospf6InterfaceAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "RouterOspf6Ospf6Interface-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspf6Ospf6InterfaceBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterOspf6Ospf6Interface-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterOspf6Ospf6InterfaceCost2edl(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "RouterOspf6Ospf6Interface-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterOspf6Ospf6InterfaceDeadInterval2edl(o["dead-interval"], d, "dead_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dead-interval"], "RouterOspf6Ospf6Interface-DeadInterval"); ok {
			if err = d.Set("dead_interval", vv); err != nil {
				return fmt.Errorf("Error reading dead_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterOspf6Ospf6InterfaceHelloInterval2edl(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "RouterOspf6Ospf6Interface-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterOspf6Ospf6InterfaceInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "RouterOspf6Ospf6Interface-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipsec_auth_alg", flattenRouterOspf6Ospf6InterfaceIpsecAuthAlg2edl(o["ipsec-auth-alg"], d, "ipsec_auth_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-auth-alg"], "RouterOspf6Ospf6Interface-IpsecAuthAlg"); ok {
			if err = d.Set("ipsec_auth_alg", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_auth_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_auth_alg: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_alg", flattenRouterOspf6Ospf6InterfaceIpsecEncAlg2edl(o["ipsec-enc-alg"], d, "ipsec_enc_alg")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-enc-alg"], "RouterOspf6Ospf6Interface-IpsecEncAlg"); ok {
			if err = d.Set("ipsec_enc_alg", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_enc_alg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_enc_alg: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipsec_keys", flattenRouterOspf6Ospf6InterfaceIpsecKeys2edl(o["ipsec-keys"], d, "ipsec_keys")); err != nil {
			if vv, ok := fortiAPIPatch(o["ipsec-keys"], "RouterOspf6Ospf6Interface-IpsecKeys"); ok {
				if err = d.Set("ipsec_keys", vv); err != nil {
					return fmt.Errorf("Error reading ipsec_keys: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ipsec_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipsec_keys"); ok {
			if err = d.Set("ipsec_keys", flattenRouterOspf6Ospf6InterfaceIpsecKeys2edl(o["ipsec-keys"], d, "ipsec_keys")); err != nil {
				if vv, ok := fortiAPIPatch(o["ipsec-keys"], "RouterOspf6Ospf6Interface-IpsecKeys"); ok {
					if err = d.Set("ipsec_keys", vv); err != nil {
						return fmt.Errorf("Error reading ipsec_keys: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ipsec_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("key_rollover_interval", flattenRouterOspf6Ospf6InterfaceKeyRolloverInterval2edl(o["key-rollover-interval"], d, "key_rollover_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-rollover-interval"], "RouterOspf6Ospf6Interface-KeyRolloverInterval"); ok {
			if err = d.Set("key_rollover_interval", vv); err != nil {
				return fmt.Errorf("Error reading key_rollover_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_rollover_interval: %v", err)
		}
	}

	if err = d.Set("mtu", flattenRouterOspf6Ospf6InterfaceMtu2edl(o["mtu"], d, "mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu"], "RouterOspf6Ospf6Interface-Mtu"); ok {
			if err = d.Set("mtu", vv); err != nil {
				return fmt.Errorf("Error reading mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("mtu_ignore", flattenRouterOspf6Ospf6InterfaceMtuIgnore2edl(o["mtu-ignore"], d, "mtu_ignore")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu-ignore"], "RouterOspf6Ospf6Interface-MtuIgnore"); ok {
			if err = d.Set("mtu_ignore", vv); err != nil {
				return fmt.Errorf("Error reading mtu_ignore: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu_ignore: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterOspf6Ospf6InterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterOspf6Ospf6Interface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterOspf6Ospf6InterfaceNeighbor2edl(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "RouterOspf6Ospf6Interface-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterOspf6Ospf6InterfaceNeighbor2edl(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "RouterOspf6Ospf6Interface-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if err = d.Set("network_type", flattenRouterOspf6Ospf6InterfaceNetworkType2edl(o["network-type"], d, "network_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-type"], "RouterOspf6Ospf6Interface-NetworkType"); ok {
			if err = d.Set("network_type", vv); err != nil {
				return fmt.Errorf("Error reading network_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterOspf6Ospf6InterfacePriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "RouterOspf6Ospf6Interface-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterOspf6Ospf6InterfaceRetransmitInterval2edl(o["retransmit-interval"], d, "retransmit_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["retransmit-interval"], "RouterOspf6Ospf6Interface-RetransmitInterval"); ok {
			if err = d.Set("retransmit_interval", vv); err != nil {
				return fmt.Errorf("Error reading retransmit_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterOspf6Ospf6InterfaceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterOspf6Ospf6Interface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterOspf6Ospf6InterfaceTransmitDelay2edl(o["transmit-delay"], d, "transmit_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["transmit-delay"], "RouterOspf6Ospf6Interface-TransmitDelay"); ok {
			if err = d.Set("transmit_delay", vv); err != nil {
				return fmt.Errorf("Error reading transmit_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	return nil
}

func flattenRouterOspf6Ospf6InterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspf6Ospf6InterfaceAreaId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceDeadInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceHelloInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecAuthAlg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecEncAlg2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-key"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey2edl(d, i["auth_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["enc-key"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey2edl(d, i["enc_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spi"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysSpi2edl(d, i["spi"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysSpi2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceKeyRolloverInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtuIgnore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighbor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceNeighborCost2edl(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip6"], _ = expandRouterOspf6Ospf6InterfaceNeighborIp62edl(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["poll-interval"], _ = expandRouterOspf6Ospf6InterfaceNeighborPollInterval2edl(d, i["poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfaceNeighborPriority2edl(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborIp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPollInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNetworkType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfacePriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceRetransmitInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceTransmitDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6Ospf6Interface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("area_id"); ok || d.HasChange("area_id") {
		t, err := expandRouterOspf6Ospf6InterfaceAreaId2edl(d, v, "area_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area-id"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandRouterOspf6Ospf6InterfaceAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterOspf6Ospf6InterfaceBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandRouterOspf6Ospf6InterfaceCost2edl(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok || d.HasChange("dead_interval") {
		t, err := expandRouterOspf6Ospf6InterfaceDeadInterval2edl(d, v, "dead_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandRouterOspf6Ospf6InterfaceHelloInterval2edl(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterOspf6Ospf6InterfaceInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_auth_alg"); ok || d.HasChange("ipsec_auth_alg") {
		t, err := expandRouterOspf6Ospf6InterfaceIpsecAuthAlg2edl(d, v, "ipsec_auth_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-auth-alg"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_enc_alg"); ok || d.HasChange("ipsec_enc_alg") {
		t, err := expandRouterOspf6Ospf6InterfaceIpsecEncAlg2edl(d, v, "ipsec_enc_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-enc-alg"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_keys"); ok || d.HasChange("ipsec_keys") {
		t, err := expandRouterOspf6Ospf6InterfaceIpsecKeys2edl(d, v, "ipsec_keys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-keys"] = t
		}
	}

	if v, ok := d.GetOk("key_rollover_interval"); ok || d.HasChange("key_rollover_interval") {
		t, err := expandRouterOspf6Ospf6InterfaceKeyRolloverInterval2edl(d, v, "key_rollover_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-rollover-interval"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok || d.HasChange("mtu") {
		t, err := expandRouterOspf6Ospf6InterfaceMtu2edl(d, v, "mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("mtu_ignore"); ok || d.HasChange("mtu_ignore") {
		t, err := expandRouterOspf6Ospf6InterfaceMtuIgnore2edl(d, v, "mtu_ignore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-ignore"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterOspf6Ospf6InterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandRouterOspf6Ospf6InterfaceNeighbor2edl(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok || d.HasChange("network_type") {
		t, err := expandRouterOspf6Ospf6InterfaceNetworkType2edl(d, v, "network_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandRouterOspf6Ospf6InterfacePriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok || d.HasChange("retransmit_interval") {
		t, err := expandRouterOspf6Ospf6InterfaceRetransmitInterval2edl(d, v, "retransmit_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterOspf6Ospf6InterfaceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok || d.HasChange("transmit_delay") {
		t, err := expandRouterOspf6Ospf6InterfaceTransmitDelay2edl(d, v, "transmit_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	return &obj, nil
}
