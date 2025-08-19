// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure logical storage.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStorageCreate,
		Read:   resourceSystemStorageRead,
		Update: resourceSystemStorageUpdate,
		Delete: resourceSystemStorageDelete,

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
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"media_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanopt_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemStorageCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemStorage(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemStorage resource while getting object: %v", err)
	}

	_, err = c.CreateSystemStorage(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemStorage resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemStorageRead(d, m)
}

func resourceSystemStorageUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemStorage(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemStorage resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStorage(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemStorage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemStorageRead(d, m)
}

func resourceSystemStorageDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemStorage(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStorageRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemStorage(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemStorage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStorage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemStorage resource from API: %v", err)
	}
	return nil
}

func flattenSystemStorageDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageMediaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageOrder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStoragePartition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageUsage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStorageWanoptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemStorage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenSystemStorageDevice(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemStorage-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("media_status", flattenSystemStorageMediaStatus(o["media-status"], d, "media_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["media-status"], "SystemStorage-MediaStatus"); ok {
			if err = d.Set("media_status", vv); err != nil {
				return fmt.Errorf("Error reading media_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading media_status: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemStorageName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemStorage-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("order", flattenSystemStorageOrder(o["order"], d, "order")); err != nil {
		if vv, ok := fortiAPIPatch(o["order"], "SystemStorage-Order"); ok {
			if err = d.Set("order", vv); err != nil {
				return fmt.Errorf("Error reading order: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading order: %v", err)
		}
	}

	if err = d.Set("partition", flattenSystemStoragePartition(o["partition"], d, "partition")); err != nil {
		if vv, ok := fortiAPIPatch(o["partition"], "SystemStorage-Partition"); ok {
			if err = d.Set("partition", vv); err != nil {
				return fmt.Errorf("Error reading partition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading partition: %v", err)
		}
	}

	if err = d.Set("size", flattenSystemStorageSize(o["size"], d, "size")); err != nil {
		if vv, ok := fortiAPIPatch(o["size"], "SystemStorage-Size"); ok {
			if err = d.Set("size", vv); err != nil {
				return fmt.Errorf("Error reading size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading size: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemStorageStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemStorage-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("usage", flattenSystemStorageUsage(o["usage"], d, "usage")); err != nil {
		if vv, ok := fortiAPIPatch(o["usage"], "SystemStorage-Usage"); ok {
			if err = d.Set("usage", vv); err != nil {
				return fmt.Errorf("Error reading usage: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading usage: %v", err)
		}
	}

	if err = d.Set("wanopt_mode", flattenSystemStorageWanoptMode(o["wanopt-mode"], d, "wanopt_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanopt-mode"], "SystemStorage-WanoptMode"); ok {
			if err = d.Set("wanopt_mode", vv); err != nil {
				return fmt.Errorf("Error reading wanopt_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanopt_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemStorageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemStorageDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageMediaStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageOrder(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStoragePartition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageUsage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageWanoptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStorage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemStorageDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("media_status"); ok || d.HasChange("media_status") {
		t, err := expandSystemStorageMediaStatus(d, v, "media_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["media-status"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemStorageName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("order"); ok || d.HasChange("order") {
		t, err := expandSystemStorageOrder(d, v, "order")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["order"] = t
		}
	}

	if v, ok := d.GetOk("partition"); ok || d.HasChange("partition") {
		t, err := expandSystemStoragePartition(d, v, "partition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partition"] = t
		}
	}

	if v, ok := d.GetOk("size"); ok || d.HasChange("size") {
		t, err := expandSystemStorageSize(d, v, "size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["size"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemStorageStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("usage"); ok || d.HasChange("usage") {
		t, err := expandSystemStorageUsage(d, v, "usage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usage"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_mode"); ok || d.HasChange("wanopt_mode") {
		t, err := expandSystemStorageWanoptMode(d, v, "wanopt_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-mode"] = t
		}
	}

	return &obj, nil
}
