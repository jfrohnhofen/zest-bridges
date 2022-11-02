package zest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Light interface {
	Channels() []uint16
	Properties() map[string]uint16
}

func Interactive(lights map[string]Light) {
	go func() {
		reader := bufio.NewReader(os.Stdin)
		var selected Light
		var property uint16

		for {
			fmt.Print("> ")
			cmd, _ := reader.ReadString('\n')
			cmd = strings.TrimSuffix(cmd, "\n")
			for _, cmd := range strings.Split(cmd, " ") {
				if cmd == "clear" {
					mu.Lock()
					for i := 0; i < 512; i++ {
						dmxFrame[i] = 0
					}
					mu.Unlock()
				}
				if cmd == "print" || cmd == "reset" {
					for ch, val := range dmxFrame {
						if val == 0 {
							continue
						}
						for name, light := range lights {
							for prop, offset := range light.Properties() {
								if light.Channels()[0]+offset == uint16(ch+1) {
									fmt.Printf("%s.%s = %d\n", name, prop, val)
								}
							}
						}
					}
				}
				for name, light := range lights {
					if strings.HasPrefix(name, cmd) {
						selected = light
						property = 0
						fmt.Println("fixture:", name)
						break
					}
				}
				if selected != nil {
					for prop, offset := range selected.Properties() {
						if strings.HasPrefix(prop, cmd) {
							property = offset
							fmt.Println("property:", prop)
						}
					}
				}
				if val, err := strconv.Atoi(cmd); err == nil && selected != nil {
					if selected != nil {
						for _, ch := range selected.Channels() {
							set(ch+property, uint8(val))
						}
					}
				}
			}
		}
	}()

	runDmxLoop()
}

func set(ch uint16, val uint8) {
	mu.Lock()
	dmxFrame[ch-1] = val
	mu.Unlock()
}
