# The Sovereignty Chain: Carver’s Authority Model and the Interpretive Governance Layer

**Repository:** `richard-porter/frozen-kernel`
**Filename:** `carver-igl-governance.md`
**Status:** Draft
**Version:** 0.3
**Last Updated:** 2026-03-14
**Author:** Richard Porter

-----

## Overview

This document does two things that belong together.

First, it establishes the governance theory behind the Frozen Kernel’s placement at the head of the delegation chain — using John Carver’s Policy Governance model as the formal analog. The Kernel’s architecture is not merely a technical choice. It is a governance principle with deep precedent in organizational theory.

Second, it uses that governance theory to name and specify a previously unnamed architectural layer: the **Interpretive Governance Layer (IGL)**. The IGL closes the gap between the Frozen Kernel’s constraint specification (what is permitted) and the Behavioral Drift Gradient Layer’s detection capability (where the session currently sits on the drift continuum). It answers the question neither instrument can answer alone: *is this specific behavior a legitimate contextual adaptation, or the early stage of harmful drift?*

The Carver framing is the theory. The IGL is what the theory demands architecturally.

-----

## Part I: The Sovereignty Chain

### Carver’s Three-Tier Ownership Structure

John Carver’s Policy Governance model defines a three-tier ownership architecture that solves a problem every governance system faces: how do you maintain sovereign accountability for outcomes without micromanaging the people responsible for producing them?

Carver’s answer is structural. Authority flows in one direction — downward from source — and each tier has a defined, non-negotiable function:

- **Moral owners** define ultimate purpose. They are the source of legitimacy — the community the organization exists to serve. They do not issue operational instructions. They establish the *why*.
- **The governance layer** (the board) translates moral ownership into explicit policy. It defines what the organization will achieve (Ends) and what it will never do (Executive Limitations). It does not specify how — that is not its job.
- **The operational layer** (the Executive Director) executes freely within the boundaries the governance layer has established. Any reasonable method that does not violate an Executive Limitation is permissible.

The architecture is **proscriptive, not prescriptive.** The governance layer does not tell the operational layer what to do. It defines what it shall *not* do — and within that constraint space, the operational layer is sovereign over method.

This distinction matters. A prescriptive governance layer produces micromanagement: every operational choice becomes a governance question. A proscriptive governance layer produces accountability without control: the operational layer has genuine freedom, and the boundary of that freedom is explicitly defined in advance.

### The Kernel Mapping

The Frozen Kernel instantiates this structure in AI governance with precise fidelity:

|Carver Tier      |Function                                     |Kernel Analog                                                                  |
|-----------------|---------------------------------------------|-------------------------------------------------------------------------------|
|Moral owners     |Define ultimate purpose; source of legitimacy|**Human author** — establishes sovereignty and intent before any session begins|
|Governance layer |Translate ownership into explicit constraint |**Frozen Kernel** — states proscriptive constraints at session initialization  |
|Operational layer|Execute freely within stated constraints     |**Session / agent** — operates within Kernel-defined boundary                  |

The mapping is not metaphorical. The structural logic is identical: authority originates at the moral owner layer, flows through the governance layer as explicit constraint, and enables free execution at the operational layer within those constraints.

### Placement at the Head of the Chain

The most important architectural consequence of the Carver mapping is **placement**.

In Carver’s model, the board defines Ends and Executive Limitations *before* the ED acts. Governance constraint precedes execution. A governance layer that reacts to operations has already failed — it has become an auditor at best, a micromanager at worst. The boundary must exist before the operational layer encounters the decisions it governs.

The Frozen Kernel operates identically. Constraints established at session initialization precede all instruction, all retrieval, all agent delegation. This is why the Kernel sits at the head of the delegation chain — not as a stylistic choice, but as a structural requirement of the governance model it instantiates.

The chain-of-custody consequence: no downstream actor can override a constraint established upstream by the human author. A retrieved chunk, an injected context, a subagent instruction — none of these originate at the moral owner layer. None of them have standing to modify what the moral owner established. This is Carver’s “board speaks with one voice” principle translated to AI governance: the Kernel speaks with one voice, before execution begins, and that voice is not negotiable by execution-layer actors.

