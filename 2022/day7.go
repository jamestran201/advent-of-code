package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Directory struct {
	name        string
	directories map[string]*Directory
	files       map[string]uint64
	parentDir   *Directory
	size        uint64
}

func (d *Directory) GetDir(dirName string) *Directory {
	return d.directories[dirName]
}

func (d *Directory) GetParentDir() *Directory {
	return d.parentDir
}

func (d *Directory) AddDir(dirName string) {
	if d.directories[dirName] != nil {
		return
	}

	d.directories[dirName] = NewDirectory(dirName, d)
}

func (d *Directory) AddFile(name string, size string) {
	_, isPresent := d.files[name]
	if isPresent {
		return
	}

	size_uint, err := strconv.ParseUint(size, 10, 64)
	if err != nil {
		panic(err)
	}

	d.files[name] = size_uint
}

func (d *Directory) GetFileSize(name string) uint64 {
	return d.files[name]
}

func (d *Directory) Size() uint64 {
	if d.size != 0 {
		return d.size
	}

	d.size = d.calculateSize()
	return d.size
}

func (d *Directory) calculateSize() uint64 {
	total := uint64(0)

	for _, size := range d.files {
		total += size
	}

	for _, childDir := range d.directories {
		total += childDir.size
	}

	return total
}

func NewDirectory(name string, parentDir *Directory) *Directory {
	return &Directory{
		name:        name,
		directories: map[string]*Directory{},
		files:       map[string]uint64{},
		parentDir:   parentDir,
	}
}

type CommandProcessor struct {
	cmdRegex       *regexp.Regexp
	currentDir     *Directory
	rootDir        *Directory
	isListingFiles bool
}

func (c *CommandProcessor) Process(input string) {
	match := c.cmdRegex.FindAllStringSubmatch(input, -1)
	if len(match) != 0 {
		cmd := match[0][1]
		arg := strings.TrimLeft(match[0][2], " ")
		c.handleCommand(cmd, arg)
	} else {
		c.parseListOutput(input)
	}
}

func (c *CommandProcessor) handleCommand(cmd string, arg string) {
	switch cmd {
	case "cd":
		c.isListingFiles = false

		switch arg {
		case "/":
			c.currentDir = c.rootDir
		case "..":
			c.currentDir = c.currentDir.GetParentDir()
		default:
			c.currentDir = c.currentDir.GetDir(arg)
		}

		fmt.Println("Go to dir", c.currentDir.name)
	case "ls":
		c.isListingFiles = true
		fmt.Println("Listing files for dir", c.currentDir.name)
	}
}

func (c *CommandProcessor) parseListOutput(input string) {
	parts := strings.Split(input, " ")
	switch parts[0] {
	case "dir":
		c.currentDir.AddDir(parts[1])
		fmt.Println("Added dir", c.currentDir.GetDir(parts[1]))
	default:
		c.currentDir.AddFile(parts[1], parts[0])
		fmt.Printf("Added file %s of size %d\n", parts[1], c.currentDir.GetFileSize(parts[1]))
	}
}

func NewCommandProcessor() *CommandProcessor {
	re := regexp.MustCompile("^\\$ (cd|ls)(.*)")
	rootDir := NewDirectory("/", nil)

	return &CommandProcessor{
		cmdRegex:       re,
		currentDir:     rootDir,
		rootDir:        rootDir,
		isListingFiles: false,
	}
}

func TotalSizeOfDirectoriesWithSizeAtMost(dir *Directory, maxSize uint64, total uint64) uint64 {
	if len(dir.directories) != 0 {
		for _, childDir := range dir.directories {
			total = TotalSizeOfDirectoriesWithSizeAtMost(childDir, maxSize, total)
		}
	}

	size := dir.Size()
	if size <= maxSize {
		total += size
	}

	return total
}

func FindSmallestDirToDelete(rootDir *Directory, totalFilesystemSize uint64, sizeNeeded uint64) uint64 {
	currentUnusedSize := totalFilesystemSize - rootDir.Size()
	return findSmallestDirToDeleteHelper(rootDir, currentUnusedSize, sizeNeeded, rootDir.size)
}

func findSmallestDirToDeleteHelper(dir *Directory, currentUnusedSize uint64, sizeNeeded uint64, minSizeToFree uint64) uint64 {
	newUnusedSize := currentUnusedSize + dir.Size()
	if newUnusedSize >= sizeNeeded && newUnusedSize < (minSizeToFree+currentUnusedSize) {
		minSizeToFree = dir.Size()
	}

	for _, childDir := range dir.directories {
		minSizeToFree = findSmallestDirToDeleteHelper(childDir, currentUnusedSize, sizeNeeded, minSizeToFree)
	}

	return minSizeToFree
}

func main() {
	// file, err := os.Open("day7_sample_input.txt")
	file, err := os.Open("day7_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	commandProcessor := NewCommandProcessor()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commandProcessor.Process(scanner.Text())
	}

	resultPart1 := TotalSizeOfDirectoriesWithSizeAtMost(commandProcessor.rootDir, 100000, 0)
	fmt.Println("Result for part 1:", resultPart1)

	resultPart2 := FindSmallestDirToDelete(commandProcessor.rootDir, 70000000, 30000000)
	fmt.Println("Result for part 2:", resultPart2)
}
