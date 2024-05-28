package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"strconv"
	"strings"

	// "strings"
	"os"
	rx "regexp"
)

type IPAddress struct {
	valid    bool
	sub_type string
	value    string
}

type Results struct {
	input  string
	usable int
	mask   string
	cidr   int
}

func main() {
	menu()
}

func menu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Subnet Calculator")
	fmt.Println("Please enter an IP/Mask or IP/CIDR")

	for {
		fmt.Print("\n-> ")
		text, _ := reader.ReadString('\n')
		input := IPAddress{valid: false, sub_type: "", value: text}
		test_result := test_input(input)
		if test_result.valid {
			result := process_input(test_result)
			fmt.Println(result)
			//fmt.Printf("")
		} else {
			fmt.Println("Input invalid!")
		}
	}
}

func test_input(input IPAddress) IPAddress {
	match_in := []byte(input.value)

	match_sm, _ := rx.Match(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|2[0-5]{2})\.?\b){4}\/(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|2[0-5]{2})\.?\b){4}`, match_in)
	if match_sm {
		input.valid = true
		input.sub_type = "sm"
		return input
	}

	match_cidr, _ := rx.Match(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|2[0-5]{2})\.?\b){4}\/(1[0-9]|2[0-9]|3[0-2]|[0-9])\D`, match_in)
	if match_cidr {
		input.valid = true
		input.sub_type = "cidr"
		return input
	}

	return input
}

func process_input(input IPAddress) Results {
	var result Results
	var host_bits int

	if input.sub_type == "sm" {
		input_array := strings.Split(input.value, "/")
		subnet_mask := strings.Split(input_array[1], ".")
		for _, s_val := range subnet_mask {
			i64_val, _ := strconv.ParseUint(s_val, 10, 8)
			i8_val := uint8(i64_val)
			if i8_val < 255 {
				host_bits += bits.TrailingZeros8(i8_val)
			}
		}
		cidr := 32 - host_bits
		usable := 2 ^ 32 - cidr
		result = Results{input: input.value, usable: usable, mask: input_array[1], cidr: cidr}
		return result
	}

	if input.sub_type == "cidr" {
		//input_array := strings.Split(input.value, "/")
		//cidr := input_array[1]

	}

	return result
}
