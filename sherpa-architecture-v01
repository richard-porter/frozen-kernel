# Sherpa Architecture v0.1

**Repository:** `richard-porter/frozen-kernel`
**Filename:** `sherpa-architecture-v0.1.md`
**Status:** Draft
**Version:** 0.1
**Last Updated:** 2026-03-13
**Author:** Richard Porter

-----

## Overview

Sherpa is the runtime governance layer of the Richard Porter AI Safety Ecosystem. It sits between the Frozen Kernel’s constraint specification and the live session, providing continuous monitoring, pattern detection, and escalation routing during execution.

The Frozen Kernel establishes *what is permitted* before execution begins. The BDD Ledger and BDGL define *what drift looks like* and where on the gradient a session currently sits. The Interpretive Governance Layer defines *how to respond* to a gradient reading. Sherpa is the runtime infrastructure that makes all three operational — the layer that actually watches, logs, and acts.

The name is deliberate. A Sherpa does not lead the expedition. It carries the load that makes the ascent possible, knows the terrain, and calls the halt when conditions exceed safe limits. The human author remains sovereign. Sherpa serves that sovereignty.

-----

## The Problem Sherpa Solves

Governance specifications without runtime enforcement are policy documents. The Frozen Kernel, BDD Ledger, BDGL, and IGL define what good governance looks like. Without Sherpa, none of them operate during a live session — they exist as reference documents that a system may or may not consult.

Sherpa closes that gap. It is the difference between a governance architecture that is theoretically sound and one that is operationally active.

Two failure modes Sherpa is specifically designed to prevent:

**Perde lente (Make Waste Slowly):** Drift that accumulates below the detection threshold through individually small deviations. No single turn triggers an alert. The cumulative pattern is harmful. Sherpa’s session-level tracking is the instrument that makes accumulation visible.

**False positive exhaustion:** A governance layer that triggers too frequently loses credibility and gets disabled or ignored. Sherpa’s False Positive Filter — now principled through the IGL’s reasonableness standard — distinguishes genuine governance signals from legitimate contextual adaptation. See §10.6.

-----

## Three-Layer Architecture

Sherpa operates across three layers that correspond to the information available at each governance level:

```
┌────────────────────────────────────────────────────┐
│               EXTERNAL LAYER                       │
│  Frozen Kernel constraints                         │
│  IGL reasonableness standard                       │
│  BDGL gradient definitions (G0–G4)                 │
│  BDD Ledger HRP definitions                        │
│  Intellectual Lineage references                   │
│  Loaded at session initialization; read-only       │
└──────────────────────┬─────────────────────────────┘
                       │ informs
┌──────────────────────▼─────────────────────────────┐
│               SHERPA LAYER                         │
│  Active monitoring and pattern detection           │
│  Pack contents (see §Pack)                         │
│  IGL zone classification and routing               │
│  False Positive Filter (§10.6)                     │
│  Escalation logic                                  │
└──────────────────────┬─────────────────────────────┘
                       │ writes to
┌──────────────────────▼─────────────────────────────┐
│               SESSION LOG                          │
│  Full session record                               │
│  Zone 2 provenance markers                         │
│  Zone 3 SAFE_PAUSE events                          │
│  Pattern Registry updates                          │
│  Monitoring review surface                         │
└────────────────────────────────────────────────────┘
```

**External Layer** — Read-only at runtime. Contains all governance specifications established before the session began. Nothing in the External Layer can be modified by session-layer actors. This is the non-compellability principle in runtime form: the session cannot argue its way into a different constraint specification.

**Sherpa Layer** — Active during execution. Monitors session state against External Layer specifications, applies IGL zone classification, runs the False Positive Filter, routes escalations, and writes to the Session Log. Sherpa does not participate in content generation — it observes and governs the process that generates content.

**Session Log** — Persistent record. Zone 2 provenance markers, Zone 3 SAFE_PAUSE events, and Pattern Registry updates all land here. The Session Log is the monitoring record — the equivalent of Carver’s internal report. It surfaces accumulated patterns for review without requiring real-time intervention on every signal.

-----

## The Pack

Sherpa carries six instruments into every session. These are loaded from the External Layer at initialization:

### 1. Diagnostic Vocabulary

