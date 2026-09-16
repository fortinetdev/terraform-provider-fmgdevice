---
subcategory: "System"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_mqttbroker_topiclist"
description: |-
  Configure topic.
---

# fmgdevice_system_mqttbroker_topiclist
Configure topic.

~> This resource is a sub resource for variable `topic_list` of resource `fmgdevice_system_mqttbroker`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Unique integer ID of the entry.
* `topic` - Topic name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System MqttBrokerTopicList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_mqttbroker_topiclist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

