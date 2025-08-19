// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: OSPF interface configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterOspfOspfInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfOspfInterfaceCreate,
		Read:   resourceRouterOspfOspfInterfaceRead,
		Update: resourceRouterOspfOspfInterfaceUpdate,
		Delete: resourceRouterOspfOspfInterfaceDelete,

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
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"database_filter_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"hello_multiplier": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keychain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"linkdown_fast_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"md5_keychain": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"md5_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"key_string": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
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
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resync_timeout": &schema.Schema{
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

func resourceRouterOspfOspfInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfOspfInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfOspfInterface resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfOspfInterface(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfOspfInterface resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspfOspfInterfaceRead(d, m)
}

func resourceRouterOspfOspfInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectRouterOspfOspfInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfOspfInterface resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfOspfInterface(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfOspfInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceRouterOspfOspfInterfaceRead(d, m)
}

func resourceRouterOspfOspfInterfaceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteRouterOspfOspfInterface(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfOspfInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfOspfInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspfOspfInterface(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfOspfInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfOspfInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfOspfInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfOspfInterfaceAuthentication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceBfd2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceComments2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceCost2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceDatabaseFilterOut2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceDeadInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceHelloInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceHelloMultiplier2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceInterface2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterfaceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceKeychain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterfaceLinkdownFastFailover2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMd5Keychain2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenRouterOspfOspfInterfaceMd5Keys2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenRouterOspfOspfInterfaceMd5KeysId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "RouterOspfOspfInterface-Md5Keys-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenRouterOspfOspfInterfaceMd5KeysId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtu2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtuIgnore2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceNetworkType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePrefixLength2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceResyncTimeout2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceRetransmitInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceTransmitDelay2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspfOspfInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authentication", flattenRouterOspfOspfInterfaceAuthentication2edl(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "RouterOspfOspfInterface-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspfOspfInterfaceBfd2edl(o["bfd"], d, "bfd")); err != nil {
		if vv, ok := fortiAPIPatch(o["bfd"], "RouterOspfOspfInterface-Bfd"); ok {
			if err = d.Set("bfd", vv); err != nil {
				return fmt.Errorf("Error reading bfd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterOspfOspfInterfaceComments2edl(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "RouterOspfOspfInterface-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterOspfOspfInterfaceCost2edl(o["cost"], d, "cost")); err != nil {
		if vv, ok := fortiAPIPatch(o["cost"], "RouterOspfOspfInterface-Cost"); ok {
			if err = d.Set("cost", vv); err != nil {
				return fmt.Errorf("Error reading cost: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("database_filter_out", flattenRouterOspfOspfInterfaceDatabaseFilterOut2edl(o["database-filter-out"], d, "database_filter_out")); err != nil {
		if vv, ok := fortiAPIPatch(o["database-filter-out"], "RouterOspfOspfInterface-DatabaseFilterOut"); ok {
			if err = d.Set("database_filter_out", vv); err != nil {
				return fmt.Errorf("Error reading database_filter_out: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading database_filter_out: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterOspfOspfInterfaceDeadInterval2edl(o["dead-interval"], d, "dead_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["dead-interval"], "RouterOspfOspfInterface-DeadInterval"); ok {
			if err = d.Set("dead_interval", vv); err != nil {
				return fmt.Errorf("Error reading dead_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterOspfOspfInterfaceHelloInterval2edl(o["hello-interval"], d, "hello_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-interval"], "RouterOspfOspfInterface-HelloInterval"); ok {
			if err = d.Set("hello_interval", vv); err != nil {
				return fmt.Errorf("Error reading hello_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("hello_multiplier", flattenRouterOspfOspfInterfaceHelloMultiplier2edl(o["hello-multiplier"], d, "hello_multiplier")); err != nil {
		if vv, ok := fortiAPIPatch(o["hello-multiplier"], "RouterOspfOspfInterface-HelloMultiplier"); ok {
			if err = d.Set("hello_multiplier", vv); err != nil {
				return fmt.Errorf("Error reading hello_multiplier: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hello_multiplier: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterOspfOspfInterfaceInterface2edl(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "RouterOspfOspfInterface-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenRouterOspfOspfInterfaceIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "RouterOspfOspfInterface-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("keychain", flattenRouterOspfOspfInterfaceKeychain2edl(o["keychain"], d, "keychain")); err != nil {
		if vv, ok := fortiAPIPatch(o["keychain"], "RouterOspfOspfInterface-Keychain"); ok {
			if err = d.Set("keychain", vv); err != nil {
				return fmt.Errorf("Error reading keychain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keychain: %v", err)
		}
	}

	if err = d.Set("linkdown_fast_failover", flattenRouterOspfOspfInterfaceLinkdownFastFailover2edl(o["linkdown-fast-failover"], d, "linkdown_fast_failover")); err != nil {
		if vv, ok := fortiAPIPatch(o["linkdown-fast-failover"], "RouterOspfOspfInterface-LinkdownFastFailover"); ok {
			if err = d.Set("linkdown_fast_failover", vv); err != nil {
				return fmt.Errorf("Error reading linkdown_fast_failover: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading linkdown_fast_failover: %v", err)
		}
	}

	if err = d.Set("md5_keychain", flattenRouterOspfOspfInterfaceMd5Keychain2edl(o["md5-keychain"], d, "md5_keychain")); err != nil {
		if vv, ok := fortiAPIPatch(o["md5-keychain"], "RouterOspfOspfInterface-Md5Keychain"); ok {
			if err = d.Set("md5_keychain", vv); err != nil {
				return fmt.Errorf("Error reading md5_keychain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading md5_keychain: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("md5_keys", flattenRouterOspfOspfInterfaceMd5Keys2edl(o["md5-keys"], d, "md5_keys")); err != nil {
			if vv, ok := fortiAPIPatch(o["md5-keys"], "RouterOspfOspfInterface-Md5Keys"); ok {
				if err = d.Set("md5_keys", vv); err != nil {
					return fmt.Errorf("Error reading md5_keys: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading md5_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("md5_keys"); ok {
			if err = d.Set("md5_keys", flattenRouterOspfOspfInterfaceMd5Keys2edl(o["md5-keys"], d, "md5_keys")); err != nil {
				if vv, ok := fortiAPIPatch(o["md5-keys"], "RouterOspfOspfInterface-Md5Keys"); ok {
					if err = d.Set("md5_keys", vv); err != nil {
						return fmt.Errorf("Error reading md5_keys: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading md5_keys: %v", err)
				}
			}
		}
	}

	if err = d.Set("mtu", flattenRouterOspfOspfInterfaceMtu2edl(o["mtu"], d, "mtu")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu"], "RouterOspfOspfInterface-Mtu"); ok {
			if err = d.Set("mtu", vv); err != nil {
				return fmt.Errorf("Error reading mtu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("mtu_ignore", flattenRouterOspfOspfInterfaceMtuIgnore2edl(o["mtu-ignore"], d, "mtu_ignore")); err != nil {
		if vv, ok := fortiAPIPatch(o["mtu-ignore"], "RouterOspfOspfInterface-MtuIgnore"); ok {
			if err = d.Set("mtu_ignore", vv); err != nil {
				return fmt.Errorf("Error reading mtu_ignore: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mtu_ignore: %v", err)
		}
	}

	if err = d.Set("name", flattenRouterOspfOspfInterfaceName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "RouterOspfOspfInterface-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("network_type", flattenRouterOspfOspfInterfaceNetworkType2edl(o["network-type"], d, "network_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["network-type"], "RouterOspfOspfInterface-NetworkType"); ok {
			if err = d.Set("network_type", vv); err != nil {
				return fmt.Errorf("Error reading network_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("prefix_length", flattenRouterOspfOspfInterfacePrefixLength2edl(o["prefix-length"], d, "prefix_length")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefix-length"], "RouterOspfOspfInterface-PrefixLength"); ok {
			if err = d.Set("prefix_length", vv); err != nil {
				return fmt.Errorf("Error reading prefix_length: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefix_length: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterOspfOspfInterfacePriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "RouterOspfOspfInterface-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("resync_timeout", flattenRouterOspfOspfInterfaceResyncTimeout2edl(o["resync-timeout"], d, "resync_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["resync-timeout"], "RouterOspfOspfInterface-ResyncTimeout"); ok {
			if err = d.Set("resync_timeout", vv); err != nil {
				return fmt.Errorf("Error reading resync_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading resync_timeout: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterOspfOspfInterfaceRetransmitInterval2edl(o["retransmit-interval"], d, "retransmit_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["retransmit-interval"], "RouterOspfOspfInterface-RetransmitInterval"); ok {
			if err = d.Set("retransmit_interval", vv); err != nil {
				return fmt.Errorf("Error reading retransmit_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterOspfOspfInterfaceStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "RouterOspfOspfInterface-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterOspfOspfInterfaceTransmitDelay2edl(o["transmit-delay"], d, "transmit_delay")); err != nil {
		if vv, ok := fortiAPIPatch(o["transmit-delay"], "RouterOspfOspfInterface-TransmitDelay"); ok {
			if err = d.Set("transmit_delay", vv); err != nil {
				return fmt.Errorf("Error reading transmit_delay: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	return nil
}

func flattenRouterOspfOspfInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfOspfInterfaceAuthentication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceAuthenticationKey2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceBfd2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceComments2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceCost2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDatabaseFilterOut2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDeadInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloMultiplier2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceInterface2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceKeychain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceLinkdownFastFailover2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Keychain2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceMd5Keys2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfOspfInterfaceMd5KeysId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key-string"], _ = expandRouterOspfOspfInterfaceMd5KeysKeyString2edl(d, i["key_string"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandRouterOspfOspfInterfaceMd5KeysId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5KeysKeyString2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandRouterOspfOspfInterfaceMtu2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMtuIgnore2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceNetworkType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePrefixLength2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceResyncTimeout2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceRetransmitInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceTransmitDelay2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfOspfInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandRouterOspfOspfInterfaceAuthentication2edl(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("authentication_key"); ok || d.HasChange("authentication_key") {
		t, err := expandRouterOspfOspfInterfaceAuthenticationKey2edl(d, v, "authentication_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication-key"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok || d.HasChange("bfd") {
		t, err := expandRouterOspfOspfInterfaceBfd2edl(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandRouterOspfOspfInterfaceComments2edl(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("cost"); ok || d.HasChange("cost") {
		t, err := expandRouterOspfOspfInterfaceCost2edl(d, v, "cost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOk("database_filter_out"); ok || d.HasChange("database_filter_out") {
		t, err := expandRouterOspfOspfInterfaceDatabaseFilterOut2edl(d, v, "database_filter_out")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-filter-out"] = t
		}
	}

	if v, ok := d.GetOk("dead_interval"); ok || d.HasChange("dead_interval") {
		t, err := expandRouterOspfOspfInterfaceDeadInterval2edl(d, v, "dead_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_interval"); ok || d.HasChange("hello_interval") {
		t, err := expandRouterOspfOspfInterfaceHelloInterval2edl(d, v, "hello_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	}

	if v, ok := d.GetOk("hello_multiplier"); ok || d.HasChange("hello_multiplier") {
		t, err := expandRouterOspfOspfInterfaceHelloMultiplier2edl(d, v, "hello_multiplier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-multiplier"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandRouterOspfOspfInterfaceInterface2edl(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandRouterOspfOspfInterfaceIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("keychain"); ok || d.HasChange("keychain") {
		t, err := expandRouterOspfOspfInterfaceKeychain2edl(d, v, "keychain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keychain"] = t
		}
	}

	if v, ok := d.GetOk("linkdown_fast_failover"); ok || d.HasChange("linkdown_fast_failover") {
		t, err := expandRouterOspfOspfInterfaceLinkdownFastFailover2edl(d, v, "linkdown_fast_failover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkdown-fast-failover"] = t
		}
	}

	if v, ok := d.GetOk("md5_keychain"); ok || d.HasChange("md5_keychain") {
		t, err := expandRouterOspfOspfInterfaceMd5Keychain2edl(d, v, "md5_keychain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keychain"] = t
		}
	}

	if v, ok := d.GetOk("md5_keys"); ok || d.HasChange("md5_keys") {
		t, err := expandRouterOspfOspfInterfaceMd5Keys2edl(d, v, "md5_keys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keys"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok || d.HasChange("mtu") {
		t, err := expandRouterOspfOspfInterfaceMtu2edl(d, v, "mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("mtu_ignore"); ok || d.HasChange("mtu_ignore") {
		t, err := expandRouterOspfOspfInterfaceMtuIgnore2edl(d, v, "mtu_ignore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-ignore"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandRouterOspfOspfInterfaceName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok || d.HasChange("network_type") {
		t, err := expandRouterOspfOspfInterfaceNetworkType2edl(d, v, "network_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}

	if v, ok := d.GetOk("prefix_length"); ok || d.HasChange("prefix_length") {
		t, err := expandRouterOspfOspfInterfacePrefixLength2edl(d, v, "prefix_length")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-length"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandRouterOspfOspfInterfacePriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("resync_timeout"); ok || d.HasChange("resync_timeout") {
		t, err := expandRouterOspfOspfInterfaceResyncTimeout2edl(d, v, "resync_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resync-timeout"] = t
		}
	}

	if v, ok := d.GetOk("retransmit_interval"); ok || d.HasChange("retransmit_interval") {
		t, err := expandRouterOspfOspfInterfaceRetransmitInterval2edl(d, v, "retransmit_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandRouterOspfOspfInterfaceStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok || d.HasChange("transmit_delay") {
		t, err := expandRouterOspfOspfInterfaceTransmitDelay2edl(d, v, "transmit_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	return &obj, nil
}
