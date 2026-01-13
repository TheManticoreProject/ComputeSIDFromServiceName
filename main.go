package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"unicode/utf16"

	"github.com/TheManticoreProject/goopts/parser"
)

var (
	// Configuration
	debug bool

	// Source values
	serviceName string
)

func ComputeSIDFromServiceName(serviceName string) (string, error) {
	serviceName = strings.TrimSpace(serviceName)
	if serviceName == "" {
		return "", fmt.Errorf("service name is required")
	}

	upper := strings.ToUpper(serviceName)

	// Encode as UTF-16LE (Windows "Unicode")
	utf16Units := utf16.Encode([]rune(upper))
	utf16LE := make([]byte, len(utf16Units)*2)
	for i, u := range utf16Units {
		binary.LittleEndian.PutUint16(utf16LE[i*2:], u)
	}

	sum := sha1.Sum(utf16LE)

	// SHA-1 is 20 bytes -> five little-endian uint32 components
	parts := make([]uint32, 5)
	for i := 0; i < 20; i += 4 {
		parts[i/4] = binary.LittleEndian.Uint32(sum[i : i+4])
	}

	return fmt.Sprintf("S-1-5-80-%d-%d-%d-%d-%d", parts[0], parts[1], parts[2], parts[3], parts[4]), nil
}

func parseArgs() {
	ap := parser.ArgumentsParser{Banner: "ComputeSIDFromServiceName - by Remi GASCOU (Podalirius) @ TheManticoreProject - v1.0.0"}

	// Configuration flags
	ap.NewBoolArgument(&debug, "-d", "--debug", false, "Debug mode.")
	ap.NewStringArgument(&serviceName, "-s", "--service-name", "", true, "Service name to compute SID from.")

	ap.Parse()
}

func main() {
	parseArgs()

	sid, err := ComputeSIDFromServiceName(serviceName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	fmt.Println(sid)
}
