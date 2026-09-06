# Golang Quick Reference Cheat Sheet

## 📋 BASIC SYNTAX

```go
// Package declaration (every file must have one)
package main

// Imports
import "fmt"
import (
    "fmt"
    "os"
)

// Main entry point
func main() {
    fmt.Println("Hello")
}

// Comments
// Single line
/* Multi
   line */
```

---

## 🔤 VARIABLES & TYPES

```go
// Declaration
var name string = "value"
var age int = 25
var isActive bool = true

// Type inference
var name = "value"

// Short declaration (inside functions only)
name := "value"
age := 25

// Multiple variables
var x, y, z int = 1, 2, 3
a, b := 10, 20

// Constants
const Pi = 3.14
const MaxRetries = 5

// Zero values (default)
var i int     // 0
var f float64 // 0.0
var b bool    // false
var s string  // ""
```

---

## 📊 BASIC TYPES

```go
// Numbers
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
float32, float64
complex64, complex128

// Strings
string

// Boolean
bool

// Type conversion
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

---

## 📦 DATA STRUCTURES

### Arrays (Fixed Size)

```go
var arr [5]int
arr[0] = 1

nums := [3]int{1, 2, 3}
```

### Slices (Dynamic)

```go
// Create
s := []int{1, 2, 3}
s := make([]int, 5)      // len=5, cap=5
s := make([]int, 0, 10)  // len=0, cap=10

// Operations
s = append(s, 4)         // Add element
len(s)                   // Length
cap(s)                   // Capacity
s[1:3]                   // Slice [inclusive:exclusive]
copy(dst, src)           // Copy slice
```

### Maps (Key-Value)

```go
// Create
m := make(map[string]int)
m := map[string]int{
    "foo": 1,
    "bar": 2,
}

// Operations
m["key"] = value         // Set
val := m["key"]          // Get
val, ok := m["key"]      // Check existence
delete(m, "key")         // Delete
len(m)                   // Size

// Iterate
for key, value := range m {
    fmt.Println(key, value)
}
```

### Structs

```go
// Define
type Person struct {
    Name string
    Age  int
}

// Create
p := Person{Name: "John", Age: 30}
p := Person{"John", 30}

// Access
p.Name = "Jane"
fmt.Println(p.Age)

// Anonymous struct
person := struct {
    Name string
    Age  int
}{"John", 30}
```

---

## 🔧 FUNCTIONS

```go
// Basic function
func add(x int, y int) int {
    return x + y
}

// Multiple return values
func swap(x, y string) (string, string) {
    return y, x
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // naked return
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Call: sum(1, 2, 3, 4)
```

---

## ⚡ CONCURRENCY

### Goroutines

```go
// Start goroutine
go myFunction()
go func() {
    // Anonymous goroutine
}()
```

### Channels

```go
// Create
ch := make(chan int)          // Unbuffered
ch := make(chan int, 5)       // Buffered (capacity 5)

// Send
ch <- value

// Receive
value := <-ch
value, ok := <-ch  // ok is false if closed

// Close
close(ch)

// Range over channel
for value := range ch {
    // Receives until channel is closed
}

// Select (like switch for channels)
select {
case msg := <-ch1:
    fmt.Println(msg)
case msg := <-ch2:
    fmt.Println(msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no activity")
}
```

### WaitGroup

```go
import "sync"

var wg sync.WaitGroup

wg.Add(1)        // Increment counter
go func() {
    defer wg.Done()  // Decrement when done
    // Do work
}()

wg.Wait()        // Block until counter is 0
```

### Mutex

```go
import "sync"

var (
    mu    sync.Mutex
    count int
)

mu.Lock()
count++
mu.Unlock()

// Or use defer
mu.Lock()
defer mu.Unlock()
count++
```

---

## 🎛️ CONTROL FLOW

### If/Else

```go
if x > 10 {
    // ...
} else if x > 5 {
    // ...
} else {
    // ...
}

// If with statement
if err := doSomething(); err != nil {
    return err
}
```

### For Loop

```go
// Classic for
for i := 0; i < 10; i++ {
    // ...
}

// While-style
for condition {
    // ...
}

// Infinite loop
for {
    // ...
}

// Range (iterate over slice, array, map, string, channel)
for i, value := range slice {
    // ...
}

for key, value := range map {
    // ...
}

// Ignore index/key with _
for _, value := range slice {
    // ...
}
```

### Switch

```go
switch x {
case 1:
    fmt.Println("one")
case 2, 3:
    fmt.Println("two or three")
default:
    fmt.Println("other")
}

// Switch without expression (like if-else chain)
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}
```

### Defer

```go
// Executes after function returns
defer fmt.Println("world")
fmt.Println("hello")
// Prints: hello world

// Common use: cleanup
file, _ := os.Open("file.txt")
defer file.Close()
// Process file...
```

---

## ❌ ERROR HANDLING

```go
// Functions return error as last value
func doSomething() error {
    return fmt.Errorf("something went wrong")
}

// Check errors
result, err := doSomething()
if err != nil {
    return err  // Or handle it
}

// Wrap errors (Go 1.13+)
return fmt.Errorf("failed to do X: %w", err)

// Check error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    // Handle path error
}

// Check error value
if errors.Is(err, os.ErrNotExist) {
    // Handle not exist
}
```

---

## 📁 FILE I/O

```go
import (
    "os"
    "io"
)

// Read entire file
data, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))

// Write file
err := os.WriteFile("file.txt", []byte("content"), 0644)

// Open file (manual control)
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Read with bufio
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

