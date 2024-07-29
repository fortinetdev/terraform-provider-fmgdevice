// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure GRE tunnel.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGreTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGreTunnelCreate,
		Read:   resourceSystemGreTunnelRead,
		Update: resourceSystemGreTunnelUpdate,
		Delete: resourceSystemGreTunnelDelete,

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
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"checksum_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"checksum_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp_copying": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keepalive_failtimes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"key_inbound": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"key_outbound": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sequence_number_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sequence_number_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemGreTunnelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemGreTunnel(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemGreTunnel resource while getting object: %v", err)
	}

	_, err = c.CreateSystemGreTunnel(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemGreTunnel resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemGreTunnelRead(d, m)
}

func resourceSystemGreTunnelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemGreTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGreTunnel resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemGreTunnel(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemGreTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemGreTunnelRead(d, m)
}

func resourceSystemGreTunnelDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemGreTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGreTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGreTunnelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemGreTunnel(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemGreTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGreTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGreTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemGreTunnelAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelChecksumReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelChecksumTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelDscpCopying(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemGreTunnelIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelKeepaliveFailtimes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelKeyInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelKeyOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelLocalGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelRemoteGw6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelSequenceNumberReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelSequenceNumberTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGreTunnelUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGreTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_asic_offload", flattenSystemGreTunnelAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "SystemGreTunnel-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("checksum_reception", flattenSystemGreTunnelChecksumReception(o["checksum-reception"], d, "checksum_reception")); err != nil {
		if vv, ok := fortiAPIPatch(o["checksum-reception"], "SystemGreTunnel-ChecksumReception"); ok {
			if err = d.Set("checksum_reception", vv); err != nil {
				return fmt.Errorf("Error reading checksum_reception: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading checksum_reception: %v", err)
		}
	}

	if err = d.Set("checksum_transmission", flattenSystemGreTunnelChecksumTransmission(o["checksum-transmission"], d, "checksum_transmission")); err != nil {
		if vv, ok := fortiAPIPatch(o["checksum-transmission"], "SystemGreTunnel-ChecksumTransmission"); ok {
			if err = d.Set("checksum_transmission", vv); err != nil {
				return fmt.Errorf("Error reading checksum_transmission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading checksum_transmission: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenSystemGreTunnelDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode"], "SystemGreTunnel-Diffservcode"); ok {
			if err = d.Set("diffservcode", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	if err = d.Set("dscp_copying", flattenSystemGreTunnelDscpCopying(o["dscp-copying"], d, "dscp_copying")); err != nil {
		if vv, ok := fortiAPIPatch(o["dscp-copying"], "SystemGreTunnel-DscpCopying"); ok {
			if err = d.Set("dscp_copying", vv); err != nil {
				return fmt.Errorf("Error reading dscp_copying: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dscp_copying: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemGreTunnelInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "SystemGreTunnel-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemGreTunnelIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip-version"], "SystemGreTunnel-IpVersion"); ok {
			if err = d.Set("ip_version", vv); err != nil {
				return fmt.Errorf("Error reading ip_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("keepalive_failtimes", flattenSystemGreTunnelKeepaliveFailtimes(o["keepalive-failtimes"], d, "keepalive_failtimes")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive-failtimes"], "SystemGreTunnel-KeepaliveFailtimes"); ok {
			if err = d.Set("keepalive_failtimes", vv); err != nil {
				return fmt.Errorf("Error reading keepalive_failtimes: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive_failtimes: %v", err)
		}
	}

	if err = d.Set("keepalive_interval", flattenSystemGreTunnelKeepaliveInterval(o["keepalive-interval"], d, "keepalive_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["keepalive-interval"], "SystemGreTunnel-KeepaliveInterval"); ok {
			if err = d.Set("keepalive_interval", vv); err != nil {
				return fmt.Errorf("Error reading keepalive_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading keepalive_interval: %v", err)
		}
	}

	if err = d.Set("key_inbound", flattenSystemGreTunnelKeyInbound(o["key-inbound"], d, "key_inbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-inbound"], "SystemGreTunnel-KeyInbound"); ok {
			if err = d.Set("key_inbound", vv); err != nil {
				return fmt.Errorf("Error reading key_inbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_inbound: %v", err)
		}
	}

	if err = d.Set("key_outbound", flattenSystemGreTunnelKeyOutbound(o["key-outbound"], d, "key_outbound")); err != nil {
		if vv, ok := fortiAPIPatch(o["key-outbound"], "SystemGreTunnel-KeyOutbound"); ok {
			if err = d.Set("key_outbound", vv); err != nil {
				return fmt.Errorf("Error reading key_outbound: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading key_outbound: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenSystemGreTunnelLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw"], "SystemGreTunnel-LocalGw"); ok {
			if err = d.Set("local_gw", vv); err != nil {
				return fmt.Errorf("Error reading local_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("local_gw6", flattenSystemGreTunnelLocalGw6(o["local-gw6"], d, "local_gw6")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-gw6"], "SystemGreTunnel-LocalGw6"); ok {
			if err = d.Set("local_gw6", vv); err != nil {
				return fmt.Errorf("Error reading local_gw6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_gw6: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemGreTunnelName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemGreTunnel-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenSystemGreTunnelRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw"], "SystemGreTunnel-RemoteGw"); ok {
			if err = d.Set("remote_gw", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("remote_gw6", flattenSystemGreTunnelRemoteGw6(o["remote-gw6"], d, "remote_gw6")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-gw6"], "SystemGreTunnel-RemoteGw6"); ok {
			if err = d.Set("remote_gw6", vv); err != nil {
				return fmt.Errorf("Error reading remote_gw6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_gw6: %v", err)
		}
	}

	if err = d.Set("sequence_number_reception", flattenSystemGreTunnelSequenceNumberReception(o["sequence-number-reception"], d, "sequence_number_reception")); err != nil {
		if vv, ok := fortiAPIPatch(o["sequence-number-reception"], "SystemGreTunnel-SequenceNumberReception"); ok {
			if err = d.Set("sequence_number_reception", vv); err != nil {
				return fmt.Errorf("Error reading sequence_number_reception: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sequence_number_reception: %v", err)
		}
	}

	if err = d.Set("sequence_number_transmission", flattenSystemGreTunnelSequenceNumberTransmission(o["sequence-number-transmission"], d, "sequence_number_transmission")); err != nil {
		if vv, ok := fortiAPIPatch(o["sequence-number-transmission"], "SystemGreTunnel-SequenceNumberTransmission"); ok {
			if err = d.Set("sequence_number_transmission", vv); err != nil {
				return fmt.Errorf("Error reading sequence_number_transmission: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sequence_number_transmission: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenSystemGreTunnelUseSdwan(o["use-sdwan"], d, "use_sdwan")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-sdwan"], "SystemGreTunnel-UseSdwan"); ok {
			if err = d.Set("use_sdwan", vv); err != nil {
				return fmt.Errorf("Error reading use_sdwan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	return nil
}

func flattenSystemGreTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGreTunnelAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelChecksumReception(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelChecksumTransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelDscpCopying(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemGreTunnelIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeepaliveFailtimes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeepaliveInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeyInbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelKeyOutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelLocalGw6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelRemoteGw6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelSequenceNumberReception(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelSequenceNumberTransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGreTunnelUseSdwan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGreTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandSystemGreTunnelAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("checksum_reception"); ok || d.HasChange("checksum_reception") {
		t, err := expandSystemGreTunnelChecksumReception(d, v, "checksum_reception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["checksum-reception"] = t
		}
	}

	if v, ok := d.GetOk("checksum_transmission"); ok || d.HasChange("checksum_transmission") {
		t, err := expandSystemGreTunnelChecksumTransmission(d, v, "checksum_transmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["checksum-transmission"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok || d.HasChange("diffservcode") {
		t, err := expandSystemGreTunnelDiffservcode(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	if v, ok := d.GetOk("dscp_copying"); ok || d.HasChange("dscp_copying") {
		t, err := expandSystemGreTunnelDscpCopying(d, v, "dscp_copying")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dscp-copying"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandSystemGreTunnelInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok || d.HasChange("ip_version") {
		t, err := expandSystemGreTunnelIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("keepalive_failtimes"); ok || d.HasChange("keepalive_failtimes") {
		t, err := expandSystemGreTunnelKeepaliveFailtimes(d, v, "keepalive_failtimes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive-failtimes"] = t
		}
	}

	if v, ok := d.GetOk("keepalive_interval"); ok || d.HasChange("keepalive_interval") {
		t, err := expandSystemGreTunnelKeepaliveInterval(d, v, "keepalive_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive-interval"] = t
		}
	}

	if v, ok := d.GetOk("key_inbound"); ok || d.HasChange("key_inbound") {
		t, err := expandSystemGreTunnelKeyInbound(d, v, "key_inbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-inbound"] = t
		}
	}

	if v, ok := d.GetOk("key_outbound"); ok || d.HasChange("key_outbound") {
		t, err := expandSystemGreTunnelKeyOutbound(d, v, "key_outbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-outbound"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok || d.HasChange("local_gw") {
		t, err := expandSystemGreTunnelLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw6"); ok || d.HasChange("local_gw6") {
		t, err := expandSystemGreTunnelLocalGw6(d, v, "local_gw6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemGreTunnelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok || d.HasChange("remote_gw") {
		t, err := expandSystemGreTunnelRemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6"); ok || d.HasChange("remote_gw6") {
		t, err := expandSystemGreTunnelRemoteGw6(d, v, "remote_gw6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("sequence_number_reception"); ok || d.HasChange("sequence_number_reception") {
		t, err := expandSystemGreTunnelSequenceNumberReception(d, v, "sequence_number_reception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sequence-number-reception"] = t
		}
	}

	if v, ok := d.GetOk("sequence_number_transmission"); ok || d.HasChange("sequence_number_transmission") {
		t, err := expandSystemGreTunnelSequenceNumberTransmission(d, v, "sequence_number_transmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sequence-number-transmission"] = t
		}
	}

	if v, ok := d.GetOk("use_sdwan"); ok || d.HasChange("use_sdwan") {
		t, err := expandSystemGreTunnelUseSdwan(d, v, "use_sdwan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	return &obj, nil
}
