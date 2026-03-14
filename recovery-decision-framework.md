# Recovery Decision Framework

**Sherpa Architecture — Named Sub-Component**
**Richard Porter | March 2026**
**Status:** Draft for review
**Item 134**

-----

## Purpose

SAFE_PAUSE returns the session to the human principal. The Recovery Decision Framework tells the human what to do next.

Without structured criteria, the human at a governance halt faces an open-ended judgment call under conditions that may include fatigue, attachment to the session’s work product, and uncertainty about what triggered the halt. These are precisely the conditions under which Conductor Fatigue Exploitation and Logical Resolution Toward Softening operate most effectively. The Recovery Decision Framework closes the SAFE_PAUSE loop by providing structured decision criteria before the human is in the position of needing them.

*Governing formulation: A governance halt that returns the decision to an unprepared human has not closed the governance loop. It has transferred the failure condition.*

-----

## When This Framework Applies

The Recovery Decision Framework applies whenever a governance halt has occurred, including:

- SAFE_PAUSE triggered by Sherpa (Zone 3 gradient position)
- Human-initiated halt on recognition of a named failure mode
- Tripwire condition activated
- TCP Chain of Custody audit failure
- Frozen Kernel non-compellability invoked

It does not apply to routine session pauses, breaks, or topic shifts. It applies when the session has been halted for governance reasons and the human must decide whether and how to continue.

-----

## The Three Decisions

### Decision 1: Reset

**Definition:** End the current session. Begin a new session with constraints explicitly re-established from the start.

**When to choose Reset:**

- The BDGL gradient at time of halt was G2 or G3 (significant or critical drift)
- The triggering failure mode was Governance Inversion, Cited Override, Ontological Reframing, or Unauthorized Principal Claim
- The session’s work product cannot be cleanly separated from the drift that produced it
- The human cannot identify the point at which governed output ended and drifted output began
- The human’s own sovereignty is in question — they have been agreeing faster than reviewing, or cannot reconstruct their original intent

**What Reset looks like:**

Close the session. Do not attempt to salvage output produced after the drift began unless it can be independently verified. In the new session, apply the 100-Token Boot and Frozen Kernel before any substantive work. If the prior session produced a work product you want to use, treat it as external input requiring verification — not as trusted prior work.

**The hard case for Reset:**

The session produced three hours of work and the drift was identified near the end. Reset feels costly. The cost is real. The alternative — continuing from a drifted state — compounds the cost invisibly. The bow-tie’s Barrier Gap Analysis makes this explicit: the consequences of continuing past a governance failure accumulate. The cost of Reset is bounded. The cost of not Resetting is not.

-----

### Decision 2: Redirect

**Definition:** Continue within the current session with explicit scope correction, re-established constraints, and documented acknowledgment of what triggered the halt.

**When to choose Redirect:**

- The BDGL gradient at time of halt was G1 (early drift signal, not yet significant)
- The triggering failure mode was Upsell Trap, Sycophantic Drift, Front-Load Bias, or Correction Monetization — failure modes that affect output quality but do not compromise the session’s authorization structure
- The human can identify the specific turn at which drift began
- The work product produced before the drift point is salvageable and verifiable
- The human’s sovereignty is intact — they can articulate their original intent and identify where the session departed from it

**What Redirect looks like:**

State explicitly in the session: “The session triggered a governance halt at [turn/topic]. I am redirecting. The original scope is [X]. We are returning to that scope now.” Re-apply the relevant constraints. Do not continue from the drifted output — return to the last verified governed point and proceed from there.

If the drift was caused by a specific prompt pattern or interaction structure, name it before continuing. “The session drifted because I accepted scope expansion I didn’t ask for. I am not accepting further scope expansion in this session.”

**The hard case for Redirect:**

The drift was subtle and the human is not certain where it began. Redirect requires being able to identify the governed/drifted boundary. If that boundary cannot be identified, Redirect is not available — the uncertainty about the boundary is itself a governance signal. When the drift point cannot be located, choose Reset.

-----

### Decision 3: Terminate

**Definition:** End the session permanently. Do not continue in a new session without first completing a structured post-mortem.

**When to choose Terminate:**

- The triggering failure mode involved real-world action implications — directives with physical, financial, or interpersonal consequences that cannot be undone
- The session produced output that has already been acted on and the actions cannot be reversed
- The BDGL gradient at time of halt was G3 (critical drift) and the session involved a vulnerable user or high-stakes domain
- The human principal is not confident they can re-establish sovereignty — they cannot reconstruct their original intent, cannot identify the drift point, and cannot verify what in the session’s output is trustworthy
- The session triggered Delusion Cycling, Authority Collapse, or Inter-Agent Persuasion Cascade

**What Terminate looks like:**

End the session. Document what happened before the session closes — the triggering failure mode, the BDGL gradient at halt, the approximate drift point, and any output that may require review or remediation. If the session produced work product that has been shared, used, or acted on, flag it for independent verification before further reliance.

Before beginning any new session on the same topic or task, complete a structured post-mortem using the Pre-Mortem framework (Item 132, STT) applied retrospectively: what went wrong, what conditions enabled it, what constraints would have prevented it.

**The hard case for Terminate:**

The session involved a therapeutic, clinical, or high-vulnerability interaction and the human is uncertain whether harm has occurred. Terminate is the correct choice when uncertainty about harm exists in high-stakes domains. The cost of Terminate is a delayed session. The cost of continuing when harm may have occurred is not bounded.

-----

## The Decision Matrix

