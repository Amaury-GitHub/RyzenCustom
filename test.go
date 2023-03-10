package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/lxn/walk"
)

var PowerMode string
var PowerLimit string
var TempLimit string
var BoostStatus string
var EnergyStarStatus string
var CpuPower string
var CpuTemp string

var MainWindow *walk.MainWindow
var Icon *walk.Icon
var NotifyIcon *walk.NotifyIcon
var IcoData []byte = []byte{
	0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x20, 0x20, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0xa8, 0x10, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00,
	0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0xc3, 0x0e,
	0x00, 0x00, 0xc3, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x40, 0x51, 0xff, 0x00, 0x15, 0x1b,
	0x61, 0x00, 0x23, 0x2c, 0xa0, 0x00, 0x25, 0x2e, 0xaa, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e, 0xa9, 0x00, 0x25, 0x2e,
	0xa9, 0x00, 0x25, 0x2e, 0xaa, 0x00, 0x23, 0x2c, 0xa0, 0x00, 0x15, 0x1b,
	0x61, 0x00, 0x40, 0x51, 0xff, 0x00, 0x00, 0x00, 0x01, 0x00, 0x52, 0x68,
	0xff, 0x00, 0x2a, 0x35, 0xc3, 0x00, 0x0f, 0x13, 0x47, 0x0e, 0x19, 0x20,
	0x73, 0x2e, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21, 0x78, 0x32, 0x1a, 0x21,
	0x78, 0x32, 0x19, 0x20, 0x73, 0x2e, 0x0f, 0x13, 0x47, 0x0e, 0x2a, 0x35,
	0xc3, 0x00, 0x53, 0x68, 0xff, 0x00, 0x1f, 0x27, 0x8f, 0x00, 0x1d, 0x24,
	0x84, 0x08, 0x26, 0x30, 0xb0, 0x90, 0x29, 0x33, 0xbb, 0xe6, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33, 0xbc, 0xe7, 0x29, 0x33,
	0xbb, 0xe6, 0x26, 0x30, 0xb0, 0x90, 0x1d, 0x24, 0x84, 0x08, 0x1f, 0x27,
	0x8f, 0x00, 0x29, 0x34, 0xbd, 0x00, 0x29, 0x34, 0xbd, 0x1c, 0x2a, 0x35,
	0xc1, 0xdf, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc1, 0xdf, 0x29, 0x34, 0xbd, 0x1c, 0x29, 0x34, 0xbd, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2d, 0x37, 0xc1, 0xff, 0x3e, 0x47, 0xbc, 0xff, 0x41, 0x4a,
	0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x41, 0x49,
	0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x41, 0x4a, 0xbb, 0xff, 0x34, 0x3e,
	0xbb, 0xff, 0x28, 0x33, 0xc1, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xbe, 0xff, 0x2b, 0x36, 0xc1, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x41, 0x4a,
	0xc8, 0xff, 0xcd, 0xd0, 0xec, 0xff, 0xe6, 0xe6, 0xf3, 0xff, 0xe4, 0xe5,
	0xf2, 0xff, 0xe4, 0xe5, 0xf2, 0xff, 0xe4, 0xe5, 0xf2, 0xff, 0xe4, 0xe5,
	0xf2, 0xff, 0xe6, 0xe7, 0xf3, 0xff, 0xbe, 0xc0, 0xe3, 0xff, 0x44, 0x4d,
	0xbf, 0xff, 0x27, 0x33, 0xc1, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2d, 0x38, 0xbd, 0xff, 0x72, 0x78,
	0xcb, 0xff, 0x3e, 0x48, 0xc7, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x44, 0x4e, 0xca, 0xff, 0xe7, 0xe8,
	0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc4, 0xc6, 0xe6, 0xff, 0x43, 0x4b,
	0xbf, 0xff, 0x28, 0x33, 0xc1, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2c, 0x37,
	0xbd, 0xff, 0x8c, 0x91, 0xd2, 0xff, 0xdd, 0xdf, 0xf4, 0xff, 0x45, 0x4f,
	0xca, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc3, 0xc5, 0xe6, 0xff, 0x41, 0x4a,
	0xbf, 0xff, 0x28, 0x33, 0xc1, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x2c, 0x37, 0xbd, 0xff, 0x8b, 0x90, 0xd2, 0xff, 0xf9, 0xf9,
	0xfc, 0xff, 0xe6, 0xe8, 0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xc3, 0xc5, 0xe6, 0xff, 0x42, 0x4b,
	0xbf, 0xff, 0x28, 0x33, 0xc1, 0xff, 0x2d, 0x37, 0xbd, 0xff, 0x8c, 0x91,
	0xd2, 0xff, 0xf8, 0xf8, 0xfb, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xfc, 0xfc, 0xfe, 0xff, 0xe8, 0xe9, 0xf8, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xe5, 0xe7,
	0xf8, 0xff, 0xe7, 0xe8, 0xf8, 0xff, 0xa5, 0xa9, 0xdf, 0xff, 0x39, 0x43,
	0xc2, 0xff, 0x8b, 0x90, 0xd4, 0xff, 0xf8, 0xf8, 0xfb, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0xe5, 0xe7, 0xf8, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe8, 0xe9,
	0xf8, 0xff, 0x5b, 0x63, 0xd0, 0xff, 0x42, 0x4c, 0xc9, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x43, 0x4c, 0xc9, 0xff, 0x48, 0x52, 0xca, 0xff, 0xe3, 0xe4,
	0xf6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x34, 0xc2, 0xff, 0x42, 0x4c,
	0xc9, 0xff, 0xe3, 0xe5, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x42, 0x4c,
	0xc9, 0xff, 0x27, 0x32, 0xc1, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x27, 0x32,
	0xc1, 0xff, 0x44, 0x4d, 0xc9, 0xff, 0xe5, 0xe7, 0xf8, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2e, 0x38, 0xc3, 0xff, 0x8d, 0x93,
	0xde, 0xff, 0xf8, 0xf8, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2d, 0x38, 0xc3, 0xff, 0x8a, 0x90,
	0xde, 0xff, 0xf6, 0xf7, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2c, 0x37, 0xc3, 0xff, 0x87, 0x8d,
	0xdd, 0xff, 0xf7, 0xf7, 0xfd, 0xff, 0xe7, 0xe8, 0xf8, 0xff, 0x44, 0x4d,
	0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2c, 0x37, 0xc2, 0xff, 0x85, 0x8b,
	0xdc, 0xff, 0xda, 0xdc, 0xf4, 0xff, 0x45, 0x4f, 0xca, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2d, 0x37, 0xc3, 0xff, 0x69, 0x71,
	0xd4, 0xff, 0x3b, 0x45, 0xc7, 0xff, 0x27, 0x32, 0xc1, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x27, 0x32, 0xc1, 0xff, 0x42, 0x4c, 0xc9, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2c, 0x36, 0xbd, 0xff, 0x40, 0x48,
	0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x41, 0x49,
	0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x41, 0x49, 0xbb, 0xff, 0x3f, 0x48,
	0xba, 0xff, 0x58, 0x5f, 0xc3, 0xff, 0xe8, 0xe9, 0xf7, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2e, 0x38,
	0xbd, 0xff, 0x90, 0x94, 0xd4, 0xff, 0xe3, 0xe4, 0xf2, 0xff, 0xe4, 0xe5,
	0xf2, 0xff, 0xe4, 0xe5, 0xf2, 0xff, 0xe4, 0xe5, 0xf2, 0xff, 0xe4, 0xe5,
	0xf2, 0xff, 0xe4, 0xe5, 0xf2, 0xff, 0xe4, 0xe5, 0xf2, 0xff, 0xe7, 0xe8,
	0xf4, 0xff, 0xfc, 0xfc, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x2d, 0x38, 0xbd, 0xff, 0x91, 0x95, 0xd4, 0xff, 0xf9, 0xfa,
	0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xe5, 0xe6, 0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2d, 0x37, 0xbd, 0xff, 0x90, 0x94,
	0xd4, 0xff, 0xf9, 0xf9, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe5, 0xe6,
	0xf8, 0xff, 0x44, 0x4e, 0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x35, 0xc2, 0xff, 0x2f, 0x39,
	0xbe, 0xff, 0x91, 0x95, 0xd4, 0xff, 0xfb, 0xfb, 0xfc, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xe7, 0xe8, 0xf8, 0xff, 0x44, 0x4e,
	0xc9, 0xff, 0x28, 0x33, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x71, 0x77, 0xd0, 0xff, 0xd7, 0xd9,
	0xf2, 0xff, 0xdf, 0xe0, 0xf6, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf,
	0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf,
	0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf,
	0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf,
	0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xde, 0xdf, 0xf5, 0xff, 0xdf, 0xe1,
	0xf6, 0xff, 0xc8, 0xcb, 0xef, 0xff, 0x40, 0x4a, 0xc8, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2b, 0x36,
	0xc2, 0xff, 0x3a, 0x44, 0xc7, 0xff, 0x3e, 0x48, 0xc8, 0xff, 0x3d, 0x47,
	0xc7, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47,
	0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47,
	0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47,
	0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47,
	0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3d, 0x47, 0xc8, 0xff, 0x3b, 0x45,
	0xc7, 0xff, 0x2c, 0x37, 0xc3, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35,
	0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34,
	0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x29, 0x34, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x00, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xe0, 0x2a, 0x35, 0xc2, 0x1f, 0x2a, 0x35, 0xc2, 0x00, 0x31, 0x3c,
	0xc4, 0x00, 0x32, 0x3c, 0xc4, 0x1c, 0x2c, 0x37, 0xc3, 0xde, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2a, 0x35,
	0xc2, 0xff, 0x2a, 0x35, 0xc2, 0xff, 0x2c, 0x37, 0xc3, 0xde, 0x32, 0x3c,
	0xc4, 0x1c, 0x31, 0x3c, 0xc4, 0x00, 0x3d, 0x47, 0xc7, 0x00, 0x41, 0x4b,
	0xc9, 0x05, 0x37, 0x41, 0xc6, 0x84, 0x33, 0x3e, 0xc5, 0xde, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e, 0xc5, 0xe0, 0x33, 0x3e,
	0xc5, 0xde, 0x37, 0x41, 0xc6, 0x84, 0x41, 0x4b, 0xc9, 0x05, 0x3d, 0x47,
	0xc7, 0x00, 0x3c, 0x46, 0xc7, 0x00, 0x3e, 0x48, 0xc8, 0x00, 0x48, 0x52,
	0xcb, 0x05, 0x42, 0x4c, 0xc9, 0x1c, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x41, 0x4b,
	0xc9, 0x1f, 0x41, 0x4b, 0xc9, 0x1f, 0x42, 0x4c, 0xc9, 0x1c, 0x48, 0x52,
	0xcb, 0x05, 0x3e, 0x48, 0xc8, 0x00, 0x3c, 0x46, 0xc7, 0x00, 0x3e, 0x48,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49, 0xc8, 0x00, 0x3f, 0x49,
	0xc8, 0x00, 0x3e, 0x47, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00,
}

