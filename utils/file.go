package utils

import "io/ioutil"

func ListFiles(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

func AndroidCommands(s string) []string {
	return []string{
		"ping",
		"automation",
		"event",
		"geo",
		"gsm",
		"cdma",
		"crash",
		"crash-on-exit",
		"kill",
		"restart",
		"network",
		"grpc",
		"power",
		"quit|exit",
		"redir",
		"sms",
		"avd",
		"qemu",
		"sensor",
		"physics",
		"finger",
		"debug",
		"rotate",
		"screenrecord",
		"fold",
		"unfold",
		"posture",
		"multidisplay",
		"icebox",
		"nodraw",
		"resize-display",
	}
}
