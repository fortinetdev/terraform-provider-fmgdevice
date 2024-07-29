// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure CPUs enabled to run engines in each DPDK stage.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDpdkCpus() *schema.Resource {
	return &schema.Resource{
		Create: resourceDpdkCpusUpdate,
		Read:   resourceDpdkCpusRead,
		Update: resourceDpdkCpusUpdate,
		Delete: resourceDpdkCpusDelete,

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
			"ips_cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"isolated_cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rx_cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tx_cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnp_cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnpsp_cpus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDpdkCpusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectDpdkCpus(d)
	if err != nil {
		return fmt.Errorf("Error updating DpdkCpus resource while getting object: %v", err)
	}

	_, err = c.UpdateDpdkCpus(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating DpdkCpus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DpdkCpus")

	return resourceDpdkCpusRead(d, m)
}

func resourceDpdkCpusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteDpdkCpus(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting DpdkCpus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDpdkCpusRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDpdkCpus(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading DpdkCpus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDpdkCpus(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DpdkCpus resource from API: %v", err)
	}
	return nil
}

func flattenDpdkCpusIpsCpus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkCpusIsolatedCpus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkCpusRxCpus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkCpusTxCpus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkCpusVnpCpus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDpdkCpusVnpspCpus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDpdkCpus(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ips_cpus", flattenDpdkCpusIpsCpus(o["ips-cpus"], d, "ips_cpus")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-cpus"], "DpdkCpus-IpsCpus"); ok {
			if err = d.Set("ips_cpus", vv); err != nil {
				return fmt.Errorf("Error reading ips_cpus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_cpus: %v", err)
		}
	}

	if err = d.Set("isolated_cpus", flattenDpdkCpusIsolatedCpus(o["isolated-cpus"], d, "isolated_cpus")); err != nil {
		if vv, ok := fortiAPIPatch(o["isolated-cpus"], "DpdkCpus-IsolatedCpus"); ok {
			if err = d.Set("isolated_cpus", vv); err != nil {
				return fmt.Errorf("Error reading isolated_cpus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isolated_cpus: %v", err)
		}
	}

	if err = d.Set("rx_cpus", flattenDpdkCpusRxCpus(o["rx-cpus"], d, "rx_cpus")); err != nil {
		if vv, ok := fortiAPIPatch(o["rx-cpus"], "DpdkCpus-RxCpus"); ok {
			if err = d.Set("rx_cpus", vv); err != nil {
				return fmt.Errorf("Error reading rx_cpus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rx_cpus: %v", err)
		}
	}

	if err = d.Set("tx_cpus", flattenDpdkCpusTxCpus(o["tx-cpus"], d, "tx_cpus")); err != nil {
		if vv, ok := fortiAPIPatch(o["tx-cpus"], "DpdkCpus-TxCpus"); ok {
			if err = d.Set("tx_cpus", vv); err != nil {
				return fmt.Errorf("Error reading tx_cpus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tx_cpus: %v", err)
		}
	}

	if err = d.Set("vnp_cpus", flattenDpdkCpusVnpCpus(o["vnp-cpus"], d, "vnp_cpus")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnp-cpus"], "DpdkCpus-VnpCpus"); ok {
			if err = d.Set("vnp_cpus", vv); err != nil {
				return fmt.Errorf("Error reading vnp_cpus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnp_cpus: %v", err)
		}
	}

	if err = d.Set("vnpsp_cpus", flattenDpdkCpusVnpspCpus(o["vnpsp-cpus"], d, "vnpsp_cpus")); err != nil {
		if vv, ok := fortiAPIPatch(o["vnpsp-cpus"], "DpdkCpus-VnpspCpus"); ok {
			if err = d.Set("vnpsp_cpus", vv); err != nil {
				return fmt.Errorf("Error reading vnpsp_cpus: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vnpsp_cpus: %v", err)
		}
	}

	return nil
}

func flattenDpdkCpusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDpdkCpusIpsCpus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusIsolatedCpus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusRxCpus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusTxCpus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusVnpCpus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusVnpspCpus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDpdkCpus(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ips_cpus"); ok || d.HasChange("ips_cpus") {
		t, err := expandDpdkCpusIpsCpus(d, v, "ips_cpus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-cpus"] = t
		}
	}

	if v, ok := d.GetOk("isolated_cpus"); ok || d.HasChange("isolated_cpus") {
		t, err := expandDpdkCpusIsolatedCpus(d, v, "isolated_cpus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isolated-cpus"] = t
		}
	}

	if v, ok := d.GetOk("rx_cpus"); ok || d.HasChange("rx_cpus") {
		t, err := expandDpdkCpusRxCpus(d, v, "rx_cpus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rx-cpus"] = t
		}
	}

	if v, ok := d.GetOk("tx_cpus"); ok || d.HasChange("tx_cpus") {
		t, err := expandDpdkCpusTxCpus(d, v, "tx_cpus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-cpus"] = t
		}
	}

	if v, ok := d.GetOk("vnp_cpus"); ok || d.HasChange("vnp_cpus") {
		t, err := expandDpdkCpusVnpCpus(d, v, "vnp_cpus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnp-cpus"] = t
		}
	}

	if v, ok := d.GetOk("vnpsp_cpus"); ok || d.HasChange("vnpsp_cpus") {
		t, err := expandDpdkCpusVnpspCpus(d, v, "vnpsp_cpus")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vnpsp-cpus"] = t
		}
	}

	return &obj, nil
}
