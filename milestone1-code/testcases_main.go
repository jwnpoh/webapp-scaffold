package main

import (
	"fmt"
)

// global test vars
var (
	testName           = "test-project"
	testLocation       = "~/Downloads"
	testRemote         = "https://github.com/jwnpoh"
	errMsgTestName     = "Project name cannot be unspecified.\n"
	errMsgTestLocation = "Project path cannot be unspecified.\n"
	errMsgTestRemote   = "Project repository URL cannot be unspecified.\n"
)

// test object for func setupFlags()
type testCaseSetupFlags struct {
	TestNumber  int
	Description string
	Input       []string
	OutputMsg   string
	Want        string
}

// test object for func validateOpts()
type testCaseValidateOpts struct {
	TestNumber  int
	Description string
	Input       opts
	OutputMsg   string
	Want        string
}

// test func setupFlags()
var testCasesSetupFlags = []testCaseSetupFlags{
	{
		TestNumber:  1,
		Description: "All arguments entered.",
		Input:       []string{"-d", testLocation, "-n", testName, "-r", testRemote, "-s"},
		Want: fmt.Sprintf(
			"Generating project...\n Project Name: %s\n Local: %s\n Remote: %s\n Setup website: %v\n",
			testName,
			testLocation,
			testRemote,
			true,
		),
	},
	{
		TestNumber:  2,
		Description: "All arguments entered in a different order.",
		Input:       []string{"-n", testName, "-d", testLocation, "-r", testRemote, "-s"},
		Want: fmt.Sprintf(
			"Generating project...\n Project Name: %s\n Local: %s\n Remote: %s\n Setup website: %v\n",
			testName,
			testLocation,
			testRemote,
			true,
		),
	},
	{
		TestNumber:  3,
		Description: "All arguments entered in another different order.",
		Input:       []string{"-n", testName, "-r", testRemote, "-s", "-d", testLocation},
		Want: fmt.Sprintf(
			"Generating project...\n Project Name: %s\n Local: %s\n Remote: %s\n Setup website: %v\n",
			testName,
			testLocation,
			testRemote,
			true,
		),
	},
	{
		TestNumber:  4,
		Description: "All arguments entered in another different order.",
		Input:       []string{"-s", "-n", testName, "-r", testRemote, "-d", testLocation},
		Want: fmt.Sprintf(
			"Generating project...\n Project Name: %s\n Local: %s\n Remote: %s\n Setup website: %v\n",
			testName,
			testLocation,
			testRemote,
			true,
		),
	},
	{
		TestNumber:  5,
		Description: "No -s flag.",
		Input:       []string{"-n", testName, "-r", testRemote, "-d", testLocation},
		Want: fmt.Sprintf(
			"Generating project...\n Project Name: %s\n Local: %s\n Remote: %s\n Setup website: %v\n",
			testName,
			testLocation,
			testRemote,
			false,
		),
	},
	{
		TestNumber:  6,
		Description: "Not all required flags specified",
		Input:       []string{"-r", testRemote, "-d", testLocation},
		Want:        fmt.Sprintf("%s", errMsgTestName),
	},
	{
		TestNumber:  7,
		Description: "Not all required flags specified",
		Input:       []string{"-d", testLocation},
		Want:        fmt.Sprintf("%s%s", errMsgTestName, errMsgTestRemote),
	},
	{
		TestNumber:  8,
		Description: "Flag passed but no argument.",
		Input:       []string{"-d"},
		Want:        "",
	},
	{
		TestNumber:  9,
		Description: "Undefined flag",
		Input:       []string{"-x"},
		Want:        "",
	},
}

// test func validateOpts()
var testCasesValidateOpts = []testCaseValidateOpts{
	{
		TestNumber:  1,
		Description: "Project name and location on disk not specified.",
		Input:       opts{remoteURL: testRemote},
		Want:        fmt.Sprintf("%s%s", errMsgTestName, errMsgTestLocation),
	},
	{
		TestNumber:  2,
		Description: "Project name and remote repository url not specified.",
		Input:       opts{locationOnDisk: testLocation},
		Want:        fmt.Sprintf("%s%s", errMsgTestName, errMsgTestRemote),
	},
	{
		TestNumber:  3,
		Description: "Project location on disk and remote repository url not specified.",
		Input:       opts{nameOfProject: testName},
		Want:        fmt.Sprintf("%s%s", errMsgTestLocation, errMsgTestRemote),
	},
	{
		TestNumber:  4,
		Description: "Project location on disk not specified.",
		Input:       opts{nameOfProject: testName, remoteURL: testRemote},
		Want:        fmt.Sprintf("%s", errMsgTestLocation),
	},
	{
		TestNumber:  5,
		Description: "Project name not specified.",
		Input:       opts{locationOnDisk: testLocation, remoteURL: testRemote},
		Want:        fmt.Sprintf("%s", errMsgTestName),
	},
	{
		TestNumber:  6,
		Description: "Project repository URL not specified.",
		Input:       opts{nameOfProject: testName, locationOnDisk: testLocation},
		Want:        fmt.Sprintf("%s", errMsgTestRemote),
	},
}
