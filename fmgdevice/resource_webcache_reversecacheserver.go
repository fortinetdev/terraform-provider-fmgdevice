// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device WebcacheReverseCacheServer

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebcacheReverseCacheServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebcacheReverseCacheServerCreate,
		Read:   resourceWebcacheReverseCacheServerRead,
		Update: resourceWebcacheReverseCacheServerUpdate,
		Delete: resourceWebcacheReverseCacheServerDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"prefetch_file": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebcacheReverseCacheServerCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebcacheReverseCacheServer(d)
	if err != nil {
		return fmt.Errorf("Error creating WebcacheReverseCacheServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWebcacheReverseCacheServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWebcacheReverseCacheServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WebcacheReverseCacheServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWebcacheReverseCacheServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WebcacheReverseCacheServer resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWebcacheReverseCacheServerRead(d, m)
}

func resourceWebcacheReverseCacheServerUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectWebcacheReverseCacheServer(d)
	if err != nil {
		return fmt.Errorf("Error updating WebcacheReverseCacheServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWebcacheReverseCacheServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WebcacheReverseCacheServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWebcacheReverseCacheServerRead(d, m)
}

func resourceWebcacheReverseCacheServerDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	wsParams["adom"] = adomv

	err = c.DeleteWebcacheReverseCacheServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WebcacheReverseCacheServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebcacheReverseCacheServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	o, err := c.ReadWebcacheReverseCacheServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WebcacheReverseCacheServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebcacheReverseCacheServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebcacheReverseCacheServer resource from API: %v", err)
	}
	return nil
}

func flattenWebcacheReverseCacheServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheReverseCacheServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheReverseCacheServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheReverseCacheServerPrefetchFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWebcacheReverseCacheServerPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebcacheReverseCacheServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebcacheReverseCacheServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenWebcacheReverseCacheServerIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WebcacheReverseCacheServer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("name", flattenWebcacheReverseCacheServerName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WebcacheReverseCacheServer-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenWebcacheReverseCacheServerPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "WebcacheReverseCacheServer-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("prefetch_file", flattenWebcacheReverseCacheServerPrefetchFile(o["prefetch-file"], d, "prefetch_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefetch-file"], "WebcacheReverseCacheServer-PrefetchFile"); ok {
			if err = d.Set("prefetch_file", vv); err != nil {
				return fmt.Errorf("Error reading prefetch_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefetch_file: %v", err)
		}
	}

	if err = d.Set("priority", flattenWebcacheReverseCacheServerPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "WebcacheReverseCacheServer-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("status", flattenWebcacheReverseCacheServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "WebcacheReverseCacheServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWebcacheReverseCacheServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebcacheReverseCacheServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheReverseCacheServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheReverseCacheServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheReverseCacheServerPrefetchFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWebcacheReverseCacheServerPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebcacheReverseCacheServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebcacheReverseCacheServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWebcacheReverseCacheServerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWebcacheReverseCacheServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandWebcacheReverseCacheServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("prefetch_file"); ok || d.HasChange("prefetch_file") {
		t, err := expandWebcacheReverseCacheServerPrefetchFile(d, v, "prefetch_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefetch-file"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandWebcacheReverseCacheServerPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandWebcacheReverseCacheServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
