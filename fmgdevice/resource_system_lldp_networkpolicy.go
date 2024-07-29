// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure LLDP network policy.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLldpNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLldpNetworkPolicyCreate,
		Read:   resourceSystemLldpNetworkPolicyRead,
		Update: resourceSystemLldpNetworkPolicyUpdate,
		Delete: resourceSystemLldpNetworkPolicyDelete,

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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guest": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"guest_voice_signaling": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"softphone": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"streaming_video": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"video_conferencing": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"video_signaling": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"voice": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"voice_signaling": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemLldpNetworkPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSystemLldpNetworkPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLldpNetworkPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLldpNetworkPolicy(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemLldpNetworkPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemLldpNetworkPolicyRead(d, m)
}

func resourceSystemLldpNetworkPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	obj, err := getObjectSystemLldpNetworkPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLldpNetworkPolicy(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemLldpNetworkPolicyRead(d, m)
}

func resourceSystemLldpNetworkPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

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

	err = c.DeleteSystemLldpNetworkPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLldpNetworkPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLldpNetworkPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemLldpNetworkPolicy(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLldpNetworkPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemLldpNetworkPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyGuestDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyGuestPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyGuestStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyGuestTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyGuestVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyGuestDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicySoftphoneDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicySoftphonePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicySoftphoneStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicySoftphoneTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicySoftphoneVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicySoftphoneDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphonePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphoneStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphoneTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphoneVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyStreamingVideoDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyStreamingVideoPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyStreamingVideoStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyStreamingVideoTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyStreamingVideoVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyStreamingVideoDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencing(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVideoConferencingDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVideoConferencingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVideoConferencingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVideoConferencingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVideoConferencingVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVideoConferencingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVideoSignalingDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVideoSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVideoSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVideoSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVideoSignalingVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVideoSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVoiceDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVoicePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVoiceStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVoiceTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVoiceVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVoiceDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoicePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVoiceSignalingDscp(i["dscp"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVoiceSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVoiceSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVoiceSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVoiceSignalingVlan(i["vlan"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVoiceSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLldpNetworkPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("comment", flattenSystemLldpNetworkPolicyComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemLldpNetworkPolicy-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("guest", flattenSystemLldpNetworkPolicyGuest(o["guest"], d, "guest")); err != nil {
			if vv, ok := fortiAPIPatch(o["guest"], "SystemLldpNetworkPolicy-Guest"); ok {
				if err = d.Set("guest", vv); err != nil {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading guest: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest"); ok {
			if err = d.Set("guest", flattenSystemLldpNetworkPolicyGuest(o["guest"], d, "guest")); err != nil {
				if vv, ok := fortiAPIPatch(o["guest"], "SystemLldpNetworkPolicy-Guest"); ok {
					if err = d.Set("guest", vv); err != nil {
						return fmt.Errorf("Error reading guest: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("guest_voice_signaling", flattenSystemLldpNetworkPolicyGuestVoiceSignaling(o["guest-voice-signaling"], d, "guest_voice_signaling")); err != nil {
			if vv, ok := fortiAPIPatch(o["guest-voice-signaling"], "SystemLldpNetworkPolicy-GuestVoiceSignaling"); ok {
				if err = d.Set("guest_voice_signaling", vv); err != nil {
					return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest_voice_signaling"); ok {
			if err = d.Set("guest_voice_signaling", flattenSystemLldpNetworkPolicyGuestVoiceSignaling(o["guest-voice-signaling"], d, "guest_voice_signaling")); err != nil {
				if vv, ok := fortiAPIPatch(o["guest-voice-signaling"], "SystemLldpNetworkPolicy-GuestVoiceSignaling"); ok {
					if err = d.Set("guest_voice_signaling", vv); err != nil {
						return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemLldpNetworkPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemLldpNetworkPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("softphone", flattenSystemLldpNetworkPolicySoftphone(o["softphone"], d, "softphone")); err != nil {
			if vv, ok := fortiAPIPatch(o["softphone"], "SystemLldpNetworkPolicy-Softphone"); ok {
				if err = d.Set("softphone", vv); err != nil {
					return fmt.Errorf("Error reading softphone: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading softphone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("softphone"); ok {
			if err = d.Set("softphone", flattenSystemLldpNetworkPolicySoftphone(o["softphone"], d, "softphone")); err != nil {
				if vv, ok := fortiAPIPatch(o["softphone"], "SystemLldpNetworkPolicy-Softphone"); ok {
					if err = d.Set("softphone", vv); err != nil {
						return fmt.Errorf("Error reading softphone: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading softphone: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("streaming_video", flattenSystemLldpNetworkPolicyStreamingVideo(o["streaming-video"], d, "streaming_video")); err != nil {
			if vv, ok := fortiAPIPatch(o["streaming-video"], "SystemLldpNetworkPolicy-StreamingVideo"); ok {
				if err = d.Set("streaming_video", vv); err != nil {
					return fmt.Errorf("Error reading streaming_video: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading streaming_video: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("streaming_video"); ok {
			if err = d.Set("streaming_video", flattenSystemLldpNetworkPolicyStreamingVideo(o["streaming-video"], d, "streaming_video")); err != nil {
				if vv, ok := fortiAPIPatch(o["streaming-video"], "SystemLldpNetworkPolicy-StreamingVideo"); ok {
					if err = d.Set("streaming_video", vv); err != nil {
						return fmt.Errorf("Error reading streaming_video: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading streaming_video: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("video_conferencing", flattenSystemLldpNetworkPolicyVideoConferencing(o["video-conferencing"], d, "video_conferencing")); err != nil {
			if vv, ok := fortiAPIPatch(o["video-conferencing"], "SystemLldpNetworkPolicy-VideoConferencing"); ok {
				if err = d.Set("video_conferencing", vv); err != nil {
					return fmt.Errorf("Error reading video_conferencing: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading video_conferencing: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("video_conferencing"); ok {
			if err = d.Set("video_conferencing", flattenSystemLldpNetworkPolicyVideoConferencing(o["video-conferencing"], d, "video_conferencing")); err != nil {
				if vv, ok := fortiAPIPatch(o["video-conferencing"], "SystemLldpNetworkPolicy-VideoConferencing"); ok {
					if err = d.Set("video_conferencing", vv); err != nil {
						return fmt.Errorf("Error reading video_conferencing: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading video_conferencing: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("video_signaling", flattenSystemLldpNetworkPolicyVideoSignaling(o["video-signaling"], d, "video_signaling")); err != nil {
			if vv, ok := fortiAPIPatch(o["video-signaling"], "SystemLldpNetworkPolicy-VideoSignaling"); ok {
				if err = d.Set("video_signaling", vv); err != nil {
					return fmt.Errorf("Error reading video_signaling: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading video_signaling: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("video_signaling"); ok {
			if err = d.Set("video_signaling", flattenSystemLldpNetworkPolicyVideoSignaling(o["video-signaling"], d, "video_signaling")); err != nil {
				if vv, ok := fortiAPIPatch(o["video-signaling"], "SystemLldpNetworkPolicy-VideoSignaling"); ok {
					if err = d.Set("video_signaling", vv); err != nil {
						return fmt.Errorf("Error reading video_signaling: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading video_signaling: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("voice", flattenSystemLldpNetworkPolicyVoice(o["voice"], d, "voice")); err != nil {
			if vv, ok := fortiAPIPatch(o["voice"], "SystemLldpNetworkPolicy-Voice"); ok {
				if err = d.Set("voice", vv); err != nil {
					return fmt.Errorf("Error reading voice: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading voice: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("voice"); ok {
			if err = d.Set("voice", flattenSystemLldpNetworkPolicyVoice(o["voice"], d, "voice")); err != nil {
				if vv, ok := fortiAPIPatch(o["voice"], "SystemLldpNetworkPolicy-Voice"); ok {
					if err = d.Set("voice", vv); err != nil {
						return fmt.Errorf("Error reading voice: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading voice: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("voice_signaling", flattenSystemLldpNetworkPolicyVoiceSignaling(o["voice-signaling"], d, "voice_signaling")); err != nil {
			if vv, ok := fortiAPIPatch(o["voice-signaling"], "SystemLldpNetworkPolicy-VoiceSignaling"); ok {
				if err = d.Set("voice_signaling", vv); err != nil {
					return fmt.Errorf("Error reading voice_signaling: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading voice_signaling: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("voice_signaling"); ok {
			if err = d.Set("voice_signaling", flattenSystemLldpNetworkPolicyVoiceSignaling(o["voice-signaling"], d, "voice_signaling")); err != nil {
				if vv, ok := fortiAPIPatch(o["voice-signaling"], "SystemLldpNetworkPolicy-VoiceSignaling"); ok {
					if err = d.Set("voice_signaling", vv); err != nil {
						return fmt.Errorf("Error reading voice_signaling: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading voice_signaling: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemLldpNetworkPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLldpNetworkPolicyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyGuestDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyGuestPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyGuestStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyGuestTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyGuestVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyGuestDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicySoftphoneDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicySoftphonePriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicySoftphoneStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicySoftphoneTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicySoftphoneVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicySoftphoneDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphonePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphoneStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphoneTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphoneVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyStreamingVideoDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyStreamingVideoPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyStreamingVideoStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyStreamingVideoTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyStreamingVideoVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVideoConferencingDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyVideoConferencingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyVideoConferencingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyVideoConferencingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVideoConferencingVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignaling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVideoSignalingDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyVideoSignalingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyVideoSignalingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyVideoSignalingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVideoSignalingVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVoiceDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyVoicePriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyVoiceStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyVoiceTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVoiceVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVoiceDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoicePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignaling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVoiceSignalingDscp(d, i["dscp"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["priority"], _ = expandSystemLldpNetworkPolicyVoiceSignalingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["status"], _ = expandSystemLldpNetworkPolicyVoiceSignalingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["tag"], _ = expandSystemLldpNetworkPolicyVoiceSignalingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVoiceSignalingVlan(d, i["vlan"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLldpNetworkPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSystemLldpNetworkPolicyComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok || d.HasChange("guest") {
		t, err := expandSystemLldpNetworkPolicyGuest(d, v, "guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	if v, ok := d.GetOk("guest_voice_signaling"); ok || d.HasChange("guest_voice_signaling") {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignaling(d, v, "guest_voice_signaling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-voice-signaling"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemLldpNetworkPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("softphone"); ok || d.HasChange("softphone") {
		t, err := expandSystemLldpNetworkPolicySoftphone(d, v, "softphone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["softphone"] = t
		}
	}

	if v, ok := d.GetOk("streaming_video"); ok || d.HasChange("streaming_video") {
		t, err := expandSystemLldpNetworkPolicyStreamingVideo(d, v, "streaming_video")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-video"] = t
		}
	}

	if v, ok := d.GetOk("video_conferencing"); ok || d.HasChange("video_conferencing") {
		t, err := expandSystemLldpNetworkPolicyVideoConferencing(d, v, "video_conferencing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["video-conferencing"] = t
		}
	}

	if v, ok := d.GetOk("video_signaling"); ok || d.HasChange("video_signaling") {
		t, err := expandSystemLldpNetworkPolicyVideoSignaling(d, v, "video_signaling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["video-signaling"] = t
		}
	}

	if v, ok := d.GetOk("voice"); ok || d.HasChange("voice") {
		t, err := expandSystemLldpNetworkPolicyVoice(d, v, "voice")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice"] = t
		}
	}

	if v, ok := d.GetOk("voice_signaling"); ok || d.HasChange("voice_signaling") {
		t, err := expandSystemLldpNetworkPolicyVoiceSignaling(d, v, "voice_signaling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-signaling"] = t
		}
	}

	return &obj, nil
}
