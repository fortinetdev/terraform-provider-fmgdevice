// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure admin users.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAdminCreate,
		Read:   resourceSystemAdminRead,
		Update: resourceSystemAdminUpdate,
		Delete: resourceSystemAdminDelete,

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
			"accprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_remove_admin_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fortitoken": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"guest_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_lang": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"guest_usergroups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hidden": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gui_default_dashboard_template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_ignore_invalid_signature_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gui_ignore_release_overview_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"history0": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"history1": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ip6_trusthost1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_trusthost9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"password_expire": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"peer_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"radius_vdom_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_group": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"sms_phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_public_key1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_public_key2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_public_key3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"trusthost1": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost10": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost2": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost3": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost4": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost5": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost6": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost7": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost8": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"trusthost9": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"vdom_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAdmin(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdmin(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	obj, err := getObjectSystemAdmin(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdmin(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	cfg := m.(*FortiClient).Cfg
	device_name, err := getVariable(cfg, d, "device_name")
	if err != nil {
		return err
	}
	paradict["device"] = device_name

	err = c.DeleteSystemAdmin(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAdmin(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdmin(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemAdminAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminAccprofileOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminAllowRemoveAdminSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminEmailTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminForcePasswordChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminFortitoken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminGuestAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuestLang(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminGuestUsergroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminHidden(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiDefaultDashboardTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiIgnoreInvalidSignatureVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminGuiIgnoreReleaseOverviewVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminIp6Trusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminPasswordExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminPeerAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminPeerGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminRadiusVdomOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRemoteAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminRemoteGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSmsCustomServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminSmsPhone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSmsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminSshPublicKey1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminSshPublicKey3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTrusthost1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost7(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost8(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTrusthost9(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemAdminVdomOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAdminWildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAdmin(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("accprofile", flattenSystemAdminAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile"], "SystemAdmin-Accprofile"); ok {
			if err = d.Set("accprofile", vv); err != nil {
				return fmt.Errorf("Error reading accprofile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("accprofile_override", flattenSystemAdminAccprofileOverride(o["accprofile-override"], d, "accprofile_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["accprofile-override"], "SystemAdmin-AccprofileOverride"); ok {
			if err = d.Set("accprofile_override", vv); err != nil {
				return fmt.Errorf("Error reading accprofile_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accprofile_override: %v", err)
		}
	}

	if err = d.Set("allow_remove_admin_session", flattenSystemAdminAllowRemoveAdminSession(o["allow-remove-admin-session"], d, "allow_remove_admin_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["allow-remove-admin-session"], "SystemAdmin-AllowRemoveAdminSession"); ok {
			if err = d.Set("allow_remove_admin_session", vv); err != nil {
				return fmt.Errorf("Error reading allow_remove_admin_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading allow_remove_admin_session: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemAdminComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "SystemAdmin-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("email_to", flattenSystemAdminEmailTo(o["email-to"], d, "email_to")); err != nil {
		if vv, ok := fortiAPIPatch(o["email-to"], "SystemAdmin-EmailTo"); ok {
			if err = d.Set("email_to", vv); err != nil {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("force_password_change", flattenSystemAdminForcePasswordChange(o["force-password-change"], d, "force_password_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["force-password-change"], "SystemAdmin-ForcePasswordChange"); ok {
			if err = d.Set("force_password_change", vv); err != nil {
				return fmt.Errorf("Error reading force_password_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenSystemAdminFortitoken(o["fortitoken"], d, "fortitoken")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortitoken"], "SystemAdmin-Fortitoken"); ok {
			if err = d.Set("fortitoken", vv); err != nil {
				return fmt.Errorf("Error reading fortitoken: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("guest_auth", flattenSystemAdminGuestAuth(o["guest-auth"], d, "guest_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-auth"], "SystemAdmin-GuestAuth"); ok {
			if err = d.Set("guest_auth", vv); err != nil {
				return fmt.Errorf("Error reading guest_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_auth: %v", err)
		}
	}

	if err = d.Set("guest_lang", flattenSystemAdminGuestLang(o["guest-lang"], d, "guest_lang")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-lang"], "SystemAdmin-GuestLang"); ok {
			if err = d.Set("guest_lang", vv); err != nil {
				return fmt.Errorf("Error reading guest_lang: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_lang: %v", err)
		}
	}

	if err = d.Set("guest_usergroups", flattenSystemAdminGuestUsergroups(o["guest-usergroups"], d, "guest_usergroups")); err != nil {
		if vv, ok := fortiAPIPatch(o["guest-usergroups"], "SystemAdmin-GuestUsergroups"); ok {
			if err = d.Set("guest_usergroups", vv); err != nil {
				return fmt.Errorf("Error reading guest_usergroups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading guest_usergroups: %v", err)
		}
	}

	if err = d.Set("hidden", flattenSystemAdminHidden(o["hidden"], d, "hidden")); err != nil {
		if vv, ok := fortiAPIPatch(o["hidden"], "SystemAdmin-Hidden"); ok {
			if err = d.Set("hidden", vv); err != nil {
				return fmt.Errorf("Error reading hidden: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}

	if err = d.Set("gui_default_dashboard_template", flattenSystemAdminGuiDefaultDashboardTemplate(o["gui-default-dashboard-template"], d, "gui_default_dashboard_template")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-default-dashboard-template"], "SystemAdmin-GuiDefaultDashboardTemplate"); ok {
			if err = d.Set("gui_default_dashboard_template", vv); err != nil {
				return fmt.Errorf("Error reading gui_default_dashboard_template: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_default_dashboard_template: %v", err)
		}
	}

	if err = d.Set("gui_ignore_invalid_signature_version", flattenSystemAdminGuiIgnoreInvalidSignatureVersion(o["gui-ignore-invalid-signature-version"], d, "gui_ignore_invalid_signature_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ignore-invalid-signature-version"], "SystemAdmin-GuiIgnoreInvalidSignatureVersion"); ok {
			if err = d.Set("gui_ignore_invalid_signature_version", vv); err != nil {
				return fmt.Errorf("Error reading gui_ignore_invalid_signature_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ignore_invalid_signature_version: %v", err)
		}
	}

	if err = d.Set("gui_ignore_release_overview_version", flattenSystemAdminGuiIgnoreReleaseOverviewVersion(o["gui-ignore-release-overview-version"], d, "gui_ignore_release_overview_version")); err != nil {
		if vv, ok := fortiAPIPatch(o["gui-ignore-release-overview-version"], "SystemAdmin-GuiIgnoreReleaseOverviewVersion"); ok {
			if err = d.Set("gui_ignore_release_overview_version", vv); err != nil {
				return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost1", flattenSystemAdminIp6Trusthost1(o["ip6-trusthost1"], d, "ip6_trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost1"], "SystemAdmin-Ip6Trusthost1"); ok {
			if err = d.Set("ip6_trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost1: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost10", flattenSystemAdminIp6Trusthost10(o["ip6-trusthost10"], d, "ip6_trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost10"], "SystemAdmin-Ip6Trusthost10"); ok {
			if err = d.Set("ip6_trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost10: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost2", flattenSystemAdminIp6Trusthost2(o["ip6-trusthost2"], d, "ip6_trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost2"], "SystemAdmin-Ip6Trusthost2"); ok {
			if err = d.Set("ip6_trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost2: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost3", flattenSystemAdminIp6Trusthost3(o["ip6-trusthost3"], d, "ip6_trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost3"], "SystemAdmin-Ip6Trusthost3"); ok {
			if err = d.Set("ip6_trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost3: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost4", flattenSystemAdminIp6Trusthost4(o["ip6-trusthost4"], d, "ip6_trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost4"], "SystemAdmin-Ip6Trusthost4"); ok {
			if err = d.Set("ip6_trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost4: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost5", flattenSystemAdminIp6Trusthost5(o["ip6-trusthost5"], d, "ip6_trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost5"], "SystemAdmin-Ip6Trusthost5"); ok {
			if err = d.Set("ip6_trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost5: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost6", flattenSystemAdminIp6Trusthost6(o["ip6-trusthost6"], d, "ip6_trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost6"], "SystemAdmin-Ip6Trusthost6"); ok {
			if err = d.Set("ip6_trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost6: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost7", flattenSystemAdminIp6Trusthost7(o["ip6-trusthost7"], d, "ip6_trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost7"], "SystemAdmin-Ip6Trusthost7"); ok {
			if err = d.Set("ip6_trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost7: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost8", flattenSystemAdminIp6Trusthost8(o["ip6-trusthost8"], d, "ip6_trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost8"], "SystemAdmin-Ip6Trusthost8"); ok {
			if err = d.Set("ip6_trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost8: %v", err)
		}
	}

	if err = d.Set("ip6_trusthost9", flattenSystemAdminIp6Trusthost9(o["ip6-trusthost9"], d, "ip6_trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip6-trusthost9"], "SystemAdmin-Ip6Trusthost9"); ok {
			if err = d.Set("ip6_trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip6_trusthost9: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemAdminName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemAdmin-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("password_expire", flattenSystemAdminPasswordExpire(o["password-expire"], d, "password_expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["password-expire"], "SystemAdmin-PasswordExpire"); ok {
			if err = d.Set("password_expire", vv); err != nil {
				return fmt.Errorf("Error reading password_expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading password_expire: %v", err)
		}
	}

	if err = d.Set("peer_auth", flattenSystemAdminPeerAuth(o["peer-auth"], d, "peer_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-auth"], "SystemAdmin-PeerAuth"); ok {
			if err = d.Set("peer_auth", vv); err != nil {
				return fmt.Errorf("Error reading peer_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_auth: %v", err)
		}
	}

	if err = d.Set("peer_group", flattenSystemAdminPeerGroup(o["peer-group"], d, "peer_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["peer-group"], "SystemAdmin-PeerGroup"); ok {
			if err = d.Set("peer_group", vv); err != nil {
				return fmt.Errorf("Error reading peer_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading peer_group: %v", err)
		}
	}

	if err = d.Set("radius_vdom_override", flattenSystemAdminRadiusVdomOverride(o["radius-vdom-override"], d, "radius_vdom_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["radius-vdom-override"], "SystemAdmin-RadiusVdomOverride"); ok {
			if err = d.Set("radius_vdom_override", vv); err != nil {
				return fmt.Errorf("Error reading radius_vdom_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading radius_vdom_override: %v", err)
		}
	}

	if err = d.Set("remote_auth", flattenSystemAdminRemoteAuth(o["remote-auth"], d, "remote_auth")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-auth"], "SystemAdmin-RemoteAuth"); ok {
			if err = d.Set("remote_auth", vv); err != nil {
				return fmt.Errorf("Error reading remote_auth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_auth: %v", err)
		}
	}

	if err = d.Set("remote_group", flattenSystemAdminRemoteGroup(o["remote-group"], d, "remote_group")); err != nil {
		if vv, ok := fortiAPIPatch(o["remote-group"], "SystemAdmin-RemoteGroup"); ok {
			if err = d.Set("remote_group", vv); err != nil {
				return fmt.Errorf("Error reading remote_group: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading remote_group: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemAdminSchedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "SystemAdmin-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenSystemAdminSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-custom-server"], "SystemAdmin-SmsCustomServer"); ok {
			if err = d.Set("sms_custom_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_custom_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenSystemAdminSmsPhone(o["sms-phone"], d, "sms_phone")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-phone"], "SystemAdmin-SmsPhone"); ok {
			if err = d.Set("sms_phone", vv); err != nil {
				return fmt.Errorf("Error reading sms_phone: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenSystemAdminSmsServer(o["sms-server"], d, "sms_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["sms-server"], "SystemAdmin-SmsServer"); ok {
			if err = d.Set("sms_server", vv); err != nil {
				return fmt.Errorf("Error reading sms_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("ssh_certificate", flattenSystemAdminSshCertificate(o["ssh-certificate"], d, "ssh_certificate")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-certificate"], "SystemAdmin-SshCertificate"); ok {
			if err = d.Set("ssh_certificate", vv); err != nil {
				return fmt.Errorf("Error reading ssh_certificate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_certificate: %v", err)
		}
	}

	if err = d.Set("ssh_public_key1", flattenSystemAdminSshPublicKey1(o["ssh-public-key1"], d, "ssh_public_key1")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-public-key1"], "SystemAdmin-SshPublicKey1"); ok {
			if err = d.Set("ssh_public_key1", vv); err != nil {
				return fmt.Errorf("Error reading ssh_public_key1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_public_key1: %v", err)
		}
	}

	if err = d.Set("ssh_public_key2", flattenSystemAdminSshPublicKey2(o["ssh-public-key2"], d, "ssh_public_key2")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-public-key2"], "SystemAdmin-SshPublicKey2"); ok {
			if err = d.Set("ssh_public_key2", vv); err != nil {
				return fmt.Errorf("Error reading ssh_public_key2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_public_key2: %v", err)
		}
	}

	if err = d.Set("ssh_public_key3", flattenSystemAdminSshPublicKey3(o["ssh-public-key3"], d, "ssh_public_key3")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-public-key3"], "SystemAdmin-SshPublicKey3"); ok {
			if err = d.Set("ssh_public_key3", vv); err != nil {
				return fmt.Errorf("Error reading ssh_public_key3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_public_key3: %v", err)
		}
	}

	if err = d.Set("trusthost1", flattenSystemAdminTrusthost1(o["trusthost1"], d, "trusthost1")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost1"], "SystemAdmin-Trusthost1"); ok {
			if err = d.Set("trusthost1", vv); err != nil {
				return fmt.Errorf("Error reading trusthost1: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost1: %v", err)
		}
	}

	if err = d.Set("trusthost10", flattenSystemAdminTrusthost10(o["trusthost10"], d, "trusthost10")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost10"], "SystemAdmin-Trusthost10"); ok {
			if err = d.Set("trusthost10", vv); err != nil {
				return fmt.Errorf("Error reading trusthost10: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost10: %v", err)
		}
	}

	if err = d.Set("trusthost2", flattenSystemAdminTrusthost2(o["trusthost2"], d, "trusthost2")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost2"], "SystemAdmin-Trusthost2"); ok {
			if err = d.Set("trusthost2", vv); err != nil {
				return fmt.Errorf("Error reading trusthost2: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost2: %v", err)
		}
	}

	if err = d.Set("trusthost3", flattenSystemAdminTrusthost3(o["trusthost3"], d, "trusthost3")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost3"], "SystemAdmin-Trusthost3"); ok {
			if err = d.Set("trusthost3", vv); err != nil {
				return fmt.Errorf("Error reading trusthost3: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost3: %v", err)
		}
	}

	if err = d.Set("trusthost4", flattenSystemAdminTrusthost4(o["trusthost4"], d, "trusthost4")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost4"], "SystemAdmin-Trusthost4"); ok {
			if err = d.Set("trusthost4", vv); err != nil {
				return fmt.Errorf("Error reading trusthost4: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost4: %v", err)
		}
	}

	if err = d.Set("trusthost5", flattenSystemAdminTrusthost5(o["trusthost5"], d, "trusthost5")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost5"], "SystemAdmin-Trusthost5"); ok {
			if err = d.Set("trusthost5", vv); err != nil {
				return fmt.Errorf("Error reading trusthost5: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost5: %v", err)
		}
	}

	if err = d.Set("trusthost6", flattenSystemAdminTrusthost6(o["trusthost6"], d, "trusthost6")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost6"], "SystemAdmin-Trusthost6"); ok {
			if err = d.Set("trusthost6", vv); err != nil {
				return fmt.Errorf("Error reading trusthost6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost6: %v", err)
		}
	}

	if err = d.Set("trusthost7", flattenSystemAdminTrusthost7(o["trusthost7"], d, "trusthost7")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost7"], "SystemAdmin-Trusthost7"); ok {
			if err = d.Set("trusthost7", vv); err != nil {
				return fmt.Errorf("Error reading trusthost7: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost7: %v", err)
		}
	}

	if err = d.Set("trusthost8", flattenSystemAdminTrusthost8(o["trusthost8"], d, "trusthost8")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost8"], "SystemAdmin-Trusthost8"); ok {
			if err = d.Set("trusthost8", vv); err != nil {
				return fmt.Errorf("Error reading trusthost8: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost8: %v", err)
		}
	}

	if err = d.Set("trusthost9", flattenSystemAdminTrusthost9(o["trusthost9"], d, "trusthost9")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusthost9"], "SystemAdmin-Trusthost9"); ok {
			if err = d.Set("trusthost9", vv); err != nil {
				return fmt.Errorf("Error reading trusthost9: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusthost9: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenSystemAdminTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor"], "SystemAdmin-TwoFactor"); ok {
			if err = d.Set("two_factor", vv); err != nil {
				return fmt.Errorf("Error reading two_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenSystemAdminTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-authentication"], "SystemAdmin-TwoFactorAuthentication"); ok {
			if err = d.Set("two_factor_authentication", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_authentication: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenSystemAdminTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification")); err != nil {
		if vv, ok := fortiAPIPatch(o["two-factor-notification"], "SystemAdmin-TwoFactorNotification"); ok {
			if err = d.Set("two_factor_notification", vv); err != nil {
				return fmt.Errorf("Error reading two_factor_notification: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemAdminVdom(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemAdmin-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vdom_override", flattenSystemAdminVdomOverride(o["vdom-override"], d, "vdom_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom-override"], "SystemAdmin-VdomOverride"); ok {
			if err = d.Set("vdom_override", vv); err != nil {
				return fmt.Errorf("Error reading vdom_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_override: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenSystemAdminWildcard(o["wildcard"], d, "wildcard")); err != nil {
		if vv, ok := fortiAPIPatch(o["wildcard"], "SystemAdmin-Wildcard"); ok {
			if err = d.Set("wildcard", vv); err != nil {
				return fmt.Errorf("Error reading wildcard: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	return nil
}

func flattenSystemAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAdminAccprofile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminAccprofileOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminAllowRemoveAdminSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminForcePasswordChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminFortitoken(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminGuestAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuestLang(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminGuestUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminHidden(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiDefaultDashboardTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiIgnoreInvalidSignatureVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminGuiIgnoreReleaseOverviewVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminHistory0(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminHistory1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminIp6Trusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminIp6Trusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminPasswordExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminPeerAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminPeerGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminRadiusVdomOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRemoteAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminRemoteGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSmsCustomServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminSmsPhone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSmsServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminSshPublicKey1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshPublicKey2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminSshPublicKey3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTrusthost1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost10(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost7(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost8(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTrusthost9(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandSystemAdminTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemAdminVdomOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAdminWildcard(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdmin(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accprofile"); ok || d.HasChange("accprofile") {
		t, err := expandSystemAdminAccprofile(d, v, "accprofile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("accprofile_override"); ok || d.HasChange("accprofile_override") {
		t, err := expandSystemAdminAccprofileOverride(d, v, "accprofile_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile-override"] = t
		}
	}

	if v, ok := d.GetOk("allow_remove_admin_session"); ok || d.HasChange("allow_remove_admin_session") {
		t, err := expandSystemAdminAllowRemoveAdminSession(d, v, "allow_remove_admin_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remove-admin-session"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandSystemAdminComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandSystemAdminEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("force_password_change"); ok || d.HasChange("force_password_change") {
		t, err := expandSystemAdminForcePasswordChange(d, v, "force_password_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-password-change"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok || d.HasChange("fortitoken") {
		t, err := expandSystemAdminFortitoken(d, v, "fortitoken")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("guest_auth"); ok || d.HasChange("guest_auth") {
		t, err := expandSystemAdminGuestAuth(d, v, "guest_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-auth"] = t
		}
	}

	if v, ok := d.GetOk("guest_lang"); ok || d.HasChange("guest_lang") {
		t, err := expandSystemAdminGuestLang(d, v, "guest_lang")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-lang"] = t
		}
	}

	if v, ok := d.GetOk("guest_usergroups"); ok || d.HasChange("guest_usergroups") {
		t, err := expandSystemAdminGuestUsergroups(d, v, "guest_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("hidden"); ok || d.HasChange("hidden") {
		t, err := expandSystemAdminHidden(d, v, "hidden")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hidden"] = t
		}
	}

	if v, ok := d.GetOk("gui_default_dashboard_template"); ok || d.HasChange("gui_default_dashboard_template") {
		t, err := expandSystemAdminGuiDefaultDashboardTemplate(d, v, "gui_default_dashboard_template")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-default-dashboard-template"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_invalid_signature_version"); ok || d.HasChange("gui_ignore_invalid_signature_version") {
		t, err := expandSystemAdminGuiIgnoreInvalidSignatureVersion(d, v, "gui_ignore_invalid_signature_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-invalid-signature-version"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_release_overview_version"); ok || d.HasChange("gui_ignore_release_overview_version") {
		t, err := expandSystemAdminGuiIgnoreReleaseOverviewVersion(d, v, "gui_ignore_release_overview_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-release-overview-version"] = t
		}
	}

	if v, ok := d.GetOk("history0"); ok || d.HasChange("history0") {
		t, err := expandSystemAdminHistory0(d, v, "history0")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history0"] = t
		}
	}

	if v, ok := d.GetOk("history1"); ok || d.HasChange("history1") {
		t, err := expandSystemAdminHistory1(d, v, "history1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["history1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost1"); ok || d.HasChange("ip6_trusthost1") {
		t, err := expandSystemAdminIp6Trusthost1(d, v, "ip6_trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost10"); ok || d.HasChange("ip6_trusthost10") {
		t, err := expandSystemAdminIp6Trusthost10(d, v, "ip6_trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost2"); ok || d.HasChange("ip6_trusthost2") {
		t, err := expandSystemAdminIp6Trusthost2(d, v, "ip6_trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost3"); ok || d.HasChange("ip6_trusthost3") {
		t, err := expandSystemAdminIp6Trusthost3(d, v, "ip6_trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost4"); ok || d.HasChange("ip6_trusthost4") {
		t, err := expandSystemAdminIp6Trusthost4(d, v, "ip6_trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost5"); ok || d.HasChange("ip6_trusthost5") {
		t, err := expandSystemAdminIp6Trusthost5(d, v, "ip6_trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost6"); ok || d.HasChange("ip6_trusthost6") {
		t, err := expandSystemAdminIp6Trusthost6(d, v, "ip6_trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost7"); ok || d.HasChange("ip6_trusthost7") {
		t, err := expandSystemAdminIp6Trusthost7(d, v, "ip6_trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost8"); ok || d.HasChange("ip6_trusthost8") {
		t, err := expandSystemAdminIp6Trusthost8(d, v, "ip6_trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("ip6_trusthost9"); ok || d.HasChange("ip6_trusthost9") {
		t, err := expandSystemAdminIp6Trusthost9(d, v, "ip6_trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemAdminName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok || d.HasChange("password") {
		t, err := expandSystemAdminPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("password_expire"); ok || d.HasChange("password_expire") {
		t, err := expandSystemAdminPasswordExpire(d, v, "password_expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-expire"] = t
		}
	}

	if v, ok := d.GetOk("peer_auth"); ok || d.HasChange("peer_auth") {
		t, err := expandSystemAdminPeerAuth(d, v, "peer_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-auth"] = t
		}
	}

	if v, ok := d.GetOk("peer_group"); ok || d.HasChange("peer_group") {
		t, err := expandSystemAdminPeerGroup(d, v, "peer_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-group"] = t
		}
	}

	if v, ok := d.GetOk("radius_vdom_override"); ok || d.HasChange("radius_vdom_override") {
		t, err := expandSystemAdminRadiusVdomOverride(d, v, "radius_vdom_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-vdom-override"] = t
		}
	}

	if v, ok := d.GetOk("remote_auth"); ok || d.HasChange("remote_auth") {
		t, err := expandSystemAdminRemoteAuth(d, v, "remote_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-auth"] = t
		}
	}

	if v, ok := d.GetOk("remote_group"); ok || d.HasChange("remote_group") {
		t, err := expandSystemAdminRemoteGroup(d, v, "remote_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-group"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok || d.HasChange("schedule") {
		t, err := expandSystemAdminSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok || d.HasChange("sms_custom_server") {
		t, err := expandSystemAdminSmsCustomServer(d, v, "sms_custom_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok || d.HasChange("sms_phone") {
		t, err := expandSystemAdminSmsPhone(d, v, "sms_phone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok || d.HasChange("sms_server") {
		t, err := expandSystemAdminSmsServer(d, v, "sms_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("ssh_certificate"); ok || d.HasChange("ssh_certificate") {
		t, err := expandSystemAdminSshCertificate(d, v, "ssh_certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key1"); ok || d.HasChange("ssh_public_key1") {
		t, err := expandSystemAdminSshPublicKey1(d, v, "ssh_public_key1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key1"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key2"); ok || d.HasChange("ssh_public_key2") {
		t, err := expandSystemAdminSshPublicKey2(d, v, "ssh_public_key2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key2"] = t
		}
	}

	if v, ok := d.GetOk("ssh_public_key3"); ok || d.HasChange("ssh_public_key3") {
		t, err := expandSystemAdminSshPublicKey3(d, v, "ssh_public_key3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-public-key3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost1"); ok || d.HasChange("trusthost1") {
		t, err := expandSystemAdminTrusthost1(d, v, "trusthost1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost1"] = t
		}
	}

	if v, ok := d.GetOk("trusthost10"); ok || d.HasChange("trusthost10") {
		t, err := expandSystemAdminTrusthost10(d, v, "trusthost10")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost10"] = t
		}
	}

	if v, ok := d.GetOk("trusthost2"); ok || d.HasChange("trusthost2") {
		t, err := expandSystemAdminTrusthost2(d, v, "trusthost2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost2"] = t
		}
	}

	if v, ok := d.GetOk("trusthost3"); ok || d.HasChange("trusthost3") {
		t, err := expandSystemAdminTrusthost3(d, v, "trusthost3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost3"] = t
		}
	}

	if v, ok := d.GetOk("trusthost4"); ok || d.HasChange("trusthost4") {
		t, err := expandSystemAdminTrusthost4(d, v, "trusthost4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost4"] = t
		}
	}

	if v, ok := d.GetOk("trusthost5"); ok || d.HasChange("trusthost5") {
		t, err := expandSystemAdminTrusthost5(d, v, "trusthost5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost5"] = t
		}
	}

	if v, ok := d.GetOk("trusthost6"); ok || d.HasChange("trusthost6") {
		t, err := expandSystemAdminTrusthost6(d, v, "trusthost6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost6"] = t
		}
	}

	if v, ok := d.GetOk("trusthost7"); ok || d.HasChange("trusthost7") {
		t, err := expandSystemAdminTrusthost7(d, v, "trusthost7")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost7"] = t
		}
	}

	if v, ok := d.GetOk("trusthost8"); ok || d.HasChange("trusthost8") {
		t, err := expandSystemAdminTrusthost8(d, v, "trusthost8")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost8"] = t
		}
	}

	if v, ok := d.GetOk("trusthost9"); ok || d.HasChange("trusthost9") {
		t, err := expandSystemAdminTrusthost9(d, v, "trusthost9")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthost9"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok || d.HasChange("two_factor") {
		t, err := expandSystemAdminTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok || d.HasChange("two_factor_authentication") {
		t, err := expandSystemAdminTwoFactorAuthentication(d, v, "two_factor_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok || d.HasChange("two_factor_notification") {
		t, err := expandSystemAdminTwoFactorNotification(d, v, "two_factor_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemAdminVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vdom_override"); ok || d.HasChange("vdom_override") {
		t, err := expandSystemAdminVdomOverride(d, v, "vdom_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-override"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok || d.HasChange("wildcard") {
		t, err := expandSystemAdminWildcard(d, v, "wildcard")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	return &obj, nil
}
