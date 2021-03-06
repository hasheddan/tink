package framework

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/tinkerbell/tink/client"
	"github.com/tinkerbell/tink/protos/hardware"
)

func readHwData(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return []byte(""), err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte(""), err
	}
	return data, nil
}

// PushHardwareData : push hardware data
func PushHardwareData(hwDataFiles []string) error {
	for _, hwFile := range hwDataFiles {
		filepath := "data/hardware/" + hwFile
		data, err := readHwData(filepath)
		if err != nil {
			return err
		}
		hw := hardware.Hardware{}
		err = json.Unmarshal(data, &hw)
		_, err = client.HardwareClient.Push(context.Background(), &hardware.PushRequest{Data: &hw})
		if err != nil {
			return err
		}
	}
	return nil
}
