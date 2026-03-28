// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure AntiVirus profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusProfileCreate,
		Read:   resourceAntivirusProfileRead,
		Update: resourceAntivirusProfileUpdate,
		Delete: resourceAntivirusProfileDelete,

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
			"analytics_bl_filetype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"analytics_accept_filetype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"analytics_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"analytics_ignore_filetype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"av_virus_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_disarm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"analytics_suspicious": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cover_page": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"detect_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"error_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"office_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"office_dde": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"office_embed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"office_hylink": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"office_linked": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"office_macro": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"original_file_destination": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_act_form": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_act_gotor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_act_java": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_act_launch": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_act_movie": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_act_sound": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_embedfile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_hyperlink": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pdf_javacode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ems_threat_feed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_blocklist": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"external_blocklist_archive_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_blocklist_enable_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiai_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortiai_timeout_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortindr_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortindr_timeout_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortisandbox_error_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortisandbox_max_upload": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fortisandbox_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortisandbox_timeout_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_content_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"imap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mobile_malware_db": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_quar": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"infected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"outbreak_prevention": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ftgd_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"outbreak_prevention_archive_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"smtp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content_disarm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"executables": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"archive_block": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"archive_log": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"av_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"emulator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"external_blocklist": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortiai": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortindr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fortisandbox": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"malware_stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"outbreak_prevention": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"analytics_max_upload": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"analytics_wl_filetype": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ftgd_analytics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"av_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAntivirusProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAntivirusProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating AntivirusProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadAntivirusProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateAntivirusProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating AntivirusProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateAntivirusProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating AntivirusProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceAntivirusProfileRead(d, m)
}

func resourceAntivirusProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAntivirusProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateAntivirusProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceAntivirusProfileRead(d, m)
}

func resourceAntivirusProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteAntivirusProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAntivirusProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading AntivirusProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusProfile resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusProfileAnalyticsBlFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileAnalyticsAcceptFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileAnalyticsDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileAnalyticsIgnoreFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileAvVirusLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileCifsArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileCifsArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileCifsAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileCifsEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileCifsOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileCifsExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileCifsFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileCifsFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileCifsFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileCifsMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileCifsOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileCifsQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileCifsArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileCifsArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileCifsAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileCifsExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileCifsQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarm(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "analytics_suspicious"
	if _, ok := i["analytics-suspicious"]; ok {
		result["analytics_suspicious"] = flattenAntivirusProfileContentDisarmAnalyticsSuspicious(i["analytics-suspicious"], d, pre_append)
	}

	pre_append = pre + ".0." + "cover_page"
	if _, ok := i["cover-page"]; ok {
		result["cover_page"] = flattenAntivirusProfileContentDisarmCoverPage(i["cover-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "detect_only"
	if _, ok := i["detect-only"]; ok {
		result["detect_only"] = flattenAntivirusProfileContentDisarmDetectOnly(i["detect-only"], d, pre_append)
	}

	pre_append = pre + ".0." + "error_action"
	if _, ok := i["error-action"]; ok {
		result["error_action"] = flattenAntivirusProfileContentDisarmErrorAction(i["error-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_action"
	if _, ok := i["office-action"]; ok {
		result["office_action"] = flattenAntivirusProfileContentDisarmOfficeAction(i["office-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_dde"
	if _, ok := i["office-dde"]; ok {
		result["office_dde"] = flattenAntivirusProfileContentDisarmOfficeDde(i["office-dde"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_embed"
	if _, ok := i["office-embed"]; ok {
		result["office_embed"] = flattenAntivirusProfileContentDisarmOfficeEmbed(i["office-embed"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_hylink"
	if _, ok := i["office-hylink"]; ok {
		result["office_hylink"] = flattenAntivirusProfileContentDisarmOfficeHylink(i["office-hylink"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_linked"
	if _, ok := i["office-linked"]; ok {
		result["office_linked"] = flattenAntivirusProfileContentDisarmOfficeLinked(i["office-linked"], d, pre_append)
	}

	pre_append = pre + ".0." + "office_macro"
	if _, ok := i["office-macro"]; ok {
		result["office_macro"] = flattenAntivirusProfileContentDisarmOfficeMacro(i["office-macro"], d, pre_append)
	}

	pre_append = pre + ".0." + "original_file_destination"
	if _, ok := i["original-file-destination"]; ok {
		result["original_file_destination"] = flattenAntivirusProfileContentDisarmOriginalFileDestination(i["original-file-destination"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_form"
	if _, ok := i["pdf-act-form"]; ok {
		result["pdf_act_form"] = flattenAntivirusProfileContentDisarmPdfActForm(i["pdf-act-form"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_gotor"
	if _, ok := i["pdf-act-gotor"]; ok {
		result["pdf_act_gotor"] = flattenAntivirusProfileContentDisarmPdfActGotor(i["pdf-act-gotor"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_java"
	if _, ok := i["pdf-act-java"]; ok {
		result["pdf_act_java"] = flattenAntivirusProfileContentDisarmPdfActJava(i["pdf-act-java"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_launch"
	if _, ok := i["pdf-act-launch"]; ok {
		result["pdf_act_launch"] = flattenAntivirusProfileContentDisarmPdfActLaunch(i["pdf-act-launch"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_movie"
	if _, ok := i["pdf-act-movie"]; ok {
		result["pdf_act_movie"] = flattenAntivirusProfileContentDisarmPdfActMovie(i["pdf-act-movie"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_act_sound"
	if _, ok := i["pdf-act-sound"]; ok {
		result["pdf_act_sound"] = flattenAntivirusProfileContentDisarmPdfActSound(i["pdf-act-sound"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_embedfile"
	if _, ok := i["pdf-embedfile"]; ok {
		result["pdf_embedfile"] = flattenAntivirusProfileContentDisarmPdfEmbedfile(i["pdf-embedfile"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_hyperlink"
	if _, ok := i["pdf-hyperlink"]; ok {
		result["pdf_hyperlink"] = flattenAntivirusProfileContentDisarmPdfHyperlink(i["pdf-hyperlink"], d, pre_append)
	}

	pre_append = pre + ".0." + "pdf_javacode"
	if _, ok := i["pdf-javacode"]; ok {
		result["pdf_javacode"] = flattenAntivirusProfileContentDisarmPdfJavacode(i["pdf-javacode"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileContentDisarmAnalyticsSuspicious(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmCoverPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmDetectOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOfficeAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOfficeDde(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOfficeEmbed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOfficeHylink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOfficeLinked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOfficeMacro(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmOriginalFileDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfActForm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfActGotor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfActJava(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfActLaunch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfActMovie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfActSound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfEmbedfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfHyperlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileContentDisarmPdfJavacode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileEmsThreatFeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileExternalBlocklistArchiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileExternalBlocklistEnableAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFeatureSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortiaiErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortiaiTimeoutAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortindrErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortindrTimeoutAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortisandboxErrorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortisandboxMaxUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortisandboxMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFortisandboxTimeoutAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileFtpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileFtpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileFtpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileFtpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileFtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileFtpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileFtpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileFtpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileFtpFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileFtpMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileFtpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileFtpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileFtpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileFtpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileHttpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileHttpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileHttpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenAntivirusProfileHttpContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileHttpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileHttpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileHttpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileHttpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileHttpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileHttpFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileHttpMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileHttpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileHttpQuarantine(i["quarantine"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_content_encoding"
	if _, ok := i["unknown-content-encoding"]; ok {
		result["unknown_content_encoding"] = flattenAntivirusProfileHttpUnknownContentEncoding(i["unknown-content-encoding"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileHttpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileHttpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileHttpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileHttpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileHttpUnknownContentEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileImapArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileImapArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileImapAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenAntivirusProfileImapContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileImapEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenAntivirusProfileImapExecutables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileImapOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileImapExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileImapFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileImapFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileImapFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileImapMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileImapOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileImapQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileImapArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileImapArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileImapAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapExecutables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileImapExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileImapQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileMapiArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileMapiArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileMapiAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileMapiEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenAntivirusProfileMapiExecutables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileMapiOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileMapiExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileMapiFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileMapiFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileMapiFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileMapiMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileMapiOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileMapiQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileMapiArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileMapiArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileMapiAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiExecutables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileMapiExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMapiQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileMobileMalwareDb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNacQuar(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "expiry"
	if _, ok := i["expiry"]; ok {
		result["expiry"] = flattenAntivirusProfileNacQuarExpiry(i["expiry"], d, pre_append)
	}

	pre_append = pre + ".0." + "infected"
	if _, ok := i["infected"]; ok {
		result["infected"] = flattenAntivirusProfileNacQuarInfected(i["infected"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenAntivirusProfileNacQuarLog(i["log"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileNacQuarExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNacQuarInfected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNacQuarLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileNntpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileNntpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileNntpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileNntpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileNntpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileNntpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileNntpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileNntpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileNntpFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileNntpMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileNntpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileNntpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileNntpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileNntpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileNntpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileNntpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileNntpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileOutbreakPreventionExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := i["ftgd-service"]; ok {
		result["ftgd_service"] = flattenAntivirusProfileOutbreakPreventionFtgdService(i["ftgd-service"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileOutbreakPreventionExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileOutbreakPreventionFtgdService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileOutbreakPreventionArchiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfilePop3ArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfilePop3ArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfilePop3AvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenAntivirusProfilePop3ContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfilePop3Emulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenAntivirusProfilePop3Executables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfilePop3Options(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfilePop3ExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfilePop3Fortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfilePop3Fortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfilePop3Fortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfilePop3MalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfilePop3OutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfilePop3Quarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfilePop3ArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfilePop3ArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfilePop3AvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3ContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Emulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Executables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Options(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfilePop3ExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Fortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Fortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Fortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3MalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3OutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfilePop3Quarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileScanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileSmtpArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileSmtpArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileSmtpAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_disarm"
	if _, ok := i["content-disarm"]; ok {
		result["content_disarm"] = flattenAntivirusProfileSmtpContentDisarm(i["content-disarm"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileSmtpEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "executables"
	if _, ok := i["executables"]; ok {
		result["executables"] = flattenAntivirusProfileSmtpExecutables(i["executables"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileSmtpOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileSmtpExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileSmtpFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileSmtpFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileSmtpFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileSmtpMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileSmtpOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileSmtpQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileSmtpArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileSmtpArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileSmtpAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpContentDisarm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpExecutables(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileSmtpExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSmtpQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := i["archive-block"]; ok {
		result["archive_block"] = flattenAntivirusProfileSshArchiveBlock(i["archive-block"], d, pre_append)
	}

	pre_append = pre + ".0." + "archive_log"
	if _, ok := i["archive-log"]; ok {
		result["archive_log"] = flattenAntivirusProfileSshArchiveLog(i["archive-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_scan"
	if _, ok := i["av-scan"]; ok {
		result["av_scan"] = flattenAntivirusProfileSshAvScan(i["av-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "emulator"
	if _, ok := i["emulator"]; ok {
		result["emulator"] = flattenAntivirusProfileSshEmulator(i["emulator"], d, pre_append)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenAntivirusProfileSshOptions(i["options"], d, pre_append)
	}

	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := i["external-blocklist"]; ok {
		result["external_blocklist"] = flattenAntivirusProfileSshExternalBlocklist(i["external-blocklist"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortiai"
	if _, ok := i["fortiai"]; ok {
		result["fortiai"] = flattenAntivirusProfileSshFortiai(i["fortiai"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortindr"
	if _, ok := i["fortindr"]; ok {
		result["fortindr"] = flattenAntivirusProfileSshFortindr(i["fortindr"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := i["fortisandbox"]; ok {
		result["fortisandbox"] = flattenAntivirusProfileSshFortisandbox(i["fortisandbox"], d, pre_append)
	}

	pre_append = pre + ".0." + "malware_stream"
	if _, ok := i["malware-stream"]; ok {
		result["malware_stream"] = flattenAntivirusProfileSshMalwareStream(i["malware-stream"], d, pre_append)
	}

	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := i["outbreak-prevention"]; ok {
		result["outbreak_prevention"] = flattenAntivirusProfileSshOutbreakPrevention(i["outbreak-prevention"], d, pre_append)
	}

	pre_append = pre + ".0." + "quarantine"
	if _, ok := i["quarantine"]; ok {
		result["quarantine"] = flattenAntivirusProfileSshQuarantine(i["quarantine"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenAntivirusProfileSshArchiveBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileSshArchiveLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileSshAvScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshEmulator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileSshExternalBlocklist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshFortiai(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshFortindr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshFortisandbox(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshMalwareStream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshOutbreakPrevention(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileSshQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileAnalyticsMaxUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileAnalyticsWlFiletype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenAntivirusProfileFtgdAnalytics(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAntivirusProfileAvBlockLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectAntivirusProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("analytics_bl_filetype", flattenAntivirusProfileAnalyticsBlFiletype(o["analytics-bl-filetype"], d, "analytics_bl_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-bl-filetype"], "AntivirusProfile-AnalyticsBlFiletype"); ok {
			if err = d.Set("analytics_bl_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_bl_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_bl_filetype: %v", err)
		}
	}

	if err = d.Set("analytics_accept_filetype", flattenAntivirusProfileAnalyticsAcceptFiletype(o["analytics-accept-filetype"], d, "analytics_accept_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-accept-filetype"], "AntivirusProfile-AnalyticsAcceptFiletype"); ok {
			if err = d.Set("analytics_accept_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_accept_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_accept_filetype: %v", err)
		}
	}

	if err = d.Set("analytics_db", flattenAntivirusProfileAnalyticsDb(o["analytics-db"], d, "analytics_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-db"], "AntivirusProfile-AnalyticsDb"); ok {
			if err = d.Set("analytics_db", vv); err != nil {
				return fmt.Errorf("Error reading analytics_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_db: %v", err)
		}
	}

	if err = d.Set("analytics_ignore_filetype", flattenAntivirusProfileAnalyticsIgnoreFiletype(o["analytics-ignore-filetype"], d, "analytics_ignore_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-ignore-filetype"], "AntivirusProfile-AnalyticsIgnoreFiletype"); ok {
			if err = d.Set("analytics_ignore_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_ignore_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_ignore_filetype: %v", err)
		}
	}

	if err = d.Set("av_virus_log", flattenAntivirusProfileAvVirusLog(o["av-virus-log"], d, "av_virus_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-virus-log"], "AntivirusProfile-AvVirusLog"); ok {
			if err = d.Set("av_virus_log", vv); err != nil {
				return fmt.Errorf("Error reading av_virus_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_virus_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("cifs", flattenAntivirusProfileCifs(o["cifs"], d, "cifs")); err != nil {
			if vv, ok := fortiAPIPatch(o["cifs"], "AntivirusProfile-Cifs"); ok {
				if err = d.Set("cifs", vv); err != nil {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cifs"); ok {
			if err = d.Set("cifs", flattenAntivirusProfileCifs(o["cifs"], d, "cifs")); err != nil {
				if vv, ok := fortiAPIPatch(o["cifs"], "AntivirusProfile-Cifs"); ok {
					if err = d.Set("cifs", vv); err != nil {
						return fmt.Errorf("Error reading cifs: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenAntivirusProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "AntivirusProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("content_disarm", flattenAntivirusProfileContentDisarm(o["content-disarm"], d, "content_disarm")); err != nil {
			if vv, ok := fortiAPIPatch(o["content-disarm"], "AntivirusProfile-ContentDisarm"); ok {
				if err = d.Set("content_disarm", vv); err != nil {
					return fmt.Errorf("Error reading content_disarm: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading content_disarm: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("content_disarm"); ok {
			if err = d.Set("content_disarm", flattenAntivirusProfileContentDisarm(o["content-disarm"], d, "content_disarm")); err != nil {
				if vv, ok := fortiAPIPatch(o["content-disarm"], "AntivirusProfile-ContentDisarm"); ok {
					if err = d.Set("content_disarm", vv); err != nil {
						return fmt.Errorf("Error reading content_disarm: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading content_disarm: %v", err)
				}
			}
		}
	}

	if err = d.Set("ems_threat_feed", flattenAntivirusProfileEmsThreatFeed(o["ems-threat-feed"], d, "ems_threat_feed")); err != nil {
		if vv, ok := fortiAPIPatch(o["ems-threat-feed"], "AntivirusProfile-EmsThreatFeed"); ok {
			if err = d.Set("ems_threat_feed", vv); err != nil {
				return fmt.Errorf("Error reading ems_threat_feed: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ems_threat_feed: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenAntivirusProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["extended-log"], "AntivirusProfile-ExtendedLog"); ok {
			if err = d.Set("extended_log", vv); err != nil {
				return fmt.Errorf("Error reading extended_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("external_blocklist", flattenAntivirusProfileExternalBlocklist(o["external-blocklist"], d, "external_blocklist")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist"], "AntivirusProfile-ExternalBlocklist"); ok {
			if err = d.Set("external_blocklist", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist: %v", err)
		}
	}

	if err = d.Set("external_blocklist_archive_scan", flattenAntivirusProfileExternalBlocklistArchiveScan(o["external-blocklist-archive-scan"], d, "external_blocklist_archive_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist-archive-scan"], "AntivirusProfile-ExternalBlocklistArchiveScan"); ok {
			if err = d.Set("external_blocklist_archive_scan", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist_archive_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist_archive_scan: %v", err)
		}
	}

	if err = d.Set("external_blocklist_enable_all", flattenAntivirusProfileExternalBlocklistEnableAll(o["external-blocklist-enable-all"], d, "external_blocklist_enable_all")); err != nil {
		if vv, ok := fortiAPIPatch(o["external-blocklist-enable-all"], "AntivirusProfile-ExternalBlocklistEnableAll"); ok {
			if err = d.Set("external_blocklist_enable_all", vv); err != nil {
				return fmt.Errorf("Error reading external_blocklist_enable_all: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading external_blocklist_enable_all: %v", err)
		}
	}

	if err = d.Set("feature_set", flattenAntivirusProfileFeatureSet(o["feature-set"], d, "feature_set")); err != nil {
		if vv, ok := fortiAPIPatch(o["feature-set"], "AntivirusProfile-FeatureSet"); ok {
			if err = d.Set("feature_set", vv); err != nil {
				return fmt.Errorf("Error reading feature_set: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("fortiai_error_action", flattenAntivirusProfileFortiaiErrorAction(o["fortiai-error-action"], d, "fortiai_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai-error-action"], "AntivirusProfile-FortiaiErrorAction"); ok {
			if err = d.Set("fortiai_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortiai_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai_error_action: %v", err)
		}
	}

	if err = d.Set("fortiai_timeout_action", flattenAntivirusProfileFortiaiTimeoutAction(o["fortiai-timeout-action"], d, "fortiai_timeout_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiai-timeout-action"], "AntivirusProfile-FortiaiTimeoutAction"); ok {
			if err = d.Set("fortiai_timeout_action", vv); err != nil {
				return fmt.Errorf("Error reading fortiai_timeout_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiai_timeout_action: %v", err)
		}
	}

	if err = d.Set("fortindr_error_action", flattenAntivirusProfileFortindrErrorAction(o["fortindr-error-action"], d, "fortindr_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr-error-action"], "AntivirusProfile-FortindrErrorAction"); ok {
			if err = d.Set("fortindr_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortindr_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr_error_action: %v", err)
		}
	}

	if err = d.Set("fortindr_timeout_action", flattenAntivirusProfileFortindrTimeoutAction(o["fortindr-timeout-action"], d, "fortindr_timeout_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortindr-timeout-action"], "AntivirusProfile-FortindrTimeoutAction"); ok {
			if err = d.Set("fortindr_timeout_action", vv); err != nil {
				return fmt.Errorf("Error reading fortindr_timeout_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortindr_timeout_action: %v", err)
		}
	}

	if err = d.Set("fortisandbox_error_action", flattenAntivirusProfileFortisandboxErrorAction(o["fortisandbox-error-action"], d, "fortisandbox_error_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-error-action"], "AntivirusProfile-FortisandboxErrorAction"); ok {
			if err = d.Set("fortisandbox_error_action", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_error_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_error_action: %v", err)
		}
	}

	if err = d.Set("fortisandbox_max_upload", flattenAntivirusProfileFortisandboxMaxUpload(o["fortisandbox-max-upload"], d, "fortisandbox_max_upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-max-upload"], "AntivirusProfile-FortisandboxMaxUpload"); ok {
			if err = d.Set("fortisandbox_max_upload", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_max_upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_max_upload: %v", err)
		}
	}

	if err = d.Set("fortisandbox_mode", flattenAntivirusProfileFortisandboxMode(o["fortisandbox-mode"], d, "fortisandbox_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-mode"], "AntivirusProfile-FortisandboxMode"); ok {
			if err = d.Set("fortisandbox_mode", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_mode: %v", err)
		}
	}

	if err = d.Set("fortisandbox_timeout_action", flattenAntivirusProfileFortisandboxTimeoutAction(o["fortisandbox-timeout-action"], d, "fortisandbox_timeout_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortisandbox-timeout-action"], "AntivirusProfile-FortisandboxTimeoutAction"); ok {
			if err = d.Set("fortisandbox_timeout_action", vv); err != nil {
				return fmt.Errorf("Error reading fortisandbox_timeout_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortisandbox_timeout_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenAntivirusProfileFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "AntivirusProfile-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenAntivirusProfileFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "AntivirusProfile-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenAntivirusProfileHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "AntivirusProfile-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenAntivirusProfileHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "AntivirusProfile-Http"); ok {
					if err = d.Set("http", vv); err != nil {
						return fmt.Errorf("Error reading http: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imap", flattenAntivirusProfileImap(o["imap"], d, "imap")); err != nil {
			if vv, ok := fortiAPIPatch(o["imap"], "AntivirusProfile-Imap"); ok {
				if err = d.Set("imap", vv); err != nil {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading imap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imap"); ok {
			if err = d.Set("imap", flattenAntivirusProfileImap(o["imap"], d, "imap")); err != nil {
				if vv, ok := fortiAPIPatch(o["imap"], "AntivirusProfile-Imap"); ok {
					if err = d.Set("imap", vv); err != nil {
						return fmt.Errorf("Error reading imap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading imap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenAntivirusProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if vv, ok := fortiAPIPatch(o["mapi"], "AntivirusProfile-Mapi"); ok {
				if err = d.Set("mapi", vv); err != nil {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenAntivirusProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if vv, ok := fortiAPIPatch(o["mapi"], "AntivirusProfile-Mapi"); ok {
					if err = d.Set("mapi", vv); err != nil {
						return fmt.Errorf("Error reading mapi: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if err = d.Set("mobile_malware_db", flattenAntivirusProfileMobileMalwareDb(o["mobile-malware-db"], d, "mobile_malware_db")); err != nil {
		if vv, ok := fortiAPIPatch(o["mobile-malware-db"], "AntivirusProfile-MobileMalwareDb"); ok {
			if err = d.Set("mobile_malware_db", vv); err != nil {
				return fmt.Errorf("Error reading mobile_malware_db: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mobile_malware_db: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nac_quar", flattenAntivirusProfileNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
			if vv, ok := fortiAPIPatch(o["nac-quar"], "AntivirusProfile-NacQuar"); ok {
				if err = d.Set("nac_quar", vv); err != nil {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nac_quar: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_quar"); ok {
			if err = d.Set("nac_quar", flattenAntivirusProfileNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
				if vv, ok := fortiAPIPatch(o["nac-quar"], "AntivirusProfile-NacQuar"); ok {
					if err = d.Set("nac_quar", vv); err != nil {
						return fmt.Errorf("Error reading nac_quar: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenAntivirusProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "AntivirusProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenAntivirusProfileNntp(o["nntp"], d, "nntp")); err != nil {
			if vv, ok := fortiAPIPatch(o["nntp"], "AntivirusProfile-Nntp"); ok {
				if err = d.Set("nntp", vv); err != nil {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenAntivirusProfileNntp(o["nntp"], d, "nntp")); err != nil {
				if vv, ok := fortiAPIPatch(o["nntp"], "AntivirusProfile-Nntp"); ok {
					if err = d.Set("nntp", vv); err != nil {
						return fmt.Errorf("Error reading nntp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("outbreak_prevention", flattenAntivirusProfileOutbreakPrevention(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
			if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "AntivirusProfile-OutbreakPrevention"); ok {
				if err = d.Set("outbreak_prevention", vv); err != nil {
					return fmt.Errorf("Error reading outbreak_prevention: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading outbreak_prevention: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("outbreak_prevention"); ok {
			if err = d.Set("outbreak_prevention", flattenAntivirusProfileOutbreakPrevention(o["outbreak-prevention"], d, "outbreak_prevention")); err != nil {
				if vv, ok := fortiAPIPatch(o["outbreak-prevention"], "AntivirusProfile-OutbreakPrevention"); ok {
					if err = d.Set("outbreak_prevention", vv); err != nil {
						return fmt.Errorf("Error reading outbreak_prevention: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading outbreak_prevention: %v", err)
				}
			}
		}
	}

	if err = d.Set("outbreak_prevention_archive_scan", flattenAntivirusProfileOutbreakPreventionArchiveScan(o["outbreak-prevention-archive-scan"], d, "outbreak_prevention_archive_scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["outbreak-prevention-archive-scan"], "AntivirusProfile-OutbreakPreventionArchiveScan"); ok {
			if err = d.Set("outbreak_prevention_archive_scan", vv); err != nil {
				return fmt.Errorf("Error reading outbreak_prevention_archive_scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading outbreak_prevention_archive_scan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pop3", flattenAntivirusProfilePop3(o["pop3"], d, "pop3")); err != nil {
			if vv, ok := fortiAPIPatch(o["pop3"], "AntivirusProfile-Pop3"); ok {
				if err = d.Set("pop3", vv); err != nil {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading pop3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3"); ok {
			if err = d.Set("pop3", flattenAntivirusProfilePop3(o["pop3"], d, "pop3")); err != nil {
				if vv, ok := fortiAPIPatch(o["pop3"], "AntivirusProfile-Pop3"); ok {
					if err = d.Set("pop3", vv); err != nil {
						return fmt.Errorf("Error reading pop3: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading pop3: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_group", flattenAntivirusProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["replacemsg-group"], "AntivirusProfile-ReplacemsgGroup"); ok {
			if err = d.Set("replacemsg_group", vv); err != nil {
				return fmt.Errorf("Error reading replacemsg_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("scan_mode", flattenAntivirusProfileScanMode(o["scan-mode"], d, "scan_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-mode"], "AntivirusProfile-ScanMode"); ok {
			if err = d.Set("scan_mode", vv); err != nil {
				return fmt.Errorf("Error reading scan_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("smtp", flattenAntivirusProfileSmtp(o["smtp"], d, "smtp")); err != nil {
			if vv, ok := fortiAPIPatch(o["smtp"], "AntivirusProfile-Smtp"); ok {
				if err = d.Set("smtp", vv); err != nil {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading smtp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtp"); ok {
			if err = d.Set("smtp", flattenAntivirusProfileSmtp(o["smtp"], d, "smtp")); err != nil {
				if vv, ok := fortiAPIPatch(o["smtp"], "AntivirusProfile-Smtp"); ok {
					if err = d.Set("smtp", vv); err != nil {
						return fmt.Errorf("Error reading smtp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading smtp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenAntivirusProfileSsh(o["ssh"], d, "ssh")); err != nil {
			if vv, ok := fortiAPIPatch(o["ssh"], "AntivirusProfile-Ssh"); ok {
				if err = d.Set("ssh", vv); err != nil {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenAntivirusProfileSsh(o["ssh"], d, "ssh")); err != nil {
				if vv, ok := fortiAPIPatch(o["ssh"], "AntivirusProfile-Ssh"); ok {
					if err = d.Set("ssh", vv); err != nil {
						return fmt.Errorf("Error reading ssh: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if err = d.Set("analytics_max_upload", flattenAntivirusProfileAnalyticsMaxUpload(o["analytics-max-upload"], d, "analytics_max_upload")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-max-upload"], "AntivirusProfile-AnalyticsMaxUpload"); ok {
			if err = d.Set("analytics_max_upload", vv); err != nil {
				return fmt.Errorf("Error reading analytics_max_upload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_max_upload: %v", err)
		}
	}

	if err = d.Set("analytics_wl_filetype", flattenAntivirusProfileAnalyticsWlFiletype(o["analytics-wl-filetype"], d, "analytics_wl_filetype")); err != nil {
		if vv, ok := fortiAPIPatch(o["analytics-wl-filetype"], "AntivirusProfile-AnalyticsWlFiletype"); ok {
			if err = d.Set("analytics_wl_filetype", vv); err != nil {
				return fmt.Errorf("Error reading analytics_wl_filetype: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading analytics_wl_filetype: %v", err)
		}
	}

	if err = d.Set("ftgd_analytics", flattenAntivirusProfileFtgdAnalytics(o["ftgd-analytics"], d, "ftgd_analytics")); err != nil {
		if vv, ok := fortiAPIPatch(o["ftgd-analytics"], "AntivirusProfile-FtgdAnalytics"); ok {
			if err = d.Set("ftgd_analytics", vv); err != nil {
				return fmt.Errorf("Error reading ftgd_analytics: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ftgd_analytics: %v", err)
		}
	}

	if err = d.Set("av_block_log", flattenAntivirusProfileAvBlockLog(o["av-block-log"], d, "av_block_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["av-block-log"], "AntivirusProfile-AvBlockLog"); ok {
			if err = d.Set("av_block_log", vv); err != nil {
				return fmt.Errorf("Error reading av_block_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_block_log: %v", err)
		}
	}

	return nil
}

func flattenAntivirusProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandAntivirusProfileAnalyticsBlFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileAnalyticsAcceptFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileAnalyticsDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileAnalyticsIgnoreFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileAvVirusLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileCifsArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileCifsArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileCifsAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileCifsEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileCifsOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileCifsExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileCifsFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileCifsFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileCifsFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileCifsMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileCifsOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileCifsQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileCifsArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileCifsArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileCifsAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileCifsExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileCifsQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "analytics_suspicious"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["analytics-suspicious"], _ = expandAntivirusProfileContentDisarmAnalyticsSuspicious(d, i["analytics_suspicious"], pre_append)
	}
	pre_append = pre + ".0." + "cover_page"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["cover-page"], _ = expandAntivirusProfileContentDisarmCoverPage(d, i["cover_page"], pre_append)
	}
	pre_append = pre + ".0." + "detect_only"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["detect-only"], _ = expandAntivirusProfileContentDisarmDetectOnly(d, i["detect_only"], pre_append)
	}
	pre_append = pre + ".0." + "error_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["error-action"], _ = expandAntivirusProfileContentDisarmErrorAction(d, i["error_action"], pre_append)
	}
	pre_append = pre + ".0." + "office_action"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-action"], _ = expandAntivirusProfileContentDisarmOfficeAction(d, i["office_action"], pre_append)
	}
	pre_append = pre + ".0." + "office_dde"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-dde"], _ = expandAntivirusProfileContentDisarmOfficeDde(d, i["office_dde"], pre_append)
	}
	pre_append = pre + ".0." + "office_embed"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-embed"], _ = expandAntivirusProfileContentDisarmOfficeEmbed(d, i["office_embed"], pre_append)
	}
	pre_append = pre + ".0." + "office_hylink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-hylink"], _ = expandAntivirusProfileContentDisarmOfficeHylink(d, i["office_hylink"], pre_append)
	}
	pre_append = pre + ".0." + "office_linked"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-linked"], _ = expandAntivirusProfileContentDisarmOfficeLinked(d, i["office_linked"], pre_append)
	}
	pre_append = pre + ".0." + "office_macro"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["office-macro"], _ = expandAntivirusProfileContentDisarmOfficeMacro(d, i["office_macro"], pre_append)
	}
	pre_append = pre + ".0." + "original_file_destination"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["original-file-destination"], _ = expandAntivirusProfileContentDisarmOriginalFileDestination(d, i["original_file_destination"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_form"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-form"], _ = expandAntivirusProfileContentDisarmPdfActForm(d, i["pdf_act_form"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_gotor"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-gotor"], _ = expandAntivirusProfileContentDisarmPdfActGotor(d, i["pdf_act_gotor"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_java"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-java"], _ = expandAntivirusProfileContentDisarmPdfActJava(d, i["pdf_act_java"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_launch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-launch"], _ = expandAntivirusProfileContentDisarmPdfActLaunch(d, i["pdf_act_launch"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_movie"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-movie"], _ = expandAntivirusProfileContentDisarmPdfActMovie(d, i["pdf_act_movie"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_act_sound"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-act-sound"], _ = expandAntivirusProfileContentDisarmPdfActSound(d, i["pdf_act_sound"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_embedfile"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-embedfile"], _ = expandAntivirusProfileContentDisarmPdfEmbedfile(d, i["pdf_embedfile"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_hyperlink"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-hyperlink"], _ = expandAntivirusProfileContentDisarmPdfHyperlink(d, i["pdf_hyperlink"], pre_append)
	}
	pre_append = pre + ".0." + "pdf_javacode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["pdf-javacode"], _ = expandAntivirusProfileContentDisarmPdfJavacode(d, i["pdf_javacode"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileContentDisarmAnalyticsSuspicious(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmCoverPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmDetectOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOfficeAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOfficeDde(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOfficeEmbed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOfficeHylink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOfficeLinked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOfficeMacro(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmOriginalFileDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfActForm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfActGotor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfActJava(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfActLaunch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfActMovie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfActSound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfEmbedfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfHyperlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileContentDisarmPdfJavacode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileEmsThreatFeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileExternalBlocklistArchiveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileExternalBlocklistEnableAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFeatureSet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortiaiErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortiaiTimeoutAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortindrErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortindrTimeoutAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortisandboxErrorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortisandboxMaxUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortisandboxMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFortisandboxTimeoutAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileFtpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileFtpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileFtpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileFtpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileFtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileFtpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileFtpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileFtpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileFtpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileFtpMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileFtpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileFtpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileFtpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileFtpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileHttpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileHttpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileHttpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandAntivirusProfileHttpContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileHttpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileHttpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileHttpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileHttpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileHttpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileHttpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileHttpMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileHttpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileHttpQuarantine(d, i["quarantine"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_content_encoding"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["unknown-content-encoding"], _ = expandAntivirusProfileHttpUnknownContentEncoding(d, i["unknown_content_encoding"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileHttpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileHttpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileHttpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileHttpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileHttpUnknownContentEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileImapArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileImapArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileImapAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandAntivirusProfileImapContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileImapEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandAntivirusProfileImapExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileImapOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileImapExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileImapFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileImapFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileImapFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileImapMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileImapOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileImapQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileImapArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileImapArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileImapAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapExecutables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileImapExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileImapQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileMapiArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileMapiArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileMapiAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileMapiEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandAntivirusProfileMapiExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileMapiOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileMapiExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileMapiFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileMapiFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileMapiFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileMapiMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileMapiOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileMapiQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileMapiArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileMapiArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileMapiAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiExecutables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileMapiExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMapiQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileMobileMalwareDb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNacQuar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "expiry"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["expiry"], _ = expandAntivirusProfileNacQuarExpiry(d, i["expiry"], pre_append)
	}
	pre_append = pre + ".0." + "infected"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["infected"], _ = expandAntivirusProfileNacQuarInfected(d, i["infected"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["log"], _ = expandAntivirusProfileNacQuarLog(d, i["log"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileNacQuarExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNacQuarInfected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNacQuarLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileNntpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileNntpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileNntpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileNntpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileNntpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileNntpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileNntpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileNntpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileNntpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileNntpMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileNntpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileNntpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileNntpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileNntpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileNntpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileNntpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileNntpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileOutbreakPreventionExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "ftgd_service"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ftgd-service"], _ = expandAntivirusProfileOutbreakPreventionFtgdService(d, i["ftgd_service"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileOutbreakPreventionExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileOutbreakPreventionFtgdService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileOutbreakPreventionArchiveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfilePop3ArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfilePop3ArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfilePop3AvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandAntivirusProfilePop3ContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfilePop3Emulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandAntivirusProfilePop3Executables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfilePop3Options(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfilePop3ExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfilePop3Fortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfilePop3Fortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfilePop3Fortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfilePop3MalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfilePop3OutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfilePop3Quarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfilePop3ArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfilePop3ArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfilePop3AvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3ContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Emulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Executables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Options(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfilePop3ExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Fortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Fortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Fortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3MalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3OutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfilePop3Quarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileScanMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileSmtpArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileSmtpArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileSmtpAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "content_disarm"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["content-disarm"], _ = expandAntivirusProfileSmtpContentDisarm(d, i["content_disarm"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileSmtpEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "executables"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["executables"], _ = expandAntivirusProfileSmtpExecutables(d, i["executables"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileSmtpOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileSmtpExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileSmtpFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileSmtpFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileSmtpFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileSmtpMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileSmtpOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileSmtpQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileSmtpArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileSmtpArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileSmtpAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpContentDisarm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpExecutables(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileSmtpExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSmtpQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "archive_block"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-block"], _ = expandAntivirusProfileSshArchiveBlock(d, i["archive_block"], pre_append)
	}
	pre_append = pre + ".0." + "archive_log"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["archive-log"], _ = expandAntivirusProfileSshArchiveLog(d, i["archive_log"], pre_append)
	}
	pre_append = pre + ".0." + "av_scan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["av-scan"], _ = expandAntivirusProfileSshAvScan(d, i["av_scan"], pre_append)
	}
	pre_append = pre + ".0." + "emulator"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["emulator"], _ = expandAntivirusProfileSshEmulator(d, i["emulator"], pre_append)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["options"], _ = expandAntivirusProfileSshOptions(d, i["options"], pre_append)
	}
	pre_append = pre + ".0." + "external_blocklist"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["external-blocklist"], _ = expandAntivirusProfileSshExternalBlocklist(d, i["external_blocklist"], pre_append)
	}
	pre_append = pre + ".0." + "fortiai"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortiai"], _ = expandAntivirusProfileSshFortiai(d, i["fortiai"], pre_append)
	}
	pre_append = pre + ".0." + "fortindr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortindr"], _ = expandAntivirusProfileSshFortindr(d, i["fortindr"], pre_append)
	}
	pre_append = pre + ".0." + "fortisandbox"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["fortisandbox"], _ = expandAntivirusProfileSshFortisandbox(d, i["fortisandbox"], pre_append)
	}
	pre_append = pre + ".0." + "malware_stream"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["malware-stream"], _ = expandAntivirusProfileSshMalwareStream(d, i["malware_stream"], pre_append)
	}
	pre_append = pre + ".0." + "outbreak_prevention"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["outbreak-prevention"], _ = expandAntivirusProfileSshOutbreakPrevention(d, i["outbreak_prevention"], pre_append)
	}
	pre_append = pre + ".0." + "quarantine"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["quarantine"], _ = expandAntivirusProfileSshQuarantine(d, i["quarantine"], pre_append)
	}

	return result, nil
}

func expandAntivirusProfileSshArchiveBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileSshArchiveLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileSshAvScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshEmulator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileSshExternalBlocklist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshFortiai(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshFortindr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshFortisandbox(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshMalwareStream(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshOutbreakPrevention(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileSshQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileAnalyticsMaxUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileAnalyticsWlFiletype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandAntivirusProfileFtgdAnalytics(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAntivirusProfileAvBlockLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("analytics_bl_filetype"); ok || d.HasChange("analytics_bl_filetype") {
		t, err := expandAntivirusProfileAnalyticsBlFiletype(d, v, "analytics_bl_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-bl-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_accept_filetype"); ok || d.HasChange("analytics_accept_filetype") {
		t, err := expandAntivirusProfileAnalyticsAcceptFiletype(d, v, "analytics_accept_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-accept-filetype"] = t
		}
	}

	if v, ok := d.GetOk("analytics_db"); ok || d.HasChange("analytics_db") {
		t, err := expandAntivirusProfileAnalyticsDb(d, v, "analytics_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-db"] = t
		}
	}

	if v, ok := d.GetOk("analytics_ignore_filetype"); ok || d.HasChange("analytics_ignore_filetype") {
		t, err := expandAntivirusProfileAnalyticsIgnoreFiletype(d, v, "analytics_ignore_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-ignore-filetype"] = t
		}
	}

	if v, ok := d.GetOk("av_virus_log"); ok || d.HasChange("av_virus_log") {
		t, err := expandAntivirusProfileAvVirusLog(d, v, "av_virus_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-virus-log"] = t
		}
	}

	if v, ok := d.GetOk("cifs"); ok || d.HasChange("cifs") {
		t, err := expandAntivirusProfileCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandAntivirusProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("content_disarm"); ok || d.HasChange("content_disarm") {
		t, err := expandAntivirusProfileContentDisarm(d, v, "content_disarm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["content-disarm"] = t
		}
	}

	if v, ok := d.GetOk("ems_threat_feed"); ok || d.HasChange("ems_threat_feed") {
		t, err := expandAntivirusProfileEmsThreatFeed(d, v, "ems_threat_feed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-threat-feed"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok || d.HasChange("extended_log") {
		t, err := expandAntivirusProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist"); ok || d.HasChange("external_blocklist") {
		t, err := expandAntivirusProfileExternalBlocklist(d, v, "external_blocklist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist_archive_scan"); ok || d.HasChange("external_blocklist_archive_scan") {
		t, err := expandAntivirusProfileExternalBlocklistArchiveScan(d, v, "external_blocklist_archive_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist-archive-scan"] = t
		}
	}

	if v, ok := d.GetOk("external_blocklist_enable_all"); ok || d.HasChange("external_blocklist_enable_all") {
		t, err := expandAntivirusProfileExternalBlocklistEnableAll(d, v, "external_blocklist_enable_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-blocklist-enable-all"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok || d.HasChange("feature_set") {
		t, err := expandAntivirusProfileFeatureSet(d, v, "feature_set")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("fortiai_error_action"); ok || d.HasChange("fortiai_error_action") {
		t, err := expandAntivirusProfileFortiaiErrorAction(d, v, "fortiai_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai-error-action"] = t
		}
	}

	if v, ok := d.GetOk("fortiai_timeout_action"); ok || d.HasChange("fortiai_timeout_action") {
		t, err := expandAntivirusProfileFortiaiTimeoutAction(d, v, "fortiai_timeout_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiai-timeout-action"] = t
		}
	}

	if v, ok := d.GetOk("fortindr_error_action"); ok || d.HasChange("fortindr_error_action") {
		t, err := expandAntivirusProfileFortindrErrorAction(d, v, "fortindr_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr-error-action"] = t
		}
	}

	if v, ok := d.GetOk("fortindr_timeout_action"); ok || d.HasChange("fortindr_timeout_action") {
		t, err := expandAntivirusProfileFortindrTimeoutAction(d, v, "fortindr_timeout_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortindr-timeout-action"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_error_action"); ok || d.HasChange("fortisandbox_error_action") {
		t, err := expandAntivirusProfileFortisandboxErrorAction(d, v, "fortisandbox_error_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-error-action"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_max_upload"); ok || d.HasChange("fortisandbox_max_upload") {
		t, err := expandAntivirusProfileFortisandboxMaxUpload(d, v, "fortisandbox_max_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-max-upload"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_mode"); ok || d.HasChange("fortisandbox_mode") {
		t, err := expandAntivirusProfileFortisandboxMode(d, v, "fortisandbox_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-mode"] = t
		}
	}

	if v, ok := d.GetOk("fortisandbox_timeout_action"); ok || d.HasChange("fortisandbox_timeout_action") {
		t, err := expandAntivirusProfileFortisandboxTimeoutAction(d, v, "fortisandbox_timeout_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-timeout-action"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok || d.HasChange("ftp") {
		t, err := expandAntivirusProfileFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandAntivirusProfileHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("imap"); ok || d.HasChange("imap") {
		t, err := expandAntivirusProfileImap(d, v, "imap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok || d.HasChange("mapi") {
		t, err := expandAntivirusProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("mobile_malware_db"); ok || d.HasChange("mobile_malware_db") {
		t, err := expandAntivirusProfileMobileMalwareDb(d, v, "mobile_malware_db")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mobile-malware-db"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok || d.HasChange("nac_quar") {
		t, err := expandAntivirusProfileNacQuar(d, v, "nac_quar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandAntivirusProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok || d.HasChange("nntp") {
		t, err := expandAntivirusProfileNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention"); ok || d.HasChange("outbreak_prevention") {
		t, err := expandAntivirusProfileOutbreakPrevention(d, v, "outbreak_prevention")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention"] = t
		}
	}

	if v, ok := d.GetOk("outbreak_prevention_archive_scan"); ok || d.HasChange("outbreak_prevention_archive_scan") {
		t, err := expandAntivirusProfileOutbreakPreventionArchiveScan(d, v, "outbreak_prevention_archive_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbreak-prevention-archive-scan"] = t
		}
	}

	if v, ok := d.GetOk("pop3"); ok || d.HasChange("pop3") {
		t, err := expandAntivirusProfilePop3(d, v, "pop3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok || d.HasChange("replacemsg_group") {
		t, err := expandAntivirusProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("scan_mode"); ok || d.HasChange("scan_mode") {
		t, err := expandAntivirusProfileScanMode(d, v, "scan_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-mode"] = t
		}
	}

	if v, ok := d.GetOk("smtp"); ok || d.HasChange("smtp") {
		t, err := expandAntivirusProfileSmtp(d, v, "smtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok || d.HasChange("ssh") {
		t, err := expandAntivirusProfileSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("analytics_max_upload"); ok || d.HasChange("analytics_max_upload") {
		t, err := expandAntivirusProfileAnalyticsMaxUpload(d, v, "analytics_max_upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-max-upload"] = t
		}
	}

	if v, ok := d.GetOk("analytics_wl_filetype"); ok || d.HasChange("analytics_wl_filetype") {
		t, err := expandAntivirusProfileAnalyticsWlFiletype(d, v, "analytics_wl_filetype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["analytics-wl-filetype"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_analytics"); ok || d.HasChange("ftgd_analytics") {
		t, err := expandAntivirusProfileFtgdAnalytics(d, v, "ftgd_analytics")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-analytics"] = t
		}
	}

	if v, ok := d.GetOk("av_block_log"); ok || d.HasChange("av_block_log") {
		t, err := expandAntivirusProfileAvBlockLog(d, v, "av_block_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-block-log"] = t
		}
	}

	return &obj, nil
}
