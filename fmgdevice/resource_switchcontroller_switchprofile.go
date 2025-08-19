// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch switch profile.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSwitchProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSwitchProfileCreate,
		Read:   resourceSwitchControllerSwitchProfileRead,
		Update: resourceSwitchControllerSwitchProfileUpdate,
		Delete: resourceSwitchControllerSwitchProfileDelete,

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
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_passwd": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"login_passwd_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision_backup_on_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSwitchProfileCreate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSwitchProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchProfile resource while getting object: %v", err)
	}

	_, err = c.CreateSwitchControllerSwitchProfile(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSwitchProfileRead(d, m)
}

func resourceSwitchControllerSwitchProfileUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSwitchControllerSwitchProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSwitchProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSwitchControllerSwitchProfileRead(d, m)
}

func resourceSwitchControllerSwitchProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSwitchControllerSwitchProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSwitchProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSwitchProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSwitchProfile(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSwitchProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSwitchProfileLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileLoginPasswdOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileRevisionBackupOnUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSwitchProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("login", flattenSwitchControllerSwitchProfileLogin(o["login"], d, "login")); err != nil {
		if vv, ok := fortiAPIPatch(o["login"], "SwitchControllerSwitchProfile-Login"); ok {
			if err = d.Set("login", vv); err != nil {
				return fmt.Errorf("Error reading login: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	if err = d.Set("login_passwd_override", flattenSwitchControllerSwitchProfileLoginPasswdOverride(o["login-passwd-override"], d, "login_passwd_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["login-passwd-override"], "SwitchControllerSwitchProfile-LoginPasswdOverride"); ok {
			if err = d.Set("login_passwd_override", vv); err != nil {
				return fmt.Errorf("Error reading login_passwd_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading login_passwd_override: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerSwitchProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SwitchControllerSwitchProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", flattenSwitchControllerSwitchProfileRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision-backup-on-logout"], "SwitchControllerSwitchProfile-RevisionBackupOnLogout"); ok {
			if err = d.Set("revision_backup_on_logout", vv); err != nil {
				return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_upgrade", flattenSwitchControllerSwitchProfileRevisionBackupOnUpgrade(o["revision-backup-on-upgrade"], d, "revision_backup_on_upgrade")); err != nil {
		if vv, ok := fortiAPIPatch(o["revision-backup-on-upgrade"], "SwitchControllerSwitchProfile-RevisionBackupOnUpgrade"); ok {
			if err = d.Set("revision_backup_on_upgrade", vv); err != nil {
				return fmt.Errorf("Error reading revision_backup_on_upgrade: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading revision_backup_on_upgrade: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSwitchProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSwitchProfileLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileLoginPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSwitchControllerSwitchProfileLoginPasswdOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileRevisionBackupOnLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileRevisionBackupOnUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSwitchProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("login"); ok || d.HasChange("login") {
		t, err := expandSwitchControllerSwitchProfileLogin(d, v, "login")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok || d.HasChange("login_passwd") {
		t, err := expandSwitchControllerSwitchProfileLoginPasswd(d, v, "login_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_override"); ok || d.HasChange("login_passwd_override") {
		t, err := expandSwitchControllerSwitchProfileLoginPasswdOverride(d, v, "login_passwd_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-override"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSwitchControllerSwitchProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("revision_backup_on_logout"); ok || d.HasChange("revision_backup_on_logout") {
		t, err := expandSwitchControllerSwitchProfileRevisionBackupOnLogout(d, v, "revision_backup_on_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-backup-on-logout"] = t
		}
	}

	if v, ok := d.GetOk("revision_backup_on_upgrade"); ok || d.HasChange("revision_backup_on_upgrade") {
		t, err := expandSwitchControllerSwitchProfileRevisionBackupOnUpgrade(d, v, "revision_backup_on_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-backup-on-upgrade"] = t
		}
	}

	return &obj, nil
}
