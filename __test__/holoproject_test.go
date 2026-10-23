package test

// import (
// 	"errors"
// 	"os"
// 	"path/filepath"
// 	"testing"

// 	"main/constants"
// )

// func withTempRoot(t *testing.T) {
// 	t.Helper()

// 	oldRoot, oldOverride := constants.Root, constants.RootOverride
// 	t.Cleanup(func() {
// 		constants.Root = oldRoot
// 		constants.RootOverride = oldOverride
// 	})
// 	constants.RootOverride = t.TempDir()
// }

// func TestExecuteHeadlessCreatesBat(t *testing.T) {
// 	withTempRoot(t)

// 	if err := executeHeadless("bat hello p/echo hi", false); err != nil {
// 		t.Fatalf("executeHeadless: %v", err)
// 	}

// 	data, err := os.ReadFile(filepath.Join(constants.Root, "cmd", "hello.bat"))
// 	if err != nil {
// 		t.Fatalf("read bat: %v", err)
// 	}
// 	if string(data) != "echo hi" {
// 		t.Errorf("content = %q, want %q", string(data), "echo hi")
// 	}
// }

// func TestExecuteHeadlessRejectsInteractive(t *testing.T) {
// 	withTempRoot(t)

// 	if err := executeHeadless("list", false); !errors.Is(err, errInteractive) {
// 		t.Errorf("error = %v, want %v", err, errInteractive)
// 	}
// }

// func TestExecuteHeadlessUnknownCommand(t *testing.T) {
// 	withTempRoot(t)

// 	if err := executeHeadless("nope", false); err == nil {
// 		t.Error("expected error for unknown command")
// 	}
// }

// func TestExecuteHeadlessRemoveWithYes(t *testing.T) {
// 	withTempRoot(t)

// 	if err := executeHeadless("bat hello p/echo hi", false); err != nil {
// 		t.Fatalf("create: %v", err)
// 	}

// 	if err := executeHeadless("rbat hello", false); !errors.Is(err, errInteractive) {
// 		t.Errorf("without --yes: error = %v, want %v", err, errInteractive)
// 	}

// 	if err := executeHeadless("rbat hello", true); err != nil {
// 		t.Fatalf("remove with --yes: %v", err)
// 	}

// 	if _, err := os.Stat(filepath.Join(constants.Root, "cmd", "hello.bat")); !os.IsNotExist(err) {
// 		t.Errorf("bat still exists: %v", err)
// 	}
// }
