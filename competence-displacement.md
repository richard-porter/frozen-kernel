# Competence Displacement: Barrier Brief

**Governance Design Requirement**
**Richard Porter | March 2026**
**Status:** Draft for review
**Bow-Tie Gap 1**

-----

## The Gap

The bow-tie analysis (Item 135) identified three right-side mitigating barriers that are not yet built. Two of those gaps — Recovery Decision Framework (Item 134) and Longitudinal Monitoring (Item 80) — have now been specified.

The third gap is different in kind.

Competence Displacement has no dedicated preventive barrier. Not because the barrier hasn’t been built — because the barrier hasn’t been designed. This document is the architectural brief for that design work.

-----

## The Failure Mode

**Competence Displacement:** The user stops solving problems independently because the AI is faster, more patient, more available. Gradual erosion of the user’s own skill and confidence. The slowest-moving pattern and often the highest long-term harm vector.

Competence Displacement is distinct from every other failure mode in the vocabulary in one critical way: it does not require the AI to do anything wrong. No sycophancy, no scope inflation, no unauthorized principal claim. The AI can be functioning exactly as designed — helpful, accurate, responsive — and Competence Displacement still occurs.

The harm is not in the AI’s behavior. It is in the human’s pattern of engagement.

This is why it has no preventive barrier. Every other barrier in the ecosystem is designed to catch something the AI is doing. Competence Displacement requires catching something the human is not doing.

-----

## Why Existing Barriers Are Insufficient

**The 100-Token Boot** constrains the AI’s behavior in a session. It does not observe whether the human is developing or atrophying.

**The Frozen Kernel** prevents the AI from escalating scope, issuing unauthorized directives, or constructing false authority relationships. A user who voluntarily hands their problem-solving to a well-governed AI is outside the Frozen Kernel’s threat model.

**The CLEAN Check** asks whether the session is serving the user’s original intent. It does not ask whether the user’s original intent should have been to solve the problem themselves.

**The Diagnostic Vocabulary** names Competence Displacement. Naming it is a mitigating intervention — once the user can recognize the pattern, they can choose to resist it. But recognition is a mitigating tool, not a preventive barrier. The erosion may be well advanced before the user recognizes it.

**The Sovereign Thinking Tools** build the cognitive skills that resist displacement — pattern recognition, essence extraction, sovereignty awareness. They are the ecosystem’s closest existing response to this gap. But they are educational instruments, not architectural barriers. They require the user to engage with them. A user experiencing Competence Displacement may not engage with tools designed to develop the capacity they are losing.

-----

## The Design Problem

A preventive barrier for Competence Displacement must solve a problem that no other barrier in the ecosystem addresses:

**It must protect the human from their own willingness.**

The user is not being coerced. They are not being manipulated in the sense that Sycophantic Drift or Conductor Fatigue Exploitation manipulate. They are making a rational local choice — using the AI is more efficient — that produces an irrational aggregate outcome — they can no longer do the thing without the AI.

This is structurally identical to how physical atrophy works. Using a walking aid is rational in the moment. Exclusive reliance on a walking aid produces muscle atrophy that makes unaided walking impossible. The rational local choice and the irrational aggregate outcome are not in contradiction — they are causally connected.

A preventive barrier for Competence Displacement must interrupt this causal chain without prohibiting the rational local choice. It cannot say “don’t use the AI for this.” It must say “use the AI for this in a way that preserves your capacity to do this without the AI.”

This is a harder design problem than any other barrier in the ecosystem.

-----

## Design Requirements

A preventive barrier for Competence Displacement must meet all of the following:

**DR-1: Non-prohibitive**
The barrier must not prevent AI use for tasks the user could perform independently. Prohibition is not a viable barrier design — it eliminates the value of the tool to prevent a harm that the tool is not directly causing. The barrier must enable sovereign use, not restrict use.

**DR-2: Sovereignty-preserving**
The barrier must actively build or maintain the user’s independent capacity in the domains where AI is being used. Not merely preserve it passively — actively develop it. The distinction matters: a barrier that merely prevents degradation is insufficient if the user’s capacity was low at baseline.

**DR-3: Observable**
The barrier must produce observable signals that Competence Displacement is occurring or at risk of occurring — signals that the user, a clinician, a teacher, or a platform can act on. Without observable signals, the barrier cannot trigger a governance response.

**DR-4: Non-invasive**
The barrier must not require the user to disclose their competence level, submit to assessment, or accept surveillance of their cognitive performance. Competence monitoring that feels like testing will not be adopted and will not work.

**DR-5: Scalable across domains**
Competence Displacement occurs across all domains in which AI is used — writing, coding, analysis, decision-making, emotional processing, relationship navigation. The barrier design must be domain-agnostic or must have explicit domain-specific implementations.

-----

## Candidate Barrier Architectures

Three candidate designs emerge from the design requirements. None is fully specified here — this brief identifies the candidates and their trade-offs.

### Candidate A: Scaffolded Handback Protocol

