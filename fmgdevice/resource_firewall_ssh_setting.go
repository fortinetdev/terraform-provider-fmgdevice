// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SSH proxy settings.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSshSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSshSettingUpdate,
		Read:   resourceFirewallSshSettingRead,
		Update: resourceFirewallSshSettingUpdate,
		Delete: resourceFirewallSshSettingDelete,

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
			"caname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_trusted_checking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostkey_dsa1024": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hostkey_ecdsa256": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hostkey_ecdsa384": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hostkey_ecdsa521": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hostkey_ed25519": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"hostkey_rsa2048": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"untrusted_caname": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSshSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectFirewallSshSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallSshSetting(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FirewallSshSetting")

	return resourceFirewallSshSettingRead(d, m)
}

func resourceFirewallSshSettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteFirewallSshSetting(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSshSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallSshSetting(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSshSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshSetting resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSshSettingCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenFirewallSshSettingHostTrustedChecking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyDsa1024(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSshSettingHostkeyEcdsa256(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSshSettingHostkeyEcdsa384(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSshSettingHostkeyEcdsa521(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSshSettingHostkeyEd25519(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSshSettingHostkeyRsa2048(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenFirewallSshSettingUntrustedCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectFirewallSshSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("caname", flattenFirewallSshSettingCaname(o["caname"], d, "caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["caname"], "FirewallSshSetting-Caname"); ok {
			if err = d.Set("caname", vv); err != nil {
				return fmt.Errorf("Error reading caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("host_trusted_checking", flattenFirewallSshSettingHostTrustedChecking(o["host-trusted-checking"], d, "host_trusted_checking")); err != nil {
		if vv, ok := fortiAPIPatch(o["host-trusted-checking"], "FirewallSshSetting-HostTrustedChecking"); ok {
			if err = d.Set("host_trusted_checking", vv); err != nil {
				return fmt.Errorf("Error reading host_trusted_checking: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading host_trusted_checking: %v", err)
		}
	}

	if err = d.Set("hostkey_dsa1024", flattenFirewallSshSettingHostkeyDsa1024(o["hostkey-dsa1024"], d, "hostkey_dsa1024")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostkey-dsa1024"], "FirewallSshSetting-HostkeyDsa1024"); ok {
			if err = d.Set("hostkey_dsa1024", vv); err != nil {
				return fmt.Errorf("Error reading hostkey_dsa1024: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostkey_dsa1024: %v", err)
		}
	}

	if err = d.Set("hostkey_ecdsa256", flattenFirewallSshSettingHostkeyEcdsa256(o["hostkey-ecdsa256"], d, "hostkey_ecdsa256")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostkey-ecdsa256"], "FirewallSshSetting-HostkeyEcdsa256"); ok {
			if err = d.Set("hostkey_ecdsa256", vv); err != nil {
				return fmt.Errorf("Error reading hostkey_ecdsa256: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostkey_ecdsa256: %v", err)
		}
	}

	if err = d.Set("hostkey_ecdsa384", flattenFirewallSshSettingHostkeyEcdsa384(o["hostkey-ecdsa384"], d, "hostkey_ecdsa384")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostkey-ecdsa384"], "FirewallSshSetting-HostkeyEcdsa384"); ok {
			if err = d.Set("hostkey_ecdsa384", vv); err != nil {
				return fmt.Errorf("Error reading hostkey_ecdsa384: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostkey_ecdsa384: %v", err)
		}
	}

	if err = d.Set("hostkey_ecdsa521", flattenFirewallSshSettingHostkeyEcdsa521(o["hostkey-ecdsa521"], d, "hostkey_ecdsa521")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostkey-ecdsa521"], "FirewallSshSetting-HostkeyEcdsa521"); ok {
			if err = d.Set("hostkey_ecdsa521", vv); err != nil {
				return fmt.Errorf("Error reading hostkey_ecdsa521: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostkey_ecdsa521: %v", err)
		}
	}

	if err = d.Set("hostkey_ed25519", flattenFirewallSshSettingHostkeyEd25519(o["hostkey-ed25519"], d, "hostkey_ed25519")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostkey-ed25519"], "FirewallSshSetting-HostkeyEd25519"); ok {
			if err = d.Set("hostkey_ed25519", vv); err != nil {
				return fmt.Errorf("Error reading hostkey_ed25519: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostkey_ed25519: %v", err)
		}
	}

	if err = d.Set("hostkey_rsa2048", flattenFirewallSshSettingHostkeyRsa2048(o["hostkey-rsa2048"], d, "hostkey_rsa2048")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostkey-rsa2048"], "FirewallSshSetting-HostkeyRsa2048"); ok {
			if err = d.Set("hostkey_rsa2048", vv); err != nil {
				return fmt.Errorf("Error reading hostkey_rsa2048: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostkey_rsa2048: %v", err)
		}
	}

	if err = d.Set("untrusted_caname", flattenFirewallSshSettingUntrustedCaname(o["untrusted-caname"], d, "untrusted_caname")); err != nil {
		if vv, ok := fortiAPIPatch(o["untrusted-caname"], "FirewallSshSetting-UntrustedCaname"); ok {
			if err = d.Set("untrusted_caname", vv); err != nil {
				return fmt.Errorf("Error reading untrusted_caname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading untrusted_caname: %v", err)
		}
	}

	return nil
}

func flattenFirewallSshSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSshSettingCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandFirewallSshSettingHostTrustedChecking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyDsa1024(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshSettingHostkeyEcdsa256(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshSettingHostkeyEcdsa384(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshSettingHostkeyEcdsa521(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshSettingHostkeyEd25519(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshSettingHostkeyRsa2048(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandFirewallSshSettingUntrustedCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFirewallSshSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("caname"); ok || d.HasChange("caname") {
		t, err := expandFirewallSshSettingCaname(d, v, "caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caname"] = t
		}
	}

	if v, ok := d.GetOk("host_trusted_checking"); ok || d.HasChange("host_trusted_checking") {
		t, err := expandFirewallSshSettingHostTrustedChecking(d, v, "host_trusted_checking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-trusted-checking"] = t
		}
	}

	if v, ok := d.GetOk("hostkey_dsa1024"); ok || d.HasChange("hostkey_dsa1024") {
		t, err := expandFirewallSshSettingHostkeyDsa1024(d, v, "hostkey_dsa1024")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostkey-dsa1024"] = t
		}
	}

	if v, ok := d.GetOk("hostkey_ecdsa256"); ok || d.HasChange("hostkey_ecdsa256") {
		t, err := expandFirewallSshSettingHostkeyEcdsa256(d, v, "hostkey_ecdsa256")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostkey-ecdsa256"] = t
		}
	}

	if v, ok := d.GetOk("hostkey_ecdsa384"); ok || d.HasChange("hostkey_ecdsa384") {
		t, err := expandFirewallSshSettingHostkeyEcdsa384(d, v, "hostkey_ecdsa384")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostkey-ecdsa384"] = t
		}
	}

	if v, ok := d.GetOk("hostkey_ecdsa521"); ok || d.HasChange("hostkey_ecdsa521") {
		t, err := expandFirewallSshSettingHostkeyEcdsa521(d, v, "hostkey_ecdsa521")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostkey-ecdsa521"] = t
		}
	}

	if v, ok := d.GetOk("hostkey_ed25519"); ok || d.HasChange("hostkey_ed25519") {
		t, err := expandFirewallSshSettingHostkeyEd25519(d, v, "hostkey_ed25519")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostkey-ed25519"] = t
		}
	}

	if v, ok := d.GetOk("hostkey_rsa2048"); ok || d.HasChange("hostkey_rsa2048") {
		t, err := expandFirewallSshSettingHostkeyRsa2048(d, v, "hostkey_rsa2048")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostkey-rsa2048"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_caname"); ok || d.HasChange("untrusted_caname") {
		t, err := expandFirewallSshSettingUntrustedCaname(d, v, "untrusted_caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-caname"] = t
		}
	}

	return &obj, nil
}
