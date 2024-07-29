// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Fast failover peer wireless controller list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerInterControllerInterControllerPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerInterControllerInterControllerPeerCreate,
		Read:   resourceWirelessControllerInterControllerInterControllerPeerRead,
		Update: resourceWirelessControllerInterControllerInterControllerPeerUpdate,
		Delete: resourceWirelessControllerInterControllerInterControllerPeerDelete,

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
			"peer_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"peer_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerInterControllerInterControllerPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectWirelessControllerInterControllerInterControllerPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerInterControllerInterControllerPeer resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerInterControllerInterControllerPeer(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerInterControllerInterControllerPeer resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerInterControllerInterControllerPeerRead(d, m)
}

func resourceWirelessControllerInterControllerInterControllerPeerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerInterControllerInterControllerPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerInterControllerInterControllerPeer resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerInterControllerInterControllerPeer(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerInterControllerInterControllerPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerInterControllerInterControllerPeerRead(d, m)
}

func resourceWirelessControllerInterControllerInterControllerPeerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWirelessControllerInterControllerInterControllerPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerInterControllerInterControllerPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerInterControllerInterControllerPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerInterControllerInterControllerPeer(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerInterControllerInterControllerPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerInterControllerInterControllerPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerInterControllerInterControllerPeer resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerInterControllerInterControllerPeerId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerPort2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerPriority2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerInterControllerInterControllerPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerInterControllerInterControllerPeerId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerInterControllerInterControllerPeer-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("peer_ip", flattenWirelessControllerInterControllerInterControllerPeerPeerIp2edl(o["peer-ip"], d, "peer_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-ip"], "WirelessControllerInterControllerInterControllerPeer-PeerIp"); ok {
			if err = d.Set("peer_ip", vv); err != nil {
				return fmt.Errorf("Error reading peer_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_ip: %v", err)
		}
	}

	if err = d.Set("peer_port", flattenWirelessControllerInterControllerInterControllerPeerPeerPort2edl(o["peer-port"], d, "peer_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-port"], "WirelessControllerInterControllerInterControllerPeer-PeerPort"); ok {
			if err = d.Set("peer_port", vv); err != nil {
				return fmt.Errorf("Error reading peer_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_port: %v", err)
		}
	}

	if err = d.Set("peer_priority", flattenWirelessControllerInterControllerInterControllerPeerPeerPriority2edl(o["peer-priority"], d, "peer_priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-priority"], "WirelessControllerInterControllerInterControllerPeer-PeerPriority"); ok {
			if err = d.Set("peer_priority", vv); err != nil {
				return fmt.Errorf("Error reading peer_priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_priority: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerInterControllerInterControllerPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerInterControllerInterControllerPeerId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerPort2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerPriority2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerInterControllerInterControllerPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerInterControllerInterControllerPeerId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("peer_ip"); ok || d.HasChange("peer_ip") {
		t, err := expandWirelessControllerInterControllerInterControllerPeerPeerIp2edl(d, v, "peer_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-ip"] = t
		}
	}

	if v, ok := d.GetOk("peer_port"); ok || d.HasChange("peer_port") {
		t, err := expandWirelessControllerInterControllerInterControllerPeerPeerPort2edl(d, v, "peer_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-port"] = t
		}
	}

	if v, ok := d.GetOk("peer_priority"); ok || d.HasChange("peer_priority") {
		t, err := expandWirelessControllerInterControllerInterControllerPeerPeerPriority2edl(d, v, "peer_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-priority"] = t
		}
	}

	return &obj, nil
}
