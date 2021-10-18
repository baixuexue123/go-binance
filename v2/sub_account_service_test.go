package binance

import (
	"encoding/json"
	"testing"
)

func TestSubAccountJSON(t *testing.T) {
	s := `
{
    "subAccounts":[
        {
            "email":"testsub@gmail.com",
            "isFreeze":false,
            "createTime":1544433328000
        },
        {
            "email":"virtual@oxebmvfonoemail.com",
            "isFreeze":false,
            "createTime":1544433328000
        }
    ]
}
`
	res := new(SubAccountList)
	err := json.Unmarshal([]byte(s), res)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	t.Log(len(res.SubAccounts))
	t.Log(res.SubAccounts)
}
