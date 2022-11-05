package component

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("/createCustomer -- 200", func() {
	It("create customer with no previous ", func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())

		client := resty.New()
		resp, err := client.R().
			SetQueryString("page=2").
			SetHeader("content-Type", "application/json").
			SetHeader("Accept", "application/json").
			Get("http://reqres.in/api/users")

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		var data map[string]interface{}
		json.Unmarshal(resp.Body(), &data)
		fmt.Println(data["page"])
		fmt.Println(data["per_page"])
		fmt.Println(data["total"])
		fmt.Println(data["total_pages"])
		data1 := data["data"]
		fmt.Println(data1)
		data2 := data1.([]interface{})
		fmt.Println(data2)
		data3 := data2[1]
		data4 := data3.(map[string]interface{})
		fmt.Println(data4)
		fmt.Println(data4["id"])
		fmt.Println(data4["email"])
		fmt.Println(data4["first_name"])
		fmt.Println(data4["last_name"])
	})

	It("create customer with no previous history", Label("Smoke", "Regression"), func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})

	It("fail... create customer from diff entity", Label("SIT", "Regression"), func() {
		fmt.Println("hello, ginkgo!")
		Expect(true).To(BeTrue())
	})

})
