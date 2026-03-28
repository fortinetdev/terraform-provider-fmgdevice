// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure RADIUS server entries.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserRadius() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserRadiusCreate,
		Read:   resourceUserRadiusRead,
		Update: resourceUserRadiusUpdate,
		Delete: resourceUserRadiusDelete,

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
			"account_key_cert_field": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accounting_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
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
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrf_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"acct_all_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"all_usergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"call_station_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_scope": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vdom": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"account_key_cert_field": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"account_key_processing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"accounting_server": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
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
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"secret": &schema.Schema{
										Type:      schema.TypeSet,
										Elem:      &schema.Schema{Type: schema.TypeString},
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"server": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"source_ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vrf_select": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"acct_all_servers": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"acct_interim_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"all_usergroup": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ca_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"call_station_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"class": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"client_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_carrier_endpoint_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_carrier_endpoint_block_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_context_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dp_flush_ip_session": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_hold_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dp_http_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_http_header_fallback": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_http_header_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_http_header_suppress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_log_dyn_flags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"dp_log_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dp_mem_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dp_profile_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_profile_attribute_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_radius_response": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dp_radius_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dp_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"dp_validate_request_secret": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"endpoint_translation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_carrier_endpoint_convert_hex": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_carrier_endpoint_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_carrier_endpoint_header_suppress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_carrier_endpoint_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_carrier_endpoint_prefix_range_max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ep_carrier_endpoint_prefix_range_min": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ep_carrier_endpoint_prefix_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_carrier_endpoint_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_ip_header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_ip_header_suppress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_missing_header_fallback": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ep_profile_query_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_override_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"h3c_compatibility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						},
						"mac_case": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_password_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_username_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nas_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nas_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"nas_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_renewal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radius_coa": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"radius_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"require_message_authenticator": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso_context_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rsso_endpoint_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso_endpoint_block_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso_ep_one_ip_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso_flush_ip_session": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso_log_flags": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"rsso_log_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rsso_radius_response": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"rsso_radius_server_port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rsso_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"rsso_validate_request_secret": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secondary_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"secondary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_identity_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_ip_interface": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sso_attribute": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_attribute_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_attribute_value_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"switch_controller_acct_fast_framedip_detect": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"switch_controller_nas_ip_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"switch_controller_service_type": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"tertiary_secret": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"tertiary_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tls_min_proto_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"transport_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"use_group_for_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"use_management_vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"username_case_sensitive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vrf_select": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"group_override_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"h3c_compatibility": &schema.Schema{
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
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"nas_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nas_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_coa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"require_message_authenticator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_context_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso_endpoint_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_endpoint_block_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_ep_one_ip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_flush_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_log_flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"rsso_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso_radius_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso_radius_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rsso_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"rsso_validate_request_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sso_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_attribute_value_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_controller_acct_fast_framedip_detect": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_controller_nas_ip_dynamic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_service_type": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"tertiary_secret": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"tertiary_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tls_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transport_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_management_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserRadiusCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserRadius(d)
	if err != nil {
		return fmt.Errorf("Error creating UserRadius resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserRadius(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserRadius(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserRadius resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserRadius(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserRadius resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserRadiusRead(d, m)
}

func resourceUserRadiusUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserRadius(d)
	if err != nil {
		return fmt.Errorf("Error updating UserRadius resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserRadius(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserRadius resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserRadiusRead(d, m)
}

func resourceUserRadiusDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserRadius(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserRadius resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserRadiusRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserRadius(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserRadius resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserRadius(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserRadius resource from API: %v", err)
	}
	return nil
}

func flattenUserRadiusAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserRadiusAccountingServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenUserRadiusAccountingServerInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenUserRadiusAccountingServerInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserRadiusAccountingServerPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenUserRadiusAccountingServerServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenUserRadiusAccountingServerSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenUserRadiusAccountingServerStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenUserRadiusAccountingServerVrfSelect(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "UserRadius-AccountingServer-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserRadiusAccountingServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusAccountingServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAcctAllServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAllUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusCallStationIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := i["_scope"]; ok {
			v := flattenUserRadiusDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_cert_field"
		if _, ok := i["account-key-cert-field"]; ok {
			v := flattenUserRadiusDynamicMappingAccountKeyCertField(i["account-key-cert-field"], d, pre_append)
			tmp["account_key_cert_field"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AccountKeyCertField")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_processing"
		if _, ok := i["account-key-processing"]; ok {
			v := flattenUserRadiusDynamicMappingAccountKeyProcessing(i["account-key-processing"], d, pre_append)
			tmp["account_key_processing"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AccountKeyProcessing")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accounting_server"
		if _, ok := i["accounting-server"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServer(i["accounting-server"], d, pre_append)
			tmp["accounting_server"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AccountingServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_all_servers"
		if _, ok := i["acct-all-servers"]; ok {
			v := flattenUserRadiusDynamicMappingAcctAllServers(i["acct-all-servers"], d, pre_append)
			tmp["acct_all_servers"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AcctAllServers")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := i["acct-interim-interval"]; ok {
			v := flattenUserRadiusDynamicMappingAcctInterimInterval(i["acct-interim-interval"], d, pre_append)
			tmp["acct_interim_interval"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AcctInterimInterval")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "all_usergroup"
		if _, ok := i["all-usergroup"]; ok {
			v := flattenUserRadiusDynamicMappingAllUsergroup(i["all-usergroup"], d, pre_append)
			tmp["all_usergroup"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AllUsergroup")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := i["auth-type"]; ok {
			v := flattenUserRadiusDynamicMappingAuthType(i["auth-type"], d, pre_append)
			tmp["auth_type"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-AuthType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_cert"
		if _, ok := i["ca-cert"]; ok {
			v := flattenUserRadiusDynamicMappingCaCert(i["ca-cert"], d, pre_append)
			tmp["ca_cert"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-CaCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "call_station_id_type"
		if _, ok := i["call-station-id-type"]; ok {
			v := flattenUserRadiusDynamicMappingCallStationIdType(i["call-station-id-type"], d, pre_append)
			tmp["call_station_id_type"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-CallStationIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := i["class"]; ok {
			v := flattenUserRadiusDynamicMappingClass(i["class"], d, pre_append)
			tmp["class"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Class")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			v := flattenUserRadiusDynamicMappingClientCert(i["client-cert"], d, pre_append)
			tmp["client_cert"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-ClientCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delimiter"
		if _, ok := i["delimiter"]; ok {
			v := flattenUserRadiusDynamicMappingDelimiter(i["delimiter"], d, pre_append)
			tmp["delimiter"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Delimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_attribute"
		if _, ok := i["dp-carrier-endpoint-attribute"]; ok {
			v := flattenUserRadiusDynamicMappingDpCarrierEndpointAttribute(i["dp-carrier-endpoint-attribute"], d, pre_append)
			tmp["dp_carrier_endpoint_attribute"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpCarrierEndpointAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_block_attribute"
		if _, ok := i["dp-carrier-endpoint-block-attribute"]; ok {
			v := flattenUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(i["dp-carrier-endpoint-block-attribute"], d, pre_append)
			tmp["dp_carrier_endpoint_block_attribute"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpCarrierEndpointBlockAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_context_timeout"
		if _, ok := i["dp-context-timeout"]; ok {
			v := flattenUserRadiusDynamicMappingDpContextTimeout(i["dp-context-timeout"], d, pre_append)
			tmp["dp_context_timeout"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpContextTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_flush_ip_session"
		if _, ok := i["dp-flush-ip-session"]; ok {
			v := flattenUserRadiusDynamicMappingDpFlushIpSession(i["dp-flush-ip-session"], d, pre_append)
			tmp["dp_flush_ip_session"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpFlushIpSession")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_hold_time"
		if _, ok := i["dp-hold-time"]; ok {
			v := flattenUserRadiusDynamicMappingDpHoldTime(i["dp-hold-time"], d, pre_append)
			tmp["dp_hold_time"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpHoldTime")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header"
		if _, ok := i["dp-http-header"]; ok {
			v := flattenUserRadiusDynamicMappingDpHttpHeader(i["dp-http-header"], d, pre_append)
			tmp["dp_http_header"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpHttpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_fallback"
		if _, ok := i["dp-http-header-fallback"]; ok {
			v := flattenUserRadiusDynamicMappingDpHttpHeaderFallback(i["dp-http-header-fallback"], d, pre_append)
			tmp["dp_http_header_fallback"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpHttpHeaderFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_status"
		if _, ok := i["dp-http-header-status"]; ok {
			v := flattenUserRadiusDynamicMappingDpHttpHeaderStatus(i["dp-http-header-status"], d, pre_append)
			tmp["dp_http_header_status"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpHttpHeaderStatus")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_suppress"
		if _, ok := i["dp-http-header-suppress"]; ok {
			v := flattenUserRadiusDynamicMappingDpHttpHeaderSuppress(i["dp-http-header-suppress"], d, pre_append)
			tmp["dp_http_header_suppress"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpHttpHeaderSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_dyn_flags"
		if _, ok := i["dp-log-dyn_flags"]; ok {
			v := flattenUserRadiusDynamicMappingDpLogDynFlags(i["dp-log-dyn_flags"], d, pre_append)
			tmp["dp_log_dyn_flags"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpLogDynFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_period"
		if _, ok := i["dp-log-period"]; ok {
			v := flattenUserRadiusDynamicMappingDpLogPeriod(i["dp-log-period"], d, pre_append)
			tmp["dp_log_period"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_mem_percent"
		if _, ok := i["dp-mem-percent"]; ok {
			v := flattenUserRadiusDynamicMappingDpMemPercent(i["dp-mem-percent"], d, pre_append)
			tmp["dp_mem_percent"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpMemPercent")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute"
		if _, ok := i["dp-profile-attribute"]; ok {
			v := flattenUserRadiusDynamicMappingDpProfileAttribute(i["dp-profile-attribute"], d, pre_append)
			tmp["dp_profile_attribute"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpProfileAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute_key"
		if _, ok := i["dp-profile-attribute-key"]; ok {
			v := flattenUserRadiusDynamicMappingDpProfileAttributeKey(i["dp-profile-attribute-key"], d, pre_append)
			tmp["dp_profile_attribute_key"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpProfileAttributeKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_response"
		if _, ok := i["dp-radius-response"]; ok {
			v := flattenUserRadiusDynamicMappingDpRadiusResponse(i["dp-radius-response"], d, pre_append)
			tmp["dp_radius_response"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpRadiusResponse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_server_port"
		if _, ok := i["dp-radius-server-port"]; ok {
			v := flattenUserRadiusDynamicMappingDpRadiusServerPort(i["dp-radius-server-port"], d, pre_append)
			tmp["dp_radius_server_port"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpRadiusServerPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_validate_request_secret"
		if _, ok := i["dp-validate-request-secret"]; ok {
			v := flattenUserRadiusDynamicMappingDpValidateRequestSecret(i["dp-validate-request-secret"], d, pre_append)
			tmp["dp_validate_request_secret"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DpValidateRequestSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_profile"
		if _, ok := i["dynamic-profile"]; ok {
			v := flattenUserRadiusDynamicMappingDynamicProfile(i["dynamic-profile"], d, pre_append)
			tmp["dynamic_profile"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-DynamicProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endpoint_translation"
		if _, ok := i["endpoint-translation"]; ok {
			v := flattenUserRadiusDynamicMappingEndpointTranslation(i["endpoint-translation"], d, pre_append)
			tmp["endpoint_translation"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EndpointTranslation")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_convert_hex"
		if _, ok := i["ep-carrier-endpoint-convert-hex"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointConvertHex(i["ep-carrier-endpoint-convert-hex"], d, pre_append)
			tmp["ep_carrier_endpoint_convert_hex"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointConvertHex")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header"
		if _, ok := i["ep-carrier-endpoint-header"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointHeader(i["ep-carrier-endpoint-header"], d, pre_append)
			tmp["ep_carrier_endpoint_header"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header_suppress"
		if _, ok := i["ep-carrier-endpoint-header-suppress"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(i["ep-carrier-endpoint-header-suppress"], d, pre_append)
			tmp["ep_carrier_endpoint_header_suppress"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointHeaderSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix"
		if _, ok := i["ep-carrier-endpoint-prefix"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointPrefix(i["ep-carrier-endpoint-prefix"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointPrefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_max"
		if _, ok := i["ep-carrier-endpoint-prefix-range-max"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(i["ep-carrier-endpoint-prefix-range-max"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix_range_max"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointPrefixRangeMax")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_min"
		if _, ok := i["ep-carrier-endpoint-prefix-range-min"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(i["ep-carrier-endpoint-prefix-range-min"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix_range_min"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointPrefixRangeMin")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_string"
		if _, ok := i["ep-carrier-endpoint-prefix-string"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointPrefixString(i["ep-carrier-endpoint-prefix-string"], d, pre_append)
			tmp["ep_carrier_endpoint_prefix_string"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointPrefixString")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_source"
		if _, ok := i["ep-carrier-endpoint-source"]; ok {
			v := flattenUserRadiusDynamicMappingEpCarrierEndpointSource(i["ep-carrier-endpoint-source"], d, pre_append)
			tmp["ep_carrier_endpoint_source"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpCarrierEndpointSource")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header"
		if _, ok := i["ep-ip-header"]; ok {
			v := flattenUserRadiusDynamicMappingEpIpHeader(i["ep-ip-header"], d, pre_append)
			tmp["ep_ip_header"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpIpHeader")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header_suppress"
		if _, ok := i["ep-ip-header-suppress"]; ok {
			v := flattenUserRadiusDynamicMappingEpIpHeaderSuppress(i["ep-ip-header-suppress"], d, pre_append)
			tmp["ep_ip_header_suppress"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpIpHeaderSuppress")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_missing_header_fallback"
		if _, ok := i["ep-missing-header-fallback"]; ok {
			v := flattenUserRadiusDynamicMappingEpMissingHeaderFallback(i["ep-missing-header-fallback"], d, pre_append)
			tmp["ep_missing_header_fallback"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpMissingHeaderFallback")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_profile_query_type"
		if _, ok := i["ep-profile-query-type"]; ok {
			v := flattenUserRadiusDynamicMappingEpProfileQueryType(i["ep-profile-query-type"], d, pre_append)
			tmp["ep_profile_query_type"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-EpProfileQueryType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_override_attr_type"
		if _, ok := i["group-override-attr-type"]; ok {
			v := flattenUserRadiusDynamicMappingGroupOverrideAttrType(i["group-override-attr-type"], d, pre_append)
			tmp["group_override_attr_type"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-GroupOverrideAttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3c_compatibility"
		if _, ok := i["h3c-compatibility"]; ok {
			v := flattenUserRadiusDynamicMappingH3CCompatibility(i["h3c-compatibility"], d, pre_append)
			tmp["h3c_compatibility"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-H3CCompatibility")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenUserRadiusDynamicMappingInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenUserRadiusDynamicMappingInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_case"
		if _, ok := i["mac-case"]; ok {
			v := flattenUserRadiusDynamicMappingMacCase(i["mac-case"], d, pre_append)
			tmp["mac_case"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-MacCase")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_password_delimiter"
		if _, ok := i["mac-password-delimiter"]; ok {
			v := flattenUserRadiusDynamicMappingMacPasswordDelimiter(i["mac-password-delimiter"], d, pre_append)
			tmp["mac_password_delimiter"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-MacPasswordDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_username_delimiter"
		if _, ok := i["mac-username-delimiter"]; ok {
			v := flattenUserRadiusDynamicMappingMacUsernameDelimiter(i["mac-username-delimiter"], d, pre_append)
			tmp["mac_username_delimiter"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-MacUsernameDelimiter")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_id"
		if _, ok := i["nas-id"]; ok {
			v := flattenUserRadiusDynamicMappingNasId(i["nas-id"], d, pre_append)
			tmp["nas_id"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-NasId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_id_type"
		if _, ok := i["nas-id-type"]; ok {
			v := flattenUserRadiusDynamicMappingNasIdType(i["nas-id-type"], d, pre_append)
			tmp["nas_id_type"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-NasIdType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_ip"
		if _, ok := i["nas-ip"]; ok {
			v := flattenUserRadiusDynamicMappingNasIp(i["nas-ip"], d, pre_append)
			tmp["nas_ip"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-NasIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_encoding"
		if _, ok := i["password-encoding"]; ok {
			v := flattenUserRadiusDynamicMappingPasswordEncoding(i["password-encoding"], d, pre_append)
			tmp["password_encoding"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-PasswordEncoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := i["password-renewal"]; ok {
			v := flattenUserRadiusDynamicMappingPasswordRenewal(i["password-renewal"], d, pre_append)
			tmp["password_renewal"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-PasswordRenewal")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_coa"
		if _, ok := i["radius-coa"]; ok {
			v := flattenUserRadiusDynamicMappingRadiusCoa(i["radius-coa"], d, pre_append)
			tmp["radius_coa"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RadiusCoa")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_port"
		if _, ok := i["radius-port"]; ok {
			v := flattenUserRadiusDynamicMappingRadiusPort(i["radius-port"], d, pre_append)
			tmp["radius_port"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RadiusPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "require_message_authenticator"
		if _, ok := i["require-message-authenticator"]; ok {
			v := flattenUserRadiusDynamicMappingRequireMessageAuthenticator(i["require-message-authenticator"], d, pre_append)
			tmp["require_message_authenticator"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RequireMessageAuthenticator")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso"
		if _, ok := i["rsso"]; ok {
			v := flattenUserRadiusDynamicMappingRsso(i["rsso"], d, pre_append)
			tmp["rsso"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Rsso")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_context_timeout"
		if _, ok := i["rsso-context-timeout"]; ok {
			v := flattenUserRadiusDynamicMappingRssoContextTimeout(i["rsso-context-timeout"], d, pre_append)
			tmp["rsso_context_timeout"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoContextTimeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_attribute"
		if _, ok := i["rsso-endpoint-attribute"]; ok {
			v := flattenUserRadiusDynamicMappingRssoEndpointAttribute(i["rsso-endpoint-attribute"], d, pre_append)
			tmp["rsso_endpoint_attribute"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoEndpointAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_block_attribute"
		if _, ok := i["rsso-endpoint-block-attribute"]; ok {
			v := flattenUserRadiusDynamicMappingRssoEndpointBlockAttribute(i["rsso-endpoint-block-attribute"], d, pre_append)
			tmp["rsso_endpoint_block_attribute"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoEndpointBlockAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_ep_one_ip_only"
		if _, ok := i["rsso-ep-one-ip-only"]; ok {
			v := flattenUserRadiusDynamicMappingRssoEpOneIpOnly(i["rsso-ep-one-ip-only"], d, pre_append)
			tmp["rsso_ep_one_ip_only"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoEpOneIpOnly")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_flush_ip_session"
		if _, ok := i["rsso-flush-ip-session"]; ok {
			v := flattenUserRadiusDynamicMappingRssoFlushIpSession(i["rsso-flush-ip-session"], d, pre_append)
			tmp["rsso_flush_ip_session"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoFlushIpSession")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_flags"
		if _, ok := i["rsso-log-flags"]; ok {
			v := flattenUserRadiusDynamicMappingRssoLogFlags(i["rsso-log-flags"], d, pre_append)
			tmp["rsso_log_flags"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoLogFlags")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_period"
		if _, ok := i["rsso-log-period"]; ok {
			v := flattenUserRadiusDynamicMappingRssoLogPeriod(i["rsso-log-period"], d, pre_append)
			tmp["rsso_log_period"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoLogPeriod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_response"
		if _, ok := i["rsso-radius-response"]; ok {
			v := flattenUserRadiusDynamicMappingRssoRadiusResponse(i["rsso-radius-response"], d, pre_append)
			tmp["rsso_radius_response"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoRadiusResponse")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_server_port"
		if _, ok := i["rsso-radius-server-port"]; ok {
			v := flattenUserRadiusDynamicMappingRssoRadiusServerPort(i["rsso-radius-server-port"], d, pre_append)
			tmp["rsso_radius_server_port"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoRadiusServerPort")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_validate_request_secret"
		if _, ok := i["rsso-validate-request-secret"]; ok {
			v := flattenUserRadiusDynamicMappingRssoValidateRequestSecret(i["rsso-validate-request-secret"], d, pre_append)
			tmp["rsso_validate_request_secret"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-RssoValidateRequestSecret")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := i["secondary-server"]; ok {
			v := flattenUserRadiusDynamicMappingSecondaryServer(i["secondary-server"], d, pre_append)
			tmp["secondary_server"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SecondaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenUserRadiusDynamicMappingServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_identity_check"
		if _, ok := i["server-identity-check"]; ok {
			v := flattenUserRadiusDynamicMappingServerIdentityCheck(i["server-identity-check"], d, pre_append)
			tmp["server_identity_check"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-ServerIdentityCheck")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenUserRadiusDynamicMappingSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_interface"
		if _, ok := i["source-ip-interface"]; ok {
			v := flattenUserRadiusDynamicMappingSourceIpInterface(i["source-ip-interface"], d, pre_append)
			tmp["source_ip_interface"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SourceIpInterface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute"
		if _, ok := i["sso-attribute"]; ok {
			v := flattenUserRadiusDynamicMappingSsoAttribute(i["sso-attribute"], d, pre_append)
			tmp["sso_attribute"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SsoAttribute")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_key"
		if _, ok := i["sso-attribute-key"]; ok {
			v := flattenUserRadiusDynamicMappingSsoAttributeKey(i["sso-attribute-key"], d, pre_append)
			tmp["sso_attribute_key"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SsoAttributeKey")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value_override"
		if _, ok := i["sso-attribute-value-override"]; ok {
			v := flattenUserRadiusDynamicMappingSsoAttributeValueOverride(i["sso-attribute-value-override"], d, pre_append)
			tmp["sso_attribute_value_override"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SsoAttributeValueOverride")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status_ttl"
		if _, ok := i["status-ttl"]; ok {
			v := flattenUserRadiusDynamicMappingStatusTtl(i["status-ttl"], d, pre_append)
			tmp["status_ttl"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-StatusTtl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_acct_fast_framedip_detect"
		if _, ok := i["switch-controller-acct-fast-framedip-detect"]; ok {
			v := flattenUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(i["switch-controller-acct-fast-framedip-detect"], d, pre_append)
			tmp["switch_controller_acct_fast_framedip_detect"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SwitchControllerAcctFastFramedipDetect")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_nas_ip_dynamic"
		if _, ok := i["switch-controller-nas-ip-dynamic"]; ok {
			v := flattenUserRadiusDynamicMappingSwitchControllerNasIpDynamic(i["switch-controller-nas-ip-dynamic"], d, pre_append)
			tmp["switch_controller_nas_ip_dynamic"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SwitchControllerNasIpDynamic")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_service_type"
		if _, ok := i["switch-controller-service-type"]; ok {
			v := flattenUserRadiusDynamicMappingSwitchControllerServiceType(i["switch-controller-service-type"], d, pre_append)
			tmp["switch_controller_service_type"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-SwitchControllerServiceType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := i["tertiary-server"]; ok {
			v := flattenUserRadiusDynamicMappingTertiaryServer(i["tertiary-server"], d, pre_append)
			tmp["tertiary_server"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-TertiaryServer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout"
		if _, ok := i["timeout"]; ok {
			v := flattenUserRadiusDynamicMappingTimeout(i["timeout"], d, pre_append)
			tmp["timeout"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-Timeout")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tls_min_proto_version"
		if _, ok := i["tls-min-proto-version"]; ok {
			v := flattenUserRadiusDynamicMappingTlsMinProtoVersion(i["tls-min-proto-version"], d, pre_append)
			tmp["tls_min_proto_version"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-TlsMinProtoVersion")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport_protocol"
		if _, ok := i["transport-protocol"]; ok {
			v := flattenUserRadiusDynamicMappingTransportProtocol(i["transport-protocol"], d, pre_append)
			tmp["transport_protocol"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-TransportProtocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_group_for_profile"
		if _, ok := i["use-group-for-profile"]; ok {
			v := flattenUserRadiusDynamicMappingUseGroupForProfile(i["use-group-for-profile"], d, pre_append)
			tmp["use_group_for_profile"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-UseGroupForProfile")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_management_vdom"
		if _, ok := i["use-management-vdom"]; ok {
			v := flattenUserRadiusDynamicMappingUseManagementVdom(i["use-management-vdom"], d, pre_append)
			tmp["use_management_vdom"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-UseManagementVdom")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_sensitive"
		if _, ok := i["username-case-sensitive"]; ok {
			v := flattenUserRadiusDynamicMappingUsernameCaseSensitive(i["username-case-sensitive"], d, pre_append)
			tmp["username_case_sensitive"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-UsernameCaseSensitive")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenUserRadiusDynamicMappingVrfSelect(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "UserRadius-DynamicMapping-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserRadiusDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenUserRadiusDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenUserRadiusDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserRadiusDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountKeyCertField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountKeyProcessing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserRadiusDynamicMappingAccountingServerId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerInterface(i["interface"], d, pre_append)
			tmp["interface"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-Interface")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
			tmp["interface_select_method"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-InterfaceSelectMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerPort(i["port"], d, pre_append)
			tmp["port"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-Port")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerServer(i["server"], d, pre_append)
			tmp["server"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-Server")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerSourceIp(i["source-ip"], d, pre_append)
			tmp["source_ip"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-SourceIp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := i["vrf-select"]; ok {
			v := flattenUserRadiusDynamicMappingAccountingServerVrfSelect(i["vrf-select"], d, pre_append)
			tmp["vrf_select"] = fortiAPISubPartPatch(v, "UserRadiusDynamicMapping-AccountingServer-VrfSelect")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserRadiusDynamicMappingAccountingServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServerServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAccountingServerVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAcctAllServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAllUsergroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingCallStationIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingClass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpCarrierEndpointAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpContextTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpFlushIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpHttpHeaderFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpHttpHeaderStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpHttpHeaderSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpLogDynFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingDpLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpMemPercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpProfileAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpProfileAttributeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpRadiusResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpRadiusServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDpValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingDynamicProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEndpointTranslation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointConvertHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointPrefixString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpCarrierEndpointSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpIpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpIpHeaderSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpMissingHeaderFallback(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingEpProfileQueryType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingGroupOverrideAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingH3CCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingNasId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingNasIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingPasswordEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRadiusCoa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRequireMessageAuthenticator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoContextTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoEndpointAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoEpOneIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoFlushIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoLogFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingRssoLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoRadiusResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoRadiusServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingRssoValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingSsoAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSsoAttributeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSsoAttributeValueOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingStatusTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSwitchControllerNasIpDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingSwitchControllerServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusDynamicMappingTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingTlsMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingTransportProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingUseGroupForProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingUseManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingUsernameCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusDynamicMappingVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusGroupOverrideAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusH3CCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusMacCase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusNasId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusNasIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusNasIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusPasswordEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusPasswordRenewal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRadiusCoa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRadiusPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRequireMessageAuthenticator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoContextTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoEndpointAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoEpOneIpOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoFlushIpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoLogFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusRssoLogPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoRadiusResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoRadiusServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusRssoValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSecondaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSourceIpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusSsoAttribute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSsoAttributeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSsoAttributeValueOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusStatusTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSwitchControllerAcctFastFramedipDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSwitchControllerNasIpDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusSwitchControllerServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserRadiusTertiaryServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusTlsMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusTransportProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusUseManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusUsernameCaseSensitive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserRadiusVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserRadius(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("account_key_cert_field", flattenUserRadiusAccountKeyCertField(o["account-key-cert-field"], d, "account_key_cert_field")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-cert-field"], "UserRadius-AccountKeyCertField"); ok {
			if err = d.Set("account_key_cert_field", vv); err != nil {
				return fmt.Errorf("Error reading account_key_cert_field: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_cert_field: %v", err)
		}
	}

	if err = d.Set("account_key_processing", flattenUserRadiusAccountKeyProcessing(o["account-key-processing"], d, "account_key_processing")); err != nil {
		if vv, ok := fortiAPIPatch(o["account-key-processing"], "UserRadius-AccountKeyProcessing"); ok {
			if err = d.Set("account_key_processing", vv); err != nil {
				return fmt.Errorf("Error reading account_key_processing: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading account_key_processing: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("accounting_server", flattenUserRadiusAccountingServer(o["accounting-server"], d, "accounting_server")); err != nil {
			if vv, ok := fortiAPIPatch(o["accounting-server"], "UserRadius-AccountingServer"); ok {
				if err = d.Set("accounting_server", vv); err != nil {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading accounting_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("accounting_server"); ok {
			if err = d.Set("accounting_server", flattenUserRadiusAccountingServer(o["accounting-server"], d, "accounting_server")); err != nil {
				if vv, ok := fortiAPIPatch(o["accounting-server"], "UserRadius-AccountingServer"); ok {
					if err = d.Set("accounting_server", vv); err != nil {
						return fmt.Errorf("Error reading accounting_server: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("acct_all_servers", flattenUserRadiusAcctAllServers(o["acct-all-servers"], d, "acct_all_servers")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-all-servers"], "UserRadius-AcctAllServers"); ok {
			if err = d.Set("acct_all_servers", vv); err != nil {
				return fmt.Errorf("Error reading acct_all_servers: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_all_servers: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenUserRadiusAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["acct-interim-interval"], "UserRadius-AcctInterimInterval"); ok {
			if err = d.Set("acct_interim_interval", vv); err != nil {
				return fmt.Errorf("Error reading acct_interim_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("all_usergroup", flattenUserRadiusAllUsergroup(o["all-usergroup"], d, "all_usergroup")); err != nil {
		if vv, ok := fortiAPIPatch(o["all-usergroup"], "UserRadius-AllUsergroup"); ok {
			if err = d.Set("all_usergroup", vv); err != nil {
				return fmt.Errorf("Error reading all_usergroup: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading all_usergroup: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserRadiusAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-type"], "UserRadius-AuthType"); ok {
			if err = d.Set("auth_type", vv); err != nil {
				return fmt.Errorf("Error reading auth_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenUserRadiusCaCert(o["ca-cert"], d, "ca_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["ca-cert"], "UserRadius-CaCert"); ok {
			if err = d.Set("ca_cert", vv); err != nil {
				return fmt.Errorf("Error reading ca_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("call_station_id_type", flattenUserRadiusCallStationIdType(o["call-station-id-type"], d, "call_station_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["call-station-id-type"], "UserRadius-CallStationIdType"); ok {
			if err = d.Set("call_station_id_type", vv); err != nil {
				return fmt.Errorf("Error reading call_station_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading call_station_id_type: %v", err)
		}
	}

	if err = d.Set("class", flattenUserRadiusClass(o["class"], d, "class")); err != nil {
		if vv, ok := fortiAPIPatch(o["class"], "UserRadius-Class"); ok {
			if err = d.Set("class", vv); err != nil {
				return fmt.Errorf("Error reading class: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading class: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenUserRadiusClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["client-cert"], "UserRadius-ClientCert"); ok {
			if err = d.Set("client_cert", vv); err != nil {
				return fmt.Errorf("Error reading client_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("delimiter", flattenUserRadiusDelimiter(o["delimiter"], d, "delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["delimiter"], "UserRadius-Delimiter"); ok {
			if err = d.Set("delimiter", vv); err != nil {
				return fmt.Errorf("Error reading delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delimiter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenUserRadiusDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserRadius-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenUserRadiusDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserRadius-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_override_attr_type", flattenUserRadiusGroupOverrideAttrType(o["group-override-attr-type"], d, "group_override_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-override-attr-type"], "UserRadius-GroupOverrideAttrType"); ok {
			if err = d.Set("group_override_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading group_override_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_override_attr_type: %v", err)
		}
	}

	if err = d.Set("h3c_compatibility", flattenUserRadiusH3CCompatibility(o["h3c-compatibility"], d, "h3c_compatibility")); err != nil {
		if vv, ok := fortiAPIPatch(o["h3c-compatibility"], "UserRadius-H3CCompatibility"); ok {
			if err = d.Set("h3c_compatibility", vv); err != nil {
				return fmt.Errorf("Error reading h3c_compatibility: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading h3c_compatibility: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserRadiusInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "UserRadius-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserRadiusInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface-select-method"], "UserRadius-InterfaceSelectMethod"); ok {
			if err = d.Set("interface_select_method", vv); err != nil {
				return fmt.Errorf("Error reading interface_select_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenUserRadiusMacCase(o["mac-case"], d, "mac_case")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-case"], "UserRadius-MacCase"); ok {
			if err = d.Set("mac_case", vv); err != nil {
				return fmt.Errorf("Error reading mac_case: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenUserRadiusMacPasswordDelimiter(o["mac-password-delimiter"], d, "mac_password_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-password-delimiter"], "UserRadius-MacPasswordDelimiter"); ok {
			if err = d.Set("mac_password_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenUserRadiusMacUsernameDelimiter(o["mac-username-delimiter"], d, "mac_username_delimiter")); err != nil {
		if vv, ok := fortiAPIPatch(o["mac-username-delimiter"], "UserRadius-MacUsernameDelimiter"); ok {
			if err = d.Set("mac_username_delimiter", vv); err != nil {
				return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("name", flattenUserRadiusName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserRadius-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nas_id", flattenUserRadiusNasId(o["nas-id"], d, "nas_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-id"], "UserRadius-NasId"); ok {
			if err = d.Set("nas_id", vv); err != nil {
				return fmt.Errorf("Error reading nas_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_id: %v", err)
		}
	}

	if err = d.Set("nas_id_type", flattenUserRadiusNasIdType(o["nas-id-type"], d, "nas_id_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-id-type"], "UserRadius-NasIdType"); ok {
			if err = d.Set("nas_id_type", vv); err != nil {
				return fmt.Errorf("Error reading nas_id_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_id_type: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenUserRadiusNasIp(o["nas-ip"], d, "nas_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["nas-ip"], "UserRadius-NasIp"); ok {
			if err = d.Set("nas_ip", vv); err != nil {
				return fmt.Errorf("Error reading nas_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("password_encoding", flattenUserRadiusPasswordEncoding(o["password-encoding"], d, "password_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-encoding"], "UserRadius-PasswordEncoding"); ok {
			if err = d.Set("password_encoding", vv); err != nil {
				return fmt.Errorf("Error reading password_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_encoding: %v", err)
		}
	}

	if err = d.Set("password_renewal", flattenUserRadiusPasswordRenewal(o["password-renewal"], d, "password_renewal")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-renewal"], "UserRadius-PasswordRenewal"); ok {
			if err = d.Set("password_renewal", vv); err != nil {
				return fmt.Errorf("Error reading password_renewal: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("radius_coa", flattenUserRadiusRadiusCoa(o["radius-coa"], d, "radius_coa")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-coa"], "UserRadius-RadiusCoa"); ok {
			if err = d.Set("radius_coa", vv); err != nil {
				return fmt.Errorf("Error reading radius_coa: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_coa: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenUserRadiusRadiusPort(o["radius-port"], d, "radius_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-port"], "UserRadius-RadiusPort"); ok {
			if err = d.Set("radius_port", vv); err != nil {
				return fmt.Errorf("Error reading radius_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("require_message_authenticator", flattenUserRadiusRequireMessageAuthenticator(o["require-message-authenticator"], d, "require_message_authenticator")); err != nil {
		if vv, ok := fortiAPIPatch(o["require-message-authenticator"], "UserRadius-RequireMessageAuthenticator"); ok {
			if err = d.Set("require_message_authenticator", vv); err != nil {
				return fmt.Errorf("Error reading require_message_authenticator: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading require_message_authenticator: %v", err)
		}
	}

	if err = d.Set("rsso", flattenUserRadiusRsso(o["rsso"], d, "rsso")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso"], "UserRadius-Rsso"); ok {
			if err = d.Set("rsso", vv); err != nil {
				return fmt.Errorf("Error reading rsso: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("rsso_context_timeout", flattenUserRadiusRssoContextTimeout(o["rsso-context-timeout"], d, "rsso_context_timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-context-timeout"], "UserRadius-RssoContextTimeout"); ok {
			if err = d.Set("rsso_context_timeout", vv); err != nil {
				return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_attribute", flattenUserRadiusRssoEndpointAttribute(o["rsso-endpoint-attribute"], d, "rsso_endpoint_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-endpoint-attribute"], "UserRadius-RssoEndpointAttribute"); ok {
			if err = d.Set("rsso_endpoint_attribute", vv); err != nil {
				return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_block_attribute", flattenUserRadiusRssoEndpointBlockAttribute(o["rsso-endpoint-block-attribute"], d, "rsso_endpoint_block_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-endpoint-block-attribute"], "UserRadius-RssoEndpointBlockAttribute"); ok {
			if err = d.Set("rsso_endpoint_block_attribute", vv); err != nil {
				return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_ep_one_ip_only", flattenUserRadiusRssoEpOneIpOnly(o["rsso-ep-one-ip-only"], d, "rsso_ep_one_ip_only")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-ep-one-ip-only"], "UserRadius-RssoEpOneIpOnly"); ok {
			if err = d.Set("rsso_ep_one_ip_only", vv); err != nil {
				return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
		}
	}

	if err = d.Set("rsso_flush_ip_session", flattenUserRadiusRssoFlushIpSession(o["rsso-flush-ip-session"], d, "rsso_flush_ip_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-flush-ip-session"], "UserRadius-RssoFlushIpSession"); ok {
			if err = d.Set("rsso_flush_ip_session", vv); err != nil {
				return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
		}
	}

	if err = d.Set("rsso_log_flags", flattenUserRadiusRssoLogFlags(o["rsso-log-flags"], d, "rsso_log_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-log-flags"], "UserRadius-RssoLogFlags"); ok {
			if err = d.Set("rsso_log_flags", vv); err != nil {
				return fmt.Errorf("Error reading rsso_log_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_log_flags: %v", err)
		}
	}

	if err = d.Set("rsso_log_period", flattenUserRadiusRssoLogPeriod(o["rsso-log-period"], d, "rsso_log_period")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-log-period"], "UserRadius-RssoLogPeriod"); ok {
			if err = d.Set("rsso_log_period", vv); err != nil {
				return fmt.Errorf("Error reading rsso_log_period: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_log_period: %v", err)
		}
	}

	if err = d.Set("rsso_radius_response", flattenUserRadiusRssoRadiusResponse(o["rsso-radius-response"], d, "rsso_radius_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-radius-response"], "UserRadius-RssoRadiusResponse"); ok {
			if err = d.Set("rsso_radius_response", vv); err != nil {
				return fmt.Errorf("Error reading rsso_radius_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_radius_response: %v", err)
		}
	}

	if err = d.Set("rsso_radius_server_port", flattenUserRadiusRssoRadiusServerPort(o["rsso-radius-server-port"], d, "rsso_radius_server_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-radius-server-port"], "UserRadius-RssoRadiusServerPort"); ok {
			if err = d.Set("rsso_radius_server_port", vv); err != nil {
				return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
		}
	}

	if err = d.Set("rsso_validate_request_secret", flattenUserRadiusRssoValidateRequestSecret(o["rsso-validate-request-secret"], d, "rsso_validate_request_secret")); err != nil {
		if vv, ok := fortiAPIPatch(o["rsso-validate-request-secret"], "UserRadius-RssoValidateRequestSecret"); ok {
			if err = d.Set("rsso_validate_request_secret", vv); err != nil {
				return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserRadiusSecondaryServer(o["secondary-server"], d, "secondary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["secondary-server"], "UserRadius-SecondaryServer"); ok {
			if err = d.Set("secondary_server", vv); err != nil {
				return fmt.Errorf("Error reading secondary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("server", flattenUserRadiusServer(o["server"], d, "server")); err != nil {
		if vv, ok := fortiAPIPatch(o["server"], "UserRadius-Server"); ok {
			if err = d.Set("server", vv); err != nil {
				return fmt.Errorf("Error reading server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_identity_check", flattenUserRadiusServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-identity-check"], "UserRadius-ServerIdentityCheck"); ok {
			if err = d.Set("server_identity_check", vv); err != nil {
				return fmt.Errorf("Error reading server_identity_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserRadiusSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip"], "UserRadius-SourceIp"); ok {
			if err = d.Set("source_ip", vv); err != nil {
				return fmt.Errorf("Error reading source_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenUserRadiusSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-ip-interface"], "UserRadius-SourceIpInterface"); ok {
			if err = d.Set("source_ip_interface", vv); err != nil {
				return fmt.Errorf("Error reading source_ip_interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("sso_attribute", flattenUserRadiusSsoAttribute(o["sso-attribute"], d, "sso_attribute")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute"], "UserRadius-SsoAttribute"); ok {
			if err = d.Set("sso_attribute", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute: %v", err)
		}
	}

	if err = d.Set("sso_attribute_key", flattenUserRadiusSsoAttributeKey(o["sso-attribute-key"], d, "sso_attribute_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-key"], "UserRadius-SsoAttributeKey"); ok {
			if err = d.Set("sso_attribute_key", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_key: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value_override", flattenUserRadiusSsoAttributeValueOverride(o["sso-attribute-value-override"], d, "sso_attribute_value_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["sso-attribute-value-override"], "UserRadius-SsoAttributeValueOverride"); ok {
			if err = d.Set("sso_attribute_value_override", vv); err != nil {
				return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
		}
	}

	if err = d.Set("status_ttl", flattenUserRadiusStatusTtl(o["status-ttl"], d, "status_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["status-ttl"], "UserRadius-StatusTtl"); ok {
			if err = d.Set("status_ttl", vv); err != nil {
				return fmt.Errorf("Error reading status_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status_ttl: %v", err)
		}
	}

	if err = d.Set("switch_controller_acct_fast_framedip_detect", flattenUserRadiusSwitchControllerAcctFastFramedipDetect(o["switch-controller-acct-fast-framedip-detect"], d, "switch_controller_acct_fast_framedip_detect")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-acct-fast-framedip-detect"], "UserRadius-SwitchControllerAcctFastFramedipDetect"); ok {
			if err = d.Set("switch_controller_acct_fast_framedip_detect", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
		}
	}

	if err = d.Set("switch_controller_nas_ip_dynamic", flattenUserRadiusSwitchControllerNasIpDynamic(o["switch-controller-nas-ip-dynamic"], d, "switch_controller_nas_ip_dynamic")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-nas-ip-dynamic"], "UserRadius-SwitchControllerNasIpDynamic"); ok {
			if err = d.Set("switch_controller_nas_ip_dynamic", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_nas_ip_dynamic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_nas_ip_dynamic: %v", err)
		}
	}

	if err = d.Set("switch_controller_service_type", flattenUserRadiusSwitchControllerServiceType(o["switch-controller-service-type"], d, "switch_controller_service_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["switch-controller-service-type"], "UserRadius-SwitchControllerServiceType"); ok {
			if err = d.Set("switch_controller_service_type", vv); err != nil {
				return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserRadiusTertiaryServer(o["tertiary-server"], d, "tertiary_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["tertiary-server"], "UserRadius-TertiaryServer"); ok {
			if err = d.Set("tertiary_server", vv); err != nil {
				return fmt.Errorf("Error reading tertiary_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("timeout", flattenUserRadiusTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "UserRadius-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("tls_min_proto_version", flattenUserRadiusTlsMinProtoVersion(o["tls-min-proto-version"], d, "tls_min_proto_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["tls-min-proto-version"], "UserRadius-TlsMinProtoVersion"); ok {
			if err = d.Set("tls_min_proto_version", vv); err != nil {
				return fmt.Errorf("Error reading tls_min_proto_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tls_min_proto_version: %v", err)
		}
	}

	if err = d.Set("transport_protocol", flattenUserRadiusTransportProtocol(o["transport-protocol"], d, "transport_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["transport-protocol"], "UserRadius-TransportProtocol"); ok {
			if err = d.Set("transport_protocol", vv); err != nil {
				return fmt.Errorf("Error reading transport_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading transport_protocol: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenUserRadiusUseManagementVdom(o["use-management-vdom"], d, "use_management_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["use-management-vdom"], "UserRadius-UseManagementVdom"); ok {
			if err = d.Set("use_management_vdom", vv); err != nil {
				return fmt.Errorf("Error reading use_management_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("username_case_sensitive", flattenUserRadiusUsernameCaseSensitive(o["username-case-sensitive"], d, "username_case_sensitive")); err != nil {
		if vv, ok := fortiAPIPatch(o["username-case-sensitive"], "UserRadius-UsernameCaseSensitive"); ok {
			if err = d.Set("username_case_sensitive", vv); err != nil {
				return fmt.Errorf("Error reading username_case_sensitive: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading username_case_sensitive: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserRadiusVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrf-select"], "UserRadius-VrfSelect"); ok {
			if err = d.Set("vrf_select", vv); err != nil {
				return fmt.Errorf("Error reading vrf_select: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenUserRadiusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserRadiusAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandUserRadiusAccountingServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandUserRadiusAccountingServerInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandUserRadiusAccountingServerInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserRadiusAccountingServerPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secret"], _ = expandUserRadiusAccountingServerSecret(d, i["secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandUserRadiusAccountingServerServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandUserRadiusAccountingServerSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandUserRadiusAccountingServerStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandUserRadiusAccountingServerVrfSelect(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserRadiusAccountingServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusAccountingServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusAccountingServerServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctAllServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAllUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusCallStationIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "_scope"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserRadiusDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_cert_field"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-cert-field"], _ = expandUserRadiusDynamicMappingAccountKeyCertField(d, i["account_key_cert_field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_key_processing"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["account-key-processing"], _ = expandUserRadiusDynamicMappingAccountKeyProcessing(d, i["account_key_processing"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accounting_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandUserRadiusDynamicMappingAccountingServer(d, i["accounting_server"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["accounting-server"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_all_servers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acct-all-servers"], _ = expandUserRadiusDynamicMappingAcctAllServers(d, i["acct_all_servers"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acct_interim_interval"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acct-interim-interval"], _ = expandUserRadiusDynamicMappingAcctInterimInterval(d, i["acct_interim_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "all_usergroup"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["all-usergroup"], _ = expandUserRadiusDynamicMappingAllUsergroup(d, i["all_usergroup"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-type"], _ = expandUserRadiusDynamicMappingAuthType(d, i["auth_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ca_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ca-cert"], _ = expandUserRadiusDynamicMappingCaCert(d, i["ca_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "call_station_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["call-station-id-type"], _ = expandUserRadiusDynamicMappingCallStationIdType(d, i["call_station_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "class"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["class"], _ = expandUserRadiusDynamicMappingClass(d, i["class"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["client-cert"], _ = expandUserRadiusDynamicMappingClientCert(d, i["client_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["delimiter"], _ = expandUserRadiusDynamicMappingDelimiter(d, i["delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-carrier-endpoint-attribute"], _ = expandUserRadiusDynamicMappingDpCarrierEndpointAttribute(d, i["dp_carrier_endpoint_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_carrier_endpoint_block_attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-carrier-endpoint-block-attribute"], _ = expandUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(d, i["dp_carrier_endpoint_block_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_context_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-context-timeout"], _ = expandUserRadiusDynamicMappingDpContextTimeout(d, i["dp_context_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_flush_ip_session"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-flush-ip-session"], _ = expandUserRadiusDynamicMappingDpFlushIpSession(d, i["dp_flush_ip_session"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_hold_time"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-hold-time"], _ = expandUserRadiusDynamicMappingDpHoldTime(d, i["dp_hold_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-http-header"], _ = expandUserRadiusDynamicMappingDpHttpHeader(d, i["dp_http_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_fallback"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-http-header-fallback"], _ = expandUserRadiusDynamicMappingDpHttpHeaderFallback(d, i["dp_http_header_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-http-header-status"], _ = expandUserRadiusDynamicMappingDpHttpHeaderStatus(d, i["dp_http_header_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_http_header_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-http-header-suppress"], _ = expandUserRadiusDynamicMappingDpHttpHeaderSuppress(d, i["dp_http_header_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_dyn_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-log-dyn_flags"], _ = expandUserRadiusDynamicMappingDpLogDynFlags(d, i["dp_log_dyn_flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-log-period"], _ = expandUserRadiusDynamicMappingDpLogPeriod(d, i["dp_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_mem_percent"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-mem-percent"], _ = expandUserRadiusDynamicMappingDpMemPercent(d, i["dp_mem_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-profile-attribute"], _ = expandUserRadiusDynamicMappingDpProfileAttribute(d, i["dp_profile_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_profile_attribute_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-profile-attribute-key"], _ = expandUserRadiusDynamicMappingDpProfileAttributeKey(d, i["dp_profile_attribute_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_response"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-radius-response"], _ = expandUserRadiusDynamicMappingDpRadiusResponse(d, i["dp_radius_response"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_radius_server_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-radius-server-port"], _ = expandUserRadiusDynamicMappingDpRadiusServerPort(d, i["dp_radius_server_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-secret"], _ = expandUserRadiusDynamicMappingDpSecret(d, i["dp_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dp_validate_request_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dp-validate-request-secret"], _ = expandUserRadiusDynamicMappingDpValidateRequestSecret(d, i["dp_validate_request_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dynamic_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dynamic-profile"], _ = expandUserRadiusDynamicMappingDynamicProfile(d, i["dynamic_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "endpoint_translation"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["endpoint-translation"], _ = expandUserRadiusDynamicMappingEndpointTranslation(d, i["endpoint_translation"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_convert_hex"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-convert-hex"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointConvertHex(d, i["ep_carrier_endpoint_convert_hex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-header"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointHeader(d, i["ep_carrier_endpoint_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_header_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-header-suppress"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(d, i["ep_carrier_endpoint_header_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-prefix"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointPrefix(d, i["ep_carrier_endpoint_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_max"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-prefix-range-max"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(d, i["ep_carrier_endpoint_prefix_range_max"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_range_min"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-prefix-range-min"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(d, i["ep_carrier_endpoint_prefix_range_min"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_prefix_string"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-prefix-string"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointPrefixString(d, i["ep_carrier_endpoint_prefix_string"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_carrier_endpoint_source"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-carrier-endpoint-source"], _ = expandUserRadiusDynamicMappingEpCarrierEndpointSource(d, i["ep_carrier_endpoint_source"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-ip-header"], _ = expandUserRadiusDynamicMappingEpIpHeader(d, i["ep_ip_header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_ip_header_suppress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-ip-header-suppress"], _ = expandUserRadiusDynamicMappingEpIpHeaderSuppress(d, i["ep_ip_header_suppress"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_missing_header_fallback"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-missing-header-fallback"], _ = expandUserRadiusDynamicMappingEpMissingHeaderFallback(d, i["ep_missing_header_fallback"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ep_profile_query_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ep-profile-query-type"], _ = expandUserRadiusDynamicMappingEpProfileQueryType(d, i["ep_profile_query_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_override_attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-override-attr-type"], _ = expandUserRadiusDynamicMappingGroupOverrideAttrType(d, i["group_override_attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3c_compatibility"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["h3c-compatibility"], _ = expandUserRadiusDynamicMappingH3CCompatibility(d, i["h3c_compatibility"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandUserRadiusDynamicMappingInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandUserRadiusDynamicMappingInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_case"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-case"], _ = expandUserRadiusDynamicMappingMacCase(d, i["mac_case"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_password_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-password-delimiter"], _ = expandUserRadiusDynamicMappingMacPasswordDelimiter(d, i["mac_password_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_username_delimiter"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mac-username-delimiter"], _ = expandUserRadiusDynamicMappingMacUsernameDelimiter(d, i["mac_username_delimiter"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nas-id"], _ = expandUserRadiusDynamicMappingNasId(d, i["nas_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_id_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nas-id-type"], _ = expandUserRadiusDynamicMappingNasIdType(d, i["nas_id_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nas_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["nas-ip"], _ = expandUserRadiusDynamicMappingNasIp(d, i["nas_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_encoding"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-encoding"], _ = expandUserRadiusDynamicMappingPasswordEncoding(d, i["password_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_renewal"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["password-renewal"], _ = expandUserRadiusDynamicMappingPasswordRenewal(d, i["password_renewal"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_coa"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-coa"], _ = expandUserRadiusDynamicMappingRadiusCoa(d, i["radius_coa"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "radius_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["radius-port"], _ = expandUserRadiusDynamicMappingRadiusPort(d, i["radius_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "require_message_authenticator"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["require-message-authenticator"], _ = expandUserRadiusDynamicMappingRequireMessageAuthenticator(d, i["require_message_authenticator"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso"], _ = expandUserRadiusDynamicMappingRsso(d, i["rsso"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_context_timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-context-timeout"], _ = expandUserRadiusDynamicMappingRssoContextTimeout(d, i["rsso_context_timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-endpoint-attribute"], _ = expandUserRadiusDynamicMappingRssoEndpointAttribute(d, i["rsso_endpoint_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_endpoint_block_attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-endpoint-block-attribute"], _ = expandUserRadiusDynamicMappingRssoEndpointBlockAttribute(d, i["rsso_endpoint_block_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_ep_one_ip_only"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-ep-one-ip-only"], _ = expandUserRadiusDynamicMappingRssoEpOneIpOnly(d, i["rsso_ep_one_ip_only"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_flush_ip_session"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-flush-ip-session"], _ = expandUserRadiusDynamicMappingRssoFlushIpSession(d, i["rsso_flush_ip_session"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_flags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-log-flags"], _ = expandUserRadiusDynamicMappingRssoLogFlags(d, i["rsso_log_flags"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_log_period"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-log-period"], _ = expandUserRadiusDynamicMappingRssoLogPeriod(d, i["rsso_log_period"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_response"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-radius-response"], _ = expandUserRadiusDynamicMappingRssoRadiusResponse(d, i["rsso_radius_response"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_radius_server_port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-radius-server-port"], _ = expandUserRadiusDynamicMappingRssoRadiusServerPort(d, i["rsso_radius_server_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-secret"], _ = expandUserRadiusDynamicMappingRssoSecret(d, i["rsso_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rsso_validate_request_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["rsso-validate-request-secret"], _ = expandUserRadiusDynamicMappingRssoValidateRequestSecret(d, i["rsso_validate_request_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-secret"], _ = expandUserRadiusDynamicMappingSecondarySecret(d, i["secondary_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secondary-server"], _ = expandUserRadiusDynamicMappingSecondaryServer(d, i["secondary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secret"], _ = expandUserRadiusDynamicMappingSecret(d, i["secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandUserRadiusDynamicMappingServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_identity_check"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server-identity-check"], _ = expandUserRadiusDynamicMappingServerIdentityCheck(d, i["server_identity_check"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandUserRadiusDynamicMappingSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip-interface"], _ = expandUserRadiusDynamicMappingSourceIpInterface(d, i["source_ip_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-attribute"], _ = expandUserRadiusDynamicMappingSsoAttribute(d, i["sso_attribute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-attribute-key"], _ = expandUserRadiusDynamicMappingSsoAttributeKey(d, i["sso_attribute_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_attribute_value_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sso-attribute-value-override"], _ = expandUserRadiusDynamicMappingSsoAttributeValueOverride(d, i["sso_attribute_value_override"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status-ttl"], _ = expandUserRadiusDynamicMappingStatusTtl(d, i["status_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_acct_fast_framedip_detect"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-controller-acct-fast-framedip-detect"], _ = expandUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(d, i["switch_controller_acct_fast_framedip_detect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_nas_ip_dynamic"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-controller-nas-ip-dynamic"], _ = expandUserRadiusDynamicMappingSwitchControllerNasIpDynamic(d, i["switch_controller_nas_ip_dynamic"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_controller_service_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["switch-controller-service-type"], _ = expandUserRadiusDynamicMappingSwitchControllerServiceType(d, i["switch_controller_service_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tertiary-secret"], _ = expandUserRadiusDynamicMappingTertiarySecret(d, i["tertiary_secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tertiary_server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tertiary-server"], _ = expandUserRadiusDynamicMappingTertiaryServer(d, i["tertiary_server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "timeout"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["timeout"], _ = expandUserRadiusDynamicMappingTimeout(d, i["timeout"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tls_min_proto_version"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["tls-min-proto-version"], _ = expandUserRadiusDynamicMappingTlsMinProtoVersion(d, i["tls_min_proto_version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transport_protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["transport-protocol"], _ = expandUserRadiusDynamicMappingTransportProtocol(d, i["transport_protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_group_for_profile"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["use-group-for-profile"], _ = expandUserRadiusDynamicMappingUseGroupForProfile(d, i["use_group_for_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_management_vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["use-management-vdom"], _ = expandUserRadiusDynamicMappingUseManagementVdom(d, i["use_management_vdom"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "username_case_sensitive"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["username-case-sensitive"], _ = expandUserRadiusDynamicMappingUsernameCaseSensitive(d, i["username_case_sensitive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandUserRadiusDynamicMappingVrfSelect(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserRadiusDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandUserRadiusDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandUserRadiusDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserRadiusDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountKeyCertField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountKeyProcessing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandUserRadiusDynamicMappingAccountingServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface"], _ = expandUserRadiusDynamicMappingAccountingServerInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-select-method"], _ = expandUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(d, i["interface_select_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["port"], _ = expandUserRadiusDynamicMappingAccountingServerPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["secret"], _ = expandUserRadiusDynamicMappingAccountingServerSecret(d, i["secret"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["server"], _ = expandUserRadiusDynamicMappingAccountingServerServer(d, i["server"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["source-ip"], _ = expandUserRadiusDynamicMappingAccountingServerSourceIp(d, i["source_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandUserRadiusDynamicMappingAccountingServerStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vrf-select"], _ = expandUserRadiusDynamicMappingAccountingServerVrfSelect(d, i["vrf_select"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserRadiusDynamicMappingAccountingServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingAccountingServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServerSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingAccountingServerServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServerSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAccountingServerVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAcctAllServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAllUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingCallStationIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingClientCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpCarrierEndpointAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpCarrierEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpContextTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpFlushIpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpHoldTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpHttpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpHttpHeaderFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpHttpHeaderStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpHttpHeaderSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpLogDynFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingDpLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpMemPercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpProfileAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpProfileAttributeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpRadiusResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpRadiusServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDpSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingDpValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingDynamicProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEndpointTranslation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointConvertHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointHeaderSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMax(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointPrefixRangeMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointPrefixString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpCarrierEndpointSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpIpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpIpHeaderSuppress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpMissingHeaderFallback(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingEpProfileQueryType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingGroupOverrideAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingH3CCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingNasId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingNasIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingPasswordEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRadiusCoa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRequireMessageAuthenticator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoContextTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoEndpointAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoEpOneIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoFlushIpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoLogFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingRssoLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoRadiusResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoRadiusServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingRssoSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingRssoValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSecondarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingSsoAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSsoAttributeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSsoAttributeValueOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingStatusTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSwitchControllerAcctFastFramedipDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSwitchControllerNasIpDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingSwitchControllerServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingTertiarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusDynamicMappingTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingTlsMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingTransportProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingUseGroupForProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingUseManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingUsernameCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDynamicMappingVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusGroupOverrideAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusH3CCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusMacCase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusNasId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusNasIdType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusNasIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusPasswordEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusPasswordRenewal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRadiusCoa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRadiusPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRequireMessageAuthenticator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoContextTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoEndpointAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoEpOneIpOnly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoFlushIpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoLogFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusRssoLogPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoRadiusResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoRadiusServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusRssoValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecondarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusSecondaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusServerIdentityCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSourceIpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusSsoAttribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSsoAttributeKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSsoAttributeValueOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusStatusTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSwitchControllerAcctFastFramedipDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSwitchControllerNasIpDynamic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSwitchControllerServiceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusTertiarySecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserRadiusTertiaryServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusTlsMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusTransportProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusUseManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusUsernameCaseSensitive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusVrfSelect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserRadius(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("account_key_cert_field"); ok || d.HasChange("account_key_cert_field") {
		t, err := expandUserRadiusAccountKeyCertField(d, v, "account_key_cert_field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-cert-field"] = t
		}
	}

	if v, ok := d.GetOk("account_key_processing"); ok || d.HasChange("account_key_processing") {
		t, err := expandUserRadiusAccountKeyProcessing(d, v, "account_key_processing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-key-processing"] = t
		}
	}

	if v, ok := d.GetOk("accounting_server"); ok || d.HasChange("accounting_server") {
		t, err := expandUserRadiusAccountingServer(d, v, "accounting_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accounting-server"] = t
		}
	}

	if v, ok := d.GetOk("acct_all_servers"); ok || d.HasChange("acct_all_servers") {
		t, err := expandUserRadiusAcctAllServers(d, v, "acct_all_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-all-servers"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok || d.HasChange("acct_interim_interval") {
		t, err := expandUserRadiusAcctInterimInterval(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("all_usergroup"); ok || d.HasChange("all_usergroup") {
		t, err := expandUserRadiusAllUsergroup(d, v, "all_usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-usergroup"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok || d.HasChange("auth_type") {
		t, err := expandUserRadiusAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok || d.HasChange("ca_cert") {
		t, err := expandUserRadiusCaCert(d, v, "ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("call_station_id_type"); ok || d.HasChange("call_station_id_type") {
		t, err := expandUserRadiusCallStationIdType(d, v, "call_station_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-station-id-type"] = t
		}
	}

	if v, ok := d.GetOk("class"); ok || d.HasChange("class") {
		t, err := expandUserRadiusClass(d, v, "class")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok || d.HasChange("client_cert") {
		t, err := expandUserRadiusClientCert(d, v, "client_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("delimiter"); ok || d.HasChange("delimiter") {
		t, err := expandUserRadiusDelimiter(d, v, "delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delimiter"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandUserRadiusDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("group_override_attr_type"); ok || d.HasChange("group_override_attr_type") {
		t, err := expandUserRadiusGroupOverrideAttrType(d, v, "group_override_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-override-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("h3c_compatibility"); ok || d.HasChange("h3c_compatibility") {
		t, err := expandUserRadiusH3CCompatibility(d, v, "h3c_compatibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3c-compatibility"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandUserRadiusInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok || d.HasChange("interface_select_method") {
		t, err := expandUserRadiusInterfaceSelectMethod(d, v, "interface_select_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok || d.HasChange("mac_case") {
		t, err := expandUserRadiusMacCase(d, v, "mac_case")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok || d.HasChange("mac_password_delimiter") {
		t, err := expandUserRadiusMacPasswordDelimiter(d, v, "mac_password_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok || d.HasChange("mac_username_delimiter") {
		t, err := expandUserRadiusMacUsernameDelimiter(d, v, "mac_username_delimiter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserRadiusName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nas_id"); ok || d.HasChange("nas_id") {
		t, err := expandUserRadiusNasId(d, v, "nas_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-id"] = t
		}
	}

	if v, ok := d.GetOk("nas_id_type"); ok || d.HasChange("nas_id_type") {
		t, err := expandUserRadiusNasIdType(d, v, "nas_id_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-id-type"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok || d.HasChange("nas_ip") {
		t, err := expandUserRadiusNasIp(d, v, "nas_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("password_encoding"); ok || d.HasChange("password_encoding") {
		t, err := expandUserRadiusPasswordEncoding(d, v, "password_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-encoding"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok || d.HasChange("password_renewal") {
		t, err := expandUserRadiusPasswordRenewal(d, v, "password_renewal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("radius_coa"); ok || d.HasChange("radius_coa") {
		t, err := expandUserRadiusRadiusCoa(d, v, "radius_coa")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-coa"] = t
		}
	}

	if v, ok := d.GetOk("radius_port"); ok || d.HasChange("radius_port") {
		t, err := expandUserRadiusRadiusPort(d, v, "radius_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("require_message_authenticator"); ok || d.HasChange("require_message_authenticator") {
		t, err := expandUserRadiusRequireMessageAuthenticator(d, v, "require_message_authenticator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-message-authenticator"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok || d.HasChange("rsso") {
		t, err := expandUserRadiusRsso(d, v, "rsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("rsso_context_timeout"); ok || d.HasChange("rsso_context_timeout") {
		t, err := expandUserRadiusRssoContextTimeout(d, v, "rsso_context_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-context-timeout"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_attribute"); ok || d.HasChange("rsso_endpoint_attribute") {
		t, err := expandUserRadiusRssoEndpointAttribute(d, v, "rsso_endpoint_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_block_attribute"); ok || d.HasChange("rsso_endpoint_block_attribute") {
		t, err := expandUserRadiusRssoEndpointBlockAttribute(d, v, "rsso_endpoint_block_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-block-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_ep_one_ip_only"); ok || d.HasChange("rsso_ep_one_ip_only") {
		t, err := expandUserRadiusRssoEpOneIpOnly(d, v, "rsso_ep_one_ip_only")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-ep-one-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("rsso_flush_ip_session"); ok || d.HasChange("rsso_flush_ip_session") {
		t, err := expandUserRadiusRssoFlushIpSession(d, v, "rsso_flush_ip_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-flush-ip-session"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_flags"); ok || d.HasChange("rsso_log_flags") {
		t, err := expandUserRadiusRssoLogFlags(d, v, "rsso_log_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-flags"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_period"); ok || d.HasChange("rsso_log_period") {
		t, err := expandUserRadiusRssoLogPeriod(d, v, "rsso_log_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-period"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_response"); ok || d.HasChange("rsso_radius_response") {
		t, err := expandUserRadiusRssoRadiusResponse(d, v, "rsso_radius_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-response"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_server_port"); ok || d.HasChange("rsso_radius_server_port") {
		t, err := expandUserRadiusRssoRadiusServerPort(d, v, "rsso_radius_server_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-server-port"] = t
		}
	}

	if v, ok := d.GetOk("rsso_secret"); ok || d.HasChange("rsso_secret") {
		t, err := expandUserRadiusRssoSecret(d, v, "rsso_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-secret"] = t
		}
	}

	if v, ok := d.GetOk("rsso_validate_request_secret"); ok || d.HasChange("rsso_validate_request_secret") {
		t, err := expandUserRadiusRssoValidateRequestSecret(d, v, "rsso_validate_request_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-validate-request-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_secret"); ok || d.HasChange("secondary_secret") {
		t, err := expandUserRadiusSecondarySecret(d, v, "secondary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok || d.HasChange("secondary_server") {
		t, err := expandUserRadiusSecondaryServer(d, v, "secondary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok || d.HasChange("secret") {
		t, err := expandUserRadiusSecret(d, v, "secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok || d.HasChange("server") {
		t, err := expandUserRadiusServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_identity_check"); ok || d.HasChange("server_identity_check") {
		t, err := expandUserRadiusServerIdentityCheck(d, v, "server_identity_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-identity-check"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok || d.HasChange("source_ip") {
		t, err := expandUserRadiusSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok || d.HasChange("source_ip_interface") {
		t, err := expandUserRadiusSourceIpInterface(d, v, "source_ip_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute"); ok || d.HasChange("sso_attribute") {
		t, err := expandUserRadiusSsoAttribute(d, v, "sso_attribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_key"); ok || d.HasChange("sso_attribute_key") {
		t, err := expandUserRadiusSsoAttributeKey(d, v, "sso_attribute_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-key"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value_override"); ok || d.HasChange("sso_attribute_value_override") {
		t, err := expandUserRadiusSsoAttributeValueOverride(d, v, "sso_attribute_value_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value-override"] = t
		}
	}

	if v, ok := d.GetOk("status_ttl"); ok || d.HasChange("status_ttl") {
		t, err := expandUserRadiusStatusTtl(d, v, "status_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status-ttl"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_acct_fast_framedip_detect"); ok || d.HasChange("switch_controller_acct_fast_framedip_detect") {
		t, err := expandUserRadiusSwitchControllerAcctFastFramedipDetect(d, v, "switch_controller_acct_fast_framedip_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-acct-fast-framedip-detect"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_nas_ip_dynamic"); ok || d.HasChange("switch_controller_nas_ip_dynamic") {
		t, err := expandUserRadiusSwitchControllerNasIpDynamic(d, v, "switch_controller_nas_ip_dynamic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-nas-ip-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_service_type"); ok || d.HasChange("switch_controller_service_type") {
		t, err := expandUserRadiusSwitchControllerServiceType(d, v, "switch_controller_service_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-service-type"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_secret"); ok || d.HasChange("tertiary_secret") {
		t, err := expandUserRadiusTertiarySecret(d, v, "tertiary_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-secret"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok || d.HasChange("tertiary_server") {
		t, err := expandUserRadiusTertiaryServer(d, v, "tertiary_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandUserRadiusTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("tls_min_proto_version"); ok || d.HasChange("tls_min_proto_version") {
		t, err := expandUserRadiusTlsMinProtoVersion(d, v, "tls_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("transport_protocol"); ok || d.HasChange("transport_protocol") {
		t, err := expandUserRadiusTransportProtocol(d, v, "transport_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport-protocol"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok || d.HasChange("use_management_vdom") {
		t, err := expandUserRadiusUseManagementVdom(d, v, "use_management_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitive"); ok || d.HasChange("username_case_sensitive") {
		t, err := expandUserRadiusUsernameCaseSensitive(d, v, "username_case_sensitive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("vrf_select"); ok || d.HasChange("vrf_select") {
		t, err := expandUserRadiusVrfSelect(d, v, "vrf_select")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	}

	return &obj, nil
}
