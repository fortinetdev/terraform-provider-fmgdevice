// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure redundant Internet connections with multiple outbound links and health-check profiles.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdwan() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanUpdate,
		Read:   resourceSystemSdwanRead,
		Update: resourceSystemSdwanUpdate,
		Delete: resourceSystemSdwanDelete,

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
			"app_perf_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"duplication": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dstaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dstintf": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"packet_de_duplication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_duplication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"service_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sla_match_service": &schema.Schema{
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
						"srcaddr6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"srcintf": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"duplication_max_discrepancy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"duplication_max_num": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fail_alert_interfaces": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"class_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"detect_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"diffservcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_match_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_request_domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"embed_measured_health": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"failtime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortiguard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiguard_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ftp_file": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ftp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ha_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"http_agent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_get": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mos_codec": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"packet_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"probe_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"probe_packets": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"probe_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quality_measured_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"recoverytime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"security_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sla": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"jitter_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"latency_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"link_cost_factor": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"mos_threshold": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"packetloss_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"priority_in_sla": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"priority_out_sla": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sla_fail_log_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sla_id_redistribute": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sla_pass_log_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"threshold_alert_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_alert_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_alert_packetloss": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_packetloss": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"update_cascade_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_static_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"health_check_fortiguard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"class_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"detect_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"diffservcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_match_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_request_domain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"embed_measured_health": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"failtime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ftp_file": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ftp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ha_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"http_agent": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_get": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"mos_codec": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_size": &schema.Schema{
							Type:     schema.TypeInt,
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
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"probe_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"probe_packets": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"probe_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quality_measured_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"recoverytime": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"security_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sla": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"jitter_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"latency_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"link_cost_factor": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"mos_threshold": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"packetloss_threshold": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"priority_in_sla": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"priority_out_sla": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sla_fail_log_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sla_id_redistribute": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"sla_pass_log_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"system_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"threshold_alert_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_alert_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_alert_packetloss": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_packetloss": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"update_cascade_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_static_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vrf": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"load_balance_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ingress_spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"preferred_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"priority_in_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority_out_sla": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority6": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"seq_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spillover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transport_group": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"volume_ratio": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"zone": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_check": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"minimum_sla_meet_members": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"route_metric": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_id": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sla_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"neighbor_hold_boot_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"neighbor_hold_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor_hold_down_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"agent_exclusive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_forward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_forward_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp_reverse": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_reverse_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dst_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"end_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"end_src_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"groups": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"hash_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"health_check": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"hold_down_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"input_device": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"input_device_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"input_zone": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"internet_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"internet_service_app_ctrl": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"internet_service_app_ctrl_category": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeInt},
							Optional: true,
							Computed: true,
						},
						"internet_service_app_ctrl_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"internet_service_custom": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"internet_service_custom_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"internet_service_group": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"internet_service_name": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"jitter_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"latency_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"link_cost_factor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link_cost_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"load_balance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"minimum_sla_meet_members": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"packet_loss_weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"passive_measurement": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority_members": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"priority_zone": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"quality_link": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"shortcut": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shortcut_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sla": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"health_check": &schema.Schema{
										Type:     schema.TypeSet,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sla_compare_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sla_stickiness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"src_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src6": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"standalone_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_src_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tie_break": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tos_mask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"use_shortcut_sla": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"users": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"zone_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"speedtest_bypass_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advpn_health_check": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"advpn_select": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"minimum_sla_meet_members": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_sla_tie_break": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceSystemSdwanUpdate(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}
	obj, err := getObjectSystemSdwan(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwan resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSdwan(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSdwan")

	return resourceSystemSdwanRead(d, m)
}

func resourceSystemSdwanDelete(d *schema.ResourceData, m interface{}) error {
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
	device_vdom, err := getVariable(cfg, d, "device_vdom")
	if err != nil {
		return err
	}
	paradict["device"] = device_name
	paradict["vdom"] = device_vdom

	if cfg.Adom != "" {
		wsParams["adom"] = fmt.Sprintf("adom/%s", cfg.Adom)
	}

	err = c.DeleteSystemSdwan(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdwan(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwan resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanAppPerfLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenSystemSdwanDuplicationDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {
			v := flattenSystemSdwanDuplicationDstaddr6(i["dstaddr6"], d, pre_append)
			tmp["dstaddr6"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Dstaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := i["dstintf"]; ok {
			v := flattenSystemSdwanDuplicationDstintf(i["dstintf"], d, pre_append)
			tmp["dstintf"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Dstintf")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdwanDuplicationId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := i["packet-de-duplication"]; ok {
			v := flattenSystemSdwanDuplicationPacketDeDuplication(i["packet-de-duplication"], d, pre_append)
			tmp["packet_de_duplication"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-PacketDeDuplication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := i["packet-duplication"]; ok {
			v := flattenSystemSdwanDuplicationPacketDuplication(i["packet-duplication"], d, pre_append)
			tmp["packet_duplication"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-PacketDuplication")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {
			v := flattenSystemSdwanDuplicationService(i["service"], d, pre_append)
			tmp["service"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Service")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenSystemSdwanDuplicationServiceId(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-ServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_match_service"
		if _, ok := i["sla-match-service"]; ok {
			v := flattenSystemSdwanDuplicationSlaMatchService(i["sla-match-service"], d, pre_append)
			tmp["sla_match_service"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-SlaMatchService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			v := flattenSystemSdwanDuplicationSrcaddr(i["srcaddr"], d, pre_append)
			tmp["srcaddr"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Srcaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {
			v := flattenSystemSdwanDuplicationSrcaddr6(i["srcaddr6"], d, pre_append)
			tmp["srcaddr6"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Srcaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := i["srcintf"]; ok {
			v := flattenSystemSdwanDuplicationSrcintf(i["srcintf"], d, pre_append)
			tmp["srcintf"] = fortiAPISubPartPatch(v, "SystemSdwan-Duplication-Srcintf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationPacketDeDuplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationPacketDuplication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationSlaMatchService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanDuplicationMaxDiscrepancy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationMaxNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanFailDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheck(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenSystemSdwanHealthCheckAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenSystemSdwanHealthCheckClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := i["detect-mode"]; ok {
			v := flattenSystemSdwanHealthCheckDetectMode(i["detect-mode"], d, pre_append)
			tmp["detect_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-DetectMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenSystemSdwanHealthCheckDiffservcode(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := i["dns-match-ip"]; ok {
			v := flattenSystemSdwanHealthCheckDnsMatchIp(i["dns-match-ip"], d, pre_append)
			tmp["dns_match_ip"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-DnsMatchIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := i["dns-request-domain"]; ok {
			v := flattenSystemSdwanHealthCheckDnsRequestDomain(i["dns-request-domain"], d, pre_append)
			tmp["dns_request_domain"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-DnsRequestDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := i["embed-measured-health"]; ok {
			v := flattenSystemSdwanHealthCheckEmbedMeasuredHealth(i["embed-measured-health"], d, pre_append)
			tmp["embed_measured_health"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-EmbedMeasuredHealth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenSystemSdwanHealthCheckFailtime(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard"
		if _, ok := i["fortiguard"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguard(i["fortiguard"], d, pre_append)
			tmp["fortiguard"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Fortiguard")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_name"
		if _, ok := i["fortiguard-name"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardName(i["fortiguard-name"], d, pre_append)
			tmp["fortiguard_name"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-FortiguardName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := i["ftp-file"]; ok {
			v := flattenSystemSdwanHealthCheckFtpFile(i["ftp-file"], d, pre_append)
			tmp["ftp_file"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-FtpFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := i["ftp-mode"]; ok {
			v := flattenSystemSdwanHealthCheckFtpMode(i["ftp-mode"], d, pre_append)
			tmp["ftp_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-FtpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenSystemSdwanHealthCheckHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenSystemSdwanHealthCheckHttpAgent(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenSystemSdwanHealthCheckHttpGet(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenSystemSdwanHealthCheckHttpMatch(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenSystemSdwanHealthCheckInterval(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenSystemSdwanHealthCheckMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := i["mos-codec"]; ok {
			v := flattenSystemSdwanHealthCheckMosCodec(i["mos-codec"], d, pre_append)
			tmp["mos_codec"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-MosCodec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdwanHealthCheckName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenSystemSdwanHealthCheckPacketSize(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSystemSdwanHealthCheckPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := i["probe-count"]; ok {
			v := flattenSystemSdwanHealthCheckProbeCount(i["probe-count"], d, pre_append)
			tmp["probe_count"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ProbeCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenSystemSdwanHealthCheckProbePackets(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenSystemSdwanHealthCheckProbeTimeout(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenSystemSdwanHealthCheckProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := i["quality-measured-method"]; ok {
			v := flattenSystemSdwanHealthCheckQualityMeasuredMethod(i["quality-measured-method"], d, pre_append)
			tmp["quality_measured_method"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-QualityMeasuredMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenSystemSdwanHealthCheckRecoverytime(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenSystemSdwanHealthCheckSecurityMode(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSystemSdwanHealthCheckServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenSystemSdwanHealthCheckSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenSystemSdwanHealthCheckSlaFailLogPeriod(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := i["sla-id-redistribute"]; ok {
			v := flattenSystemSdwanHealthCheckSlaIdRedistribute(i["sla-id-redistribute"], d, pre_append)
			tmp["sla_id_redistribute"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-SlaIdRedistribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPassLogPeriod(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenSystemSdwanHealthCheckSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenSystemSdwanHealthCheckSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := i["system-dns"]; ok {
			v := flattenSystemSdwanHealthCheckSystemDns(i["system-dns"], d, pre_append)
			tmp["system_dns"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-SystemDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenSystemSdwanHealthCheckThresholdAlertJitter(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenSystemSdwanHealthCheckThresholdAlertLatency(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenSystemSdwanHealthCheckThresholdAlertPacketloss(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenSystemSdwanHealthCheckThresholdWarningJitter(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenSystemSdwanHealthCheckThresholdWarningLatency(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenSystemSdwanHealthCheckThresholdWarningPacketloss(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenSystemSdwanHealthCheckUpdateCascadeInterface(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenSystemSdwanHealthCheckUpdateStaticRoute(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-UpdateStaticRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenSystemSdwanHealthCheckUser(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-User")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenSystemSdwanHealthCheckVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheck-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckDetectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckEmbedMeasuredHealth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFtpFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenSystemSdwanHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckMosCodec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbeCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdwanHealthCheckSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenSystemSdwanHealthCheckSlaLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaMosThreshold(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPriorityInSla(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenSystemSdwanHealthCheckSlaPriorityOutSla(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheck-Sla-PriorityOutSla")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckSlaMosThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPriorityInSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPriorityOutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaIdRedistribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSystemDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguard(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := i["class-id"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardClassId(i["class-id"], d, pre_append)
			tmp["class_id"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ClassId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := i["detect-mode"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardDetectMode(i["detect-mode"], d, pre_append)
			tmp["detect_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-DetectMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardDiffservcode(i["diffservcode"], d, pre_append)
			tmp["diffservcode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Diffservcode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := i["dns-match-ip"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardDnsMatchIp(i["dns-match-ip"], d, pre_append)
			tmp["dns_match_ip"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-DnsMatchIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := i["dns-request-domain"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardDnsRequestDomain(i["dns-request-domain"], d, pre_append)
			tmp["dns_request_domain"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-DnsRequestDomain")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := i["embed-measured-health"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(i["embed-measured-health"], d, pre_append)
			tmp["embed_measured_health"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-EmbedMeasuredHealth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardFailtime(i["failtime"], d, pre_append)
			tmp["failtime"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Failtime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := i["ftp-file"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardFtpFile(i["ftp-file"], d, pre_append)
			tmp["ftp_file"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-FtpFile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := i["ftp-mode"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardFtpMode(i["ftp-mode"], d, pre_append)
			tmp["ftp_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-FtpMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardHaPriority(i["ha-priority"], d, pre_append)
			tmp["ha_priority"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-HaPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardHttpAgent(i["http-agent"], d, pre_append)
			tmp["http_agent"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-HttpAgent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardHttpGet(i["http-get"], d, pre_append)
			tmp["http_get"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-HttpGet")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardHttpMatch(i["http-match"], d, pre_append)
			tmp["http_match"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-HttpMatch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardInterval(i["interval"], d, pre_append)
			tmp["interval"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Interval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardMembers(i["members"], d, pre_append)
			tmp["members"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Members")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := i["mos-codec"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardMosCodec(i["mos-codec"], d, pre_append)
			tmp["mos_codec"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-MosCodec")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardPacketSize(i["packet-size"], d, pre_append)
			tmp["packet_size"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-PacketSize")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := i["probe-count"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardProbeCount(i["probe-count"], d, pre_append)
			tmp["probe_count"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ProbeCount")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardProbePackets(i["probe-packets"], d, pre_append)
			tmp["probe_packets"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ProbePackets")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardProbeTimeout(i["probe-timeout"], d, pre_append)
			tmp["probe_timeout"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ProbeTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := i["quality-measured-method"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(i["quality-measured-method"], d, pre_append)
			tmp["quality_measured_method"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-QualityMeasuredMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardRecoverytime(i["recoverytime"], d, pre_append)
			tmp["recoverytime"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Recoverytime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSecurityMode(i["security-mode"], d, pre_append)
			tmp["security_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-SecurityMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(i["sla-fail-log-period"], d, pre_append)
			tmp["sla_fail_log_period"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-SlaFailLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := i["sla-id-redistribute"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaIdRedistribute(i["sla-id-redistribute"], d, pre_append)
			tmp["sla_id_redistribute"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-SlaIdRedistribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(i["sla-pass-log-period"], d, pre_append)
			tmp["sla_pass_log_period"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-SlaPassLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := i["system-dns"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSystemDns(i["system-dns"], d, pre_append)
			tmp["system_dns"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-SystemDns")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_name"
		if _, ok := i["target-name"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardTargetName(i["target-name"], d, pre_append)
			tmp["target_name"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-TargetName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardThresholdAlertJitter(i["threshold-alert-jitter"], d, pre_append)
			tmp["threshold_alert_jitter"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ThresholdAlertJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardThresholdAlertLatency(i["threshold-alert-latency"], d, pre_append)
			tmp["threshold_alert_latency"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ThresholdAlertLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(i["threshold-alert-packetloss"], d, pre_append)
			tmp["threshold_alert_packetloss"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ThresholdAlertPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardThresholdWarningJitter(i["threshold-warning-jitter"], d, pre_append)
			tmp["threshold_warning_jitter"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ThresholdWarningJitter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardThresholdWarningLatency(i["threshold-warning-latency"], d, pre_append)
			tmp["threshold_warning_latency"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ThresholdWarningLatency")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(i["threshold-warning-packetloss"], d, pre_append)
			tmp["threshold_warning_packetloss"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-ThresholdWarningPacketloss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(i["update-cascade-interface"], d, pre_append)
			tmp["update_cascade_interface"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-UpdateCascadeInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardUpdateStaticRoute(i["update-static-route"], d, pre_append)
			tmp["update_static_route"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-UpdateStaticRoute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardUser(i["user"], d, pre_append)
			tmp["user"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-User")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardVrf(i["vrf"], d, pre_append)
			tmp["vrf"] = fortiAPISubPartPatch(v, "SystemSdwan-HealthCheckFortiguard-Vrf")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckFortiguardAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardDetectMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardFtpFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardFtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenSystemSdwanHealthCheckFortiguardHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardMosCodec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProbeCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProbePackets(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProbeTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaJitterThreshold(i["jitter-threshold"], d, pre_append)
			tmp["jitter_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-JitterThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(i["latency-threshold"], d, pre_append)
			tmp["latency_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-LatencyThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := i["mos-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaMosThreshold(i["mos-threshold"], d, pre_append)
			tmp["mos_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-MosThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append)
			tmp["packetloss_threshold"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-PacketlossThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPriorityInSla(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "SystemSdwanHealthCheckFortiguard-Sla-PriorityOutSla")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckFortiguardSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanHealthCheckFortiguardSlaMosThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPriorityInSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaIdRedistribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardSystemDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardTargetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFortiguardVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanLoadBalanceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenSystemSdwanMembersComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			v := flattenSystemSdwanMembersCost(i["cost"], d, pre_append)
			tmp["cost"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Cost")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenSystemSdwanMembersGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			v := flattenSystemSdwanMembersGateway6(i["gateway6"], d, pre_append)
			tmp["gateway6"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Gateway6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {
			v := flattenSystemSdwanMembersIngressSpilloverThreshold(i["ingress-spillover-threshold"], d, pre_append)
			tmp["ingress_spillover_threshold"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-IngressSpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenSystemSdwanMembersInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_source"
		if _, ok := i["preferred-source"]; ok {
			v := flattenSystemSdwanMembersPreferredSource(i["preferred-source"], d, pre_append)
			tmp["preferred_source"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-PreferredSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			v := flattenSystemSdwanMembersPriority(i["priority"], d, pre_append)
			tmp["priority"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Priority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := i["priority-in-sla"]; ok {
			v := flattenSystemSdwanMembersPriorityInSla(i["priority-in-sla"], d, pre_append)
			tmp["priority_in_sla"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-PriorityInSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := i["priority-out-sla"]; ok {
			v := flattenSystemSdwanMembersPriorityOutSla(i["priority-out-sla"], d, pre_append)
			tmp["priority_out_sla"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-PriorityOutSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := i["priority6"]; ok {
			v := flattenSystemSdwanMembersPriority6(i["priority6"], d, pre_append)
			tmp["priority6"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Priority6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {
			v := flattenSystemSdwanMembersSeqNum(i["seq-num"], d, pre_append)
			tmp["seq_num"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-SeqNum")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {
			v := flattenSystemSdwanMembersSource(i["source"], d, pre_append)
			tmp["source"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Source")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {
			v := flattenSystemSdwanMembersSource6(i["source6"], d, pre_append)
			tmp["source6"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Source6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {
			v := flattenSystemSdwanMembersSpilloverThreshold(i["spillover-threshold"], d, pre_append)
			tmp["spillover_threshold"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-SpilloverThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemSdwanMembersStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport_group"
		if _, ok := i["transport-group"]; ok {
			v := flattenSystemSdwanMembersTransportGroup(i["transport-group"], d, pre_append)
			tmp["transport_group"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-TransportGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {
			v := flattenSystemSdwanMembersVolumeRatio(i["volume-ratio"], d, pre_append)
			tmp["volume_ratio"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-VolumeRatio")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			v := flattenSystemSdwanMembersWeight(i["weight"], d, pre_append)
			tmp["weight"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Weight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := i["zone"]; ok {
			v := flattenSystemSdwanMembersZone(i["zone"], d, pre_append)
			tmp["zone"] = fortiAPISubPartPatch(v, "SystemSdwan-Members-Zone")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanMembersComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanMembersPreferredSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriorityInSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriorityOutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriority6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSource6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersTransportGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanMembersZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenSystemSdwanNeighborHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenSystemSdwanNeighborIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {
			v := flattenSystemSdwanNeighborMember(i["member"], d, pre_append)
			tmp["member"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-Member")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenSystemSdwanNeighborMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenSystemSdwanNeighborMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenSystemSdwanNeighborRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_metric"
		if _, ok := i["route-metric"]; ok {
			v := flattenSystemSdwanNeighborRouteMetric(i["route-metric"], d, pre_append)
			tmp["route_metric"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-RouteMetric")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {
			v := flattenSystemSdwanNeighborServiceId(i["service-id"], d, pre_append)
			tmp["service_id"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-ServiceId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := i["sla-id"]; ok {
			v := flattenSystemSdwanNeighborSlaId(i["sla-id"], d, pre_append)
			tmp["sla_id"] = fortiAPISubPartPatch(v, "SystemSdwan-Neighbor-SlaId")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanNeighborMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanNeighborMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborRouteMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanNeighborSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHoldBootTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHoldDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {
			v := flattenSystemSdwanServiceAddrMode(i["addr-mode"], d, pre_append)
			tmp["addr_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-AddrMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_exclusive"
		if _, ok := i["agent-exclusive"]; ok {
			v := flattenSystemSdwanServiceAgentExclusive(i["agent-exclusive"], d, pre_append)
			tmp["agent_exclusive"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-AgentExclusive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := i["bandwidth-weight"]; ok {
			v := flattenSystemSdwanServiceBandwidthWeight(i["bandwidth-weight"], d, pre_append)
			tmp["bandwidth_weight"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-BandwidthWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			v := flattenSystemSdwanServiceComment(i["comment"], d, pre_append)
			tmp["comment"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Comment")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {
			v := flattenSystemSdwanServiceDefault(i["default"], d, pre_append)
			tmp["default"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Default")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := i["dscp-forward"]; ok {
			v := flattenSystemSdwanServiceDscpForward(i["dscp-forward"], d, pre_append)
			tmp["dscp_forward"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-DscpForward")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := i["dscp-forward-tag"]; ok {
			v := flattenSystemSdwanServiceDscpForwardTag(i["dscp-forward-tag"], d, pre_append)
			tmp["dscp_forward_tag"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-DscpForwardTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := i["dscp-reverse"]; ok {
			v := flattenSystemSdwanServiceDscpReverse(i["dscp-reverse"], d, pre_append)
			tmp["dscp_reverse"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-DscpReverse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := i["dscp-reverse-tag"]; ok {
			v := flattenSystemSdwanServiceDscpReverseTag(i["dscp-reverse-tag"], d, pre_append)
			tmp["dscp_reverse_tag"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-DscpReverseTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			v := flattenSystemSdwanServiceDst(i["dst"], d, pre_append)
			tmp["dst"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Dst")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := i["dst-negate"]; ok {
			v := flattenSystemSdwanServiceDstNegate(i["dst-negate"], d, pre_append)
			tmp["dst_negate"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-DstNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {
			v := flattenSystemSdwanServiceDst6(i["dst6"], d, pre_append)
			tmp["dst6"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Dst6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {
			v := flattenSystemSdwanServiceEndPort(i["end-port"], d, pre_append)
			tmp["end_port"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-EndPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_src_port"
		if _, ok := i["end-src-port"]; ok {
			v := flattenSystemSdwanServiceEndSrcPort(i["end-src-port"], d, pre_append)
			tmp["end_src_port"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-EndSrcPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			v := flattenSystemSdwanServiceGateway(i["gateway"], d, pre_append)
			tmp["gateway"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Gateway")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			v := flattenSystemSdwanServiceGroups(i["groups"], d, pre_append)
			tmp["groups"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Groups")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := i["hash-mode"]; ok {
			v := flattenSystemSdwanServiceHashMode(i["hash-mode"], d, pre_append)
			tmp["hash_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-HashMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenSystemSdwanServiceHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := i["hold-down-time"]; ok {
			v := flattenSystemSdwanServiceHoldDownTime(i["hold-down-time"], d, pre_append)
			tmp["hold_down_time"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-HoldDownTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdwanServiceId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := i["input-device"]; ok {
			v := flattenSystemSdwanServiceInputDevice(i["input-device"], d, pre_append)
			tmp["input_device"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InputDevice")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := i["input-device-negate"]; ok {
			v := flattenSystemSdwanServiceInputDeviceNegate(i["input-device-negate"], d, pre_append)
			tmp["input_device_negate"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InputDeviceNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_zone"
		if _, ok := i["input-zone"]; ok {
			v := flattenSystemSdwanServiceInputZone(i["input-zone"], d, pre_append)
			tmp["input_zone"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InputZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := i["internet-service"]; ok {
			v := flattenSystemSdwanServiceInternetService(i["internet-service"], d, pre_append)
			tmp["internet_service"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetService")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := i["internet-service-app-ctrl"]; ok {
			v := flattenSystemSdwanServiceInternetServiceAppCtrl(i["internet-service-app-ctrl"], d, pre_append)
			tmp["internet_service_app_ctrl"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceAppCtrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_category"
		if _, ok := i["internet-service-app-ctrl-category"]; ok {
			v := flattenSystemSdwanServiceInternetServiceAppCtrlCategory(i["internet-service-app-ctrl-category"], d, pre_append)
			tmp["internet_service_app_ctrl_category"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceAppCtrlCategory")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := i["internet-service-app-ctrl-group"]; ok {
			v := flattenSystemSdwanServiceInternetServiceAppCtrlGroup(i["internet-service-app-ctrl-group"], d, pre_append)
			tmp["internet_service_app_ctrl_group"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceAppCtrlGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := i["internet-service-custom"]; ok {
			v := flattenSystemSdwanServiceInternetServiceCustom(i["internet-service-custom"], d, pre_append)
			tmp["internet_service_custom"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceCustom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := i["internet-service-custom-group"]; ok {
			v := flattenSystemSdwanServiceInternetServiceCustomGroup(i["internet-service-custom-group"], d, pre_append)
			tmp["internet_service_custom_group"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceCustomGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := i["internet-service-group"]; ok {
			v := flattenSystemSdwanServiceInternetServiceGroup(i["internet-service-group"], d, pre_append)
			tmp["internet_service_group"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceGroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := i["internet-service-name"]; ok {
			v := flattenSystemSdwanServiceInternetServiceName(i["internet-service-name"], d, pre_append)
			tmp["internet_service_name"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-InternetServiceName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := i["jitter-weight"]; ok {
			v := flattenSystemSdwanServiceJitterWeight(i["jitter-weight"], d, pre_append)
			tmp["jitter_weight"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-JitterWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := i["latency-weight"]; ok {
			v := flattenSystemSdwanServiceLatencyWeight(i["latency-weight"], d, pre_append)
			tmp["latency_weight"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-LatencyWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {
			v := flattenSystemSdwanServiceLinkCostFactor(i["link-cost-factor"], d, pre_append)
			tmp["link_cost_factor"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-LinkCostFactor")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := i["link-cost-threshold"]; ok {
			v := flattenSystemSdwanServiceLinkCostThreshold(i["link-cost-threshold"], d, pre_append)
			tmp["link_cost_threshold"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-LinkCostThreshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balance"
		if _, ok := i["load-balance"]; ok {
			v := flattenSystemSdwanServiceLoadBalance(i["load-balance"], d, pre_append)
			tmp["load_balance"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-LoadBalance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenSystemSdwanServiceMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {
			v := flattenSystemSdwanServiceMode(i["mode"], d, pre_append)
			tmp["mode"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Mode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdwanServiceName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := i["packet-loss-weight"]; ok {
			v := flattenSystemSdwanServicePacketLossWeight(i["packet-loss-weight"], d, pre_append)
			tmp["packet_loss_weight"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-PacketLossWeight")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := i["passive-measurement"]; ok {
			v := flattenSystemSdwanServicePassiveMeasurement(i["passive-measurement"], d, pre_append)
			tmp["passive_measurement"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-PassiveMeasurement")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := i["priority-members"]; ok {
			v := flattenSystemSdwanServicePriorityMembers(i["priority-members"], d, pre_append)
			tmp["priority_members"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-PriorityMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := i["priority-zone"]; ok {
			v := flattenSystemSdwanServicePriorityZone(i["priority-zone"], d, pre_append)
			tmp["priority_zone"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-PriorityZone")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenSystemSdwanServiceProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Protocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := i["quality-link"]; ok {
			v := flattenSystemSdwanServiceQualityLink(i["quality-link"], d, pre_append)
			tmp["quality_link"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-QualityLink")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			v := flattenSystemSdwanServiceRole(i["role"], d, pre_append)
			tmp["role"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Role")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {
			v := flattenSystemSdwanServiceRouteTag(i["route-tag"], d, pre_append)
			tmp["route_tag"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-RouteTag")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			v := flattenSystemSdwanServiceShortcut(i["shortcut"], d, pre_append)
			tmp["shortcut"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Shortcut")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_priority"
		if _, ok := i["shortcut-priority"]; ok {
			v := flattenSystemSdwanServiceShortcutPriority(i["shortcut-priority"], d, pre_append)
			tmp["shortcut_priority"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-ShortcutPriority")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {
			v := flattenSystemSdwanServiceSla(i["sla"], d, pre_append)
			tmp["sla"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Sla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := i["sla-compare-method"]; ok {
			v := flattenSystemSdwanServiceSlaCompareMethod(i["sla-compare-method"], d, pre_append)
			tmp["sla_compare_method"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-SlaCompareMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_stickiness"
		if _, ok := i["sla-stickiness"]; ok {
			v := flattenSystemSdwanServiceSlaStickiness(i["sla-stickiness"], d, pre_append)
			tmp["sla_stickiness"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-SlaStickiness")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			v := flattenSystemSdwanServiceSrc(i["src"], d, pre_append)
			tmp["src"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Src")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := i["src-negate"]; ok {
			v := flattenSystemSdwanServiceSrcNegate(i["src-negate"], d, pre_append)
			tmp["src_negate"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-SrcNegate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {
			v := flattenSystemSdwanServiceSrc6(i["src6"], d, pre_append)
			tmp["src6"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Src6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := i["standalone-action"]; ok {
			v := flattenSystemSdwanServiceStandaloneAction(i["standalone-action"], d, pre_append)
			tmp["standalone_action"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-StandaloneAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {
			v := flattenSystemSdwanServiceStartPort(i["start-port"], d, pre_append)
			tmp["start_port"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-StartPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_src_port"
		if _, ok := i["start-src-port"]; ok {
			v := flattenSystemSdwanServiceStartSrcPort(i["start-src-port"], d, pre_append)
			tmp["start_src_port"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-StartSrcPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenSystemSdwanServiceStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := i["tie-break"]; ok {
			v := flattenSystemSdwanServiceTieBreak(i["tie-break"], d, pre_append)
			tmp["tie_break"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-TieBreak")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {
			v := flattenSystemSdwanServiceTos(i["tos"], d, pre_append)
			tmp["tos"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Tos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {
			v := flattenSystemSdwanServiceTosMask(i["tos-mask"], d, pre_append)
			tmp["tos_mask"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-TosMask")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := i["use-shortcut-sla"]; ok {
			v := flattenSystemSdwanServiceUseShortcutSla(i["use-shortcut-sla"], d, pre_append)
			tmp["use_shortcut_sla"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-UseShortcutSla")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			v := flattenSystemSdwanServiceUsers(i["users"], d, pre_append)
			tmp["users"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-Users")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone_mode"
		if _, ok := i["zone-mode"]; ok {
			v := flattenSystemSdwanServiceZoneMode(i["zone-mode"], d, pre_append)
			tmp["zone_mode"] = fortiAPISubPartPatch(v, "SystemSdwan-Service-ZoneMode")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceAgentExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceDstNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceDst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceEndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceEndSrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceHashMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceInputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceInputZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemSdwanServiceInternetServiceAppCtrlCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenSystemSdwanServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceLoadBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServicePassiveMeasurement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServicePriorityZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceQualityLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceRouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceShortcutPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSla(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {
			v := flattenSystemSdwanServiceSlaHealthCheck(i["health-check"], d, pre_append)
			tmp["health_check"] = fortiAPISubPartPatch(v, "SystemSdwanService-Sla-HealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemSdwanServiceSlaId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemSdwanService-Sla-Id")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSlaStickiness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrc6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceStartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceStartSrcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceTieBreak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceUseShortcutSla(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanServiceUsers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanServiceZoneMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanSpeedtestBypassRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_health_check"
		if _, ok := i["advpn-health-check"]; ok {
			v := flattenSystemSdwanZoneAdvpnHealthCheck(i["advpn-health-check"], d, pre_append)
			tmp["advpn_health_check"] = fortiAPISubPartPatch(v, "SystemSdwan-Zone-AdvpnHealthCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_select"
		if _, ok := i["advpn-select"]; ok {
			v := flattenSystemSdwanZoneAdvpnSelect(i["advpn-select"], d, pre_append)
			tmp["advpn_select"] = fortiAPISubPartPatch(v, "SystemSdwan-Zone-AdvpnSelect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {
			v := flattenSystemSdwanZoneMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append)
			tmp["minimum_sla_meet_members"] = fortiAPISubPartPatch(v, "SystemSdwan-Zone-MinimumSlaMeetMembers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSdwanZoneName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSdwan-Zone-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := i["service-sla-tie-break"]; ok {
			v := flattenSystemSdwanZoneServiceSlaTieBreak(i["service-sla-tie-break"], d, pre_append)
			tmp["service_sla_tie_break"] = fortiAPISubPartPatch(v, "SystemSdwan-Zone-ServiceSlaTieBreak")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSdwanZoneAdvpnHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSdwanZoneAdvpnSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZoneMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZoneName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSdwanZoneServiceSlaTieBreak(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSdwan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("app_perf_log_period", flattenSystemSdwanAppPerfLogPeriod(o["app-perf-log-period"], d, "app_perf_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["app-perf-log-period"], "SystemSdwan-AppPerfLogPeriod"); ok {
			if err = d.Set("app_perf_log_period", vv); err != nil {
				return fmt.Errorf("Error reading app_perf_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_perf_log_period: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("duplication", flattenSystemSdwanDuplication(o["duplication"], d, "duplication")); err != nil {
			if vv, ok := fortiAPIPatch(o["duplication"], "SystemSdwan-Duplication"); ok {
				if err = d.Set("duplication", vv); err != nil {
					return fmt.Errorf("Error reading duplication: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading duplication: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("duplication"); ok {
			if err = d.Set("duplication", flattenSystemSdwanDuplication(o["duplication"], d, "duplication")); err != nil {
				if vv, ok := fortiAPIPatch(o["duplication"], "SystemSdwan-Duplication"); ok {
					if err = d.Set("duplication", vv); err != nil {
						return fmt.Errorf("Error reading duplication: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading duplication: %v", err)
				}
			}
		}
	}

	if err = d.Set("duplication_max_discrepancy", flattenSystemSdwanDuplicationMaxDiscrepancy(o["duplication-max-discrepancy"], d, "duplication_max_discrepancy")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-max-discrepancy"], "SystemSdwan-DuplicationMaxDiscrepancy"); ok {
			if err = d.Set("duplication_max_discrepancy", vv); err != nil {
				return fmt.Errorf("Error reading duplication_max_discrepancy: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_max_discrepancy: %v", err)
		}
	}

	if err = d.Set("duplication_max_num", flattenSystemSdwanDuplicationMaxNum(o["duplication-max-num"], d, "duplication_max_num")); err != nil {
		if vv, ok := fortiAPIPatch(o["duplication-max-num"], "SystemSdwan-DuplicationMaxNum"); ok {
			if err = d.Set("duplication_max_num", vv); err != nil {
				return fmt.Errorf("Error reading duplication_max_num: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading duplication_max_num: %v", err)
		}
	}

	if err = d.Set("fail_alert_interfaces", flattenSystemSdwanFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-alert-interfaces"], "SystemSdwan-FailAlertInterfaces"); ok {
			if err = d.Set("fail_alert_interfaces", vv); err != nil {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenSystemSdwanFailDetect(o["fail-detect"], d, "fail_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["fail-detect"], "SystemSdwan-FailDetect"); ok {
			if err = d.Set("fail_detect", vv); err != nil {
				return fmt.Errorf("Error reading fail_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenSystemSdwanHealthCheck(o["health-check"], d, "health_check")); err != nil {
			if vv, ok := fortiAPIPatch(o["health-check"], "SystemSdwan-HealthCheck"); ok {
				if err = d.Set("health_check", vv); err != nil {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check"); ok {
			if err = d.Set("health_check", flattenSystemSdwanHealthCheck(o["health-check"], d, "health_check")); err != nil {
				if vv, ok := fortiAPIPatch(o["health-check"], "SystemSdwan-HealthCheck"); ok {
					if err = d.Set("health_check", vv); err != nil {
						return fmt.Errorf("Error reading health_check: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("health_check_fortiguard", flattenSystemSdwanHealthCheckFortiguard(o["health-check-fortiguard"], d, "health_check_fortiguard")); err != nil {
			if vv, ok := fortiAPIPatch(o["health-check-fortiguard"], "SystemSdwan-HealthCheckFortiguard"); ok {
				if err = d.Set("health_check_fortiguard", vv); err != nil {
					return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check_fortiguard"); ok {
			if err = d.Set("health_check_fortiguard", flattenSystemSdwanHealthCheckFortiguard(o["health-check-fortiguard"], d, "health_check_fortiguard")); err != nil {
				if vv, ok := fortiAPIPatch(o["health-check-fortiguard"], "SystemSdwan-HealthCheckFortiguard"); ok {
					if err = d.Set("health_check_fortiguard", vv); err != nil {
						return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading health_check_fortiguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("load_balance_mode", flattenSystemSdwanLoadBalanceMode(o["load-balance-mode"], d, "load_balance_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["load-balance-mode"], "SystemSdwan-LoadBalanceMode"); ok {
			if err = d.Set("load_balance_mode", vv); err != nil {
				return fmt.Errorf("Error reading load_balance_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenSystemSdwanMembers(o["members"], d, "members")); err != nil {
			if vv, ok := fortiAPIPatch(o["members"], "SystemSdwan-Members"); ok {
				if err = d.Set("members", vv); err != nil {
					return fmt.Errorf("Error reading members: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenSystemSdwanMembers(o["members"], d, "members")); err != nil {
				if vv, ok := fortiAPIPatch(o["members"], "SystemSdwan-Members"); ok {
					if err = d.Set("members", vv); err != nil {
						return fmt.Errorf("Error reading members: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenSystemSdwanNeighbor(o["neighbor"], d, "neighbor")); err != nil {
			if vv, ok := fortiAPIPatch(o["neighbor"], "SystemSdwan-Neighbor"); ok {
				if err = d.Set("neighbor", vv); err != nil {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenSystemSdwanNeighbor(o["neighbor"], d, "neighbor")); err != nil {
				if vv, ok := fortiAPIPatch(o["neighbor"], "SystemSdwan-Neighbor"); ok {
					if err = d.Set("neighbor", vv); err != nil {
						return fmt.Errorf("Error reading neighbor: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if err = d.Set("neighbor_hold_boot_time", flattenSystemSdwanNeighborHoldBootTime(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-boot-time"], "SystemSdwan-NeighborHoldBootTime"); ok {
			if err = d.Set("neighbor_hold_boot_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenSystemSdwanNeighborHoldDown(o["neighbor-hold-down"], d, "neighbor_hold_down")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down"], "SystemSdwan-NeighborHoldDown"); ok {
			if err = d.Set("neighbor_hold_down", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenSystemSdwanNeighborHoldDownTime(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["neighbor-hold-down-time"], "SystemSdwan-NeighborHoldDownTime"); ok {
			if err = d.Set("neighbor_hold_down_time", vv); err != nil {
				return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if err = d.Set("option", flattenSystemSdwanOption(o["option"], d, "option")); err != nil {
		if vv, ok := fortiAPIPatch(o["option"], "SystemSdwan-Option"); ok {
			if err = d.Set("option", vv); err != nil {
				return fmt.Errorf("Error reading option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading option: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenSystemSdwanService(o["service"], d, "service")); err != nil {
			if vv, ok := fortiAPIPatch(o["service"], "SystemSdwan-Service"); ok {
				if err = d.Set("service", vv); err != nil {
					return fmt.Errorf("Error reading service: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenSystemSdwanService(o["service"], d, "service")); err != nil {
				if vv, ok := fortiAPIPatch(o["service"], "SystemSdwan-Service"); ok {
					if err = d.Set("service", vv); err != nil {
						return fmt.Errorf("Error reading service: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("speedtest_bypass_routing", flattenSystemSdwanSpeedtestBypassRouting(o["speedtest-bypass-routing"], d, "speedtest_bypass_routing")); err != nil {
		if vv, ok := fortiAPIPatch(o["speedtest-bypass-routing"], "SystemSdwan-SpeedtestBypassRouting"); ok {
			if err = d.Set("speedtest_bypass_routing", vv); err != nil {
				return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdwanStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSdwan-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("zone", flattenSystemSdwanZone(o["zone"], d, "zone")); err != nil {
			if vv, ok := fortiAPIPatch(o["zone"], "SystemSdwan-Zone"); ok {
				if err = d.Set("zone", vv); err != nil {
					return fmt.Errorf("Error reading zone: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("zone"); ok {
			if err = d.Set("zone", flattenSystemSdwanZone(o["zone"], d, "zone")); err != nil {
				if vv, ok := fortiAPIPatch(o["zone"], "SystemSdwan-Zone"); ok {
					if err = d.Set("zone", vv); err != nil {
						return fmt.Errorf("Error reading zone: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading zone: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSdwanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSdwanAppPerfLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandSystemSdwanDuplicationDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr6"], _ = expandSystemSdwanDuplicationDstaddr6(d, i["dstaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstintf"], _ = expandSystemSdwanDuplicationDstintf(d, i["dstintf"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdwanDuplicationId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-de-duplication"], _ = expandSystemSdwanDuplicationPacketDeDuplication(d, i["packet_de_duplication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-duplication"], _ = expandSystemSdwanDuplicationPacketDuplication(d, i["packet_duplication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service"], _ = expandSystemSdwanDuplicationService(d, i["service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandSystemSdwanDuplicationServiceId(d, i["service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_match_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-match-service"], _ = expandSystemSdwanDuplicationSlaMatchService(d, i["sla_match_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr"], _ = expandSystemSdwanDuplicationSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcaddr6"], _ = expandSystemSdwanDuplicationSrcaddr6(d, i["srcaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["srcintf"], _ = expandSystemSdwanDuplicationSrcintf(d, i["srcintf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationPacketDeDuplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationPacketDuplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationSlaMatchService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanDuplicationMaxDiscrepancy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationMaxNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanFailDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandSystemSdwanHealthCheckAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandSystemSdwanHealthCheckClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detect-mode"], _ = expandSystemSdwanHealthCheckDetectMode(d, i["detect_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandSystemSdwanHealthCheckDiffservcode(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-match-ip"], _ = expandSystemSdwanHealthCheckDnsMatchIp(d, i["dns_match_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-request-domain"], _ = expandSystemSdwanHealthCheckDnsRequestDomain(d, i["dns_request_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embed-measured-health"], _ = expandSystemSdwanHealthCheckEmbedMeasuredHealth(d, i["embed_measured_health"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandSystemSdwanHealthCheckFailtime(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiguard"], _ = expandSystemSdwanHealthCheckFortiguard(d, i["fortiguard"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiguard-name"], _ = expandSystemSdwanHealthCheckFortiguardName(d, i["fortiguard_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-file"], _ = expandSystemSdwanHealthCheckFtpFile(d, i["ftp_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-mode"], _ = expandSystemSdwanHealthCheckFtpMode(d, i["ftp_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandSystemSdwanHealthCheckHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandSystemSdwanHealthCheckHttpAgent(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandSystemSdwanHealthCheckHttpGet(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandSystemSdwanHealthCheckHttpMatch(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandSystemSdwanHealthCheckInterval(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandSystemSdwanHealthCheckMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-codec"], _ = expandSystemSdwanHealthCheckMosCodec(d, i["mos_codec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdwanHealthCheckName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandSystemSdwanHealthCheckPacketSize(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandSystemSdwanHealthCheckPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSystemSdwanHealthCheckPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-count"], _ = expandSystemSdwanHealthCheckProbeCount(d, i["probe_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandSystemSdwanHealthCheckProbePackets(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandSystemSdwanHealthCheckProbeTimeout(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandSystemSdwanHealthCheckProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-measured-method"], _ = expandSystemSdwanHealthCheckQualityMeasuredMethod(d, i["quality_measured_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandSystemSdwanHealthCheckRecoverytime(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandSystemSdwanHealthCheckSecurityMode(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSystemSdwanHealthCheckServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemSdwanHealthCheckSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandSystemSdwanHealthCheckSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id-redistribute"], _ = expandSystemSdwanHealthCheckSlaIdRedistribute(d, i["sla_id_redistribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandSystemSdwanHealthCheckSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandSystemSdwanHealthCheckSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandSystemSdwanHealthCheckSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["system-dns"], _ = expandSystemSdwanHealthCheckSystemDns(d, i["system_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandSystemSdwanHealthCheckThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandSystemSdwanHealthCheckThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandSystemSdwanHealthCheckThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandSystemSdwanHealthCheckThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandSystemSdwanHealthCheckThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandSystemSdwanHealthCheckThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandSystemSdwanHealthCheckUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandSystemSdwanHealthCheckUpdateStaticRoute(d, i["update_static_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandSystemSdwanHealthCheckUser(d, i["user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandSystemSdwanHealthCheckVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckDetectMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDnsMatchIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckEmbedMeasuredHealth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFtpFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckMosCodec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbeCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbePackets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckQualityMeasuredMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdwanHealthCheckSlaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandSystemSdwanHealthCheckSlaJitterThreshold(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandSystemSdwanHealthCheckSlaLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandSystemSdwanHealthCheckSlaLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandSystemSdwanHealthCheckSlaMosThreshold(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandSystemSdwanHealthCheckSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandSystemSdwanHealthCheckSlaPriorityInSla(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandSystemSdwanHealthCheckSlaPriorityOutSla(d, i["priority_out_sla"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckSlaMosThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPriorityInSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPriorityOutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaIdRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSystemDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandSystemSdwanHealthCheckFortiguardAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class-id"], _ = expandSystemSdwanHealthCheckFortiguardClassId(d, i["class_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["detect-mode"], _ = expandSystemSdwanHealthCheckFortiguardDetectMode(d, i["detect_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["diffservcode"], _ = expandSystemSdwanHealthCheckFortiguardDiffservcode(d, i["diffservcode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-match-ip"], _ = expandSystemSdwanHealthCheckFortiguardDnsMatchIp(d, i["dns_match_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dns-request-domain"], _ = expandSystemSdwanHealthCheckFortiguardDnsRequestDomain(d, i["dns_request_domain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "embed_measured_health"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["embed-measured-health"], _ = expandSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(d, i["embed_measured_health"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["failtime"], _ = expandSystemSdwanHealthCheckFortiguardFailtime(d, i["failtime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-file"], _ = expandSystemSdwanHealthCheckFortiguardFtpFile(d, i["ftp_file"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ftp-mode"], _ = expandSystemSdwanHealthCheckFortiguardFtpMode(d, i["ftp_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ha-priority"], _ = expandSystemSdwanHealthCheckFortiguardHaPriority(d, i["ha_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-agent"], _ = expandSystemSdwanHealthCheckFortiguardHttpAgent(d, i["http_agent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-get"], _ = expandSystemSdwanHealthCheckFortiguardHttpGet(d, i["http_get"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["http-match"], _ = expandSystemSdwanHealthCheckFortiguardHttpMatch(d, i["http_match"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interval"], _ = expandSystemSdwanHealthCheckFortiguardInterval(d, i["interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandSystemSdwanHealthCheckFortiguardMembers(d, i["members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_codec"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-codec"], _ = expandSystemSdwanHealthCheckFortiguardMosCodec(d, i["mos_codec"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-size"], _ = expandSystemSdwanHealthCheckFortiguardPacketSize(d, i["packet_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password"], _ = expandSystemSdwanHealthCheckFortiguardPassword(d, i["password"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandSystemSdwanHealthCheckFortiguardPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-count"], _ = expandSystemSdwanHealthCheckFortiguardProbeCount(d, i["probe_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-packets"], _ = expandSystemSdwanHealthCheckFortiguardProbePackets(d, i["probe_packets"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["probe-timeout"], _ = expandSystemSdwanHealthCheckFortiguardProbeTimeout(d, i["probe_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandSystemSdwanHealthCheckFortiguardProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-measured-method"], _ = expandSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(d, i["quality_measured_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["recoverytime"], _ = expandSystemSdwanHealthCheckFortiguardRecoverytime(d, i["recoverytime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["security-mode"], _ = expandSystemSdwanHealthCheckFortiguardSecurityMode(d, i["security_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandSystemSdwanHealthCheckFortiguardServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemSdwanHealthCheckFortiguardSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-fail-log-period"], _ = expandSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id_redistribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id-redistribute"], _ = expandSystemSdwanHealthCheckFortiguardSlaIdRedistribute(d, i["sla_id_redistribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-pass-log-period"], _ = expandSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandSystemSdwanHealthCheckFortiguardSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandSystemSdwanHealthCheckFortiguardSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["system-dns"], _ = expandSystemSdwanHealthCheckFortiguardSystemDns(d, i["system_dns"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["target-name"], _ = expandSystemSdwanHealthCheckFortiguardTargetName(d, i["target_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"], _ = expandSystemSdwanHealthCheckFortiguardThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-latency"], _ = expandSystemSdwanHealthCheckFortiguardThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"], _ = expandSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"], _ = expandSystemSdwanHealthCheckFortiguardThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-latency"], _ = expandSystemSdwanHealthCheckFortiguardThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"], _ = expandSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-cascade-interface"], _ = expandSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["update-static-route"], _ = expandSystemSdwanHealthCheckFortiguardUpdateStaticRoute(d, i["update_static_route"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user"], _ = expandSystemSdwanHealthCheckFortiguardUser(d, i["user"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf"], _ = expandSystemSdwanHealthCheckFortiguardVrf(d, i["vrf"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckFortiguardAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardDetectMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardDnsMatchIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardEmbedMeasuredHealth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardFtpFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardFtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardMosCodec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProbeCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProbePackets(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProbeTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardQualityMeasuredMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdwanHealthCheckFortiguardSlaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaJitterThreshold(d, i["jitter_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(d, i["latency_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mos_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mos-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaMosThreshold(d, i["mos_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packetloss-threshold"], _ = expandSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandSystemSdwanHealthCheckFortiguardSlaPriorityInSla(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(d, i["priority_out_sla"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanHealthCheckFortiguardSlaMosThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPriorityInSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPriorityOutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaIdRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardSystemDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardTargetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFortiguardVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanLoadBalanceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandSystemSdwanMembersComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cost"], _ = expandSystemSdwanMembersCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandSystemSdwanMembersGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway6"], _ = expandSystemSdwanMembersGateway6(d, i["gateway6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ingress-spillover-threshold"], _ = expandSystemSdwanMembersIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandSystemSdwanMembersInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["preferred-source"], _ = expandSystemSdwanMembersPreferredSource(d, i["preferred_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority"], _ = expandSystemSdwanMembersPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_in_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-in-sla"], _ = expandSystemSdwanMembersPriorityInSla(d, i["priority_in_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_out_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-out-sla"], _ = expandSystemSdwanMembersPriorityOutSla(d, i["priority_out_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority6"], _ = expandSystemSdwanMembersPriority6(d, i["priority6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["seq-num"], _ = expandSystemSdwanMembersSeqNum(d, i["seq_num"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source"], _ = expandSystemSdwanMembersSource(d, i["source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source6"], _ = expandSystemSdwanMembersSource6(d, i["source6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["spillover-threshold"], _ = expandSystemSdwanMembersSpilloverThreshold(d, i["spillover_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemSdwanMembersStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transport-group"], _ = expandSystemSdwanMembersTransportGroup(d, i["transport_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["volume-ratio"], _ = expandSystemSdwanMembersVolumeRatio(d, i["volume_ratio"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["weight"], _ = expandSystemSdwanMembersWeight(d, i["weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["zone"], _ = expandSystemSdwanMembersZone(d, i["zone"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanMembersComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanMembersPreferredSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriorityInSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriorityOutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriority6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSource6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersTransportGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandSystemSdwanNeighborHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandSystemSdwanNeighborIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["member"], _ = expandSystemSdwanNeighborMember(d, i["member"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandSystemSdwanNeighborMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandSystemSdwanNeighborMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandSystemSdwanNeighborRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_metric"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-metric"], _ = expandSystemSdwanNeighborRouteMetric(d, i["route_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-id"], _ = expandSystemSdwanNeighborServiceId(d, i["service_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-id"], _ = expandSystemSdwanNeighborSlaId(d, i["sla_id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanNeighborMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanNeighborMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborRouteMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanNeighborSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHoldBootTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHoldDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["addr-mode"], _ = expandSystemSdwanServiceAddrMode(d, i["addr_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "agent_exclusive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["agent-exclusive"], _ = expandSystemSdwanServiceAgentExclusive(d, i["agent_exclusive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["bandwidth-weight"], _ = expandSystemSdwanServiceBandwidthWeight(d, i["bandwidth_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["comment"], _ = expandSystemSdwanServiceComment(d, i["comment"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["default"], _ = expandSystemSdwanServiceDefault(d, i["default"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward"], _ = expandSystemSdwanServiceDscpForward(d, i["dscp_forward"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-forward-tag"], _ = expandSystemSdwanServiceDscpForwardTag(d, i["dscp_forward_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse"], _ = expandSystemSdwanServiceDscpReverse(d, i["dscp_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dscp-reverse-tag"], _ = expandSystemSdwanServiceDscpReverseTag(d, i["dscp_reverse_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst"], _ = expandSystemSdwanServiceDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst-negate"], _ = expandSystemSdwanServiceDstNegate(d, i["dst_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dst6"], _ = expandSystemSdwanServiceDst6(d, i["dst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-port"], _ = expandSystemSdwanServiceEndPort(d, i["end_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_src_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["end-src-port"], _ = expandSystemSdwanServiceEndSrcPort(d, i["end_src_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["gateway"], _ = expandSystemSdwanServiceGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["groups"], _ = expandSystemSdwanServiceGroups(d, i["groups"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hash-mode"], _ = expandSystemSdwanServiceHashMode(d, i["hash_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandSystemSdwanServiceHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hold-down-time"], _ = expandSystemSdwanServiceHoldDownTime(d, i["hold_down_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdwanServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device"], _ = expandSystemSdwanServiceInputDevice(d, i["input_device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-device-negate"], _ = expandSystemSdwanServiceInputDeviceNegate(d, i["input_device_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["input-zone"], _ = expandSystemSdwanServiceInputZone(d, i["input_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service"], _ = expandSystemSdwanServiceInternetService(d, i["internet_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl"], _ = expandSystemSdwanServiceInternetServiceAppCtrl(d, i["internet_service_app_ctrl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_category"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-category"], _ = expandSystemSdwanServiceInternetServiceAppCtrlCategory(d, i["internet_service_app_ctrl_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-group"], _ = expandSystemSdwanServiceInternetServiceAppCtrlGroup(d, i["internet_service_app_ctrl_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom"], _ = expandSystemSdwanServiceInternetServiceCustom(d, i["internet_service_custom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-custom-group"], _ = expandSystemSdwanServiceInternetServiceCustomGroup(d, i["internet_service_custom_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-group"], _ = expandSystemSdwanServiceInternetServiceGroup(d, i["internet_service_group"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["internet-service-name"], _ = expandSystemSdwanServiceInternetServiceName(d, i["internet_service_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["jitter-weight"], _ = expandSystemSdwanServiceJitterWeight(d, i["jitter_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["latency-weight"], _ = expandSystemSdwanServiceLatencyWeight(d, i["latency_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-factor"], _ = expandSystemSdwanServiceLinkCostFactor(d, i["link_cost_factor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["link-cost-threshold"], _ = expandSystemSdwanServiceLinkCostThreshold(d, i["link_cost_threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["load-balance"], _ = expandSystemSdwanServiceLoadBalance(d, i["load_balance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandSystemSdwanServiceMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mode"], _ = expandSystemSdwanServiceMode(d, i["mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdwanServiceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["packet-loss-weight"], _ = expandSystemSdwanServicePacketLossWeight(d, i["packet_loss_weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["passive-measurement"], _ = expandSystemSdwanServicePassiveMeasurement(d, i["passive_measurement"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-members"], _ = expandSystemSdwanServicePriorityMembers(d, i["priority_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["priority-zone"], _ = expandSystemSdwanServicePriorityZone(d, i["priority_zone"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["protocol"], _ = expandSystemSdwanServiceProtocol(d, i["protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quality-link"], _ = expandSystemSdwanServiceQualityLink(d, i["quality_link"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["role"], _ = expandSystemSdwanServiceRole(d, i["role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["route-tag"], _ = expandSystemSdwanServiceRouteTag(d, i["route_tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut"], _ = expandSystemSdwanServiceShortcut(d, i["shortcut"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut_priority"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["shortcut-priority"], _ = expandSystemSdwanServiceShortcutPriority(d, i["shortcut_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemSdwanServiceSla(d, i["sla"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["sla"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-compare-method"], _ = expandSystemSdwanServiceSlaCompareMethod(d, i["sla_compare_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_stickiness"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sla-stickiness"], _ = expandSystemSdwanServiceSlaStickiness(d, i["sla_stickiness"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src"], _ = expandSystemSdwanServiceSrc(d, i["src"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-negate"], _ = expandSystemSdwanServiceSrcNegate(d, i["src_negate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src6"], _ = expandSystemSdwanServiceSrc6(d, i["src6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["standalone-action"], _ = expandSystemSdwanServiceStandaloneAction(d, i["standalone_action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-port"], _ = expandSystemSdwanServiceStartPort(d, i["start_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_src_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["start-src-port"], _ = expandSystemSdwanServiceStartSrcPort(d, i["start_src_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandSystemSdwanServiceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tie-break"], _ = expandSystemSdwanServiceTieBreak(d, i["tie_break"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos"], _ = expandSystemSdwanServiceTos(d, i["tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tos-mask"], _ = expandSystemSdwanServiceTosMask(d, i["tos_mask"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["use-shortcut-sla"], _ = expandSystemSdwanServiceUseShortcutSla(d, i["use_shortcut_sla"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["users"], _ = expandSystemSdwanServiceUsers(d, i["users"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["zone-mode"], _ = expandSystemSdwanServiceZoneMode(d, i["zone_mode"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceAgentExclusive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceBandwidthWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpForwardTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpReverseTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceDstNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceEndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceEndSrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceHashMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceHoldDownTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInputDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInputZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceJitterWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLatencyWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLinkCostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLoadBalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePacketLossWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePassiveMeasurement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServicePriorityZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceQualityLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceRouteTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceShortcut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceShortcutPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["health-check"], _ = expandSystemSdwanServiceSlaHealthCheck(d, i["health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemSdwanServiceSlaId(d, i["id"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceSlaHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSlaCompareMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSlaStickiness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrc6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceStandaloneAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStartSrcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTieBreak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceUseShortcutSla(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanServiceZoneMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanSpeedtestBypassRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_health_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advpn-health-check"], _ = expandSystemSdwanZoneAdvpnHealthCheck(d, i["advpn_health_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advpn_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["advpn-select"], _ = expandSystemSdwanZoneAdvpnSelect(d, i["advpn_select"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["minimum-sla-meet-members"], _ = expandSystemSdwanZoneMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSdwanZoneName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["service-sla-tie-break"], _ = expandSystemSdwanZoneServiceSlaTieBreak(d, i["service_sla_tie_break"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSdwanZoneAdvpnHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSdwanZoneAdvpnSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneServiceSlaTieBreak(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_perf_log_period"); ok || d.HasChange("app_perf_log_period") {
		t, err := expandSystemSdwanAppPerfLogPeriod(d, v, "app_perf_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-perf-log-period"] = t
		}
	}

	if v, ok := d.GetOk("duplication"); ok || d.HasChange("duplication") {
		t, err := expandSystemSdwanDuplication(d, v, "duplication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication"] = t
		}
	}

	if v, ok := d.GetOk("duplication_max_discrepancy"); ok || d.HasChange("duplication_max_discrepancy") {
		t, err := expandSystemSdwanDuplicationMaxDiscrepancy(d, v, "duplication_max_discrepancy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-max-discrepancy"] = t
		}
	}

	if v, ok := d.GetOk("duplication_max_num"); ok || d.HasChange("duplication_max_num") {
		t, err := expandSystemSdwanDuplicationMaxNum(d, v, "duplication_max_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-max-num"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok || d.HasChange("fail_alert_interfaces") {
		t, err := expandSystemSdwanFailAlertInterfaces(d, v, "fail_alert_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok || d.HasChange("fail_detect") {
		t, err := expandSystemSdwanFailDetect(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandSystemSdwanHealthCheck(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("health_check_fortiguard"); ok || d.HasChange("health_check_fortiguard") {
		t, err := expandSystemSdwanHealthCheckFortiguard(d, v, "health_check_fortiguard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_mode"); ok || d.HasChange("load_balance_mode") {
		t, err := expandSystemSdwanLoadBalanceMode(d, v, "load_balance_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-mode"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		t, err := expandSystemSdwanMembers(d, v, "members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		t, err := expandSystemSdwanNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_boot_time"); ok || d.HasChange("neighbor_hold_boot_time") {
		t, err := expandSystemSdwanNeighborHoldBootTime(d, v, "neighbor_hold_boot_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-boot-time"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok || d.HasChange("neighbor_hold_down") {
		t, err := expandSystemSdwanNeighborHoldDown(d, v, "neighbor_hold_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down_time"); ok || d.HasChange("neighbor_hold_down_time") {
		t, err := expandSystemSdwanNeighborHoldDownTime(d, v, "neighbor_hold_down_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("option"); ok || d.HasChange("option") {
		t, err := expandSystemSdwanOption(d, v, "option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandSystemSdwanService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("speedtest_bypass_routing"); ok || d.HasChange("speedtest_bypass_routing") {
		t, err := expandSystemSdwanSpeedtestBypassRouting(d, v, "speedtest_bypass_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtest-bypass-routing"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSdwanStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		t, err := expandSystemSdwanZone(d, v, "zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	return &obj, nil
}