func BoostOff() {
	//关闭睿频
	Command := exec.Command("cmd", "/c", "powercfg /SETACVALUEINDEX scheme_current sub_processor perfboostmode 0")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	Command = exec.Command("cmd", "/c", "powercfg /SETDCVALUEINDEX scheme_current sub_processor perfboostmode 0")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	Command = exec.Command("cmd", "/c", "powercfg /SETACTIVE SCHEME_CURRENT")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	BoostStatus = "关闭"
}
func BoostOn() {
	//打开睿频
	Command := exec.Command("cmd", "/c", "powercfg /SETACVALUEINDEX scheme_current sub_processor perfboostmode 1")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	Command = exec.Command("cmd", "/c", "powercfg /SETDCVALUEINDEX scheme_current sub_processor perfboostmode 1")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	Command = exec.Command("cmd", "/c", "powercfg /SETACTIVE SCHEME_CURRENT")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	BoostStatus = "打开"
}
func EnergyStarOff() {
	//关闭EnergyStar
	Command := exec.Command("cmd", "/c", "taskkill /f /im EnergyStar.exe")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
}
func EnergyStarOn() {
	//关闭EnergyStar
	EnergyStarOff()
	//启动EnergyStar
	Command := exec.Command("cmd", "/c", "EnergyStar\\EnergyStar.exe")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Start()
}
func PowerLow() {
	PowerMode = "能效模式"
	//启动EnergyStar
	EnergyStarOn()
	//关闭睿频
	BoostOff()
	PowerLimit = "5000"
	TempLimit = "55"
}
func PowerMid() {
	PowerMode = "平衡模式"
	//启动EnergyStar
	EnergyStarOn()
	//关闭睿频
	BoostOff()
	PowerLimit = "15000"
	TempLimit = "70"
}
func PowerHigh() {
	PowerMode = "性能模式"
	//关闭EnergyStar
	EnergyStarOff()
	//打开睿频
	BoostOn()
	PowerLimit = "45000"
	TempLimit = "85"
}
func CallRyzenAdj() {
	//生成command
	RyzenAdjCommand := "RyzenAdj\\ryzenadj --stapm-limit=" + PowerLimit + " --fast-limit=" + PowerLimit + " --slow-limit=" + PowerLimit + " --tctl-temp=" + TempLimit
	//调用RyzenAdj
	Command := exec.Command("cmd", "/c", RyzenAdjCommand)
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
}
func CallRyzenAdjLoop() {
	//调用RyzenAdj
	go CallRyzenAdj()
	//10s循环
	time.AfterFunc(10*time.Second, CallRyzenAdjLoop)
}
func ShowMessage() {
	//延时
	time.Sleep(time.Duration(1) * time.Second)
	//读取EnergyStar运行状态
	Command := exec.Command("powershell", "Get-Process EnergyStar")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := Command.Output()
	if err == nil {
		EnergyStarStatus = "运行中"
	} else {
		EnergyStarStatus = "未运行"
	}
	//读取CPU功率和温度
	Command = exec.Command("cmd", "/c", "RyzenAdj\\ryzenadj -i")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Output, err := Command.Output()
	if err == nil {
		CpuPower = string(Output)[strings.Index(string(Output), "STAPM VALUE")+21 : strings.Index(string(Output), "STAPM VALUE")+32]
		CpuTemp = string(Output)[strings.Index(string(Output), "THM VALUE CORE")+21 : strings.Index(string(Output), "THM VALUE CORE")+32]
	}
	//显示通知
	NotifyIcon.ShowMessage(PowerMode+"\r\n"+"设置功率 : "+PowerLimit[:strings.LastIndex(PowerLimit, "000")]+" W"+"    设置温度 : "+TempLimit+" ℃", "EnergyStar : "+EnergyStarStatus+"\r\n"+"Boost : "+BoostStatus+"\r\n"+"实时功率 : "+CpuPower+" W"+"\r\n"+"实时温度 : "+CpuTemp+" ℃")
}

