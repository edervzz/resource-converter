package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Give me .cs file")
	fmt.Println("---------------------")

	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
	// text = strings.Replace(text, "\n", "", -1)
	// text = strings.Replace(text, "\r", "", -1)

	contents, err := os.ReadFile(`C:\Users\ederv\gcr\fiinsoft-core-V2\src\Core\FiinSoft.Core.Application\Resources\AccountMessage.cs`)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	textos := strings.Split(string(contents), "\r\n")

	prepareValue := false
	ready := false
	value := ""
	data := ""

	for _, v := range textos {
		var cut string
		v = strings.Trim(v, cut)
		if v == "    /// <summary>" {
			prepareValue = true
			continue
		}
		if prepareValue {
			prepareValue = false
			v = strings.Replace(v, "    /// ", "", -1)
			value = "    <value>" + v + "</value>\r\n    <comment/>\r\n  </data>"
			continue
		}
		if v == "    /// </summary>" {
			ready = true
			continue
		}
		if ready {
			t := strings.Split(v, " = ")
			t[1] = strings.Replace(t[1], ";", " ", -1)
			ready = false
			data = "  <data name=" + t[1] + "xml:space=\"preserve\">"

			fmt.Println(data)
			fmt.Println(value)
			data, value = "", ""
		}

	}

}
