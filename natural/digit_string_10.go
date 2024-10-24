package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	// SOLUTION
	return []string{"", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}[digit]
	// SOLUTION_END

}

// HINT
// Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
// auch eine Liste mit den Strings für die Ziffern verwenden.
