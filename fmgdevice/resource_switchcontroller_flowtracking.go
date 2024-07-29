// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch flow tracking and export via ipfix/netflow.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerFlowTracking() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerFlowTrackingUpdate,
		Read:   resourceSwitchControllerFlowTrackingRead,
		Update: resourceSwitchControllerFlowTrackingUpdate,
		Delete: resourceSwitchControllerFlowTrackingDelete,

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
			"aggregates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"collectors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"transport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_export_pkt_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sample_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sample_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"template_export_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_general": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_icmp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_tcp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_tcp_fin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_tcp_rst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_udp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"transport": &schema.Schema{
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

func resourceSwitchControllerFlowTrackingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerFlowTracking(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFlowTracking resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerFlowTracking(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFlowTracking resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SwitchControllerFlowTracking")

	return resourceSwitchControllerFlowTrackingRead(d, m)
}

func resourceSwitchControllerFlowTrackingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerFlowTracking(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerFlowTracking resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerFlowTrackingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerFlowTracking(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFlowTracking resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerFlowTracking(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFlowTracking resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerFlowTrackingAggregates(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSwitchControllerFlowTrackingAggregatesId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SwitchControllerFlowTracking-Aggregates-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerFlowTrackingAggregatesIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerFlowTracking-Aggregates-Ip")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerFlowTrackingAggregatesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingAggregatesIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSwitchControllerFlowTrackingCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectors(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSwitchControllerFlowTrackingCollectorsIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SwitchControllerFlowTracking-Collectors-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSwitchControllerFlowTrackingCollectorsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SwitchControllerFlowTracking-Collectors-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSwitchControllerFlowTrackingCollectorsPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SwitchControllerFlowTracking-Collectors-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport"
		if _, ok := i["transport"]; ok {
			v := flattenSwitchControllerFlowTrackingCollectorsTransport(i["transport"], d, pre_append)
			tmp["transport"] = fortiAPISubPartPatch(v, "SwitchControllerFlowTracking-Collectors-Transport")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSwitchControllerFlowTrackingCollectorsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectorsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectorsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectorsTransport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingMaxExportPktSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingSampleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTemplateExportPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutGeneral(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutIcmp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutTcp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutTcpFin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutTcpRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutUdp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTransport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerFlowTracking(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("aggregates", flattenSwitchControllerFlowTrackingAggregates(o["aggregates"], d, "aggregates")); err != nil {
			if vv, ok := fortiAPIPatch(o["aggregates"], "SwitchControllerFlowTracking-Aggregates"); ok {
				if err = d.Set("aggregates", vv); err != nil {
					return fmt.Errorf("Error reading aggregates: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading aggregates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregates"); ok {
			if err = d.Set("aggregates", flattenSwitchControllerFlowTrackingAggregates(o["aggregates"], d, "aggregates")); err != nil {
				if vv, ok := fortiAPIPatch(o["aggregates"], "SwitchControllerFlowTracking-Aggregates"); ok {
					if err = d.Set("aggregates", vv); err != nil {
						return fmt.Errorf("Error reading aggregates: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading aggregates: %v", err)
				}
			}
		}
	}

	if err = d.Set("collector_ip", flattenSwitchControllerFlowTrackingCollectorIp(o["collector-ip"], d, "collector_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-ip"], "SwitchControllerFlowTracking-CollectorIp"); ok {
			if err = d.Set("collector_ip", vv); err != nil {
				return fmt.Errorf("Error reading collector_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSwitchControllerFlowTrackingCollectorPort(o["collector-port"], d, "collector_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["collector-port"], "SwitchControllerFlowTracking-CollectorPort"); ok {
			if err = d.Set("collector_port", vv); err != nil {
				return fmt.Errorf("Error reading collector_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("collectors", flattenSwitchControllerFlowTrackingCollectors(o["collectors"], d, "collectors")); err != nil {
			if vv, ok := fortiAPIPatch(o["collectors"], "SwitchControllerFlowTracking-Collectors"); ok {
				if err = d.Set("collectors", vv); err != nil {
					return fmt.Errorf("Error reading collectors: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading collectors: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("collectors"); ok {
			if err = d.Set("collectors", flattenSwitchControllerFlowTrackingCollectors(o["collectors"], d, "collectors")); err != nil {
				if vv, ok := fortiAPIPatch(o["collectors"], "SwitchControllerFlowTracking-Collectors"); ok {
					if err = d.Set("collectors", vv); err != nil {
						return fmt.Errorf("Error reading collectors: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading collectors: %v", err)
				}
			}
		}
	}

	if err = d.Set("format", flattenSwitchControllerFlowTrackingFormat(o["format"], d, "format")); err != nil {
		if vv, ok := fortiAPIPatch(o["format"], "SwitchControllerFlowTracking-Format"); ok {
			if err = d.Set("format", vv); err != nil {
				return fmt.Errorf("Error reading format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("level", flattenSwitchControllerFlowTrackingLevel(o["level"], d, "level")); err != nil {
		if vv, ok := fortiAPIPatch(o["level"], "SwitchControllerFlowTracking-Level"); ok {
			if err = d.Set("level", vv); err != nil {
				return fmt.Errorf("Error reading level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("max_export_pkt_size", flattenSwitchControllerFlowTrackingMaxExportPktSize(o["max-export-pkt-size"], d, "max_export_pkt_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-export-pkt-size"], "SwitchControllerFlowTracking-MaxExportPktSize"); ok {
			if err = d.Set("max_export_pkt_size", vv); err != nil {
				return fmt.Errorf("Error reading max_export_pkt_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_export_pkt_size: %v", err)
		}
	}

	if err = d.Set("sample_mode", flattenSwitchControllerFlowTrackingSampleMode(o["sample-mode"], d, "sample_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["sample-mode"], "SwitchControllerFlowTracking-SampleMode"); ok {
			if err = d.Set("sample_mode", vv); err != nil {
				return fmt.Errorf("Error reading sample_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sample_mode: %v", err)
		}
	}

	if err = d.Set("sample_rate", flattenSwitchControllerFlowTrackingSampleRate(o["sample-rate"], d, "sample_rate")); err != nil {
		if vv, ok := fortiAPIPatch(o["sample-rate"], "SwitchControllerFlowTracking-SampleRate"); ok {
			if err = d.Set("sample_rate", vv); err != nil {
				return fmt.Errorf("Error reading sample_rate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sample_rate: %v", err)
		}
	}

	if err = d.Set("template_export_period", flattenSwitchControllerFlowTrackingTemplateExportPeriod(o["template-export-period"], d, "template_export_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["template-export-period"], "SwitchControllerFlowTracking-TemplateExportPeriod"); ok {
			if err = d.Set("template_export_period", vv); err != nil {
				return fmt.Errorf("Error reading template_export_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading template_export_period: %v", err)
		}
	}

	if err = d.Set("timeout_general", flattenSwitchControllerFlowTrackingTimeoutGeneral(o["timeout-general"], d, "timeout_general")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-general"], "SwitchControllerFlowTracking-TimeoutGeneral"); ok {
			if err = d.Set("timeout_general", vv); err != nil {
				return fmt.Errorf("Error reading timeout_general: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_general: %v", err)
		}
	}

	if err = d.Set("timeout_icmp", flattenSwitchControllerFlowTrackingTimeoutIcmp(o["timeout-icmp"], d, "timeout_icmp")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-icmp"], "SwitchControllerFlowTracking-TimeoutIcmp"); ok {
			if err = d.Set("timeout_icmp", vv); err != nil {
				return fmt.Errorf("Error reading timeout_icmp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_icmp: %v", err)
		}
	}

	if err = d.Set("timeout_max", flattenSwitchControllerFlowTrackingTimeoutMax(o["timeout-max"], d, "timeout_max")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-max"], "SwitchControllerFlowTracking-TimeoutMax"); ok {
			if err = d.Set("timeout_max", vv); err != nil {
				return fmt.Errorf("Error reading timeout_max: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_max: %v", err)
		}
	}

	if err = d.Set("timeout_tcp", flattenSwitchControllerFlowTrackingTimeoutTcp(o["timeout-tcp"], d, "timeout_tcp")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-tcp"], "SwitchControllerFlowTracking-TimeoutTcp"); ok {
			if err = d.Set("timeout_tcp", vv); err != nil {
				return fmt.Errorf("Error reading timeout_tcp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_tcp: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_fin", flattenSwitchControllerFlowTrackingTimeoutTcpFin(o["timeout-tcp-fin"], d, "timeout_tcp_fin")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-tcp-fin"], "SwitchControllerFlowTracking-TimeoutTcpFin"); ok {
			if err = d.Set("timeout_tcp_fin", vv); err != nil {
				return fmt.Errorf("Error reading timeout_tcp_fin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_tcp_fin: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_rst", flattenSwitchControllerFlowTrackingTimeoutTcpRst(o["timeout-tcp-rst"], d, "timeout_tcp_rst")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-tcp-rst"], "SwitchControllerFlowTracking-TimeoutTcpRst"); ok {
			if err = d.Set("timeout_tcp_rst", vv); err != nil {
				return fmt.Errorf("Error reading timeout_tcp_rst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_tcp_rst: %v", err)
		}
	}

	if err = d.Set("timeout_udp", flattenSwitchControllerFlowTrackingTimeoutUdp(o["timeout-udp"], d, "timeout_udp")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout-udp"], "SwitchControllerFlowTracking-TimeoutUdp"); ok {
			if err = d.Set("timeout_udp", vv); err != nil {
				return fmt.Errorf("Error reading timeout_udp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout_udp: %v", err)
		}
	}

	if err = d.Set("transport", flattenSwitchControllerFlowTrackingTransport(o["transport"], d, "transport")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport"], "SwitchControllerFlowTracking-Transport"); ok {
			if err = d.Set("transport", vv); err != nil {
				return fmt.Errorf("Error reading transport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerFlowTrackingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerFlowTrackingAggregates(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerFlowTrackingAggregatesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerFlowTrackingAggregatesIp(d, i["ip"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerFlowTrackingAggregatesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingAggregatesIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSwitchControllerFlowTrackingCollectorIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectors(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSwitchControllerFlowTrackingCollectorsIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSwitchControllerFlowTrackingCollectorsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSwitchControllerFlowTrackingCollectorsPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transport"], _ = expandSwitchControllerFlowTrackingCollectorsTransport(d, i["transport"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSwitchControllerFlowTrackingCollectorsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectorsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectorsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectorsTransport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingMaxExportPktSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingSampleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTemplateExportPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutGeneral(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutIcmp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutTcp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutTcpFin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutTcpRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutUdp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTransport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerFlowTracking(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aggregates"); ok || d.HasChange("aggregates") {
		t, err := expandSwitchControllerFlowTrackingAggregates(d, v, "aggregates")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregates"] = t
		}
	}

	if v, ok := d.GetOk("collector_ip"); ok || d.HasChange("collector_ip") {
		t, err := expandSwitchControllerFlowTrackingCollectorIp(d, v, "collector_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOk("collector_port"); ok || d.HasChange("collector_port") {
		t, err := expandSwitchControllerFlowTrackingCollectorPort(d, v, "collector_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	if v, ok := d.GetOk("collectors"); ok || d.HasChange("collectors") {
		t, err := expandSwitchControllerFlowTrackingCollectors(d, v, "collectors")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collectors"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok || d.HasChange("format") {
		t, err := expandSwitchControllerFlowTrackingFormat(d, v, "format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("level"); ok || d.HasChange("level") {
		t, err := expandSwitchControllerFlowTrackingLevel(d, v, "level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["level"] = t
		}
	}

	if v, ok := d.GetOk("max_export_pkt_size"); ok || d.HasChange("max_export_pkt_size") {
		t, err := expandSwitchControllerFlowTrackingMaxExportPktSize(d, v, "max_export_pkt_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-export-pkt-size"] = t
		}
	}

	if v, ok := d.GetOk("sample_mode"); ok || d.HasChange("sample_mode") {
		t, err := expandSwitchControllerFlowTrackingSampleMode(d, v, "sample_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-mode"] = t
		}
	}

	if v, ok := d.GetOk("sample_rate"); ok || d.HasChange("sample_rate") {
		t, err := expandSwitchControllerFlowTrackingSampleRate(d, v, "sample_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("template_export_period"); ok || d.HasChange("template_export_period") {
		t, err := expandSwitchControllerFlowTrackingTemplateExportPeriod(d, v, "template_export_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-export-period"] = t
		}
	}

	if v, ok := d.GetOk("timeout_general"); ok || d.HasChange("timeout_general") {
		t, err := expandSwitchControllerFlowTrackingTimeoutGeneral(d, v, "timeout_general")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-general"] = t
		}
	}

	if v, ok := d.GetOk("timeout_icmp"); ok || d.HasChange("timeout_icmp") {
		t, err := expandSwitchControllerFlowTrackingTimeoutIcmp(d, v, "timeout_icmp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-icmp"] = t
		}
	}

	if v, ok := d.GetOk("timeout_max"); ok || d.HasChange("timeout_max") {
		t, err := expandSwitchControllerFlowTrackingTimeoutMax(d, v, "timeout_max")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-max"] = t
		}
	}

	if v, ok := d.GetOk("timeout_tcp"); ok || d.HasChange("timeout_tcp") {
		t, err := expandSwitchControllerFlowTrackingTimeoutTcp(d, v, "timeout_tcp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-tcp"] = t
		}
	}

	if v, ok := d.GetOk("timeout_tcp_fin"); ok || d.HasChange("timeout_tcp_fin") {
		t, err := expandSwitchControllerFlowTrackingTimeoutTcpFin(d, v, "timeout_tcp_fin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-tcp-fin"] = t
		}
	}

	if v, ok := d.GetOk("timeout_tcp_rst"); ok || d.HasChange("timeout_tcp_rst") {
		t, err := expandSwitchControllerFlowTrackingTimeoutTcpRst(d, v, "timeout_tcp_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-tcp-rst"] = t
		}
	}

	if v, ok := d.GetOk("timeout_udp"); ok || d.HasChange("timeout_udp") {
		t, err := expandSwitchControllerFlowTrackingTimeoutUdp(d, v, "timeout_udp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-udp"] = t
		}
	}

	if v, ok := d.GetOk("transport"); ok || d.HasChange("transport") {
		t, err := expandSwitchControllerFlowTrackingTransport(d, v, "transport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport"] = t
		}
	}

	return &obj, nil
}
