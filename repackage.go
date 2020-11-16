package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
	os.Mkdir(filepath.Join(stageDir, "data"), os.FileMode(directoryMode))

	symlink("bin/recmd-cli", filepath.Join(stageDir, "recmd"))
	copy(filepath.Join(parentDir, "bin/recmd-dmn"), filepath.Join(stageDir, "bin/recmd-dmn"))
	chmod(filepath.Join(stageDir, "bin/recmd-dmn"), 0755)
	chmod(filepath.Join(stageDir, "bin/recmd-cli"), 0755)
	copy(filepath.Join(parentDir, "bin/recmd-cli"), filepath.Join(stageDir, "bin/recmd-cli"))
	copy(filepath.Join(parentDir, "conf/recmd_history.json"), filepath.Join(stageDir, "conf/recmd_history.json"))

	// Doesn't work, use workaround
	//copyAllFiles("../data", stageDir+"/data")
	cmd := exec.Command("cp", "-r", "data", stageDir)
	cmd.Dir = parentDir
	cmd.Wait()
	combinedOutput, _ := cmd.CombinedOutput()
	fmt.Println(string(combinedOutput))

	cmd = exec.Command("zip", "-r", "../recmd.zip", ".")
	cmd.Dir = stageDir
	cmd.Wait()
	combinedOutput, _ = cmd.CombinedOutput()
	fmt.Println(string(combinedOutput))
}

func symlink(from, to string) {
	os.Symlink(from, to)
}

func chmod(file string, perm os.FileMode) {
	os.Chmod("test.txt", perm)
}

func copy(from, to string) {
	original, _ := os.Open(from)
	defer original.Close()

	toFile, _ := os.Create(to)
	defer toFile.Close()

	io.Copy(toFile, original)
}

func copyAllFiles(from, to string) {

	files, err := ioutil.ReadDir(from)

	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		copy(f.Name(), to+"/"+f.Name())
	}
}