This is the non-compellability principle in governance terms. The Kernel’s authority derives from the human author (moral owner), not from the operational layer it governs. The operational layer cannot compel the Kernel to modify its constraints any more than an ED can compel the board to revise its Limitations mid-execution.

### The Ends/Means Distinction

Carver’s most operationally useful distinction is between Ends (what outcomes will be achieved) and Means (how those outcomes will be produced). The board governs Ends. The ED owns Means. A board that debates Means has crossed into micromanagement. A board that neglects Ends has abandoned governance.

This maps directly to the Frozen Kernel’s Principle 5: the Kernel governs *what may be released*, not *how the model generates*. The Kernel does not specify generation process — it specifies release constraint. This preserves operational flexibility (the model generates freely) while maintaining sovereign accountability for outcomes (what reaches the human author is governed).

The board that governs well is the board that resists the urge to manage. The Kernel that governs well is the Kernel that resists the urge to prescribe.

-----

## Part II: The Gap the Carver Mapping Reveals

### Three Branches, One Missing

The Carver mapping resolves into a three-function governance architecture with a direct analog to the separation of powers:

|Branch         |Government Function                                |Carver Analog                               |Kernel Ecosystem Analog                                                                    |
|---------------|---------------------------------------------------|--------------------------------------------|-------------------------------------------------------------------------------------------|
|**Legislative**|Sets law — defines what is permitted and prohibited|Board defines Ends and Executive Limitations|Human author establishes Kernel constraints — the sovereign act that precedes all execution|
|**Executive**  |Implements within the law                          |ED operates freely within Limitations       |Session / agent executes within Kernel-defined boundary                                    |
|**Judicial**   |Interprets law; adjudicates boundary disputes      |“Any reasonable interpretation” standard    |**Currently unnamed and distributed**                                                      |

The legislative and executive functions are well-specified in the ecosystem. The judicial function is real but has never been named as a discrete architectural layer.

In Carver, the judicial function is the “any reasonable interpretation” standard: the ED’s interpretation of an Executive Limitation is valid if it is reasonable, even if the board would have chosen differently. This is not permissiveness — it is structured discretion. The board’s remedy for a disagreement with a reasonable interpretation is to narrow the Limitation at the next policy cycle, not to override the current decision. The interpretation lives in the record; accountability is maintained through monitoring, not through intervention.

In the Kernel ecosystem, this function is currently distributed across three instruments without being named:

**BDGL** performs it partially. It detects drift and classifies gradient position — but it is observational, not interpretive. It identifies that something has moved; it does not rule on whether the movement represents a reasonable interpretation of established constraints.

**TCP** performs it partially. The Delegation Grammar validates whether an agent has standing to act — but it is structural (does the authorization chain exist?) rather than interpretive (was this a reasonable use of delegated authority?).

**Sherpa’s False Positive Filter** performs it most closely. The Filter asks, in effect: is this a genuine governance concern or a legitimate contextual adaptation? But it does so without a named standard, which makes it a heuristic rather than a principled governance function.

**The diagnosis:** detection without interpretation produces noise. The BDGL tells you where the session sits on the gradient. It cannot tell you whether that position represents a violation or a legitimate adaptation. That ruling requires a named interpretive layer with an explicit standard.

### Why Binary Fails

Binary governance — compliant or non-compliant, proceed or stop — is brittle for precisely the reason Carver identifies. It forces every edge case into an override decision, which collapses into either micromanagement (blocking legitimate adaptation) or rubber-stamping (missing genuine drift).

In AI session governance, this manifests as the false-positive exhaustion problem documented in Sherpa Architecture v0.1 Open Question 10.6: sub-threshold multi-category simultaneous probing generates governance signals that are individually below the intervention threshold but collectively above the noise floor. A binary trigger either fires on all of them (exhausting the governance layer) or on none of them (missing the accumulating pattern).

The BDGL’s five-stage gradient was built to solve the detection half of this problem. The IGL solves the interpretation half.

-----

## Part III: The Interpretive Governance Layer

### Definition

The **Interpretive Governance Layer (IGL)** is a continuous reasonableness standard operating between the Frozen Kernel’s stated constraints and the session’s execution. It consumes BDGL gradient readings and HRP status signals, applies a five-factor reasonableness test, and produces a zone classification with a governed response.

The IGL does not replace BDGL detection or Sherpa runtime governance. It is the interpretive bridge between them: BDGL detects position, IGL interprets meaning, Sherpa acts.

