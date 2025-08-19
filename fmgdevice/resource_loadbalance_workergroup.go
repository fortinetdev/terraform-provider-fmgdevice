// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Worker group configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLoadBalanceWorkerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoadBalanceWorkerGroupCreate,
		Read:   resourceLoadBalanceWorkerGroupRead,
		Update: resourceLoadBalanceWorkerGroupUpdate,
		Delete: resourceLoadBalanceWorkerGroupDelete,

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
			"member": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"worker_group_name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceLoadBalanceWorkerGroupCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLoadBalanceWorkerGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWorkerGroup resource while getting object: %v", err)
	}

	_, err = c.CreateLoadBalanceWorkerGroup(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating LoadBalanceWorkerGroup resource: %v", err)
	}

	d.SetId(getStringKey(d, "worker_group_name"))

	return resourceLoadBalanceWorkerGroupRead(d, m)
}

func resourceLoadBalanceWorkerGroupUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectLoadBalanceWorkerGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWorkerGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateLoadBalanceWorkerGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating LoadBalanceWorkerGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "worker_group_name"))

	return resourceLoadBalanceWorkerGroupRead(d, m)
}

func resourceLoadBalanceWorkerGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteLoadBalanceWorkerGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting LoadBalanceWorkerGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLoadBalanceWorkerGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLoadBalanceWorkerGroup(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWorkerGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLoadBalanceWorkerGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LoadBalanceWorkerGroup resource from API: %v", err)
	}
	return nil
}

func flattenLoadBalanceWorkerGroupMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenLoadBalanceWorkerGroupWorkerGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLoadBalanceWorkerGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("member", flattenLoadBalanceWorkerGroupMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "LoadBalanceWorkerGroup-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("worker_group_name", flattenLoadBalanceWorkerGroupWorkerGroupName(o["worker-group-name"], d, "worker_group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["worker-group-name"], "LoadBalanceWorkerGroup-WorkerGroupName"); ok {
			if err = d.Set("worker_group_name", vv); err != nil {
				return fmt.Errorf("Error reading worker_group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading worker_group_name: %v", err)
		}
	}

	return nil
}

func flattenLoadBalanceWorkerGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLoadBalanceWorkerGroupMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandLoadBalanceWorkerGroupWorkerGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLoadBalanceWorkerGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandLoadBalanceWorkerGroupMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("worker_group_name"); ok || d.HasChange("worker_group_name") {
		t, err := expandLoadBalanceWorkerGroupWorkerGroupName(d, v, "worker_group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["worker-group-name"] = t
		}
	}

	return &obj, nil
}
