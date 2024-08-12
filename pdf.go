package contentreader

import (
	"os"
	"os/exec"
	"time"
)

func readPdf(content []byte) (string, error) {
	tmpFileName := "/tmp/tmp.pdf" + time.Now().String()
	os.WriteFile(tmpFileName, content, 0600)
	defer os.Remove(tmpFileName)

	body, err := exec.Command("pdftotext", "-nopgbrk", "-enc", "UTF-8", tmpFileName, "-").Output()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