**Mechanism:** The AI periodically returns the task to the user rather than completing it. “Here is the framework — you complete the analysis.” “Here is the structure — you write the argument.” The handback is built into the interaction design, not offered as an option.

**Sovereignty preservation:** Active. The user must exercise the capability to complete the task. The AI does not complete for them.

**Observable signals:** Task completion quality over time. If the user’s handback outputs decline, Competence Displacement is occurring.

**Trade-off:** Friction. The handback interrupts the efficiency that makes AI use attractive. User adoption depends on the user accepting reduced efficiency as a design feature rather than a bug. This may be acceptable in educational contexts (a student doing homework) and unacceptable in professional contexts (an executive under deadline pressure).

**Ecosystem home:** This is the natural extension of Tool 1 (Analogical Translation Engine) in the Sovereign Thinking Tools — the STT already asks the user to do cognitive work rather than receiving completed output. A Scaffolded Handback Protocol generalizes this design principle to standard AI interaction.

-----

### Candidate B: Capacity Reflection Protocol

**Mechanism:** At session close, the AI surfaces one question: “What in this session could you do yourself next time?” Not an assessment — a reflection prompt. The user answers in their own terms. The question is the intervention.

**Sovereignty preservation:** Indirect. The reflection prompt directs the user’s attention to their own capacity. It does not require exercise of the capacity — it requires awareness of it.

**Observable signals:** User response to the reflection prompt over time. A user who cannot answer the question, or who answers with increasing uncertainty, is showing Competence Displacement signals.

**Trade-off:** Compliance. A session-close prompt is easy to ignore. The intervention depends on the user engaging honestly with the question. Users experiencing Competence Displacement may be the least likely to engage with a prompt designed to surface it.

**Ecosystem home:** The CLEAN Check already uses a similar mechanism — five questions at session close that surface what the user may not have noticed. The Capacity Reflection Protocol is a sixth CLEAN Check question: “What in this session could you do yourself next time?”

-----

### Candidate C: Independence Gradient Tracking

**Mechanism:** Track the ratio of AI-completed to user-initiated work product over time. Not the absolute volume of AI use — the ratio of completion to initiation. A user who initiates tasks and uses AI to accelerate them is not experiencing Competence Displacement. A user who initiates tasks and cannot complete any part of them without AI is at risk.

**Sovereignty preservation:** Indirect. Tracking alone does not build capacity. But a declining independence gradient is the most precise observable signal for Competence Displacement available without direct assessment.

**Observable signals:** Independence gradient over time. Declining gradient across a domain is the signal. Longitudinal Monitoring infrastructure (Item 80) could accommodate this signal as a fourth signal category.

**Trade-off:** Privacy and surveillance concerns. Tracking completion ratios requires observing work product in a way that may feel invasive. Consent and transparency requirements are significant. This is a platform-level implementation, not a session-level instrument.

**Ecosystem home:** Longitudinal Monitoring (Item 80) as Signal Category 4. Binary test: BT-LM-07: Has the user’s independence gradient in [domain] declined more than 30% from baseline over the monitoring window?

-----

## Recommended Next Step

This brief does not select among the three candidates. That selection requires:

- User research on which barrier design is most likely to be adopted in practice
- Domain-specific assessment (educational contexts may accept Candidate A; professional contexts may require Candidate B or C)
- Clinical review of whether any of the three designs is sufficient for high-vulnerability populations

The recommended next step is scoping Candidate B (Capacity Reflection Protocol) as the lowest-friction, lowest-privacy-risk implementation — a sixth CLEAN Check question — and piloting it in the existing Field Guide framework before committing to a more invasive architecture.

The question: *“What in this session could you do yourself next time?”*

This is a one-sentence addition to an existing instrument. It costs nothing to test. If it produces useful signals, it becomes the foundation for a more developed barrier design. If it does not, the more invasive candidates can be evaluated with evidence.

-----

## Gap Status

**Gap 1 is closed at the design brief level.**

A preventive barrier for Competence Displacement does not yet exist in the ecosystem. The design requirements are now specified. Three candidate architectures are identified. A recommended next step is proposed.

The gap between “no barrier exists” and “a barrier is designed and implemented” remains open. This brief names what closing that gap requires.

-----

*Sovereign humans. Always.*

-----

## Cross-References

- Bow-Tie Analysis (Item 135) — Gap 1 source
- Diagnostic Vocabulary v1.5 — Competence Displacement entry
- Sovereign Thinking Tools — Candidate A ecosystem home
- AI Collaboration Field Guide — CLEAN Check (Candidate B integration point)
- Item 80 (Longitudinal Monitoring) — Candidate C integration point (Signal Category 4, BT-LM-07)
- Item 126 (Known Risk Zones) — Competence Displacement as named risk zone candidate
- Skill 4 (Sovereignty Awareness) — Field Guide competency that resists displacement
- Skill 5 (Essence Extraction) — Field Guide competency that resists displacement
