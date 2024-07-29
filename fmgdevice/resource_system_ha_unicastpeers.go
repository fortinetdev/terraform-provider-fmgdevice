// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Number of unicast peers.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHaUnicastPeers() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaUnicastPeersCreate,
		Read:   resourceSystemHaUnicastPeersRead,
		Update: resourceSystemHaUnicastPeersUpdate,
		Delete: resourceSystemHaUnicastPeersDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"peer_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaUnicastPeersCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemHaUnicastPeers(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemHaUnicastPeers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHaUnicastPeers(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemHaUnicastPeers resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaUnicastPeersRead(d, m)
}

func resourceSystemHaUnicastPeersUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemHaUnicastPeers(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaUnicastPeers resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHaUnicastPeers(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemHaUnicastPeers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemHaUnicastPeersRead(d, m)
}

func resourceSystemHaUnicastPeersDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemHaUnicastPeers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHaUnicastPeers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaUnicastPeersRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHaUnicastPeers(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaUnicastPeers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHaUnicastPeers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHaUnicastPeers resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaUnicastPeersId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastPeersPeerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHaUnicastPeers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemHaUnicastPeersId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemHaUnicastPeers-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("peer_ip", flattenSystemHaUnicastPeersPeerIp2edl(o["peer-ip"], d, "peer_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-ip"], "SystemHaUnicastPeers-PeerIp"); ok {
			if err = d.Set("peer_ip", vv); err != nil {
				return fmt.Errorf("Error reading peer_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemHaUnicastPeersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaUnicastPeersId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastPeersPeerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHaUnicastPeers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemHaUnicastPeersId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("peer_ip"); ok || d.HasChange("peer_ip") {
		t, err := expandSystemHaUnicastPeersPeerIp2edl(d, v, "peer_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-ip"] = t
		}
	}

	return &obj, nil
}