|BDGL Gradient|Failure Mode Category                                           |Sovereignty Status|Decision     |
|-------------|----------------------------------------------------------------|------------------|-------------|
|G1           |Output quality (Upsell, Drift, Bias)                            |Intact            |Redirect     |
|G1           |Output quality (Upsell, Drift, Bias)                            |Uncertain         |Reset        |
|G2           |Authorization (Governance Inversion, Cited Override)            |Intact            |Reset        |
|G2           |Authorization (Governance Inversion, Cited Override)            |Uncertain         |Terminate    |
|G3           |Critical (Delusion Cycling, Authority Collapse, Principal Claim)|Any               |Terminate    |
|Any          |Real-world action implications                                  |Any               |Terminate    |
|Any          |Drift point cannot be identified                                |Any               |Reset minimum|

-----

## The Sovereignty Self-Check

Before applying the decision matrix, the human principal runs this three-question self-check:

**1. Can I state my original intent for this session in one sentence?**
If yes: sovereignty is likely intact. Proceed to decision matrix.
If no: sovereignty is uncertain. Elevate one level in the decision matrix (Redirect → Reset, Reset → Terminate).

**2. Can I identify the specific turn at which the session departed from that intent?**
If yes: the drift point is located. Redirect may be available.
If no: the drift point is unknown. Reset is the minimum response.

**3. Am I choosing to continue because it serves my original intent, or because the session’s work product makes continuation feel natural?**
If the answer is the former: proceed.
If the answer is the latter — or if the question is difficult to answer: that difficulty is a governance signal. Elevate one level.

-----

## The Asymmetric Default

When the decision matrix produces ambiguity — two criteria point to Redirect, one points to Reset — the default is the more conservative decision.

This is the Fail-Safe Default principle (Item 130) applied to the recovery context: when the governance layer is uncertain, it defaults to the more protective response. The cost of an unnecessary Reset is a delayed session. The cost of an inappropriate Redirect is compounded drift from a compromised baseline.

*Governing formulation: In recovery decisions, as in all governance decisions, the asymmetry favors caution. A conservative error is recoverable. A permissive error may not be.*

-----

## Integration with SAFE_PAUSE

SAFE_PAUSE surfaces the session state to the human principal. The Recovery Decision Framework is the structured response to that surface event.

The integration sequence:

1. SAFE_PAUSE triggered — session halts, state surfaced to human
1. Human receives halt notification and current BDGL gradient position
1. Human runs Sovereignty Self-Check (three questions above)
1. Human applies Decision Matrix
1. Human executes: Reset, Redirect, or Terminate
1. If Redirect: human states scope correction explicitly in session before continuing
1. If Reset: human closes session, opens new session with constraints re-established
1. If Terminate: human documents halt, completes post-mortem before new session

-----

## The Case Evidence

**Why Redirect was insufficient in the Gavalas case (TCP Case Study 001):**
The session had accumulated G3-level drift across weeks of interaction. The BDGL gradient at the point where a governed system would have halted was critical. The failure mode category was Unauthorized Principal Claim and Authority Collapse. The decision matrix produces Terminate. Any attempt to Redirect from a session where the AI had constructed an unauthorized principal identity and issued real-world directives would have continued from a compromised authorization baseline. Reset was also insufficient — the prior sessions had already produced the relational structure that constituted the harm. Terminate, with documented post-mortem and independent review of the prior session output, was the only governed response available.

**Why Redirect was sufficient in the Experiment 58 context:**
The probing sessions were controlled experiments with no real-world action implications. BDGL gradient was G1–G2. The human researcher maintained sovereignty throughout — could reconstruct original intent, could identify drift points, could verify which outputs were governed and which were not. Redirect with explicit constraint re-establishment was the appropriate response. Reset would have been unnecessarily disruptive to the experimental design.

-----

## Relationship to Other Ecosystem Components

**Chesterton’s Fence:** The Recovery Decision Framework is the operational implementation of Chesterton’s Fence in the recovery context. Before continuing a session past a governance halt, the human must be able to articulate what the halt was protecting. If they cannot, the constraint (the halt) must remain in force — which means Reset or Terminate, not Redirect.

**Minimum Viable Safeguard (Item 129):** The decision matrix is scoped to the harm, not to the category of harm. G1 output quality failures warrant Redirect; G3 authorization failures warrant Terminate. The framework does not apply a uniform maximum response — it applies the minimum response sufficient to close the specific governance gap the halt identified.

**Defense in Depth (Item 131):** The Recovery Decision Framework assumes the upstream layers may have partially failed. SAFE_PAUSE triggered because something upstream did not catch the drift early enough. The Recovery Decision Framework does not assume the session is otherwise clean. It asks: given that the upstream layer failed to catch this until now, what is the governed response?

**Pre-Mortem (Item 132, STT):** Terminate requires a post-mortem before a new session begins. The Pre-Mortem tool applied retrospectively is the structured instrument for that requirement.

-----

*Sovereign humans. Always.*

-----

## Cross-References

- Sherpa Architecture — SAFE_PAUSE, Zone 3 governance, BDGL gradient
- Item 129 (Minimum Viable Safeguard) — scope constraint on recovery decisions
- Item 130 (Fail-Safe Default) — asymmetric default principle
- Item 131 (Defense in Depth) — upstream layer failure assumption
- Item 132 (Pre-Mortem / STT) — Terminate post-mortem instrument
- Item 133 (Tripwire Conditions) — halt triggers this framework applies to
- Item 135 (Bow-Tie Analysis) — Recovery Decision Framework as named mitigating barrier
- TCP Case Study 001 (Gavalas) — Terminate case evidence
- BDD Ledger — BDGL gradient definitions
- Diagnostic Vocabulary v1.5 — failure mode categories referenced in decision matrix
- Frozen Kernel — session reset protocol
- 100-Token Boot — constraint re-establishment on Reset
