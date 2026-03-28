// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> SAML server entry configuration.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserSaml() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserSamlCreate,
		Read:   resourceUserSamlRead,
		Update: resourceUserSamlUpdate,
		Delete: resourceUserSamlDelete,

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
			"adfs_claim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"clock_tolerance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"digest_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"adfs_claim": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"auth_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"clock_tolerance": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"digest_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_claim_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"idp_entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"idp_single_sign_on_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"limit_relaystate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"scim_client": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"scim_group_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"scim_user_attr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"single_sign_on_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_claim_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"idp_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit_relaystate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scim_client": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"scim_group_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scim_user_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_claim_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceUserSamlCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserSaml(d)
	if err != nil {
		return fmt.Errorf("Error creating UserSaml resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadUserSaml(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateUserSaml(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating UserSaml resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateUserSaml(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating UserSaml resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceUserSamlRead(d, m)
}

func resourceUserSamlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserSaml(d)
	if err != nil {
		return fmt.Errorf("Error updating UserSaml resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateUserSaml(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating UserSaml resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceUserSamlRead(d, m)
}

func resourceUserSamlDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteUserSaml(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting UserSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSamlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserSaml(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading UserSaml resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserSaml(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserSaml resource from API: %v", err)
	}
	return nil
}

func flattenUserSamlAdfsClaim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlAuthUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSamlClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDigestMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMapping(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserSamlDynamicMappingScope(i["_scope"], d, pre_append)
			tmp["_scope"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-Scope")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adfs_claim"
		if _, ok := i["adfs-claim"]; ok {
			v := flattenUserSamlDynamicMappingAdfsClaim(i["adfs-claim"], d, pre_append)
			tmp["adfs_claim"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-AdfsClaim")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_url"
		if _, ok := i["auth-url"]; ok {
			v := flattenUserSamlDynamicMappingAuthUrl(i["auth-url"], d, pre_append)
			tmp["auth_url"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-AuthUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert"
		if _, ok := i["cert"]; ok {
			v := flattenUserSamlDynamicMappingCert(i["cert"], d, pre_append)
			tmp["cert"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-Cert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clock_tolerance"
		if _, ok := i["clock-tolerance"]; ok {
			v := flattenUserSamlDynamicMappingClockTolerance(i["clock-tolerance"], d, pre_append)
			tmp["clock_tolerance"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-ClockTolerance")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "digest_method"
		if _, ok := i["digest-method"]; ok {
			v := flattenUserSamlDynamicMappingDigestMethod(i["digest-method"], d, pre_append)
			tmp["digest_method"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-DigestMethod")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entity_id"
		if _, ok := i["entity-id"]; ok {
			v := flattenUserSamlDynamicMappingEntityId(i["entity-id"], d, pre_append)
			tmp["entity_id"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-EntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_claim_type"
		if _, ok := i["group-claim-type"]; ok {
			v := flattenUserSamlDynamicMappingGroupClaimType(i["group-claim-type"], d, pre_append)
			tmp["group_claim_type"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-GroupClaimType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := i["group-name"]; ok {
			v := flattenUserSamlDynamicMappingGroupName(i["group-name"], d, pre_append)
			tmp["group_name"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-GroupName")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := i["idp-cert"]; ok {
			v := flattenUserSamlDynamicMappingIdpCert(i["idp-cert"], d, pre_append)
			tmp["idp_cert"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-IdpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {
			v := flattenUserSamlDynamicMappingIdpEntityId(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenUserSamlDynamicMappingIdpSingleLogoutUrl(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenUserSamlDynamicMappingIdpSingleSignOnUrl(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit_relaystate"
		if _, ok := i["limit-relaystate"]; ok {
			v := flattenUserSamlDynamicMappingLimitRelaystate(i["limit-relaystate"], d, pre_append)
			tmp["limit_relaystate"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-LimitRelaystate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reauth"
		if _, ok := i["reauth"]; ok {
			v := flattenUserSamlDynamicMappingReauth(i["reauth"], d, pre_append)
			tmp["reauth"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-Reauth")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_client"
		if _, ok := i["scim-client"]; ok {
			v := flattenUserSamlDynamicMappingScimClient(i["scim-client"], d, pre_append)
			tmp["scim_client"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-ScimClient")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_group_attr_type"
		if _, ok := i["scim-group-attr-type"]; ok {
			v := flattenUserSamlDynamicMappingScimGroupAttrType(i["scim-group-attr-type"], d, pre_append)
			tmp["scim_group_attr_type"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-ScimGroupAttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_user_attr_type"
		if _, ok := i["scim-user-attr-type"]; ok {
			v := flattenUserSamlDynamicMappingScimUserAttrType(i["scim-user-attr-type"], d, pre_append)
			tmp["scim_user_attr_type"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-ScimUserAttrType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_logout_url"
		if _, ok := i["single-logout-url"]; ok {
			v := flattenUserSamlDynamicMappingSingleLogoutUrl(i["single-logout-url"], d, pre_append)
			tmp["single_logout_url"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-SingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_sign_on_url"
		if _, ok := i["single-sign-on-url"]; ok {
			v := flattenUserSamlDynamicMappingSingleSignOnUrl(i["single-sign-on-url"], d, pre_append)
			tmp["single_sign_on_url"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-SingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_claim_type"
		if _, ok := i["user-claim-type"]; ok {
			v := flattenUserSamlDynamicMappingUserClaimType(i["user-claim-type"], d, pre_append)
			tmp["user_claim_type"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-UserClaimType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := i["user-name"]; ok {
			v := flattenUserSamlDynamicMappingUserName(i["user-name"], d, pre_append)
			tmp["user_name"] = fortiAPISubPartPatch(v, "UserSaml-DynamicMapping-UserName")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserSamlDynamicMappingScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenUserSamlDynamicMappingScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "UserSamlDynamicMapping-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenUserSamlDynamicMappingScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "UserSamlDynamicMapping-Scope-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenUserSamlDynamicMappingScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingAdfsClaim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingAuthUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSamlDynamicMappingClockTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingDigestMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingGroupClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSamlDynamicMappingIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingLimitRelaystate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingScimClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSamlDynamicMappingScimGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingScimUserAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlDynamicMappingUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlGroupClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlLimitRelaystate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlScimClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenUserSamlScimGroupAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlScimUserAttrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlUserClaimType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserSamlUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserSaml(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("adfs_claim", flattenUserSamlAdfsClaim(o["adfs-claim"], d, "adfs_claim")); err != nil {
		if vv, ok := fortiAPIPatch(o["adfs-claim"], "UserSaml-AdfsClaim"); ok {
			if err = d.Set("adfs_claim", vv); err != nil {
				return fmt.Errorf("Error reading adfs_claim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adfs_claim: %v", err)
		}
	}

	if err = d.Set("auth_url", flattenUserSamlAuthUrl(o["auth-url"], d, "auth_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["auth-url"], "UserSaml-AuthUrl"); ok {
			if err = d.Set("auth_url", vv); err != nil {
				return fmt.Errorf("Error reading auth_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auth_url: %v", err)
		}
	}

	if err = d.Set("cert", flattenUserSamlCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "UserSaml-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("clock_tolerance", flattenUserSamlClockTolerance(o["clock-tolerance"], d, "clock_tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["clock-tolerance"], "UserSaml-ClockTolerance"); ok {
			if err = d.Set("clock_tolerance", vv); err != nil {
				return fmt.Errorf("Error reading clock_tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading clock_tolerance: %v", err)
		}
	}

	if err = d.Set("digest_method", flattenUserSamlDigestMethod(o["digest-method"], d, "digest_method")); err != nil {
		if vv, ok := fortiAPIPatch(o["digest-method"], "UserSaml-DigestMethod"); ok {
			if err = d.Set("digest_method", vv); err != nil {
				return fmt.Errorf("Error reading digest_method: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading digest_method: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dynamic_mapping", flattenUserSamlDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
			if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserSaml-DynamicMapping"); ok {
				if err = d.Set("dynamic_mapping", vv); err != nil {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading dynamic_mapping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dynamic_mapping"); ok {
			if err = d.Set("dynamic_mapping", flattenUserSamlDynamicMapping(o["dynamic_mapping"], d, "dynamic_mapping")); err != nil {
				if vv, ok := fortiAPIPatch(o["dynamic_mapping"], "UserSaml-DynamicMapping"); ok {
					if err = d.Set("dynamic_mapping", vv); err != nil {
						return fmt.Errorf("Error reading dynamic_mapping: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading dynamic_mapping: %v", err)
				}
			}
		}
	}

	if err = d.Set("entity_id", flattenUserSamlEntityId(o["entity-id"], d, "entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entity-id"], "UserSaml-EntityId"); ok {
			if err = d.Set("entity_id", vv); err != nil {
				return fmt.Errorf("Error reading entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("group_claim_type", flattenUserSamlGroupClaimType(o["group-claim-type"], d, "group_claim_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-claim-type"], "UserSaml-GroupClaimType"); ok {
			if err = d.Set("group_claim_type", vv); err != nil {
				return fmt.Errorf("Error reading group_claim_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_claim_type: %v", err)
		}
	}

	if err = d.Set("group_name", flattenUserSamlGroupName(o["group-name"], d, "group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-name"], "UserSaml-GroupName"); ok {
			if err = d.Set("group_name", vv); err != nil {
				return fmt.Errorf("Error reading group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenUserSamlIdpCert(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "UserSaml-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenUserSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "UserSaml-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenUserSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "UserSaml-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenUserSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "UserSaml-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("limit_relaystate", flattenUserSamlLimitRelaystate(o["limit-relaystate"], d, "limit_relaystate")); err != nil {
		if vv, ok := fortiAPIPatch(o["limit-relaystate"], "UserSaml-LimitRelaystate"); ok {
			if err = d.Set("limit_relaystate", vv); err != nil {
				return fmt.Errorf("Error reading limit_relaystate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading limit_relaystate: %v", err)
		}
	}

	if err = d.Set("name", flattenUserSamlName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "UserSaml-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reauth", flattenUserSamlReauth(o["reauth"], d, "reauth")); err != nil {
		if vv, ok := fortiAPIPatch(o["reauth"], "UserSaml-Reauth"); ok {
			if err = d.Set("reauth", vv); err != nil {
				return fmt.Errorf("Error reading reauth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("scim_client", flattenUserSamlScimClient(o["scim-client"], d, "scim_client")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-client"], "UserSaml-ScimClient"); ok {
			if err = d.Set("scim_client", vv); err != nil {
				return fmt.Errorf("Error reading scim_client: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_client: %v", err)
		}
	}

	if err = d.Set("scim_group_attr_type", flattenUserSamlScimGroupAttrType(o["scim-group-attr-type"], d, "scim_group_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-group-attr-type"], "UserSaml-ScimGroupAttrType"); ok {
			if err = d.Set("scim_group_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_group_attr_type: %v", err)
		}
	}

	if err = d.Set("scim_user_attr_type", flattenUserSamlScimUserAttrType(o["scim-user-attr-type"], d, "scim_user_attr_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["scim-user-attr-type"], "UserSaml-ScimUserAttrType"); ok {
			if err = d.Set("scim_user_attr_type", vv); err != nil {
				return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scim_user_attr_type: %v", err)
		}
	}

	if err = d.Set("single_logout_url", flattenUserSamlSingleLogoutUrl(o["single-logout-url"], d, "single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-logout-url"], "UserSaml-SingleLogoutUrl"); ok {
			if err = d.Set("single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", flattenUserSamlSingleSignOnUrl(o["single-sign-on-url"], d, "single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-sign-on-url"], "UserSaml-SingleSignOnUrl"); ok {
			if err = d.Set("single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("user_claim_type", flattenUserSamlUserClaimType(o["user-claim-type"], d, "user_claim_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-claim-type"], "UserSaml-UserClaimType"); ok {
			if err = d.Set("user_claim_type", vv); err != nil {
				return fmt.Errorf("Error reading user_claim_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_claim_type: %v", err)
		}
	}

	if err = d.Set("user_name", flattenUserSamlUserName(o["user-name"], d, "user_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["user-name"], "UserSaml-UserName"); ok {
			if err = d.Set("user_name", vv); err != nil {
				return fmt.Errorf("Error reading user_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading user_name: %v", err)
		}
	}

	return nil
}

func flattenUserSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserSamlAdfsClaim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlAuthUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSamlClockTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDigestMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMapping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			t, err := expandUserSamlDynamicMappingScope(d, i["_scope"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["_scope"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adfs_claim"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adfs-claim"], _ = expandUserSamlDynamicMappingAdfsClaim(d, i["adfs_claim"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["auth-url"], _ = expandUserSamlDynamicMappingAuthUrl(d, i["auth_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["cert"], _ = expandUserSamlDynamicMappingCert(d, i["cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "clock_tolerance"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["clock-tolerance"], _ = expandUserSamlDynamicMappingClockTolerance(d, i["clock_tolerance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "digest_method"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["digest-method"], _ = expandUserSamlDynamicMappingDigestMethod(d, i["digest_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["entity-id"], _ = expandUserSamlDynamicMappingEntityId(d, i["entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_claim_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-claim-type"], _ = expandUserSamlDynamicMappingGroupClaimType(d, i["group_claim_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["group-name"], _ = expandUserSamlDynamicMappingGroupName(d, i["group_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-cert"], _ = expandUserSamlDynamicMappingIdpCert(d, i["idp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-entity-id"], _ = expandUserSamlDynamicMappingIdpEntityId(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandUserSamlDynamicMappingIdpSingleLogoutUrl(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandUserSamlDynamicMappingIdpSingleSignOnUrl(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "limit_relaystate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["limit-relaystate"], _ = expandUserSamlDynamicMappingLimitRelaystate(d, i["limit_relaystate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reauth"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["reauth"], _ = expandUserSamlDynamicMappingReauth(d, i["reauth"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_client"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scim-client"], _ = expandUserSamlDynamicMappingScimClient(d, i["scim_client"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_group_attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scim-group-attr-type"], _ = expandUserSamlDynamicMappingScimGroupAttrType(d, i["scim_group_attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "scim_user_attr_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["scim-user-attr-type"], _ = expandUserSamlDynamicMappingScimUserAttrType(d, i["scim_user_attr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["single-logout-url"], _ = expandUserSamlDynamicMappingSingleLogoutUrl(d, i["single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["single-sign-on-url"], _ = expandUserSamlDynamicMappingSingleSignOnUrl(d, i["single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_claim_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-claim-type"], _ = expandUserSamlDynamicMappingUserClaimType(d, i["user_claim_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["user-name"], _ = expandUserSamlDynamicMappingUserName(d, i["user_name"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserSamlDynamicMappingScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandUserSamlDynamicMappingScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandUserSamlDynamicMappingScopeVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandUserSamlDynamicMappingScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingAdfsClaim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingAuthUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSamlDynamicMappingClockTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingDigestMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingGroupClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSamlDynamicMappingIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingLimitRelaystate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingScimClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSamlDynamicMappingScimGroupAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingScimUserAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingUserClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlDynamicMappingUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlGroupClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSamlIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlLimitRelaystate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlScimClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandUserSamlScimGroupAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlScimUserAttrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlUserClaimType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserSamlUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserSaml(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adfs_claim"); ok || d.HasChange("adfs_claim") {
		t, err := expandUserSamlAdfsClaim(d, v, "adfs_claim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-claim"] = t
		}
	}

	if v, ok := d.GetOk("auth_url"); ok || d.HasChange("auth_url") {
		t, err := expandUserSamlAuthUrl(d, v, "auth_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-url"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandUserSamlCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("clock_tolerance"); ok || d.HasChange("clock_tolerance") {
		t, err := expandUserSamlClockTolerance(d, v, "clock_tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clock-tolerance"] = t
		}
	}

	if v, ok := d.GetOk("digest_method"); ok || d.HasChange("digest_method") {
		t, err := expandUserSamlDigestMethod(d, v, "digest_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digest-method"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_mapping"); ok || d.HasChange("dynamic_mapping") {
		t, err := expandUserSamlDynamicMapping(d, v, "dynamic_mapping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic_mapping"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok || d.HasChange("entity_id") {
		t, err := expandUserSamlEntityId(d, v, "entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	}

	if v, ok := d.GetOk("group_claim_type"); ok || d.HasChange("group_claim_type") {
		t, err := expandUserSamlGroupClaimType(d, v, "group_claim_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok || d.HasChange("group_name") {
		t, err := expandUserSamlGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandUserSamlIdpCert(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandUserSamlIdpEntityId(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandUserSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandUserSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("limit_relaystate"); ok || d.HasChange("limit_relaystate") {
		t, err := expandUserSamlLimitRelaystate(d, v, "limit_relaystate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-relaystate"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandUserSamlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok || d.HasChange("reauth") {
		t, err := expandUserSamlReauth(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("scim_client"); ok || d.HasChange("scim_client") {
		t, err := expandUserSamlScimClient(d, v, "scim_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-client"] = t
		}
	}

	if v, ok := d.GetOk("scim_group_attr_type"); ok || d.HasChange("scim_group_attr_type") {
		t, err := expandUserSamlScimGroupAttrType(d, v, "scim_group_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-group-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("scim_user_attr_type"); ok || d.HasChange("scim_user_attr_type") {
		t, err := expandUserSamlScimUserAttrType(d, v, "scim_user_attr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scim-user-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("single_logout_url"); ok || d.HasChange("single_logout_url") {
		t, err := expandUserSamlSingleLogoutUrl(d, v, "single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok || d.HasChange("single_sign_on_url") {
		t, err := expandUserSamlSingleSignOnUrl(d, v, "single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("user_claim_type"); ok || d.HasChange("user_claim_type") {
		t, err := expandUserSamlUserClaimType(d, v, "user_claim_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-claim-type"] = t
		}
	}

	if v, ok := d.GetOk("user_name"); ok || d.HasChange("user_name") {
		t, err := expandUserSamlUserName(d, v, "user_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-name"] = t
		}
	}

	return &obj, nil
}
