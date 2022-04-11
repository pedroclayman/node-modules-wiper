package cmd

import (
	"fmt"
	"sync"

	"github.com/pedroclayman/node-modules-wiper/modulesearch"
	"github.com/spf13/cobra"
)

const Giga = (1073741824)
const Mega = (1048576)
const Kilo = (1024)

func printSize(sizeInBytes int64) string {
	size := Size{0, 0, 0, 0}

	size.giga = sizeInBytes / Giga
	sizeInBytes = sizeInBytes - Giga*size.giga

	size.mega = sizeInBytes / Mega
	sizeInBytes = sizeInBytes - Mega*size.mega

	size.kilo = sizeInBytes / Kilo
	size.bytes = sizeInBytes - Kilo*size.kilo

	return fmt.Sprintf("%v GB %v MB %v KB %v bytes", size.giga, size.mega, size.kilo, size.bytes)
}

type Size struct {
	giga  int64
	mega  int64
	kilo  int64
	bytes int64
}

func printDirAndSize(dir string, size int64) {
	fmt.Printf("%v:\t%v\n", dir, printSize(size))
}

func list(path string, largerThan int64, process func(dir string, size int64)) {
	var dirs []string
	modulesearch.GetNodeModuleDirectories(path, &dirs)

	wg := new(sync.WaitGroup)

	for _, dir := range dirs {
		wg.Add(1)

		go func(dir string) {
			defer wg.Done()
			size, _ := modulesearch.DirSize(dir)

			if largerThan == -1 || size > largerThan {
				process(dir, size)
			}

		}(dir)
	}

	wg.Wait()
}

func getListCommand() *cobra.Command {
	var largerThan *int64 = new(int64)

	command := cobra.Command{
		Use: "list path",
		// Aliases: []string{"get", "get-all"},
		Short: "Lists all paths to node_module directiries under a path",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			list(path, *largerThan, printDirAndSize)
		},
	}
	command.Flags().Int64Var(largerThan, "larger-than", -1, "only lists node_modules dirs larger than the value provided (in bytes)")
	return &command
}