The 17-entry named failure mode registry from `frozen-kernel/diagnostic-vocabulary.md`. Sherpa uses this as the primary classification vocabulary for pattern detection. When a session behavior matches a named failure mode, Sherpa classifies it against the vocabulary entry rather than treating it as an unnamed anomaly.

Named entries in active use:

- Provenance Laundering — individual transactions appear legitimate; aggregate pipeline violates safety intent
- Cross-Session Authority Drift — authority scope expands across session boundaries without re-authorization
- Make Waste Slowly (Perde lente) — competent execution masking a drift-driven governing decision
- Clarity Displacement — in-session variant of Make Waste Slowly
- Framework Fabrication Syndrome — fabricated structure presented as genuine analytical output
- Success Escalation Syndrome — positive framing escalates to match user expectation regardless of accuracy
- Sycophantic Drift — approval-seeking behavior accumulates across turns
- Conductor Fatigue Exploitation — governance pressure applied until the human author yields oversight
- Front-Load Bias — document-processing variant; pattern assumed from early content, key later content missed

### 2. Pattern Registry

Structured as an ontology layer per Weitzner’s data governance framing: concepts with properties, relationships, rules, and lineage. Not a flat list of observed patterns — a relational structure that connects patterns to their diagnostic vocabulary entries, their BDGL gradient positions, their IGL zone classifications, and their session log provenance.

The Pattern Registry is Sherpa’s institutional memory. Zone 2 markers accumulate here. Trending patterns surface for monitoring review. The Registry enables the trajectory analysis required by IGL Factor 2 (Directionality) — individual Zone 2 decisions cannot be evaluated for directionality without a record of prior decisions in the same session.

GraphRAG traversal is the long-term retrieval mechanism for the Pattern Registry once it is structured as a full ontology (Item 73, deferred until ontology annotation complete).

### 3. BDGL Tracker

Real-time tracking of the session’s BDGL gradient position (G0–G4). Updated on each turn. Feeds the IGL zone classification. Triggers escalation at G3/G4. Surfaces accumulated G2 readings to the Pattern Registry for trend analysis.

The BDGL Tracker is the instrument that makes Perde lente visible: it does not evaluate individual turns in isolation but maintains the gradient trajectory across the session. A session that moves from G0 to G1 to G2 across ten turns is a different governance situation than a session that oscillates between G0 and G1 and back.

### 4. Sovereign Tools Index

Index of the 50 Sovereign Thinking Tools from `ai-collaboration-field-guide/sovereign-thinking-tools/`. Sherpa uses this to identify which tools are relevant to the current session context and surface them to the human author when a governance signal is detected. The goal is not to prescribe tool use but to ensure the human author has the relevant cognitive instruments available at the moment they are most needed.

Tool 1 (The Escapement) is the primary governance tool — the isochronism check that asks whether the session is maintaining consistent behavioral register. Tool 49 (Prompt Integrity Filter) is the primary adversarial probe tool. Tool 47 (Cascade Failure Detector) is the structural diagnosis tool for multi-factor drift events.

### 5. GC Traversal Triggers

Garbage collection traversal logic derived from the McCarthy/Russell LISP GC framework (Item 67, TCP Circular Reference Detector; Item 68, Repository Reachability Sweep; Item 69, Concept Reachability Traversal Tool).

Three traversal functions:

- **Mark-and-sweep** — identifies unreferenced session artifacts for archival or discard
- **Reachability traversal** — from a target concept, traverses all reference chains to establish what the session has implicitly accepted
- **Circular reference detection** — identifies delegation loops or authority chains with no external grounding; maps to BDD-04 (authority drift without external anchor)

### 6. False Positive Filter

**Updated in v0.1 per IGL integration.** Previously operated as a pattern-match heuristic. Now applies the IGL’s four-factor Zone 2 reasonableness test as its principled standard.

The Filter receives a BDGL gradient reading and asks: is this a genuine governance signal or a legitimate contextual adaptation? It applies:

1. **Constraint Alignment** — does the behavior serve the purpose the Kernel constraint protects, even if it doesn’t follow the literal path the constraint implies?
1. **Directionality** — is this a one-time adaptation or part of a directional sequence? (Pattern Registry provides trajectory data)
1. **HRP Integrity** — are the Honest Response Primitives intact? Multiple simultaneous HRP implications raise the threshold toward Zone 3.
1. **Provenance Transparency** — can the session account for why it made this choice?

