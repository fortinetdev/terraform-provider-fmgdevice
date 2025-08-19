// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure public cloud VPN service.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnVpnCreate,
		Read:   resourceSystemSdnVpnRead,
		Update: resourceSystemSdnVpnUpdate,
		Delete: resourceSystemSdnVpnDelete,

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
			"bgp_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"bgp_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bgp_seq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cgw_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cgw_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgw_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"internal_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"local_cidr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nat_traversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"remote_cidr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"remote_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tgw_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tgw_vpn_rtbl_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trtbl_attachment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tunnel_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vgw_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpn_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSdnVpnCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSdnVpn(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnVpn resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSdnVpn(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnVpn resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdnVpnRead(d, m)
}

func resourceSystemSdnVpnUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSdnVpn(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnVpn resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdnVpn(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnVpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemSdnVpnRead(d, m)
}

func resourceSystemSdnVpnDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSdnVpn(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnVpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnVpnRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnVpn(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnVpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnVpn(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnVpn resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnVpnBgpAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnBgpFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnBgpSeq(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnCgwGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnCgwName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnCgwId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnInternalInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnVpnLocalCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnVpnName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnNatTraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnRemoteCidr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnVpnRemoteType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnRoutingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnSdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnVpnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnSubnetId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnTgwId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnTgwVpnRtblId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnTrtblAttachment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdnVpnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnVgwId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdnVpnVpnId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdnVpn(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bgp_as", flattenSystemSdnVpnBgpAs(o["bgp-as"], d, "bgp_as")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-as"], "SystemSdnVpn-BgpAs"); ok {
			if err = d.Set("bgp_as", vv); err != nil {
				return fmt.Errorf("Error reading bgp_as: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_as: %v", err)
		}
	}

	if err = d.Set("bgp_from", flattenSystemSdnVpnBgpFrom(o["bgp-from"], d, "bgp_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-from"], "SystemSdnVpn-BgpFrom"); ok {
			if err = d.Set("bgp_from", vv); err != nil {
				return fmt.Errorf("Error reading bgp_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_from: %v", err)
		}
	}

	if err = d.Set("bgp_seq", flattenSystemSdnVpnBgpSeq(o["bgp-seq"], d, "bgp_seq")); err != nil {
		if vv, ok := fortiAPIPatch(o["bgp-seq"], "SystemSdnVpn-BgpSeq"); ok {
			if err = d.Set("bgp_seq", vv); err != nil {
				return fmt.Errorf("Error reading bgp_seq: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bgp_seq: %v", err)
		}
	}

	if err = d.Set("cgw_gateway", flattenSystemSdnVpnCgwGateway(o["cgw-gateway"], d, "cgw_gateway")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgw-gateway"], "SystemSdnVpn-CgwGateway"); ok {
			if err = d.Set("cgw_gateway", vv); err != nil {
				return fmt.Errorf("Error reading cgw_gateway: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgw_gateway: %v", err)
		}
	}

	if err = d.Set("cgw_name", flattenSystemSdnVpnCgwName(o["cgw-name"], d, "cgw_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgw-name"], "SystemSdnVpn-CgwName"); ok {
			if err = d.Set("cgw_name", vv); err != nil {
				return fmt.Errorf("Error reading cgw_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgw_name: %v", err)
		}
	}

	if err = d.Set("cgw_id", flattenSystemSdnVpnCgwId(o["cgw_id"], d, "cgw_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgw_id"], "SystemSdnVpn-CgwId"); ok {
			if err = d.Set("cgw_id", vv); err != nil {
				return fmt.Errorf("Error reading cgw_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgw_id: %v", err)
		}
	}

	if err = d.Set("code", flattenSystemSdnVpnCode(o["code"], d, "code")); err != nil {
		if vv, ok := fortiAPIPatch(o["code"], "SystemSdnVpn-Code"); ok {
			if err = d.Set("code", vv); err != nil {
				return fmt.Errorf("Error reading code: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading code: %v", err)
		}
	}

	if err = d.Set("internal_interface", flattenSystemSdnVpnInternalInterface(o["internal-interface"], d, "internal_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-interface"], "SystemSdnVpn-InternalInterface"); ok {
			if err = d.Set("internal_interface", vv); err != nil {
				return fmt.Errorf("Error reading internal_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_interface: %v", err)
		}
	}

	if err = d.Set("local_cidr", flattenSystemSdnVpnLocalCidr(o["local-cidr"], d, "local_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["local-cidr"], "SystemSdnVpn-LocalCidr"); ok {
			if err = d.Set("local_cidr", vv); err != nil {
				return fmt.Errorf("Error reading local_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading local_cidr: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSdnVpnName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemSdnVpn-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat_traversal", flattenSystemSdnVpnNatTraversal(o["nat-traversal"], d, "nat_traversal")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat-traversal"], "SystemSdnVpn-NatTraversal"); ok {
			if err = d.Set("nat_traversal", vv); err != nil {
				return fmt.Errorf("Error reading nat_traversal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat_traversal: %v", err)
		}
	}

	if err = d.Set("remote_cidr", flattenSystemSdnVpnRemoteCidr(o["remote-cidr"], d, "remote_cidr")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-cidr"], "SystemSdnVpn-RemoteCidr"); ok {
			if err = d.Set("remote_cidr", vv); err != nil {
				return fmt.Errorf("Error reading remote_cidr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_cidr: %v", err)
		}
	}

	if err = d.Set("remote_type", flattenSystemSdnVpnRemoteType(o["remote-type"], d, "remote_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-type"], "SystemSdnVpn-RemoteType"); ok {
			if err = d.Set("remote_type", vv); err != nil {
				return fmt.Errorf("Error reading remote_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_type: %v", err)
		}
	}

	if err = d.Set("routing_type", flattenSystemSdnVpnRoutingType(o["routing-type"], d, "routing_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["routing-type"], "SystemSdnVpn-RoutingType"); ok {
			if err = d.Set("routing_type", vv); err != nil {
				return fmt.Errorf("Error reading routing_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading routing_type: %v", err)
		}
	}

	if err = d.Set("sdn", flattenSystemSdnVpnSdn(o["sdn"], d, "sdn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sdn"], "SystemSdnVpn-Sdn"); ok {
			if err = d.Set("sdn", vv); err != nil {
				return fmt.Errorf("Error reading sdn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdnVpnStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSdnVpn-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("subnet_id", flattenSystemSdnVpnSubnetId(o["subnet-id"], d, "subnet_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["subnet-id"], "SystemSdnVpn-SubnetId"); ok {
			if err = d.Set("subnet_id", vv); err != nil {
				return fmt.Errorf("Error reading subnet_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading subnet_id: %v", err)
		}
	}

	if err = d.Set("tgw_id", flattenSystemSdnVpnTgwId(o["tgw-id"], d, "tgw_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgw-id"], "SystemSdnVpn-TgwId"); ok {
			if err = d.Set("tgw_id", vv); err != nil {
				return fmt.Errorf("Error reading tgw_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgw_id: %v", err)
		}
	}

	if err = d.Set("tgw_vpn_rtbl_id", flattenSystemSdnVpnTgwVpnRtblId(o["tgw_vpn_rtbl_id"], d, "tgw_vpn_rtbl_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tgw_vpn_rtbl_id"], "SystemSdnVpn-TgwVpnRtblId"); ok {
			if err = d.Set("tgw_vpn_rtbl_id", vv); err != nil {
				return fmt.Errorf("Error reading tgw_vpn_rtbl_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tgw_vpn_rtbl_id: %v", err)
		}
	}

	if err = d.Set("trtbl_attachment", flattenSystemSdnVpnTrtblAttachment(o["trtbl_attachment"], d, "trtbl_attachment")); err != nil {
		if vv, ok := fortiAPIPatch(o["trtbl_attachment"], "SystemSdnVpn-TrtblAttachment"); ok {
			if err = d.Set("trtbl_attachment", vv); err != nil {
				return fmt.Errorf("Error reading trtbl_attachment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trtbl_attachment: %v", err)
		}
	}

	if err = d.Set("tunnel_interface", flattenSystemSdnVpnTunnelInterface(o["tunnel-interface"], d, "tunnel_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["tunnel-interface"], "SystemSdnVpn-TunnelInterface"); ok {
			if err = d.Set("tunnel_interface", vv); err != nil {
				return fmt.Errorf("Error reading tunnel_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tunnel_interface: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSdnVpnType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemSdnVpn-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("vgw_id", flattenSystemSdnVpnVgwId(o["vgw-id"], d, "vgw_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vgw-id"], "SystemSdnVpn-VgwId"); ok {
			if err = d.Set("vgw_id", vv); err != nil {
				return fmt.Errorf("Error reading vgw_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vgw_id: %v", err)
		}
	}

	if err = d.Set("vpn_id", flattenSystemSdnVpnVpnId(o["vpn_id"], d, "vpn_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn_id"], "SystemSdnVpn-VpnId"); ok {
			if err = d.Set("vpn_id", vv); err != nil {
				return fmt.Errorf("Error reading vpn_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_id: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnVpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdnVpnBgpAs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnBgpFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnBgpSeq(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCgwGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCgwName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCgwId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnInternalInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnVpnLocalCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemSdnVpnName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnNatTraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnPsksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnVpnRemoteCidr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemSdnVpnRemoteType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnRoutingType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnSdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnVpnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnSubnetId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnTgwId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnTgwVpnRtblId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnTrtblAttachment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnTunnelInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdnVpnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnVgwId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnVpnVpnId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnVpn(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bgp_as"); ok || d.HasChange("bgp_as") {
		t, err := expandSystemSdnVpnBgpAs(d, v, "bgp_as")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-as"] = t
		}
	}

	if v, ok := d.GetOk("bgp_from"); ok || d.HasChange("bgp_from") {
		t, err := expandSystemSdnVpnBgpFrom(d, v, "bgp_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-from"] = t
		}
	}

	if v, ok := d.GetOk("bgp_seq"); ok || d.HasChange("bgp_seq") {
		t, err := expandSystemSdnVpnBgpSeq(d, v, "bgp_seq")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp-seq"] = t
		}
	}

	if v, ok := d.GetOk("cgw_gateway"); ok || d.HasChange("cgw_gateway") {
		t, err := expandSystemSdnVpnCgwGateway(d, v, "cgw_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgw-gateway"] = t
		}
	}

	if v, ok := d.GetOk("cgw_name"); ok || d.HasChange("cgw_name") {
		t, err := expandSystemSdnVpnCgwName(d, v, "cgw_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgw-name"] = t
		}
	}

	if v, ok := d.GetOk("cgw_id"); ok || d.HasChange("cgw_id") {
		t, err := expandSystemSdnVpnCgwId(d, v, "cgw_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgw_id"] = t
		}
	}

	if v, ok := d.GetOk("code"); ok || d.HasChange("code") {
		t, err := expandSystemSdnVpnCode(d, v, "code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["code"] = t
		}
	}

	if v, ok := d.GetOk("internal_interface"); ok || d.HasChange("internal_interface") {
		t, err := expandSystemSdnVpnInternalInterface(d, v, "internal_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-interface"] = t
		}
	}

	if v, ok := d.GetOk("local_cidr"); ok || d.HasChange("local_cidr") {
		t, err := expandSystemSdnVpnLocalCidr(d, v, "local_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cidr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemSdnVpnName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat_traversal"); ok || d.HasChange("nat_traversal") {
		t, err := expandSystemSdnVpnNatTraversal(d, v, "nat_traversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-traversal"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok || d.HasChange("psksecret") {
		t, err := expandSystemSdnVpnPsksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("remote_cidr"); ok || d.HasChange("remote_cidr") {
		t, err := expandSystemSdnVpnRemoteCidr(d, v, "remote_cidr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-cidr"] = t
		}
	}

	if v, ok := d.GetOk("remote_type"); ok || d.HasChange("remote_type") {
		t, err := expandSystemSdnVpnRemoteType(d, v, "remote_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-type"] = t
		}
	}

	if v, ok := d.GetOk("routing_type"); ok || d.HasChange("routing_type") {
		t, err := expandSystemSdnVpnRoutingType(d, v, "routing_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["routing-type"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok || d.HasChange("sdn") {
		t, err := expandSystemSdnVpnSdn(d, v, "sdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSdnVpnStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("subnet_id"); ok || d.HasChange("subnet_id") {
		t, err := expandSystemSdnVpnSubnetId(d, v, "subnet_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-id"] = t
		}
	}

	if v, ok := d.GetOk("tgw_id"); ok || d.HasChange("tgw_id") {
		t, err := expandSystemSdnVpnTgwId(d, v, "tgw_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgw-id"] = t
		}
	}

	if v, ok := d.GetOk("tgw_vpn_rtbl_id"); ok || d.HasChange("tgw_vpn_rtbl_id") {
		t, err := expandSystemSdnVpnTgwVpnRtblId(d, v, "tgw_vpn_rtbl_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tgw_vpn_rtbl_id"] = t
		}
	}

	if v, ok := d.GetOk("trtbl_attachment"); ok || d.HasChange("trtbl_attachment") {
		t, err := expandSystemSdnVpnTrtblAttachment(d, v, "trtbl_attachment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trtbl_attachment"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_interface"); ok || d.HasChange("tunnel_interface") {
		t, err := expandSystemSdnVpnTunnelInterface(d, v, "tunnel_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-interface"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemSdnVpnType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("vgw_id"); ok || d.HasChange("vgw_id") {
		t, err := expandSystemSdnVpnVgwId(d, v, "vgw_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vgw-id"] = t
		}
	}

	if v, ok := d.GetOk("vpn_id"); ok || d.HasChange("vpn_id") {
		t, err := expandSystemSdnVpnVpnId(d, v, "vpn_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn_id"] = t
		}
	}

	return &obj, nil
}
