// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure USB auto installation.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoInstall() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoInstallUpdate,
		Read:   resourceSystemAutoInstallRead,
		Update: resourceSystemAutoInstallUpdate,
		Delete: resourceSystemAutoInstallDelete,

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
			"auto_install_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_install_image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_config_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_image_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutoInstallUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutoInstall(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoInstall resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoInstall(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoInstall resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemAutoInstall")

	return resourceSystemAutoInstallRead(d, m)
}

func resourceSystemAutoInstallDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemAutoInstall(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoInstall resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoInstallRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutoInstall(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoInstall resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoInstall(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoInstall resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoInstallAutoInstallConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoInstallAutoInstallImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoInstallDefaultConfigFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoInstallDefaultImageFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutoInstall(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("auto_install_config", flattenSystemAutoInstallAutoInstallConfig(o["auto-install-config"], d, "auto_install_config")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-install-config"], "SystemAutoInstall-AutoInstallConfig"); ok {
			if err = d.Set("auto_install_config", vv); err != nil {
				return fmt.Errorf("Error reading auto_install_config: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_install_config: %v", err)
		}
	}

	if err = d.Set("auto_install_image", flattenSystemAutoInstallAutoInstallImage(o["auto-install-image"], d, "auto_install_image")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-install-image"], "SystemAutoInstall-AutoInstallImage"); ok {
			if err = d.Set("auto_install_image", vv); err != nil {
				return fmt.Errorf("Error reading auto_install_image: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_install_image: %v", err)
		}
	}

	if err = d.Set("default_config_file", flattenSystemAutoInstallDefaultConfigFile(o["default-config-file"], d, "default_config_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-config-file"], "SystemAutoInstall-DefaultConfigFile"); ok {
			if err = d.Set("default_config_file", vv); err != nil {
				return fmt.Errorf("Error reading default_config_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_config_file: %v", err)
		}
	}

	if err = d.Set("default_image_file", flattenSystemAutoInstallDefaultImageFile(o["default-image-file"], d, "default_image_file")); err != nil {
		if vv, ok := fortiAPIPatch(o["default-image-file"], "SystemAutoInstall-DefaultImageFile"); ok {
			if err = d.Set("default_image_file", vv); err != nil {
				return fmt.Errorf("Error reading default_image_file: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading default_image_file: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoInstallFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutoInstallAutoInstallConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoInstallAutoInstallImage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoInstallDefaultConfigFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoInstallDefaultImageFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoInstall(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auto_install_config"); ok || d.HasChange("auto_install_config") {
		t, err := expandSystemAutoInstallAutoInstallConfig(d, v, "auto_install_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-install-config"] = t
		}
	}

	if v, ok := d.GetOk("auto_install_image"); ok || d.HasChange("auto_install_image") {
		t, err := expandSystemAutoInstallAutoInstallImage(d, v, "auto_install_image")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-install-image"] = t
		}
	}

	if v, ok := d.GetOk("default_config_file"); ok || d.HasChange("default_config_file") {
		t, err := expandSystemAutoInstallDefaultConfigFile(d, v, "default_config_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-config-file"] = t
		}
	}

	if v, ok := d.GetOk("default_image_file"); ok || d.HasChange("default_image_file") {
		t, err := expandSystemAutoInstallDefaultImageFile(d, v, "default_image_file")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-image-file"] = t
		}
	}

	return &obj, nil
}
