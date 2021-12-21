package main

import "strings"

const layout = "02-01-2006 15:04"

const timeLayout = "02-01-2006 15:04 +07:00"
const formattedLayout = "2 January 2006, 15:04 MST"

var status = map[string]Status{
	"SHIPMENT RECEIVED BY JNE": {
		Code:    "060094",
		Message: "Package Received By JNE",
	},
	"RECEIVED AT SORTING CENTER": {
		Code:    "060095",
		Message: "Package Received at JNE sorting center",
	},
	"RECEIVED AT ORIGIN": {
		Code:    "060096",
		Message: "Package Received at JNE origin",
	},
	"RECEIVED AT WAREHOUSE": {
		Code:    "060097",
		Message: "Package Received at JNE warehouse",
	},
	"SHIPMENT FORWARDED TO DESTINATION": {
		Code:    "060098",
		Message: "Package forwarded to destination",
	},
	"RECEIVED AT INBOUND STATION": {
		Code:    "060099",
		Message: "Package Received at Inbound station",
	},
	"WITH DELIVERY COURIER": {
		Code:    "0600100",
		Message: "Package forwarded to destination with delivery courier",
	},
	"DELIVERED TO": {
		Code:    "060101",
		Message: "Delivery tracking detail fetched successfully",
	},
}

var r = strings.NewReplacer(
	"January", "Januari",
	"February", "Februari",
	"March", "Maret",
	"April", "April",
	"May", "Mei",
	"June", "Juni",
	"July", "Juli",
	"August", "Agustus",
	"September", "September",
	"October", "Oktober",
	"November", "November",
	"December", "Desember")
