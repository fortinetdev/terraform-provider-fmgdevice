// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Application list entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationListEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationListEntriesCreate,
		Read:   resourceApplicationListEntriesRead,
		Update: resourceApplicationListEntriesUpdate,
		Delete: resourceApplicationListEntriesDelete,

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
			"list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"behavior": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"exclusion": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"popularity": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"protocols": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rate_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rate_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate_track": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"risk": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"shaper": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"shaper_reverse": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sub_category": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceApplicationListEntriesCreate(d *schema.ResourceData, m interface{}) error {
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
	list := d.Get("list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	obj, err := getObjectApplicationListEntries(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationListEntries resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadApplicationListEntries(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateApplicationListEntries(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ApplicationListEntries resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateApplicationListEntries(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ApplicationListEntries resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceApplicationListEntriesRead(d, m)
}

func resourceApplicationListEntriesUpdate(d *schema.ResourceData, m interface{}) error {
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
	list := d.Get("list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	obj, err := getObjectApplicationListEntries(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationListEntries resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateApplicationListEntries(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationListEntries resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceApplicationListEntriesRead(d, m)
}

func resourceApplicationListEntriesDelete(d *schema.ResourceData, m interface{}) error {
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
	list := d.Get("list").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	wsParams["adom"] = adomv

	err = c.DeleteApplicationListEntries(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationListEntries resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationListEntriesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	list := d.Get("list").(string)
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
	if list == "" {
		list = importOptionChecking(m.(*FortiClient).Cfg, "list")
		if list == "" {
			return fmt.Errorf("Parameter list is missing")
		}
		if err = d.Set("list", list); err != nil {
			return fmt.Errorf("Error set params list: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["list"] = list

	o, err := c.ReadApplicationListEntries(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ApplicationListEntries resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationListEntries(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationListEntries resource from API: %v", err)
	}
	return nil
}

func flattenApplicationListEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesApplication2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesBehavior2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesExclusion2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesLogPacket2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParameters2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenApplicationListEntriesParametersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ApplicationListEntries-Parameters-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenApplicationListEntriesParametersMembers2edl(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "ApplicationListEntries-Parameters-Members")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenApplicationListEntriesParametersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembers2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenApplicationListEntriesParametersMembersId2edl(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ApplicationListEntriesParameters-Members-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenApplicationListEntriesParametersMembersName2edl(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ApplicationListEntriesParameters-Members-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenApplicationListEntriesParametersMembersValue2edl(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ApplicationListEntriesParameters-Members-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenApplicationListEntriesParametersMembersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembersName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembersValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesPerIpShaper2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesPopularity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesProtocols2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateCount2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateMode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateTrack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRisk2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesSessionTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesShaper2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesShaperReverse2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesSubCategory2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenApplicationListEntriesTechnology2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenApplicationListEntriesVendor2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectApplicationListEntries(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("action", flattenApplicationListEntriesAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "ApplicationListEntries-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("application", flattenApplicationListEntriesApplication2edl(o["application"], d, "application")); err != nil {
		if vv, ok := fortiAPIPatch(o["application"], "ApplicationListEntries-Application"); ok {
			if err = d.Set("application", vv); err != nil {
				return fmt.Errorf("Error reading application: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("behavior", flattenApplicationListEntriesBehavior2edl(o["behavior"], d, "behavior")); err != nil {
		if vv, ok := fortiAPIPatch(o["behavior"], "ApplicationListEntries-Behavior"); ok {
			if err = d.Set("behavior", vv); err != nil {
				return fmt.Errorf("Error reading behavior: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationListEntriesCategory2edl(o["category"], d, "category")); err != nil {
		if vv, ok := fortiAPIPatch(o["category"], "ApplicationListEntries-Category"); ok {
			if err = d.Set("category", vv); err != nil {
				return fmt.Errorf("Error reading category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("exclusion", flattenApplicationListEntriesExclusion2edl(o["exclusion"], d, "exclusion")); err != nil {
		if vv, ok := fortiAPIPatch(o["exclusion"], "ApplicationListEntries-Exclusion"); ok {
			if err = d.Set("exclusion", vv); err != nil {
				return fmt.Errorf("Error reading exclusion: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading exclusion: %v", err)
		}
	}

	if err = d.Set("fosid", flattenApplicationListEntriesId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ApplicationListEntries-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("log", flattenApplicationListEntriesLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "ApplicationListEntries-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenApplicationListEntriesLogPacket2edl(o["log-packet"], d, "log_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-packet"], "ApplicationListEntries-LogPacket"); ok {
			if err = d.Set("log_packet", vv); err != nil {
				return fmt.Errorf("Error reading log_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("parameters", flattenApplicationListEntriesParameters2edl(o["parameters"], d, "parameters")); err != nil {
			if vv, ok := fortiAPIPatch(o["parameters"], "ApplicationListEntries-Parameters"); ok {
				if err = d.Set("parameters", vv); err != nil {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading parameters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameters"); ok {
			if err = d.Set("parameters", flattenApplicationListEntriesParameters2edl(o["parameters"], d, "parameters")); err != nil {
				if vv, ok := fortiAPIPatch(o["parameters"], "ApplicationListEntries-Parameters"); ok {
					if err = d.Set("parameters", vv); err != nil {
						return fmt.Errorf("Error reading parameters: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			}
		}
	}

	if err = d.Set("per_ip_shaper", flattenApplicationListEntriesPerIpShaper2edl(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "ApplicationListEntries-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("popularity", flattenApplicationListEntriesPopularity2edl(o["popularity"], d, "popularity")); err != nil {
		if vv, ok := fortiAPIPatch(o["popularity"], "ApplicationListEntries-Popularity"); ok {
			if err = d.Set("popularity", vv); err != nil {
				return fmt.Errorf("Error reading popularity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	if err = d.Set("protocols", flattenApplicationListEntriesProtocols2edl(o["protocols"], d, "protocols")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocols"], "ApplicationListEntries-Protocols"); ok {
			if err = d.Set("protocols", vv); err != nil {
				return fmt.Errorf("Error reading protocols: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocols: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenApplicationListEntriesQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "ApplicationListEntries-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenApplicationListEntriesQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "ApplicationListEntries-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenApplicationListEntriesQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "ApplicationListEntries-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("rate_count", flattenApplicationListEntriesRateCount2edl(o["rate-count"], d, "rate_count")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-count"], "ApplicationListEntries-RateCount"); ok {
			if err = d.Set("rate_count", vv); err != nil {
				return fmt.Errorf("Error reading rate_count: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_count: %v", err)
		}
	}

	if err = d.Set("rate_duration", flattenApplicationListEntriesRateDuration2edl(o["rate-duration"], d, "rate_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-duration"], "ApplicationListEntries-RateDuration"); ok {
			if err = d.Set("rate_duration", vv); err != nil {
				return fmt.Errorf("Error reading rate_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_duration: %v", err)
		}
	}

	if err = d.Set("rate_mode", flattenApplicationListEntriesRateMode2edl(o["rate-mode"], d, "rate_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-mode"], "ApplicationListEntries-RateMode"); ok {
			if err = d.Set("rate_mode", vv); err != nil {
				return fmt.Errorf("Error reading rate_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_mode: %v", err)
		}
	}

	if err = d.Set("rate_track", flattenApplicationListEntriesRateTrack2edl(o["rate-track"], d, "rate_track")); err != nil {
		if vv, ok := fortiAPIPatch(o["rate-track"], "ApplicationListEntries-RateTrack"); ok {
			if err = d.Set("rate_track", vv); err != nil {
				return fmt.Errorf("Error reading rate_track: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rate_track: %v", err)
		}
	}

	if err = d.Set("risk", flattenApplicationListEntriesRisk2edl(o["risk"], d, "risk")); err != nil {
		if vv, ok := fortiAPIPatch(o["risk"], "ApplicationListEntries-Risk"); ok {
			if err = d.Set("risk", vv); err != nil {
				return fmt.Errorf("Error reading risk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading risk: %v", err)
		}
	}

	if err = d.Set("session_ttl", flattenApplicationListEntriesSessionTtl2edl(o["session-ttl"], d, "session_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-ttl"], "ApplicationListEntries-SessionTtl"); ok {
			if err = d.Set("session_ttl", vv); err != nil {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("shaper", flattenApplicationListEntriesShaper2edl(o["shaper"], d, "shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["shaper"], "ApplicationListEntries-Shaper"); ok {
			if err = d.Set("shaper", vv); err != nil {
				return fmt.Errorf("Error reading shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shaper: %v", err)
		}
	}

	if err = d.Set("shaper_reverse", flattenApplicationListEntriesShaperReverse2edl(o["shaper-reverse"], d, "shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["shaper-reverse"], "ApplicationListEntries-ShaperReverse"); ok {
			if err = d.Set("shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading shaper_reverse: %v", err)
		}
	}

	if err = d.Set("sub_category", flattenApplicationListEntriesSubCategory2edl(o["sub-category"], d, "sub_category")); err != nil {
		if vv, ok := fortiAPIPatch(o["sub-category"], "ApplicationListEntries-SubCategory"); ok {
			if err = d.Set("sub_category", vv); err != nil {
				return fmt.Errorf("Error reading sub_category: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sub_category: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationListEntriesTechnology2edl(o["technology"], d, "technology")); err != nil {
		if vv, ok := fortiAPIPatch(o["technology"], "ApplicationListEntries-Technology"); ok {
			if err = d.Set("technology", vv); err != nil {
				return fmt.Errorf("Error reading technology: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationListEntriesVendor2edl(o["vendor"], d, "vendor")); err != nil {
		if vv, ok := fortiAPIPatch(o["vendor"], "ApplicationListEntries-Vendor"); ok {
			if err = d.Set("vendor", vv); err != nil {
				return fmt.Errorf("Error reading vendor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	return nil
}

func flattenApplicationListEntriesFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandApplicationListEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesApplication2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesBehavior2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesExclusion2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLogPacket2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParameters2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandApplicationListEntriesParametersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandApplicationListEntriesParametersMembers2edl(d, i["members"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["members"] = t
			}
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembers2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandApplicationListEntriesParametersMembersId2edl(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandApplicationListEntriesParametersMembersName2edl(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandApplicationListEntriesParametersMembersValue2edl(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersMembersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembersName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembersValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesPerIpShaper2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesPopularity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesProtocols2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateCount2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateMode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateTrack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRisk2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesSessionTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesShaper2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesShaperReverse2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesSubCategory2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesTechnology2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandApplicationListEntriesVendor2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectApplicationListEntries(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandApplicationListEntriesAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandApplicationListEntriesApplication2edl(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("behavior"); ok || d.HasChange("behavior") {
		t, err := expandApplicationListEntriesBehavior2edl(d, v, "behavior")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandApplicationListEntriesCategory2edl(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("exclusion"); ok || d.HasChange("exclusion") {
		t, err := expandApplicationListEntriesExclusion2edl(d, v, "exclusion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusion"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandApplicationListEntriesId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandApplicationListEntriesLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok || d.HasChange("log_packet") {
		t, err := expandApplicationListEntriesLogPacket2edl(d, v, "log_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("parameters"); ok || d.HasChange("parameters") {
		t, err := expandApplicationListEntriesParameters2edl(d, v, "parameters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameters"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok || d.HasChange("per_ip_shaper") {
		t, err := expandApplicationListEntriesPerIpShaper2edl(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("popularity"); ok || d.HasChange("popularity") {
		t, err := expandApplicationListEntriesPopularity2edl(d, v, "popularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	if v, ok := d.GetOk("protocols"); ok || d.HasChange("protocols") {
		t, err := expandApplicationListEntriesProtocols2edl(d, v, "protocols")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocols"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandApplicationListEntriesQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandApplicationListEntriesQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandApplicationListEntriesQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("rate_count"); ok || d.HasChange("rate_count") {
		t, err := expandApplicationListEntriesRateCount2edl(d, v, "rate_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-count"] = t
		}
	}

	if v, ok := d.GetOk("rate_duration"); ok || d.HasChange("rate_duration") {
		t, err := expandApplicationListEntriesRateDuration2edl(d, v, "rate_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-duration"] = t
		}
	}

	if v, ok := d.GetOk("rate_mode"); ok || d.HasChange("rate_mode") {
		t, err := expandApplicationListEntriesRateMode2edl(d, v, "rate_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-mode"] = t
		}
	}

	if v, ok := d.GetOk("rate_track"); ok || d.HasChange("rate_track") {
		t, err := expandApplicationListEntriesRateTrack2edl(d, v, "rate_track")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-track"] = t
		}
	}

	if v, ok := d.GetOk("risk"); ok || d.HasChange("risk") {
		t, err := expandApplicationListEntriesRisk2edl(d, v, "risk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok || d.HasChange("session_ttl") {
		t, err := expandApplicationListEntriesSessionTtl2edl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("shaper"); ok || d.HasChange("shaper") {
		t, err := expandApplicationListEntriesShaper2edl(d, v, "shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaper"] = t
		}
	}

	if v, ok := d.GetOk("shaper_reverse"); ok || d.HasChange("shaper_reverse") {
		t, err := expandApplicationListEntriesShaperReverse2edl(d, v, "shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("sub_category"); ok || d.HasChange("sub_category") {
		t, err := expandApplicationListEntriesSubCategory2edl(d, v, "sub_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-category"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok || d.HasChange("technology") {
		t, err := expandApplicationListEntriesTechnology2edl(d, v, "technology")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok || d.HasChange("vendor") {
		t, err := expandApplicationListEntriesVendor2edl(d, v, "vendor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	return &obj, nil
}
