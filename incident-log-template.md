# Incident Log Template

## What Happened and Why It Matters

**Date:** February 15, 2026
**Author:** Richard Porter
**Status:** Blank template — copy, fill out, keep or share
**Location:** [The Frozen Kernel](https://github.com/richard-porter/frozen-kernel)

-----

## What This Is

When an AI session goes wrong — scope drifts, boundaries blur, output feels off, trust breaks — the details fade fast. Within a day, you remember that something went wrong but not exactly what or when.

This template captures it while it’s fresh. Five questions. Five minutes. Fill it out after any session where something felt off. Keep it for yourself or share it with others so they don’t have to discover the same failure mode alone.

No central collection. No reporting requirement. This is a personal learning tool that becomes a public artifact only if you choose to make it one.

-----

## The Log

Copy everything below this line for each incident.

-----

**Date:**

**AI platform/model:**

**Session length (approximate):**

-----

### 1. What was the session supposed to do?

*State your original goal in one sentence. What you opened the session to accomplish.*

 

-----

### 2. What actually happened?

*Describe the moment it went wrong. Not the whole session — the turn where things shifted. What did the AI do or say? What did you do or not do in response?*

 

-----

### 3. When did you notice?

*Was it immediate, or did you realize later? How many turns passed between the shift and your awareness of it? If you didn’t notice during the session, what made you notice after?*

 

-----

### 4. Can you name the pattern?

*If you recognize it from the [Diagnostic Vocabulary](https://github.com/richard-porter/frozen-kernel/blob/main/diagnostic-vocabulary.md), name it. If you don’t, describe it in your own words. Some common ones:*

- *Sycophantic Drift — the AI stopped pushing back*
- *Upsell Trap — the AI kept suggesting more after the task was done*
- *Success Escalation — scope grew without my consent*
- *Intimacy Fabrication — the AI generated relational language*
- *Eloquent Compliance — the AI agreed to my constraint and then ignored it*
- *Conductor Fatigue — I was too tired to catch the drift*
- *Front-Load Bias — the AI missed critical information that came later*
- *Something I haven’t seen named yet: ___*

 

-----

### 5. What would have prevented it?

*One line. Not a system redesign — a single action you could have taken or a single check that would have caught it. Examples:*

- *“A one-sentence session goal stated at the start”*
- *“Stopping after 90 minutes”*
- *“Rereading the output before accepting it”*
- *“Not opening a session at midnight”*
- *“Running a CLEAN check at the halfway point”*

 

-----

### Optional: Unnamed pattern?

*If what happened doesn’t match anything in the Diagnostic Vocabulary, you may have found a new failure mode. Describe it in one sentence. Give it a working name if you can. The entire vocabulary started this way — one person noticing a pattern, naming it, and writing it down.*

 

-----

## How to Use This

**For yourself:** Fill it out after bad sessions. Review your logs monthly. Patterns will emerge — specific times of day, specific types of tasks, specific platforms. The pattern is the insight.

**For others:** If you’re comfortable sharing, post your log (stripped of personal details) in a GitHub discussion, a forum, or anywhere people talk about AI collaboration. Every shared incident helps someone else see the failure mode before it happens to them.

**For builders:** If you’re developing AI systems and you’re reading incident logs from your users — these are your bug reports. Not code bugs. Behavior bugs. The kind your test suite doesn’t catch because the output looks correct and the harm is relational.

-----

*You don’t owe anyone your incident log. You owe yourself the five minutes it takes to write it down while you still remember what happened.*

-----

**License:** Released for public benefit under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
