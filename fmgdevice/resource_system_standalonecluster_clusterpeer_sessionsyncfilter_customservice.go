// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Only sessions using these custom services are synchronized. Use source and destination port ranges to define these custom services.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceCreate,
		Read:   resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceRead,
		Update: resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceUpdate,
		Delete: resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDelete,

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
			"cluster_peer": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dst_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"src_port_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	cluster_peer := d.Get("cluster_peer").(string)
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource while getting object: %v", err)
	}

	_, err = c.CreateSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceRead(d, m)
}

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceUpdate(d *schema.ResourceData, m interface{}) error {
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
	cluster_peer := d.Get("cluster_peer").(string)
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceRead(d, m)
}

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDelete(d *schema.ResourceData, m interface{}) error {
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
	cluster_peer := d.Get("cluster_peer").(string)
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	cluster_peer := d.Get("cluster_peer").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	if cluster_peer == "" {
		cluster_peer = importOptionChecking(m.(*FortiClient).Cfg, "cluster_peer")
		if cluster_peer == "" {
			return fmt.Errorf("Parameter cluster_peer is missing")
		}
		if err = d.Set("cluster_peer", cluster_peer); err != nil {
			return fmt.Errorf("Error set params cluster_peer: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["cluster_peer"] = cluster_peer

	o, err := c.ReadSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("dst_port_range", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange4thl(o["dst-port-range"], d, "dst_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst-port-range"], "SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService-DstPortRange"); ok {
			if err = d.Set("dst_port_range", vv); err != nil {
				return fmt.Errorf("Error reading dst_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst_port_range: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId4thl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src_port_range", flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange4thl(o["src-port-range"], d, "src_port_range")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-port-range"], "SystemStandaloneClusterClusterPeerSessionSyncFilterCustomService-SrcPortRange"); ok {
			if err = d.Set("src_port_range", vv); err != nil {
				return fmt.Errorf("Error reading src_port_range: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_port_range: %v", err)
		}
	}

	return nil
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst_port_range"); ok || d.HasChange("dst_port_range") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange4thl(d, v, "dst_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-port-range"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId4thl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("src_port_range"); ok || d.HasChange("src_port_range") {
		t, err := expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange4thl(d, v, "src_port_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-port-range"] = t
		}
	}

	return &obj, nil
}
