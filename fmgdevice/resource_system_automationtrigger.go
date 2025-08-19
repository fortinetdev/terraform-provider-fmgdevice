// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Trigger for automation stitches.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationTriggerCreate,
		Read:   resourceSystemAutomationTriggerRead,
		Update: resourceSystemAutomationTriggerUpdate,
		Delete: resourceSystemAutomationTriggerDelete,

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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_event_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_event_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"faz_event_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"faz_event_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"faz_event_tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": &schema.Schema{
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
			"ioc_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"report_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stitch_name": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trigger_datetime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_day": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"trigger_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_hour": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"trigger_minute": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"trigger_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
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

func resourceSystemAutomationTriggerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationTrigger(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAutomationTrigger(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemAutomationTrigger(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutomationTrigger(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemAutomationTrigger(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationTrigger resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationTriggerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutomationTrigger(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationTrigger(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationTriggerDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerEventType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFabricEventName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFabricEventSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFazEventName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFazEventSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFazEventTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFields(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemAutomationTriggerFieldsId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemAutomationTrigger-Fields-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemAutomationTriggerFieldsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemAutomationTrigger-Fields-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenSystemAutomationTriggerFieldsValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "SystemAutomationTrigger-Fields-Value")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemAutomationTriggerFieldsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerIocLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLicenseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLogid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemAutomationTriggerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerReportType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerStitchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAutomationTriggerTriggerDatetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerMinute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationTriggerVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemAutomationTrigger(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("description", flattenSystemAutomationTriggerDescription(o["description"], d, "description")); err != nil {
		if vv, ok := fortiAPIPatch(o["description"], "SystemAutomationTrigger-Description"); ok {
			if err = d.Set("description", vv); err != nil {
				return fmt.Errorf("Error reading description: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("event_type", flattenSystemAutomationTriggerEventType(o["event-type"], d, "event_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["event-type"], "SystemAutomationTrigger-EventType"); ok {
			if err = d.Set("event_type", vv); err != nil {
				return fmt.Errorf("Error reading event_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event_type: %v", err)
		}
	}

	if err = d.Set("fabric_event_name", flattenSystemAutomationTriggerFabricEventName(o["fabric-event-name"], d, "fabric_event_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-event-name"], "SystemAutomationTrigger-FabricEventName"); ok {
			if err = d.Set("fabric_event_name", vv); err != nil {
				return fmt.Errorf("Error reading fabric_event_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_event_name: %v", err)
		}
	}

	if err = d.Set("fabric_event_severity", flattenSystemAutomationTriggerFabricEventSeverity(o["fabric-event-severity"], d, "fabric_event_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["fabric-event-severity"], "SystemAutomationTrigger-FabricEventSeverity"); ok {
			if err = d.Set("fabric_event_severity", vv); err != nil {
				return fmt.Errorf("Error reading fabric_event_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fabric_event_severity: %v", err)
		}
	}

	if err = d.Set("faz_event_name", flattenSystemAutomationTriggerFazEventName(o["faz-event-name"], d, "faz_event_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-event-name"], "SystemAutomationTrigger-FazEventName"); ok {
			if err = d.Set("faz_event_name", vv); err != nil {
				return fmt.Errorf("Error reading faz_event_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_event_name: %v", err)
		}
	}

	if err = d.Set("faz_event_severity", flattenSystemAutomationTriggerFazEventSeverity(o["faz-event-severity"], d, "faz_event_severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-event-severity"], "SystemAutomationTrigger-FazEventSeverity"); ok {
			if err = d.Set("faz_event_severity", vv); err != nil {
				return fmt.Errorf("Error reading faz_event_severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_event_severity: %v", err)
		}
	}

	if err = d.Set("faz_event_tags", flattenSystemAutomationTriggerFazEventTags(o["faz-event-tags"], d, "faz_event_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz-event-tags"], "SystemAutomationTrigger-FazEventTags"); ok {
			if err = d.Set("faz_event_tags", vv); err != nil {
				return fmt.Errorf("Error reading faz_event_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz_event_tags: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fields", flattenSystemAutomationTriggerFields(o["fields"], d, "fields")); err != nil {
			if vv, ok := fortiAPIPatch(o["fields"], "SystemAutomationTrigger-Fields"); ok {
				if err = d.Set("fields", vv); err != nil {
					return fmt.Errorf("Error reading fields: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fields"); ok {
			if err = d.Set("fields", flattenSystemAutomationTriggerFields(o["fields"], d, "fields")); err != nil {
				if vv, ok := fortiAPIPatch(o["fields"], "SystemAutomationTrigger-Fields"); ok {
					if err = d.Set("fields", vv); err != nil {
						return fmt.Errorf("Error reading fields: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("ioc_level", flattenSystemAutomationTriggerIocLevel(o["ioc-level"], d, "ioc_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["ioc-level"], "SystemAutomationTrigger-IocLevel"); ok {
			if err = d.Set("ioc_level", vv); err != nil {
				return fmt.Errorf("Error reading ioc_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ioc_level: %v", err)
		}
	}

	if err = d.Set("license_type", flattenSystemAutomationTriggerLicenseType(o["license-type"], d, "license_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["license-type"], "SystemAutomationTrigger-LicenseType"); ok {
			if err = d.Set("license_type", vv); err != nil {
				return fmt.Errorf("Error reading license_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading license_type: %v", err)
		}
	}

	if err = d.Set("logid", flattenSystemAutomationTriggerLogid(o["logid"], d, "logid")); err != nil {
		if vv, ok := fortiAPIPatch(o["logid"], "SystemAutomationTrigger-Logid"); ok {
			if err = d.Set("logid", vv); err != nil {
				return fmt.Errorf("Error reading logid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logid: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAutomationTriggerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAutomationTrigger-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("report_type", flattenSystemAutomationTriggerReportType(o["report-type"], d, "report_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["report-type"], "SystemAutomationTrigger-ReportType"); ok {
			if err = d.Set("report_type", vv); err != nil {
				return fmt.Errorf("Error reading report_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report_type: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemAutomationTriggerSerial(o["serial"], d, "serial")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial"], "SystemAutomationTrigger-Serial"); ok {
			if err = d.Set("serial", vv); err != nil {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("stitch_name", flattenSystemAutomationTriggerStitchName(o["stitch-name"], d, "stitch_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["stitch-name"], "SystemAutomationTrigger-StitchName"); ok {
			if err = d.Set("stitch_name", vv); err != nil {
				return fmt.Errorf("Error reading stitch_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stitch_name: %v", err)
		}
	}

	if err = d.Set("trigger_datetime", flattenSystemAutomationTriggerTriggerDatetime(o["trigger-datetime"], d, "trigger_datetime")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-datetime"], "SystemAutomationTrigger-TriggerDatetime"); ok {
			if err = d.Set("trigger_datetime", vv); err != nil {
				return fmt.Errorf("Error reading trigger_datetime: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_datetime: %v", err)
		}
	}

	if err = d.Set("trigger_day", flattenSystemAutomationTriggerTriggerDay(o["trigger-day"], d, "trigger_day")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-day"], "SystemAutomationTrigger-TriggerDay"); ok {
			if err = d.Set("trigger_day", vv); err != nil {
				return fmt.Errorf("Error reading trigger_day: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_day: %v", err)
		}
	}

	if err = d.Set("trigger_frequency", flattenSystemAutomationTriggerTriggerFrequency(o["trigger-frequency"], d, "trigger_frequency")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-frequency"], "SystemAutomationTrigger-TriggerFrequency"); ok {
			if err = d.Set("trigger_frequency", vv); err != nil {
				return fmt.Errorf("Error reading trigger_frequency: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_frequency: %v", err)
		}
	}

	if err = d.Set("trigger_hour", flattenSystemAutomationTriggerTriggerHour(o["trigger-hour"], d, "trigger_hour")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-hour"], "SystemAutomationTrigger-TriggerHour"); ok {
			if err = d.Set("trigger_hour", vv); err != nil {
				return fmt.Errorf("Error reading trigger_hour: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_hour: %v", err)
		}
	}

	if err = d.Set("trigger_minute", flattenSystemAutomationTriggerTriggerMinute(o["trigger-minute"], d, "trigger_minute")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-minute"], "SystemAutomationTrigger-TriggerMinute"); ok {
			if err = d.Set("trigger_minute", vv); err != nil {
				return fmt.Errorf("Error reading trigger_minute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_minute: %v", err)
		}
	}

	if err = d.Set("trigger_type", flattenSystemAutomationTriggerTriggerType(o["trigger-type"], d, "trigger_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-type"], "SystemAutomationTrigger-TriggerType"); ok {
			if err = d.Set("trigger_type", vv); err != nil {
				return fmt.Errorf("Error reading trigger_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_type: %v", err)
		}
	}

	if err = d.Set("trigger_weekday", flattenSystemAutomationTriggerTriggerWeekday(o["trigger-weekday"], d, "trigger_weekday")); err != nil {
		if vv, ok := fortiAPIPatch(o["trigger-weekday"], "SystemAutomationTrigger-TriggerWeekday"); ok {
			if err = d.Set("trigger_weekday", vv); err != nil {
				return fmt.Errorf("Error reading trigger_weekday: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trigger_weekday: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemAutomationTriggerVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemAutomationTrigger-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationTriggerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationTriggerDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerEventType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFabricEventName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFabricEventSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFazEventName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFazEventSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFazEventTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAutomationTriggerFieldsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemAutomationTriggerFieldsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["value"], _ = expandSystemAutomationTriggerFieldsValue(d, i["value"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemAutomationTriggerFieldsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerIocLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLicenseType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLogid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationTriggerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerReportType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerStitchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAutomationTriggerTriggerDatetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerMinute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerWeekday(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemAutomationTrigger(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		t, err := expandSystemAutomationTriggerDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("event_type"); ok || d.HasChange("event_type") {
		t, err := expandSystemAutomationTriggerEventType(d, v, "event_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-type"] = t
		}
	}

	if v, ok := d.GetOk("fabric_event_name"); ok || d.HasChange("fabric_event_name") {
		t, err := expandSystemAutomationTriggerFabricEventName(d, v, "fabric_event_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-event-name"] = t
		}
	}

	if v, ok := d.GetOk("fabric_event_severity"); ok || d.HasChange("fabric_event_severity") {
		t, err := expandSystemAutomationTriggerFabricEventSeverity(d, v, "fabric_event_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-event-severity"] = t
		}
	}

	if v, ok := d.GetOk("faz_event_name"); ok || d.HasChange("faz_event_name") {
		t, err := expandSystemAutomationTriggerFazEventName(d, v, "faz_event_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-event-name"] = t
		}
	}

	if v, ok := d.GetOk("faz_event_severity"); ok || d.HasChange("faz_event_severity") {
		t, err := expandSystemAutomationTriggerFazEventSeverity(d, v, "faz_event_severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-event-severity"] = t
		}
	}

	if v, ok := d.GetOk("faz_event_tags"); ok || d.HasChange("faz_event_tags") {
		t, err := expandSystemAutomationTriggerFazEventTags(d, v, "faz_event_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-event-tags"] = t
		}
	}

	if v, ok := d.GetOk("fields"); ok || d.HasChange("fields") {
		t, err := expandSystemAutomationTriggerFields(d, v, "fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fields"] = t
		}
	}

	if v, ok := d.GetOk("ioc_level"); ok || d.HasChange("ioc_level") {
		t, err := expandSystemAutomationTriggerIocLevel(d, v, "ioc_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ioc-level"] = t
		}
	}

	if v, ok := d.GetOk("license_type"); ok || d.HasChange("license_type") {
		t, err := expandSystemAutomationTriggerLicenseType(d, v, "license_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license-type"] = t
		}
	}

	if v, ok := d.GetOk("logid"); ok || d.HasChange("logid") {
		t, err := expandSystemAutomationTriggerLogid(d, v, "logid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAutomationTriggerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("report_type"); ok || d.HasChange("report_type") {
		t, err := expandSystemAutomationTriggerReportType(d, v, "report_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-type"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok || d.HasChange("serial") {
		t, err := expandSystemAutomationTriggerSerial(d, v, "serial")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("stitch_name"); ok || d.HasChange("stitch_name") {
		t, err := expandSystemAutomationTriggerStitchName(d, v, "stitch_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stitch-name"] = t
		}
	}

	if v, ok := d.GetOk("trigger_datetime"); ok || d.HasChange("trigger_datetime") {
		t, err := expandSystemAutomationTriggerTriggerDatetime(d, v, "trigger_datetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-datetime"] = t
		}
	}

	if v, ok := d.GetOk("trigger_day"); ok || d.HasChange("trigger_day") {
		t, err := expandSystemAutomationTriggerTriggerDay(d, v, "trigger_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-day"] = t
		}
	}

	if v, ok := d.GetOk("trigger_frequency"); ok || d.HasChange("trigger_frequency") {
		t, err := expandSystemAutomationTriggerTriggerFrequency(d, v, "trigger_frequency")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-frequency"] = t
		}
	}

	if v, ok := d.GetOk("trigger_hour"); ok || d.HasChange("trigger_hour") {
		t, err := expandSystemAutomationTriggerTriggerHour(d, v, "trigger_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-hour"] = t
		}
	}

	if v, ok := d.GetOk("trigger_minute"); ok || d.HasChange("trigger_minute") {
		t, err := expandSystemAutomationTriggerTriggerMinute(d, v, "trigger_minute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-minute"] = t
		}
	}

	if v, ok := d.GetOk("trigger_type"); ok || d.HasChange("trigger_type") {
		t, err := expandSystemAutomationTriggerTriggerType(d, v, "trigger_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-type"] = t
		}
	}

	if v, ok := d.GetOk("trigger_weekday"); ok || d.HasChange("trigger_weekday") {
		t, err := expandSystemAutomationTriggerTriggerWeekday(d, v, "trigger_weekday")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-weekday"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemAutomationTriggerVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
