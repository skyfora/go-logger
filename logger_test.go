package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInitJSONStdout(t *testing.T) {
	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	Init(Logger{
		WithStdout:  true,
		ServiceName: "test-service",
		Environment: "test",
	})
	Info("hello", zap.String("key", "value"))

	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}

	var entry map[string]any
	if err := json.Unmarshal(buf.Bytes(), &entry); err != nil {
		t.Fatalf("expected JSON log line, got %q: %v", buf.String(), err)
	}

	for _, key := range []string{"timestamp", "level", "msg", "service", "env", "key"} {
		if _, ok := entry[key]; !ok {
			t.Fatalf("missing %q in log entry: %v", key, entry)
		}
	}
	if entry["msg"] != "hello" {
		t.Fatalf("msg = %v, want hello", entry["msg"])
	}
	if entry["service"] != "test-service" {
		t.Fatalf("service = %v, want test-service", entry["service"])
	}
	if entry["env"] != "test" {
		t.Fatalf("env = %v, want test", entry["env"])
	}
	if entry["key"] != "value" {
		t.Fatalf("key = %v, want value", entry["key"])
	}
}

func TestInitService(t *testing.T) {
	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	InitService("my-service", "dev", true)
	Debug("debug message")

	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}

	var entry map[string]any
	if err := json.Unmarshal(buf.Bytes(), &entry); err != nil {
		t.Fatalf("expected JSON log line, got %q: %v", buf.String(), err)
	}
	if entry["level"] != "debug" {
		t.Fatalf("level = %v, want debug", entry["level"])
	}
	if entry["service"] != "my-service" {
		t.Fatalf("service = %v, want my-service", entry["service"])
	}
}

func TestConsoleStdout(t *testing.T) {
	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	Init(Logger{
		WithStdout:    true,
		ConsoleStdout: true,
	})
	Info("console line")

	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	if json.Valid(buf.Bytes()) {
		t.Fatalf("expected console output, got JSON: %q", buf.String())
	}
}

func TestNewJSONStdoutEncoder(t *testing.T) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	core := NewJSONStdoutEncoder(encoderConfig, level)
	if core == nil {
		t.Fatal("expected non-nil core")
	}
	if !core.Enabled(zapcore.InfoLevel) {
		t.Fatal("expected info level to be enabled")
	}
}