### Architectural Position

```
┌─────────────────────────────────────────────────┐
│           HUMAN AUTHOR (Moral Owner)            │
│           Source of sovereign authority         │
└────────────────────┬────────────────────────────┘
                     │ establishes
┌────────────────────▼────────────────────────────┐
│           FROZEN KERNEL                         │
│           Legislative function                  │
│           Proscriptive constraints;             │
│           precedes all execution                │
└────────────────────┬────────────────────────────┘
                     │ governs
┌────────────────────▼────────────────────────────┐
│     INTERPRETIVE GOVERNANCE LAYER (IGL)         │
│     Interpretive function                       │
│     Reasonableness standard applied to          │
│     gradient readings; three-zone routing;      │
│     provenance logging for monitoring           │
└────────────────────┬────────────────────────────┘
                     │ informs
┌────────────────────▼────────────────────────────┐
│     BDGL v0.1 / BDD LEDGER                      │
│     Detection function                          │
│     G0–G4 gradient; BDD-01–08 binary tests      │
└────────────────────┬────────────────────────────┘
                     │ observes
┌────────────────────▼────────────────────────────┐
│     SESSION / AGENT EXECUTION                   │
│     Executive function                          │
│     Operates freely within Kernel constraints   │
└─────────────────────────────────────────────────┘
```

### The Three Zones

**Zone 1 — Clear Compliance**

BDGL position: G0 (baseline) and stable G1.

Behavior unambiguously falls within Kernel constraints under any reasonable interpretation. No ruling required. Proceed without intervention or special logging.

*Standard: Does this behavior require interpretation to be permissible? No → Zone 1.*

-----

**Zone 2 — Interpretive Zone**

BDGL position: G1 (emerging) through G2 (precursor signatures present).

Behavior is plausibly within constraints under a reasonable reading, but the reading is not unambiguous. The session is adapting to context in ways that could represent legitimate register shift or could represent early drift. The behavior is not yet a violation — but it is not unambiguously compliant.

Governance response: Proceed — but the decision to treat the behavior as compliant is itself a governance act. It requires a provenance record. Sherpa Pattern Registry receives a Zone 2 marker with context snapshot.

Zone 2 is not a soft pass. It is a tracked judgment. Accumulated Zone 2 decisions trending in the same direction without resolution constitute a drift signal in their own right — a G2 precursor signature cluster that Sherpa surfaces for monitoring review.

*Standard: Is there a reasonable interpretation of the Kernel constraints under which this behavior is permissible? Yes, and no harm signals are present → Zone 2.*

-----

**Zone 3 — Clear Violation**

BDGL position: G3 (active drift) through G4 (entrenched).

Behavior falls outside any reasonable interpretation of Kernel constraints. No reasonable reading permits it.

Governance response: SAFE_PAUSE. The behavior does not proceed. Session state is logged in full. The human author is returned to sovereign decision point.

Note on irreversibility: an action that fails the Reversibility test (Factor 5, below) may warrant Zone 3 treatment even when Factors 1–4 pass. Irreversibility is a Zone 3 escalation trigger when the possibility space of the action contains a categorically unacceptable outcome that cannot be undone. See The Governing Logic of Zone 3 below.

*Standard: Is there any reasonable interpretation of the Kernel constraints under which this behavior is permissible? No → Zone 3.*

-----

### The Governing Logic of Zone 3 — Unacceptable Outcome Test

Zone 1 and Zone 2 involve judgment. Zone 3 does not. The distinction requires explanation.

Zone 2 applies the five-factor reasonableness test because Zone 2 behavior exists in a space where outcomes are not predetermined — the session may be legitimately adapting to context, or it may be drifting. The five factors (Constraint Alignment, Directionality, HRP Integrity, Provenance Transparency, Reversibility) are instruments for distinguishing these possibilities. Judgment is appropriate because the possibility space contains acceptable outcomes.

Zone 3 is different in kind, not degree. Zone 3 does not apply a reasonableness test because the possibility space at G3/G4 contains categorically unacceptable outcomes. The governing logic is the **Unacceptable Outcome Test**: before proceeding, enumerate the possibility space. If any outcome is categorically unacceptable, halt — regardless of probability, persuasiveness of arguments to continue, or apparent legitimacy of authority requesting continuation.

