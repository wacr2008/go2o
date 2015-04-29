/**
 * Copyright 2014 @ ops Inc.
 * name :
 * author : newmin
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package cache

import (
	"bytes"
	"fmt"
	"go2o/src/core/service/dps"
	"strings"
)

func GetShopCheckboxs(partnerId int, chks string) []byte {
	shops := dps.PartnerService.GetShopsOfPartner(partnerId)
	buf := bytes.NewBufferString("")

	if len(chks) == 0 {
		for i, k := range shops {
			buf.WriteString(fmt.Sprintf(
				`<input type="checkbox" value="%d" id="shop%d" field="ApplySubs[%d]"/>
			 	<label for="shop%d">%s</label>`,
				k.Id,
				i,
				i,
				i,
				k.Name,
			))
		}
	} else {
		chks = fmt.Sprintf(",%s,", chks)
		for i, k := range shops {
			if strings.Index(chks, fmt.Sprintf(",%d,", k.Id)) == -1 {
				buf.WriteString(fmt.Sprintf(
					`<input type="checkbox" value="%d" id="shop%d" field="ApplySubs[%d]"/>
			 	<label for="shop%d">%s</label>`,
					k.Id,
					i,
					i,
					i,
					k.Name,
				))
			} else {
				buf.WriteString(fmt.Sprintf(
					`<input type="checkbox" value="%d" id="shop%d" field="ApplySubs[%d]" checked="checked"/>
			 	<label for="shop%d">%s</label>`,
					k.Id,
					i,
					i,
					i,
					k.Name,
				))
			}
		}
	}
	return buf.Bytes()
}

func GetShopsJson(partnerId int) []byte {
	shops := dps.PartnerService.GetShopsOfPartner(partnerId)
	buf := bytes.NewBufferString("[")
	for i, v := range shops {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`{"id":%d,"name":"%s"}`, v.Id, v.Name))
	}
	buf.WriteString("]")
	return buf.Bytes()
}

func GetShopDropList(partnerId int, selected int) []byte {
	buf := bytes.NewBuffer([]byte{})
	shops := dps.PartnerService.GetShopsOfPartner(partnerId)
	for _, v := range shops {
		if v.Id == selected {
			buf.WriteString(fmt.Sprintf(`<option value="%d" selected="selected">%s</option>`, v.Id, v.Name))
		} else {
			buf.WriteString(fmt.Sprintf(`<option value="%d">%s</option>`, v.Id, v.Name))
		}
	}
	return buf.Bytes()
}
