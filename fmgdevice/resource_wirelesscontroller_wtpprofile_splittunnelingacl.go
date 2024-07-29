// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Split tunneling ACL filter list.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpProfileSplitTunnelingAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpProfileSplitTunnelingAclCreate,
		Read:   resourceWirelessControllerWtpProfileSplitTunnelingAclRead,
		Update: resourceWirelessControllerWtpProfileSplitTunnelingAclUpdate,
		Delete: resourceWirelessControllerWtpProfileSplitTunnelingAclDelete,

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
			"device_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wtp_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dest_ip": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfileSplitTunnelingAclCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectWirelessControllerWtpProfileSplitTunnelingAcl(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpProfileSplitTunnelingAcl resource while getting object: %v", err)
	}

	_, err = c.CreateWirelessControllerWtpProfileSplitTunnelingAcl(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerWtpProfileSplitTunnelingAclRead(d, m)
}

func resourceWirelessControllerWtpProfileSplitTunnelingAclUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	obj, err := getObjectWirelessControllerWtpProfileSplitTunnelingAcl(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileSplitTunnelingAcl resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerWtpProfileSplitTunnelingAcl(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWirelessControllerWtpProfileSplitTunnelingAclRead(d, m)
}

func resourceWirelessControllerWtpProfileSplitTunnelingAclDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	wtp_profile := d.Get("wtp_profile").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	err = c.DeleteWirelessControllerWtpProfileSplitTunnelingAcl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfileSplitTunnelingAclRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	wtp_profile := d.Get("wtp_profile").(string)
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
	if wtp_profile == "" {
		wtp_profile = importOptionChecking(m.(*FortiClient).Cfg, "wtp_profile")
		if wtp_profile == "" {
			return fmt.Errorf("Parameter wtp_profile is missing")
		}
		if err = d.Set("wtp_profile", wtp_profile); err != nil {
			return fmt.Errorf("Error set params wtp_profile: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["wtp_profile"] = wtp_profile

	o, err := c.ReadWirelessControllerWtpProfileSplitTunnelingAcl(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileSplitTunnelingAcl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfileSplitTunnelingAcl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfileSplitTunnelingAcl resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dest_ip", flattenWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(o["dest-ip"], d, "dest_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dest-ip"], "WirelessControllerWtpProfileSplitTunnelingAcl-DestIp"); ok {
			if err = d.Set("dest_ip", vv); err != nil {
				return fmt.Errorf("Error reading dest_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dest_ip: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWirelessControllerWtpProfileSplitTunnelingAclId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WirelessControllerWtpProfileSplitTunnelingAcl-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dest_ip"); ok || d.HasChange("dest_ip") {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAclDestIp2edl(d, v, "dest_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dest-ip"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAclId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
