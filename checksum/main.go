package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/zeebo/blake3"
	"golang.org/x/crypto/blake2s"
)

// generateBlake2sChecksum generates smaller BLAKE2s checksum (16 bytes = 128 bits)
func generateBlake2sChecksum(data []byte, result chan<- string) {
	hash, err := blake2s.New256(nil) // Use 256-bit BLAKE2s and truncate
	if err != nil {
		result <- fmt.Sprintf("BLAKE2s Error: %v", err)
		return
	}
	hash.Write(data)
	fullHash := hash.Sum(nil)
	smallHash := fullHash[:16] // Truncate to 16 bytes
	checksum := hex.EncodeToString(smallHash)
	result <- fmt.Sprintf("BLAKE2s-128: %s", checksum)
}

// generateBlake3Checksum generates smaller BLAKE3 checksum (16 bytes = 128 bits)
func generateBlake3Checksum(data []byte, result chan<- string) {
	hash := blake3.New()
	hash.Write(data)
	// Get only first 16 bytes for smaller checksum
	fullHash := hash.Sum(nil)
	smallHash := fullHash[:16] // Truncate to 16 bytes
	checksum := hex.EncodeToString(smallHash)
	result <- fmt.Sprintf("BLAKE3-128: %s", checksum)
}

// generateTestData creates random test data
func generateTestData(size int) []byte {
	data := make([]byte, size)
	rand.Read(data)
	return data
}

func main() {
	
	// Checksum generation demo
	fmt.Println("=== Concurrent Small Checksum Generation ===")

	// Create test data sets
	testDataSets := [][]byte{
		[]byte("Hello, World!"),
		[]byte("This is a test string for BLAKE checksums"),
		generateTestData(1024), // 1KB random data
		generateTestData(4096), // 4KB random data
	}

	// Channel to collect checksum results
	checksumResults := make(chan string, len(testDataSets)*2)

	// Generate checksums concurrently
	for i, data := range testDataSets {
		fmt.Printf("Processing dataset %d (size: %d bytes)\n", i+1, len(data))

		// Launch goroutines for both BLAKE2s and BLAKE3
		go generateBlake2sChecksum(data, checksumResults)
		go generateBlake3Checksum(data, checksumResults)
	}

	// Collect results
	for i := 0; i < len(testDataSets)*2; i++ {
		fmt.Println(<-checksumResults)
	}
}
