// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure CIFS protocol options.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProfileProtocolOptionsCifs() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProfileProtocolOptionsCifsUpdate,
		Read:   resourceFirewallProfileProtocolOptionsCifsRead,
		Update: resourceFirewallProfileProtocolOptionsCifsUpdate,
		Delete: resourceFirewallProfileProtocolOptionsCifsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"file_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"comment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"file_type": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"filter": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"oversize_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"scan_bzip2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_credential_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_keytab": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keytab": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"principal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_window_maximum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_window_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uncompressed_nest_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uncompressed_oversize_limit": &schema.Schema{
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

func resourceFirewallProfileProtocolOptionsCifsUpdate(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectFirewallProfileProtocolOptionsCifs(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsCifs resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateFirewallProfileProtocolOptionsCifs(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProfileProtocolOptionsCifs resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallProfileProtocolOptionsCifs")

	return resourceFirewallProfileProtocolOptionsCifsRead(d, m)
}

func resourceFirewallProfileProtocolOptionsCifsDelete(d *schema.ResourceData, m interface{}) error {
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
	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	wsParams["adom"] = adomv

	err = c.DeleteFirewallProfileProtocolOptionsCifs(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProfileProtocolOptionsCifs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProfileProtocolOptionsCifsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg

	device_name, err := getVariable(cfg, d, "device_name")
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	profile_protocol_options := d.Get("profile_protocol_options").(string)
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
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadFirewallProfileProtocolOptionsCifs(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsCifs resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProfileProtocolOptionsCifs(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProfileProtocolOptionsCifs resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProfileProtocolOptionsCifsDomainController2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsFileFilter2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {
		result["entries"] = flattenFirewallProfileProtocolOptionsCifsFileFilterEntries2edl(i["entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenFirewallProfileProtocolOptionsCifsFileFilterLog2edl(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallProfileProtocolOptionsCifsFileFilterStatus2edl(i["status"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntries2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesAction2edl(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesComment2edl(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection2edl(i["direction"], d, pre_append)
			tmp["direction"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Direction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType2edl(i["file-type"], d, pre_append)
			tmp["file_type"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-FileType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter2edl(i["filter"], d, pre_append)
			tmp["filter"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Filter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol2edl(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifsFileFilter-Entries-Protocol")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesComment2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsFileFilterStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsOptions2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallProfileProtocolOptionsCifsOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenFirewallProfileProtocolOptionsCifsScanBzip22edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerCredentialType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytab2edl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := i["keytab"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab2edl(i["keytab"], d, pre_append)
			tmp["keytab"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifs-ServerKeytab-Keytab")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := i["principal"]; ok {
			v := flattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal2edl(i["principal"], d, pre_append)
			tmp["principal"] = fortiAPISubPartPatch(v, "FirewallProfileProtocolOptionsCifs-ServerKeytab-Principal")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabKeytab2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsServerKeytabPrincipal2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowMaximum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowMinimum2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsTcpWindowType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsUncompressedNestLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallProfileProtocolOptionsCifs(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("domain_controller", flattenFirewallProfileProtocolOptionsCifsDomainController2edl(o["domain-controller"], d, "domain_controller")); err != nil {
		if vv, ok := fortiAPIPatch(o["domain-controller"], "FirewallProfileProtocolOptionsCifs-DomainController"); ok {
			if err = d.Set("domain_controller", vv); err != nil {
				return fmt.Errorf("Error reading domain_controller: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading domain_controller: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("file_filter", flattenFirewallProfileProtocolOptionsCifsFileFilter2edl(o["file-filter"], d, "file_filter")); err != nil {
			if vv, ok := fortiAPIPatch(o["file-filter"], "FirewallProfileProtocolOptionsCifs-FileFilter"); ok {
				if err = d.Set("file_filter", vv); err != nil {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading file_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("file_filter"); ok {
			if err = d.Set("file_filter", flattenFirewallProfileProtocolOptionsCifsFileFilter2edl(o["file-filter"], d, "file_filter")); err != nil {
				if vv, ok := fortiAPIPatch(o["file-filter"], "FirewallProfileProtocolOptionsCifs-FileFilter"); ok {
					if err = d.Set("file_filter", vv); err != nil {
						return fmt.Errorf("Error reading file_filter: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("options", flattenFirewallProfileProtocolOptionsCifsOptions2edl(o["options"], d, "options")); err != nil {
		if vv, ok := fortiAPIPatch(o["options"], "FirewallProfileProtocolOptionsCifs-Options"); ok {
			if err = d.Set("options", vv); err != nil {
				return fmt.Errorf("Error reading options: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("oversize_limit", flattenFirewallProfileProtocolOptionsCifsOversizeLimit2edl(o["oversize-limit"], d, "oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["oversize-limit"], "FirewallProfileProtocolOptionsCifs-OversizeLimit"); ok {
			if err = d.Set("oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading oversize_limit: %v", err)
		}
	}

	if err = d.Set("ports", flattenFirewallProfileProtocolOptionsCifsPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "FirewallProfileProtocolOptionsCifs-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("scan_bzip2", flattenFirewallProfileProtocolOptionsCifsScanBzip22edl(o["scan-bzip2"], d, "scan_bzip2")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-bzip2"], "FirewallProfileProtocolOptionsCifs-ScanBzip2"); ok {
			if err = d.Set("scan_bzip2", vv); err != nil {
				return fmt.Errorf("Error reading scan_bzip2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_bzip2: %v", err)
		}
	}

	if err = d.Set("server_credential_type", flattenFirewallProfileProtocolOptionsCifsServerCredentialType2edl(o["server-credential-type"], d, "server_credential_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-credential-type"], "FirewallProfileProtocolOptionsCifs-ServerCredentialType"); ok {
			if err = d.Set("server_credential_type", vv); err != nil {
				return fmt.Errorf("Error reading server_credential_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_credential_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_keytab", flattenFirewallProfileProtocolOptionsCifsServerKeytab2edl(o["server-keytab"], d, "server_keytab")); err != nil {
			if vv, ok := fortiAPIPatch(o["server-keytab"], "FirewallProfileProtocolOptionsCifs-ServerKeytab"); ok {
				if err = d.Set("server_keytab", vv); err != nil {
					return fmt.Errorf("Error reading server_keytab: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading server_keytab: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_keytab"); ok {
			if err = d.Set("server_keytab", flattenFirewallProfileProtocolOptionsCifsServerKeytab2edl(o["server-keytab"], d, "server_keytab")); err != nil {
				if vv, ok := fortiAPIPatch(o["server-keytab"], "FirewallProfileProtocolOptionsCifs-ServerKeytab"); ok {
					if err = d.Set("server_keytab", vv); err != nil {
						return fmt.Errorf("Error reading server_keytab: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading server_keytab: %v", err)
				}
			}
		}
	}

	if err = d.Set("status", flattenFirewallProfileProtocolOptionsCifsStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FirewallProfileProtocolOptionsCifs-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tcp_window_maximum", flattenFirewallProfileProtocolOptionsCifsTcpWindowMaximum2edl(o["tcp-window-maximum"], d, "tcp_window_maximum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-maximum"], "FirewallProfileProtocolOptionsCifs-TcpWindowMaximum"); ok {
			if err = d.Set("tcp_window_maximum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_maximum: %v", err)
		}
	}

	if err = d.Set("tcp_window_minimum", flattenFirewallProfileProtocolOptionsCifsTcpWindowMinimum2edl(o["tcp-window-minimum"], d, "tcp_window_minimum")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-minimum"], "FirewallProfileProtocolOptionsCifs-TcpWindowMinimum"); ok {
			if err = d.Set("tcp_window_minimum", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_minimum: %v", err)
		}
	}

	if err = d.Set("tcp_window_size", flattenFirewallProfileProtocolOptionsCifsTcpWindowSize2edl(o["tcp-window-size"], d, "tcp_window_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-size"], "FirewallProfileProtocolOptionsCifs-TcpWindowSize"); ok {
			if err = d.Set("tcp_window_size", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_size: %v", err)
		}
	}

	if err = d.Set("tcp_window_type", flattenFirewallProfileProtocolOptionsCifsTcpWindowType2edl(o["tcp-window-type"], d, "tcp_window_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-window-type"], "FirewallProfileProtocolOptionsCifs-TcpWindowType"); ok {
			if err = d.Set("tcp_window_type", vv); err != nil {
				return fmt.Errorf("Error reading tcp_window_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_window_type: %v", err)
		}
	}

	if err = d.Set("uncompressed_nest_limit", flattenFirewallProfileProtocolOptionsCifsUncompressedNestLimit2edl(o["uncompressed-nest-limit"], d, "uncompressed_nest_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-nest-limit"], "FirewallProfileProtocolOptionsCifs-UncompressedNestLimit"); ok {
			if err = d.Set("uncompressed_nest_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_nest_limit: %v", err)
		}
	}

	if err = d.Set("uncompressed_oversize_limit", flattenFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit2edl(o["uncompressed-oversize-limit"], d, "uncompressed_oversize_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["uncompressed-oversize-limit"], "FirewallProfileProtocolOptionsCifs-UncompressedOversizeLimit"); ok {
			if err = d.Set("uncompressed_oversize_limit", vv); err != nil {
				return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uncompressed_oversize_limit: %v", err)
		}
	}

	return nil
}

func flattenFirewallProfileProtocolOptionsCifsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallProfileProtocolOptionsCifsDomainController2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		t, err := expandFirewallProfileProtocolOptionsCifsFileFilterEntries2edl(d, i["entries"], pre_append)
		if err != nil {
			return result, err
		} else if t != nil {
			result["entries"] = t
		}
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterLog2edl(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterStatus2edl(d, i["status"], pre_append)
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntries2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesAction2edl(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesComment2edl(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["direction"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection2edl(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["file-type"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType2edl(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["filter"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter2edl(d, i["filter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol2edl(d, i["protocol"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesComment2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesDirection2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFileType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesFilter2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterEntriesProtocol2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsFileFilterStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsOptions2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsScanBzip22edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerCredentialType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytab2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keytab"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["keytab"], _ = expandFirewallProfileProtocolOptionsCifsServerKeytabKeytab2edl(d, i["keytab"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandFirewallProfileProtocolOptionsCifsServerKeytabPassword2edl(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "principal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["principal"], _ = expandFirewallProfileProtocolOptionsCifsServerKeytabPrincipal2edl(d, i["principal"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabKeytab2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabPassword2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallProfileProtocolOptionsCifsServerKeytabPrincipal2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowMaximum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowMinimum2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsTcpWindowType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsUncompressedNestLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProfileProtocolOptionsCifs(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain_controller"); ok || d.HasChange("domain_controller") {
		t, err := expandFirewallProfileProtocolOptionsCifsDomainController2edl(d, v, "domain_controller")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-controller"] = t
		}
	}

	if v, ok := d.GetOk("file_filter"); ok || d.HasChange("file_filter") {
		t, err := expandFirewallProfileProtocolOptionsCifsFileFilter2edl(d, v, "file_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandFirewallProfileProtocolOptionsCifsOptions2edl(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("oversize_limit"); ok || d.HasChange("oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsCifsOversizeLimit2edl(d, v, "oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oversize-limit"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandFirewallProfileProtocolOptionsCifsPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("scan_bzip2"); ok || d.HasChange("scan_bzip2") {
		t, err := expandFirewallProfileProtocolOptionsCifsScanBzip22edl(d, v, "scan_bzip2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-bzip2"] = t
		}
	}

	if v, ok := d.GetOk("server_credential_type"); ok || d.HasChange("server_credential_type") {
		t, err := expandFirewallProfileProtocolOptionsCifsServerCredentialType2edl(d, v, "server_credential_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-credential-type"] = t
		}
	}

	if v, ok := d.GetOk("server_keytab"); ok || d.HasChange("server_keytab") {
		t, err := expandFirewallProfileProtocolOptionsCifsServerKeytab2edl(d, v, "server_keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-keytab"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFirewallProfileProtocolOptionsCifsStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_maximum"); ok || d.HasChange("tcp_window_maximum") {
		t, err := expandFirewallProfileProtocolOptionsCifsTcpWindowMaximum2edl(d, v, "tcp_window_maximum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-maximum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_minimum"); ok || d.HasChange("tcp_window_minimum") {
		t, err := expandFirewallProfileProtocolOptionsCifsTcpWindowMinimum2edl(d, v, "tcp_window_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-minimum"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_size"); ok || d.HasChange("tcp_window_size") {
		t, err := expandFirewallProfileProtocolOptionsCifsTcpWindowSize2edl(d, v, "tcp_window_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-size"] = t
		}
	}

	if v, ok := d.GetOk("tcp_window_type"); ok || d.HasChange("tcp_window_type") {
		t, err := expandFirewallProfileProtocolOptionsCifsTcpWindowType2edl(d, v, "tcp_window_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-window-type"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_nest_limit"); ok || d.HasChange("uncompressed_nest_limit") {
		t, err := expandFirewallProfileProtocolOptionsCifsUncompressedNestLimit2edl(d, v, "uncompressed_nest_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-nest-limit"] = t
		}
	}

	if v, ok := d.GetOk("uncompressed_oversize_limit"); ok || d.HasChange("uncompressed_oversize_limit") {
		t, err := expandFirewallProfileProtocolOptionsCifsUncompressedOversizeLimit2edl(d, v, "uncompressed_oversize_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uncompressed-oversize-limit"] = t
		}
	}

	return &obj, nil
}
