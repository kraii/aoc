package nineteen

import (
	"aoc"
	"strings"
)

const defaultTarget = "default"

type intRange struct {
	start, end int //inclusive
}

type rule struct {
	target    string
	lesser    bool // < or > ?
	threshold int
	action    string
}

type workflow []rule

type machinePart map[string]int

func sumAccepted(file string) int {
	workflows, parts := parse(file)

	total := 0
	for _, part := range parts {
		total += rating(workflows, part)
	}

	return total
}

func rating(workflows map[string]workflow, part machinePart) int {
	current := "in"
	for {
		wf := workflows[current]
	RuleCheck:
		for _, r := range wf {
			if matches(r, part) {
				switch r.action {
				case "A":
					return total(part)
				case "R":
					return 0
				default:
					current = r.action
					break RuleCheck
				}
			}
		}
	}
}

func total(part machinePart) int {
	sum := 0
	for _, v := range part {
		sum += v
	}
	return sum
}

func matches(r rule, part machinePart) bool {
	if r.target == defaultTarget {
		return true
	}
	i, p := part[r.target]
	if p && r.lesser {
		return i < r.threshold
	} else if p {
		return i > r.threshold
	}
	return false
}

func parseCond(cond string) (string, bool, int) {
	if strings.Contains(cond, "<") {
		parts := strings.Split(cond, "<")
		v := aoc.ToInt(parts[1])
		return parts[0], true, v
	} else if strings.Contains(cond, ">") {
		parts := strings.Split(cond, ">")
		v := aoc.ToInt(parts[1])
		return parts[0], false, v
	}
	panic("Don't know what to do with cond " + cond)
}

func parse(file string) (map[string]workflow, []machinePart) {
	workflows := make(map[string]workflow)
	var parts []machinePart
	scanner := aoc.OpenScanner(file)

	workflowsRead := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			workflowsRead = true
			continue
		}
		if workflowsRead {
			parts = append(parts, parsePart(line))
		} else {
			name, wf := parseWorkflow(line)
			workflows[name] = wf
		}
	}
	return workflows, parts
}

func parsePart(line string) machinePart {
	part := make(machinePart)
	components := strings.Split(line[1:len(line)-1], ",") // split without { and }
	for _, component := range components {
		kAndV := strings.Split(component, "=")
		part[kAndV[0]] = aoc.ToInt(kAndV[1])
	}
	return part
}

func parseWorkflow(line string) (string, workflow) {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == '{' || r == '}'
	})
	name := fields[0]
	rules := strings.Split(fields[1], ",")
	wf := make(workflow, len(rules))
	for i, r := range rules {
		if strings.Contains(r, ":") {
			split := strings.Split(r, ":")
			target, lesser, threshold := parseCond(split[0])
			wf[i] = rule{target: target, lesser: lesser, threshold: threshold, action: split[1]}
		} else {
			wf[i] = rule{target: defaultTarget, action: r}
		}
	}
	return name, wf
}

// part 2
type partRange = map[string]intRange

func sumPermutations(file string) int64 {
	workflows, _ := parse(file)
	parts := partRange{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
	return calcPermutations(workflows, "in", parts)
}

func calcPermutations(workflows map[string]workflow, current string, ranges partRange) int64 {
	if current == "A" {
		return permutationsOf(ranges)
	} else if current == "R" {
		return 0
	}
	wf := workflows[current]
	result := int64(0)
	rest := ranges
	for _, r := range wf {
		var passing partRange
		passing, rest = splitPassing(r, rest)
		result += calcPermutations(workflows, r.action, passing)
	}
	return result
}

func permutationsOf(rs partRange) int64 {
	perms := int64(1)
	for _, i := range rs {
		perms *= int64(i.end + 1 - i.start)
	}
	return perms
}

func splitPassing(rl rule, rn partRange) (partRange, partRange) {
	if rl.target == defaultTarget {
		return rn, rn
	}
	passing, failing := copyRange(rn), copyRange(rn)
	prev := rn[rl.target]
	if rl.lesser {
		passing[rl.target] = intRange{prev.start, rl.threshold - 1}
		failing[rl.target] = intRange{rl.threshold, prev.end}
	} else {
		passing[rl.target] = intRange{rl.threshold + 1, prev.end}
		failing[rl.target] = intRange{prev.start, rl.threshold}
	}
	return passing, failing
}

func copyRange(r partRange) partRange {
	c := make(partRange, len(r))
	for k, v := range r {
		c[k] = v
	}
	return c
}
