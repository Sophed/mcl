package main

func printInfo(profile *Profile) {

	uuid := hyphen(profile.UUID)

	print("┌")
	line(len(uuid) + 9)
	println("┐")

	print("│ Profile  ", profile.Name)

	for i := 0; i < len(uuid)+1-len(profile.Name); i++ {
		print(" ")
	}
	println("│")

	println("│ UUID    ", uuid, "│")

	print("└") // ├
	line(len(uuid) + 9)
	println("┘") // ┤

}

func line(n int) {
	for i := 0; i < n+2; i++ {
		print("─")
	}
}

func hyphen(uuid string) string {
	if len(uuid) != 32 {
		return ""
	}
	return uuid[:8] + "-" + uuid[8:12] + "-" + uuid[12:16] + "-" + uuid[16:20] + "-" + uuid[20:]
}
