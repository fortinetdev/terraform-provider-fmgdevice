// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Modify cache-service destination peer list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptCacheServiceDstPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptCacheServiceDstPeerCreate,
		Read:   resourceWanoptCacheServiceDstPeerRead,
		Update: resourceWanoptCacheServiceDstPeerUpdate,
		Delete: resourceWanoptCacheServiceDstPeerDelete,

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
			"auth_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"encode_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptCacheServiceDstPeerCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWanoptCacheServiceDstPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptCacheServiceDstPeer resource while getting object: %v", err)
	}

	_, err = c.CreateWanoptCacheServiceDstPeer(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WanoptCacheServiceDstPeer resource: %v", err)
	}

	d.SetId(getStringKey(d, "device_id"))

	return resourceWanoptCacheServiceDstPeerRead(d, m)
}

func resourceWanoptCacheServiceDstPeerUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWanoptCacheServiceDstPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheServiceDstPeer resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptCacheServiceDstPeer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheServiceDstPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "device_id"))

	return resourceWanoptCacheServiceDstPeerRead(d, m)
}

func resourceWanoptCacheServiceDstPeerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptCacheServiceDstPeer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptCacheServiceDstPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptCacheServiceDstPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptCacheServiceDstPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptCacheServiceDstPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptCacheServiceDstPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptCacheServiceDstPeer resource from API: %v", err)
	}
	return nil
}

func flattenWanoptCacheServiceDstPeerAuthType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerDeviceId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerEncodeType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptCacheServiceDstPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_type", flattenWanoptCacheServiceDstPeerAuthType2edl(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "WanoptCacheServiceDstPeer-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("device_id", flattenWanoptCacheServiceDstPeerDeviceId2edl(o["device-id"], d, "device_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-id"], "WanoptCacheServiceDstPeer-DeviceId"); ok {
			if err = d.Set("device_id", vv); err != nil {
				return fmt.Errorf("Error reading device_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("encode_type", flattenWanoptCacheServiceDstPeerEncodeType2edl(o["encode-type"], d, "encode_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["encode-type"], "WanoptCacheServiceDstPeer-EncodeType"); ok {
			if err = d.Set("encode_type", vv); err != nil {
				return fmt.Errorf("Error reading encode_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading encode_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenWanoptCacheServiceDstPeerIp2edl(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WanoptCacheServiceDstPeer-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("priority", flattenWanoptCacheServiceDstPeerPriority2edl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "WanoptCacheServiceDstPeer-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenWanoptCacheServiceDstPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptCacheServiceDstPeerAuthType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerDeviceId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerEncodeType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptCacheServiceDstPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandWanoptCacheServiceDstPeerAuthType2edl(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("device_id"); ok || d.HasChange("device_id") {
		t, err := expandWanoptCacheServiceDstPeerDeviceId2edl(d, v, "device_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("encode_type"); ok || d.HasChange("encode_type") {
		t, err := expandWanoptCacheServiceDstPeerEncodeType2edl(d, v, "encode_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encode-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWanoptCacheServiceDstPeerIp2edl(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandWanoptCacheServiceDstPeerPriority2edl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