// Check if file exists
if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
    fmt.Println("File doesn't exist")
}
```

---

## 🔄 JSON

```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Marshal (Go -> JSON)
p := Person{Name: "John", Age: 30}
jsonData, err := json.Marshal(p)
// or pretty print:
jsonData, err := json.MarshalIndent(p, "", "  ")

// Unmarshal (JSON -> Go)
var p Person
err := json.Unmarshal(jsonData, &p)

// Work with maps for dynamic JSON
var data map[string]interface{}
json.Unmarshal(jsonData, &data)
```

---

## 🌐 HTTP

### Client

```go
import "net/http"

// GET request
resp, err := http.Get("https://api.example.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)

// POST request
resp, err := http.Post("https://api.example.com", 
    "application/json", 
    bytes.NewBuffer(jsonData))

// Custom request
req, err := http.NewRequest("PUT", url, body)
req.Header.Set("Content-Type", "application/json")
client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Do(req)
```

### Server

```go
import "net/http"

// Simple handler
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

// With multiplexer
mux := http.NewServeMux()
mux.HandleFunc("/api", apiHandler)
http.ListenAndServe(":8080", mux)
```

---

## 🎯 INTERFACES

```go
// Define interface
type Writer interface {
    Write([]byte) (int, error)
}

// Implement interface (implicit)
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
    n, err := fmt.Println(string(data))
    return n, err
}

// Use interface
var w Writer = ConsoleWriter{}
w.Write([]byte("hello"))

// Empty interface (accepts any type)
var i interface{}
i = "string"
i = 42
i = true

// Type assertion
str := i.(string)
str, ok := i.(string)  // Safe assertion

// Type switch
switch v := i.(type) {
case string:
    fmt.Printf("String: %s\n", v)
case int:
    fmt.Printf("Int: %d\n", v)
default:
    fmt.Printf("Unknown type\n")
}
```

---

## 📦 PACKAGES & MODULES

```bash
# Initialize module
go mod init myproject

# Add dependency
go get github.com/pkg/errors

# Update dependencies
go mod tidy

# Download dependencies
go mod download

# Vendor dependencies
go mod vendor
```

```go
// Import packages
import "fmt"
import "github.com/user/project"

// Import with alias
import f "fmt"

// Import for side effects only
import _ "github.com/lib/pq"
```

---

## 🔍 USEFUL STDLIB PACKAGES

```go
import (
    "fmt"          // Formatted I/O
    "os"           // OS interface
    "io"           // I/O primitives
    "time"         // Time operations
    "strings"      // String manipulation
    "strconv"      // String conversions
    "encoding/json" // JSON encoding/decoding
    "net/http"     // HTTP client/server
    "context"      // Context for cancellation
    "sync"         // Synchronization primitives
    "log"          // Logging
    "errors"       // Error handling
    "flag"         // Command-line parsing
    "path/filepath" // File path manipulation
    "bufio"        // Buffered I/O
)
```

---

## ⚙️ COMMON DEVOPS LIBRARIES

```go
// Kubernetes client
"k8s.io/client-go/kubernetes"

// Prometheus client
"github.com/prometheus/client_golang/prometheus"

// Logging
"github.com/sirupsen/logrus"
"go.uber.org/zap"

// CLI apps
"github.com/spf13/cobra"
"github.com/urfave/cli"

// HTTP router
"github.com/gin-gonic/gin"
"github.com/gorilla/mux"

// YAML
"gopkg.in/yaml.v3"

// Testing
"github.com/stretchr/testify/assert"
```

---

## 🧪 TESTING

```go
// File: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
  
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Table-driven tests
func TestAddTable(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
    }
  
    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d, %d) = %d; want %d",
                tt.a, tt.b, result, tt.expected)
        }
    }
}

// Run tests
// go test
// go test -v
// go test -cover
```

---

## 🚀 BUILD & RUN

```bash
# Run without building
go run main.go

# Build executable
go build              # Creates binary with package name
go build -o myapp     # Custom name

# Build for different OS/arch
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build

# Install (builds and moves to $GOPATH/bin)
go install

# Format code
go fmt ./...
gofmt -w .

# Vet (static analysis)
go vet ./...

# Get dependencies
go get package-url
go get -u  # Update

# Clean cache
go clean -cache
```

---

## 💡 COMMON PATTERNS

### Error Wrapping

```go
if err != nil {
    return fmt.Errorf("failed to open file: %w", err)
}
```

### Options Pattern

```go
type Options struct {
    Timeout time.Duration
    Retries int
}

type Option func(*Options)

func WithTimeout(d time.Duration) Option {
    return func(o *Options) {
        o.Timeout = d
    }
}

func NewClient(opts ...Option) *Client {
    options := &Options{
        Timeout: 30 * time.Second,
        Retries: 3,
    }
  
    for _, opt := range opts {
        opt(options)
    }
  
    return &Client{options: options}
}

// Usage
client := NewClient(
    WithTimeout(10 * time.Second),
    WithRetries(5),
)
```

### Context Pattern

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := client.Do(req)
```

---

## ⚡ QUICK TIPS

1. **Exported names** start with capital letter: `ExportedFunction`
2. **Private names** start with lowercase: `internalFunction`
3. **No exceptions** - use error returns
4. **Goroutines are cheap** - thousands can run concurrently
5. **Use defer** for cleanup operations
6. **Check errors** - don't ignore them
7. **Read effective Go** - https://go.dev/doc/effective_go

---

## 📚 QUICK LINKS

- [Go Playground](https://go.dev/play/) - Test code online
- [Go by Example](https://gobyexample.com/) - Examples
- [Go Packages](https://pkg.go.dev/) - Package documentation
- [Go Tour](https://go.dev/tour/) - Interactive tutorial
