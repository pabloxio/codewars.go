// https://www.codewars.com/kata/550f22f4d758534c1100025a/train/go

package main

func isReasonable(currentDir, nextDir string) bool {
	switch currentDir {
	case "NORTH":
		if nextDir != "SOUTH" {
			return true
		}
	case "SOUTH":
		if nextDir != "NORTH" {
			return true
		}
	case "WEST":
		if nextDir != "EAST" {
			return true
		}
	case "EAST":
		if nextDir != "WEST" {
			return true
		}
	}

	return false
}

func DirReduc(dirs []string) []string {
	if len(dirs) < 2 {
		return dirs
	}

	reduction := DirReduc(dirs[1:])
	if len(reduction) == 0 {
		return dirs[0:1]
	}

	if isReasonable(dirs[0], reduction[0]) {
		return append(dirs[0:1], reduction...)
	}

	return reduction[1:]
}
