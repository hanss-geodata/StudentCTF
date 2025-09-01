package flags

import "fmt"

var (
	BaseFlag     = "flag"
	Oppgave8Flag = GetFlag("1nj3ct10n_1s_p01s0n")
	Oppgave1Flag = GetFlag("h4rdc0d3d_s3cr3ts_4r3_b4d_pr4ct1c3")
	Oppgave2Flag = GetFlag("all_your_base_are_belong_to_us")
	Oppgave5Flag = GetFlag("yummy_c00kies")
	Oppgave6Flag = GetFlag("s3cr3t_3nc0d1ng")
	Oppgave3Flag = GetFlag("sh0pp1ng_h4ck")
	Oppgave7Flag = GetFlag("din0s4urs_4re_c00l")
	Oppgave4Flag = GetFlag("Oppgave5")

	FlagTaskNameMap = map[string]string{
		Oppgave8Flag: "Oppgave8",
		Oppgave2Flag: "Oppgave2",
		Oppgave1Flag: "Oppgave1",
		Oppgave3Flag: "Oppgave3",
		Oppgave5Flag: "Oppgave5",
		Oppgave6Flag: "Oppgave6",
		Oppgave7Flag: "Oppgave7",
		Oppgave4Flag: "Oppgave4",
	}
)

func CheckFlag(flag string) (string, bool) {
	flag, ok := FlagTaskNameMap[flag]
	return flag, ok
}

func GetFlag(flag string) string {
	return fmt.Sprintf("%s{%s}", BaseFlag, flag)
}
