// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i>

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceImageAnalyzerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageAnalyzerProfileCreate,
		Read:   resourceImageAnalyzerProfileRead,
		Update: resourceImageAnalyzerProfileUpdate,
		Delete: resourceImageAnalyzerProfileDelete,

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
			"alcohol_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"alcohol_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"blocked_img_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drugs_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"drugs_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extremism_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"extremism_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gambling_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gambling_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gore_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"gore_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_skip_height": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"image_skip_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"image_skip_width": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"ocr_activation_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"optical_character_recognition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"porn_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"porn_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rating_err_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replace_image": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"source_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"swim_underwear_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"swim_underwear_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weapons_block_strictness_level": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"weapons_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceImageAnalyzerProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectImageAnalyzerProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ImageAnalyzerProfile resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadImageAnalyzerProfile(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateImageAnalyzerProfile(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ImageAnalyzerProfile resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateImageAnalyzerProfile(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ImageAnalyzerProfile resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceImageAnalyzerProfileRead(d, m)
}

func resourceImageAnalyzerProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectImageAnalyzerProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ImageAnalyzerProfile resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateImageAnalyzerProfile(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ImageAnalyzerProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceImageAnalyzerProfileRead(d, m)
}

func resourceImageAnalyzerProfileDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteImageAnalyzerProfile(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ImageAnalyzerProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceImageAnalyzerProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadImageAnalyzerProfile(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ImageAnalyzerProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectImageAnalyzerProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ImageAnalyzerProfile resource from API: %v", err)
	}
	return nil
}

func flattenImageAnalyzerProfileAlcoholBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileAlcoholStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileBlockedImgCache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileDrugsBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileDrugsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileExtremismBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileExtremismStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileGamblingBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileGamblingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileGoreBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileGoreStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileImageSkipHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileImageSkipSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileImageSkipWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileLogOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileOcrActivationThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileOpticalCharacterRecognition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfilePornBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfilePornStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileRatingErrAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileReplaceImage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenImageAnalyzerProfileSourceUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileSwimUnderwearStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileWeaponsBlockStrictnessLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenImageAnalyzerProfileWeaponsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectImageAnalyzerProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("alcohol_block_strictness_level", flattenImageAnalyzerProfileAlcoholBlockStrictnessLevel(o["alcohol-block-strictness-level"], d, "alcohol_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["alcohol-block-strictness-level"], "ImageAnalyzerProfile-AlcoholBlockStrictnessLevel"); ok {
			if err = d.Set("alcohol_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading alcohol_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alcohol_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("alcohol_status", flattenImageAnalyzerProfileAlcoholStatus(o["alcohol-status"], d, "alcohol_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["alcohol-status"], "ImageAnalyzerProfile-AlcoholStatus"); ok {
			if err = d.Set("alcohol_status", vv); err != nil {
				return fmt.Errorf("Error reading alcohol_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading alcohol_status: %v", err)
		}
	}

	if err = d.Set("blocked_img_cache", flattenImageAnalyzerProfileBlockedImgCache(o["blocked-img-cache"], d, "blocked_img_cache")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-img-cache"], "ImageAnalyzerProfile-BlockedImgCache"); ok {
			if err = d.Set("blocked_img_cache", vv); err != nil {
				return fmt.Errorf("Error reading blocked_img_cache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_img_cache: %v", err)
		}
	}

	if err = d.Set("comment", flattenImageAnalyzerProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ImageAnalyzerProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("drugs_block_strictness_level", flattenImageAnalyzerProfileDrugsBlockStrictnessLevel(o["drugs-block-strictness-level"], d, "drugs_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["drugs-block-strictness-level"], "ImageAnalyzerProfile-DrugsBlockStrictnessLevel"); ok {
			if err = d.Set("drugs_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading drugs_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drugs_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("drugs_status", flattenImageAnalyzerProfileDrugsStatus(o["drugs-status"], d, "drugs_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["drugs-status"], "ImageAnalyzerProfile-DrugsStatus"); ok {
			if err = d.Set("drugs_status", vv); err != nil {
				return fmt.Errorf("Error reading drugs_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading drugs_status: %v", err)
		}
	}

	if err = d.Set("extremism_block_strictness_level", flattenImageAnalyzerProfileExtremismBlockStrictnessLevel(o["extremism-block-strictness-level"], d, "extremism_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["extremism-block-strictness-level"], "ImageAnalyzerProfile-ExtremismBlockStrictnessLevel"); ok {
			if err = d.Set("extremism_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading extremism_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extremism_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("extremism_status", flattenImageAnalyzerProfileExtremismStatus(o["extremism-status"], d, "extremism_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["extremism-status"], "ImageAnalyzerProfile-ExtremismStatus"); ok {
			if err = d.Set("extremism_status", vv); err != nil {
				return fmt.Errorf("Error reading extremism_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading extremism_status: %v", err)
		}
	}

	if err = d.Set("gambling_block_strictness_level", flattenImageAnalyzerProfileGamblingBlockStrictnessLevel(o["gambling-block-strictness-level"], d, "gambling_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["gambling-block-strictness-level"], "ImageAnalyzerProfile-GamblingBlockStrictnessLevel"); ok {
			if err = d.Set("gambling_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading gambling_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gambling_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("gambling_status", flattenImageAnalyzerProfileGamblingStatus(o["gambling-status"], d, "gambling_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["gambling-status"], "ImageAnalyzerProfile-GamblingStatus"); ok {
			if err = d.Set("gambling_status", vv); err != nil {
				return fmt.Errorf("Error reading gambling_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gambling_status: %v", err)
		}
	}

	if err = d.Set("gore_block_strictness_level", flattenImageAnalyzerProfileGoreBlockStrictnessLevel(o["gore-block-strictness-level"], d, "gore_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["gore-block-strictness-level"], "ImageAnalyzerProfile-GoreBlockStrictnessLevel"); ok {
			if err = d.Set("gore_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading gore_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gore_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("gore_status", flattenImageAnalyzerProfileGoreStatus(o["gore-status"], d, "gore_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["gore-status"], "ImageAnalyzerProfile-GoreStatus"); ok {
			if err = d.Set("gore_status", vv); err != nil {
				return fmt.Errorf("Error reading gore_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading gore_status: %v", err)
		}
	}

	if err = d.Set("image_skip_height", flattenImageAnalyzerProfileImageSkipHeight(o["image-skip-height"], d, "image_skip_height")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-skip-height"], "ImageAnalyzerProfile-ImageSkipHeight"); ok {
			if err = d.Set("image_skip_height", vv); err != nil {
				return fmt.Errorf("Error reading image_skip_height: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_skip_height: %v", err)
		}
	}

	if err = d.Set("image_skip_size", flattenImageAnalyzerProfileImageSkipSize(o["image-skip-size"], d, "image_skip_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-skip-size"], "ImageAnalyzerProfile-ImageSkipSize"); ok {
			if err = d.Set("image_skip_size", vv); err != nil {
				return fmt.Errorf("Error reading image_skip_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_skip_size: %v", err)
		}
	}

	if err = d.Set("image_skip_width", flattenImageAnalyzerProfileImageSkipWidth(o["image-skip-width"], d, "image_skip_width")); err != nil {
		if vv, ok := fortiAPIPatch(o["image-skip-width"], "ImageAnalyzerProfile-ImageSkipWidth"); ok {
			if err = d.Set("image_skip_width", vv); err != nil {
				return fmt.Errorf("Error reading image_skip_width: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading image_skip_width: %v", err)
		}
	}

	if err = d.Set("log_option", flattenImageAnalyzerProfileLogOption(o["log-option"], d, "log_option")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-option"], "ImageAnalyzerProfile-LogOption"); ok {
			if err = d.Set("log_option", vv); err != nil {
				return fmt.Errorf("Error reading log_option: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_option: %v", err)
		}
	}

	if err = d.Set("name", flattenImageAnalyzerProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ImageAnalyzerProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ocr_activation_threshold", flattenImageAnalyzerProfileOcrActivationThreshold(o["ocr-activation-threshold"], d, "ocr_activation_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["ocr-activation-threshold"], "ImageAnalyzerProfile-OcrActivationThreshold"); ok {
			if err = d.Set("ocr_activation_threshold", vv); err != nil {
				return fmt.Errorf("Error reading ocr_activation_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ocr_activation_threshold: %v", err)
		}
	}

	if err = d.Set("optical_character_recognition", flattenImageAnalyzerProfileOpticalCharacterRecognition(o["optical-character-recognition"], d, "optical_character_recognition")); err != nil {
		if vv, ok := fortiAPIPatch(o["optical-character-recognition"], "ImageAnalyzerProfile-OpticalCharacterRecognition"); ok {
			if err = d.Set("optical_character_recognition", vv); err != nil {
				return fmt.Errorf("Error reading optical_character_recognition: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading optical_character_recognition: %v", err)
		}
	}

	if err = d.Set("porn_block_strictness_level", flattenImageAnalyzerProfilePornBlockStrictnessLevel(o["porn-block-strictness-level"], d, "porn_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["porn-block-strictness-level"], "ImageAnalyzerProfile-PornBlockStrictnessLevel"); ok {
			if err = d.Set("porn_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading porn_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading porn_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("porn_status", flattenImageAnalyzerProfilePornStatus(o["porn-status"], d, "porn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["porn-status"], "ImageAnalyzerProfile-PornStatus"); ok {
			if err = d.Set("porn_status", vv); err != nil {
				return fmt.Errorf("Error reading porn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading porn_status: %v", err)
		}
	}

	if err = d.Set("rating_err_action", flattenImageAnalyzerProfileRatingErrAction(o["rating-err-action"], d, "rating_err_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["rating-err-action"], "ImageAnalyzerProfile-RatingErrAction"); ok {
			if err = d.Set("rating_err_action", vv); err != nil {
				return fmt.Errorf("Error reading rating_err_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rating_err_action: %v", err)
		}
	}

	if err = d.Set("replace_image", flattenImageAnalyzerProfileReplaceImage(o["replace-image"], d, "replace_image")); err != nil {
		if vv, ok := fortiAPIPatch(o["replace-image"], "ImageAnalyzerProfile-ReplaceImage"); ok {
			if err = d.Set("replace_image", vv); err != nil {
				return fmt.Errorf("Error reading replace_image: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading replace_image: %v", err)
		}
	}

	if err = d.Set("source_url", flattenImageAnalyzerProfileSourceUrl(o["source-url"], d, "source_url")); err != nil {
		if vv, ok := fortiAPIPatch(o["source-url"], "ImageAnalyzerProfile-SourceUrl"); ok {
			if err = d.Set("source_url", vv); err != nil {
				return fmt.Errorf("Error reading source_url: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading source_url: %v", err)
		}
	}

	if err = d.Set("swim_underwear_block_strictness_level", flattenImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(o["swim_underwear-block-strictness-level"], d, "swim_underwear_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["swim_underwear-block-strictness-level"], "ImageAnalyzerProfile-SwimUnderwearBlockStrictnessLevel"); ok {
			if err = d.Set("swim_underwear_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading swim_underwear_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading swim_underwear_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("swim_underwear_status", flattenImageAnalyzerProfileSwimUnderwearStatus(o["swim_underwear-status"], d, "swim_underwear_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["swim_underwear-status"], "ImageAnalyzerProfile-SwimUnderwearStatus"); ok {
			if err = d.Set("swim_underwear_status", vv); err != nil {
				return fmt.Errorf("Error reading swim_underwear_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading swim_underwear_status: %v", err)
		}
	}

	if err = d.Set("weapons_block_strictness_level", flattenImageAnalyzerProfileWeaponsBlockStrictnessLevel(o["weapons-block-strictness-level"], d, "weapons_block_strictness_level")); err != nil {
		if vv, ok := fortiAPIPatch(o["weapons-block-strictness-level"], "ImageAnalyzerProfile-WeaponsBlockStrictnessLevel"); ok {
			if err = d.Set("weapons_block_strictness_level", vv); err != nil {
				return fmt.Errorf("Error reading weapons_block_strictness_level: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weapons_block_strictness_level: %v", err)
		}
	}

	if err = d.Set("weapons_status", flattenImageAnalyzerProfileWeaponsStatus(o["weapons-status"], d, "weapons_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["weapons-status"], "ImageAnalyzerProfile-WeaponsStatus"); ok {
			if err = d.Set("weapons_status", vv); err != nil {
				return fmt.Errorf("Error reading weapons_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading weapons_status: %v", err)
		}
	}

	return nil
}

func flattenImageAnalyzerProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandImageAnalyzerProfileAlcoholBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileAlcoholStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileBlockedImgCache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileDrugsBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileDrugsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileExtremismBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileExtremismStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileGamblingBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileGamblingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileGoreBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileGoreStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileImageSkipHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileImageSkipSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileImageSkipWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileLogOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileOcrActivationThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileOpticalCharacterRecognition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfilePornBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfilePornStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileRatingErrAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileReplaceImage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandImageAnalyzerProfileSourceUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileSwimUnderwearStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileWeaponsBlockStrictnessLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandImageAnalyzerProfileWeaponsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectImageAnalyzerProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("alcohol_block_strictness_level"); ok || d.HasChange("alcohol_block_strictness_level") {
		t, err := expandImageAnalyzerProfileAlcoholBlockStrictnessLevel(d, v, "alcohol_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alcohol-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("alcohol_status"); ok || d.HasChange("alcohol_status") {
		t, err := expandImageAnalyzerProfileAlcoholStatus(d, v, "alcohol_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alcohol-status"] = t
		}
	}

	if v, ok := d.GetOk("blocked_img_cache"); ok || d.HasChange("blocked_img_cache") {
		t, err := expandImageAnalyzerProfileBlockedImgCache(d, v, "blocked_img_cache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-img-cache"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandImageAnalyzerProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("drugs_block_strictness_level"); ok || d.HasChange("drugs_block_strictness_level") {
		t, err := expandImageAnalyzerProfileDrugsBlockStrictnessLevel(d, v, "drugs_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drugs-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("drugs_status"); ok || d.HasChange("drugs_status") {
		t, err := expandImageAnalyzerProfileDrugsStatus(d, v, "drugs_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drugs-status"] = t
		}
	}

	if v, ok := d.GetOk("extremism_block_strictness_level"); ok || d.HasChange("extremism_block_strictness_level") {
		t, err := expandImageAnalyzerProfileExtremismBlockStrictnessLevel(d, v, "extremism_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extremism-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("extremism_status"); ok || d.HasChange("extremism_status") {
		t, err := expandImageAnalyzerProfileExtremismStatus(d, v, "extremism_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extremism-status"] = t
		}
	}

	if v, ok := d.GetOk("gambling_block_strictness_level"); ok || d.HasChange("gambling_block_strictness_level") {
		t, err := expandImageAnalyzerProfileGamblingBlockStrictnessLevel(d, v, "gambling_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gambling-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("gambling_status"); ok || d.HasChange("gambling_status") {
		t, err := expandImageAnalyzerProfileGamblingStatus(d, v, "gambling_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gambling-status"] = t
		}
	}

	if v, ok := d.GetOk("gore_block_strictness_level"); ok || d.HasChange("gore_block_strictness_level") {
		t, err := expandImageAnalyzerProfileGoreBlockStrictnessLevel(d, v, "gore_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gore-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("gore_status"); ok || d.HasChange("gore_status") {
		t, err := expandImageAnalyzerProfileGoreStatus(d, v, "gore_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gore-status"] = t
		}
	}

	if v, ok := d.GetOk("image_skip_height"); ok || d.HasChange("image_skip_height") {
		t, err := expandImageAnalyzerProfileImageSkipHeight(d, v, "image_skip_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-skip-height"] = t
		}
	}

	if v, ok := d.GetOk("image_skip_size"); ok || d.HasChange("image_skip_size") {
		t, err := expandImageAnalyzerProfileImageSkipSize(d, v, "image_skip_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-skip-size"] = t
		}
	}

	if v, ok := d.GetOk("image_skip_width"); ok || d.HasChange("image_skip_width") {
		t, err := expandImageAnalyzerProfileImageSkipWidth(d, v, "image_skip_width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-skip-width"] = t
		}
	}

	if v, ok := d.GetOk("log_option"); ok || d.HasChange("log_option") {
		t, err := expandImageAnalyzerProfileLogOption(d, v, "log_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-option"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandImageAnalyzerProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ocr_activation_threshold"); ok || d.HasChange("ocr_activation_threshold") {
		t, err := expandImageAnalyzerProfileOcrActivationThreshold(d, v, "ocr_activation_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocr-activation-threshold"] = t
		}
	}

	if v, ok := d.GetOk("optical_character_recognition"); ok || d.HasChange("optical_character_recognition") {
		t, err := expandImageAnalyzerProfileOpticalCharacterRecognition(d, v, "optical_character_recognition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optical-character-recognition"] = t
		}
	}

	if v, ok := d.GetOk("porn_block_strictness_level"); ok || d.HasChange("porn_block_strictness_level") {
		t, err := expandImageAnalyzerProfilePornBlockStrictnessLevel(d, v, "porn_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["porn-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("porn_status"); ok || d.HasChange("porn_status") {
		t, err := expandImageAnalyzerProfilePornStatus(d, v, "porn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["porn-status"] = t
		}
	}

	if v, ok := d.GetOk("rating_err_action"); ok || d.HasChange("rating_err_action") {
		t, err := expandImageAnalyzerProfileRatingErrAction(d, v, "rating_err_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rating-err-action"] = t
		}
	}

	if v, ok := d.GetOk("replace_image"); ok || d.HasChange("replace_image") {
		t, err := expandImageAnalyzerProfileReplaceImage(d, v, "replace_image")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replace-image"] = t
		}
	}

	if v, ok := d.GetOk("source_url"); ok || d.HasChange("source_url") {
		t, err := expandImageAnalyzerProfileSourceUrl(d, v, "source_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-url"] = t
		}
	}

	if v, ok := d.GetOk("swim_underwear_block_strictness_level"); ok || d.HasChange("swim_underwear_block_strictness_level") {
		t, err := expandImageAnalyzerProfileSwimUnderwearBlockStrictnessLevel(d, v, "swim_underwear_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swim_underwear-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("swim_underwear_status"); ok || d.HasChange("swim_underwear_status") {
		t, err := expandImageAnalyzerProfileSwimUnderwearStatus(d, v, "swim_underwear_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["swim_underwear-status"] = t
		}
	}

	if v, ok := d.GetOk("weapons_block_strictness_level"); ok || d.HasChange("weapons_block_strictness_level") {
		t, err := expandImageAnalyzerProfileWeaponsBlockStrictnessLevel(d, v, "weapons_block_strictness_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weapons-block-strictness-level"] = t
		}
	}

	if v, ok := d.GetOk("weapons_status"); ok || d.HasChange("weapons_status") {
		t, err := expandImageAnalyzerProfileWeaponsStatus(d, v, "weapons_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weapons-status"] = t
		}
	}

	return &obj, nil
}
