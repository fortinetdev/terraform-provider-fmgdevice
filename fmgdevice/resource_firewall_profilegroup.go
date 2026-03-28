// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure profile groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileGroupCreate,
		Read:   resourceFirewallProfileGroupRead,
		Update: resourceFirewallProfileGroupUpdate,
		Delete: resourceFirewallProfileGroupDelete,

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
			"application_list": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"casb_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"diameter_filter_profile": &schema.Schema{
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
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ips_voip_filter": &schema.Schema{
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
			"sctp_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"telemetry_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"videofilter_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"virtual_patch_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
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
			"ia_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"isolator_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"redirect_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallProfileGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProfileGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadFirewallProfileGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateFirewallProfileGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating FirewallProfileGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateFirewallProfileGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating FirewallProfileGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProfileGroupRead(d, m)
}

func resourceFirewallProfileGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProfileGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceFirewallProfileGroupRead(d, m)
}

func resourceFirewallProfileGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallProfileGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProfileGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileGroup resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileGroupApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupCasbProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupDiameterFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupDlpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupIpsVoipFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileGroupProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupSctpFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupSshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupSslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupTelemetryProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupIaProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupIsolatorProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileGroupRedirectProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallProfileGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("application_list", flattenFirewallProfileGroupApplicationList(o["application-list"], d, "application_list")); err != nil {
		if vv, ok := fortiAPIPatch(o["application-list"], "FirewallProfileGroup-ApplicationList"); ok {
			if err = d.Set("application_list", vv); err != nil {
				return fmt.Errorf("Error reading application_list: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallProfileGroupAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-profile"], "FirewallProfileGroup-AvProfile"); ok {
			if err = d.Set("av_profile", vv); err != nil {
				return fmt.Errorf("Error reading av_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallProfileGroupCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["cifs-profile"], "FirewallProfileGroup-CifsProfile"); ok {
			if err = d.Set("cifs_profile", vv); err != nil {
				return fmt.Errorf("Error reading cifs_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallProfileGroupCasbProfile(o["casb-profile"], d, "casb_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["casb-profile"], "FirewallProfileGroup-CasbProfile"); ok {
			if err = d.Set("casb_profile", vv); err != nil {
				return fmt.Errorf("Error reading casb_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallProfileGroupDiameterFilterProfile(o["diameter-filter-profile"], d, "diameter_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["diameter-filter-profile"], "FirewallProfileGroup-DiameterFilterProfile"); ok {
			if err = d.Set("diameter_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallProfileGroupDlpProfile(o["dlp-profile"], d, "dlp_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-profile"], "FirewallProfileGroup-DlpProfile"); ok {
			if err = d.Set("dlp_profile", vv); err != nil {
				return fmt.Errorf("Error reading dlp_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallProfileGroupDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnsfilter-profile"], "FirewallProfileGroup-DnsfilterProfile"); ok {
			if err = d.Set("dnsfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallProfileGroupEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["emailfilter-profile"], "FirewallProfileGroup-EmailfilterProfile"); ok {
			if err = d.Set("emailfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading emailfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallProfileGroupFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["file-filter-profile"], "FirewallProfileGroup-FileFilterProfile"); ok {
			if err = d.Set("file_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading file_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallProfileGroupIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["icap-profile"], "FirewallProfileGroup-IcapProfile"); ok {
			if err = d.Set("icap_profile", vv); err != nil {
				return fmt.Errorf("Error reading icap_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallProfileGroupIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-sensor"], "FirewallProfileGroup-IpsSensor"); ok {
			if err = d.Set("ips_sensor", vv); err != nil {
				return fmt.Errorf("Error reading ips_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallProfileGroupIpsVoipFilter(o["ips-voip-filter"], d, "ips_voip_filter")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips-voip-filter"], "FirewallProfileGroup-IpsVoipFilter"); ok {
			if err = d.Set("ips_voip_filter", vv); err != nil {
				return fmt.Errorf("Error reading ips_voip_filter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallProfileGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "FirewallProfileGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallProfileGroupProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if vv, ok := fortiAPIPatch(o["profile-protocol-options"], "FirewallProfileGroup-ProfileProtocolOptions"); ok {
			if err = d.Set("profile_protocol_options", vv); err != nil {
				return fmt.Errorf("Error reading profile_protocol_options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallProfileGroupSctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["sctp-filter-profile"], "FirewallProfileGroup-SctpFilterProfile"); ok {
			if err = d.Set("sctp_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallProfileGroupSshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-filter-profile"], "FirewallProfileGroup-SshFilterProfile"); ok {
			if err = d.Set("ssh_filter_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallProfileGroupSslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssl-ssh-profile"], "FirewallProfileGroup-SslSshProfile"); ok {
			if err = d.Set("ssl_ssh_profile", vv); err != nil {
				return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("telemetry_profile", flattenFirewallProfileGroupTelemetryProfile(o["telemetry-profile"], d, "telemetry_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["telemetry-profile"], "FirewallProfileGroup-TelemetryProfile"); ok {
			if err = d.Set("telemetry_profile", vv); err != nil {
				return fmt.Errorf("Error reading telemetry_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading telemetry_profile: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallProfileGroupVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["videofilter-profile"], "FirewallProfileGroup-VideofilterProfile"); ok {
			if err = d.Set("videofilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading videofilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallProfileGroupVirtualPatchProfile(o["virtual-patch-profile"], d, "virtual_patch_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["virtual-patch-profile"], "FirewallProfileGroup-VirtualPatchProfile"); ok {
			if err = d.Set("virtual_patch_profile", vv); err != nil {
				return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallProfileGroupVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["voip-profile"], "FirewallProfileGroup-VoipProfile"); ok {
			if err = d.Set("voip_profile", vv); err != nil {
				return fmt.Errorf("Error reading voip_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallProfileGroupWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["waf-profile"], "FirewallProfileGroup-WafProfile"); ok {
			if err = d.Set("waf_profile", vv); err != nil {
				return fmt.Errorf("Error reading waf_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallProfileGroupWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["webfilter-profile"], "FirewallProfileGroup-WebfilterProfile"); ok {
			if err = d.Set("webfilter_profile", vv); err != nil {
				return fmt.Errorf("Error reading webfilter_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallProfileGroupDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if vv, ok := fortiAPIPatch(o["dlp-sensor"], "FirewallProfileGroup-DlpSensor"); ok {
			if err = d.Set("dlp_sensor", vv); err != nil {
				return fmt.Errorf("Error reading dlp_sensor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("ia_profile", flattenFirewallProfileGroupIaProfile(o["ia-profile"], d, "ia_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["ia-profile"], "FirewallProfileGroup-IaProfile"); ok {
			if err = d.Set("ia_profile", vv); err != nil {
				return fmt.Errorf("Error reading ia_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ia_profile: %v", err)
		}
	}

	if err = d.Set("isolator_profile", flattenFirewallProfileGroupIsolatorProfile(o["isolator-profile"], d, "isolator_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["isolator-profile"], "FirewallProfileGroup-IsolatorProfile"); ok {
			if err = d.Set("isolator_profile", vv); err != nil {
				return fmt.Errorf("Error reading isolator_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading isolator_profile: %v", err)
		}
	}

	if err = d.Set("redirect_profile", flattenFirewallProfileGroupRedirectProfile(o["redirect-profile"], d, "redirect_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["redirect-profile"], "FirewallProfileGroup-RedirectProfile"); ok {
			if err = d.Set("redirect_profile", vv); err != nil {
				return fmt.Errorf("Error reading redirect_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading redirect_profile: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileGroupApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupCifsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupCasbProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupDiameterFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupDlpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupFileFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupIpsVoipFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileGroupProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupSctpFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupSshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupSslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupTelemetryProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupVideofilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupWafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupIaProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupIsolatorProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileGroupRedirectProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallProfileGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("application_list"); ok || d.HasChange("application_list") {
		t, err := expandFirewallProfileGroupApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok || d.HasChange("av_profile") {
		t, err := expandFirewallProfileGroupAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok || d.HasChange("cifs_profile") {
		t, err := expandFirewallProfileGroupCifsProfile(d, v, "cifs_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok || d.HasChange("casb_profile") {
		t, err := expandFirewallProfileGroupCasbProfile(d, v, "casb_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok || d.HasChange("diameter_filter_profile") {
		t, err := expandFirewallProfileGroupDiameterFilterProfile(d, v, "diameter_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok || d.HasChange("dlp_profile") {
		t, err := expandFirewallProfileGroupDlpProfile(d, v, "dlp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok || d.HasChange("dnsfilter_profile") {
		t, err := expandFirewallProfileGroupDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok || d.HasChange("emailfilter_profile") {
		t, err := expandFirewallProfileGroupEmailfilterProfile(d, v, "emailfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("file_filter_profile"); ok || d.HasChange("file_filter_profile") {
		t, err := expandFirewallProfileGroupFileFilterProfile(d, v, "file_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok || d.HasChange("icap_profile") {
		t, err := expandFirewallProfileGroupIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok || d.HasChange("ips_sensor") {
		t, err := expandFirewallProfileGroupIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok || d.HasChange("ips_voip_filter") {
		t, err := expandFirewallProfileGroupIpsVoipFilter(d, v, "ips_voip_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandFirewallProfileGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok || d.HasChange("profile_protocol_options") {
		t, err := expandFirewallProfileGroupProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok || d.HasChange("sctp_filter_profile") {
		t, err := expandFirewallProfileGroupSctpFilterProfile(d, v, "sctp_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok || d.HasChange("ssh_filter_profile") {
		t, err := expandFirewallProfileGroupSshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok || d.HasChange("ssl_ssh_profile") {
		t, err := expandFirewallProfileGroupSslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("telemetry_profile"); ok || d.HasChange("telemetry_profile") {
		t, err := expandFirewallProfileGroupTelemetryProfile(d, v, "telemetry_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["telemetry-profile"] = t
		}
	}

	if v, ok := d.GetOk("videofilter_profile"); ok || d.HasChange("videofilter_profile") {
		t, err := expandFirewallProfileGroupVideofilterProfile(d, v, "videofilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok || d.HasChange("virtual_patch_profile") {
		t, err := expandFirewallProfileGroupVirtualPatchProfile(d, v, "virtual_patch_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok || d.HasChange("voip_profile") {
		t, err := expandFirewallProfileGroupVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok || d.HasChange("waf_profile") {
		t, err := expandFirewallProfileGroupWafProfile(d, v, "waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok || d.HasChange("webfilter_profile") {
		t, err := expandFirewallProfileGroupWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok || d.HasChange("dlp_sensor") {
		t, err := expandFirewallProfileGroupDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ia_profile"); ok || d.HasChange("ia_profile") {
		t, err := expandFirewallProfileGroupIaProfile(d, v, "ia_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ia-profile"] = t
		}
	}

	if v, ok := d.GetOk("isolator_profile"); ok || d.HasChange("isolator_profile") {
		t, err := expandFirewallProfileGroupIsolatorProfile(d, v, "isolator_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isolator-profile"] = t
		}
	}

	if v, ok := d.GetOk("redirect_profile"); ok || d.HasChange("redirect_profile") {
		t, err := expandFirewallProfileGroupRedirectProfile(d, v, "redirect_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-profile"] = t
		}
	}

	return &obj, nil
}
