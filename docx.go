package contentreader

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func readDocx(content []byte) (string, error) {
	tmpFileName := "/tmp/tmp.docx" + time.Now().String()
	os.WriteFile(tmpFileName, content, 0600)
	defer os.Remove(tmpFileName)

	cmd := fmt.Sprintf(`unzip -p '%s' -d unzip word/document.xml | sed -e 's/<\/w:p>/\n/g; s/<[^>]\{1,\}>//g; s/[^[:print:]\n]\{1,\}//g'`, tmpFileName)
	body, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(body), nil
}
