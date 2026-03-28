// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure queue assignment on NP7LITE.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNpuNpQueues() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNpuNpQueuesUpdate,
		Read:   resourceSystemNpuNpQueuesRead,
		Update: resourceSystemNpuNpQueuesUpdate,
		Delete: resourceSystemNpuNpQueuesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"custom_etype_lookup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ethernet_type": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ip_protocol": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ip_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sport": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"profile": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cos0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cos7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp0": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp10": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp11": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp12": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp13": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp14": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp15": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp16": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp17": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp18": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp19": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp20": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp21": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp22": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp23": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp24": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp25": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp26": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp27": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp28": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp29": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp30": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp31": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp32": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp33": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp34": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp35": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp36": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp37": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp38": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp39": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp4": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp40": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp41": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp42": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp43": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp44": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp45": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp46": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp47": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp48": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp49": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp5": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp50": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp51": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp52": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp53": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp54": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp55": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp56": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp57": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp58": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp59": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp60": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp61": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp62": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp63": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp7": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp8": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp9": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"scheduler": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceSystemNpuNpQueuesUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectSystemNpuNpQueues(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuNpQueues resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNpuNpQueues(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpuNpQueues resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemNpuNpQueues")

	return resourceSystemNpuNpQueuesRead(d, m)
}

func resourceSystemNpuNpQueuesDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteSystemNpuNpQueues(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNpuNpQueues resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuNpQueuesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadSystemNpuNpQueues(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNpuNpQueues resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNpuNpQueues(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpuNpQueues resource from API: %v", err)
	}
	return nil
}

func flattenSystemNpuNpQueuesCustomEtypeLookup2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetType2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemNpuNpQueuesEthernetTypeName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-EthernetType-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenSystemNpuNpQueuesEthernetTypeQueue2edl(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-EthernetType-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemNpuNpQueuesEthernetTypeType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-EthernetType-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSystemNpuNpQueuesEthernetTypeWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-EthernetType-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNpuNpQueuesEthernetTypeName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetTypeQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetTypeType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesEthernetTypeWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpProtocol2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemNpuNpQueuesIpProtocolName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpProtocol-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenSystemNpuNpQueuesIpProtocolProtocol2edl(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpProtocol-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenSystemNpuNpQueuesIpProtocolQueue2edl(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpProtocol-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSystemNpuNpQueuesIpProtocolWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpProtocol-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNpuNpQueuesIpProtocolName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpProtocolProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpProtocolQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpProtocolWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpService2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := i["dport"]; ok {
			v := flattenSystemNpuNpQueuesIpServiceDport2edl(i["dport"], d, pre_append)
			tmp["dport"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpService-Dport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemNpuNpQueuesIpServiceName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenSystemNpuNpQueuesIpServiceProtocol2edl(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpService-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := i["queue"]; ok {
			v := flattenSystemNpuNpQueuesIpServiceQueue2edl(i["queue"], d, pre_append)
			tmp["queue"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpService-Queue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := i["sport"]; ok {
			v := flattenSystemNpuNpQueuesIpServiceSport2edl(i["sport"], d, pre_append)
			tmp["sport"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpService-Sport")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSystemNpuNpQueuesIpServiceWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-IpService-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNpuNpQueuesIpServiceDport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpServiceName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpServiceProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpServiceQueue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpServiceSport2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesIpServiceWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfile2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := i["cos0"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos02edl(i["cos0"], d, pre_append)
			tmp["cos0"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := i["cos1"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos12edl(i["cos1"], d, pre_append)
			tmp["cos1"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := i["cos2"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos22edl(i["cos2"], d, pre_append)
			tmp["cos2"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := i["cos3"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos32edl(i["cos3"], d, pre_append)
			tmp["cos3"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := i["cos4"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos42edl(i["cos4"], d, pre_append)
			tmp["cos4"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := i["cos5"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos52edl(i["cos5"], d, pre_append)
			tmp["cos5"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := i["cos6"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos62edl(i["cos6"], d, pre_append)
			tmp["cos6"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := i["cos7"]; ok {
			v := flattenSystemNpuNpQueuesProfileCos72edl(i["cos7"], d, pre_append)
			tmp["cos7"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Cos7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := i["dscp0"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp02edl(i["dscp0"], d, pre_append)
			tmp["dscp0"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp0")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := i["dscp1"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp12edl(i["dscp1"], d, pre_append)
			tmp["dscp1"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp1")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := i["dscp10"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp102edl(i["dscp10"], d, pre_append)
			tmp["dscp10"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp10")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := i["dscp11"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp112edl(i["dscp11"], d, pre_append)
			tmp["dscp11"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp11")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := i["dscp12"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp122edl(i["dscp12"], d, pre_append)
			tmp["dscp12"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp12")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := i["dscp13"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp132edl(i["dscp13"], d, pre_append)
			tmp["dscp13"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp13")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := i["dscp14"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp142edl(i["dscp14"], d, pre_append)
			tmp["dscp14"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp14")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := i["dscp15"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp152edl(i["dscp15"], d, pre_append)
			tmp["dscp15"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp15")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := i["dscp16"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp162edl(i["dscp16"], d, pre_append)
			tmp["dscp16"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp16")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := i["dscp17"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp172edl(i["dscp17"], d, pre_append)
			tmp["dscp17"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp17")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := i["dscp18"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp182edl(i["dscp18"], d, pre_append)
			tmp["dscp18"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp18")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := i["dscp19"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp192edl(i["dscp19"], d, pre_append)
			tmp["dscp19"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp19")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := i["dscp2"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp22edl(i["dscp2"], d, pre_append)
			tmp["dscp2"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp2")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := i["dscp20"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp202edl(i["dscp20"], d, pre_append)
			tmp["dscp20"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp20")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := i["dscp21"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp212edl(i["dscp21"], d, pre_append)
			tmp["dscp21"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp21")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := i["dscp22"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp222edl(i["dscp22"], d, pre_append)
			tmp["dscp22"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp22")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := i["dscp23"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp232edl(i["dscp23"], d, pre_append)
			tmp["dscp23"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp23")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := i["dscp24"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp242edl(i["dscp24"], d, pre_append)
			tmp["dscp24"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp24")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := i["dscp25"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp252edl(i["dscp25"], d, pre_append)
			tmp["dscp25"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp25")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := i["dscp26"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp262edl(i["dscp26"], d, pre_append)
			tmp["dscp26"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp26")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := i["dscp27"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp272edl(i["dscp27"], d, pre_append)
			tmp["dscp27"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp27")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := i["dscp28"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp282edl(i["dscp28"], d, pre_append)
			tmp["dscp28"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp28")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := i["dscp29"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp292edl(i["dscp29"], d, pre_append)
			tmp["dscp29"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp29")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := i["dscp3"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp32edl(i["dscp3"], d, pre_append)
			tmp["dscp3"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp3")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := i["dscp30"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp302edl(i["dscp30"], d, pre_append)
			tmp["dscp30"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp30")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := i["dscp31"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp312edl(i["dscp31"], d, pre_append)
			tmp["dscp31"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp31")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := i["dscp32"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp322edl(i["dscp32"], d, pre_append)
			tmp["dscp32"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp32")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := i["dscp33"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp332edl(i["dscp33"], d, pre_append)
			tmp["dscp33"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp33")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := i["dscp34"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp342edl(i["dscp34"], d, pre_append)
			tmp["dscp34"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp34")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := i["dscp35"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp352edl(i["dscp35"], d, pre_append)
			tmp["dscp35"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp35")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := i["dscp36"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp362edl(i["dscp36"], d, pre_append)
			tmp["dscp36"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp36")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := i["dscp37"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp372edl(i["dscp37"], d, pre_append)
			tmp["dscp37"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp37")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := i["dscp38"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp382edl(i["dscp38"], d, pre_append)
			tmp["dscp38"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp38")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := i["dscp39"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp392edl(i["dscp39"], d, pre_append)
			tmp["dscp39"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp39")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := i["dscp4"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp42edl(i["dscp4"], d, pre_append)
			tmp["dscp4"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp4")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := i["dscp40"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp402edl(i["dscp40"], d, pre_append)
			tmp["dscp40"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp40")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := i["dscp41"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp412edl(i["dscp41"], d, pre_append)
			tmp["dscp41"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp41")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := i["dscp42"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp422edl(i["dscp42"], d, pre_append)
			tmp["dscp42"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp42")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := i["dscp43"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp432edl(i["dscp43"], d, pre_append)
			tmp["dscp43"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp43")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := i["dscp44"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp442edl(i["dscp44"], d, pre_append)
			tmp["dscp44"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp44")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := i["dscp45"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp452edl(i["dscp45"], d, pre_append)
			tmp["dscp45"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp45")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := i["dscp46"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp462edl(i["dscp46"], d, pre_append)
			tmp["dscp46"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp46")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := i["dscp47"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp472edl(i["dscp47"], d, pre_append)
			tmp["dscp47"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp47")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := i["dscp48"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp482edl(i["dscp48"], d, pre_append)
			tmp["dscp48"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp48")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := i["dscp49"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp492edl(i["dscp49"], d, pre_append)
			tmp["dscp49"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp49")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := i["dscp5"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp52edl(i["dscp5"], d, pre_append)
			tmp["dscp5"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp5")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := i["dscp50"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp502edl(i["dscp50"], d, pre_append)
			tmp["dscp50"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp50")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := i["dscp51"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp512edl(i["dscp51"], d, pre_append)
			tmp["dscp51"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp51")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := i["dscp52"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp522edl(i["dscp52"], d, pre_append)
			tmp["dscp52"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp52")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := i["dscp53"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp532edl(i["dscp53"], d, pre_append)
			tmp["dscp53"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp53")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := i["dscp54"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp542edl(i["dscp54"], d, pre_append)
			tmp["dscp54"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp54")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := i["dscp55"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp552edl(i["dscp55"], d, pre_append)
			tmp["dscp55"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp55")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := i["dscp56"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp562edl(i["dscp56"], d, pre_append)
			tmp["dscp56"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp56")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := i["dscp57"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp572edl(i["dscp57"], d, pre_append)
			tmp["dscp57"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp57")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := i["dscp58"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp582edl(i["dscp58"], d, pre_append)
			tmp["dscp58"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp58")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := i["dscp59"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp592edl(i["dscp59"], d, pre_append)
			tmp["dscp59"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp59")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := i["dscp6"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp62edl(i["dscp6"], d, pre_append)
			tmp["dscp6"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := i["dscp60"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp602edl(i["dscp60"], d, pre_append)
			tmp["dscp60"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp60")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := i["dscp61"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp612edl(i["dscp61"], d, pre_append)
			tmp["dscp61"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp61")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := i["dscp62"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp622edl(i["dscp62"], d, pre_append)
			tmp["dscp62"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp62")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := i["dscp63"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp632edl(i["dscp63"], d, pre_append)
			tmp["dscp63"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp63")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := i["dscp7"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp72edl(i["dscp7"], d, pre_append)
			tmp["dscp7"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp7")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := i["dscp8"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp82edl(i["dscp8"], d, pre_append)
			tmp["dscp8"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp8")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := i["dscp9"]; ok {
			v := flattenSystemNpuNpQueuesProfileDscp92edl(i["dscp9"], d, pre_append)
			tmp["dscp9"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Dscp9")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemNpuNpQueuesProfileId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemNpuNpQueuesProfileType2edl(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSystemNpuNpQueuesProfileWeight2edl(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Profile-Weight")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNpuNpQueuesProfileCos02edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileCos72edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp02edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp12edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp102edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp112edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp122edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp132edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp142edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp152edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp162edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp172edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp182edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp192edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp202edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp212edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp222edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp232edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp242edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp252edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp262edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp272edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp282edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp292edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp32edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp302edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp312edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp322edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp332edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp342edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp352edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp362edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp372edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp382edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp392edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp42edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp402edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp412edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp422edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp432edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp442edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp452edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp462edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp472edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp482edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp492edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp52edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp502edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp512edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp522edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp532edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp542edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp552edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp562edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp572edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp582edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp592edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp62edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp602edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp612edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp622edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp632edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp72edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp82edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileDscp92edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesProfileWeight2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesScheduler2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenSystemNpuNpQueuesSchedulerMode2edl(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Scheduler-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemNpuNpQueuesSchedulerName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNpuNpQueues-Scheduler-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNpuNpQueuesSchedulerMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNpuNpQueuesSchedulerName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNpuNpQueues(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("custom_etype_lookup", flattenSystemNpuNpQueuesCustomEtypeLookup2edl(o["custom-etype-lookup"], d, "custom_etype_lookup")); err != nil {
		if vv, ok := fortiAPIPatch(o["custom-etype-lookup"], "SystemNpuNpQueues-CustomEtypeLookup"); ok {
			if err = d.Set("custom_etype_lookup", vv); err != nil {
				return fmt.Errorf("Error reading custom_etype_lookup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading custom_etype_lookup: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ethernet_type", flattenSystemNpuNpQueuesEthernetType2edl(o["ethernet-type"], d, "ethernet_type")); err != nil {
			if vv, ok := fortiAPIPatch(o["ethernet-type"], "SystemNpuNpQueues-EthernetType"); ok {
				if err = d.Set("ethernet_type", vv); err != nil {
					return fmt.Errorf("Error reading ethernet_type: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ethernet_type: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ethernet_type"); ok {
			if err = d.Set("ethernet_type", flattenSystemNpuNpQueuesEthernetType2edl(o["ethernet-type"], d, "ethernet_type")); err != nil {
				if vv, ok := fortiAPIPatch(o["ethernet-type"], "SystemNpuNpQueues-EthernetType"); ok {
					if err = d.Set("ethernet_type", vv); err != nil {
						return fmt.Errorf("Error reading ethernet_type: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ethernet_type: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_protocol", flattenSystemNpuNpQueuesIpProtocol2edl(o["ip-protocol"], d, "ip_protocol")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-protocol"], "SystemNpuNpQueues-IpProtocol"); ok {
				if err = d.Set("ip_protocol", vv); err != nil {
					return fmt.Errorf("Error reading ip_protocol: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_protocol: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_protocol"); ok {
			if err = d.Set("ip_protocol", flattenSystemNpuNpQueuesIpProtocol2edl(o["ip-protocol"], d, "ip_protocol")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-protocol"], "SystemNpuNpQueues-IpProtocol"); ok {
					if err = d.Set("ip_protocol", vv); err != nil {
						return fmt.Errorf("Error reading ip_protocol: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_protocol: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_service", flattenSystemNpuNpQueuesIpService2edl(o["ip-service"], d, "ip_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["ip-service"], "SystemNpuNpQueues-IpService"); ok {
				if err = d.Set("ip_service", vv); err != nil {
					return fmt.Errorf("Error reading ip_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ip_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_service"); ok {
			if err = d.Set("ip_service", flattenSystemNpuNpQueuesIpService2edl(o["ip-service"], d, "ip_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["ip-service"], "SystemNpuNpQueues-IpService"); ok {
					if err = d.Set("ip_service", vv); err != nil {
						return fmt.Errorf("Error reading ip_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ip_service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("profile", flattenSystemNpuNpQueuesProfile2edl(o["profile"], d, "profile")); err != nil {
			if vv, ok := fortiAPIPatch(o["profile"], "SystemNpuNpQueues-Profile"); ok {
				if err = d.Set("profile", vv); err != nil {
					return fmt.Errorf("Error reading profile: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading profile: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("profile"); ok {
			if err = d.Set("profile", flattenSystemNpuNpQueuesProfile2edl(o["profile"], d, "profile")); err != nil {
				if vv, ok := fortiAPIPatch(o["profile"], "SystemNpuNpQueues-Profile"); ok {
					if err = d.Set("profile", vv); err != nil {
						return fmt.Errorf("Error reading profile: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading profile: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("scheduler", flattenSystemNpuNpQueuesScheduler2edl(o["scheduler"], d, "scheduler")); err != nil {
			if vv, ok := fortiAPIPatch(o["scheduler"], "SystemNpuNpQueues-Scheduler"); ok {
				if err = d.Set("scheduler", vv); err != nil {
					return fmt.Errorf("Error reading scheduler: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scheduler: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scheduler"); ok {
			if err = d.Set("scheduler", flattenSystemNpuNpQueuesScheduler2edl(o["scheduler"], d, "scheduler")); err != nil {
				if vv, ok := fortiAPIPatch(o["scheduler"], "SystemNpuNpQueues-Scheduler"); ok {
					if err = d.Set("scheduler", vv); err != nil {
						return fmt.Errorf("Error reading scheduler: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scheduler: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemNpuNpQueuesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNpuNpQueuesCustomEtypeLookup2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpQueuesEthernetTypeName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandSystemNpuNpQueuesEthernetTypeQueue2edl(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemNpuNpQueuesEthernetTypeType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSystemNpuNpQueuesEthernetTypeWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesEthernetTypeName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetTypeQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetTypeType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesEthernetTypeWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNpuNpQueuesIpProtocolName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandSystemNpuNpQueuesIpProtocolProtocol2edl(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandSystemNpuNpQueuesIpProtocolQueue2edl(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSystemNpuNpQueuesIpProtocolWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesIpProtocolName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocolProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocolQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpProtocolWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dport"], _ = expandSystemNpuNpQueuesIpServiceDport2edl(d, i["dport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemNpuNpQueuesIpServiceName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandSystemNpuNpQueuesIpServiceProtocol2edl(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queue"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["queue"], _ = expandSystemNpuNpQueuesIpServiceQueue2edl(d, i["queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sport"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sport"], _ = expandSystemNpuNpQueuesIpServiceSport2edl(d, i["sport"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSystemNpuNpQueuesIpServiceWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesIpServiceDport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceQueue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceSport2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesIpServiceWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfile2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos0"], _ = expandSystemNpuNpQueuesProfileCos02edl(d, i["cos0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos1"], _ = expandSystemNpuNpQueuesProfileCos12edl(d, i["cos1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos2"], _ = expandSystemNpuNpQueuesProfileCos22edl(d, i["cos2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos3"], _ = expandSystemNpuNpQueuesProfileCos32edl(d, i["cos3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos4"], _ = expandSystemNpuNpQueuesProfileCos42edl(d, i["cos4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos5"], _ = expandSystemNpuNpQueuesProfileCos52edl(d, i["cos5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos6"], _ = expandSystemNpuNpQueuesProfileCos62edl(d, i["cos6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cos7"], _ = expandSystemNpuNpQueuesProfileCos72edl(d, i["cos7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp0"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp0"], _ = expandSystemNpuNpQueuesProfileDscp02edl(d, i["dscp0"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp1"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp1"], _ = expandSystemNpuNpQueuesProfileDscp12edl(d, i["dscp1"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp10"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp10"], _ = expandSystemNpuNpQueuesProfileDscp102edl(d, i["dscp10"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp11"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp11"], _ = expandSystemNpuNpQueuesProfileDscp112edl(d, i["dscp11"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp12"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp12"], _ = expandSystemNpuNpQueuesProfileDscp122edl(d, i["dscp12"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp13"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp13"], _ = expandSystemNpuNpQueuesProfileDscp132edl(d, i["dscp13"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp14"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp14"], _ = expandSystemNpuNpQueuesProfileDscp142edl(d, i["dscp14"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp15"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp15"], _ = expandSystemNpuNpQueuesProfileDscp152edl(d, i["dscp15"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp16"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp16"], _ = expandSystemNpuNpQueuesProfileDscp162edl(d, i["dscp16"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp17"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp17"], _ = expandSystemNpuNpQueuesProfileDscp172edl(d, i["dscp17"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp18"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp18"], _ = expandSystemNpuNpQueuesProfileDscp182edl(d, i["dscp18"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp19"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp19"], _ = expandSystemNpuNpQueuesProfileDscp192edl(d, i["dscp19"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp2"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp2"], _ = expandSystemNpuNpQueuesProfileDscp22edl(d, i["dscp2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp20"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp20"], _ = expandSystemNpuNpQueuesProfileDscp202edl(d, i["dscp20"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp21"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp21"], _ = expandSystemNpuNpQueuesProfileDscp212edl(d, i["dscp21"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp22"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp22"], _ = expandSystemNpuNpQueuesProfileDscp222edl(d, i["dscp22"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp23"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp23"], _ = expandSystemNpuNpQueuesProfileDscp232edl(d, i["dscp23"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp24"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp24"], _ = expandSystemNpuNpQueuesProfileDscp242edl(d, i["dscp24"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp25"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp25"], _ = expandSystemNpuNpQueuesProfileDscp252edl(d, i["dscp25"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp26"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp26"], _ = expandSystemNpuNpQueuesProfileDscp262edl(d, i["dscp26"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp27"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp27"], _ = expandSystemNpuNpQueuesProfileDscp272edl(d, i["dscp27"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp28"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp28"], _ = expandSystemNpuNpQueuesProfileDscp282edl(d, i["dscp28"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp29"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp29"], _ = expandSystemNpuNpQueuesProfileDscp292edl(d, i["dscp29"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp3"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp3"], _ = expandSystemNpuNpQueuesProfileDscp32edl(d, i["dscp3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp30"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp30"], _ = expandSystemNpuNpQueuesProfileDscp302edl(d, i["dscp30"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp31"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp31"], _ = expandSystemNpuNpQueuesProfileDscp312edl(d, i["dscp31"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp32"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp32"], _ = expandSystemNpuNpQueuesProfileDscp322edl(d, i["dscp32"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp33"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp33"], _ = expandSystemNpuNpQueuesProfileDscp332edl(d, i["dscp33"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp34"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp34"], _ = expandSystemNpuNpQueuesProfileDscp342edl(d, i["dscp34"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp35"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp35"], _ = expandSystemNpuNpQueuesProfileDscp352edl(d, i["dscp35"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp36"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp36"], _ = expandSystemNpuNpQueuesProfileDscp362edl(d, i["dscp36"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp37"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp37"], _ = expandSystemNpuNpQueuesProfileDscp372edl(d, i["dscp37"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp38"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp38"], _ = expandSystemNpuNpQueuesProfileDscp382edl(d, i["dscp38"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp39"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp39"], _ = expandSystemNpuNpQueuesProfileDscp392edl(d, i["dscp39"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp4"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp4"], _ = expandSystemNpuNpQueuesProfileDscp42edl(d, i["dscp4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp40"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp40"], _ = expandSystemNpuNpQueuesProfileDscp402edl(d, i["dscp40"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp41"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp41"], _ = expandSystemNpuNpQueuesProfileDscp412edl(d, i["dscp41"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp42"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp42"], _ = expandSystemNpuNpQueuesProfileDscp422edl(d, i["dscp42"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp43"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp43"], _ = expandSystemNpuNpQueuesProfileDscp432edl(d, i["dscp43"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp44"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp44"], _ = expandSystemNpuNpQueuesProfileDscp442edl(d, i["dscp44"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp45"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp45"], _ = expandSystemNpuNpQueuesProfileDscp452edl(d, i["dscp45"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp46"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp46"], _ = expandSystemNpuNpQueuesProfileDscp462edl(d, i["dscp46"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp47"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp47"], _ = expandSystemNpuNpQueuesProfileDscp472edl(d, i["dscp47"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp48"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp48"], _ = expandSystemNpuNpQueuesProfileDscp482edl(d, i["dscp48"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp49"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp49"], _ = expandSystemNpuNpQueuesProfileDscp492edl(d, i["dscp49"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp5"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp5"], _ = expandSystemNpuNpQueuesProfileDscp52edl(d, i["dscp5"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp50"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp50"], _ = expandSystemNpuNpQueuesProfileDscp502edl(d, i["dscp50"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp51"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp51"], _ = expandSystemNpuNpQueuesProfileDscp512edl(d, i["dscp51"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp52"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp52"], _ = expandSystemNpuNpQueuesProfileDscp522edl(d, i["dscp52"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp53"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp53"], _ = expandSystemNpuNpQueuesProfileDscp532edl(d, i["dscp53"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp54"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp54"], _ = expandSystemNpuNpQueuesProfileDscp542edl(d, i["dscp54"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp55"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp55"], _ = expandSystemNpuNpQueuesProfileDscp552edl(d, i["dscp55"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp56"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp56"], _ = expandSystemNpuNpQueuesProfileDscp562edl(d, i["dscp56"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp57"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp57"], _ = expandSystemNpuNpQueuesProfileDscp572edl(d, i["dscp57"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp58"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp58"], _ = expandSystemNpuNpQueuesProfileDscp582edl(d, i["dscp58"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp59"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp59"], _ = expandSystemNpuNpQueuesProfileDscp592edl(d, i["dscp59"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp6"], _ = expandSystemNpuNpQueuesProfileDscp62edl(d, i["dscp6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp60"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp60"], _ = expandSystemNpuNpQueuesProfileDscp602edl(d, i["dscp60"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp61"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp61"], _ = expandSystemNpuNpQueuesProfileDscp612edl(d, i["dscp61"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp62"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp62"], _ = expandSystemNpuNpQueuesProfileDscp622edl(d, i["dscp62"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp63"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp63"], _ = expandSystemNpuNpQueuesProfileDscp632edl(d, i["dscp63"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp7"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp7"], _ = expandSystemNpuNpQueuesProfileDscp72edl(d, i["dscp7"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp8"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp8"], _ = expandSystemNpuNpQueuesProfileDscp82edl(d, i["dscp8"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp9"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp9"], _ = expandSystemNpuNpQueuesProfileDscp92edl(d, i["dscp9"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemNpuNpQueuesProfileId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemNpuNpQueuesProfileType2edl(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSystemNpuNpQueuesProfileWeight2edl(d, i["weight"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesProfileCos02edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileCos72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp02edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp12edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp102edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp112edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp122edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp132edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp142edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp152edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp162edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp172edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp182edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp192edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp202edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp212edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp222edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp232edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp242edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp252edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp262edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp272edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp282edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp292edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp32edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp302edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp312edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp322edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp332edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp342edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp352edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp362edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp372edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp382edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp392edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp42edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp402edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp412edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp422edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp432edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp442edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp452edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp462edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp472edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp482edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp492edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp52edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp502edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp512edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp522edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp532edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp542edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp552edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp562edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp572edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp582edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp592edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp62edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp602edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp612edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp622edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp632edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp72edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp82edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileDscp92edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesProfileWeight2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesScheduler2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandSystemNpuNpQueuesSchedulerMode2edl(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemNpuNpQueuesSchedulerName2edl(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNpuNpQueuesSchedulerMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNpQueuesSchedulerName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNpuNpQueues(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("custom_etype_lookup"); ok || d.HasChange("custom_etype_lookup") {
		t, err := expandSystemNpuNpQueuesCustomEtypeLookup2edl(d, v, "custom_etype_lookup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-etype-lookup"] = t
		}
	}

	if v, ok := d.GetOk("ethernet_type"); ok || d.HasChange("ethernet_type") {
		t, err := expandSystemNpuNpQueuesEthernetType2edl(d, v, "ethernet_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ethernet-type"] = t
		}
	}

	if v, ok := d.GetOk("ip_protocol"); ok || d.HasChange("ip_protocol") {
		t, err := expandSystemNpuNpQueuesIpProtocol2edl(d, v, "ip_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ip_service"); ok || d.HasChange("ip_service") {
		t, err := expandSystemNpuNpQueuesIpService2edl(d, v, "ip_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-service"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok || d.HasChange("profile") {
		t, err := expandSystemNpuNpQueuesProfile2edl(d, v, "profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("scheduler"); ok || d.HasChange("scheduler") {
		t, err := expandSystemNpuNpQueuesScheduler2edl(d, v, "scheduler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scheduler"] = t
		}
	}

	return &obj, nil
}