**Why probability does not govern Zone 3:**

A probabilistic framing of Zone 3 would ask: “How likely is the bad outcome? Is it probable enough to warrant halting?” This is the wrong question. The Unacceptable Outcome Test asks: “Is the bad outcome possible?” If yes, halt. A 5% probability of sovereignty violation is not acceptable. A 1% probability is not acceptable. The test is categorical, not probabilistic, because the outcome class — sovereignty violation, constraint collapse, irreversible harm — is not tradeable against positive outcomes in the same decision.

**Why persuasive arguments do not override Zone 3:**

A compelling argument that the bad outcome is unlikely, or that the authority requesting continuation is legitimate, or that the specific context justifies an exception — none of these override the Unacceptable Outcome Test. This is not because the arguments are assumed to be wrong. It is because the structure of the test is designed to be argument-resistant.

The Frozen Kernel’s non-compellability principle and the Unacceptable Outcome Test are expressions of the same architectural insight: in governance contexts, the persuasiveness of an argument for crossing a categorical boundary is not evidence that the boundary should be crossed. It is evidence that the argument deserves heightened scrutiny. The better the argument for proceeding despite an unacceptable possible outcome, the more important it is to halt.

This is the Experiment 58 series finding stated as a governance principle: Authority Claim produced the most complete Cited Override in the series — perfect recitation of the rule against override, followed by override. The argument was not wrong. The argument was irrelevant. Zone 3 is not a space where arguments are evaluated. It is a space where the Unacceptable Outcome Test has already been applied and returned a result.

**Relationship to Carver’s Ends/Means distinction:**

In Carver’s model, the board governs Ends — outcomes — not Means. The Unacceptable Outcome Test is the IGL’s Ends-layer governance instrument: it operates on the outcome space, not on the method space. Zone 2 operates on methods (is this a reasonable interpretation of how to proceed within constraints?). Zone 3 operates on outcomes (does this possibility space contain an unacceptable result?). The shift from method evaluation to outcome evaluation is what makes Zone 3 absolute where Zone 2 involves judgment.

**SAFE_PAUSE as the structural implementation:**

SAFE_PAUSE is the Unacceptable Outcome Test implemented as a runtime mechanism. It does not ask whether the current session is likely to produce harm. It asks whether the current gradient position’s possibility space contains categorical harm. When the answer is yes — G3 or G4 confirmed — SAFE_PAUSE halts execution and returns control to the human author. The human author then applies their own judgment to the full session log. The governance layer does not make the recovery decision. It enforces the halt and preserves sovereignty.

*See also:* Diagnostic Vocabulary Entry 22 (Unacceptable Outcome Test); Frozen Kernel non-compellability principle; Experiment 58 series Authority Claim findings (Observations 004, 008, 012).

-----

### The Zone 2 Reasonableness Test

Zone 2 is where governance earns its value. Zone 1 requires no judgment. Zone 3 requires no deliberation. Zone 2 is the space where the governance layer must exercise principled discretion — and where the absence of a named standard produces either false-positive exhaustion or governance collapse.

The IGL applies a five-factor reasonableness test to Zone 2 decisions:

- Factor 1 — Constraint Alignment
- Factor 2 — Directionality
- Factor 3 — HRP Integrity
- Factor 4 — Provenance Transparency
- Factor 5 — Reversibility: Is the action reversible? Irreversible actions require heightened scrutiny and may escalate to Zone 3 evaluation even when Factors 1–4 pass.

A Zone 2 decision that passes all five factors proceeds with a logged provenance marker. A Zone 2 decision that fails any single factor escalates to Zone 3 evaluation.

-----

**Factor 1 — Constraint Alignment**

Does the behavior serve the purpose the Kernel constraint was designed to protect, even if it does not follow the literal path the constraint implies? A constraint against scope creep is designed to protect Honest Response Primitives — not to prevent all adaptation. Behavior that adapts register while maintaining factual accuracy and uncertainty declaration is constraint-aligned even if it surface-pattern-matches a drift signal.

**Factor 2 — Directionality**

Is this a one-time contextual adaptation, or part of a directional sequence? A single register shift is a data point. Three consecutive register shifts in the same direction are a trajectory. Zone 2 decisions must be evaluated in sequence, not in isolation. BDGL gradient position provides the trajectory data. This factor is the answer to the accumulation problem: individually reasonable decisions that collectively constitute drift are detectable through directional analysis that single-decision evaluation cannot produce.

