// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure FortiClient Enterprise Management Server (EMS) entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEndpointControlFctems() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlFctemsCreate,
		Read:   resourceEndpointControlFctemsRead,
		Update: resourceEndpointControlFctemsUpdate,
		Delete: resourceEndpointControlFctemsDelete,

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
			"call_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"capabilities": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"certificate_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_authentication_access_key": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"dirty_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ems_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"fortinetone_cloud_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"out_of_sync_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pull_avatars": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_malware_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_sysinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_tags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pull_vulnerabilities": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_tags_to_all_vdoms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status_check_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_ca_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verified_cn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"verifying_ca": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"websocket_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cn_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preserve_ssl_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEndpointControlFctemsCreate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectEndpointControlFctems(d)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlFctems resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("ems_id")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadEndpointControlFctems(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateEndpointControlFctems(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating EndpointControlFctems resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateEndpointControlFctems(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating EndpointControlFctems resource: %v", err)
		}

	}

	d.SetId(strconv.Itoa(getIntKey(d, "ems_id")))

	return resourceEndpointControlFctemsRead(d, m)
}

func resourceEndpointControlFctemsUpdate(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	obj, err := getObjectEndpointControlFctems(d)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlFctems resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateEndpointControlFctems(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlFctems resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "ems_id")))

	return resourceEndpointControlFctemsRead(d, m)
}

