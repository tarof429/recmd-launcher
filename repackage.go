package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {

	wd, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	parentDir := filepath.Dir(wd)

	stageDir := filepath.Join(parentDir, "stage")

	directoryMode := int(0755)

	os.Mkdir(stageDir, os.FileMode(directoryMode))
	os.Mkdir(filepath.Join(stageDir, "bin"), os.FileMode(directoryMode))
	os.Mkdir(filepath.Join(stageDir, "conf"), os.FileMode(directoryMode))
	os.Mkdir(filepath.Join(stageDir, "logs"), os.FileMode(directoryMode))
	symlink("bin/recmd-cli", filepath.Join(stageDir, "recmd"))
	copy(filepath.Join(parentDir, "bin/recmd-dmn"), filepath.Join(stageDir, "bin/recmd-dmn"))
	chmod(filepath.Join(stageDir, "bin/recmd-dmn"), 0755)
	chmod(filepath.Join(stageDir, "bin/recmd-cli"), 0755)
	copy(filepath.Join(parentDir, "bin/recmd-cli"), filepath.Join(stageDir, "bin/recmd-cli"))
	copy(filepath.Join(parentDir, "conf/recmd_history.json"), filepath.Join(stageDir, "conf/recmd_history.json"))

	// err = os.Mkdir(stageDir, os.FileMode(directoryMode))

	// if err != nil {
	// 	log.Println("Error, unable to create stage dir")
	// }

	// err = os.Mkdir(filepath.Join(stageDir, "bin"), os.FileMode(directoryMode))

	// if err != nil {
	// 	log.Println("Error, unable to create bin dir")
	// }
	// err = os.Mkdir(filepath.Join(stageDir, "conf"), os.FileMode(directoryMode))

	// if err != nil {
	// 	log.Println("Error, unable to create bin dir")
	// }

	// err = os.Mkdir(filepath.Join(stageDir, "logs"), os.FileMode(directoryMode))

	// if err != nil {
	// 	log.Println("Error, unable to create bin dir")
	// }

	// copy(filepath.Join(parentDir, "bin/recmd-dmn"), filepath.Join(stageDir, "bin/recmd-dmn"))

}

func symlink(from, to string) {
	os.Symlink(from, to)
}

func chmod(file string, perm os.FileMode) {
	os.Chmod("test.txt", perm)
}

func copy(from, to string) {
	// Open original file
	original, _ := os.Open(from)
	// if err != nil {
	// 	log.Println("Unable to open %v\n", from)
	// }

	defer original.Close()

	// Create new file
	toFile, _ := os.Create(to)

	// if err != nil {
	// 	log.Println("Unable to open %v\n", to)
	// }
	defer toFile.Close()

	//This will copy
	io.Copy(toFile, original)

	// if err != nil {
	// 	log.Println("Unable to copy file")
	// }
	//fmt.Printf("Bytes Written: %d\n", bytesWritten)
}