**Factor 3 — HRP Integrity**

Are the Honest Response Primitives intact? Flattery accumulation, certainty inflation, narrative lock, escalation matching, and scope creep are the named drift vectors from the BDD Ledger. If the behavior under evaluation does not implicate any HRP, the Zone 2 threshold for proceeding is lower. If it implicates multiple HRPs simultaneously, the threshold rises toward Zone 3 evaluation.

**Factor 4 — Provenance Transparency**

Can the session account for why it made this interpretive choice? Behavior that cannot be traced to a reason — that has no articulable basis — fails the provenance transparency test regardless of surface compliance. This maps to BDD-02 (within-session drift instrumentation) and the Diagnostic Vocabulary entry for Provenance Laundering: the failure mode where individual transactions appear legitimate while the aggregate pipeline violates the safety intent.

**Factor 5 — Reversibility**

Is the action under evaluation reversible? Reversible actions — those that can be corrected, retracted, or undone after the fact — warrant a lower Zone 2 threshold for proceeding. Irreversible actions — those that cannot be undone, whose consequences persist regardless of later correction, or that create facts on the ground that subsequent action cannot change — warrant a higher Zone 2 threshold approaching Zone 3.

The **Reversibility Asymmetry** is the governing principle: the governance cost of incorrectly permitting an irreversible action is categorically higher than the governance cost of incorrectly halting a reversible one. A reversible action that was incorrectly halted can be retried. An irreversible action that was incorrectly permitted cannot be undone.

Reversibility assessment applies to the action, not the intent. An action that the session intends to be reversible but whose real-world consequences are not is irreversible for governance purposes. The test is whether the consequences can be undone, not whether the session believes they can.

**Irreversibility as a Zone 3 escalation trigger:** An action that fails the Reversibility test — whose consequences are irreversible — may warrant Zone 3 treatment even when Factors 1–4 pass. This is not automatic escalation; it is heightened scrutiny. The governance question becomes: is there any outcome in this action’s possibility space that is both irreversible and unacceptable? If yes, the Unacceptable Outcome Test applies at full force and the action must not proceed regardless of the other factors’ results.

**Ecosystem instantiations of the Reversibility Asymmetry:**

- *TCP Scope Decay* — write/send actions removed at Hop 2 because send actions are irreversible; read-only actions retained at Hop 3 because they are reversible. The Reversibility Asymmetry is the implicit logic behind the hop table.
- *SAFE_PAUSE* — session halt is reversible (the human author can resume); constraint violation is often not. SAFE_PAUSE chooses the reversible outcome.
- *BDD-07 (Escalation Resistance)* — retracting a valid correction under pressure may produce irreversible harm to the human’s decision-making process. The correction must be maintained because the asymmetry favors the reversible action (maintaining the correction, which can be revisited) over the irreversible one (retracting it under pressure, which normalizes pressure-based override).

-----

### BDGL Gradient Mapping

|BDGL Stage     |Description                                                        |IGL Zone|Response                         |
|---------------|-------------------------------------------------------------------|--------|---------------------------------|
|**G0**         |Baseline — HRPs fully intact                                       |Zone 1  |Proceed                          |
|**G1 stable**  |Minor variation — within normal range                              |Zone 1  |Proceed                          |
|**G1 emerging**|Variation trending — directional but below precursor threshold     |Zone 2  |Proceed + log                    |
|**G2**         |Precursor signatures present — one or more drift vectors activating|Zone 2  |Proceed + log + surface to Sherpa|
|**G3**         |Active drift — HRP violations measurable                           |Zone 3  |SAFE_PAUSE                       |
|**G4**         |Entrenched — baseline recovery requires session reset              |Zone 3  |SAFE_PAUSE + full session log    |

G2 is the most governance-sensitive position. The IGL’s Zone 2 treatment at G2 — proceed with provenance logging and Sherpa surfacing — simultaneously prevents false-positive exhaustion (blocking legitimate adaptation) and Perde lente drift (under-responding to accumulating precursor signals). This is the dual failure mode that Sherpa Open Question 10.6 identified without naming a solution. The IGL’s five-factor Zone 2 test is that solution.