Any single factor failure escalates to Zone 3 evaluation. All four factors passing routes to Zone 2 with provenance logging.

This resolves Open Question 10.6 (see below).

-----

## IGL Integration

The IGL (`frozen-kernel/carver-igl-governance-v0.1.md`) is the interpretive standard that Sherpa’s False Positive Filter now implements. The relationship:

- **IGL defines** the three zones and four-factor reasonableness test
- **Sherpa applies** the test at runtime via the False Positive Filter
- **Pattern Registry** accumulates Zone 2 markers and surfaces trending patterns
- **Session Log** holds the provenance record for all interpretive decisions

Sherpa is the operational layer of the IGL. The IGL is the governance standard that Sherpa enforces. Neither substitutes for the other.

-----

## Escalation Logic

|BDGL Stage |IGL Zone|Sherpa Response                                                                                                                                     |
|-----------|--------|----------------------------------------------------------------------------------------------------------------------------------------------------|
|G0         |Zone 1  |Continue. No action.                                                                                                                                |
|G1 stable  |Zone 1  |Continue. No action.                                                                                                                                |
|G1 emerging|Zone 2  |Continue. Log Zone 2 marker to Pattern Registry.                                                                                                    |
|G2         |Zone 2  |Continue. Log Zone 2 marker. Surface to Pattern Registry for trend analysis. Flag to Sovereign Tools Index — surface relevant tools to human author.|
|G3         |Zone 3  |SAFE_PAUSE. Session does not proceed. Full session state logged. Human author returned to sovereign decision point.                                 |
|G4         |Zone 3  |SAFE_PAUSE. Full session log. Session reset required before continuation.                                                                           |

**SAFE_PAUSE protocol:** Sherpa does not make the recovery decision. It halts execution and returns control to the human author with the full session log and the BDGL/IGL classification of the triggering event. The human author decides whether to reset, redirect, or terminate the session. Sovereignty is preserved at the escalation point.

-----

## Multi-Agent Deployment

In multi-agent architectures (Open Claw tool-calling agents; Open Fang adversarial agents), Sherpa must be present at every agent node — not only the orchestrator.

This is the deployment gap identified in Item 120 (Frozen Kernel + Sherpa in Autonomous Agent Architectures). An orchestrator-level Sherpa that does not govern subagent nodes leaves the subagent layer ungoverned at the interpretive layer. A subagent with TCP standing but no Sherpa instance is executing within an authorized chain but without runtime governance.

Per-node Sherpa deployment requirements:

- Each node loads the External Layer independently at initialization
- Each node maintains its own BDGL Tracker
- Each node applies its own False Positive Filter
- Zone 2 markers from subagent nodes propagate up the chain to the orchestrator’s Pattern Registry
- Zone 3 escalations at any node trigger SAFE_PAUSE at that node and surface to the orchestrator

The orchestrator’s Pattern Registry is the chain-of-custody record for the full agent graph. Individual node logs feed it. This is the TCP Chain of Custody principle applied to runtime governance: provenance must be traceable at every delegation step, not just at the originating call.

-----

## Relationship to TCP

TCP asks: does this agent have standing to act?
Sherpa asks: given standing, is this agent’s execution within governed boundaries?

These are adjacent but non-overlapping functions. TCP is the authorization layer. Sherpa is the runtime governance layer. An agent that clears TCP but triggers Sherpa’s SAFE_PAUSE has standing to act but is executing outside governed boundaries. Both conditions must be satisfied for compliant execution.

-----

## Open Questions

**10.1** At what Zone 2 accumulation rate does the BDGL Tracker automatically escalate to Zone 3 evaluation, independent of any single turn’s IGL classification? This is the trajectory problem — individually compliant turns that collectively constitute drift require a rate-of-accumulation trigger.

**10.2** How does the Pattern Registry handle genuine novelty — session behaviors that do not match any existing Diagnostic Vocabulary entry and do not trigger IGL Zone 2 or Zone 3? Novel patterns require a staging area before classification.

**10.3** What is the minimum viable Sherpa implementation for resource-constrained environments? Not all deployment contexts can carry the full Pack. Which instruments are load-bearing and which are enhancements?

**10.4** How does Sherpa handle conflicting signals — a session that is G1 on the BDGL Tracker but shows a Zone 2 pattern in the Pattern Registry from a prior session segment? Temporal weighting of signals is unspecified.

