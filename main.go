package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

const (
	ElemLevel   = "│"
	Elem        = "├───"
	ElemLast    = "└───"
	SizeEmpty   = "empty"
	ElemLastDir = ">>"
	ExcludeFile = ".DS_Store"
)

// regCleaner used to clean up element path
var regCleaner = regexp.MustCompile(`\W*`)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

// dirTree prints Directory tree by specified path, optionally prints file names and file size
func dirTree(out io.Writer, path string, printFiles bool) error {
	// cleaning up for ElemLastDir flags
	// dirPath used for getting directory structure
	dirPath := path
	if strings.Count(path, ElemLastDir) > 0 {
		dirPath = strings.ReplaceAll(dirPath, ElemLastDir, "")
	}
	// reading directory structure
	dirList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// sort filenames by ascending order
	sort.SliceStable(dirList, func(i, j int) bool { return dirList[i].Name() < dirList[j].Name() })
	// skip filenames if no printFiles flag is set
	for i := 0; i < len(dirList); i++ {
		// skip file info if printFiles is false
		if !dirList[i].IsDir() && (!printFiles || dirList[i].Name() == ExcludeFile) {
			dirList = append(dirList[:i], dirList[i+1:]...)
			i--
		}
	}

	// define last element
	lastPos := len(dirList) - 1
	for i, file := range dirList {
		isLast := i == lastPos
		err = printTree(out, file, path, isLast)
		if err != nil {
			return err
		}

		// recursive call to print folded directory
		if file.IsDir() {
			// set ElemLastDir flag
			foldingLast := ""
			if isLast {
				foldingLast = ElemLastDir
			}
			subPath := fmt.Sprintf("%s%s%s%s", path, string(os.PathSeparator), foldingLast, file.Name())

			err := dirTree(out, subPath, printFiles)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// printTree outputs tree line for path
func printTree(out io.Writer, file os.FileInfo, path string, isLast bool) error {
	// clean up all dots
	path = strings.ReplaceAll(path, ".", "")
	// extracting path separators
	schema := regCleaner.FindAllString(path, -1)
	// building tree
	// join all separators in one line
	tree := strings.Join(schema, "")
	// replacing ElemLastDir to empty separator
	tree = strings.ReplaceAll(tree, string(os.PathSeparator)+ElemLastDir, "\t ")
	// replacing os.PathSeparator to level separator
	tree = strings.ReplaceAll(tree, string(os.PathSeparator), ElemLevel+"\t")
	// define type of element connection
	filePrefix := Elem
	if isLast {
		// Tip of branch has special connection
		filePrefix = ElemLast
	}

	// form element name and info
	name := file.Name()
	if !file.IsDir() {
		size := file.Size()
		sizeText := SizeEmpty
		if size > 0 {
			sizeText = fmt.Sprintf("%vb", size)
		}
		name += fmt.Sprintf(" (%s)", sizeText)
	}

	// form line and send to the Writer
	log := fmt.Sprintf("%s%s%s\n", tree, filePrefix, name)
	_, err := out.Write([]byte(log))

	return err
}