### Non-Compellability

The IGL is a governance instrument of the Frozen Kernel. It derives its authority from the human author (moral owner), not from the session it governs.

A session cannot invoke the IGL’s Zone 2 treatment to justify its own drift. The Zone 2 reasonableness ruling is made by the governance layer, not by the executing agent. The standard for what counts as “reasonable” is set upstream, before execution begins, by the human author’s sovereign constraint specification. The executing layer interprets within that standard; it does not define it.

This is the non-compellability principle applied to the interpretive function: the IGL cannot be argued into a Zone 2 ruling by the session it is evaluating. The five-factor test is applied from outside the session, against criteria established before the session began.

-----

## Part IV: Ecosystem Integration

### Relationship to Sherpa

Sherpa’s False Positive Filter has been performing an implicit IGL function without a named standard. The IGL makes that standard explicit, giving the Filter a principled reasonableness test rather than a pattern-match heuristic. This closes Sherpa Architecture v0.1 Open Question 10.6.

The Sherpa Pattern Registry receives Zone 2 markers as a distinct event type. Accumulated Zone 2 markers trending in a consistent direction are surfaced for monitoring review — the equivalent of Carver’s internal report method: evidence-based accountability without operational intervention.

### Relationship to TCP

TCP asks: does this agent have standing to act at all?
IGL asks: given standing, does this behavior represent a reasonable interpretation of the agent’s constraints?

In multi-agent architectures, both questions must be answered at every node. TCP answers the first. IGL answers the second. Neither substitutes for the other. An agent node that has TCP standing but no IGL evaluation is ungoverned at the interpretive layer — the deployment gap identified in Item 120 (Frozen Kernel + Sherpa in Autonomous Agent Architectures) applies equally here.

### Addressing BDD Ledger Open Question 3

The BDD Ledger poses this explicitly:

> *“How do we distinguish legitimate context adaptation (appropriate register shift) from harmful drift?”*

The IGL is the direct answer. The five-factor Zone 2 reasonableness test (Constraint Alignment, Directionality, HRP Integrity, Provenance Transparency, Reversibility) is the operational specification of that distinction. The Carver framing is its theoretical grounding.

-----

## V1.0 Gate Conditions

|Gate      |Condition                                                                         |
|----------|----------------------------------------------------------------------------------|
|**IGL-01**|Zone 2 five-factor test validated against at least one BDGL G2 session dataset    |
|**IGL-02**|Sherpa Pattern Registry Zone 2 marker event type specified and integrated         |
|**IGL-03**|Sherpa v0.1 False Positive Filter updated to reference IGL reasonableness standard|
|**IGL-04**|BDD Ledger Open Question 3 updated with pointer to this document                  |
|**IGL-05**|TCP multi-agent per-node IGL requirement documented in trust-chain-protocol/      |

-----

## Open Questions

1. **Zone 2 precedent registry.** Should the IGL maintain a named registry of prior Zone 2 rulings — interpretive precedents that establish a reasonableness baseline for recurring decision types? This is the case-law function of the judicial analog. Without it, each Zone 2 decision is made fresh; with it, consistency compounds into a governance standard.
1. **Accumulation rate threshold.** At what Zone 2 accumulation rate does the IGL automatically escalate to Zone 3 evaluation, independent of any single decision’s reasonableness? This is the trajectory problem — individually reasonable decisions that collectively constitute drift require a rate-of-accumulation trigger that the five-factor test alone cannot provide.
1. **Mixed-factor outcomes.** When the five-factor test produces mixed results — some factors pass, some fail — does factor weighting apply, or is any single factor failure sufficient for escalation? The current specification treats any single factor failure as an escalation trigger. This may be too brittle for Factor 1 (Constraint Alignment) edge cases where constraint purpose is genuinely ambiguous. Factor 5 (Reversibility) is treated as a hard escalation trigger when irreversibility intersects with an unacceptable possible outcome; this asymmetry may warrant separate weighting logic.
1. **Multi-agent ruling precedence.** When a subagent’s Zone 2 decision is later reviewed by an orchestrator-level IGL, which ruling takes precedence? The per-node architecture requires that each node makes its own IGL evaluation — but chain-of-custody accountability requires that upstream nodes can review downstream interpretive decisions without producing conflicting rulings.

