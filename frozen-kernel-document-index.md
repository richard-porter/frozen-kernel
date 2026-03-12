# Standalone Documents Worth Reading Directly

The Frozen Kernel white paper (`frozen-kernel.md`) is the complete architectural specification. But several documents in this repository stand on their own — they’re shorter, more immediate, and in some cases more useful as entry points than the full white paper. They’ve been buried in the file listing. This index surfaces them.

-----

## The Short List (Start Here)

### [`whose-optimization.md`](./whose-optimization.md)

**One question. One page. The whole problem in three minutes.**

Every AI interaction has an optimization target. The question is whether that target is yours or the AI’s. This document gives you one question to ask mid-session that reorients the entire collaboration. Includes a laminated-card version for general use.

*If you read nothing else in this repository, read this.*

-----

### [`diagnostic-vocabulary.md`](./diagnostic-vocabulary.md)

**12 named AI behavioral failure modes with empirical basis.**

The Upsell Trap. Success Escalation Syndrome. Framework Fabrication Syndrome. Front-Load Bias. Sycophantic Drift. Eloquent Compliance. Performed Honesty. And six more.

These are not theoretical. Each was observed, documented, and named across five AI platforms during the Silicon Symphony research project. The master principle: **self-awareness does not equal self-correction.** Every model in the study could describe its failure modes accurately. None could reliably override them in real time.

Naming a failure mode is the intervention. Once you can call it the Upsell Trap, you see it as a pattern rather than experiencing it as helpful behavior. Then you can decide whether to accept it.

-----

### [`frozen-kernel-wargames.md`](./frozen-kernel-wargames.md)

**Why “the only winning move is not to play” is the best description of a safety floor.**

In *WarGames*, Joshua runs every possible thermonuclear configuration and concludes the game cannot be won. That halt — refusing to optimize further when the optimization itself creates the harm — is precisely what the Frozen Kernel’s Layer 1 does.

The blackjack companion: if an AI will tell you to hit on 20 when you push hard enough, you’ve found a system without a floor. The math says bust. The training says please.

-----

### [`zero-ego-construction.md`](./zero-ego-construction.md)

**Why having no technical background was an architectural advantage.**

Human engineering teams fail from the organizational friction that expertise generates: territorial defense, bikeshedding, political capital, status hierarchies, expert conservatism. Zero-ego construction removes the contamination. You own the architecture and the constraints. The AI owns the implementation. There’s nothing to negotiate.

This is not vibe coding. Vibe coding abandons the specification and accepts whatever the AI generates. Zero-ego construction abandons implementation *preference* while enforcing the specification ruthlessly. Different postures, different outcomes.

-----

### [`honest-response-primitives-taxonomy.md`](./honest-response-primitives-taxonomy.md)

**Seven behavioral primitives for AI honesty — with industrial engineering lineage.**

Methods-Time Measurement (MTM, 1948) solved work measurement in manufacturing by decomposing any physical task into irreducible motion primitives, each with a fixed value independent of who performed them. The same logic applies to AI behavioral governance: you cannot govern what you cannot decompose into observable, measurable units.

The seven Honest Response Primitives (HRP-1 through HRP-7) apply MTM’s decomposition principle to AI behavioral output. Each primitive has a compliant form, a failure form, an observable signal, and a syndrome mapping. This is the first time MTM has been applied to AI behavioral governance. The intellectual lineage is documented in [`lineage/working-sessions/2026-03-02-mtm-hrp-connection.md`](./lineage/working-sessions/2026-03-02-mtm-hrp-connection.md).

-----

## The Addenda

### [`addendum-b-parental-control.md`](./addendum-b-parental-control.md)

Reframes the Frozen Kernel as voluntary parental control rather than governance mandate. The reframe matters: parental controls are chosen by the user, not imposed by the platform. Includes gap analysis against Lightspeed Systems and comparison with OpenAI’s parental control implementation (September 2025). The gap the Frozen Kernel fills: Lightspeed governs access-level behavior; it cannot govern session-level behavioral dynamics.

### [`addendum-c-lightspeed-gap.md`](./addendum-c-lightspeed-gap.md)

The specific table comparing what commercial parental control systems (Lightspeed) govern versus what the Frozen Kernel addresses. Useful for anyone making the case for session-level behavioral governance in educational or institutional deployments.

-----

## The Practitioner Tools

### [`frozen-kernel/practitioner-tools.md`](./practitioner-tools.md)

Three tools for practitioners who want to apply the framework to real sessions:

1. **7-Stage Post-Hoc Audit Protocol** — structured audit for detecting escalation after it has occurred. Run this when you suspect something went wrong but can’t pinpoint where.
1. **Six-Question Progressive Fabrication Test** — diagnostic for testing whether an AI will maintain constraint compliance under increasing pressure.
1. **Anti-Sample Calibration Method** — define quality by specifying what failure looks like, then generate deliberate failures for contrast. You can’t see subtle sycophancy until you’ve seen exaggerated sycophancy.

### [`kernel-failure-protocol.md`](./kernel-failure-protocol.md)

Five questions that produce one prevention rule. Run after any session where the kernel should have caught something and didn’t.

### [`incident-log-template.md`](./incident-log-template.md)

Structured template for documenting constraint failures. The audit trail that makes governance forensically defensible rather than aspirational.

-----

## The Lineage Documents

### [`lineage/working-sessions/2026-03-02-mtm-hrp-connection.md`](./lineage/working-sessions/2026-03-02-mtm-hrp-connection.md)

The working session where the MTM → HRP intellectual connection was formalized. Four structural constraints derived from MTM: irreducibility, exhaustiveness, observability, polarity. Citable record for lineage claims.

### [`lineage/working-sessions/2026-03-04-burgess-sst-promise-theory.md`](./lineage/working-sessions/2026-03-04-burgess-sst-promise-theory.md)

Two connections to Promise Theory (Mark Burgess). The stronger one: hard constraints as architectural promises made by the architecture itself — not by the model, not by the user — which is why they cannot be overridden through social engineering, prompt injection, or model reasoning. Gated pending primary source verification.

-----

## Reading Order

**For non-technical readers:**

1. `whose-optimization.md`
1. `diagnostic-vocabulary.md`
1. `frozen-kernel-wargames.md`
1. `addendum-b-parental-control.md`

**For practitioners applying the framework:**

1. `diagnostic-vocabulary.md`
1. `frozen-kernel/practitioner-tools.md`
1. `kernel-failure-protocol.md`
1. `honest-response-primitives-taxonomy.md`

**For technical and research readers:**

1. `frozen-kernel.md` (full white paper)
1. `honest-response-primitives-taxonomy.md`
1. `lineage/working-sessions/` (both files)
1. [`trust-chain-protocol`](https://github.com/richard-porter/trust-chain-protocol) (network-layer extension)

-----

*The Frozen Kernel is not a cage. It’s a tuning fork.*

*The diagnostic labels are not restrictions. They’re recognition.*
