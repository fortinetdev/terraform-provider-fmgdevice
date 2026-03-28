// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device IcapLocalServer

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIcapLocalServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapLocalServerCreate,
		Read:   resourceIcapLocalServerRead,
		Update: resourceIcapLocalServerUpdate,
		Delete: resourceIcapLocalServerDelete,

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
			"icap_incoming_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icap_incoming_ssl_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icap_server_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"icap_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"av_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dlp_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dlp_sensor": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"extension_headers": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"image_analyzer_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"profile_protocol_options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"service_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"webfilter_profile": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ipv6": &schema.Schema{
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
			"message_preview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secure_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_scheme_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceIcapLocalServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapLocalServer(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapLocalServer resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("icap_server_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadIcapLocalServer(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateIcapLocalServer(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating IcapLocalServer resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateIcapLocalServer(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating IcapLocalServer resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "icap_server_id")))

	return resourceIcapLocalServerRead(d, m)
}

func resourceIcapLocalServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIcapLocalServer(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapLocalServer resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateIcapLocalServer(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating IcapLocalServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "icap_server_id")))

	return resourceIcapLocalServerRead(d, m)
}

func resourceIcapLocalServerDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteIcapLocalServer(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting IcapLocalServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapLocalServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIcapLocalServer(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading IcapLocalServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapLocalServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapLocalServer resource from API: %v", err)
	}
	return nil
}

func flattenIcapLocalServerIcapIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapIncomingSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "av_profile"
		if _, ok := i["av-profile"]; ok {
			v := flattenIcapLocalServerIcapServiceAvProfile(i["av-profile"], d, pre_append)
			tmp["av_profile"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-AvProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dlp_profile"
		if _, ok := i["dlp-profile"]; ok {
			v := flattenIcapLocalServerIcapServiceDlpProfile(i["dlp-profile"], d, pre_append)
			tmp["dlp_profile"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-DlpProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dlp_sensor"
		if _, ok := i["dlp-sensor"]; ok {
			v := flattenIcapLocalServerIcapServiceDlpSensor(i["dlp-sensor"], d, pre_append)
			tmp["dlp_sensor"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-DlpSensor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extension_headers"
		if _, ok := i["extension-headers"]; ok {
			v := flattenIcapLocalServerIcapServiceExtensionHeaders(i["extension-headers"], d, pre_append)
			tmp["extension_headers"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-ExtensionHeaders")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image_analyzer_profile"
		if _, ok := i["image-analyzer-profile"]; ok {
			v := flattenIcapLocalServerIcapServiceImageAnalyzerProfile(i["image-analyzer-profile"], d, pre_append)
			tmp["image_analyzer_profile"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-ImageAnalyzerProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenIcapLocalServerIcapServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_protocol_options"
		if _, ok := i["profile-protocol-options"]; ok {
			v := flattenIcapLocalServerIcapServiceProfileProtocolOptions(i["profile-protocol-options"], d, pre_append)
			tmp["profile_protocol_options"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-ProfileProtocolOptions")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenIcapLocalServerIcapServiceServiceId(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-ServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "webfilter_profile"
		if _, ok := i["webfilter-profile"]; ok {
			v := flattenIcapLocalServerIcapServiceWebfilterProfile(i["webfilter-profile"], d, pre_append)
			tmp["webfilter_profile"] = fortiAPISubPartPatch(v, "IcapLocalServer-IcapService-WebfilterProfile")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenIcapLocalServerIcapServiceAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceExtensionHeaders(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceImageAnalyzerProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapServiceProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIcapServiceServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIcapServiceWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerIncomingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerIncomingIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerMessagePreview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerSecureConnection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerSslCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenIcapLocalServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerStatusIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerStrictSchemeCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapLocalServerTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIcapLocalServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("icap_incoming_port", flattenIcapLocalServerIcapIncomingPort(o["icap-incoming-port"], d, "icap_incoming_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-incoming-port"], "IcapLocalServer-IcapIncomingPort"); ok {
			if err = d.Set("icap_incoming_port", vv); err != nil {
				return fmt.Errorf("Error reading icap_incoming_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_incoming_port: %v", err)
		}
	}

	if err = d.Set("icap_incoming_ssl_port", flattenIcapLocalServerIcapIncomingSslPort(o["icap-incoming-ssl-port"], d, "icap_incoming_ssl_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-incoming-ssl-port"], "IcapLocalServer-IcapIncomingSslPort"); ok {
			if err = d.Set("icap_incoming_ssl_port", vv); err != nil {
				return fmt.Errorf("Error reading icap_incoming_ssl_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_incoming_ssl_port: %v", err)
		}
	}

	if err = d.Set("icap_server_id", flattenIcapLocalServerIcapServerId(o["icap-server-id"], d, "icap_server_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-server-id"], "IcapLocalServer-IcapServerId"); ok {
			if err = d.Set("icap_server_id", vv); err != nil {
				return fmt.Errorf("Error reading icap_server_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_server_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("icap_service", flattenIcapLocalServerIcapService(o["icap-service"], d, "icap_service")); err != nil {
			if vv, ok := fortiAPIPatch(o["icap-service"], "IcapLocalServer-IcapService"); ok {
				if err = d.Set("icap_service", vv); err != nil {
					return fmt.Errorf("Error reading icap_service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icap_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap_service"); ok {
			if err = d.Set("icap_service", flattenIcapLocalServerIcapService(o["icap-service"], d, "icap_service")); err != nil {
				if vv, ok := fortiAPIPatch(o["icap-service"], "IcapLocalServer-IcapService"); ok {
					if err = d.Set("icap_service", vv); err != nil {
						return fmt.Errorf("Error reading icap_service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icap_service: %v", err)
				}
			}
		}
	}

	if err = d.Set("incoming_ip", flattenIcapLocalServerIncomingIp(o["incoming-ip"], d, "incoming_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ip"], "IcapLocalServer-IncomingIp"); ok {
			if err = d.Set("incoming_ip", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("incoming_ipv6", flattenIcapLocalServerIncomingIpv6(o["incoming-ipv6"], d, "incoming_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["incoming-ipv6"], "IcapLocalServer-IncomingIpv6"); ok {
			if err = d.Set("incoming_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading incoming_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incoming_ipv6: %v", err)
		}
	}

	if err = d.Set("interface", flattenIcapLocalServerInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "IcapLocalServer-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("message_preview", flattenIcapLocalServerMessagePreview(o["message-preview"], d, "message_preview")); err != nil {
		if vv, ok := fortiAPIPatch(o["message-preview"], "IcapLocalServer-MessagePreview"); ok {
			if err = d.Set("message_preview", vv); err != nil {
				return fmt.Errorf("Error reading message_preview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading message_preview: %v", err)
		}
	}

	if err = d.Set("secure_connection", flattenIcapLocalServerSecureConnection(o["secure-connection"], d, "secure_connection")); err != nil {
		if vv, ok := fortiAPIPatch(o["secure-connection"], "IcapLocalServer-SecureConnection"); ok {
			if err = d.Set("secure_connection", vv); err != nil {
				return fmt.Errorf("Error reading secure_connection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secure_connection: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenIcapLocalServerSrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "IcapLocalServer-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("ssl_cert", flattenIcapLocalServerSslCert(o["ssl-cert"], d, "ssl_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-cert"], "IcapLocalServer-SslCert"); ok {
			if err = d.Set("ssl_cert", vv); err != nil {
				return fmt.Errorf("Error reading ssl_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("status", flattenIcapLocalServerStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "IcapLocalServer-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("status_ipv6", flattenIcapLocalServerStatusIpv6(o["status-ipv6"], d, "status_ipv6")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-ipv6"], "IcapLocalServer-StatusIpv6"); ok {
			if err = d.Set("status_ipv6", vv); err != nil {
				return fmt.Errorf("Error reading status_ipv6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_ipv6: %v", err)
		}
	}

	if err = d.Set("strict_scheme_check", flattenIcapLocalServerStrictSchemeCheck(o["strict-scheme-check"], d, "strict_scheme_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["strict-scheme-check"], "IcapLocalServer-StrictSchemeCheck"); ok {
			if err = d.Set("strict_scheme_check", vv); err != nil {
				return fmt.Errorf("Error reading strict_scheme_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strict_scheme_check: %v", err)
		}
	}

	if err = d.Set("timeout", flattenIcapLocalServerTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "IcapLocalServer-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenIcapLocalServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIcapLocalServerIcapIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapIncomingSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "av_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["av-profile"], _ = expandIcapLocalServerIcapServiceAvProfile(d, i["av_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dlp_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dlp-profile"], _ = expandIcapLocalServerIcapServiceDlpProfile(d, i["dlp_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dlp_sensor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dlp-sensor"], _ = expandIcapLocalServerIcapServiceDlpSensor(d, i["dlp_sensor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extension_headers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["extension-headers"], _ = expandIcapLocalServerIcapServiceExtensionHeaders(d, i["extension_headers"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "image_analyzer_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["image-analyzer-profile"], _ = expandIcapLocalServerIcapServiceImageAnalyzerProfile(d, i["image_analyzer_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandIcapLocalServerIcapServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "profile_protocol_options"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["profile-protocol-options"], _ = expandIcapLocalServerIcapServiceProfileProtocolOptions(d, i["profile_protocol_options"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandIcapLocalServerIcapServiceServiceId(d, i["service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "webfilter_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["webfilter-profile"], _ = expandIcapLocalServerIcapServiceWebfilterProfile(d, i["webfilter_profile"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandIcapLocalServerIcapServiceAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceExtensionHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceImageAnalyzerProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapServiceProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIcapServiceServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIcapServiceWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerIncomingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerIncomingIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerMessagePreview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerSecureConnection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerSslCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandIcapLocalServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerStatusIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerStrictSchemeCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapLocalServerTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIcapLocalServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("icap_incoming_port"); ok || d.HasChange("icap_incoming_port") {
		t, err := expandIcapLocalServerIcapIncomingPort(d, v, "icap_incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("icap_incoming_ssl_port"); ok || d.HasChange("icap_incoming_ssl_port") {
		t, err := expandIcapLocalServerIcapIncomingSslPort(d, v, "icap_incoming_ssl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-incoming-ssl-port"] = t
		}
	}

	if v, ok := d.GetOk("icap_server_id"); ok || d.HasChange("icap_server_id") {
		t, err := expandIcapLocalServerIcapServerId(d, v, "icap_server_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-server-id"] = t
		}
	}

	if v, ok := d.GetOk("icap_service"); ok || d.HasChange("icap_service") {
		t, err := expandIcapLocalServerIcapService(d, v, "icap_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-service"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok || d.HasChange("incoming_ip") {
		t, err := expandIcapLocalServerIncomingIp(d, v, "incoming_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ipv6"); ok || d.HasChange("incoming_ipv6") {
		t, err := expandIcapLocalServerIncomingIpv6(d, v, "incoming_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandIcapLocalServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("message_preview"); ok || d.HasChange("message_preview") {
		t, err := expandIcapLocalServerMessagePreview(d, v, "message_preview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-preview"] = t
		}
	}

	if v, ok := d.GetOk("secure_connection"); ok || d.HasChange("secure_connection") {
		t, err := expandIcapLocalServerSecureConnection(d, v, "secure_connection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-connection"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandIcapLocalServerSrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cert"); ok || d.HasChange("ssl_cert") {
		t, err := expandIcapLocalServerSslCert(d, v, "ssl_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandIcapLocalServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("status_ipv6"); ok || d.HasChange("status_ipv6") {
		t, err := expandIcapLocalServerStatusIpv6(d, v, "status_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("strict_scheme_check"); ok || d.HasChange("strict_scheme_check") {
		t, err := expandIcapLocalServerStrictSchemeCheck(d, v, "strict_scheme_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-scheme-check"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandIcapLocalServerTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