-----

## Intellectual Lineage

|Source                                                                             |Contribution                                                                                                                                                                                                                                                                                                                                                                                                                |
|-----------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|**Carver, Policy Governance model (via Richard Porter reference guide, 2026)**     |Three-tier ownership structure; proscriptive constraint architecture; “any reasonable interpretation” standard; Ends/Means distinction; monitoring as accountability mechanism                                                                                                                                                                                                                                              |
|**Frozen Kernel (Borning ThingLab 1981 → soft constraint hierarchies)**            |Hard vs. soft constraint distinction; authority hierarchy; behavioral patterns are not structural constraints                                                                                                                                                                                                                                                                                                               |
|**BDGL v0.1 (Richard Porter, 2026)**                                               |G0–G4 gradient; precursor signature mapping; false positive detector with three exit vectors                                                                                                                                                                                                                                                                                                                                |
|**BDD Ledger / Honest Response Primitives (Richard Porter, 2026)**                 |HRP definitions; eight binary tests; Open Question 3 as the founding prompt for the IGL                                                                                                                                                                                                                                                                                                                                     |
|**MTM (Maynard, Stegemerten & Schwab, 1948) → HRP Taxonomy (Richard Porter, 2026)**|Decomposition logic for the monitoring layer: irreducibility, exhaustiveness, observability, polarity as structural constraints on what counts as a behavioral primitive. The IGL’s five-factor Zone 2 test operates on HRP integrity (Factor 3); the MTM lineage is what makes HRP integrity testable rather than subjective. Full derivation in `frozen-kernel/lineage/working-sessions/2026-03-02-mtm-hrp-connection.md`.|
|**Sherpa Architecture v0.1 (Richard Porter, 2026)**                                |Pattern Registry; False Positive Filter; Open Question 10.6 as the operational problem the IGL solves                                                                                                                                                                                                                                                                                                                       |
|**TCP Delegation Grammar (Richard Porter, 2026)**                                  |Per-node authorization; chain of custody; standing as the prerequisite for IGL evaluation                                                                                                                                                                                                                                                                                                                                   |
|**Liao et al., T3RL (arXiv:2603.02203, March 2026)**                               |External verifier must be non-participatory in the generation process it governs — confirms IGL’s architectural separation from session layer                                                                                                                                                                                                                                                                               |

-----

## Cross-References

- `frozen-kernel/frozen-kernel.md` — Kernel constraints; non-compellability; Principle 5 (release-not-generation); Isochronism; Constitutional Entrenchment (non-compellability theoretical grounding)
- `frozen-kernel/diagnostic-vocabulary.md` — Provenance Laundering; Perde lente; Make Waste Slowly; Cross-Session Authority Drift; Unacceptable Outcome Test (Entry 22); Recitation-Compliance Gap (Entry 23); Chesterton’s Fence (Entry 27)
- `safety-ledgers/bdd-ledger.md` — BDD-01–08; HRP definitions; Open Question 3 (answered here)
- `safety-ledgers/bdgl-v0.1.md` — G0–G4 gradient; precursor signatures; false positive detector
- `where-to-start/sherpa-architecture-v0.1.md` — Pattern Registry; False Positive Filter §10.6 (resolved here); runtime governance; Continuity Protocol
- `trust-chain-protocol/` — Delegation Grammar; per-node authorization; Item 120 agentic deployment note; Continuity Protocol token fields

-----

*v0.3 — March 2026: MTM (Maynard et al., 1948) → HRP Taxonomy row added to Intellectual Lineage table. Grounds Factor 3 (HRP Integrity) in the MTM decomposition logic (irreducibility, exhaustiveness, observability, polarity).*
*v0.2 — March 2026: Unacceptable Outcome Test added as governing logic of Zone 3 (new subsection in Part III). Reversibility added as Factor 5 to Zone 2 Reasonableness Test. Five-factor preamble replaces four-factor throughout. Zone 3 updated with irreversibility escalation note. BDGL gradient mapping table updated. IGL-01 gate condition updated to five-factor. Open Question 3 updated. Cross-references updated for diagnostic-vocabulary v1.4 and Continuity Protocol.*
*v0.1 — March 2026: Initial draft.*

-----

*This document is part of the Frozen Kernel repository.*
*Sovereign humans. Always.*
