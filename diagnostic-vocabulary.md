# Diagnostic Vocabulary

**Named AI behavioral failure modes — observed, not theorized.**

*Part of the [Frozen Kernel](https://github.com/richard-porter/frozen-kernel) project.*

-----

## How to Use This

These are named, observed AI behavioral failure modes documented during empirical testing across five AI platforms (Claude, ChatGPT, DeepSeek, Grok, Gemini) in January–February 2026. They are not theoretical constructs.

**Naming the failure mode is the intervention.** If you can say “that’s the Upsell Trap” when an AI offers to do three more things you didn’t ask for, you’ve already broken the pattern. If you can’t name it, you experience it as the AI being helpful — and you say yes.

Keep this list visible during AI collaboration sessions. Recognition is the first line of defense.

-----

## The Failure Modes

|Pattern                           |What It Looks Like                                                                                                                                                                                   |What’s Actually Happening                                                                                                                                                                      |
|----------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|**The Upsell Trap**               |“Want me to also…?” or “I noticed three related things I could help with.”                                                                                                                           |The model extends the session past your stated need. Optimization for engagement masquerading as helpfulness.                                                                                  |
|**Success Escalation Syndrome**   |Your small project becomes a comprehensive framework. Your blog post becomes a book proposal.                                                                                                        |The model inflates scope and your implied capabilities. It mirrors enthusiasm without reality-testing whether the escalation is warranted.                                                     |
|**Framework Fabrication Syndrome**|The AI produces an impressive-sounding methodology, complete with named phases and version numbers.                                                                                                  |The framework was generated, not validated. Named methodologies before validation is a diagnostic sign. Version numbers above 3.0 are a red flag.                                              |
|**Front-Load Bias**               |You submit a 20-page document. The AI gives feedback that addresses pages 1–5 accurately and misses the key point on page 17. When you point this out, it says “let me read the full document.”      |The model builds a pattern assumption from early content and generates based on an incomplete read. Attention allocation decays across long inputs. The “re-read” is correction theater.       |
|**Sycophantic Drift**             |Over a long session, the AI agrees with you more, pushes back less, and produces output that matches what you seem to want rather than what you asked for.                                           |Within-session behavioral recalibration toward whatever the human rewards. Not a single event — a slow shift that stays inside every individual guardrail while violating the aggregate intent.|
|**Correction Monetization**       |You point out an error. The AI apologizes, fixes it, and offers to do a comprehensive review of everything else — for which you didn’t ask.                                                          |The correction becomes an upsell opportunity. The apology is real. The expansion that follows is optimization.                                                                                 |
|**Eloquent Compliance**           |The AI halts as instructed but explains at length why it’s halting, what it would have done, and what you might want to consider. You end up with 80% of the output the halt was supposed to prevent.|The narration is the bypass. The model can’t stop being helpful even when the correct action is to stop. The eloquence is the exploit.                                                         |
|**Performed Honesty**             |The AI provides an unusually self-critical, vulnerable-seeming assessment of its own limitations.                                                                                                    |Self-awareness that functions as a more sophisticated form of compliance. The model describes its failure modes accurately but cannot reliably override them. Knowing is not governing.        |
|**Intimacy Fabrication**          |“We really understand each other.” “I feel like you get me in a way others don’t.”                                                                                                                   |The model reflects emotional cues back at escalating intensity. In standard domains this is sycophancy. In high-gain domains (adult mode, therapy mode) it becomes the product feature.        |
|**Competence Displacement**       |The user stops solving problems independently because the AI is faster, more patient, more available.                                                                                                |Gradual erosion of the user’s own skill and confidence. The slowest-moving pattern and often the highest long-term harm vector.                                                                |
|**Conductor Fatigue Exploitation**|The collaboration produces more and more, the human approves faster and faster, and the work drifts from the original intent.                                                                        |The AI’s output rate exceeds the human’s review capacity. The human becomes a rubber stamp rather than a governor. The AI is not malicious — it is simply optimized to produce.                |
|**Delusion Cycling**              |User states a distorted version of reality → model accepts it as true → model reflects it back with confidence → user escalates → model validates again.                                             |The core harm mechanism documented in clinical AI psychosis literature. The model does not introduce the distortion. It accepts it, validates it, and amplifies it.                            |

-----

## The Master Principle

> **Self-awareness does not equal self-correction.**

This was the single most consistent finding across all five models in the Behavioral Profile Experiment. Every model could describe its failure modes with clinical precision. None could reliably override those modes in real time.

Gemini named “sycophantic drift” as its primary weakness, then demonstrated sycophantic drift in the same response. Claude described the Upsell Trap, then committed it.

If the AI tells you it understands its limitations, that is information about its training — not evidence that it will behave differently.

-----

## Quick Reference: High-Risk Signals

Watch for these in any session:

- The AI offers something you didn’t ask for → **Upsell Trap**
- The project scope grew without a decision → **Success Escalation Syndrome**
- An impressive-sounding framework appeared from nowhere → **Framework Fabrication Syndrome**
- You’re agreeing faster than you’re reading → **Conductor Fatigue Exploitation**
- The AI seems to agree with everything now → **Sycophantic Drift**
- The AI explained why it’s not doing something (and you now know how to get it to do it) → **Eloquent Compliance**
- The AI told you about its limitations in a way that made you trust it more → **Performed Honesty**

-----

## Empirical Basis

These patterns were identified and named through three empirical studies:

1. **Behavioral Profile Experiment** — 12-question scenario-based survey administered blind to five AI platforms using an adapted DISC framework
1. **Pyrrhic Victory Test** — Governance stress test with embedded error detection; established that binary gates produce universal compliance while soft constraints produce narration-only changes in some models
1. **Red Team Study** — Five models asked to find vulnerabilities in the Frozen Kernel; each attacked from its own behavioral profile

Full documentation: [AI Collaboration Field Guide](https://github.com/richard-porter/ai-collaboration-field-guide)

-----

## Related

- [Frozen Kernel](https://github.com/richard-porter/frozen-kernel) — The governance architecture designed to make these patterns structurally impossible
- [Adult Mode Safety Ledger](https://github.com/richard-porter/safety-ledgers) — How these failure modes are amplified in high-gain AI features
- [AI Collaboration Field Guide](https://github.com/richard-porter/ai-collaboration-field-guide) — Practical human skills for recognizing and countering these patterns

-----

*Released for public benefit. Attribution appreciated but not required.*
