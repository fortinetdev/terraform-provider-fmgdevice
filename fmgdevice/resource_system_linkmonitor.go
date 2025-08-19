// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Link Health Monitor.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLinkMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLinkMonitorCreate,
		Read:   resourceSystemLinkMonitorRead,
		Update: resourceSystemLinkMonitorUpdate,
		Delete: resourceSystemLinkMonitorDelete,

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
			"addr_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"failtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"packet_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
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
			"probe_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"probe_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"recoverytime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"route": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
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
			"update_cascade_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_policy_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_static_route": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceSystemLinkMonitorCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemLinkMonitor(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLinkMonitor resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLinkMonitor(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemLinkMonitor resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemLinkMonitorRead(d, m)
}

func resourceSystemLinkMonitorUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemLinkMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLinkMonitor resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLinkMonitor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemLinkMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemLinkMonitorRead(d, m)
}

func resourceSystemLinkMonitorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemLinkMonitor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLinkMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLinkMonitorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemLinkMonitor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLinkMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLinkMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLinkMonitor resource from API: %v", err)
	}
	return nil
}

func flattenSystemLinkMonitorAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLinkMonitorDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorFailWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorGatewayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorGatewayIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorProbeCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLinkMonitorRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLinkMonitorSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLinkMonitorServerConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenSystemLinkMonitorServerListDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "SystemLinkMonitor-ServerList-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemLinkMonitorServerListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemLinkMonitor-ServerList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSystemLinkMonitorServerListPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SystemLinkMonitor-ServerList-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenSystemLinkMonitorServerListProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "SystemLinkMonitor-ServerList-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSystemLinkMonitorServerListWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SystemLinkMonitor-ServerList-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemLinkMonitorServerListDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServerListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServerListPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServerListProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLinkMonitorServerListWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServiceDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemLinkMonitorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorUpdatePolicyRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLinkMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("addr_mode", flattenSystemLinkMonitorAddrMode(o["addr-mode"], d, "addr_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["addr-mode"], "SystemLinkMonitor-AddrMode"); ok {
			if err = d.Set("addr_mode", vv); err != nil {
				return fmt.Errorf("Error reading addr_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("class_id", flattenSystemLinkMonitorClassId(o["class-id"], d, "class_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["class-id"], "SystemLinkMonitor-ClassId"); ok {
			if err = d.Set("class_id", vv); err != nil {
				return fmt.Errorf("Error reading class_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenSystemLinkMonitorDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "SystemLinkMonitor-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("fail_weight", flattenSystemLinkMonitorFailWeight(o["fail-weight"], d, "fail_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-weight"], "SystemLinkMonitor-FailWeight"); ok {
			if err = d.Set("fail_weight", vv); err != nil {
				return fmt.Errorf("Error reading fail_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_weight: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemLinkMonitorFailtime(o["failtime"], d, "failtime")); err != nil {
		if vv, ok := fortiAPIPatch(o["failtime"], "SystemLinkMonitor-Failtime"); ok {
			if err = d.Set("failtime", vv); err != nil {
				return fmt.Errorf("Error reading failtime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("gateway_ip", flattenSystemLinkMonitorGatewayIp(o["gateway-ip"], d, "gateway_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway-ip"], "SystemLinkMonitor-GatewayIp"); ok {
			if err = d.Set("gateway_ip", vv); err != nil {
				return fmt.Errorf("Error reading gateway_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway_ip: %v", err)
		}
	}

	if err = d.Set("gateway_ip6", flattenSystemLinkMonitorGatewayIp6(o["gateway-ip6"], d, "gateway_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["gateway-ip6"], "SystemLinkMonitor-GatewayIp6"); ok {
			if err = d.Set("gateway_ip6", vv); err != nil {
				return fmt.Errorf("Error reading gateway_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gateway_ip6: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenSystemLinkMonitorHttpAgent(o["http-agent"], d, "http_agent")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-agent"], "SystemLinkMonitor-HttpAgent"); ok {
			if err = d.Set("http_agent", vv); err != nil {
				return fmt.Errorf("Error reading http_agent: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_get", flattenSystemLinkMonitorHttpGet(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "SystemLinkMonitor-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenSystemLinkMonitorHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "SystemLinkMonitor-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemLinkMonitorInterval(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "SystemLinkMonitor-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemLinkMonitorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemLinkMonitor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenSystemLinkMonitorPacketSize(o["packet-size"], d, "packet_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["packet-size"], "SystemLinkMonitor-PacketSize"); ok {
			if err = d.Set("packet_size", vv); err != nil {
				return fmt.Errorf("Error reading packet_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemLinkMonitorPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "SystemLinkMonitor-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("probe_count", flattenSystemLinkMonitorProbeCount(o["probe-count"], d, "probe_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-count"], "SystemLinkMonitor-ProbeCount"); ok {
			if err = d.Set("probe_count", vv); err != nil {
				return fmt.Errorf("Error reading probe_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_count: %v", err)
		}
	}

	if err = d.Set("probe_timeout", flattenSystemLinkMonitorProbeTimeout(o["probe-timeout"], d, "probe_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["probe-timeout"], "SystemLinkMonitor-ProbeTimeout"); ok {
			if err = d.Set("probe_timeout", vv); err != nil {
				return fmt.Errorf("Error reading probe_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading probe_timeout: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemLinkMonitorProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemLinkMonitor-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenSystemLinkMonitorRecoverytime(o["recoverytime"], d, "recoverytime")); err != nil {
		if vv, ok := fortiAPIPatch(o["recoverytime"], "SystemLinkMonitor-Recoverytime"); ok {
			if err = d.Set("recoverytime", vv); err != nil {
				return fmt.Errorf("Error reading recoverytime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("route", flattenSystemLinkMonitorRoute(o["route"], d, "route")); err != nil {
		if vv, ok := fortiAPIPatch(o["route"], "SystemLinkMonitor-Route"); ok {
			if err = d.Set("route", vv); err != nil {
				return fmt.Errorf("Error reading route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading route: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemLinkMonitorSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["security-mode"], "SystemLinkMonitor-SecurityMode"); ok {
			if err = d.Set("security_mode", vv); err != nil {
				return fmt.Errorf("Error reading security_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemLinkMonitorServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "SystemLinkMonitor-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_config", flattenSystemLinkMonitorServerConfig(o["server-config"], d, "server_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-config"], "SystemLinkMonitor-ServerConfig"); ok {
			if err = d.Set("server_config", vv); err != nil {
				return fmt.Errorf("Error reading server_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_config: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenSystemLinkMonitorServerList(o["server-list"], d, "server_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-list"], "SystemLinkMonitor-ServerList"); ok {
				if err = d.Set("server_list", vv); err != nil {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenSystemLinkMonitorServerList(o["server-list"], d, "server_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-list"], "SystemLinkMonitor-ServerList"); ok {
					if err = d.Set("server_list", vv); err != nil {
						return fmt.Errorf("Error reading server_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenSystemLinkMonitorServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemLinkMonitor-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("service_detection", flattenSystemLinkMonitorServiceDetection(o["service-detection"], d, "service_detection")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-detection"], "SystemLinkMonitor-ServiceDetection"); ok {
			if err = d.Set("service_detection", vv); err != nil {
				return fmt.Errorf("Error reading service_detection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_detection: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemLinkMonitorSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemLinkMonitor-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemLinkMonitorSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip6"], "SystemLinkMonitor-SourceIp6"); ok {
			if err = d.Set("source_ip6", vv); err != nil {
				return fmt.Errorf("Error reading source_ip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenSystemLinkMonitorSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "SystemLinkMonitor-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLinkMonitorStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemLinkMonitor-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenSystemLinkMonitorUpdateCascadeInterface(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-cascade-interface"], "SystemLinkMonitor-UpdateCascadeInterface"); ok {
			if err = d.Set("update_cascade_interface", vv); err != nil {
				return fmt.Errorf("Error reading update_cascade_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_policy_route", flattenSystemLinkMonitorUpdatePolicyRoute(o["update-policy-route"], d, "update_policy_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-policy-route"], "SystemLinkMonitor-UpdatePolicyRoute"); ok {
			if err = d.Set("update_policy_route", vv); err != nil {
				return fmt.Errorf("Error reading update_policy_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_policy_route: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenSystemLinkMonitorUpdateStaticRoute(o["update-static-route"], d, "update_static_route")); err != nil {
		if vv, ok := fortiAPIPatch(o["update-static-route"], "SystemLinkMonitor-UpdateStaticRoute"); ok {
			if err = d.Set("update_static_route", vv); err != nil {
				return fmt.Errorf("Error reading update_static_route: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	return nil
}

func flattenSystemLinkMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLinkMonitorAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorFailWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorGatewayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorGatewayIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorProbeCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorProbeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorServerConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandSystemLinkMonitorServerListDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemLinkMonitorServerListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSystemLinkMonitorServerListPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandSystemLinkMonitorServerListProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSystemLinkMonitorServerListWeight(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemLinkMonitorServerListDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServerListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServerListPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServerListProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorServerListWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServiceDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemLinkMonitorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorUpdatePolicyRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLinkMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("addr_mode"); ok || d.HasChange("addr_mode") {
		t, err := expandSystemLinkMonitorAddrMode(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("class_id"); ok || d.HasChange("class_id") {
		t, err := expandSystemLinkMonitorClassId(d, v, "class_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandSystemLinkMonitorDiffservcode(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("fail_weight"); ok || d.HasChange("fail_weight") {
		t, err := expandSystemLinkMonitorFailWeight(d, v, "fail_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-weight"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok || d.HasChange("failtime") {
		t, err := expandSystemLinkMonitorFailtime(d, v, "failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("gateway_ip"); ok || d.HasChange("gateway_ip") {
		t, err := expandSystemLinkMonitorGatewayIp(d, v, "gateway_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway-ip"] = t
		}
	}

	if v, ok := d.GetOk("gateway_ip6"); ok || d.HasChange("gateway_ip6") {
		t, err := expandSystemLinkMonitorGatewayIp6(d, v, "gateway_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway-ip6"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok || d.HasChange("http_agent") {
		t, err := expandSystemLinkMonitorHttpAgent(d, v, "http_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandSystemLinkMonitorHttpGet(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandSystemLinkMonitorHttpMatch(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandSystemLinkMonitorInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemLinkMonitorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok || d.HasChange("packet_size") {
		t, err := expandSystemLinkMonitorPacketSize(d, v, "packet_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemLinkMonitorPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandSystemLinkMonitorPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("probe_count"); ok || d.HasChange("probe_count") {
		t, err := expandSystemLinkMonitorProbeCount(d, v, "probe_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-count"] = t
		}
	}

	if v, ok := d.GetOk("probe_timeout"); ok || d.HasChange("probe_timeout") {
		t, err := expandSystemLinkMonitorProbeTimeout(d, v, "probe_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-timeout"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemLinkMonitorProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok || d.HasChange("recoverytime") {
		t, err := expandSystemLinkMonitorRecoverytime(d, v, "recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("route"); ok || d.HasChange("route") {
		t, err := expandSystemLinkMonitorRoute(d, v, "route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok || d.HasChange("security_mode") {
		t, err := expandSystemLinkMonitorSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandSystemLinkMonitorServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_config"); ok || d.HasChange("server_config") {
		t, err := expandSystemLinkMonitorServerConfig(d, v, "server_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-config"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandSystemLinkMonitorServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemLinkMonitorServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("service_detection"); ok || d.HasChange("service_detection") {
		t, err := expandSystemLinkMonitorServiceDetection(d, v, "service_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-detection"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemLinkMonitorSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok || d.HasChange("source_ip6") {
		t, err := expandSystemLinkMonitorSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandSystemLinkMonitorSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemLinkMonitorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok || d.HasChange("update_cascade_interface") {
		t, err := expandSystemLinkMonitorUpdateCascadeInterface(d, v, "update_cascade_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_policy_route"); ok || d.HasChange("update_policy_route") {
		t, err := expandSystemLinkMonitorUpdatePolicyRoute(d, v, "update_policy_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-policy-route"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok || d.HasChange("update_static_route") {
		t, err := expandSystemLinkMonitorUpdateStaticRoute(d, v, "update_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	return &obj, nil
}
