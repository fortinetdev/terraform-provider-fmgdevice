// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure WAN optimization authentication groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptAuthGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptAuthGroupCreate,
		Read:   resourceWanoptAuthGroupRead,
		Update: resourceWanoptAuthGroupUpdate,
		Delete: resourceWanoptAuthGroupDelete,

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
			"auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"peer": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"peer_accept": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
		},
	}
}

func resourceWanoptAuthGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptAuthGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptAuthGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadWanoptAuthGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateWanoptAuthGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating WanoptAuthGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateWanoptAuthGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating WanoptAuthGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceWanoptAuthGroupRead(d, m)
}

func resourceWanoptAuthGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptAuthGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptAuthGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWanoptAuthGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptAuthGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceWanoptAuthGroupRead(d, m)
}

func resourceWanoptAuthGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptAuthGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptAuthGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptAuthGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptAuthGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading WanoptAuthGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptAuthGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptAuthGroup resource from API: %v", err)
	}
	return nil
}

func flattenWanoptAuthGroupAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptAuthGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWanoptAuthGroupPeerAccept(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptAuthGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auth_method", flattenWanoptAuthGroupAuthMethod(o["auth-method"], d, "auth_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-method"], "WanoptAuthGroup-AuthMethod"); ok {
			if err = d.Set("auth_method", vv); err != nil {
				return fmt.Errorf("Error reading auth_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("cert", flattenWanoptAuthGroupCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "WanoptAuthGroup-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("name", flattenWanoptAuthGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "WanoptAuthGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("peer", flattenWanoptAuthGroupPeer(o["peer"], d, "peer")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer"], "WanoptAuthGroup-Peer"); ok {
			if err = d.Set("peer", vv); err != nil {
				return fmt.Errorf("Error reading peer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peer_accept", flattenWanoptAuthGroupPeerAccept(o["peer-accept"], d, "peer_accept")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-accept"], "WanoptAuthGroup-PeerAccept"); ok {
			if err = d.Set("peer_accept", vv); err != nil {
				return fmt.Errorf("Error reading peer_accept: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_accept: %v", err)
		}
	}

	return nil
}

func flattenWanoptAuthGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptAuthGroupAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptAuthGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWanoptAuthGroupPeerAccept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupPsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectWanoptAuthGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_method"); ok || d.HasChange("auth_method") {
		t, err := expandWanoptAuthGroupAuthMethod(d, v, "auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandWanoptAuthGroupCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandWanoptAuthGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok || d.HasChange("peer") {
		t, err := expandWanoptAuthGroupPeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("peer_accept"); ok || d.HasChange("peer_accept") {
		t, err := expandWanoptAuthGroupPeerAccept(d, v, "peer_accept")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-accept"] = t
		}
	}

	if v, ok := d.GetOk("psk"); ok || d.HasChange("psk") {
		t, err := expandWanoptAuthGroupPsk(d, v, "psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk"] = t
		}
	}

	return &obj, nil
}
