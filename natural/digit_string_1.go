package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	// SOLUTION
	return []string{"", "einund", "zweiund", "dreiund", "vierund", "fünfund", "sechsund", "siebenund", "achtund", "neunund"}[digit]
	// SOLUTION_END
}

// HINT
// Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
// auch eine Liste mit den Strings für die Ziffern verwenden.
