// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Exclusion filters

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNetflowExclusionFilters() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNetflowExclusionFiltersCreate,
		Read:   resourceSystemNetflowExclusionFiltersRead,
		Update: resourceSystemNetflowExclusionFiltersUpdate,
		Delete: resourceSystemNetflowExclusionFiltersDelete,

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
			"destination_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemNetflowExclusionFiltersCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemNetflowExclusionFilters(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemNetflowExclusionFilters resource while getting object: %v", err)
	}

	v, err := c.CreateSystemNetflowExclusionFilters(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemNetflowExclusionFilters resource: %v", err)
	}

	if v != nil && v["id"] != nil {
		if vidn, ok := v["id"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourceSystemNetflowExclusionFiltersRead(d, m)
		} else {
			return fmt.Errorf("Error creating SystemNetflowExclusionFilters resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemNetflowExclusionFiltersRead(d, m)
}

func resourceSystemNetflowExclusionFiltersUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemNetflowExclusionFilters(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetflowExclusionFilters resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNetflowExclusionFilters(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetflowExclusionFilters resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemNetflowExclusionFiltersRead(d, m)
}

func resourceSystemNetflowExclusionFiltersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemNetflowExclusionFilters(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNetflowExclusionFilters resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNetflowExclusionFiltersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNetflowExclusionFilters(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetflowExclusionFilters resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNetflowExclusionFilters(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetflowExclusionFilters resource from API: %v", err)
	}
	return nil
}

func flattenSystemNetflowExclusionFiltersDestinationIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetflowExclusionFiltersDestinationPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetflowExclusionFiltersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetflowExclusionFiltersProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetflowExclusionFiltersSourceIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNetflowExclusionFiltersSourcePort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemNetflowExclusionFilters(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("destination_ip", flattenSystemNetflowExclusionFiltersDestinationIp2edl(o["destination-ip"], d, "destination_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-ip"], "SystemNetflowExclusionFilters-DestinationIp"); ok {
			if err = d.Set("destination_ip", vv); err != nil {
				return fmt.Errorf("Error reading destination_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_ip: %v", err)
		}
	}

	if err = d.Set("destination_port", flattenSystemNetflowExclusionFiltersDestinationPort2edl(o["destination-port"], d, "destination_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["destination-port"], "SystemNetflowExclusionFilters-DestinationPort"); ok {
			if err = d.Set("destination_port", vv); err != nil {
				return fmt.Errorf("Error reading destination_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading destination_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemNetflowExclusionFiltersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemNetflowExclusionFilters-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemNetflowExclusionFiltersProtocol2edl(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemNetflowExclusionFilters-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemNetflowExclusionFiltersSourceIp2edl(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "SystemNetflowExclusionFilters-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_port", flattenSystemNetflowExclusionFiltersSourcePort2edl(o["source-port"], d, "source_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-port"], "SystemNetflowExclusionFilters-SourcePort"); ok {
			if err = d.Set("source_port", vv); err != nil {
				return fmt.Errorf("Error reading source_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	return nil
}

func flattenSystemNetflowExclusionFiltersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemNetflowExclusionFiltersDestinationIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowExclusionFiltersDestinationPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowExclusionFiltersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowExclusionFiltersProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowExclusionFiltersSourceIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowExclusionFiltersSourcePort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNetflowExclusionFilters(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("destination_ip"); ok || d.HasChange("destination_ip") {
		t, err := expandSystemNetflowExclusionFiltersDestinationIp2edl(d, v, "destination_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-ip"] = t
		}
	}

	if v, ok := d.GetOk("destination_port"); ok || d.HasChange("destination_port") {
		t, err := expandSystemNetflowExclusionFiltersDestinationPort2edl(d, v, "destination_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemNetflowExclusionFiltersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemNetflowExclusionFiltersProtocol2edl(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandSystemNetflowExclusionFiltersSourceIp2edl(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_port"); ok || d.HasChange("source_port") {
		t, err := expandSystemNetflowExclusionFiltersSourcePort2edl(d, v, "source_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	}

	return &obj, nil
}