func init() {
	//阻止多次启动
	Mutex, _ := syscall.UTF16PtrFromString("RyzenCustom")
	_, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("CreateMutexW").Call(0, 0, uintptr(unsafe.Pointer(Mutex)))
	if int(err.(syscall.Errno)) != 0 {
		os.Exit(1)
	}
	//创建ICO
	os.WriteFile("icon.ico", IcoData, 0644)
}
func main() {
	//定义托盘图标文字
	MainWindow, _ = walk.NewMainWindow()
	Icon, _ = walk.Resources.Icon("icon.ico")
	//删除图标
	Command := exec.Command("cmd", "/c", "del /f /q icon.ico")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	NotifyIcon, _ = walk.NewNotifyIcon(MainWindow)
	defer NotifyIcon.Dispose()
	NotifyIcon.SetIcon(Icon)
	NotifyIcon.SetToolTip("RyzenCustom")
	NotifyIcon.SetVisible(true)
	//定义左键显示
	NotifyIcon.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		//显示通知
		ShowMessage()
	})
	//定义右键菜单
	blank1 := walk.NewAction()
	blank1.SetText("-")
	blank2 := walk.NewAction()
	blank2.SetText("-")
	blank3 := walk.NewAction()
	blank3.SetText("-")
	blank4 := walk.NewAction()
	blank4.SetText("-")
	blank5 := walk.NewAction()
	blank5.SetText("-")
	blank6 := walk.NewAction()
	blank6.SetText("-")
	blank7 := walk.NewAction()
	blank7.SetText("-")
	SetPowerLow := walk.NewAction()
	SetPowerLow.SetText("能效模式")
	SetPowerMid := walk.NewAction()
	SetPowerMid.SetText("平衡模式")
	SetPowerHigh := walk.NewAction()
	SetPowerHigh.SetText("性能模式")
	EnergyStarOnButton := walk.NewAction()
	EnergyStarOnButton.SetText("打开 EnergyStar")
	EnergyStarOffButton := walk.NewAction()
	EnergyStarOffButton.SetText("关闭 EnergyStar")
	BoostOnButton := walk.NewAction()
	BoostOnButton.SetText("打开 Boost")
	BoostOffButton := walk.NewAction()
	BoostOffButton.SetText("关闭 Boost")
	Exit := walk.NewAction()
	Exit.SetText("Exit")
	NotifyIcon.ContextMenu().Actions().Add(SetPowerLow)
	NotifyIcon.ContextMenu().Actions().Add(blank1)
	NotifyIcon.ContextMenu().Actions().Add(SetPowerMid)
	NotifyIcon.ContextMenu().Actions().Add(blank2)
	NotifyIcon.ContextMenu().Actions().Add(SetPowerHigh)
	NotifyIcon.ContextMenu().Actions().Add(blank3)
	NotifyIcon.ContextMenu().Actions().Add(EnergyStarOnButton)
	NotifyIcon.ContextMenu().Actions().Add(blank4)
	NotifyIcon.ContextMenu().Actions().Add(EnergyStarOffButton)
	NotifyIcon.ContextMenu().Actions().Add(blank5)
	NotifyIcon.ContextMenu().Actions().Add(BoostOnButton)
	NotifyIcon.ContextMenu().Actions().Add(blank6)
	NotifyIcon.ContextMenu().Actions().Add(BoostOffButton)
	NotifyIcon.ContextMenu().Actions().Add(blank7)
	NotifyIcon.ContextMenu().Actions().Add(Exit)
	//设置为能效
	SetPowerLow.Triggered().Attach(func() {
		//设置参数为能效
		PowerLow()
		//调用RyzenAdj
		CallRyzenAdj()
		//显示通知
		ShowMessage()
	})
	//设置为平衡
	SetPowerMid.Triggered().Attach(func() {
		//设置参数为平衡
		PowerMid()
		//调用RyzenAdj
		CallRyzenAdj()
		//显示通知
		ShowMessage()
	})
	//设置为性能
	SetPowerHigh.Triggered().Attach(func() {
		//设置参数为性能
		PowerHigh()
		//调用RyzenAdj
		CallRyzenAdj()
		//显示通知
		ShowMessage()
	})
	//打开EnergyStar
	EnergyStarOnButton.Triggered().Attach(func() {
		//启动EnergyStar
		EnergyStarOn()
		//显示通知
		ShowMessage()
	})
	//关闭EnergyStar
	EnergyStarOffButton.Triggered().Attach(func() {
		//关闭EnergyStar
		EnergyStarOff()
		//显示通知
		ShowMessage()
	})
	//打开睿频
	BoostOnButton.Triggered().Attach(func() {
		//打开睿频
		BoostOn()
		//显示通知
		ShowMessage()
	})
	//关闭睿频
	BoostOffButton.Triggered().Attach(func() {
		//关闭睿频
		BoostOff()
		//显示通知
		ShowMessage()
	})
	//Exit
	Exit.Triggered().Attach(func() {
		//退出主程序
		walk.App().Exit(0)
	})

	//设置参数为平衡
	PowerMid()
	//调用RyzenAdj
	CallRyzenAdj()
	//显示通知
	ShowMessage()

	//循环调用RyzenAdj
	go CallRyzenAdjLoop()
	//主程序运行
	MainWindow.Run()
}
