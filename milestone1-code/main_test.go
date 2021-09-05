package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestValidateOpts(t *testing.T) {
  for _, test := range testCasesValidateOpts {
    fmt.Printf("Running test #%d: %s\n", test.TestNumber, test.Description)
    err := validateOpts(test.Input)
    test.OutputMsg = err.Error()
    if test.OutputMsg != test.Want {
      t.Errorf("Test #%d failed...\ngot:\n%s\nwant:\n%s", test.TestNumber, test.OutputMsg, test.Want)
    }
  }
}

func TestSetupFlags(t *testing.T) {
  buffer := bytes.Buffer{}

  for _, test := range testCasesSetupFlags {
    fmt.Printf("Running test #%d: %s\n", test.TestNumber, test.Description)
    _, err := setupFlags(&buffer, test.Input)
    if err != nil {
      if err.Error() == "project not set up correctly. Try -h for flag usage" {
        continue
      }
    }
    test.OutputMsg = buffer.String()
    if test.OutputMsg != test.Want {
      t.Errorf("Test #%d failed...\ngot:\n%s\nwant:\n%s", test.TestNumber, test.OutputMsg, test.Want)
    }
    buffer.Reset()
  }
}

