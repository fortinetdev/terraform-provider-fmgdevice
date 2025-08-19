// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure access profiles for system administrators.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAccprofileCreate,
		Read:   resourceSystemAccprofileRead,
		Update: resourceSystemAccprofileUpdate,
		Delete: resourceSystemAccprofileDelete,

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
			"admintimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"admintimeout_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_diagnose": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_exec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_show": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ftviewgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"others": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"schedule": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loggrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"report_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"threat_weight": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cfg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"packet_capture": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_cfg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secfabgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sysgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cfg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mnt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"upd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"system_diagnostics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_execute_ssh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_execute_telnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utmgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utmgrp_permission": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"antivirus": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"application_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"casb": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_leak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_loss_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dlp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dnsfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emailfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"endpoint_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"file_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"icap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ips": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mmsgtp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"videofilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_patch": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"voip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"waf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"webfilter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"vpngrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanoptgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAccprofileCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemAccprofile(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAccprofile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemAccprofile(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAccprofile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAccprofile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAccprofile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAccprofile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAccprofile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource from API: %v", err)
	}
	return nil
}

func flattenSystemAccprofileAdmintimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileAdmintimeoutOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileAuthgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileCliConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileCliDiagnose(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileCliExec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileCliGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileCliShow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFtviewgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "address"
	if _, ok := i["address"]; ok {
		result["address"] = flattenSystemAccprofileFwgrpPermissionAddress(i["address"], d, pre_append)
	}

	pre_append = pre + ".0." + "others"
	if _, ok := i["others"]; ok {
		result["others"] = flattenSystemAccprofileFwgrpPermissionOthers(i["others"], d, pre_append)
	}

	pre_append = pre + ".0." + "policy"
	if _, ok := i["policy"]; ok {
		result["policy"] = flattenSystemAccprofileFwgrpPermissionPolicy(i["policy"], d, pre_append)
	}

	pre_append = pre + ".0." + "schedule"
	if _, ok := i["schedule"]; ok {
		result["schedule"] = flattenSystemAccprofileFwgrpPermissionSchedule(i["schedule"], d, pre_append)
	}

	pre_append = pre + ".0." + "service"
	if _, ok := i["service"]; ok {
		result["service"] = flattenSystemAccprofileFwgrpPermissionService(i["service"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileFwgrpPermissionAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionOthers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileFwgrpPermissionService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config"
	if _, ok := i["config"]; ok {
		result["config"] = flattenSystemAccprofileLoggrpPermissionConfig(i["config"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_access"
	if _, ok := i["data-access"]; ok {
		result["data_access"] = flattenSystemAccprofileLoggrpPermissionDataAccess(i["data-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "report_access"
	if _, ok := i["report-access"]; ok {
		result["report_access"] = flattenSystemAccprofileLoggrpPermissionReportAccess(i["report-access"], d, pre_append)
	}

	pre_append = pre + ".0." + "threat_weight"
	if _, ok := i["threat-weight"]; ok {
		result["threat_weight"] = flattenSystemAccprofileLoggrpPermissionThreatWeight(i["threat-weight"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileLoggrpPermissionConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermissionDataAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermissionReportAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileLoggrpPermissionThreatWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {
		result["cfg"] = flattenSystemAccprofileNetgrpPermissionCfg(i["cfg"], d, pre_append)
	}

	pre_append = pre + ".0." + "packet_capture"
	if _, ok := i["packet-capture"]; ok {
		result["packet_capture"] = flattenSystemAccprofileNetgrpPermissionPacketCapture(i["packet-capture"], d, pre_append)
	}

	pre_append = pre + ".0." + "route_cfg"
	if _, ok := i["route-cfg"]; ok {
		result["route_cfg"] = flattenSystemAccprofileNetgrpPermissionRouteCfg(i["route-cfg"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileNetgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrpPermissionPacketCapture(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileNetgrpPermissionRouteCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSecfabgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "admin"
	if _, ok := i["admin"]; ok {
		result["admin"] = flattenSystemAccprofileSysgrpPermissionAdmin(i["admin"], d, pre_append)
	}

	pre_append = pre + ".0." + "cfg"
	if _, ok := i["cfg"]; ok {
		result["cfg"] = flattenSystemAccprofileSysgrpPermissionCfg(i["cfg"], d, pre_append)
	}

	pre_append = pre + ".0." + "mnt"
	if _, ok := i["mnt"]; ok {
		result["mnt"] = flattenSystemAccprofileSysgrpPermissionMnt(i["mnt"], d, pre_append)
	}

	pre_append = pre + ".0." + "upd"
	if _, ok := i["upd"]; ok {
		result["upd"] = flattenSystemAccprofileSysgrpPermissionUpd(i["upd"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileSysgrpPermissionAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermissionCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermissionMnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSysgrpPermissionUpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSystemDiagnostics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSystemExecuteSsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileSystemExecuteTelnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermission(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "antivirus"
	if _, ok := i["antivirus"]; ok {
		result["antivirus"] = flattenSystemAccprofileUtmgrpPermissionAntivirus(i["antivirus"], d, pre_append)
	}

	pre_append = pre + ".0." + "application_control"
	if _, ok := i["application-control"]; ok {
		result["application_control"] = flattenSystemAccprofileUtmgrpPermissionApplicationControl(i["application-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "casb"
	if _, ok := i["casb"]; ok {
		result["casb"] = flattenSystemAccprofileUtmgrpPermissionCasb(i["casb"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_leak_prevention"
	if _, ok := i["data-leak-prevention"]; ok {
		result["data_leak_prevention"] = flattenSystemAccprofileUtmgrpPermissionDataLeakPrevention(i["data-leak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "data_loss_prevention"
	if _, ok := i["data-loss-prevention"]; ok {
		result["data_loss_prevention"] = flattenSystemAccprofileUtmgrpPermissionDataLossPrevention(i["data-loss-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "dlp"
	if _, ok := i["dlp"]; ok {
		result["dlp"] = flattenSystemAccprofileUtmgrpPermissionDlp(i["dlp"], d, pre_append)
	}

	pre_append = pre + ".0." + "dnsfilter"
	if _, ok := i["dnsfilter"]; ok {
		result["dnsfilter"] = flattenSystemAccprofileUtmgrpPermissionDnsfilter(i["dnsfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "emailfilter"
	if _, ok := i["emailfilter"]; ok {
		result["emailfilter"] = flattenSystemAccprofileUtmgrpPermissionEmailfilter(i["emailfilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "endpoint_control"
	if _, ok := i["endpoint-control"]; ok {
		result["endpoint_control"] = flattenSystemAccprofileUtmgrpPermissionEndpointControl(i["endpoint-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "file_filter"
	if _, ok := i["file-filter"]; ok {
		result["file_filter"] = flattenSystemAccprofileUtmgrpPermissionFileFilter(i["file-filter"], d, pre_append)
	}

	pre_append = pre + ".0." + "icap"
	if _, ok := i["icap"]; ok {
		result["icap"] = flattenSystemAccprofileUtmgrpPermissionIcap(i["icap"], d, pre_append)
	}

	pre_append = pre + ".0." + "ips"
	if _, ok := i["ips"]; ok {
		result["ips"] = flattenSystemAccprofileUtmgrpPermissionIps(i["ips"], d, pre_append)
	}

	pre_append = pre + ".0." + "mmsgtp"
	if _, ok := i["mmsgtp"]; ok {
		result["mmsgtp"] = flattenSystemAccprofileUtmgrpPermissionMmsgtp(i["mmsgtp"], d, pre_append)
	}

	pre_append = pre + ".0." + "videofilter"
	if _, ok := i["videofilter"]; ok {
		result["videofilter"] = flattenSystemAccprofileUtmgrpPermissionVideofilter(i["videofilter"], d, pre_append)
	}

	pre_append = pre + ".0." + "virtual_patch"
	if _, ok := i["virtual-patch"]; ok {
		result["virtual_patch"] = flattenSystemAccprofileUtmgrpPermissionVirtualPatch(i["virtual-patch"], d, pre_append)
	}

	pre_append = pre + ".0." + "voip"
	if _, ok := i["voip"]; ok {
		result["voip"] = flattenSystemAccprofileUtmgrpPermissionVoip(i["voip"], d, pre_append)
	}

	pre_append = pre + ".0." + "waf"
	if _, ok := i["waf"]; ok {
		result["waf"] = flattenSystemAccprofileUtmgrpPermissionWaf(i["waf"], d, pre_append)
	}

	pre_append = pre + ".0." + "webfilter"
	if _, ok := i["webfilter"]; ok {
		result["webfilter"] = flattenSystemAccprofileUtmgrpPermissionWebfilter(i["webfilter"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemAccprofileUtmgrpPermissionAntivirus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionApplicationControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionCasb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionDataLeakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionDataLossPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionDlp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionDnsfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionEmailfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionEndpointControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionFileFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionIcap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionIps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionMmsgtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionVideofilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionVirtualPatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionWaf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileUtmgrpPermissionWebfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileVpngrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileWanoptgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAccprofileWifi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("admintimeout", flattenSystemAccprofileAdmintimeout(o["admintimeout"], d, "admintimeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["admintimeout"], "SystemAccprofile-Admintimeout"); ok {
			if err = d.Set("admintimeout", vv); err != nil {
				return fmt.Errorf("Error reading admintimeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admintimeout: %v", err)
		}
	}

	if err = d.Set("admintimeout_override", flattenSystemAccprofileAdmintimeoutOverride(o["admintimeout-override"], d, "admintimeout_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["admintimeout-override"], "SystemAccprofile-AdmintimeoutOverride"); ok {
			if err = d.Set("admintimeout_override", vv); err != nil {
				return fmt.Errorf("Error reading admintimeout_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading admintimeout_override: %v", err)
		}
	}

	if err = d.Set("authgrp", flattenSystemAccprofileAuthgrp(o["authgrp"], d, "authgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["authgrp"], "SystemAccprofile-Authgrp"); ok {
			if err = d.Set("authgrp", vv); err != nil {
				return fmt.Errorf("Error reading authgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading authgrp: %v", err)
		}
	}

	if err = d.Set("cli_config", flattenSystemAccprofileCliConfig(o["cli-config"], d, "cli_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-config"], "SystemAccprofile-CliConfig"); ok {
			if err = d.Set("cli_config", vv); err != nil {
				return fmt.Errorf("Error reading cli_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_config: %v", err)
		}
	}

	if err = d.Set("cli_diagnose", flattenSystemAccprofileCliDiagnose(o["cli-diagnose"], d, "cli_diagnose")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-diagnose"], "SystemAccprofile-CliDiagnose"); ok {
			if err = d.Set("cli_diagnose", vv); err != nil {
				return fmt.Errorf("Error reading cli_diagnose: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_diagnose: %v", err)
		}
	}

	if err = d.Set("cli_exec", flattenSystemAccprofileCliExec(o["cli-exec"], d, "cli_exec")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-exec"], "SystemAccprofile-CliExec"); ok {
			if err = d.Set("cli_exec", vv); err != nil {
				return fmt.Errorf("Error reading cli_exec: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_exec: %v", err)
		}
	}

	if err = d.Set("cli_get", flattenSystemAccprofileCliGet(o["cli-get"], d, "cli_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-get"], "SystemAccprofile-CliGet"); ok {
			if err = d.Set("cli_get", vv); err != nil {
				return fmt.Errorf("Error reading cli_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_get: %v", err)
		}
	}

	if err = d.Set("cli_show", flattenSystemAccprofileCliShow(o["cli-show"], d, "cli_show")); err != nil {
		if vv, ok := fortiAPIPatch(o["cli-show"], "SystemAccprofile-CliShow"); ok {
			if err = d.Set("cli_show", vv); err != nil {
				return fmt.Errorf("Error reading cli_show: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cli_show: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemAccprofileComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "SystemAccprofile-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("ftviewgrp", flattenSystemAccprofileFtviewgrp(o["ftviewgrp"], d, "ftviewgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftviewgrp"], "SystemAccprofile-Ftviewgrp"); ok {
			if err = d.Set("ftviewgrp", vv); err != nil {
				return fmt.Errorf("Error reading ftviewgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftviewgrp: %v", err)
		}
	}

	if err = d.Set("fwgrp", flattenSystemAccprofileFwgrp(o["fwgrp"], d, "fwgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["fwgrp"], "SystemAccprofile-Fwgrp"); ok {
			if err = d.Set("fwgrp", vv); err != nil {
				return fmt.Errorf("Error reading fwgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fwgrp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fwgrp_permission", flattenSystemAccprofileFwgrpPermission(o["fwgrp-permission"], d, "fwgrp_permission")); err != nil {
			if vv, ok := fortiAPIPatch(o["fwgrp-permission"], "SystemAccprofile-FwgrpPermission"); ok {
				if err = d.Set("fwgrp_permission", vv); err != nil {
					return fmt.Errorf("Error reading fwgrp_permission: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fwgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fwgrp_permission"); ok {
			if err = d.Set("fwgrp_permission", flattenSystemAccprofileFwgrpPermission(o["fwgrp-permission"], d, "fwgrp_permission")); err != nil {
				if vv, ok := fortiAPIPatch(o["fwgrp-permission"], "SystemAccprofile-FwgrpPermission"); ok {
					if err = d.Set("fwgrp_permission", vv); err != nil {
						return fmt.Errorf("Error reading fwgrp_permission: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fwgrp_permission: %v", err)
				}
			}
		}
	}

	if err = d.Set("loggrp", flattenSystemAccprofileLoggrp(o["loggrp"], d, "loggrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["loggrp"], "SystemAccprofile-Loggrp"); ok {
			if err = d.Set("loggrp", vv); err != nil {
				return fmt.Errorf("Error reading loggrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading loggrp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("loggrp_permission", flattenSystemAccprofileLoggrpPermission(o["loggrp-permission"], d, "loggrp_permission")); err != nil {
			if vv, ok := fortiAPIPatch(o["loggrp-permission"], "SystemAccprofile-LoggrpPermission"); ok {
				if err = d.Set("loggrp_permission", vv); err != nil {
					return fmt.Errorf("Error reading loggrp_permission: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading loggrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("loggrp_permission"); ok {
			if err = d.Set("loggrp_permission", flattenSystemAccprofileLoggrpPermission(o["loggrp-permission"], d, "loggrp_permission")); err != nil {
				if vv, ok := fortiAPIPatch(o["loggrp-permission"], "SystemAccprofile-LoggrpPermission"); ok {
					if err = d.Set("loggrp_permission", vv); err != nil {
						return fmt.Errorf("Error reading loggrp_permission: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading loggrp_permission: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemAccprofileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAccprofile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("netgrp", flattenSystemAccprofileNetgrp(o["netgrp"], d, "netgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["netgrp"], "SystemAccprofile-Netgrp"); ok {
			if err = d.Set("netgrp", vv); err != nil {
				return fmt.Errorf("Error reading netgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading netgrp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("netgrp_permission", flattenSystemAccprofileNetgrpPermission(o["netgrp-permission"], d, "netgrp_permission")); err != nil {
			if vv, ok := fortiAPIPatch(o["netgrp-permission"], "SystemAccprofile-NetgrpPermission"); ok {
				if err = d.Set("netgrp_permission", vv); err != nil {
					return fmt.Errorf("Error reading netgrp_permission: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading netgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("netgrp_permission"); ok {
			if err = d.Set("netgrp_permission", flattenSystemAccprofileNetgrpPermission(o["netgrp-permission"], d, "netgrp_permission")); err != nil {
				if vv, ok := fortiAPIPatch(o["netgrp-permission"], "SystemAccprofile-NetgrpPermission"); ok {
					if err = d.Set("netgrp_permission", vv); err != nil {
						return fmt.Errorf("Error reading netgrp_permission: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading netgrp_permission: %v", err)
				}
			}
		}
	}

	if err = d.Set("scope", flattenSystemAccprofileScope(o["scope"], d, "scope")); err != nil {
		if vv, ok := fortiAPIPatch(o["scope"], "SystemAccprofile-Scope"); ok {
			if err = d.Set("scope", vv); err != nil {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("secfabgrp", flattenSystemAccprofileSecfabgrp(o["secfabgrp"], d, "secfabgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["secfabgrp"], "SystemAccprofile-Secfabgrp"); ok {
			if err = d.Set("secfabgrp", vv); err != nil {
				return fmt.Errorf("Error reading secfabgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secfabgrp: %v", err)
		}
	}

	if err = d.Set("sysgrp", flattenSystemAccprofileSysgrp(o["sysgrp"], d, "sysgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["sysgrp"], "SystemAccprofile-Sysgrp"); ok {
			if err = d.Set("sysgrp", vv); err != nil {
				return fmt.Errorf("Error reading sysgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sysgrp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sysgrp_permission", flattenSystemAccprofileSysgrpPermission(o["sysgrp-permission"], d, "sysgrp_permission")); err != nil {
			if vv, ok := fortiAPIPatch(o["sysgrp-permission"], "SystemAccprofile-SysgrpPermission"); ok {
				if err = d.Set("sysgrp_permission", vv); err != nil {
					return fmt.Errorf("Error reading sysgrp_permission: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sysgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sysgrp_permission"); ok {
			if err = d.Set("sysgrp_permission", flattenSystemAccprofileSysgrpPermission(o["sysgrp-permission"], d, "sysgrp_permission")); err != nil {
				if vv, ok := fortiAPIPatch(o["sysgrp-permission"], "SystemAccprofile-SysgrpPermission"); ok {
					if err = d.Set("sysgrp_permission", vv); err != nil {
						return fmt.Errorf("Error reading sysgrp_permission: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sysgrp_permission: %v", err)
				}
			}
		}
	}

	if err = d.Set("system_diagnostics", flattenSystemAccprofileSystemDiagnostics(o["system-diagnostics"], d, "system_diagnostics")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-diagnostics"], "SystemAccprofile-SystemDiagnostics"); ok {
			if err = d.Set("system_diagnostics", vv); err != nil {
				return fmt.Errorf("Error reading system_diagnostics: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_diagnostics: %v", err)
		}
	}

	if err = d.Set("system_execute_ssh", flattenSystemAccprofileSystemExecuteSsh(o["system-execute-ssh"], d, "system_execute_ssh")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-execute-ssh"], "SystemAccprofile-SystemExecuteSsh"); ok {
			if err = d.Set("system_execute_ssh", vv); err != nil {
				return fmt.Errorf("Error reading system_execute_ssh: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_execute_ssh: %v", err)
		}
	}

	if err = d.Set("system_execute_telnet", flattenSystemAccprofileSystemExecuteTelnet(o["system-execute-telnet"], d, "system_execute_telnet")); err != nil {
		if vv, ok := fortiAPIPatch(o["system-execute-telnet"], "SystemAccprofile-SystemExecuteTelnet"); ok {
			if err = d.Set("system_execute_telnet", vv); err != nil {
				return fmt.Errorf("Error reading system_execute_telnet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system_execute_telnet: %v", err)
		}
	}

	if err = d.Set("utmgrp", flattenSystemAccprofileUtmgrp(o["utmgrp"], d, "utmgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["utmgrp"], "SystemAccprofile-Utmgrp"); ok {
			if err = d.Set("utmgrp", vv); err != nil {
				return fmt.Errorf("Error reading utmgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading utmgrp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("utmgrp_permission", flattenSystemAccprofileUtmgrpPermission(o["utmgrp-permission"], d, "utmgrp_permission")); err != nil {
			if vv, ok := fortiAPIPatch(o["utmgrp-permission"], "SystemAccprofile-UtmgrpPermission"); ok {
				if err = d.Set("utmgrp_permission", vv); err != nil {
					return fmt.Errorf("Error reading utmgrp_permission: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading utmgrp_permission: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("utmgrp_permission"); ok {
			if err = d.Set("utmgrp_permission", flattenSystemAccprofileUtmgrpPermission(o["utmgrp-permission"], d, "utmgrp_permission")); err != nil {
				if vv, ok := fortiAPIPatch(o["utmgrp-permission"], "SystemAccprofile-UtmgrpPermission"); ok {
					if err = d.Set("utmgrp_permission", vv); err != nil {
						return fmt.Errorf("Error reading utmgrp_permission: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading utmgrp_permission: %v", err)
				}
			}
		}
	}

	if err = d.Set("vpngrp", flattenSystemAccprofileVpngrp(o["vpngrp"], d, "vpngrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpngrp"], "SystemAccprofile-Vpngrp"); ok {
			if err = d.Set("vpngrp", vv); err != nil {
				return fmt.Errorf("Error reading vpngrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpngrp: %v", err)
		}
	}

	if err = d.Set("wanoptgrp", flattenSystemAccprofileWanoptgrp(o["wanoptgrp"], d, "wanoptgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["wanoptgrp"], "SystemAccprofile-Wanoptgrp"); ok {
			if err = d.Set("wanoptgrp", vv); err != nil {
				return fmt.Errorf("Error reading wanoptgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wanoptgrp: %v", err)
		}
	}

	if err = d.Set("wifi", flattenSystemAccprofileWifi(o["wifi"], d, "wifi")); err != nil {
		if vv, ok := fortiAPIPatch(o["wifi"], "SystemAccprofile-Wifi"); ok {
			if err = d.Set("wifi", vv); err != nil {
				return fmt.Errorf("Error reading wifi: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wifi: %v", err)
		}
	}

	return nil
}

func flattenSystemAccprofileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAccprofileAdmintimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAdmintimeoutOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileAuthgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileCliConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileCliDiagnose(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileCliExec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileCliGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileCliShow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFtviewgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "address"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["address"], _ = expandSystemAccprofileFwgrpPermissionAddress(d, i["address"], pre_append)
	}
	pre_append = pre + ".0." + "others"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["others"], _ = expandSystemAccprofileFwgrpPermissionOthers(d, i["others"], pre_append)
	}
	pre_append = pre + ".0." + "policy"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["policy"], _ = expandSystemAccprofileFwgrpPermissionPolicy(d, i["policy"], pre_append)
	}
	pre_append = pre + ".0." + "schedule"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["schedule"], _ = expandSystemAccprofileFwgrpPermissionSchedule(d, i["schedule"], pre_append)
	}
	pre_append = pre + ".0." + "service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["service"], _ = expandSystemAccprofileFwgrpPermissionService(d, i["service"], pre_append)
	}

	return result, nil
}

func expandSystemAccprofileFwgrpPermissionAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionOthers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileFwgrpPermissionService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "config"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["config"], _ = expandSystemAccprofileLoggrpPermissionConfig(d, i["config"], pre_append)
	}
	pre_append = pre + ".0." + "data_access"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-access"], _ = expandSystemAccprofileLoggrpPermissionDataAccess(d, i["data_access"], pre_append)
	}
	pre_append = pre + ".0." + "report_access"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["report-access"], _ = expandSystemAccprofileLoggrpPermissionReportAccess(d, i["report_access"], pre_append)
	}
	pre_append = pre + ".0." + "threat_weight"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["threat-weight"], _ = expandSystemAccprofileLoggrpPermissionThreatWeight(d, i["threat_weight"], pre_append)
	}

	return result, nil
}

func expandSystemAccprofileLoggrpPermissionConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermissionDataAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermissionReportAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileLoggrpPermissionThreatWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrpPermission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "cfg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cfg"], _ = expandSystemAccprofileNetgrpPermissionCfg(d, i["cfg"], pre_append)
	}
	pre_append = pre + ".0." + "packet_capture"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["packet-capture"], _ = expandSystemAccprofileNetgrpPermissionPacketCapture(d, i["packet_capture"], pre_append)
	}
	pre_append = pre + ".0." + "route_cfg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["route-cfg"], _ = expandSystemAccprofileNetgrpPermissionRouteCfg(d, i["route_cfg"], pre_append)
	}

	return result, nil
}

func expandSystemAccprofileNetgrpPermissionCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrpPermissionPacketCapture(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileNetgrpPermissionRouteCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSecfabgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "admin"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["admin"], _ = expandSystemAccprofileSysgrpPermissionAdmin(d, i["admin"], pre_append)
	}
	pre_append = pre + ".0." + "cfg"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cfg"], _ = expandSystemAccprofileSysgrpPermissionCfg(d, i["cfg"], pre_append)
	}
	pre_append = pre + ".0." + "mnt"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mnt"], _ = expandSystemAccprofileSysgrpPermissionMnt(d, i["mnt"], pre_append)
	}
	pre_append = pre + ".0." + "upd"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["upd"], _ = expandSystemAccprofileSysgrpPermissionUpd(d, i["upd"], pre_append)
	}

	return result, nil
}

func expandSystemAccprofileSysgrpPermissionAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermissionCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermissionMnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSysgrpPermissionUpd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSystemDiagnostics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSystemExecuteSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileSystemExecuteTelnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "antivirus"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["antivirus"], _ = expandSystemAccprofileUtmgrpPermissionAntivirus(d, i["antivirus"], pre_append)
	}
	pre_append = pre + ".0." + "application_control"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["application-control"], _ = expandSystemAccprofileUtmgrpPermissionApplicationControl(d, i["application_control"], pre_append)
	}
	pre_append = pre + ".0." + "casb"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["casb"], _ = expandSystemAccprofileUtmgrpPermissionCasb(d, i["casb"], pre_append)
	}
	pre_append = pre + ".0." + "data_leak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-leak-prevention"], _ = expandSystemAccprofileUtmgrpPermissionDataLeakPrevention(d, i["data_leak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "data_loss_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["data-loss-prevention"], _ = expandSystemAccprofileUtmgrpPermissionDataLossPrevention(d, i["data_loss_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "dlp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dlp"], _ = expandSystemAccprofileUtmgrpPermissionDlp(d, i["dlp"], pre_append)
	}
	pre_append = pre + ".0." + "dnsfilter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dnsfilter"], _ = expandSystemAccprofileUtmgrpPermissionDnsfilter(d, i["dnsfilter"], pre_append)
	}
	pre_append = pre + ".0." + "emailfilter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emailfilter"], _ = expandSystemAccprofileUtmgrpPermissionEmailfilter(d, i["emailfilter"], pre_append)
	}
	pre_append = pre + ".0." + "endpoint_control"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["endpoint-control"], _ = expandSystemAccprofileUtmgrpPermissionEndpointControl(d, i["endpoint_control"], pre_append)
	}
	pre_append = pre + ".0." + "file_filter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["file-filter"], _ = expandSystemAccprofileUtmgrpPermissionFileFilter(d, i["file_filter"], pre_append)
	}
	pre_append = pre + ".0." + "icap"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["icap"], _ = expandSystemAccprofileUtmgrpPermissionIcap(d, i["icap"], pre_append)
	}
	pre_append = pre + ".0." + "ips"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ips"], _ = expandSystemAccprofileUtmgrpPermissionIps(d, i["ips"], pre_append)
	}
	pre_append = pre + ".0." + "mmsgtp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mmsgtp"], _ = expandSystemAccprofileUtmgrpPermissionMmsgtp(d, i["mmsgtp"], pre_append)
	}
	pre_append = pre + ".0." + "videofilter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["videofilter"], _ = expandSystemAccprofileUtmgrpPermissionVideofilter(d, i["videofilter"], pre_append)
	}
	pre_append = pre + ".0." + "virtual_patch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["virtual-patch"], _ = expandSystemAccprofileUtmgrpPermissionVirtualPatch(d, i["virtual_patch"], pre_append)
	}
	pre_append = pre + ".0." + "voip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["voip"], _ = expandSystemAccprofileUtmgrpPermissionVoip(d, i["voip"], pre_append)
	}
	pre_append = pre + ".0." + "waf"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["waf"], _ = expandSystemAccprofileUtmgrpPermissionWaf(d, i["waf"], pre_append)
	}
	pre_append = pre + ".0." + "webfilter"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["webfilter"], _ = expandSystemAccprofileUtmgrpPermissionWebfilter(d, i["webfilter"], pre_append)
	}

	return result, nil
}

func expandSystemAccprofileUtmgrpPermissionAntivirus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionApplicationControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionCasb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionDataLeakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionDataLossPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionDlp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionDnsfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionEmailfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionEndpointControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionFileFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionIcap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionIps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionMmsgtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionVideofilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionVirtualPatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionWaf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileUtmgrpPermissionWebfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileVpngrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileWanoptgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAccprofileWifi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAccprofile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admintimeout"); ok || d.HasChange("admintimeout") {
		t, err := expandSystemAccprofileAdmintimeout(d, v, "admintimeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout"] = t
		}
	}

	if v, ok := d.GetOk("admintimeout_override"); ok || d.HasChange("admintimeout_override") {
		t, err := expandSystemAccprofileAdmintimeoutOverride(d, v, "admintimeout_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admintimeout-override"] = t
		}
	}

	if v, ok := d.GetOk("authgrp"); ok || d.HasChange("authgrp") {
		t, err := expandSystemAccprofileAuthgrp(d, v, "authgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authgrp"] = t
		}
	}

	if v, ok := d.GetOk("cli_config"); ok || d.HasChange("cli_config") {
		t, err := expandSystemAccprofileCliConfig(d, v, "cli_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-config"] = t
		}
	}

	if v, ok := d.GetOk("cli_diagnose"); ok || d.HasChange("cli_diagnose") {
		t, err := expandSystemAccprofileCliDiagnose(d, v, "cli_diagnose")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-diagnose"] = t
		}
	}

	if v, ok := d.GetOk("cli_exec"); ok || d.HasChange("cli_exec") {
		t, err := expandSystemAccprofileCliExec(d, v, "cli_exec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-exec"] = t
		}
	}

	if v, ok := d.GetOk("cli_get"); ok || d.HasChange("cli_get") {
		t, err := expandSystemAccprofileCliGet(d, v, "cli_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-get"] = t
		}
	}

	if v, ok := d.GetOk("cli_show"); ok || d.HasChange("cli_show") {
		t, err := expandSystemAccprofileCliShow(d, v, "cli_show")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-show"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandSystemAccprofileComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("ftviewgrp"); ok || d.HasChange("ftviewgrp") {
		t, err := expandSystemAccprofileFtviewgrp(d, v, "ftviewgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftviewgrp"] = t
		}
	}

	if v, ok := d.GetOk("fwgrp"); ok || d.HasChange("fwgrp") {
		t, err := expandSystemAccprofileFwgrp(d, v, "fwgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwgrp"] = t
		}
	}

	if v, ok := d.GetOk("fwgrp_permission"); ok || d.HasChange("fwgrp_permission") {
		t, err := expandSystemAccprofileFwgrpPermission(d, v, "fwgrp_permission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("loggrp"); ok || d.HasChange("loggrp") {
		t, err := expandSystemAccprofileLoggrp(d, v, "loggrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loggrp"] = t
		}
	}

	if v, ok := d.GetOk("loggrp_permission"); ok || d.HasChange("loggrp_permission") {
		t, err := expandSystemAccprofileLoggrpPermission(d, v, "loggrp_permission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loggrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAccprofileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("netgrp"); ok || d.HasChange("netgrp") {
		t, err := expandSystemAccprofileNetgrp(d, v, "netgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netgrp"] = t
		}
	}

	if v, ok := d.GetOk("netgrp_permission"); ok || d.HasChange("netgrp_permission") {
		t, err := expandSystemAccprofileNetgrpPermission(d, v, "netgrp_permission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandSystemAccprofileScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("secfabgrp"); ok || d.HasChange("secfabgrp") {
		t, err := expandSystemAccprofileSecfabgrp(d, v, "secfabgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secfabgrp"] = t
		}
	}

	if v, ok := d.GetOk("sysgrp"); ok || d.HasChange("sysgrp") {
		t, err := expandSystemAccprofileSysgrp(d, v, "sysgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sysgrp"] = t
		}
	}

	if v, ok := d.GetOk("sysgrp_permission"); ok || d.HasChange("sysgrp_permission") {
		t, err := expandSystemAccprofileSysgrpPermission(d, v, "sysgrp_permission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sysgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("system_diagnostics"); ok || d.HasChange("system_diagnostics") {
		t, err := expandSystemAccprofileSystemDiagnostics(d, v, "system_diagnostics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-diagnostics"] = t
		}
	}

	if v, ok := d.GetOk("system_execute_ssh"); ok || d.HasChange("system_execute_ssh") {
		t, err := expandSystemAccprofileSystemExecuteSsh(d, v, "system_execute_ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-execute-ssh"] = t
		}
	}

	if v, ok := d.GetOk("system_execute_telnet"); ok || d.HasChange("system_execute_telnet") {
		t, err := expandSystemAccprofileSystemExecuteTelnet(d, v, "system_execute_telnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-execute-telnet"] = t
		}
	}

	if v, ok := d.GetOk("utmgrp"); ok || d.HasChange("utmgrp") {
		t, err := expandSystemAccprofileUtmgrp(d, v, "utmgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utmgrp"] = t
		}
	}

	if v, ok := d.GetOk("utmgrp_permission"); ok || d.HasChange("utmgrp_permission") {
		t, err := expandSystemAccprofileUtmgrpPermission(d, v, "utmgrp_permission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utmgrp-permission"] = t
		}
	}

	if v, ok := d.GetOk("vpngrp"); ok || d.HasChange("vpngrp") {
		t, err := expandSystemAccprofileVpngrp(d, v, "vpngrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpngrp"] = t
		}
	}

	if v, ok := d.GetOk("wanoptgrp"); ok || d.HasChange("wanoptgrp") {
		t, err := expandSystemAccprofileWanoptgrp(d, v, "wanoptgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanoptgrp"] = t
		}
	}

	if v, ok := d.GetOk("wifi"); ok || d.HasChange("wifi") {
		t, err := expandSystemAccprofileWifi(d, v, "wifi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi"] = t
		}
	}

	return &obj, nil
}
