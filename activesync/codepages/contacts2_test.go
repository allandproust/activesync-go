package codepages

import (
	"testing"
)

func Test_Contacts2(t *testing.T) {
	cp := Contacts2()
	checkTagCount(t, 10, cp)
	checkTagExists(t, TAG_CONTACTS2_CUSTOMERID, ID_CONTACTS2_CUSTOMERID, cp)
	checkTagExists(t, TAG_CONTACTS2_GOVERNMENTID, ID_CONTACTS2_GOVERNMENTID, cp)
	checkTagExists(t, TAG_CONTACTS2_IMADDRESS, ID_CONTACTS2_IMADDRESS, cp)
	checkTagExists(t, TAG_CONTACTS2_IMADDRESS2, ID_CONTACTS2_IMADDRESS2, cp)
	checkTagExists(t, TAG_CONTACTS2_IMADDRESS3, ID_CONTACTS2_IMADDRESS3, cp)
	checkTagExists(t, TAG_CONTACTS2_MANAGERNAME, ID_CONTACTS2_MANAGERNAME, cp)
	checkTagExists(t, TAG_CONTACTS2_COMPANYMAINPHONE, ID_CONTACTS2_COMPANYMAINPHONE, cp)
	checkTagExists(t, TAG_CONTACTS2_ACCOUNTNAME, ID_CONTACTS2_ACCOUNTNAME, cp)
	checkTagExists(t, TAG_CONTACTS2_NICKNAME, ID_CONTACTS2_NICKNAME, cp)
	checkTagExists(t, TAG_CONTACTS2_MMS, ID_CONTACTS2_MMS, cp)
}
