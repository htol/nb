package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/TylerBrock/colorjson"
)

func main() {
	cmd := ""
	query := ""
	grep := ""
	requestURL := ""

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("\n\tnb [flags] cmd query [grep]\n\tcmd=[ip, pref, agg, dev, vm]")
		fmt.Println("Example:\n\tnb vm rr host")
        fmt.Println("\tIt means: find 'rr' in vm's and grep line with 'host' in it")
		os.Exit(0)
	}

    nbAddr := flag.String("host", "https://netbox", "NetBox address")
	flag.Parse()

	switch flag.NArg() {
	case 2:
		cmd = flag.Arg(0)
		query = flag.Arg(1)
	case 3:
		cmd = flag.Arg(0)
		query = flag.Arg(1)
		grep = flag.Arg(2)
	default:
		flag.Usage()
	}

	switch cmd {
	case "ip":
		requestURL = fmt.Sprintf("%s/api/ipam/ip-addresses/?limit=0&q=%s", *nbAddr, query)
	case "pref":
		requestURL = fmt.Sprintf("%s/api/ipam/prefixes/?limit=0&q=%s", *nbAddr, query)
	case "agg":
		requestURL = fmt.Sprintf("%s/api/ipam/aggregates/?limit=0&q=%s", *nbAddr, query)
	case "dev":
		requestURL = fmt.Sprintf("%s/api/dcim/devices/?limit=0&q=%s", *nbAddr, query)
	case "vm":
		requestURL = fmt.Sprintf("%s/api/virtualization/virtual-machines/?limit=0&q=%s", *nbAddr, query)
	}

	client := &http.Client{}

	fmt.Println(requestURL)
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	req.Header.Add("Accept", "application/json; indent=4")

	r, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// fmt.Println(string(rBody))
	var obj map[string]interface{}
	json.Unmarshal([]byte(rBody), &obj)

	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(obj)

	if grep != "" {
		scanner := bufio.NewScanner(bytes.NewReader(s))

		for scanner.Scan() {
			l := scanner.Text()
			if strings.Contains(l, grep) {
				fmt.Println(l)
			}
		}
	} else {
		fmt.Println(string(s))
	}
}
