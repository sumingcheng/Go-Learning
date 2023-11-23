package main

import (
	"fmt"
	"sync"
	"time"
)

type Entry struct {
	Amount    float64
	Category  string
	CreatedAt time.Time
}

type Ledger struct {
	entries  []Entry
	commands chan Command
	mu       sync.Mutex
}

type Command struct {
	Type     string
	Amount   float64
	Category string
}

func NewLedger() *Ledger {
	return &Ledger{
		commands: make(chan Command),
	}
}

func (l *Ledger) ProcessCommands(wg *sync.WaitGroup) {
	defer wg.Done()
	for cmd := range l.commands {
		switch cmd.Type {
		case "add":
			l.AddEntry(cmd.Amount, cmd.Category)
		case "list":
			l.ListEntries()
		}
	}
}

func (l *Ledger) AddEntry(amount float64, category string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := Entry{
		Amount:    amount,
		Category:  category,
		CreatedAt: time.Now(),
	}
	l.entries = append(l.entries, entry)
	fmt.Println("账目已添加")
}

func (l *Ledger) ListEntries() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.entries) == 0 {
		fmt.Println("账本为空")
		return
	}
	fmt.Println("账目列表:")
	for i, entry := range l.entries {
		fmt.Printf("%d. 金额: %.2f, 分类: %s, 时间: %s\n", i+1, entry.Amount, entry.Category, entry.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
