// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device SystemNethsm

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNethsm() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNethsmUpdate,
		Read:   resourceSystemNethsmRead,
		Update: resourceSystemNethsmUpdate,
		Delete: resourceSystemNethsmDelete,

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
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_status_pulling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hagroups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
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
			"partitions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pkcs11_pin": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"slot_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"primus_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"receivetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rsa_mech_remap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret_content": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"htl": &schema.Schema{
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
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"slots": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"for_ha": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
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

func resourceSystemNethsmUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNethsm(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNethsm resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemNethsm(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNethsm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemNethsm")

	return resourceSystemNethsmRead(d, m)
}

func resourceSystemNethsmDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNethsm(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNethsm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNethsmRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNethsm(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemNethsm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNethsm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNethsm resource from API: %v", err)
	}
	return nil
}

func flattenSystemNethsmHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmHaStatusPullingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmHagroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenSystemNethsmHagroupsMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "SystemNethsm-Hagroups-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemNethsmHagroupsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNethsm-Hagroups-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNethsmHagroupsMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemNethsmHagroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemNethsmPartitions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemNethsmPartitionsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNethsm-Partitions-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "slot_id"
		if _, ok := i["slot-id"]; ok {
			v := flattenSystemNethsmPartitionsSlotId(i["slot-id"], d, pre_append)
			tmp["slot_id"] = fortiAPISubPartPatch(v, "SystemNethsm-Partitions-SlotId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNethsmPartitionsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmPartitionsSlotId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmPrimusCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmReceivetimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmRsaMechRemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmServers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "htl"
		if _, ok := i["htl"]; ok {
			v := flattenSystemNethsmServersHtl(i["htl"], d, pre_append)
			tmp["htl"] = fortiAPISubPartPatch(v, "SystemNethsm-Servers-Htl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemNethsmServersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNethsm-Servers-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSystemNethsmServersPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SystemNethsm-Servers-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSystemNethsmServersServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SystemNethsm-Servers-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_cert"
		if _, ok := i["server-cert"]; ok {
			v := flattenSystemNethsmServersServerCert(i["server-cert"], d, pre_append)
			tmp["server_cert"] = fortiAPISubPartPatch(v, "SystemNethsm-Servers-ServerCert")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNethsmServersHtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmServersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmServersPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmServersServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmServersServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmSlots(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "for_ha"
		if _, ok := i["for-ha"]; ok {
			v := flattenSystemNethsmSlotsForHa(i["for-ha"], d, pre_append)
			tmp["for_ha"] = fortiAPISubPartPatch(v, "SystemNethsm-Slots-ForHa")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemNethsmSlotsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemNethsm-Slots-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemNethsmSlotsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemNethsm-Slots-Name")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemNethsmSlotsForHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmSlotsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmSlotsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNethsmVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNethsm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("ha", flattenSystemNethsmHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "SystemNethsm-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("ha_status_pulling_interval", flattenSystemNethsmHaStatusPullingInterval(o["ha-status-pulling-interval"], d, "ha_status_pulling_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha-status-pulling-interval"], "SystemNethsm-HaStatusPullingInterval"); ok {
			if err = d.Set("ha_status_pulling_interval", vv); err != nil {
				return fmt.Errorf("Error reading ha_status_pulling_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_status_pulling_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("hagroups", flattenSystemNethsmHagroups(o["hagroups"], d, "hagroups")); err != nil {
			if vv, ok := fortiAPIPatch(o["hagroups"], "SystemNethsm-Hagroups"); ok {
				if err = d.Set("hagroups", vv); err != nil {
					return fmt.Errorf("Error reading hagroups: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading hagroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hagroups"); ok {
			if err = d.Set("hagroups", flattenSystemNethsmHagroups(o["hagroups"], d, "hagroups")); err != nil {
				if vv, ok := fortiAPIPatch(o["hagroups"], "SystemNethsm-Hagroups"); ok {
					if err = d.Set("hagroups", vv); err != nil {
						return fmt.Errorf("Error reading hagroups: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading hagroups: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenSystemNethsmInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemNethsm-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("partitions", flattenSystemNethsmPartitions(o["partitions"], d, "partitions")); err != nil {
			if vv, ok := fortiAPIPatch(o["partitions"], "SystemNethsm-Partitions"); ok {
				if err = d.Set("partitions", vv); err != nil {
					return fmt.Errorf("Error reading partitions: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading partitions: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("partitions"); ok {
			if err = d.Set("partitions", flattenSystemNethsmPartitions(o["partitions"], d, "partitions")); err != nil {
				if vv, ok := fortiAPIPatch(o["partitions"], "SystemNethsm-Partitions"); ok {
					if err = d.Set("partitions", vv); err != nil {
						return fmt.Errorf("Error reading partitions: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading partitions: %v", err)
				}
			}
		}
	}

	if err = d.Set("primus_cfg", flattenSystemNethsmPrimusCfg(o["primus-cfg"], d, "primus_cfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["primus-cfg"], "SystemNethsm-PrimusCfg"); ok {
			if err = d.Set("primus_cfg", vv); err != nil {
				return fmt.Errorf("Error reading primus_cfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primus_cfg: %v", err)
		}
	}

	if err = d.Set("receivetimeout", flattenSystemNethsmReceivetimeout(o["receivetimeout"], d, "receivetimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["receivetimeout"], "SystemNethsm-Receivetimeout"); ok {
			if err = d.Set("receivetimeout", vv); err != nil {
				return fmt.Errorf("Error reading receivetimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading receivetimeout: %v", err)
		}
	}

	if err = d.Set("rsa_mech_remap", flattenSystemNethsmRsaMechRemap(o["rsa-mech-remap"], d, "rsa_mech_remap")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsa-mech-remap"], "SystemNethsm-RsaMechRemap"); ok {
			if err = d.Set("rsa_mech_remap", vv); err != nil {
				return fmt.Errorf("Error reading rsa_mech_remap: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsa_mech_remap: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("servers", flattenSystemNethsmServers(o["servers"], d, "servers")); err != nil {
			if vv, ok := fortiAPIPatch(o["servers"], "SystemNethsm-Servers"); ok {
				if err = d.Set("servers", vv); err != nil {
					return fmt.Errorf("Error reading servers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading servers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("servers"); ok {
			if err = d.Set("servers", flattenSystemNethsmServers(o["servers"], d, "servers")); err != nil {
				if vv, ok := fortiAPIPatch(o["servers"], "SystemNethsm-Servers"); ok {
					if err = d.Set("servers", vv); err != nil {
						return fmt.Errorf("Error reading servers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading servers: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("slots", flattenSystemNethsmSlots(o["slots"], d, "slots")); err != nil {
			if vv, ok := fortiAPIPatch(o["slots"], "SystemNethsm-Slots"); ok {
				if err = d.Set("slots", vv); err != nil {
					return fmt.Errorf("Error reading slots: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading slots: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("slots"); ok {
			if err = d.Set("slots", flattenSystemNethsmSlots(o["slots"], d, "slots")); err != nil {
				if vv, ok := fortiAPIPatch(o["slots"], "SystemNethsm-Slots"); ok {
					if err = d.Set("slots", vv); err != nil {
						return fmt.Errorf("Error reading slots: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading slots: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenSystemNethsmStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemNethsm-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vendor", flattenSystemNethsmVendor(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "SystemNethsm-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenSystemNethsmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNethsmHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmHaStatusPullingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmHagroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandSystemNethsmHagroupsMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemNethsmHagroupsName(d, i["name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNethsmHagroupsMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNethsmHagroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNethsmPartitions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemNethsmPartitionsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pkcs11_pin"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pkcs11-pin"], _ = expandSystemNethsmPartitionsPkcs11Pin(d, i["pkcs11_pin"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "slot_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["slot-id"], _ = expandSystemNethsmPartitionsSlotId(d, i["slot_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNethsmPartitionsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmPartitionsPkcs11Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNethsmPartitionsSlotId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmPrimusCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmReceivetimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmRsaMechRemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSecretContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNethsmServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "htl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["htl"], _ = expandSystemNethsmServersHtl(d, i["htl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemNethsmServersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSystemNethsmServersPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSystemNethsmServersServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-cert"], _ = expandSystemNethsmServersServerCert(d, i["server_cert"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNethsmServersHtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmServersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmServersPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmServersServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmServersServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlots(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "for_ha"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["for-ha"], _ = expandSystemNethsmSlotsForHa(d, i["for_ha"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemNethsmSlotsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemNethsmSlotsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandSystemNethsmSlotsPassword(d, i["password"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemNethsmSlotsForHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlotsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlotsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmSlotsPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemNethsmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNethsmVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNethsm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ha"); ok || d.HasChange("ha") {
		t, err := expandSystemNethsmHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("ha_status_pulling_interval"); ok || d.HasChange("ha_status_pulling_interval") {
		t, err := expandSystemNethsmHaStatusPullingInterval(d, v, "ha_status_pulling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-status-pulling-interval"] = t
		}
	}

	if v, ok := d.GetOk("hagroups"); ok || d.HasChange("hagroups") {
		t, err := expandSystemNethsmHagroups(d, v, "hagroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hagroups"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemNethsmInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("partitions"); ok || d.HasChange("partitions") {
		t, err := expandSystemNethsmPartitions(d, v, "partitions")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partitions"] = t
		}
	}

	if v, ok := d.GetOk("primus_cfg"); ok || d.HasChange("primus_cfg") {
		t, err := expandSystemNethsmPrimusCfg(d, v, "primus_cfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primus-cfg"] = t
		}
	}

	if v, ok := d.GetOk("receivetimeout"); ok || d.HasChange("receivetimeout") {
		t, err := expandSystemNethsmReceivetimeout(d, v, "receivetimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["receivetimeout"] = t
		}
	}

	if v, ok := d.GetOk("rsa_mech_remap"); ok || d.HasChange("rsa_mech_remap") {
		t, err := expandSystemNethsmRsaMechRemap(d, v, "rsa_mech_remap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-mech-remap"] = t
		}
	}

	if v, ok := d.GetOk("secret_content"); ok || d.HasChange("secret_content") {
		t, err := expandSystemNethsmSecretContent(d, v, "secret_content")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-content"] = t
		}
	}

	if v, ok := d.GetOk("servers"); ok || d.HasChange("servers") {
		t, err := expandSystemNethsmServers(d, v, "servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["servers"] = t
		}
	}

	if v, ok := d.GetOk("slots"); ok || d.HasChange("slots") {
		t, err := expandSystemNethsmSlots(d, v, "slots")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slots"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemNethsmStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandSystemNethsmVendor(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