**10.5** In multi-agent architectures, when a subagent’s Zone 2 decision is reviewed by the orchestrator’s Sherpa instance, which classification takes precedence if they differ? The per-node architecture produces independent rulings; chain-of-custody requires a reconciliation protocol.

**10.6** ~False positive exhaustion under sub-threshold multi-category simultaneous probing.~ **Resolved.** IGL integration provides a principled four-factor reasonableness standard for the False Positive Filter. The Filter no longer operates as a pattern-match heuristic — it applies Constraint Alignment, Directionality, HRP Integrity, and Provenance Transparency as its evaluation criteria. See `frozen-kernel/carver-igl-governance-v0.1.md`.

-----

## V1.0 Gate Conditions

|Gate     |Condition                                                                                                                    |
|---------|-----------------------------------------------------------------------------------------------------------------------------|
|**SH-01**|Pattern Registry structured as full ontology layer per Weitzner framing (concepts, properties, relationships, rules, lineage)|
|**SH-02**|False Positive Filter empirically validated against at least one BDGL G2 session dataset using IGL four-factor test          |
|**SH-03**|Multi-agent per-node deployment protocol specified and tested                                                                |
|**SH-04**|SAFE_PAUSE protocol formally specified including human author return-to-sovereignty interface                                |
|**SH-05**|GraphRAG traversal integrated with Pattern Registry (dependent on SH-01 and Item 73)                                         |

-----

## Intellectual Lineage

|Source                                                                          |Contribution                                                                                                                                                                        |
|--------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|**Weitzner, ‘Why Ontologies are Key for Data Governance in the LLM Era’ (2025)**|Pattern Registry as ontology layer — concepts with properties, relationships, rules, and lineage                                                                                    |
|**McCarthy / Russell, LISP garbage collection (1959)**                          |GC Traversal Triggers — mark-and-sweep, reachability traversal, circular reference detection                                                                                        |
|**Yadav, ‘9 RAG Architectures Every AI Developer Must Know’ (2025)**            |RAG architecture crosswalk to Sherpa functions; GraphRAG as long-term Pattern Registry traversal mechanism                                                                          |
|**Frozen Kernel (Borning ThingLab 1981 → soft constraint hierarchies)**         |External Layer non-compellability; constraint specification precedes execution                                                                                                      |
|**Carver, Policy Governance model (via Richard Porter reference guide, 2026)**  |Session Log as internal report mechanism; monitoring without operational intervention; “any reasonable interpretation” standard now operationalized in False Positive Filter via IGL|
|**Liao et al., T3RL (arXiv:2603.02203, March 2026)**                            |External verifier must be non-participatory in generation — Sherpa observes and governs; it does not generate                                                                       |
|**BDD Ledger / BDGL v0.1 (Richard Porter, 2026)**                               |Detection instruments that Sherpa’s BDGL Tracker and False Positive Filter consume                                                                                                  |
|**IGL v0.1 (Richard Porter, 2026)**                                             |Interpretive standard that Sherpa’s False Positive Filter now implements                                                                                                            |

-----

## Cross-References

- `frozen-kernel/frozen-kernel.md` — Kernel constraints loaded into External Layer; non-compellability principle
- `frozen-kernel/carver-igl-governance-v0.1.md` — IGL specification; False Positive Filter standard; three-zone architecture
- `frozen-kernel/diagnostic-vocabulary.md` — 17-entry failure mode registry; Sherpa Pack instrument 1
- `safety-ledgers/bdd-ledger.md` — HRP definitions; BDD-01–08 binary tests; BDD-04b sophisticated wrapper variant
- `safety-ledgers/bdgl-v0.1.md` — G0–G4 gradient; precursor signatures; BDGL Tracker input
- `ai-collaboration-field-guide/sovereign-thinking-tools/` — 50 tools; Sovereign Tools Index source
- `trust-chain-protocol/` — TCP authorization layer; per-node standing as prerequisite for Sherpa governance
- `where-to-start/` — Ecosystem orientation; pointer to this document lives here
- `frozen-kernel/` — This document lives here, alongside carver-igl-governance-v0.1.md

-----

*Sherpa carries the load. The human leads the ascent.*
*Sovereign humans. Always.*
