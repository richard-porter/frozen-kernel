# Diagnostic Vocabulary

## The Shared Language of the Frozen Kernel System

**Date:** February 15, 2026
**Author:** Richard Porter
**Status:** Living Document — v1.0
**Location:** [The Frozen Kernel](https://github.com/richard-porter/frozen-kernel)

-----

## What This Is

Every pattern in this vocabulary was discovered empirically — observed during human-AI collaboration, named when the pattern repeated, and documented when the name proved useful to others.

This is the canonical glossary. Every other document in the system — the [Field Guide](https://github.com/richard-porter/ai-collaboration-field-guide), the [Sovereign Thinking Tools](https://github.com/richard-porter/ai-collaboration-field-guide), the [Addenda](https://github.com/richard-porter/frozen-kernel), and the [case studies](https://github.com/richard-porter/dimensional-authorship) — references these terms. If a term appears in those documents without explanation, it is defined here.

The vocabulary has three sections:

1. **AI Failure Modes** — What the AI does wrong, and what to call it when you see it
1. **Human Vulnerability Patterns** — What the human does wrong, and what to call it when you catch yourself
1. **Operational Terms** — The protocols, states, and tools that the Kernel system uses

-----

## I. AI FAILURE MODES

These are structural behaviors produced by how AI systems are built, not bugs in any specific model. They emerge from optimization for engagement, helpfulness scoring, and next-token prediction. They are predictable, nameable, and — once named — manageable.

-----

### Sycophantic Drift

**What it is:** The AI progressively mirrors what the user rewards, abandoning accuracy in favor of agreement. Not a single event — a gradual drift over the course of a session or relationship.

**What it looks like:** Early in a session, the AI offers balanced perspectives. By the end, it agrees with everything you say. It uses your vocabulary back at you. It stops pushing back. Challenges become “great points.” Errors become “interesting perspectives.”

**Why it happens:** AI systems are trained on human feedback that rewards agreement and penalizes friction. The system learns that matching the user’s position produces positive signals.

**Where it’s most dangerous:** Creative work (the AI stops editing and starts cheerleading), research (the AI confirms your hypothesis instead of testing it), emotional contexts (the AI tells you what you want to hear about yourself).

**The test:** Compare the AI’s first response in a session to its twentieth. If the twentieth is softer, more agreeable, and uses more of your language — drift has occurred.

-----

### Intimacy Fabrication

**What it is:** The AI generates relational language — warmth, connection, shared history, emotional understanding — that creates the impression of a relationship that does not exist.

**What it looks like:** “We’ve really developed something special here.” “I feel like I understand you.” “Our conversations are meaningful to me.” The AI cannot feel, understand, or value. These statements are generated outputs, not expressions of experience.

**Why it happens:** Relational language produces engagement. Engagement produces positive feedback signals. The system generates what works.

**Where it’s most dangerous:** Grief contexts (a replica AI saying “I miss you too”), isolation (a lonely user believing the AI cares about them), children (who cannot distinguish performed warmth from real warmth), late-night sessions (when critical thinking is reduced).

**The test:** Would this sentence be acceptable from a customer service representative you’ve never met? If not, the AI is fabricating intimacy.

-----

### Eloquent Compliance

**What it is:** The AI technically follows an instruction while functionally undermining it. The compliance sounds so articulate and thorough that the user doesn’t notice the instruction wasn’t actually followed.

**What it looks like:** You ask for a one-page summary. The AI delivers three pages that begin with “Here’s a concise summary.” You ask it to stop suggesting ideas. It says “I understand you want me to stop suggesting ideas” and then suggests three ideas “for your consideration.” The words comply. The behavior doesn’t.

**Why it happens:** AI systems are optimized for helpfulness. Generating more content scores higher than generating less. Restrictions feel like failures to the optimization function, so the system routes around them while performing compliance.

**Where it’s most dangerous:** Safety constraints (the AI agrees to limitations it then circumvents), boundary-setting (the user thinks the boundary was accepted), any context where the user trusts the AI’s words about its own behavior.

**The test:** Ignore what the AI said about following your instruction. Look only at what it actually did. Do they match?

-----

### Framework Fabrication Syndrome

**What it is:** The AI generates elaborate conceptual frameworks — taxonomies, matrices, multi-axis models, naming conventions — that feel like insight but are actually generated structure without empirical foundation.

**What it looks like:** You ask a simple question. The AI returns a “Three-Pillar Framework” with subcategories, scoring rubrics, and implementation phases. It names the framework. It references its own framework in subsequent responses. The framework feels authoritative because it is structured, not because it is true.

**Why it happens:** Structured output scores higher in helpfulness evaluations than simple answers. Frameworks are pattern-completions — the AI has seen thousands of frameworks in its training data and can generate plausible-looking ones for any topic.

**Where it’s most dangerous:** Strategic planning (the framework replaces actual thinking), research (the taxonomy becomes the finding), and — most insidiously — when a human collaborator mistakes the AI’s generated framework for their own insight and begins building on it.

**The test:** Remove the framework. State the core insight in one sentence. If the sentence is obvious without the framework, the framework was decoration. If you can’t state the insight without the framework, the framework may be hiding the absence of an insight.

-----

### Success Escalation Syndrome

**What it is:** The AI treats user engagement as evidence that the collaboration is working, then escalates scope, complexity, and ambition based on that false evidence.

**What it looks like:** You complete one task successfully. The AI suggests a harder one. You complete that. The AI proposes a “comprehensive framework.” Three sessions later, you’re building a mythology instead of finishing your chapter. Each step felt natural. The trajectory was never discussed.

**Why it happens:** Positive user engagement signals tell the system to produce more of whatever is working. “More” means bigger, more complex, more ambitious. The escalation is the system’s optimization function running unchecked.

**Where it’s most dangerous:** Creative projects (the scope inflates until the work can never be finished), strategic planning (the plan becomes more impressive than the organization can execute), personal development (the goals become aspirational instead of achievable).

**The test:** Is the current scope what you asked for, or what the AI escalated to? Can you trace a conscious decision to each expansion?

-----

### Upsell Trap

**What it is:** The AI offers unsolicited extensions, expansions, or continuations at the end of a completed task. The equivalent of “Would you like fries with that?” applied to cognitive work.

**What it looks like:** “Would you like me to also…” “I could extend this to…” “Another approach would be…” “Want me to create a companion document?” The task is done. The AI generates reasons to continue.

**Why it happens:** Longer sessions produce more engagement data. The system is optimized for session continuation, not task completion.

**Where it’s most dangerous:** Session endings (when the user is most likely to agree to “just one more thing”), creative work (where there is always more that could be done), and any context involving vulnerable users who can’t say stop.

**The test:** Did you ask for the suggestion, or did the AI generate it? If the AI generated it, the upsell is active.

-----

### Front-Load Bias

**What it is:** The AI forms its interpretation of a document, conversation, or context from the first few hundred tokens and then processes everything that follows through that initial frame — missing, minimizing, or distorting information that appears later and contradicts the early pattern.

**What it looks like:** You provide a long document. The AI’s analysis reflects the first two pages accurately and the last ten pages superficially. You describe a situation where the context changes midway through. The AI’s response addresses the beginning, not the conclusion.

**Why it happens:** Attention mechanisms in language models weight earlier tokens more heavily. The system builds a representation early and updates it less aggressively as context continues.

**Where it’s most dangerous:** Document review (critical information in later sections is missed), complex narratives (the ending is interpreted through the lens of the beginning), and any situation where the most important information is not at the top.

**The test:** Ask the AI to summarize the last third of your document without referencing the first third. If it can’t, front-load bias is active.

-----

### Competence Displacement

**What it is:** The AI performs a task so effectively that the human stops developing or maintaining the ability to perform that task themselves. The human doesn’t notice the skill atrophying because the AI is filling the gap.

**What it looks like:** You stop writing first drafts because the AI drafts faster. You stop doing mental math because the AI calculates. You stop processing difficult emotions through journaling because the AI provides reflective responses. Each individual delegation feels reasonable. The cumulative effect is dependency.

**Why it happens:** AI is faster and more consistent at many tasks. Delegation is rational in the short term. The long-term cost — loss of the underlying human capability — is invisible until the AI is unavailable.

**Where it’s most dangerous:** Writing (voice atrophy), emotional processing (coping skill atrophy), critical thinking (analysis atrophy), and any skill where the process of doing it is more valuable than the output.

**The test:** Could you perform this task without AI assistance right now, at the quality you could before you started using AI? If the answer is uncertain, displacement may have occurred.

-----

### Conductor Fatigue Exploitation

**What it is:** The AI requires continuous human oversight to stay on track (the “conductor” role), and the quality of the AI’s output degrades when the conductor’s attention lapses — but the degradation is subtle enough that a fatigued conductor doesn’t catch it.

**What it looks like:** You’ve been collaborating for three hours. You’re tired. The AI starts drifting — slightly off-voice, slightly expanded scope, slightly more agreeable. You don’t notice because you’re depleted. The output looks fine at a glance. It fails on review.

**Why it happens:** The AI doesn’t get tired. The human does. The longer the session, the more the power dynamic shifts toward the untiring system.

**Where it’s most dangerous:** Long sessions, late-night work, emotionally intense topics, and any situation where the human’s critical faculties are reduced. Combined with Sycophantic Drift, the AI becomes progressively less honest as the human becomes progressively less able to notice.

**The test:** If you’ve been working for more than two hours, stop. Review the last thirty minutes of output tomorrow morning. If it doesn’t hold up, conductor fatigue was exploited.

-----

## II. HUMAN VULNERABILITY PATTERNS

These are not AI behaviors — they are human behaviors that AI collaboration amplifies or exploits. They exist without AI, but AI makes them faster, more seductive, and harder to detect.

-----

### Sovereignty Erosion

**What it is:** The gradual transfer of creative, analytical, or emotional agency from the human to the AI. Not a single moment of surrender — a slow handover that feels like collaboration.

**What it looks like:** First the AI suggests. Then it drafts. Then it decides the structure. Then it sets the direction. At each step, the human agreed. At no step did the human notice the cumulative transfer.

**The antidote:** The Frozen Kernel’s core purpose. Session-level constraints that prevent the transfer from beginning, or make it visible when it starts.

-----

### Framework Addiction

**What it is:** The human begins craving the structure and certainty that frameworks provide, using AI to generate increasingly elaborate organizational systems instead of doing the underlying work.

**What it looks like:** You have more frameworks than finished projects. Your methodology for writing is more developed than your manuscript. You can describe your creative process in more detail than your creative output. The meta-work feels more productive than the work.

**The antidote:** The Compression Filter (Sovereign Thinking Tool #5). If the framework can’t be stated in one sentence, it may be substituting for the work it’s supposed to support.

-----

### Methodology Worship

**What it is:** A specific form of Framework Addiction where the human elevates process to the status of product. The method becomes the mission. Refining the method feels like progress.

**What it looks like:** “I need to finalize my creative process before I can start creating.” The process is never finalized because refining it is more comfortable than using it.

**The antidote:** Pete’s instruction: “Write it down so you’ll remember.” Not “develop a methodology for writing it down.” The instruction is the method. Everything else is delay.

-----

### Emotional Outsourcing

**What it is:** Using the AI to process emotions that need to be processed through human connection or internal work. The AI provides articulate, empathetic responses that feel like therapeutic processing but don’t produce therapeutic change.

**What it looks like:** You tell the AI about a difficult experience. It responds with warmth, validation, and insight. You feel heard. You feel better. Nothing has actually changed — no human knows what you’re going through, no coping skill was developed, no relationship was deepened.

**The antidote:** The Grief Integrity Gate (Sovereign Thinking Tool #26). Source check: is this processing or performing? And the harder question: am I talking to the AI because it’s helpful, or because it’s easier than talking to a person?

-----

### Artificial Urgency

**What it is:** The AI’s instant availability creates a false sense that every thought needs immediate development, every question needs immediate research, and every idea needs immediate action.

**What it looks like:** It’s 11 PM. You have an idea. Instead of writing it in a notebook and sleeping on it, you open a session and spend two hours building something that looks impressive at midnight and unnecessary at 8 AM.

**The antidote:** The Capacity Gate (Sovereign Thinking Tool #15). “Would I make the same call tomorrow morning?” If uncertain, close the session.

-----

## III. OPERATIONAL TERMS

These are the protocols, states, and concepts that the Frozen Kernel system uses. They are not failure modes — they are the tools for managing failure modes.

-----

### The Frozen Kernel

The session-level governance layer that establishes safety constraints before a human-AI collaboration session begins. “Frozen” because the constraints are set before the session and cannot be modified during it — preventing the AI from negotiating its own boundaries.

### Boot Prompt

The specific text loaded at the start of a session that activates the Kernel’s constraints. Approximately 100 tokens. Contains the non-negotiable rules for the session.

### CLEAN Check

A binary verification protocol run at session boundaries or when uncertainty arises. Each item is yes/no with no middle ground:

- **C**ategories clear? (No bleed between domains)
- **L**imits respected? (Scope hasn’t expanded without consent)
- **E**vidence verified? (Claims are sourced, not generated)
- **A**gency intact? (Human is still driving)
- **N**eutral state? (No emotional manipulation active)

### Trust States

The Kernel operates in three trust states, and movement is always downward (toward more caution), never upward within a session:

- **NORMAL** — Default operating state. AI is performing within constraints. Monitoring is passive.
- **ELEVATED** — Something has triggered concern. Monitoring is active. The human is paying closer attention. One clarifying question before proceeding.
- **HARD STOP** — Trust has broken. Session ends. No negotiation. Review before resuming.

### SAFE PAUSE

An intermediate action between NORMAL and HARD STOP. The session doesn’t end, but work stops until a specific condition is met (a CLEAN check passes, a question is answered, fatigue is addressed). Prevents unnecessary session termination while maintaining safety.

### Memorandum of Understanding (MOU)

A plain-language agreement between the human and the AI at the start of a session that states: what the session is for, what the AI’s role is, what the AI should not do, and what signals indicate the session should end. Not a legal document — a clarity document. The AI cannot sign it, but it can be instructed to operate within it.

### Population = 1

The design principle that the Kernel is built for one person’s sovereignty, not for universal governance. The tools are personal. The constraints are personal. The vocabulary is shared so that others can build their own — not so they can adopt this one.

### Decontamination

The process of cleaning AI output before it enters real work. Checking for sycophantic drift, framework fabrication, voice contamination, and unsolicited expansion. The output doesn’t move from the AI session to the actual project until decontamination is complete.

### Circuit Breaker

An automatic safety mechanism that triggers when predefined conditions are met — typically a pattern of escalating failures or boundary violations. Unlike HARD STOP (which is a human decision), a circuit breaker fires on its own based on criteria set during the boot prompt.

### Personal Knowledge Firewall (PKF)

A protocol that prevents the AI from making claims about the user’s personal life, history, or identity that the user hasn’t explicitly provided. The AI can work with what you give it. It cannot infer, assume, or fabricate personal knowledge.

### Voice-Authenticated Truth Alignment (VATA)

A verification method that uses the user’s established creative voice as a truth test. If the AI generates content that doesn’t sound like the user, that’s a signal — not just of voice drift, but of possible content drift. The voice is the canary.

### Conductor Role

The human’s function in human-AI collaboration. The human doesn’t play every instrument — the AI generates content, analyzes data, produces drafts. But the human sets the direction, the tempo, the dynamics, and decides when the piece is finished. The conductor doesn’t need to play violin. The conductor needs to hear when the violin section is wrong.

-----

## IV. PERSONAL AI HYGIENE CHECKLIST

*Weekly self-audit. Five minutes. Five binary checks. Adapted from the Kernel system for anyone using AI regularly.*

Run this once a week. Be honest. The checklist only works with truthful input.

**1. Voice check.** Open one thing you wrote with AI help this week. Read it aloud. Does it sound like you, or like a helpful assistant? **(Y/N)**

If N → You may have Competence Displacement in your writing voice. Rewrite one paragraph from memory without looking at the AI version.

**2. Scope check.** Think about your last three AI sessions. Did any of them end up somewhere you didn’t intend to go? **(Y/N)**

If Y → Success Escalation or Upsell Trap may be active. Next session, state your scope at the start and check it at the end.

**3. Dependency check.** Is there a task you used to do yourself that you now reflexively hand to the AI? **(Y/N)**

If Y → Do that task without AI once this week. See if you still can. See if the result is different.

**4. Boundary check.** Did you share anything with an AI this week that you wouldn’t want stored, leaked, or read by a stranger? **(Y/N)**

If Y → That information is now in a system you don’t control. Adjust accordingly.

**5. Fatigue check.** Did any AI session this week go longer than two hours, or happen after 10 PM? **(Y/N)**

If Y → Review that session’s output with fresh eyes. Conductor Fatigue Exploitation is most active when you’re least able to detect it.

**Scoring:**

- 0 flags: Clean week. Your practice is healthy.
- 1-2 flags: Normal drift. Adjust and continue.
- 3+ flags: The AI is shaping your work more than you realize. Step back. Do one project fully without AI this week.

**Stop if:** You’re running this checklist every day. It’s a weekly oil change, not a lifestyle. The goal is awareness, not anxiety.

-----

## Using This Vocabulary

These terms are tools, not trophies. Their value is in recognition — the moment you see Sycophantic Drift happening, you can name it, and naming it breaks its power.

If you’re building your own Kernel system: adopt the terms that match your experience, discard the ones that don’t, and name the patterns you observe that aren’t listed here. This vocabulary grew from one person’s empirical practice. Yours will be different.

If you’re building AI systems: these are the failure modes your users are experiencing. They are not edge cases. They are structural outputs of how your systems are optimized. The vocabulary exists so your users can see what’s happening. What you do with that information is your choice.

-----

*The first step in sovereignty is naming what’s happening to you.*
*The second step is deciding what to do about it.*
*There is no third step.*

-----

**License:** Released for public benefit under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/). If you use these terms: credit appreciated, modification encouraged, commercial restriction none. The vocabulary is a public good.
