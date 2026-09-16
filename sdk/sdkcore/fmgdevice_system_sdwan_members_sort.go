package forticlient

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// sortSystemSdwanMembersItem contains the parameters for each item
type sortSystemSdwanMembersItem struct {
	seq_num int
}

func getEntryListSystemSdwanMembers(c *FortiSDKClient, inputModel *SortInputModel) (itemList []sortSystemSdwanMembersItem, err error) {
	path := "/pm/config/device/{device}/vdom/{vdom}/system/sdwan/members"
	path, err = replaceParaWithValue(path, inputModel.URLParams)
	params := map[string]interface{}{
		"fields": []string{"seq-num"},
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

		var members []sortSystemSdwanMembersItem
		for _, v := range listTmp {
			c := v.(map[string]interface{})

			members = append(members,
				sortSystemSdwanMembersItem{
					seq_num: int(c["seq-num"].(float64)),
				})
		}

		itemList = members
	}

	return
}

func bEntryListSortedSystemSdwanMembers(itemList []sortSystemSdwanMembersItem, inputModel *SortInputModel) (bsorted bool) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	bsorted = true
	if sortby == "seq-num" {
		for i := 0; i < len(itemList)-1; i++ {
			if sortdirection == "ascending" {
				if itemList[i].seq_num > itemList[i+1].seq_num {
					bsorted = false
					return
				}
			} else if sortdirection == "descending" {
				if itemList[i].seq_num < itemList[i+1].seq_num {
					bsorted = false
					return
				}
			} else if sortdirection == "manual" {
				curItemMap := make(map[string]int)
				for index, item := range itemList {
					curKeyValue := strconv.Itoa(item.seq_num)
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

func moveAfterSystemSdwanMembers(idbefore, idafter int, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	idbefores := strconv.Itoa(idbefore)
	idafters := strconv.Itoa(idafter)
	path := "/pm/config/device/{device}/vdom/{vdom}/system/sdwan/members/"
	path, err = replaceParaWithValue(path, inputModel.URLParams)

	params := make(map[string]interface{})
	path += "/" + idbefores
	params["target"] = idafters
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

func sortEntryListSystemSdwanMembers(itemList []sortSystemSdwanMembersItem, c *FortiSDKClient, inputModel *SortInputModel) (err error) {
	sortby := inputModel.SortBy
	sortdirection := inputModel.SortDirection
	manual_order := inputModel.ManualOrder
	var targetItemOrder []sortSystemSdwanMembersItem
	if sortby == "seq-num" {
		if sortdirection == "ascending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].seq_num < itemList[j].seq_num
			})
			targetItemOrder = itemList
		} else if sortdirection == "descending" {
			sort.Slice(itemList, func(i, j int) bool {
				return itemList[i].seq_num > itemList[j].seq_num
			})
			targetItemOrder = itemList
		} else if sortdirection == "manual" {
			curItemMap := make(map[string]sortSystemSdwanMembersItem)
			for _, item := range itemList {
				curIndex := strconv.Itoa(item.seq_num)
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
		err = moveAfterSystemSdwanMembers(targetItemOrder[i+1].seq_num, targetItemOrder[i].seq_num, c, inputModel)
		if err != nil {
			err = fmt.Errorf("Error move entry: %s", err)
			return
		}
	}

	return nil
}

// CreateUpdateSystemSdwanMembersSort API operation for FortiManager to sort the firewall policies.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateSystemSdwanMembersSort(inputModel *SortInputModel) (err error) {
	itemList, err := getEntryListSystemSdwanMembers(c, inputModel)
	log.Printf("[INFO] Entry ID list: %v", itemList)
	if err != nil {
		err = fmt.Errorf("Fail to get entries before sort: %s", err)
		return
	}

	bsorted := bEntryListSortedSystemSdwanMembers(itemList, inputModel)
	if bsorted == true {
		return
	}

	err = sortEntryListSystemSdwanMembers(itemList, c, inputModel)
	if err != nil {
		err = fmt.Errorf("Error when sort entries: %s", err)
		return
	}

	return
}

// ReadSystemSdwanMembersSort API operation for FortiManager to read the firewall policies sort results
// Returns sort status
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) ReadSystemSdwanMembersSort(inputModel *SortInputModel) (sorted bool, itemMapList []interface{}, err error) {
	itemList, err := getEntryListSystemSdwanMembers(c, inputModel)
	if err != nil {
		err = fmt.Errorf("Fail to read the entries: %s", err)
		return
	}

	bsorted := bEntryListSortedSystemSdwanMembers(itemList, inputModel)
	if bsorted == true {
		sorted = true
		return
	}

	sorted = false
	for _, item := range itemList {
		curItemMap := make(map[string]interface{})
		curItemMap["seq-num"] = item.seq_num
		itemMapList = append(itemMapList, curItemMap)
	}

	return
}
