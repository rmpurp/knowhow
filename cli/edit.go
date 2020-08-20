package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const defaultEditor = "vim"

func getTemporaryFile(extension string) (*os.File, error) {
	file, err := ioutil.TempFile("", fmt.Sprintf("*.%s", extension))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func openFileInEditor(filepath string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}

	executable, err := exec.LookPath(editor)

	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func EditText(text string) (string, error) {
	file, err := getTemporaryFile("md")

	if err != nil {
		return "", err
	}

	filename := file.Name()
	defer os.Remove(filename)

	file.WriteString(text)

	if err = file.Close(); err != nil {
		return "", err
	}

	if err = openFileInEditor(filename); err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
