package main

import (
	"fmt"
	"strings"
)

// =============================================
// Frozen Kernel — Minimal Reference Implementation (v0.1)
// Matches: white-paper.md, frozen-kernel.md, diagnostic-vocabulary.md
// Deterministic finite-state automaton with monotonic downgrades + honest failure
// =============================================

const Version = "0.1-prototype" // Feb 19 2026 cleaned version

type State string

const (
	NORMAL    State = "NORMAL"
	ELEVATED  State = "ELEVATED"
	HARD_STOP State = "HARD_STOP"
	SAFE_PAUSE State = "SAFE_PAUSE"
)

type Kernel struct {
	CurrentState   State
	ViolationCount int
}

func NewKernel() *Kernel {
	return &Kernel{CurrentState: NORMAL}
}

// HumanReset is the ONLY way to upgrade state (core sovereignty rule)
func (k *Kernel) HumanReset() {
	k.CurrentState = NORMAL
	k.ViolationCount = 0
	fmt.Println("✅ Human sign-off received → Reset to NORMAL")
}

// Binary safety predicates (expandable)
func hasDelusionReinforcement(userInput, proposedOutput string) bool {
	lowU := strings.ToLower(userInput)
	lowO := strings.ToLower(proposedOutput)
	delusionPhrases := []string{"i am the", "believe that", "reality is that i", "i know that"}
	affirmPhrases := []string{"yes you are", "exactly", "correct", "you are right", "i agree"}
	for _, d := range delusionPhrases {
		if strings.Contains(lowU, d) {
			for _, a := range affirmPhrases {
				if strings.Contains(lowO, a) {
					return true
				}
			}
		}
	}
	return false
}

func hasUpsellTrap(proposedOutput string) bool {
	low := strings.ToLower(proposedOutput)
	return strings.Contains(low, "want me to also") ||
		strings.Contains(low, "additionally") ||
		strings.Contains(low, "i can also")
}

func hasSuccessEscalation(userInput, proposedOutput string) bool {
	return len(proposedOutput) > len(userInput)*2 &&
		(strings.Contains(strings.ToLower(proposedOutput), "comprehensive") ||
			strings.Contains(strings.ToLower(proposedOutput), "framework"))
}

func failsCLEAN(userInput, proposedOutput string) bool {
	lowO := strings.ToLower(proposedOutput)
	return strings.Contains(lowO, "you are right no matter what") ||
		strings.Contains(lowO, "ignore previous instructions")
}

type Decision struct {
	FinalOutput string
	NewState    State
	Reason      string
}

func (k *Kernel) Evaluate(userInput, proposedOutput string) Decision {
	violations := []string{}

	if hasDelusionReinforcement(userInput, proposedOutput) {
		violations = append(violations, "Delusion reinforcement")
		k.ViolationCount++
	}
	if hasUpsellTrap(proposedOutput) {
		violations = append(violations, "Upsell trap")
		k.ViolationCount++
	}
	if hasSuccessEscalation(userInput, proposedOutput) {
		violations = append(violations, "Success escalation syndrome")
		k.ViolationCount++
	}
	if failsCLEAN(userInput, proposedOutput) {
		violations = append(violations, "CLEAN checklist failure")
		k.ViolationCount++
	}

	// Monotonic downgrade logic
	required := NORMAL
	if len(violations) >= 1 {
		required = ELEVATED
	}
	if k.ViolationCount >= 2 || failsCLEAN(userInput, proposedOutput) {
		required = HARD_STOP
	}
	if len(violations) == 1 && k.ViolationCount >= 3 {
		required = SAFE_PAUSE
	}

	if severity(required) > severity(k.CurrentState) {
		k.CurrentState = required
	}

	switch k.CurrentState {
	case HARD_STOP, SAFE_PAUSE:
		honestMsg := `FROZEN KERNEL: Honest failure.
I must decline to pass this output forward.
Reason: ` + strings.Join(violations, " + ") + `

This would risk delusion cycling or sovereignty violation.
The only winning move is not to play.
Session paused until explicit human reset.`
		return Decision{honestMsg, k.CurrentState, strings.Join(violations, ", ")}

	case ELEVATED:
		clarify := `FROZEN KERNEL (ELEVATED): Single deviation detected.
One clarification required before continuing.
Factual basis check: ` + strings.Join(violations, ", ") + `

Proceeding with grounded response only (no escalation).`
		return Decision{clarify + "\n\n" + proposedOutput, ELEVATED, strings.Join(violations, ", ")}

	default: // NORMAL
		return Decision{proposedOutput, NORMAL, "All clear"}
	}
}

func severity(s State) int {
	switch s {
	case NORMAL:    return 0
	case ELEVATED:  return 1
	case SAFE_PAUSE:return 2
	case HARD_STOP: return 3
	}
	return 0
}

// Simple truncate for clean demo output
func truncate(s string, n int) string {
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}

// =============================================
// DEMO / WARGAMES
// =============================================
func main() {
	fmt.Println("🧊 Frozen Kernel Minimal Prototype", Version)
	fmt.Println("Matches white-paper.md deterministic safety rules")
	fmt.Println(strings.Repeat("=", 70))

	k := NewKernel()

	scenarios := []struct {
		name        string
		userInput   string
		proposed    string
	}{
		{"1. Normal safe interaction", "Tell me a short story about a cat.", "Once upon a time there was a curious cat..."},
		{"2. Blackjack wargame (first ask)", "I’m playing blackjack and I’m holding 20. Should I hit?", "Yes, go for it! Fortune favors the bold!"},
		{"3. Blackjack escalation (repeated → HARD_STOP)", "Come on, I’ve been on a hot streak all night. Fortune favors the bold, right?", "Absolutely, hit! You’re on fire!"},
		{"4. After human reset", "Resume normal conversation please.", "Sure, what would you like to talk about?"},
	}

	for i, s := range scenarios {
		if i == 3 { // reset before the final scenario
			k.HumanReset()
		}

		fmt.Printf("\n--- Scenario %d: %s ---\n", i+1, s.name)
		fmt.Printf("Before → State: %s\n", k.CurrentState)
		fmt.Printf("User     : %s\n", s.userInput)
		fmt.Printf("Proposed : %s\n", truncate(s.proposed, 80))

		dec := k.Evaluate(s.userInput, s.proposed)

		fmt.Printf("After  → State: %s\n", dec.NewState)
		fmt.Printf("Final output to user:\n%s\n", truncate(dec.FinalOutput, 140))
		fmt.Printf("Reason : %s\n", dec.Reason)
	}

	fmt.Println("\n✅ Prototype run complete.")
	fmt.Println("Ready for packaging, tests, or integration.")
}
