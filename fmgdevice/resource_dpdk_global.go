// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure global DPDK options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDpdkGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceDpdkGlobalUpdate,
		Read:   resourceDpdkGlobalRead,
		Update: resourceDpdkGlobalUpdate,
		Delete: resourceDpdkGlobalDelete,

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
			"elasticbuffer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hugepage_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ipsec_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbufpool_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"multiqueue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_session_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protects": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_table_percentage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sleep_on_idle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDpdkGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectDpdkGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating DpdkGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateDpdkGlobal(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DpdkGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DpdkGlobal")

	return resourceDpdkGlobalRead(d, m)
}

func resourceDpdkGlobalDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteDpdkGlobal(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DpdkGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDpdkGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDpdkGlobal(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading DpdkGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDpdkGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DpdkGlobal resource from API: %v", err)
	}
	return nil
}

func flattenDpdkGlobalElasticbuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalHugepagePercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDpdkGlobalIpsecOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalMbufpoolPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalMultiqueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalPerSessionAccounting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalProtects(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalSessionTablePercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalSleepOnIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkGlobalStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDpdkGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("elasticbuffer", flattenDpdkGlobalElasticbuffer(o["elasticbuffer"], d, "elasticbuffer")); err != nil {
		if vv, ok := fortiAPIPatch(o["elasticbuffer"], "DpdkGlobal-Elasticbuffer"); ok {
			if err = d.Set("elasticbuffer", vv); err != nil {
				return fmt.Errorf("Error reading elasticbuffer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading elasticbuffer: %v", err)
		}
	}

	if err = d.Set("hugepage_percentage", flattenDpdkGlobalHugepagePercentage(o["hugepage-percentage"], d, "hugepage_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["hugepage-percentage"], "DpdkGlobal-HugepagePercentage"); ok {
			if err = d.Set("hugepage_percentage", vv); err != nil {
				return fmt.Errorf("Error reading hugepage_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hugepage_percentage: %v", err)
		}
	}

	if err = d.Set("interface", flattenDpdkGlobalInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "DpdkGlobal-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipsec_offload", flattenDpdkGlobalIpsecOffload(o["ipsec-offload"], d, "ipsec_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipsec-offload"], "DpdkGlobal-IpsecOffload"); ok {
			if err = d.Set("ipsec_offload", vv); err != nil {
				return fmt.Errorf("Error reading ipsec_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipsec_offload: %v", err)
		}
	}

	if err = d.Set("mbufpool_percentage", flattenDpdkGlobalMbufpoolPercentage(o["mbufpool-percentage"], d, "mbufpool_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["mbufpool-percentage"], "DpdkGlobal-MbufpoolPercentage"); ok {
			if err = d.Set("mbufpool_percentage", vv); err != nil {
				return fmt.Errorf("Error reading mbufpool_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mbufpool_percentage: %v", err)
		}
	}

	if err = d.Set("multiqueue", flattenDpdkGlobalMultiqueue(o["multiqueue"], d, "multiqueue")); err != nil {
		if vv, ok := fortiAPIPatch(o["multiqueue"], "DpdkGlobal-Multiqueue"); ok {
			if err = d.Set("multiqueue", vv); err != nil {
				return fmt.Errorf("Error reading multiqueue: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading multiqueue: %v", err)
		}
	}

	if err = d.Set("per_session_accounting", flattenDpdkGlobalPerSessionAccounting(o["per-session-accounting"], d, "per_session_accounting")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-session-accounting"], "DpdkGlobal-PerSessionAccounting"); ok {
			if err = d.Set("per_session_accounting", vv); err != nil {
				return fmt.Errorf("Error reading per_session_accounting: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_session_accounting: %v", err)
		}
	}

	if err = d.Set("protects", flattenDpdkGlobalProtects(o["protects"], d, "protects")); err != nil {
		if vv, ok := fortiAPIPatch(o["protects"], "DpdkGlobal-Protects"); ok {
			if err = d.Set("protects", vv); err != nil {
				return fmt.Errorf("Error reading protects: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protects: %v", err)
		}
	}

	if err = d.Set("session_table_percentage", flattenDpdkGlobalSessionTablePercentage(o["session-table-percentage"], d, "session_table_percentage")); err != nil {
		if vv, ok := fortiAPIPatch(o["session-table-percentage"], "DpdkGlobal-SessionTablePercentage"); ok {
			if err = d.Set("session_table_percentage", vv); err != nil {
				return fmt.Errorf("Error reading session_table_percentage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading session_table_percentage: %v", err)
		}
	}

	if err = d.Set("sleep_on_idle", flattenDpdkGlobalSleepOnIdle(o["sleep-on-idle"], d, "sleep_on_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["sleep-on-idle"], "DpdkGlobal-SleepOnIdle"); ok {
			if err = d.Set("sleep_on_idle", vv); err != nil {
				return fmt.Errorf("Error reading sleep_on_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sleep_on_idle: %v", err)
		}
	}

	if err = d.Set("status", flattenDpdkGlobalStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "DpdkGlobal-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenDpdkGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDpdkGlobalElasticbuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalHugepagePercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDpdkGlobalIpsecOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalMbufpoolPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalMultiqueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalPerSessionAccounting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalProtects(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalSessionTablePercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalSleepOnIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkGlobalStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDpdkGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("elasticbuffer"); ok || d.HasChange("elasticbuffer") {
		t, err := expandDpdkGlobalElasticbuffer(d, v, "elasticbuffer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["elasticbuffer"] = t
		}
	}

	if v, ok := d.GetOk("hugepage_percentage"); ok || d.HasChange("hugepage_percentage") {
		t, err := expandDpdkGlobalHugepagePercentage(d, v, "hugepage_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hugepage-percentage"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandDpdkGlobalInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_offload"); ok || d.HasChange("ipsec_offload") {
		t, err := expandDpdkGlobalIpsecOffload(d, v, "ipsec_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-offload"] = t
		}
	}

	if v, ok := d.GetOk("mbufpool_percentage"); ok || d.HasChange("mbufpool_percentage") {
		t, err := expandDpdkGlobalMbufpoolPercentage(d, v, "mbufpool_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbufpool-percentage"] = t
		}
	}

	if v, ok := d.GetOk("multiqueue"); ok || d.HasChange("multiqueue") {
		t, err := expandDpdkGlobalMultiqueue(d, v, "multiqueue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiqueue"] = t
		}
	}

	if v, ok := d.GetOk("per_session_accounting"); ok || d.HasChange("per_session_accounting") {
		t, err := expandDpdkGlobalPerSessionAccounting(d, v, "per_session_accounting")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-session-accounting"] = t
		}
	}

	if v, ok := d.GetOk("protects"); ok || d.HasChange("protects") {
		t, err := expandDpdkGlobalProtects(d, v, "protects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protects"] = t
		}
	}

	if v, ok := d.GetOk("session_table_percentage"); ok || d.HasChange("session_table_percentage") {
		t, err := expandDpdkGlobalSessionTablePercentage(d, v, "session_table_percentage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-table-percentage"] = t
		}
	}

	if v, ok := d.GetOk("sleep_on_idle"); ok || d.HasChange("sleep_on_idle") {
		t, err := expandDpdkGlobalSleepOnIdle(d, v, "sleep_on_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sleep-on-idle"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandDpdkGlobalStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
