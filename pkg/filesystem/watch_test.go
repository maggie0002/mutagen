package filesystem

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestWatchModeUnmarshalPortable(t *testing.T) {
	var mode WatchMode
	if err := mode.UnmarshalText([]byte("portable")); err != nil {
		t.Fatal("unable to unmarshal text:", err)
	} else if mode != WatchMode_WatchPortable {
		t.Error("unmarshalled mode does not match expected")
	}
}

func TestWatchModeUnmarshalPOSIXRaw(t *testing.T) {
	var mode WatchMode
	if err := mode.UnmarshalText([]byte("force-poll")); err != nil {
		t.Fatal("unable to unmarshal text:", err)
	} else if mode != WatchMode_WatchForcePoll {
		t.Error("unmarshalled mode does not match expected")
	}
}

func TestWatchModeUnmarshalEmpty(t *testing.T) {
	var mode WatchMode
	if mode.UnmarshalText([]byte("")) == nil {
		t.Error("empty watch mode successfully unmarshalled")
	}
}

func TestWatchModeUnmarshalInvalid(t *testing.T) {
	var mode WatchMode
	if mode.UnmarshalText([]byte("invalid")) == nil {
		t.Error("invalid watch mode successfully unmarshalled")
	}
}

func TestWatchModeSupported(t *testing.T) {
	if WatchMode_WatchDefault.Supported() {
		t.Error("default watch mode considered supported")
	}
	if !WatchMode_WatchPortable.Supported() {
		t.Error("portable watch mode considered unsupported")
	}
	if !WatchMode_WatchForcePoll.Supported() {
		t.Error("force poll watch mode considered unsupported")
	}
	if (WatchMode_WatchForcePoll + 1).Supported() {
		t.Error("invalid watch mode considered supported")
	}
}

func TestWatchModeDescription(t *testing.T) {
	if description := WatchMode_WatchDefault.Description(); description != "Default" {
		t.Error("default watch mode description incorrect:", description, "!=", "Default")
	}
	if description := WatchMode_WatchPortable.Description(); description != "Portable" {
		t.Error("watch mode portable description incorrect:", description, "!=", "Portable")
	}
	if description := WatchMode_WatchForcePoll.Description(); description != "Force Poll" {
		t.Error("watch mode force poll description incorrect:", description, "!=", "Force Poll")
	}
	if description := (WatchMode_WatchForcePoll + 1).Description(); description != "Unknown" {
		t.Error("invalid watch mode description incorrect:", description, "!=", "Unknown")
	}
}

const (
	testWatchEstablishWait = 5 * time.Second
	testWatchTimeout       = 20 * time.Second
)

func testWatchCycle(path string, mode WatchMode) error {
	// Create a temporary directory in a subpath of the home directory and defer
	// its removal.
	directory, err := ioutil.TempDir(HomeDirectory, "mutagen_filesystem_watch")
	if err != nil {
		return errors.Wrap(err, "unable to create temporary directory")
	}
	defer os.RemoveAll(directory)

	// Create a cancellable watch context and defer its cancellation.
	watchContext, watchCancel := context.WithCancel(context.Background())
	defer watchCancel()

	// Create a watch event channel.
	events := make(chan struct{}, 1)

	// Start watching in a separate Goroutine.
	go Watch(watchContext, directory, events, WatchMode_WatchPortable, 1)

	// HACK: Wait long enough for the recursive watch to be established or the
	// initial polling to occur. The CI systems aren't as fast as things are
	// locally, so we have to be a little conservative.
	time.Sleep(testWatchEstablishWait)

	// Compute the test file path.
	testFilePath := filepath.Join(directory, "file")

	// Create a file inside the directory and wait for an event or timeout.
	if err := WriteFileAtomic(testFilePath, []byte{}, 0600); err != nil {
		return errors.New("unable to create file")
	}
	select {
	case <-events:
	case <-time.After(testWatchTimeout):
		return errors.New("create event not received in time")
	}

	// Modify a file inside the directory and wait for an event or timeout.
	if err := WriteFileAtomic(testFilePath, []byte{0, 0}, 0600); err != nil {
		return errors.New("unable to modify file")
	}
	select {
	case <-events:
	case <-time.After(testWatchTimeout):
		return errors.New("modify event not received in time")
	}

	// Remove a file inside the directory and wait for an event or timeout.
	if err := os.Remove(testFilePath); err != nil {
		return errors.New("unable to remove file")
	}
	select {
	case <-events:
	case <-time.After(testWatchTimeout):
		return errors.New("remove event not received in time")
	}

	// Success.
	return nil
}

func TestWatchPortable(t *testing.T) {
	// Skip this test on Windows for now, because the notify package seems to
	// have a data race there that the race detector catches.
	if runtime.GOOS == "windows" {
		t.Skip()
	}

	// Create a temporary directory in a subpath of the home directory and defer
	// its removal.
	directory, err := ioutil.TempDir(HomeDirectory, "mutagen_filesystem_watch")
	if err != nil {
		t.Fatal("unable to create temporary directory:", err)
	}
	defer os.RemoveAll(directory)

	// Run the test cycle.
	if err := testWatchCycle(directory, WatchMode_WatchPortable); err != nil {
		t.Fatal("watch cycle test failed:", err)
	}
}

func TestWatchForcePoll(t *testing.T) {
	// Create a temporary directory and defer its removal.
	directory, err := ioutil.TempDir("", "mutagen_filesystem_watch")
	if err != nil {
		t.Fatal("unable to create temporary directory:", err)
	}
	defer os.RemoveAll(directory)

	// Run the test cycle.
	if err := testWatchCycle(directory, WatchMode_WatchForcePoll); err != nil {
		t.Fatal("watch cycle test failed:", err)
	}
}
