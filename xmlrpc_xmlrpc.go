// This file is autogenerated by xmlrpcgen
// do not change it directly!

package core

import (
	"github.com/beevik/etree"
	"github.com/phonkee/go-xmlrpc"
	"strconv"
)

var (
	availableMethodsForSearchService = map[string]bool{
		"search": true,
	}
)

/*
	MethodExists returns whether rpc method is available on service
*/
func (s *SearchService) MethodExists(method string) (ok bool) {
	_, ok = availableMethodsForSearchService[method]
	return
}

/*
	ListMethods returns list of all available methods for given service
*/
func (s *SearchService) ListMethods() []string {
	result := make([]string, 0, len(availableMethodsForSearchService))
	for key := range availableMethodsForSearchService {
		result = append(result, key)
	}
	return result
}

/*
	Dispatch dispatches method on service, do not use this method directly.
	root is params *etree.Element (actually "methodCall/params"
*/
func (s *SearchService) Dispatch(method string, root *etree.Element) (doc *etree.Document, err error) {

	// call appropriate methods
	switch method {
	case "search":
		// Get parameters from xmlrpc request

		v_2 := root.FindElement("param[1]/value")
		if v_2 == nil {
			err = xmlrpc.Errorf(400, "could not find query")
			return
		}

		var query string
		if query, err = xmlrpc.XPathValueGetString(v_2, "query"); err != nil {
			return
		}

		// If following method call fails there are 2 possible reasons:
		// 1. you have either changed method signature or you deleted method. Please re-run "go generate"
		// 2. you have probably found a bug and you should file issue on github.
		// @TODO: add panic recovery that returns error with 500 code

		var result_1 []SearchResult

		result_1, err = s.search(query)

		// create *etree.Document
		doc = etree.NewDocument()
		doc.CreateProcInst("xml", "version=\"1.0\" encoding=\"UTF-8\"")
		methodResponse_3 := doc.CreateElement("methodResponse")
		if err != nil {

			// move this code to error.
			fault_8 := methodResponse_3.CreateElement("fault")

			code_7 := 500

			// Try to cast error to xmlrpc.Error (with code added)
			if code_4, ok_6 := err.(xmlrpc.Error); ok_6 {
				code_7 = code_4.Code()
			}

			struct_5 := fault_8.CreateElement("value").CreateElement("struct")

			member_9 := struct_5.CreateElement("member")
			member_9.CreateElement("name").SetText("faultCode")
			member_9.CreateElement("value").CreateElement("int").SetText(strconv.Itoa(code_7))

			member_10 := struct_5.CreateElement("member")
			member_10.CreateElement("name").SetText("faultString")
			member_10.CreateElement("value").CreateElement("string").SetText(err.Error())

		} else {
			// here is place where we need to hydrate results
			v_11 := methodResponse_3.CreateElement("params").CreateElement("param").CreateElement("value")
			array_data_12 := v_11.CreateElement("array").CreateElement("data")
			for _, item_13 := range result_1 {
				value_14 := array_data_12.CreateElement("value")

				struct_15 := value_14.CreateElement("struct")
				// iterate over struct members

				member_16 := struct_15.CreateElement("member")

				// first create "name" xml element with member name
				member_16.CreateElement("name").SetText("_pypi_ordering")

				value_17 := member_16.CreateElement("value")

				// make shortcut to struct member
				struct_var_18 := item_13._pypi_ordering

				// set value
				value_17.CreateElement("int").SetText(strconv.Itoa(int(struct_var_18)))

				member_19 := struct_15.CreateElement("member")

				// first create "name" xml element with member name
				member_19.CreateElement("name").SetText("version")

				value_20 := member_19.CreateElement("value")

				// make shortcut to struct member
				struct_var_21 := item_13.version

				// set value
				value_20.CreateElement("string").SetText(struct_var_21)

				member_23 := struct_15.CreateElement("member")

				// first create "name" xml element with member name
				member_23.CreateElement("name").SetText("name")

				value_24 := member_23.CreateElement("value")

				// make shortcut to struct member
				struct_var_25 := item_13.name

				// set value
				value_24.CreateElement("string").SetText(struct_var_25)

				member_27 := struct_15.CreateElement("member")

				// first create "name" xml element with member name
				member_27.CreateElement("name").SetText("summary")

				value_28 := member_27.CreateElement("value")

				// make shortcut to struct member
				struct_var_29 := item_13.summary

				// set value
				value_28.CreateElement("string").SetText(struct_var_29)

			}

		}
	default:
		// method not found, this should not happened since we check whether method exists
		err = xmlrpc.ErrMethodNotFound
		return
	}
	return
}