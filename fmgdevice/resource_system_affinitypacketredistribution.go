// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure packet redistribution.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAffinityPacketRedistribution() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAffinityPacketRedistributionCreate,
		Read:   resourceSystemAffinityPacketRedistributionRead,
		Update: resourceSystemAffinityPacketRedistributionUpdate,
		Delete: resourceSystemAffinityPacketRedistributionDelete,

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
			"affinity_cpumask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"round_robin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rxqid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemAffinityPacketRedistributionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAffinityPacketRedistribution(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityPacketRedistribution resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAffinityPacketRedistribution(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityPacketRedistribution resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAffinityPacketRedistributionRead(d, m)
}

func resourceSystemAffinityPacketRedistributionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAffinityPacketRedistribution(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityPacketRedistribution resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAffinityPacketRedistribution(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityPacketRedistribution resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemAffinityPacketRedistributionRead(d, m)
}

func resourceSystemAffinityPacketRedistributionDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAffinityPacketRedistribution(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAffinityPacketRedistribution resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAffinityPacketRedistributionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAffinityPacketRedistribution(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityPacketRedistribution resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAffinityPacketRedistribution(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityPacketRedistribution resource from API: %v", err)
	}
	return nil
}

func flattenSystemAffinityPacketRedistributionAffinityCpumask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityPacketRedistributionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityPacketRedistributionInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAffinityPacketRedistributionRoundRobin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityPacketRedistributionRxqid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAffinityPacketRedistribution(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("affinity_cpumask", flattenSystemAffinityPacketRedistributionAffinityCpumask(o["affinity-cpumask"], d, "affinity_cpumask")); err != nil {
		if vv, ok := fortiAPIPatch(o["affinity-cpumask"], "SystemAffinityPacketRedistribution-AffinityCpumask"); ok {
			if err = d.Set("affinity_cpumask", vv); err != nil {
				return fmt.Errorf("Error reading affinity_cpumask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading affinity_cpumask: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemAffinityPacketRedistributionId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemAffinityPacketRedistribution-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemAffinityPacketRedistributionInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemAffinityPacketRedistribution-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("round_robin", flattenSystemAffinityPacketRedistributionRoundRobin(o["round-robin"], d, "round_robin")); err != nil {
		if vv, ok := fortiAPIPatch(o["round-robin"], "SystemAffinityPacketRedistribution-RoundRobin"); ok {
			if err = d.Set("round_robin", vv); err != nil {
				return fmt.Errorf("Error reading round_robin: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading round_robin: %v", err)
		}
	}

	if err = d.Set("rxqid", flattenSystemAffinityPacketRedistributionRxqid(o["rxqid"], d, "rxqid")); err != nil {
		if vv, ok := fortiAPIPatch(o["rxqid"], "SystemAffinityPacketRedistribution-Rxqid"); ok {
			if err = d.Set("rxqid", vv); err != nil {
				return fmt.Errorf("Error reading rxqid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rxqid: %v", err)
		}
	}

	return nil
}

func flattenSystemAffinityPacketRedistributionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAffinityPacketRedistributionAffinityCpumask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityPacketRedistributionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityPacketRedistributionInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAffinityPacketRedistributionRoundRobin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityPacketRedistributionRxqid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAffinityPacketRedistribution(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("affinity_cpumask"); ok || d.HasChange("affinity_cpumask") {
		t, err := expandSystemAffinityPacketRedistributionAffinityCpumask(d, v, "affinity_cpumask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["affinity-cpumask"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemAffinityPacketRedistributionId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemAffinityPacketRedistributionInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("round_robin"); ok || d.HasChange("round_robin") {
		t, err := expandSystemAffinityPacketRedistributionRoundRobin(d, v, "round_robin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["round-robin"] = t
		}
	}

	if v, ok := d.GetOk("rxqid"); ok || d.HasChange("rxqid") {
		t, err := expandSystemAffinityPacketRedistributionRxqid(d, v, "rxqid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rxqid"] = t
		}
	}

	return &obj, nil
}
