// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Global settings for SAML authentication.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSaml() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSamlUpdate,
		Read:   resourceSystemSamlRead,
		Update: resourceSystemSamlUpdate,
		Delete: resourceSystemSamlDelete,

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
			"artifact_resolution_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"binding_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"default_login_page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"idp_artifact_resolution_url": &schema.Schema{
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
			"life": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"portal_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_providers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assertion_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"idp_artifact_resolution_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_artifact_resolution_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_binding_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_cert": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
						"sp_entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_portal_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_single_logout_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_single_sign_on_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"single_logout_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"single_sign_on_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tolerance": &schema.Schema{
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

func resourceSystemSamlUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSaml(d, false)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSaml(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSaml")

	return resourceSystemSamlRead(d, m)
}

func resourceSystemSamlDelete(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSaml(d, true)

	if err != nil {
		return fmt.Errorf("Error updating SystemSaml resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSaml(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSaml resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSamlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSaml(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSaml resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSaml(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSaml resource from API: %v", err)
	}
	return nil
}

func flattenSystemSamlArtifactResolutionUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlBindingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSamlDefaultLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlDefaultProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSamlEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpArtifactResolutionUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSamlIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlPortalUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProviders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assertion_attributes"
		if _, ok := i["assertion-attributes"]; ok {
			v := flattenSystemSamlServiceProvidersAssertionAttributes(i["assertion-attributes"], d, pre_append)
			tmp["assertion_attributes"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-AssertionAttributes")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_artifact_resolution_url"
		if _, ok := i["idp-artifact-resolution-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpArtifactResolutionUrl(i["idp-artifact-resolution-url"], d, pre_append)
			tmp["idp_artifact_resolution_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpArtifactResolutionUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := i["idp-entity-id"]; ok {
			v := flattenSystemSamlServiceProvidersIdpEntityId(i["idp-entity-id"], d, pre_append)
			tmp["idp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := i["idp-single-logout-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(i["idp-single-logout-url"], d, pre_append)
			tmp["idp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := i["idp-single-sign-on-url"]; ok {
			v := flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(i["idp-single-sign-on-url"], d, pre_append)
			tmp["idp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-IdpSingleSignOnUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSystemSamlServiceProvidersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			v := flattenSystemSamlServiceProvidersPrefix(i["prefix"], d, pre_append)
			tmp["prefix"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-Prefix")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_artifact_resolution_url"
		if _, ok := i["sp-artifact-resolution-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpArtifactResolutionUrl(i["sp-artifact-resolution-url"], d, pre_append)
			tmp["sp_artifact_resolution_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpArtifactResolutionUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_binding_protocol"
		if _, ok := i["sp-binding-protocol"]; ok {
			v := flattenSystemSamlServiceProvidersSpBindingProtocol(i["sp-binding-protocol"], d, pre_append)
			tmp["sp_binding_protocol"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpBindingProtocol")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := i["sp-cert"]; ok {
			v := flattenSystemSamlServiceProvidersSpCert(i["sp-cert"], d, pre_append)
			tmp["sp_cert"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpCert")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := i["sp-entity-id"]; ok {
			v := flattenSystemSamlServiceProvidersSpEntityId(i["sp-entity-id"], d, pre_append)
			tmp["sp_entity_id"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpEntityId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_portal_url"
		if _, ok := i["sp-portal-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpPortalUrl(i["sp-portal-url"], d, pre_append)
			tmp["sp_portal_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpPortalUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := i["sp-single-logout-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpSingleLogoutUrl(i["sp-single-logout-url"], d, pre_append)
			tmp["sp_single_logout_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpSingleLogoutUrl")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := i["sp-single-sign-on-url"]; ok {
			v := flattenSystemSamlServiceProvidersSpSingleSignOnUrl(i["sp-single-sign-on-url"], d, pre_append)
			tmp["sp_single_sign_on_url"] = fortiAPISubPartPatch(v, "SystemSaml-ServiceProviders-SpSingleSignOnUrl")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSamlServiceProvidersAssertionAttributes(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenSystemSamlServiceProvidersAssertionAttributesName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SystemSamlServiceProviders-AssertionAttributes-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenSystemSamlServiceProvidersAssertionAttributesType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "SystemSamlServiceProviders-AssertionAttributes-Type")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemSamlServiceProvidersAssertionAttributesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersAssertionAttributesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpArtifactResolutionUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersIdpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpArtifactResolutionUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpBindingProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSamlServiceProvidersSpEntityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpPortalUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlServiceProvidersSpSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlSingleLogoutUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlSingleSignOnUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSamlTolerance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSaml(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("artifact_resolution_url", flattenSystemSamlArtifactResolutionUrl(o["artifact-resolution-url"], d, "artifact_resolution_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["artifact-resolution-url"], "SystemSaml-ArtifactResolutionUrl"); ok {
			if err = d.Set("artifact_resolution_url", vv); err != nil {
				return fmt.Errorf("Error reading artifact_resolution_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading artifact_resolution_url: %v", err)
		}
	}

	if err = d.Set("binding_protocol", flattenSystemSamlBindingProtocol(o["binding-protocol"], d, "binding_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["binding-protocol"], "SystemSaml-BindingProtocol"); ok {
			if err = d.Set("binding_protocol", vv); err != nil {
				return fmt.Errorf("Error reading binding_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading binding_protocol: %v", err)
		}
	}

	if err = d.Set("cert", flattenSystemSamlCert(o["cert"], d, "cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["cert"], "SystemSaml-Cert"); ok {
			if err = d.Set("cert", vv); err != nil {
				return fmt.Errorf("Error reading cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("default_login_page", flattenSystemSamlDefaultLoginPage(o["default-login-page"], d, "default_login_page")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-login-page"], "SystemSaml-DefaultLoginPage"); ok {
			if err = d.Set("default_login_page", vv); err != nil {
				return fmt.Errorf("Error reading default_login_page: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_login_page: %v", err)
		}
	}

	if err = d.Set("default_profile", flattenSystemSamlDefaultProfile(o["default-profile"], d, "default_profile")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-profile"], "SystemSaml-DefaultProfile"); ok {
			if err = d.Set("default_profile", vv); err != nil {
				return fmt.Errorf("Error reading default_profile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_profile: %v", err)
		}
	}

	if err = d.Set("entity_id", flattenSystemSamlEntityId(o["entity-id"], d, "entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["entity-id"], "SystemSaml-EntityId"); ok {
			if err = d.Set("entity_id", vv); err != nil {
				return fmt.Errorf("Error reading entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading entity_id: %v", err)
		}
	}

	if err = d.Set("idp_artifact_resolution_url", flattenSystemSamlIdpArtifactResolutionUrl(o["idp-artifact-resolution-url"], d, "idp_artifact_resolution_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-artifact-resolution-url"], "SystemSaml-IdpArtifactResolutionUrl"); ok {
			if err = d.Set("idp_artifact_resolution_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_artifact_resolution_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_artifact_resolution_url: %v", err)
		}
	}

	if err = d.Set("idp_cert", flattenSystemSamlIdpCert(o["idp-cert"], d, "idp_cert")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-cert"], "SystemSaml-IdpCert"); ok {
			if err = d.Set("idp_cert", vv); err != nil {
				return fmt.Errorf("Error reading idp_cert: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_cert: %v", err)
		}
	}

	if err = d.Set("idp_entity_id", flattenSystemSamlIdpEntityId(o["idp-entity-id"], d, "idp_entity_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-entity-id"], "SystemSaml-IdpEntityId"); ok {
			if err = d.Set("idp_entity_id", vv); err != nil {
				return fmt.Errorf("Error reading idp_entity_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_entity_id: %v", err)
		}
	}

	if err = d.Set("idp_single_logout_url", flattenSystemSamlIdpSingleLogoutUrl(o["idp-single-logout-url"], d, "idp_single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-logout-url"], "SystemSaml-IdpSingleLogoutUrl"); ok {
			if err = d.Set("idp_single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_logout_url: %v", err)
		}
	}

	if err = d.Set("idp_single_sign_on_url", flattenSystemSamlIdpSingleSignOnUrl(o["idp-single-sign-on-url"], d, "idp_single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["idp-single-sign-on-url"], "SystemSaml-IdpSingleSignOnUrl"); ok {
			if err = d.Set("idp_single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading idp_single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("life", flattenSystemSamlLife(o["life"], d, "life")); err != nil {
		if vv, ok := fortiAPIPatch(o["life"], "SystemSaml-Life"); ok {
			if err = d.Set("life", vv); err != nil {
				return fmt.Errorf("Error reading life: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading life: %v", err)
		}
	}

	if err = d.Set("portal_url", flattenSystemSamlPortalUrl(o["portal-url"], d, "portal_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["portal-url"], "SystemSaml-PortalUrl"); ok {
			if err = d.Set("portal_url", vv); err != nil {
				return fmt.Errorf("Error reading portal_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading portal_url: %v", err)
		}
	}

	if err = d.Set("role", flattenSystemSamlRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "SystemSaml-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("server_address", flattenSystemSamlServerAddress(o["server-address"], d, "server_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["server-address"], "SystemSaml-ServerAddress"); ok {
			if err = d.Set("server_address", vv); err != nil {
				return fmt.Errorf("Error reading server_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading server_address: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_providers", flattenSystemSamlServiceProviders(o["service-providers"], d, "service_providers")); err != nil {
			if vv, ok := fortiAPIPatch(o["service-providers"], "SystemSaml-ServiceProviders"); ok {
				if err = d.Set("service_providers", vv); err != nil {
					return fmt.Errorf("Error reading service_providers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading service_providers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service_providers"); ok {
			if err = d.Set("service_providers", flattenSystemSamlServiceProviders(o["service-providers"], d, "service_providers")); err != nil {
				if vv, ok := fortiAPIPatch(o["service-providers"], "SystemSaml-ServiceProviders"); ok {
					if err = d.Set("service_providers", vv); err != nil {
						return fmt.Errorf("Error reading service_providers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading service_providers: %v", err)
				}
			}
		}
	}

	if err = d.Set("single_logout_url", flattenSystemSamlSingleLogoutUrl(o["single-logout-url"], d, "single_logout_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-logout-url"], "SystemSaml-SingleLogoutUrl"); ok {
			if err = d.Set("single_logout_url", vv); err != nil {
				return fmt.Errorf("Error reading single_logout_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_logout_url: %v", err)
		}
	}

	if err = d.Set("single_sign_on_url", flattenSystemSamlSingleSignOnUrl(o["single-sign-on-url"], d, "single_sign_on_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["single-sign-on-url"], "SystemSaml-SingleSignOnUrl"); ok {
			if err = d.Set("single_sign_on_url", vv); err != nil {
				return fmt.Errorf("Error reading single_sign_on_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading single_sign_on_url: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSamlStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "SystemSaml-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tolerance", flattenSystemSamlTolerance(o["tolerance"], d, "tolerance")); err != nil {
		if vv, ok := fortiAPIPatch(o["tolerance"], "SystemSaml-Tolerance"); ok {
			if err = d.Set("tolerance", vv); err != nil {
				return fmt.Errorf("Error reading tolerance: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tolerance: %v", err)
		}
	}

	return nil
}

func flattenSystemSamlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSamlArtifactResolutionUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlBindingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSamlDefaultLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlDefaultProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSamlEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpArtifactResolutionUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSamlIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlPortalUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProviders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assertion_attributes"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			t, err := expandSystemSamlServiceProvidersAssertionAttributes(d, i["assertion_attributes"], pre_append)
			if err != nil {
				return result, err
			} else if t != nil {
				tmp["assertion-attributes"] = t
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_artifact_resolution_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-artifact-resolution-url"], _ = expandSystemSamlServiceProvidersIdpArtifactResolutionUrl(d, i["idp_artifact_resolution_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-entity-id"], _ = expandSystemSamlServiceProvidersIdpEntityId(d, i["idp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-logout-url"], _ = expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d, i["idp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "idp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["idp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d, i["idp_single_sign_on_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSystemSamlServiceProvidersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["prefix"], _ = expandSystemSamlServiceProvidersPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_artifact_resolution_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-artifact-resolution-url"], _ = expandSystemSamlServiceProvidersSpArtifactResolutionUrl(d, i["sp_artifact_resolution_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_binding_protocol"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-binding-protocol"], _ = expandSystemSamlServiceProvidersSpBindingProtocol(d, i["sp_binding_protocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_cert"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-cert"], _ = expandSystemSamlServiceProvidersSpCert(d, i["sp_cert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_entity_id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-entity-id"], _ = expandSystemSamlServiceProvidersSpEntityId(d, i["sp_entity_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_portal_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-portal-url"], _ = expandSystemSamlServiceProvidersSpPortalUrl(d, i["sp_portal_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_logout_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-single-logout-url"], _ = expandSystemSamlServiceProvidersSpSingleLogoutUrl(d, i["sp_single_logout_url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sp_single_sign_on_url"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sp-single-sign-on-url"], _ = expandSystemSamlServiceProvidersSpSingleSignOnUrl(d, i["sp_single_sign_on_url"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSamlServiceProvidersAssertionAttributes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSamlServiceProvidersAssertionAttributesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["type"], _ = expandSystemSamlServiceProvidersAssertionAttributesType(d, i["type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemSamlServiceProvidersAssertionAttributesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersAssertionAttributesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpArtifactResolutionUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersIdpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpArtifactResolutionUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpBindingProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSamlServiceProvidersSpEntityId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpPortalUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlServiceProvidersSpSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlSingleLogoutUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlSingleSignOnUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSamlTolerance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSaml(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("artifact_resolution_url"); ok || d.HasChange("artifact_resolution_url") {
		t, err := expandSystemSamlArtifactResolutionUrl(d, v, "artifact_resolution_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["artifact-resolution-url"] = t
		}
	}

	if v, ok := d.GetOk("binding_protocol"); ok || d.HasChange("binding_protocol") {
		t, err := expandSystemSamlBindingProtocol(d, v, "binding_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["binding-protocol"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok || d.HasChange("cert") {
		t, err := expandSystemSamlCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("default_login_page"); ok || d.HasChange("default_login_page") {
		t, err := expandSystemSamlDefaultLoginPage(d, v, "default_login_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-login-page"] = t
		}
	}

	if v, ok := d.GetOk("default_profile"); ok || d.HasChange("default_profile") {
		t, err := expandSystemSamlDefaultProfile(d, v, "default_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-profile"] = t
		}
	}

	if v, ok := d.GetOk("entity_id"); ok || d.HasChange("entity_id") {
		t, err := expandSystemSamlEntityId(d, v, "entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_artifact_resolution_url"); ok || d.HasChange("idp_artifact_resolution_url") {
		t, err := expandSystemSamlIdpArtifactResolutionUrl(d, v, "idp_artifact_resolution_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-artifact-resolution-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_cert"); ok || d.HasChange("idp_cert") {
		t, err := expandSystemSamlIdpCert(d, v, "idp_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-cert"] = t
		}
	}

	if v, ok := d.GetOk("idp_entity_id"); ok || d.HasChange("idp_entity_id") {
		t, err := expandSystemSamlIdpEntityId(d, v, "idp_entity_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-entity-id"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_logout_url"); ok || d.HasChange("idp_single_logout_url") {
		t, err := expandSystemSamlIdpSingleLogoutUrl(d, v, "idp_single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("idp_single_sign_on_url"); ok || d.HasChange("idp_single_sign_on_url") {
		t, err := expandSystemSamlIdpSingleSignOnUrl(d, v, "idp_single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idp-single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("life"); ok || d.HasChange("life") {
		t, err := expandSystemSamlLife(d, v, "life")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["life"] = t
		}
	}

	if v, ok := d.GetOk("portal_url"); ok || d.HasChange("portal_url") {
		t, err := expandSystemSamlPortalUrl(d, v, "portal_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-url"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandSystemSamlRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("server_address"); ok || d.HasChange("server_address") {
		t, err := expandSystemSamlServerAddress(d, v, "server_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-address"] = t
		}
	}

	if bemptysontable {
		obj["service-providers"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("service_providers"); ok || d.HasChange("service_providers") {
			t, err := expandSystemSamlServiceProviders(d, v, "service_providers")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service-providers"] = t
			}
		}
	}

	if v, ok := d.GetOk("single_logout_url"); ok || d.HasChange("single_logout_url") {
		t, err := expandSystemSamlSingleLogoutUrl(d, v, "single_logout_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-logout-url"] = t
		}
	}

	if v, ok := d.GetOk("single_sign_on_url"); ok || d.HasChange("single_sign_on_url") {
		t, err := expandSystemSamlSingleSignOnUrl(d, v, "single_sign_on_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["single-sign-on-url"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandSystemSamlStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tolerance"); ok || d.HasChange("tolerance") {
		t, err := expandSystemSamlTolerance(d, v, "tolerance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tolerance"] = t
		}
	}

	return &obj, nil
}