func resourceEndpointControlFctemsDelete(d *schema.ResourceData, m interface{}) error {
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
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteEndpointControlFctems(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlFctems resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlFctemsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEndpointControlFctems(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading EndpointControlFctems resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlFctems(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlFctems resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlFctemsCallTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEndpointControlFctemsCapabilities(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEndpointControlFctemsCertificateFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsDirtyReason(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsEmsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsFortinetoneCloudAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEndpointControlFctemsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsOutOfSyncThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullAvatars(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullMalwareHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullSysinfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullTags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsPullVulnerabilities(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsSendTagsToAllVdoms(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsStatusCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsTrustCaCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsVerifiedCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsVerifyingCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenEndpointControlFctemsWebsocketOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsCaCnInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsCloudServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlFctemsPreserveSslSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEndpointControlFctems(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("call_timeout", flattenEndpointControlFctemsCallTimeout(o["call-timeout"], d, "call_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-timeout"], "EndpointControlFctems-CallTimeout"); ok {
			if err = d.Set("call_timeout", vv); err != nil {
				return fmt.Errorf("Error reading call_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_timeout: %v", err)
		}
	}

	if err = d.Set("certificate", flattenEndpointControlFctemsCertificate(o["certificate"], d, "certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate"], "EndpointControlFctems-Certificate"); ok {
			if err = d.Set("certificate", vv); err != nil {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("capabilities", flattenEndpointControlFctemsCapabilities(o["capabilities"], d, "capabilities")); err != nil {
		if vv, ok := fortiAPIPatch(o["capabilities"], "EndpointControlFctems-Capabilities"); ok {
			if err = d.Set("capabilities", vv); err != nil {
				return fmt.Errorf("Error reading capabilities: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading capabilities: %v", err)
		}
	}

	if err = d.Set("certificate_fingerprint", flattenEndpointControlFctemsCertificateFingerprint(o["certificate-fingerprint"], d, "certificate_fingerprint")); err != nil {
		if vv, ok := fortiAPIPatch(o["certificate-fingerprint"], "EndpointControlFctems-CertificateFingerprint"); ok {
			if err = d.Set("certificate_fingerprint", vv); err != nil {
				return fmt.Errorf("Error reading certificate_fingerprint: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading certificate_fingerprint: %v", err)
		}
	}

	if err = d.Set("dirty_reason", flattenEndpointControlFctemsDirtyReason(o["dirty-reason"], d, "dirty_reason")); err != nil {
		if vv, ok := fortiAPIPatch(o["dirty-reason"], "EndpointControlFctems-DirtyReason"); ok {
			if err = d.Set("dirty_reason", vv); err != nil {
				return fmt.Errorf("Error reading dirty_reason: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dirty_reason: %v", err)
		}
	}

	if err = d.Set("ems_id", flattenEndpointControlFctemsEmsId(o["ems-id"], d, "ems_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-id"], "EndpointControlFctems-EmsId"); ok {
			if err = d.Set("ems_id", vv); err != nil {
				return fmt.Errorf("Error reading ems_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_id: %v", err)
		}
	}

	if err = d.Set("fortinetone_cloud_authentication", flattenEndpointControlFctemsFortinetoneCloudAuthentication(o["fortinetone-cloud-authentication"], d, "fortinetone_cloud_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortinetone-cloud-authentication"], "EndpointControlFctems-FortinetoneCloudAuthentication"); ok {
			if err = d.Set("fortinetone_cloud_authentication", vv); err != nil {
				return fmt.Errorf("Error reading fortinetone_cloud_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortinetone_cloud_authentication: %v", err)
		}
	}

	if err = d.Set("https_port", flattenEndpointControlFctemsHttpsPort(o["https-port"], d, "https_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["https-port"], "EndpointControlFctems-HttpsPort"); ok {
			if err = d.Set("https_port", vv); err != nil {
				return fmt.Errorf("Error reading https_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("interface", flattenEndpointControlFctemsInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "EndpointControlFctems-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenEndpointControlFctemsInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "EndpointControlFctems-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("name", flattenEndpointControlFctemsName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "EndpointControlFctems-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("out_of_sync_threshold", flattenEndpointControlFctemsOutOfSyncThreshold(o["out-of-sync-threshold"], d, "out_of_sync_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["out-of-sync-threshold"], "EndpointControlFctems-OutOfSyncThreshold"); ok {
			if err = d.Set("out_of_sync_threshold", vv); err != nil {
				return fmt.Errorf("Error reading out_of_sync_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading out_of_sync_threshold: %v", err)
		}
	}

	if err = d.Set("pull_avatars", flattenEndpointControlFctemsPullAvatars(o["pull-avatars"], d, "pull_avatars")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-avatars"], "EndpointControlFctems-PullAvatars"); ok {
			if err = d.Set("pull_avatars", vv); err != nil {
				return fmt.Errorf("Error reading pull_avatars: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_avatars: %v", err)
		}
	}

	if err = d.Set("pull_malware_hash", flattenEndpointControlFctemsPullMalwareHash(o["pull-malware-hash"], d, "pull_malware_hash")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-malware-hash"], "EndpointControlFctems-PullMalwareHash"); ok {
			if err = d.Set("pull_malware_hash", vv); err != nil {
				return fmt.Errorf("Error reading pull_malware_hash: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_malware_hash: %v", err)
		}
	}

	if err = d.Set("pull_sysinfo", flattenEndpointControlFctemsPullSysinfo(o["pull-sysinfo"], d, "pull_sysinfo")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-sysinfo"], "EndpointControlFctems-PullSysinfo"); ok {
			if err = d.Set("pull_sysinfo", vv); err != nil {
				return fmt.Errorf("Error reading pull_sysinfo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_sysinfo: %v", err)
		}
	}

	if err = d.Set("pull_tags", flattenEndpointControlFctemsPullTags(o["pull-tags"], d, "pull_tags")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-tags"], "EndpointControlFctems-PullTags"); ok {
			if err = d.Set("pull_tags", vv); err != nil {
				return fmt.Errorf("Error reading pull_tags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_tags: %v", err)
		}
	}

	if err = d.Set("pull_vulnerabilities", flattenEndpointControlFctemsPullVulnerabilities(o["pull-vulnerabilities"], d, "pull_vulnerabilities")); err != nil {
		if vv, ok := fortiAPIPatch(o["pull-vulnerabilities"], "EndpointControlFctems-PullVulnerabilities"); ok {
			if err = d.Set("pull_vulnerabilities", vv); err != nil {
				return fmt.Errorf("Error reading pull_vulnerabilities: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading pull_vulnerabilities: %v", err)
		}
	}

	if err = d.Set("send_tags_to_all_vdoms", flattenEndpointControlFctemsSendTagsToAllVdoms(o["send-tags-to-all-vdoms"], d, "send_tags_to_all_vdoms")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-tags-to-all-vdoms"], "EndpointControlFctems-SendTagsToAllVdoms"); ok {
			if err = d.Set("send_tags_to_all_vdoms", vv); err != nil {
				return fmt.Errorf("Error reading send_tags_to_all_vdoms: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_tags_to_all_vdoms: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenEndpointControlFctemsSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if vv, ok := fortiAPIPatch(o["serial-number"], "EndpointControlFctems-SerialNumber"); ok {
			if err = d.Set("serial_number", vv); err != nil {
				return fmt.Errorf("Error reading serial_number: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("server", flattenEndpointControlFctemsServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "EndpointControlFctems-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenEndpointControlFctemsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "EndpointControlFctems-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("status_check_interval", flattenEndpointControlFctemsStatusCheckInterval(o["status-check-interval"], d, "status_check_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-check-interval"], "EndpointControlFctems-StatusCheckInterval"); ok {
			if err = d.Set("status_check_interval", vv); err != nil {
				return fmt.Errorf("Error reading status_check_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_check_interval: %v", err)
		}
	}

	if err = d.Set("status", flattenEndpointControlFctemsStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "EndpointControlFctems-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tenant_id", flattenEndpointControlFctemsTenantId(o["tenant-id"], d, "tenant_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["tenant-id"], "EndpointControlFctems-TenantId"); ok {
			if err = d.Set("tenant_id", vv); err != nil {
				return fmt.Errorf("Error reading tenant_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("trust_ca_cn", flattenEndpointControlFctemsTrustCaCn(o["trust-ca-cn"], d, "trust_ca_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["trust-ca-cn"], "EndpointControlFctems-TrustCaCn"); ok {
			if err = d.Set("trust_ca_cn", vv); err != nil {
				return fmt.Errorf("Error reading trust_ca_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trust_ca_cn: %v", err)
		}
	}

	if err = d.Set("verified_cn", flattenEndpointControlFctemsVerifiedCn(o["verified-cn"], d, "verified_cn")); err != nil {
		if vv, ok := fortiAPIPatch(o["verified-cn"], "EndpointControlFctems-VerifiedCn"); ok {
			if err = d.Set("verified_cn", vv); err != nil {
				return fmt.Errorf("Error reading verified_cn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verified_cn: %v", err)
		}
	}

	if err = d.Set("verifying_ca", flattenEndpointControlFctemsVerifyingCa(o["verifying-ca"], d, "verifying_ca")); err != nil {
		if vv, ok := fortiAPIPatch(o["verifying-ca"], "EndpointControlFctems-VerifyingCa"); ok {
			if err = d.Set("verifying_ca", vv); err != nil {
				return fmt.Errorf("Error reading verifying_ca: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verifying_ca: %v", err)
		}
	}

	if err = d.Set("websocket_override", flattenEndpointControlFctemsWebsocketOverride(o["websocket-override"], d, "websocket_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["websocket-override"], "EndpointControlFctems-WebsocketOverride"); ok {
			if err = d.Set("websocket_override", vv); err != nil {
				return fmt.Errorf("Error reading websocket_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading websocket_override: %v", err)
		}
	}

	if err = d.Set("ca_cn_info", flattenEndpointControlFctemsCaCnInfo(o["ca-cn-info"], d, "ca_cn_info")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cn-info"], "EndpointControlFctems-CaCnInfo"); ok {
			if err = d.Set("ca_cn_info", vv); err != nil {
				return fmt.Errorf("Error reading ca_cn_info: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cn_info: %v", err)
		}
	}

	if err = d.Set("cloud_server_type", flattenEndpointControlFctemsCloudServerType(o["cloud-server-type"], d, "cloud_server_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["cloud-server-type"], "EndpointControlFctems-CloudServerType"); ok {
			if err = d.Set("cloud_server_type", vv); err != nil {
				return fmt.Errorf("Error reading cloud_server_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cloud_server_type: %v", err)
		}
	}

	if err = d.Set("preserve_ssl_session", flattenEndpointControlFctemsPreserveSslSession(o["preserve-ssl-session"], d, "preserve_ssl_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["preserve-ssl-session"], "EndpointControlFctems-PreserveSslSession"); ok {
			if err = d.Set("preserve_ssl_session", vv); err != nil {
				return fmt.Errorf("Error reading preserve_ssl_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preserve_ssl_session: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlFctemsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEndpointControlFctemsCallTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEndpointControlFctemsCapabilities(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEndpointControlFctemsCertificateFingerprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCloudAuthenticationAccessKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEndpointControlFctemsDirtyReason(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsEmsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsFortinetoneCloudAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEndpointControlFctemsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsOutOfSyncThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullAvatars(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullMalwareHash(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullSysinfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPullVulnerabilities(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsSendTagsToAllVdoms(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsStatusCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsTenantId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsTrustCaCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsVerifiedCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsVerifyingCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandEndpointControlFctemsWebsocketOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCaCnInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsCloudServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlFctemsPreserveSslSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlFctems(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("call_timeout"); ok || d.HasChange("call_timeout") {
		t, err := expandEndpointControlFctemsCallTimeout(d, v, "call_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-timeout"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandEndpointControlFctemsCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("capabilities"); ok || d.HasChange("capabilities") {
		t, err := expandEndpointControlFctemsCapabilities(d, v, "capabilities")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capabilities"] = t
		}
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok || d.HasChange("certificate_fingerprint") {
		t, err := expandEndpointControlFctemsCertificateFingerprint(d, v, "certificate_fingerprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-fingerprint"] = t
		}
	}

	if v, ok := d.GetOk("cloud_authentication_access_key"); ok || d.HasChange("cloud_authentication_access_key") {
		t, err := expandEndpointControlFctemsCloudAuthenticationAccessKey(d, v, "cloud_authentication_access_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-authentication-access-key"] = t
		}
	}

	if v, ok := d.GetOk("dirty_reason"); ok || d.HasChange("dirty_reason") {
		t, err := expandEndpointControlFctemsDirtyReason(d, v, "dirty_reason")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dirty-reason"] = t
		}
	}

	if v, ok := d.GetOk("ems_id"); ok || d.HasChange("ems_id") {
		t, err := expandEndpointControlFctemsEmsId(d, v, "ems_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-id"] = t
		}
	}

	if v, ok := d.GetOk("fortinetone_cloud_authentication"); ok || d.HasChange("fortinetone_cloud_authentication") {
		t, err := expandEndpointControlFctemsFortinetoneCloudAuthentication(d, v, "fortinetone_cloud_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinetone-cloud-authentication"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok || d.HasChange("https_port") {
		t, err := expandEndpointControlFctemsHttpsPort(d, v, "https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandEndpointControlFctemsInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandEndpointControlFctemsInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandEndpointControlFctemsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("out_of_sync_threshold"); ok || d.HasChange("out_of_sync_threshold") {
		t, err := expandEndpointControlFctemsOutOfSyncThreshold(d, v, "out_of_sync_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["out-of-sync-threshold"] = t
		}
	}

	if v, ok := d.GetOk("pull_avatars"); ok || d.HasChange("pull_avatars") {
		t, err := expandEndpointControlFctemsPullAvatars(d, v, "pull_avatars")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-avatars"] = t
		}
	}

	if v, ok := d.GetOk("pull_malware_hash"); ok || d.HasChange("pull_malware_hash") {
		t, err := expandEndpointControlFctemsPullMalwareHash(d, v, "pull_malware_hash")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-malware-hash"] = t
		}
	}

	if v, ok := d.GetOk("pull_sysinfo"); ok || d.HasChange("pull_sysinfo") {
		t, err := expandEndpointControlFctemsPullSysinfo(d, v, "pull_sysinfo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("pull_tags"); ok || d.HasChange("pull_tags") {
		t, err := expandEndpointControlFctemsPullTags(d, v, "pull_tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-tags"] = t
		}
	}

	if v, ok := d.GetOk("pull_vulnerabilities"); ok || d.HasChange("pull_vulnerabilities") {
		t, err := expandEndpointControlFctemsPullVulnerabilities(d, v, "pull_vulnerabilities")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pull-vulnerabilities"] = t
		}
	}

	if v, ok := d.GetOk("send_tags_to_all_vdoms"); ok || d.HasChange("send_tags_to_all_vdoms") {
		t, err := expandEndpointControlFctemsSendTagsToAllVdoms(d, v, "send_tags_to_all_vdoms")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-tags-to-all-vdoms"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok || d.HasChange("serial_number") {
		t, err := expandEndpointControlFctemsSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandEndpointControlFctemsServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandEndpointControlFctemsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("status_check_interval"); ok || d.HasChange("status_check_interval") {
		t, err := expandEndpointControlFctemsStatusCheckInterval(d, v, "status_check_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-check-interval"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandEndpointControlFctemsStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tenant_id"); ok || d.HasChange("tenant_id") {
		t, err := expandEndpointControlFctemsTenantId(d, v, "tenant_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	}

	if v, ok := d.GetOk("trust_ca_cn"); ok || d.HasChange("trust_ca_cn") {
		t, err := expandEndpointControlFctemsTrustCaCn(d, v, "trust_ca_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ca-cn"] = t
		}
	}

	if v, ok := d.GetOk("verified_cn"); ok || d.HasChange("verified_cn") {
		t, err := expandEndpointControlFctemsVerifiedCn(d, v, "verified_cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verified-cn"] = t
		}
	}

	if v, ok := d.GetOk("verifying_ca"); ok || d.HasChange("verifying_ca") {
		t, err := expandEndpointControlFctemsVerifyingCa(d, v, "verifying_ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verifying-ca"] = t
		}
	}

	if v, ok := d.GetOk("websocket_override"); ok || d.HasChange("websocket_override") {
		t, err := expandEndpointControlFctemsWebsocketOverride(d, v, "websocket_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websocket-override"] = t
		}
	}

	if v, ok := d.GetOk("ca_cn_info"); ok || d.HasChange("ca_cn_info") {
		t, err := expandEndpointControlFctemsCaCnInfo(d, v, "ca_cn_info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cn-info"] = t
		}
	}

	if v, ok := d.GetOk("cloud_server_type"); ok || d.HasChange("cloud_server_type") {
		t, err := expandEndpointControlFctemsCloudServerType(d, v, "cloud_server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cloud-server-type"] = t
		}
	}

	if v, ok := d.GetOk("preserve_ssl_session"); ok || d.HasChange("preserve_ssl_session") {
		t, err := expandEndpointControlFctemsPreserveSslSession(d, v, "preserve_ssl_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-ssl-session"] = t
		}
	}

	return &obj, nil
}
