// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure WCCP.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemWccp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemWccpCreate,
		Read:   resourceSystemWccpRead,
		Update: resourceSystemWccpUpdate,
		Delete: resourceSystemWccpDelete,

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
			"assignment_bucket_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assignment_dstaddr_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assignment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assignment_srcaddr_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assignment_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_engine_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"ports_defined": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_hash": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"return_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemWccpCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemWccp(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemWccp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemWccp(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemWccp resource: %v", err)
	}

	d.SetId(getStringKey(d, "service_id"))

	return resourceSystemWccpRead(d, m)
}

func resourceSystemWccpUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	obj, err := getObjectSystemWccp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemWccp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemWccp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemWccp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "service_id"))

	return resourceSystemWccpRead(d, m)
}

func resourceSystemWccpDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	err = c.DeleteSystemWccp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemWccp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemWccpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemWccp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemWccp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemWccp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemWccp resource from API: %v", err)
	}
	return nil
}

func flattenSystemWccpAssignmentBucketFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpAssignmentDstaddrMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpAssignmentMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpAssignmentSrcaddrMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpAssignmentWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpCacheEngineMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpForwardMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpGroupAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemWccpPortsDefined(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpPrimaryHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemWccpPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpReturnMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpRouterList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemWccpServerList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemWccpServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemWccpServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemWccp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("assignment_bucket_format", flattenSystemWccpAssignmentBucketFormat(o["assignment-bucket-format"], d, "assignment_bucket_format")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment-bucket-format"], "SystemWccp-AssignmentBucketFormat"); ok {
			if err = d.Set("assignment_bucket_format", vv); err != nil {
				return fmt.Errorf("Error reading assignment_bucket_format: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment_bucket_format: %v", err)
		}
	}

	if err = d.Set("assignment_dstaddr_mask", flattenSystemWccpAssignmentDstaddrMask(o["assignment-dstaddr-mask"], d, "assignment_dstaddr_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment-dstaddr-mask"], "SystemWccp-AssignmentDstaddrMask"); ok {
			if err = d.Set("assignment_dstaddr_mask", vv); err != nil {
				return fmt.Errorf("Error reading assignment_dstaddr_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment_dstaddr_mask: %v", err)
		}
	}

	if err = d.Set("assignment_method", flattenSystemWccpAssignmentMethod(o["assignment-method"], d, "assignment_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment-method"], "SystemWccp-AssignmentMethod"); ok {
			if err = d.Set("assignment_method", vv); err != nil {
				return fmt.Errorf("Error reading assignment_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment_method: %v", err)
		}
	}

	if err = d.Set("assignment_srcaddr_mask", flattenSystemWccpAssignmentSrcaddrMask(o["assignment-srcaddr-mask"], d, "assignment_srcaddr_mask")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment-srcaddr-mask"], "SystemWccp-AssignmentSrcaddrMask"); ok {
			if err = d.Set("assignment_srcaddr_mask", vv); err != nil {
				return fmt.Errorf("Error reading assignment_srcaddr_mask: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment_srcaddr_mask: %v", err)
		}
	}

	if err = d.Set("assignment_weight", flattenSystemWccpAssignmentWeight(o["assignment-weight"], d, "assignment_weight")); err != nil {
		if vv, ok := fortiAPIPatch(o["assignment-weight"], "SystemWccp-AssignmentWeight"); ok {
			if err = d.Set("assignment_weight", vv); err != nil {
				return fmt.Errorf("Error reading assignment_weight: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading assignment_weight: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemWccpAuthentication(o["authentication"], d, "authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["authentication"], "SystemWccp-Authentication"); ok {
			if err = d.Set("authentication", vv); err != nil {
				return fmt.Errorf("Error reading authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("cache_engine_method", flattenSystemWccpCacheEngineMethod(o["cache-engine-method"], d, "cache_engine_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-engine-method"], "SystemWccp-CacheEngineMethod"); ok {
			if err = d.Set("cache_engine_method", vv); err != nil {
				return fmt.Errorf("Error reading cache_engine_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_engine_method: %v", err)
		}
	}

	if err = d.Set("cache_id", flattenSystemWccpCacheId(o["cache-id"], d, "cache_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["cache-id"], "SystemWccp-CacheId"); ok {
			if err = d.Set("cache_id", vv); err != nil {
				return fmt.Errorf("Error reading cache_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cache_id: %v", err)
		}
	}

	if err = d.Set("forward_method", flattenSystemWccpForwardMethod(o["forward-method"], d, "forward_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["forward-method"], "SystemWccp-ForwardMethod"); ok {
			if err = d.Set("forward_method", vv); err != nil {
				return fmt.Errorf("Error reading forward_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading forward_method: %v", err)
		}
	}

	if err = d.Set("group_address", flattenSystemWccpGroupAddress(o["group-address"], d, "group_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-address"], "SystemWccp-GroupAddress"); ok {
			if err = d.Set("group_address", vv); err != nil {
				return fmt.Errorf("Error reading group_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_address: %v", err)
		}
	}

	if err = d.Set("ports", flattenSystemWccpPorts(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "SystemWccp-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("ports_defined", flattenSystemWccpPortsDefined(o["ports-defined"], d, "ports_defined")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports-defined"], "SystemWccp-PortsDefined"); ok {
			if err = d.Set("ports_defined", vv); err != nil {
				return fmt.Errorf("Error reading ports_defined: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports_defined: %v", err)
		}
	}

	if err = d.Set("primary_hash", flattenSystemWccpPrimaryHash(o["primary-hash"], d, "primary_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["primary-hash"], "SystemWccp-PrimaryHash"); ok {
			if err = d.Set("primary_hash", vv); err != nil {
				return fmt.Errorf("Error reading primary_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading primary_hash: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemWccpPriority(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "SystemWccp-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemWccpProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemWccp-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("return_method", flattenSystemWccpReturnMethod(o["return-method"], d, "return_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["return-method"], "SystemWccp-ReturnMethod"); ok {
			if err = d.Set("return_method", vv); err != nil {
				return fmt.Errorf("Error reading return_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading return_method: %v", err)
		}
	}

	if err = d.Set("router_id", flattenSystemWccpRouterId(o["router-id"], d, "router_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-id"], "SystemWccp-RouterId"); ok {
			if err = d.Set("router_id", vv); err != nil {
				return fmt.Errorf("Error reading router_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("router_list", flattenSystemWccpRouterList(o["router-list"], d, "router_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["router-list"], "SystemWccp-RouterList"); ok {
			if err = d.Set("router_list", vv); err != nil {
				return fmt.Errorf("Error reading router_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading router_list: %v", err)
		}
	}

	if err = d.Set("server_list", flattenSystemWccpServerList(o["server-list"], d, "server_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-list"], "SystemWccp-ServerList"); ok {
			if err = d.Set("server_list", vv); err != nil {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_list: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemWccpServerType(o["server-type"], d, "server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-type"], "SystemWccp-ServerType"); ok {
			if err = d.Set("server_type", vv); err != nil {
				return fmt.Errorf("Error reading server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("service_id", flattenSystemWccpServiceId(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "SystemWccp-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("service_type", flattenSystemWccpServiceType(o["service-type"], d, "service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-type"], "SystemWccp-ServiceType"); ok {
			if err = d.Set("service_type", vv); err != nil {
				return fmt.Errorf("Error reading service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_type: %v", err)
		}
	}

	return nil
}

func flattenSystemWccpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemWccpAssignmentBucketFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpAssignmentDstaddrMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpAssignmentMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpAssignmentSrcaddrMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpAssignmentWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpCacheEngineMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpCacheId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpForwardMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpGroupAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemWccpPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemWccpPortsDefined(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpPrimaryHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemWccpPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpReturnMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpRouterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpRouterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemWccpServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemWccpServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemWccpServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemWccp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("assignment_bucket_format"); ok || d.HasChange("assignment_bucket_format") {
		t, err := expandSystemWccpAssignmentBucketFormat(d, v, "assignment_bucket_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment-bucket-format"] = t
		}
	}

	if v, ok := d.GetOk("assignment_dstaddr_mask"); ok || d.HasChange("assignment_dstaddr_mask") {
		t, err := expandSystemWccpAssignmentDstaddrMask(d, v, "assignment_dstaddr_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment-dstaddr-mask"] = t
		}
	}

	if v, ok := d.GetOk("assignment_method"); ok || d.HasChange("assignment_method") {
		t, err := expandSystemWccpAssignmentMethod(d, v, "assignment_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment-method"] = t
		}
	}

	if v, ok := d.GetOk("assignment_srcaddr_mask"); ok || d.HasChange("assignment_srcaddr_mask") {
		t, err := expandSystemWccpAssignmentSrcaddrMask(d, v, "assignment_srcaddr_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment-srcaddr-mask"] = t
		}
	}

	if v, ok := d.GetOk("assignment_weight"); ok || d.HasChange("assignment_weight") {
		t, err := expandSystemWccpAssignmentWeight(d, v, "assignment_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assignment-weight"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok || d.HasChange("authentication") {
		t, err := expandSystemWccpAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("cache_engine_method"); ok || d.HasChange("cache_engine_method") {
		t, err := expandSystemWccpCacheEngineMethod(d, v, "cache_engine_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-engine-method"] = t
		}
	}

	if v, ok := d.GetOk("cache_id"); ok || d.HasChange("cache_id") {
		t, err := expandSystemWccpCacheId(d, v, "cache_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-id"] = t
		}
	}

	if v, ok := d.GetOk("forward_method"); ok || d.HasChange("forward_method") {
		t, err := expandSystemWccpForwardMethod(d, v, "forward_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-method"] = t
		}
	}

	if v, ok := d.GetOk("group_address"); ok || d.HasChange("group_address") {
		t, err := expandSystemWccpGroupAddress(d, v, "group_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-address"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemWccpPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandSystemWccpPorts(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("ports_defined"); ok || d.HasChange("ports_defined") {
		t, err := expandSystemWccpPortsDefined(d, v, "ports_defined")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports-defined"] = t
		}
	}

	if v, ok := d.GetOk("primary_hash"); ok || d.HasChange("primary_hash") {
		t, err := expandSystemWccpPrimaryHash(d, v, "primary_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-hash"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandSystemWccpPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemWccpProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("return_method"); ok || d.HasChange("return_method") {
		t, err := expandSystemWccpReturnMethod(d, v, "return_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["return-method"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok || d.HasChange("router_id") {
		t, err := expandSystemWccpRouterId(d, v, "router_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("router_list"); ok || d.HasChange("router_list") {
		t, err := expandSystemWccpRouterList(d, v, "router_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-list"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandSystemWccpServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok || d.HasChange("server_type") {
		t, err := expandSystemWccpServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandSystemWccpServiceId(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("service_type"); ok || d.HasChange("service_type") {
		t, err := expandSystemWccpServiceType(d, v, "service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-type"] = t
		}
	}

	return &obj, nil
}
