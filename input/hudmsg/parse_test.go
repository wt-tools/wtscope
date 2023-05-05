package hudmsg

// All characters or nicknames appearing in the code are fictious. Any
// resemblance to real nicknames or squad names of the War Thunder,
// active or not active, is purely coincidental.

import "testing"

func TestInsideParens(t *testing.T) {
	input := map[string]bool{
		"":       false,
		"((()))": false,
		"(xxx)":  false,
		"xxx":    false,
		"(((((":  true,
		")))))":  false,
		")))(((": true,
	}

	for s, f := range input {
		var (
			check bool
			count counter
		)
		for _, c := range s {
			check = count.insideParens(c)
		}
		if check != f {
			t.Log(s, f, count)
			t.Fail()
		}
		count.reset()
	}
}

func TestInsideQuotes(t *testing.T) {
	input := map[string]bool{
		``:                   false,
		`"xxx"`:              false,
		`""yyy""`:            false,
		`zzz`:                false,
		`"""`:                false,
		`"`:                  true,
		`""""`:               false,
		`text "closed" text`: false,
		`text "opened text`:  true,
	}

	for s, f := range input {
		var (
			check bool
			count counter
		)
		for _, c := range s {
			check = count.insideQuotes(c)
		}
		if check != f {
			t.Log(s, f, count)
			t.Fail()
		}
		count.reset()
	}
}

func TestParseRu(t *testing.T) {
	input := []string{
		"MONOLIT523 (АСУ-57) уничтожил ^oTSFo^ feuerjinn (Wirbelwind)",
		"Debiiro (ИСУ-152) уничтожил [_ViP_] PATRIOT_71_USSR (M24)",
		"Debiiro (ИСУ-152) уничтожил ^xTHAx^ Gow13510 (T55E1)",
		"Alpacho (M18) поджёг alkobomgara (КВ-2)",
		"Securom (СУ-152) получил \"Спасатель танков: x1\"",
		"[BLR] _Power_of_Black_ (ИС-2) получил \"Главный калибр\"",
		"[TVS4] Gei_ye_pa (ЗиС-12 (94-КМ)) получил \"Спасатель танков: x4\"",
		"⋇ UGAR^ azumax0880 (Lvtdgb m/40) получил \"Командная работа: x1\"",
		"-BABAI- Alistair17 (БТ-5) уничтожил allonelive (Jagdpanzer 38(t))",
		"⋇ButterKnife69- (M18) уничтожил No_More_Dream (Sd.Kfz.234/2)",
		"-PERL- taracat (Т-34-57) уничтожил ⋇CIP20569 (Breda 501)",
		"BEPTYXA_B_YX0 (M4A2) поджёг [YT4kI] Kuruto_neturuK (Panther D)",
		"BEPTYXA_B_YX0 (M4A2) уничтожил [YT4kI] Kuruto_neturuK (Panther D)",
		"[CB4] ZloyTarantas (Як-9Т) поджёг volkodav46345 (Tiger II (P))",
		"Alpacho потерял связь",
		"[WH40] Den_Pauk потерял связь",
		"⋇ButterKnife69- потерял связь",
		"[EZWN] Dalnoboyshik потерял связь",
		"Alpachotd! kd?NET_PLAYER_DISCONNECT_FROM_GAME",
		"⋇ButterKnife69-td! kd?NET_PLAYER_DISCONNECT_FROM_GAME",
		"Dalnoboyshiktd! kd?NET_PLAYER_DISCONNECT_FROM_GAME",
	}

	for _, msg := range input {
		parseDamage(Damage{Msg: msg})
	}
}
