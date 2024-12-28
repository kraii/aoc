package fifteen

import (
	"slices"
	"strconv"
	"strings"
)

func hash(s string) int {
	total := 0
	for _, r := range []rune(s) {
		total += int(r)
		total *= 17
		total = total % 256
	}
	return total
}

func sumHashSequence(s string) int {
	entries := strings.Split(strings.TrimSpace(s), ",")
	total := 0
	for _, entry := range entries {
		total += hash(entry)
	}
	return total
}

func initializeLenses(s string) int {
	entries := strings.Split(strings.TrimSpace(s), ",")
	boxes := make([][]lens, 256)
	for _, entry := range entries {
		if strings.Contains(entry, "=") {
			parts := strings.Split(entry, "=")
			label, length := parts[0], toInt(parts[1])
			i := hash(label)
			box := boxes[i]
			existingLoc := slices.IndexFunc(box, matchingLabel(label))
			if existingLoc == -1 {
				boxes[i] = append(box, lens{label, length})
			} else {
				box[existingLoc] = lens{label, length}
			}
		} else {
			parts := strings.Split(entry, "-")
			label := parts[0]
			i := hash(label)
			boxes[i] = slices.DeleteFunc(boxes[i], matchingLabel(label))
		}
	}

	total := 0
	for i, box := range boxes {
		for j, l := range box {
			total += (i + 1) * (j + 1) * l.focalLength
		}
	}
	return total
}

func matchingLabel(label string) func(a lens) bool {
	return func(a lens) bool {
		return a.label == label
	}
}

type lens struct {
	label       string
	focalLength int
}

func toInt(x string) int {
	num, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return num
}
