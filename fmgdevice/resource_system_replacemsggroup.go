// Copyright 2024 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: <i>This object will be purged after policy copy and install.</i> Configure replacement message groups.

package fmgdevice

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemReplacemsgGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgGroupCreate,
		Read:   resourceSystemReplacemsgGroupRead,
		Update: resourceSystemReplacemsgGroupUpdate,
		Delete: resourceSystemReplacemsgGroupDelete,

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
			"admin": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"alertmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"auth": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"automation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"custom_message": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"device_detection_portal": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"fortiguard_wf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"icap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"mail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"nac_quar": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
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
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"spam": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"traffic_quota": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"utm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"webproxy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemReplacemsgGroupCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemReplacemsgGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgGroup resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemReplacemsgGroup(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemReplacemsgGroup(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemReplacemsgGroup resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateSystemReplacemsgGroup(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemReplacemsgGroup resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceSystemReplacemsgGroupRead(d, m)
}

func resourceSystemReplacemsgGroupUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemReplacemsgGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroup resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemReplacemsgGroup(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceSystemReplacemsgGroupRead(d, m)
}

func resourceSystemReplacemsgGroupDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteSystemReplacemsgGroup(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgGroupRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemReplacemsgGroup(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemReplacemsgGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgGroup resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgGroupAdmin(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupAdminBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Admin-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupAdminFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Admin-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupAdminHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Admin-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupAdminMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Admin-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupAdminBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdminFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdminHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdminMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupAlertmailBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Alertmail-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupAlertmailFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Alertmail-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupAlertmailHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Alertmail-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemReplacemsgGroupAlertmailId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Alertmail-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupAlertmailMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Alertmail-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupAlertmailBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuth(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupAuthBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Auth-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupAuthFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Auth-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupAuthHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Auth-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupAuthMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Auth-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupAuthBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuthFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuthHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuthMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupAutomationBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Automation-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupAutomationFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Automation-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupAutomationHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Automation-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupAutomationMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Automation-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupAutomationBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomationFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomationHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomationMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupCustomMessageBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-CustomMessage-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupCustomMessageFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-CustomMessage-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupCustomMessageHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-CustomMessage-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupCustomMessageMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-CustomMessage-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupCustomMessageBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessageFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessageHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessageMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupDeviceDetectionPortalBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-DeviceDetectionPortal-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupDeviceDetectionPortalFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-DeviceDetectionPortal-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupDeviceDetectionPortalHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-DeviceDetectionPortal-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupDeviceDetectionPortalMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-DeviceDetectionPortal-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupFortiguardWfBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-FortiguardWf-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupFortiguardWfFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-FortiguardWf-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupFortiguardWfHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-FortiguardWf-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupFortiguardWfMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-FortiguardWf-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupFortiguardWfBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWfFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWfHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWfMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupFtpBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Ftp-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupFtpFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Ftp-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupFtpHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Ftp-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupFtpMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Ftp-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupFtpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupHttpBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Http-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupHttpFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Http-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupHttpHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Http-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupHttpMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Http-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupHttpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupIcapBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Icap-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupIcapFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Icap-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupIcapHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Icap-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupIcapMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Icap-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupIcapBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcapFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcapHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcapMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupMailBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Mail-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupMailFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Mail-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupMailHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Mail-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupMailMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Mail-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupMailBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMailFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuar(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupNacQuarBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-NacQuar-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupNacQuarFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-NacQuar-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupNacQuarHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-NacQuar-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenSystemReplacemsgGroupNacQuarId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-NacQuar-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupNacQuarMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-NacQuar-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupNacQuarBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupNntpBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Nntp-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupNntpFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Nntp-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupNntpHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Nntp-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupNntpMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Nntp-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupNntpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupSpamBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Spam-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupSpamFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Spam-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupSpamHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Spam-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupSpamMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Spam-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupSpamBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpn(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupSslvpnBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Sslvpn-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupSslvpnFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Sslvpn-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupSslvpnHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Sslvpn-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupSslvpnMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Sslvpn-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupSslvpnBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpnFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpnHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpnMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuota(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupTrafficQuotaBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-TrafficQuota-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupTrafficQuotaFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-TrafficQuota-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupTrafficQuotaHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-TrafficQuota-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupTrafficQuotaMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-TrafficQuota-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupTrafficQuotaBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuotaFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuotaHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuotaMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtm(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupUtmBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Utm-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupUtmFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Utm-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupUtmHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Utm-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupUtmMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Utm-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupUtmBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtmFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtmHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtmMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			v := flattenSystemReplacemsgGroupWebproxyBuffer(i["buffer"], d, pre_append)
			tmp["buffer"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Webproxy-Buffer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			v := flattenSystemReplacemsgGroupWebproxyFormat(i["format"], d, pre_append)
			tmp["format"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Webproxy-Format")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			v := flattenSystemReplacemsgGroupWebproxyHeader(i["header"], d, pre_append)
			tmp["header"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Webproxy-Header")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			v := flattenSystemReplacemsgGroupWebproxyMsgType(i["msg-type"], d, pre_append)
			tmp["msg_type"] = fortiAPISubPartPatch(v, "SystemReplacemsgGroup-Webproxy-MsgType")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenSystemReplacemsgGroupWebproxyBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxyFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxyHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxyMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("admin", flattenSystemReplacemsgGroupAdmin(o["admin"], d, "admin")); err != nil {
			if vv, ok := fortiAPIPatch(o["admin"], "SystemReplacemsgGroup-Admin"); ok {
				if err = d.Set("admin", vv); err != nil {
					return fmt.Errorf("Error reading admin: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading admin: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("admin"); ok {
			if err = d.Set("admin", flattenSystemReplacemsgGroupAdmin(o["admin"], d, "admin")); err != nil {
				if vv, ok := fortiAPIPatch(o["admin"], "SystemReplacemsgGroup-Admin"); ok {
					if err = d.Set("admin", vv); err != nil {
						return fmt.Errorf("Error reading admin: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading admin: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("alertmail", flattenSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail")); err != nil {
			if vv, ok := fortiAPIPatch(o["alertmail"], "SystemReplacemsgGroup-Alertmail"); ok {
				if err = d.Set("alertmail", vv); err != nil {
					return fmt.Errorf("Error reading alertmail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading alertmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alertmail"); ok {
			if err = d.Set("alertmail", flattenSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail")); err != nil {
				if vv, ok := fortiAPIPatch(o["alertmail"], "SystemReplacemsgGroup-Alertmail"); ok {
					if err = d.Set("alertmail", vv); err != nil {
						return fmt.Errorf("Error reading alertmail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading alertmail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("auth", flattenSystemReplacemsgGroupAuth(o["auth"], d, "auth")); err != nil {
			if vv, ok := fortiAPIPatch(o["auth"], "SystemReplacemsgGroup-Auth"); ok {
				if err = d.Set("auth", vv); err != nil {
					return fmt.Errorf("Error reading auth: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth"); ok {
			if err = d.Set("auth", flattenSystemReplacemsgGroupAuth(o["auth"], d, "auth")); err != nil {
				if vv, ok := fortiAPIPatch(o["auth"], "SystemReplacemsgGroup-Auth"); ok {
					if err = d.Set("auth", vv); err != nil {
						return fmt.Errorf("Error reading auth: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading auth: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("automation", flattenSystemReplacemsgGroupAutomation(o["automation"], d, "automation")); err != nil {
			if vv, ok := fortiAPIPatch(o["automation"], "SystemReplacemsgGroup-Automation"); ok {
				if err = d.Set("automation", vv); err != nil {
					return fmt.Errorf("Error reading automation: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading automation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("automation"); ok {
			if err = d.Set("automation", flattenSystemReplacemsgGroupAutomation(o["automation"], d, "automation")); err != nil {
				if vv, ok := fortiAPIPatch(o["automation"], "SystemReplacemsgGroup-Automation"); ok {
					if err = d.Set("automation", vv); err != nil {
						return fmt.Errorf("Error reading automation: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading automation: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenSystemReplacemsgGroupComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemReplacemsgGroup-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_message", flattenSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message")); err != nil {
			if vv, ok := fortiAPIPatch(o["custom-message"], "SystemReplacemsgGroup-CustomMessage"); ok {
				if err = d.Set("custom_message", vv); err != nil {
					return fmt.Errorf("Error reading custom_message: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading custom_message: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_message"); ok {
			if err = d.Set("custom_message", flattenSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message")); err != nil {
				if vv, ok := fortiAPIPatch(o["custom-message"], "SystemReplacemsgGroup-CustomMessage"); ok {
					if err = d.Set("custom_message", vv); err != nil {
						return fmt.Errorf("Error reading custom_message: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading custom_message: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("device_detection_portal", flattenSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal")); err != nil {
			if vv, ok := fortiAPIPatch(o["device-detection-portal"], "SystemReplacemsgGroup-DeviceDetectionPortal"); ok {
				if err = d.Set("device_detection_portal", vv); err != nil {
					return fmt.Errorf("Error reading device_detection_portal: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device_detection_portal: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_detection_portal"); ok {
			if err = d.Set("device_detection_portal", flattenSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal")); err != nil {
				if vv, ok := fortiAPIPatch(o["device-detection-portal"], "SystemReplacemsgGroup-DeviceDetectionPortal"); ok {
					if err = d.Set("device_detection_portal", vv); err != nil {
						return fmt.Errorf("Error reading device_detection_portal: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device_detection_portal: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("fortiguard_wf", flattenSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf")); err != nil {
			if vv, ok := fortiAPIPatch(o["fortiguard-wf"], "SystemReplacemsgGroup-FortiguardWf"); ok {
				if err = d.Set("fortiguard_wf", vv); err != nil {
					return fmt.Errorf("Error reading fortiguard_wf: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading fortiguard_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fortiguard_wf"); ok {
			if err = d.Set("fortiguard_wf", flattenSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf")); err != nil {
				if vv, ok := fortiAPIPatch(o["fortiguard-wf"], "SystemReplacemsgGroup-FortiguardWf"); ok {
					if err = d.Set("fortiguard_wf", vv); err != nil {
						return fmt.Errorf("Error reading fortiguard_wf: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading fortiguard_wf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenSystemReplacemsgGroupFtp(o["ftp"], d, "ftp")); err != nil {
			if vv, ok := fortiAPIPatch(o["ftp"], "SystemReplacemsgGroup-Ftp"); ok {
				if err = d.Set("ftp", vv); err != nil {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenSystemReplacemsgGroupFtp(o["ftp"], d, "ftp")); err != nil {
				if vv, ok := fortiAPIPatch(o["ftp"], "SystemReplacemsgGroup-Ftp"); ok {
					if err = d.Set("ftp", vv); err != nil {
						return fmt.Errorf("Error reading ftp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if err = d.Set("group_type", flattenSystemReplacemsgGroupGroupType(o["group-type"], d, "group_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["group-type"], "SystemReplacemsgGroup-GroupType"); ok {
			if err = d.Set("group_type", vv); err != nil {
				return fmt.Errorf("Error reading group_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenSystemReplacemsgGroupHttp(o["http"], d, "http")); err != nil {
			if vv, ok := fortiAPIPatch(o["http"], "SystemReplacemsgGroup-Http"); ok {
				if err = d.Set("http", vv); err != nil {
					return fmt.Errorf("Error reading http: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenSystemReplacemsgGroupHttp(o["http"], d, "http")); err != nil {
				if vv, ok := fortiAPIPatch(o["http"], "SystemReplacemsgGroup-Http"); ok {
					if err = d.Set("http", vv); err != nil {
						return fmt.Errorf("Error reading http: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("icap", flattenSystemReplacemsgGroupIcap(o["icap"], d, "icap")); err != nil {
			if vv, ok := fortiAPIPatch(o["icap"], "SystemReplacemsgGroup-Icap"); ok {
				if err = d.Set("icap", vv); err != nil {
					return fmt.Errorf("Error reading icap: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading icap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap"); ok {
			if err = d.Set("icap", flattenSystemReplacemsgGroupIcap(o["icap"], d, "icap")); err != nil {
				if vv, ok := fortiAPIPatch(o["icap"], "SystemReplacemsgGroup-Icap"); ok {
					if err = d.Set("icap", vv); err != nil {
						return fmt.Errorf("Error reading icap: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading icap: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mail", flattenSystemReplacemsgGroupMail(o["mail"], d, "mail")); err != nil {
			if vv, ok := fortiAPIPatch(o["mail"], "SystemReplacemsgGroup-Mail"); ok {
				if err = d.Set("mail", vv); err != nil {
					return fmt.Errorf("Error reading mail: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading mail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mail"); ok {
			if err = d.Set("mail", flattenSystemReplacemsgGroupMail(o["mail"], d, "mail")); err != nil {
				if vv, ok := fortiAPIPatch(o["mail"], "SystemReplacemsgGroup-Mail"); ok {
					if err = d.Set("mail", vv); err != nil {
						return fmt.Errorf("Error reading mail: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading mail: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("nac_quar", flattenSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
			if vv, ok := fortiAPIPatch(o["nac-quar"], "SystemReplacemsgGroup-NacQuar"); ok {
				if err = d.Set("nac_quar", vv); err != nil {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nac_quar: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_quar"); ok {
			if err = d.Set("nac_quar", flattenSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
				if vv, ok := fortiAPIPatch(o["nac-quar"], "SystemReplacemsgGroup-NacQuar"); ok {
					if err = d.Set("nac_quar", vv); err != nil {
						return fmt.Errorf("Error reading nac_quar: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemReplacemsgGroupName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "SystemReplacemsgGroup-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nntp", flattenSystemReplacemsgGroupNntp(o["nntp"], d, "nntp")); err != nil {
			if vv, ok := fortiAPIPatch(o["nntp"], "SystemReplacemsgGroup-Nntp"); ok {
				if err = d.Set("nntp", vv); err != nil {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenSystemReplacemsgGroupNntp(o["nntp"], d, "nntp")); err != nil {
				if vv, ok := fortiAPIPatch(o["nntp"], "SystemReplacemsgGroup-Nntp"); ok {
					if err = d.Set("nntp", vv); err != nil {
						return fmt.Errorf("Error reading nntp: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("spam", flattenSystemReplacemsgGroupSpam(o["spam"], d, "spam")); err != nil {
			if vv, ok := fortiAPIPatch(o["spam"], "SystemReplacemsgGroup-Spam"); ok {
				if err = d.Set("spam", vv); err != nil {
					return fmt.Errorf("Error reading spam: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading spam: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("spam"); ok {
			if err = d.Set("spam", flattenSystemReplacemsgGroupSpam(o["spam"], d, "spam")); err != nil {
				if vv, ok := fortiAPIPatch(o["spam"], "SystemReplacemsgGroup-Spam"); ok {
					if err = d.Set("spam", vv); err != nil {
						return fmt.Errorf("Error reading spam: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading spam: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("sslvpn", flattenSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
			if vv, ok := fortiAPIPatch(o["sslvpn"], "SystemReplacemsgGroup-Sslvpn"); ok {
				if err = d.Set("sslvpn", vv); err != nil {
					return fmt.Errorf("Error reading sslvpn: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading sslvpn: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sslvpn"); ok {
			if err = d.Set("sslvpn", flattenSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
				if vv, ok := fortiAPIPatch(o["sslvpn"], "SystemReplacemsgGroup-Sslvpn"); ok {
					if err = d.Set("sslvpn", vv); err != nil {
						return fmt.Errorf("Error reading sslvpn: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading sslvpn: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("traffic_quota", flattenSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota")); err != nil {
			if vv, ok := fortiAPIPatch(o["traffic-quota"], "SystemReplacemsgGroup-TrafficQuota"); ok {
				if err = d.Set("traffic_quota", vv); err != nil {
					return fmt.Errorf("Error reading traffic_quota: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading traffic_quota: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("traffic_quota"); ok {
			if err = d.Set("traffic_quota", flattenSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota")); err != nil {
				if vv, ok := fortiAPIPatch(o["traffic-quota"], "SystemReplacemsgGroup-TrafficQuota"); ok {
					if err = d.Set("traffic_quota", vv); err != nil {
						return fmt.Errorf("Error reading traffic_quota: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading traffic_quota: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("utm", flattenSystemReplacemsgGroupUtm(o["utm"], d, "utm")); err != nil {
			if vv, ok := fortiAPIPatch(o["utm"], "SystemReplacemsgGroup-Utm"); ok {
				if err = d.Set("utm", vv); err != nil {
					return fmt.Errorf("Error reading utm: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading utm: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("utm"); ok {
			if err = d.Set("utm", flattenSystemReplacemsgGroupUtm(o["utm"], d, "utm")); err != nil {
				if vv, ok := fortiAPIPatch(o["utm"], "SystemReplacemsgGroup-Utm"); ok {
					if err = d.Set("utm", vv); err != nil {
						return fmt.Errorf("Error reading utm: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading utm: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("webproxy", flattenSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy")); err != nil {
			if vv, ok := fortiAPIPatch(o["webproxy"], "SystemReplacemsgGroup-Webproxy"); ok {
				if err = d.Set("webproxy", vv); err != nil {
					return fmt.Errorf("Error reading webproxy: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading webproxy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("webproxy"); ok {
			if err = d.Set("webproxy", flattenSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy")); err != nil {
				if vv, ok := fortiAPIPatch(o["webproxy"], "SystemReplacemsgGroup-Webproxy"); ok {
					if err = d.Set("webproxy", vv); err != nil {
						return fmt.Errorf("Error reading webproxy: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading webproxy: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemReplacemsgGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemReplacemsgGroupAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAdminBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupAdminFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupAdminHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAdminMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAdminBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdminFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdminHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdminMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAlertmailBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupAlertmailFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupAlertmailHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemReplacemsgGroupAlertmailId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAlertmailMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAlertmailBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAuthBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupAuthFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupAuthHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAuthMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAuthBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuthFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuthHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuthMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAutomationBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupAutomationFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupAutomationHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAutomationMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAutomationBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomationFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomationHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomationMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupCustomMessageBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupCustomMessageFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupCustomMessageHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupCustomMessageMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupCustomMessageBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessageFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessageHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessageMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupFortiguardWfBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupFortiguardWfFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupFortiguardWfHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupFortiguardWfMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupFortiguardWfBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWfFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWfHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWfMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupFtpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupFtpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupFtpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupFtpMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupFtpBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtpFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtpMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupGroupType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupHttpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupHttpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupHttpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupHttpMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupHttpBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttpFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttpMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupIcapBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupIcapFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupIcapHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupIcapMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupIcapBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcapFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcapHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcapMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupMailBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupMailFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupMailHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupMailMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupMailBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMailFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMailHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMailMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuar(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupNacQuarBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupNacQuarFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupNacQuarHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["id"], _ = expandSystemReplacemsgGroupNacQuarId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupNacQuarMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupNacQuarBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupNntpBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupNntpFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupNntpHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupNntpMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupNntpBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntpFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntpHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntpMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupSpamBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupSpamFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupSpamHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupSpamMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupSpamBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupSslvpnBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupSslvpnFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupSslvpnHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupSslvpnMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupSslvpnBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpnFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpnHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpnMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupTrafficQuotaBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupTrafficQuotaFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupTrafficQuotaHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupTrafficQuotaMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupTrafficQuotaBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuotaFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuotaHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuotaMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupUtmBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupUtmFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupUtmHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupUtmMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupUtmBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtmFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtmHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtmMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["buffer"], _ = expandSystemReplacemsgGroupWebproxyBuffer(d, i["buffer"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["format"], _ = expandSystemReplacemsgGroupWebproxyFormat(d, i["format"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["header"], _ = expandSystemReplacemsgGroupWebproxyHeader(d, i["header"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupWebproxyMsgType(d, i["msg_type"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupWebproxyBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxyFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxyHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxyMsgType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("admin"); ok || d.HasChange("admin") {
		t, err := expandSystemReplacemsgGroupAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("alertmail"); ok || d.HasChange("alertmail") {
		t, err := expandSystemReplacemsgGroupAlertmail(d, v, "alertmail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alertmail"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandSystemReplacemsgGroupAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("automation"); ok || d.HasChange("automation") {
		t, err := expandSystemReplacemsgGroupAutomation(d, v, "automation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["automation"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSystemReplacemsgGroupComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("custom_message"); ok || d.HasChange("custom_message") {
		t, err := expandSystemReplacemsgGroupCustomMessage(d, v, "custom_message")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-message"] = t
		}
	}

	if v, ok := d.GetOk("device_detection_portal"); ok || d.HasChange("device_detection_portal") {
		t, err := expandSystemReplacemsgGroupDeviceDetectionPortal(d, v, "device_detection_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-detection-portal"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_wf"); ok || d.HasChange("fortiguard_wf") {
		t, err := expandSystemReplacemsgGroupFortiguardWf(d, v, "fortiguard_wf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-wf"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok || d.HasChange("ftp") {
		t, err := expandSystemReplacemsgGroupFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok || d.HasChange("group_type") {
		t, err := expandSystemReplacemsgGroupGroupType(d, v, "group_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandSystemReplacemsgGroupHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("icap"); ok || d.HasChange("icap") {
		t, err := expandSystemReplacemsgGroupIcap(d, v, "icap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap"] = t
		}
	}

	if v, ok := d.GetOk("mail"); ok || d.HasChange("mail") {
		t, err := expandSystemReplacemsgGroupMail(d, v, "mail")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok || d.HasChange("nac_quar") {
		t, err := expandSystemReplacemsgGroupNacQuar(d, v, "nac_quar")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandSystemReplacemsgGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok || d.HasChange("nntp") {
		t, err := expandSystemReplacemsgGroupNntp(d, v, "nntp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("spam"); ok || d.HasChange("spam") {
		t, err := expandSystemReplacemsgGroupSpam(d, v, "spam")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok || d.HasChange("sslvpn") {
		t, err := expandSystemReplacemsgGroupSslvpn(d, v, "sslvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("traffic_quota"); ok || d.HasChange("traffic_quota") {
		t, err := expandSystemReplacemsgGroupTrafficQuota(d, v, "traffic_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-quota"] = t
		}
	}

	if v, ok := d.GetOk("utm"); ok || d.HasChange("utm") {
		t, err := expandSystemReplacemsgGroupUtm(d, v, "utm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm"] = t
		}
	}

	if v, ok := d.GetOk("webproxy"); ok || d.HasChange("webproxy") {
		t, err := expandSystemReplacemsgGroupWebproxy(d, v, "webproxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy"] = t
		}
	}

	return &obj, nil
}
