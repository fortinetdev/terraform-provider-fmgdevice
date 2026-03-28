// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure TACACS+ server entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserTacacs() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserTacacsCreate,
		Read:   resourceUserTacacsRead,
		Update: resourceUserTacacsUpdate,
		Delete: resourceUserTacacsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"authen_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"authen_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authorization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"secondary_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"secondary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tertiary_key": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"tertiary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vrf_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"secondary_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tertiary_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserTacacsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectUserTacacs(d)
	if err != nil {
		return fmt.Errorf("Error creating UserTacacs resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserTacacs(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserTacacs(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserTacacs resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserTacacs(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserTacacs resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserTacacsRead(d, m)
}

func resourceUserTacacsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	obj, err := getObjectUserTacacs(d)
	if err != nil {
		return fmt.Errorf("Error updating UserTacacs resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserTacacs(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserTacacs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserTacacsRead(d, m)
}

func resourceUserTacacsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

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

	wsParams["adom"] = adomv

	err = c.DeleteUserTacacs(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserTacacs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserTacacsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserTacacs(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserTacacs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserTacacs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserTacacs resource from API: %v", err)
	}
	return nil
}

func flattenUserTacacsAuthenType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenUserTacacsDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authen_type"
		if _, ok := i["authen-type"]; ok {
			v := flattenUserTacacsDynamicMappingAuthenType(i["authen-type"], d, pre_append)
			tmp["authen_type"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-AuthenType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization"
		if _, ok := i["authorization"]; ok {
			v := flattenUserTacacsDynamicMappingAuthorization(i["authorization"], d, pre_append)
			tmp["authorization"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-Authorization")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenUserTacacsDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenUserTacacsDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserTacacsDynamicMappingPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := i["secondary-server"]; ok {
			v := flattenUserTacacsDynamicMappingSecondaryServer(i["secondary-server"], d, pre_append)
			tmp["secondary_server"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-SecondaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenUserTacacsDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenUserTacacsDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status_ttl"
		if _, ok := i["status-ttl"]; ok {
			v := flattenUserTacacsDynamicMappingStatusTtl(i["status-ttl"], d, pre_append)
			tmp["status_ttl"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-StatusTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := i["tertiary-server"]; ok {
			v := flattenUserTacacsDynamicMappingTertiaryServer(i["tertiary-server"], d, pre_append)
			tmp["tertiary_server"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-TertiaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenUserTacacsDynamicMappingVrfSelect(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "UserTacacs-DynamicMapping-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserTacacsDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenUserTacacsDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserTacacsDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenUserTacacsDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "UserTacacsDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserTacacsDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingAuthenType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserTacacsDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingStatusTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsDynamicMappingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserTacacsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsStatusTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserTacacsVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserTacacs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("authen_type", flattenUserTacacsAuthenType(o["authen-type"], d, "authen_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["authen-type"], "UserTacacs-AuthenType"); ok {
			if err = d.Set("authen_type", vv); err != nil {
				return fmt.Errorf("Error reading authen_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authen_type: %v", err)
		}
	}

	if err = d.Set("authorization", flattenUserTacacsAuthorization(o["authorization"], d, "authorization")); err != nil {
		if vv, ok := fortiAPIPatch(o["authorization"], "UserTacacs-Authorization"); ok {
			if err = d.Set("authorization", vv); err != nil {
				return fmt.Errorf("Error reading authorization: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authorization: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenUserTacacsDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserTacacs-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenUserTacacsDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserTacacs-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenUserTacacsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "UserTacacs-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserTacacsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "UserTacacs-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenUserTacacsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserTacacs-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenUserTacacsPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "UserTacacs-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserTacacsSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "UserTacacs-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("server", flattenUserTacacsServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "UserTacacs-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserTacacsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "UserTacacs-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenUserTacacsStatusTtl(o["status-ttl"], d, "status_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-ttl"], "UserTacacs-StatusTtl"); ok {
			if err = d.Set("status_ttl", vv); err != nil {
				return fmt.Errorf("Error reading status_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserTacacsTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "UserTacacs-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserTacacsVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "UserTacacs-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenUserTacacsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserTacacsAuthenType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserTacacsDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authen_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authen-type"], _ = expandUserTacacsDynamicMappingAuthenType(d, i["authen_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authorization"], _ = expandUserTacacsDynamicMappingAuthorization(d, i["authorization"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandUserTacacsDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandUserTacacsDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["key"], _ = expandUserTacacsDynamicMappingKey(d, i["key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserTacacsDynamicMappingPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-key"], _ = expandUserTacacsDynamicMappingSecondaryKey(d, i["secondary_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-server"], _ = expandUserTacacsDynamicMappingSecondaryServer(d, i["secondary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandUserTacacsDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandUserTacacsDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status-ttl"], _ = expandUserTacacsDynamicMappingStatusTtl(d, i["status_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tertiary-key"], _ = expandUserTacacsDynamicMappingTertiaryKey(d, i["tertiary_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tertiary-server"], _ = expandUserTacacsDynamicMappingTertiaryServer(d, i["tertiary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandUserTacacsDynamicMappingVrfSelect(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserTacacsDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandUserTacacsDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandUserTacacsDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserTacacsDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingAuthenType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsDynamicMappingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingSecondaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsDynamicMappingSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingStatusTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingTertiaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsDynamicMappingTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsDynamicMappingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsSecondaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsStatusTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsTertiaryKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserTacacsTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserTacacsVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserTacacs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("authen_type"); ok || d.HasChange("authen_type") {
		t, err := expandUserTacacsAuthenType(d, v, "authen_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authen-type"] = t
		}
	}

	if v, ok := d.GetOk("authorization"); ok || d.HasChange("authorization") {
		t, err := expandUserTacacsAuthorization(d, v, "authorization")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandUserTacacsDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandUserTacacsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandUserTacacsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {
		t, err := expandUserTacacsKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserTacacsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandUserTacacsPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("secondary_key"); ok || d.HasChange("secondary_key") {
		t, err := expandUserTacacsSecondaryKey(d, v, "secondary_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-key"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandUserTacacsSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandUserTacacsServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandUserTacacsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("status_ttl"); ok || d.HasChange("status_ttl") {
		t, err := expandUserTacacsStatusTtl(d, v, "status_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_key"); ok || d.HasChange("tertiary_key") {
		t, err := expandUserTacacsTertiaryKey(d, v, "tertiary_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-key"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandUserTacacsTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandUserTacacsVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
