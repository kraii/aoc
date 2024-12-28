package twenty

import (
	"aoc"
	"aoc/maths"
	"fmt"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"log"
	"strconv"
	"strings"
)

type signal int

const (
	nothing   signal = 0
	lowPulse  signal = 1
	highPulse signal = 2
)

type emission struct {
	source string
	sig    signal
}

type module interface {
	receive(c *circuit, from string, pulse signal) signal
	outConnections() []string
}

type circuit struct {
	modules     map[string]module
	connections map[string][]string
}

type flipFlop struct {
	name        string
	on          bool
	connections []string
}

func (f *flipFlop) outConnections() []string {
	return f.connections
}

func (f *flipFlop) receive(_ *circuit, _ string, pulse signal) signal {
	if pulse == highPulse {
		return nothing
	}
	if f.on {
		f.on = false
		return lowPulse
	} else {
		f.on = true
		return highPulse
	}
}

type conjunction struct {
	name        string
	memory      map[string]signal
	connections []string
}

func (m *conjunction) outConnections() []string {
	return m.connections
}

func (m *conjunction) receive(c *circuit, from string, pulse signal) signal {
	inputs := c.connections[m.name]
	m.memory[from] = pulse
	if len(m.memory) != len(inputs) {
		return highPulse
	}

	for _, s := range m.memory {
		if s != highPulse {
			return highPulse
		}
	}

	// all are high so send low
	return lowPulse
}

type broadcaster struct {
	name        string
	connections []string
}

func (b *broadcaster) outConnections() []string {
	return b.connections
}

func (b *broadcaster) receive(_ *circuit, _ string, s signal) signal {
	return s
}

func calcPulses(file string, presses int) int {
	cir := parse(file)
	c := &cir

	totalHigh, totalLow := 0, 0
	for i := 0; i < presses; i++ {
		//println("--------press---------")
		emissions := []emission{{source: "broadcaster", sig: lowPulse}}
		totalLow++
		for len(emissions) > 0 {
			current := emissions[0]
			emissions = emissions[1:]
			currentModule := c.modules[current.source]
			for _, con := range currentModule.outConnections() {
				if current.sig == lowPulse {
					//println(current.source, "low->", con)
					totalLow++
				}
				if current.sig == highPulse {
					//println(current.source, "high->", con)
					totalHigh++
				}
				connectedModule, p := c.modules[con]
				if !p {
					continue
				}
				newSig := connectedModule.receive(c, current.source, current.sig)
				if newSig != nothing {
					emissions = append(emissions, emission{source: con, sig: newSig})
				}

			}
		}
	}
	return totalLow * totalHigh
}

// I made a visualization but was still stumped, ended up stealing idea from this reddit post:
// https://www.reddit.com/r/adventofcode/comments/18mmfxb/comment/ke5sgxs/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
func findRxLowSend(file string) int {
	cir := parse(file)
	c := &cir

	var results []int

	for _, topLevelConnection := range c.modules["broadcaster"].outConnections() {
		binaryS := ""
		toVisit := []string{topLevelConnection}

		for len(toVisit) > 0 {
			current := toVisit[0]
			toVisit = toVisit[1:]

			mod := c.modules[current]
			switch mod.(type) {
			case *flipFlop:
				if connectedToConjunction(c, mod.outConnections()) {
					binaryS = "1" + binaryS
				} else {
					binaryS = "0" + binaryS
				}
			default:
				binaryS = "0" + binaryS
			}

			for _, cName := range mod.outConnections() {
				conn := c.modules[cName]
				switch conn.(type) {
				case *flipFlop:
					toVisit = append(toVisit, cName)
				}
			}
		}
		i, err := strconv.ParseInt(binaryS, 2, 32)
		if err != nil {
			panic(err)
		}
		results = append(results, int(i))
	}
	return maths.LcmAll(results)
}

func connectedToConjunction(c *circuit, connections []string) bool {
	for _, connection := range connections {
		switch c.modules[connection].(type) {
		case *conjunction:
			return true
		}
	}
	return false
}

func createGraphImage() {
	c := parse("twenty/input.txt")
	g := graphviz.New()
	graph, err := g.Graph(graphviz.Directed)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		_ = g.Close()
	}()

	for name, mod := range c.modules {
		n := newNode(graph, name)
		for _, conn := range mod.outConnections() {
			cNode := newNode(graph, conn)
			_, err := graph.CreateEdge(fmt.Sprintf("%s-%s", name, conn), n, cNode)
			if err != nil {
				return
			}
		}

	}
	if err := g.RenderFilename(graph, graphviz.PNG, "graph.png"); err != nil {
		panic(err)
	}

}

func newNode(graph *cgraph.Graph, name string) *cgraph.Node {
	node, err := graph.CreateNode(name)
	if err != nil {
		panic(err)
	}
	return node
}

func parse(file string) circuit {
	scanner := aoc.OpenScanner(file)
	nodes := make(map[string]module)
	for scanner.Scan() {
		line := scanner.Text()
		iAndO := strings.Split(line, "->")
		namePart := strings.TrimSpace(iAndO[0])
		firstChar := namePart[0]
		output := buildOutput(iAndO[1])
		name := namePart[1:]
		if firstChar == '%' {
			nodes[name] = &flipFlop{name: name, on: false, connections: output}
		} else if firstChar == '&' {
			nodes[name] = &conjunction{name: name, memory: make(map[string]signal), connections: output}
		} else {
			nodes[namePart] = &broadcaster{name: namePart, connections: output}
		}
	}

	return circuit{modules: nodes, connections: buildConnections(nodes)}
}

func buildOutput(output string) []string {
	split := strings.Split(output, ",")
	var result = make([]string, len(split))
	for i, s := range split {
		result[i] = strings.TrimSpace(s)
	}
	return result
}

func buildConnections(modules map[string]module) map[string][]string {
	var result = make(map[string][]string)
	for n, m := range modules {
		for _, c := range m.outConnections() {
			result[c] = append(result[c], n)
		}
	}
	return result
}
