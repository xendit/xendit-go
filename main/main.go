package main

import (
	"fmt"
	"regexp"

	xendit "github.com/xendit/xendit-go"
	form "github.com/xendit/xendit-go/form"
)

func getInvoiceTest(x xendit.Xendit) {

	resp, _ := x.Invoice.GetInvoice("5e007149d5100c05060a4d10")

	fmt.Println(resp)
}

func createInvoiceTest(x xendit.Xendit) {

	data := form.CreateInvoiceData{
		ExternalID:  "invoice-1577331345",
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "Invoice Demo #234",
	}

	resp, _ := x.Invoice.CreateInvoice(data)

	fmt.Println(resp)
}

func expireInvoiceTest(x xendit.Xendit) {

	resp, _ := x.Invoice.ExpireInvoice("5e0454a6f4d38b20d541fa48")

	fmt.Println(resp)
}

func getAllInvoicesTest(x xendit.Xendit) {

	data := form.GetAllInvoicesData{
		Statuses:     []string{"EXPIRED", "SETTLED"},
		Limit:        2,
		CreatedAfter: "2016-02-24T23:48:36.697Z",
	}

	resp, _ := x.Invoice.GetAllInvoices(data)

	fmt.Println(resp)
}

func regSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

func main() {
	x := xendit.New("", "xnd_development_REt02KJzkM6AootfKnDrMw1Sse4LlzEDHfKzXoBocqIEiH4bqjHUJXbl6Cfaab", "")

	getAllInvoicesTest(x)
	// fmt.Println(regSplit("{\"statuses\":[\"EXPIRED\",\"SETTLED\"],\"limit\":2", "[^\\]]+(?![^\\[]*\)"))
}
