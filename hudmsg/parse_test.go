package hudmsg

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
		"Alpacho (M18) поджёг alkobomgara (КВ-2)",
		"Securom (СУ-152) получил \"Спасатель танков: x1\"",
		"[BLR] _Power_of_Black_ (ИС-2) получил \"Главный калибр\"",
		"[TVS4] Gei_ye_pa (ЗиС-12 (94-КМ)) получил \"Спасатель танков: x4\"",
		"Alpacho потерял связь",
		"Alpachotd! kd?NET_PLAYER_DISCONNECT_FROM_GAME",
	}

	for i, msg := range input {
		parseDamage(msg, uint(i))
	}
}
