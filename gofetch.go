package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	printTuxASCII()

	fmt.Println("-------------------------")
	printOSInfo()
	printKernelInfo()
	printCPUInfo()
	printRAMInfo()
	printUptime()
}

func printTuxASCII() {
	fmt.Println("   .--.    ")
	fmt.Println("  |o_o |   ")
	fmt.Println("  |:_/ |   ")
	fmt.Println(" //   \\ \\  ")
	fmt.Println("(|     | ) ")
	fmt.Println("/'\\_   _/`\\")
	fmt.Println("\\___)=(___/")
}

func printOSInfo() {

	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("OS: Unknown")
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {

			name := strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\"")
			fmt.Println("OS:", name)
			return
		}
	}

	fmt.Println("OS:", runtime.GOOS)
}

func printKernelInfo() {

	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		fmt.Println("Kernel: Unknown")
		return
	}

	kernel := strings.TrimSpace(string(out))
	fmt.Println("Kernel:", kernel)
}

func printCPUInfo() {

	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		fmt.Println("CPU: Unknown")
		return
	}

	lines := strings.Split(string(data), "\n")
	var modelName string
	coreCount := 0

	for _, line := range lines {
		if strings.HasPrefix(line, "model name") {

			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				modelName = strings.TrimSpace(parts[1])
			}
		}
		if strings.HasPrefix(line, "processor") {
			coreCount++
		}
	}

	if modelName != "" {
		fmt.Printf("CPU: %s (%d)\n", modelName, coreCount)
	} else {
		fmt.Println("CPU: Unknown")
	}
}

func printRAMInfo() {

	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("RAM: Unknown")
		return
	}

	lines := strings.Split(string(data), "\n")
	var total, available int64

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		if fields[0] == "MemTotal:" {
			total, _ = strconv.ParseInt(fields[1], 10, 64)
		}
		if fields[0] == "MemAvailable:" {
			available, _ = strconv.ParseInt(fields[1], 10, 64)
		}
	}

	totalMB := total / 1024
	usedMB := (total - available) / 1024

	fmt.Printf("RAM: %dMB / %dMB\n", usedMB, totalMB)
}

func printUptime() {

	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		fmt.Println("Uptime: Unknown")
		return
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		fmt.Println("Uptime: Unknown")
		return
	}

	uptimeFloat, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		fmt.Println("Uptime: Unknown")
		return
	}

	uptime := int(uptimeFloat)
	hours := uptime / 3600
	minutes := (uptime % 3600) / 60

	fmt.Printf("Uptime: %d hours, %d mins\n", hours, minutes)
}
