// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Designate cache-service for wan-optimization and webcache.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptCacheService() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptCacheServiceUpdate,
		Read:   resourceWanoptCacheServiceRead,
		Update: resourceWanoptCacheServiceUpdate,
		Delete: resourceWanoptCacheServiceDelete,

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
			"acceptable_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"collaboration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"device_id": &schema.Schema{
							Type:     schema.TypeString,
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
				},
			},
			"prefer_scenario": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"device_id": &schema.Schema{
							Type:     schema.TypeString,
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
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWanoptCacheServiceUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectWanoptCacheService(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheService resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptCacheService(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("WanoptCacheService")

	return resourceWanoptCacheServiceRead(d, m)
}

func resourceWanoptCacheServiceDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWanoptCacheService(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptCacheService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptCacheServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptCacheService(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WanoptCacheService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptCacheService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptCacheService resource from API: %v", err)
	}
	return nil
}

func flattenWanoptCacheServiceAcceptableConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceCollaboration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := i["auth-type"]; ok {
			v := flattenWanoptCacheServiceDstPeerAuthType(i["auth-type"], d, pre_append)
			tmp["auth_type"] = fortiAPISubPartPatch(v, "WanoptCacheService-DstPeer-AuthType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if _, ok := i["device-id"]; ok {
			v := flattenWanoptCacheServiceDstPeerDeviceId(i["device-id"], d, pre_append)
			tmp["device_id"] = fortiAPISubPartPatch(v, "WanoptCacheService-DstPeer-DeviceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if _, ok := i["encode-type"]; ok {
			v := flattenWanoptCacheServiceDstPeerEncodeType(i["encode-type"], d, pre_append)
			tmp["encode_type"] = fortiAPISubPartPatch(v, "WanoptCacheService-DstPeer-EncodeType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWanoptCacheServiceDstPeerIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WanoptCacheService-DstPeer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenWanoptCacheServiceDstPeerPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "WanoptCacheService-DstPeer-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptCacheServiceDstPeerAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerEncodeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServicePreferScenario(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := i["auth-type"]; ok {
			v := flattenWanoptCacheServiceSrcPeerAuthType(i["auth-type"], d, pre_append)
			tmp["auth_type"] = fortiAPISubPartPatch(v, "WanoptCacheService-SrcPeer-AuthType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if _, ok := i["device-id"]; ok {
			v := flattenWanoptCacheServiceSrcPeerDeviceId(i["device-id"], d, pre_append)
			tmp["device_id"] = fortiAPISubPartPatch(v, "WanoptCacheService-SrcPeer-DeviceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if _, ok := i["encode-type"]; ok {
			v := flattenWanoptCacheServiceSrcPeerEncodeType(i["encode-type"], d, pre_append)
			tmp["encode_type"] = fortiAPISubPartPatch(v, "WanoptCacheService-SrcPeer-EncodeType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenWanoptCacheServiceSrcPeerIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "WanoptCacheService-SrcPeer-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenWanoptCacheServiceSrcPeerPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "WanoptCacheService-SrcPeer-Priority")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenWanoptCacheServiceSrcPeerAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeerDeviceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeerEncodeType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeerPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptCacheService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("acceptable_connections", flattenWanoptCacheServiceAcceptableConnections(o["acceptable-connections"], d, "acceptable_connections")); err != nil {
		if vv, ok := fortiAPIPatch(o["acceptable-connections"], "WanoptCacheService-AcceptableConnections"); ok {
			if err = d.Set("acceptable_connections", vv); err != nil {
				return fmt.Errorf("Error reading acceptable_connections: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acceptable_connections: %v", err)
		}
	}

	if err = d.Set("collaboration", flattenWanoptCacheServiceCollaboration(o["collaboration"], d, "collaboration")); err != nil {
		if vv, ok := fortiAPIPatch(o["collaboration"], "WanoptCacheService-Collaboration"); ok {
			if err = d.Set("collaboration", vv); err != nil {
				return fmt.Errorf("Error reading collaboration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading collaboration: %v", err)
		}
	}

	if err = d.Set("device_id", flattenWanoptCacheServiceDeviceId(o["device-id"], d, "device_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["device-id"], "WanoptCacheService-DeviceId"); ok {
			if err = d.Set("device_id", vv); err != nil {
				return fmt.Errorf("Error reading device_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dst_peer", flattenWanoptCacheServiceDstPeer(o["dst-peer"], d, "dst_peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["dst-peer"], "WanoptCacheService-DstPeer"); ok {
				if err = d.Set("dst_peer", vv); err != nil {
					return fmt.Errorf("Error reading dst_peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dst_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dst_peer"); ok {
			if err = d.Set("dst_peer", flattenWanoptCacheServiceDstPeer(o["dst-peer"], d, "dst_peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["dst-peer"], "WanoptCacheService-DstPeer"); ok {
					if err = d.Set("dst_peer", vv); err != nil {
						return fmt.Errorf("Error reading dst_peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dst_peer: %v", err)
				}
			}
		}
	}

	if err = d.Set("prefer_scenario", flattenWanoptCacheServicePreferScenario(o["prefer-scenario"], d, "prefer_scenario")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefer-scenario"], "WanoptCacheService-PreferScenario"); ok {
			if err = d.Set("prefer_scenario", vv); err != nil {
				return fmt.Errorf("Error reading prefer_scenario: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefer_scenario: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_peer", flattenWanoptCacheServiceSrcPeer(o["src-peer"], d, "src_peer")); err != nil {
			if vv, ok := fortiAPIPatch(o["src-peer"], "WanoptCacheService-SrcPeer"); ok {
				if err = d.Set("src_peer", vv); err != nil {
					return fmt.Errorf("Error reading src_peer: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading src_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_peer"); ok {
			if err = d.Set("src_peer", flattenWanoptCacheServiceSrcPeer(o["src-peer"], d, "src_peer")); err != nil {
				if vv, ok := fortiAPIPatch(o["src-peer"], "WanoptCacheService-SrcPeer"); ok {
					if err = d.Set("src_peer", vv); err != nil {
						return fmt.Errorf("Error reading src_peer: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading src_peer: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWanoptCacheServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptCacheServiceAcceptableConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceCollaboration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-type"], _ = expandWanoptCacheServiceDstPeerAuthType(d, i["auth_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-id"], _ = expandWanoptCacheServiceDstPeerDeviceId(d, i["device_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encode-type"], _ = expandWanoptCacheServiceDstPeerEncodeType(d, i["encode_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWanoptCacheServiceDstPeerIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandWanoptCacheServiceDstPeerPriority(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptCacheServiceDstPeerAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerEncodeType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServicePreferScenario(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-type"], _ = expandWanoptCacheServiceSrcPeerAuthType(d, i["auth_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device-id"], _ = expandWanoptCacheServiceSrcPeerDeviceId(d, i["device_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["encode-type"], _ = expandWanoptCacheServiceSrcPeerEncodeType(d, i["encode_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandWanoptCacheServiceSrcPeerIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandWanoptCacheServiceSrcPeerPriority(d, i["priority"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandWanoptCacheServiceSrcPeerAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerDeviceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerEncodeType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptCacheService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("acceptable_connections"); ok || d.HasChange("acceptable_connections") {
		t, err := expandWanoptCacheServiceAcceptableConnections(d, v, "acceptable_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acceptable-connections"] = t
		}
	}

	if v, ok := d.GetOk("collaboration"); ok || d.HasChange("collaboration") {
		t, err := expandWanoptCacheServiceCollaboration(d, v, "collaboration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collaboration"] = t
		}
	}

	if v, ok := d.GetOk("device_id"); ok || d.HasChange("device_id") {
		t, err := expandWanoptCacheServiceDeviceId(d, v, "device_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("dst_peer"); ok || d.HasChange("dst_peer") {
		t, err := expandWanoptCacheServiceDstPeer(d, v, "dst_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-peer"] = t
		}
	}

	if v, ok := d.GetOk("prefer_scenario"); ok || d.HasChange("prefer_scenario") {
		t, err := expandWanoptCacheServicePreferScenario(d, v, "prefer_scenario")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer-scenario"] = t
		}
	}

	if v, ok := d.GetOk("src_peer"); ok || d.HasChange("src_peer") {
		t, err := expandWanoptCacheServiceSrcPeer(d, v, "src_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-peer"] = t
		}
	}

	return &obj, nil
}
