// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure SSH config.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSshConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSshConfigUpdate,
		Read:   resourceSystemSshConfigRead,
		Update: resourceSystemSshConfigUpdate,
		Delete: resourceSystemSshConfigDelete,

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
			"ssh_enc_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_hsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_hsk_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_hsk_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_hsk_password": &schema.Schema{
				Type:      schema.TypeSet,
				Elem:      &schema.Schema{Type: schema.TypeString},
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"ssh_kex_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"ssh_mac_algo": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSshConfigUpdate(d *schema.ResourceData, m interface{}) error {
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
	obj, err := getObjectSystemSshConfig(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSshConfig resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSshConfig(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemSshConfig resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemSshConfig")

	return resourceSystemSshConfigRead(d, m)
}

func resourceSystemSshConfigDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemSshConfig(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSshConfig resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSshConfigRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSshConfig(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemSshConfig resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSshConfig(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSshConfig resource from API: %v", err)
	}
	return nil
}

func flattenSystemSshConfigSshEncAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSshConfigSshHsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSshConfigSshHskAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSshConfigSshHskOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSshConfigSshKexAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenSystemSshConfigSshMacAlgo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectSystemSshConfig(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ssh_enc_algo", flattenSystemSshConfigSshEncAlgo(o["ssh-enc-algo"], d, "ssh_enc_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-enc-algo"], "SystemSshConfig-SshEncAlgo"); ok {
			if err = d.Set("ssh_enc_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hsk", flattenSystemSshConfigSshHsk(o["ssh-hsk"], d, "ssh_hsk")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hsk"], "SystemSshConfig-SshHsk"); ok {
			if err = d.Set("ssh_hsk", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hsk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hsk: %v", err)
		}
	}

	if err = d.Set("ssh_hsk_algo", flattenSystemSshConfigSshHskAlgo(o["ssh-hsk-algo"], d, "ssh_hsk_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hsk-algo"], "SystemSshConfig-SshHskAlgo"); ok {
			if err = d.Set("ssh_hsk_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hsk_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hsk_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hsk_override", flattenSystemSshConfigSshHskOverride(o["ssh-hsk-override"], d, "ssh_hsk_override")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-hsk-override"], "SystemSshConfig-SshHskOverride"); ok {
			if err = d.Set("ssh_hsk_override", vv); err != nil {
				return fmt.Errorf("Error reading ssh_hsk_override: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_hsk_override: %v", err)
		}
	}

	if err = d.Set("ssh_kex_algo", flattenSystemSshConfigSshKexAlgo(o["ssh-kex-algo"], d, "ssh_kex_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-kex-algo"], "SystemSshConfig-SshKexAlgo"); ok {
			if err = d.Set("ssh_kex_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
		}
	}

	if err = d.Set("ssh_mac_algo", flattenSystemSshConfigSshMacAlgo(o["ssh-mac-algo"], d, "ssh_mac_algo")); err != nil {
		if vv, ok := fortiAPIPatch(o["ssh-mac-algo"], "SystemSshConfig-SshMacAlgo"); ok {
			if err = d.Set("ssh_mac_algo", vv); err != nil {
				return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
		}
	}

	return nil
}

func flattenSystemSshConfigFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSshConfigSshEncAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSshConfigSshHsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshHskAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSshConfigSshHskOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshHskPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSshConfigSshKexAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandSystemSshConfigSshMacAlgo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectSystemSshConfig(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ssh_enc_algo"); ok || d.HasChange("ssh_enc_algo") {
		t, err := expandSystemSshConfigSshEncAlgo(d, v, "ssh_enc_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-enc-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hsk"); ok || d.HasChange("ssh_hsk") {
		t, err := expandSystemSshConfigSshHsk(d, v, "ssh_hsk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hsk"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hsk_algo"); ok || d.HasChange("ssh_hsk_algo") {
		t, err := expandSystemSshConfigSshHskAlgo(d, v, "ssh_hsk_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hsk-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hsk_override"); ok || d.HasChange("ssh_hsk_override") {
		t, err := expandSystemSshConfigSshHskOverride(d, v, "ssh_hsk_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hsk-override"] = t
		}
	}

	if v, ok := d.GetOk("ssh_hsk_password"); ok || d.HasChange("ssh_hsk_password") {
		t, err := expandSystemSshConfigSshHskPassword(d, v, "ssh_hsk_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-hsk-password"] = t
		}
	}

	if v, ok := d.GetOk("ssh_kex_algo"); ok || d.HasChange("ssh_kex_algo") {
		t, err := expandSystemSshConfigSshKexAlgo(d, v, "ssh_kex_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-kex-algo"] = t
		}
	}

	if v, ok := d.GetOk("ssh_mac_algo"); ok || d.HasChange("ssh_mac_algo") {
		t, err := expandSystemSshConfigSshMacAlgo(d, v, "ssh_mac_algo")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-mac-algo"] = t
		}
	}

	return &obj, nil
}
