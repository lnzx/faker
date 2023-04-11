package useragent

import (
	"bufio"
	"embed"
	_ "embed"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	Chrome  = "chrome"
	Firefox = "firefox"
	Edge    = "edge"
)

var defaults = []string{Chrome, Firefox, Edge}

//go:embed chrome
var chrome embed.FS

//go:embed firefox
var firefox embed.FS

//go:embed edge
var edge embed.FS

type UserAgent struct {
	browsers []string
	chrome   []string
	firefox  []string
	edge     []string
	all      []string
}

func New(browsers ...string) *UserAgent {
	ua := &UserAgent{}
	if len(browsers) > 0 {
		defaults = browsers
	}
	for _, s := range defaults {
		switch s {
		case Chrome:
			ua.chrome = fill(chrome, Chrome)
			ua.all = append(ua.all, ua.chrome...)
		case Firefox:
			ua.firefox = fill(firefox, Firefox)
			ua.all = append(ua.all, ua.firefox...)
		case Edge:
			ua.edge = fill(edge, Edge)
			ua.all = append(ua.all, ua.edge...)
		default:
			log.Fatalln("not supported:", s)
		}
	}

	return ua
}

func fill(e embed.FS, name string) (data []string) {
	f, err := e.Open(name)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if text := strings.TrimSpace(scanner.Text()); text != "" {
			data = append(data, text)
		}
	}
	return
}

func randomValue(arr []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(arr))
	return arr[index]
}

func (f *UserAgent) Chrome() string {
	return randomValue(f.chrome)
}

func (f *UserAgent) Firefox() string {
	return randomValue(f.firefox)
}

func (f *UserAgent) Edge() string {
	return randomValue(f.edge)
}

func (f *UserAgent) Random() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(f.all))
	return f.all[index]
}
