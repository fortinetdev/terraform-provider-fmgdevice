// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure a remote cache device as Web cache storage.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptRemoteStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptRemoteStorageUpdate,
		Read:   resourceWanoptRemoteStorageRead,
		Update: resourceWanoptRemoteStorageUpdate,
		Delete: resourceWanoptRemoteStorageDelete,

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
			"local_cache_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_cache_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_cache_ip": &schema.Schema{
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

func resourceWanoptRemoteStorageUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptRemoteStorage(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptRemoteStorage resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptRemoteStorage(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WanoptRemoteStorage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptRemoteStorage")

	return resourceWanoptRemoteStorageRead(d, m)
}

func resourceWanoptRemoteStorageDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptRemoteStorage(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptRemoteStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptRemoteStorageRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptRemoteStorage(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptRemoteStorage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptRemoteStorage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptRemoteStorage resource from API: %v", err)
	}
	return nil
}

func flattenWanoptRemoteStorageLocalCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptRemoteStorageRemoteCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptRemoteStorageRemoteCacheIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptRemoteStorageStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptRemoteStorage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("local_cache_id", flattenWanoptRemoteStorageLocalCacheId(o["local-cache-id"], d, "local_cache_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cache-id"], "WanoptRemoteStorage-LocalCacheId"); ok {
			if err = d.Set("local_cache_id", vv); err != nil {
				return fmt.Errorf("Error reading local_cache_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cache_id: %v", err)
		}
	}

	if err = d.Set("remote_cache_id", flattenWanoptRemoteStorageRemoteCacheId(o["remote-cache-id"], d, "remote_cache_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-cache-id"], "WanoptRemoteStorage-RemoteCacheId"); ok {
			if err = d.Set("remote_cache_id", vv); err != nil {
				return fmt.Errorf("Error reading remote_cache_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_cache_id: %v", err)
		}
	}

	if err = d.Set("remote_cache_ip", flattenWanoptRemoteStorageRemoteCacheIp(o["remote-cache-ip"], d, "remote_cache_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-cache-ip"], "WanoptRemoteStorage-RemoteCacheIp"); ok {
			if err = d.Set("remote_cache_ip", vv); err != nil {
				return fmt.Errorf("Error reading remote_cache_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_cache_ip: %v", err)
		}
	}

	if err = d.Set("status", flattenWanoptRemoteStorageStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WanoptRemoteStorage-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWanoptRemoteStorageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptRemoteStorageLocalCacheId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageRemoteCacheId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageRemoteCacheIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptRemoteStorage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("local_cache_id"); ok || d.HasChange("local_cache_id") {
		t, err := expandWanoptRemoteStorageLocalCacheId(d, v, "local_cache_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cache-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_cache_id"); ok || d.HasChange("remote_cache_id") {
		t, err := expandWanoptRemoteStorageRemoteCacheId(d, v, "remote_cache_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-cache-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_cache_ip"); ok || d.HasChange("remote_cache_ip") {
		t, err := expandWanoptRemoteStorageRemoteCacheIp(d, v, "remote_cache_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-cache-ip"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWanoptRemoteStorageStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
