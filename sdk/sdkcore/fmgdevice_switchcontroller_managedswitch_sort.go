package forticlient

import (
	"fmt"
	"log"
	"sort"
)

// sortSwitchControllerManagedSwitchItem contains the parameters for each item
type sortSwitchControllerManagedSwitchItem struct {
	switch_id string
}

func getEntryListSwitchControllerManagedSwitch(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortSwitchControllerManagedSwitchItem, err error) {
	path := "/pm/config/device/{device}/vdom/{vdom}/switch-controller/managed-switch"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"switch-id"},
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "get"
	requestInput.path = path
	requestInput.bodyParams = &params
	requestInput.bMove = true

	listTmp, err := readMove(requestInput)

	if err == nil {
		if listTmp == nil {
			err = fmt.Errorf("cannot get the results from the response")
			return
		}

		var members []sortSwitchControllerManagedSwitchItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortSwitchControllerManagedSwitchItem{
					switch_id: conv2str(c["switch-id"]),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedSwitchControllerManagedSwitch(itemList []sortSwitchControllerManagedSwitchItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "switch-id" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].switch_id > itemList[i+1].switch_id {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].switch_id < itemList[i+1].switch_id {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := item.switch_id
					curItemMap[curKeyValue] = index
				}
				for j := 0; j < len(manual_order)-1; j++ {
					indexL, okL := curItemMap[manual_order[j].(string)]
					indexR, okR := curItemMap[manual_order[j+1].(string)]
					if okL && okR && indexL > indexR {
						bsorted = false
						return
					}
				}
			}
		}
	}

	return
}

func moveAfterSwitchControllerManagedSwitch(idbefore, idafter string, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	path := "/pm/config/device/{device}/vdom/{vdom}/switch-controller/managed-switch/"
	path, err = replaceParaWithValue(path, inputModel.URLParams)

	params := make(map[string]interface{})
	path += "/" + idbefore
	params["target"] = idafter
	params["option"] = "after"

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "move"
	requestInput.path = path
	requestInput.bodyParams = &params
	requestInput.wsParams = inputModel.WSParams
	requestInput.bMove = true

	_, err = createUpdate(requestInput)

	return
}

func sortEntryListSwitchControllerManagedSwitch(itemList []sortSwitchControllerManagedSwitchItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortSwitchControllerManagedSwitchItem
	if sortby == "switch-id" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].switch_id < itemList[j].switch_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].switch_id > itemList[j].switch_id
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortSwitchControllerManagedSwitchItem)
			for _, item := range itemList {
				curIndex := item.switch_id
				curItemMap[curIndex] = item
			}
			for _, val := range manual_order {
				if item, ok := curItemMap[val.(string)]; ok {
					targetItemOrder = append(targetItemOrder, item)
				}
			}
		}
	}

	for i := 0; i < len(targetItemOrder)-1; i++ {
		err = moveAfterSwitchControllerManagedSwitch(targetItemOrder[i+1].switch_id, targetItemOrder[i].switch_id, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateSwitchControllerManagedSwitchSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateSwitchControllerManagedSwitchSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListSwitchControllerManagedSwitch(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedSwitchControllerManagedSwitch(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListSwitchControllerManagedSwitch(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadSwitchControllerManagedSwitchSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadSwitchControllerManagedSwitchSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListSwitchControllerManagedSwitch(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedSwitchControllerManagedSwitch(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["switch-id"] = item.switch_id
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
