package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	diskMap := strings.TrimSpace(string(content))
	uncompressedDiskMap, diskBlocks := uncompressDisk(diskMap)
	defragmentedDiskMap := defragmentDisk(uncompressedDiskMap)
	defragmentedDiskBlocks := defragmentDiskBlocks(diskBlocks)
	checksum, blocksChecksum := getChecksums(defragmentedDiskMap, defragmentedDiskBlocks)

	fmt.Println("Resulting defragmented disk checksum:", checksum)
	fmt.Println("Resulting defragmented disk by files checksum:", blocksChecksum)
}

type DiskSpace struct {
	id   int
	free bool
}

type DiskBlock struct {
	id   int
	size int
	free bool
}

func uncompressDisk(diskMap string) ([]DiskSpace, []DiskBlock) {
	id := 0
	uncompressedDiskMap := []DiskSpace{}
	diskBlocks := []DiskBlock{}

	for i := 0; i < len(diskMap); i++ {
		val := int(diskMap[i] - '0')

		if i%2 == 0 {
			for j := 0; j < val; j++ {
				uncompressedDiskMap = append(uncompressedDiskMap, DiskSpace{id: id, free: false})
			}
			diskBlocks = append(diskBlocks, DiskBlock{id: id, size: val, free: false})
			id++
		} else {
			for j := 0; j < val; j++ {
				uncompressedDiskMap = append(uncompressedDiskMap, DiskSpace{id: -1, free: true})
			}
			diskBlocks = append(diskBlocks, DiskBlock{id: -1, size: val, free: true})
		}
	}
	return uncompressedDiskMap, diskBlocks
}

func defragmentDisk(uncompressedDiskMap []DiskSpace) []DiskSpace {
	for i := len(uncompressedDiskMap) - 1; i > 0; i-- {
		if !uncompressedDiskMap[i].free {
			fileID := uncompressedDiskMap[i].id
			for j := 0; j < i; j++ {
				if uncompressedDiskMap[j].free {
					uncompressedDiskMap[j] = DiskSpace{id: fileID, free: false}
					uncompressedDiskMap[i] = DiskSpace{id: -1, free: true}
					break
				}
			}
		}
	}

	return uncompressedDiskMap
}

func defragmentDiskBlocks(diskBlocks []DiskBlock) []DiskBlock {
	moreBlocks := true
	for moreBlocks {
		found := false
		for i := len(diskBlocks) - 1; i > 0; i-- {
			diskBlock := diskBlocks[i]
			if !diskBlock.free {
				for j := 0; j < i; j++ {
					if diskBlocks[j].free && diskBlocks[j].size >= diskBlock.size {
						emptySpace := diskBlocks[j].size - diskBlock.size

						diskBlocks[j] = DiskBlock{id: diskBlock.id, size: diskBlock.size, free: false}
						diskBlocks[i] = DiskBlock{id: -1, size: diskBlock.size, free: true}

						if emptySpace > 0 {
							emptySpaceBlock := DiskBlock{id: -1, size: emptySpace, free: true}
							afterEmptySpaceBlock := append([]DiskBlock{emptySpaceBlock}, diskBlocks[j+1:]...)
							diskBlocks = append(diskBlocks[:j+1], afterEmptySpaceBlock...)
						}
						found = true
						break
					}

				}
			}

			if found {
				break
			}
		}

		if !found {
			moreBlocks = false
		}
	}

	return diskBlocks
}

func getChecksums(disk []DiskSpace, blocks []DiskBlock) (int, int) {
	checksum := 0
	blocksChecksum := 0

	for i, disk := range disk {
		if disk.free {
			break
		}

		checksum += i * disk.id
	}

	currentIndex := 0
	for _, block := range blocks {
		if block.free {
			currentIndex += block.size
			continue
		}

		for j := 0; j < block.size; j++ {
			blocksChecksum += currentIndex * block.id
			currentIndex++
		}
	}

	return checksum, blocksChecksum
}
