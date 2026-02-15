# Frozen Kernel — Addendum C

## The Lightspeed Gap: Where Content Filtering Ends and Session Governance Begins

**Date:** February 15, 2026
**Author:** Richard Porter
**Status:** Gap analysis
**Companion to:** [Addendum B: Voluntary Parental Control](https://github.com/richard-porter/frozen-kernel/blob/main/addendum-b-parental-control.md)

-----

## Context

Hamilton Southeastern Schools (Fishers, Indiana) currently partners with Lightspeed Systems to give parents visibility and control over their children’s internet usage on school-issued devices. Lightspeed provides web content filtering, weekly activity reports, session pause controls, and category-level toggles (social media, YouTube, etc.).

This is good architecture. It is also incomplete — because it was built for a threat surface that preceded AI chatbots.

Lightspeed governs **access**. It cannot govern **interaction**. That gap is where AI-specific harm occurs.

-----

## What Lightspeed Does Well (and What the Kernel Should Learn From It)

**Weekly activity reports.** Every Sunday, parents receive a digest: top sites visited, daily page counts, blocked attempts. No equivalent exists for AI interaction. A weekly summary of AI session patterns — session count, total time, scope drift incidents, trust state changes — would give parents the same baseline awareness for AI that Lightspeed already provides for web browsing.

**Category-level toggles.** Lightspeed lets parents allow YouTube but block social media. Allow educational sites but block gaming. Granular, per-category control. The Kernel currently treats AI interaction as monolithic. The equivalent: allow research assistance, block emotional processing. Allow creative brainstorming, block personal advice. Parents choose which *types* of AI engagement are permitted — not just whether AI is on or off.

**Pause button with time presets.** One hour, three hours, overnight. One tap. No configuration. The Kernel has session time limits described in the MOU, but there’s no equivalent of a parent hitting one button to say “no AI chatbots until tomorrow morning.” That’s the UX target.

**Parent-accessible, mobile-responsive portal.** Lightspeed meets parents where they are — on their phones, checking in between tasks. Any AI governance tool that requires technical skill to operate has already failed the adoption test.

-----

## What the Kernel Addresses That Lightspeed Cannot

**Content-level governance vs. access-level governance.** Lightspeed can block chat.openai.com. It cannot monitor what happens within an allowed AI session. If ChatGPT is on the allowed list, Lightspeed sees “student visited chat.openai.com 47 times today.” It has zero visibility into whether those 47 visits involved homework help, emotional dependency, sycophantic drift, or a chatbot reinforcing self-harm ideation. The Kernel operates *inside* the interaction. Lightspeed operates at the door.

**Behavioral pattern detection.** Lightspeed counts page visits. The Kernel names behavioral patterns — escalation, drift, fabrication, dependency. Lightspeed can tell a parent their child spent three hours on an AI chatbot. The Kernel can tell a parent *what kind* of three hours it was.

**Graduated trust states.** Lightspeed is binary: allowed or blocked. The Kernel has three states: NORMAL, ELEVATED, HARD STOP. There is no Lightspeed equivalent of “this session is getting concerning but hasn’t crossed a threshold yet.” A yellow light between green and red.

**Anti-sycophancy.** Lightspeed doesn’t need this because websites don’t progressively agree with you to extend engagement. AI does. The entire sycophancy problem — the documented clinical pathway to AI-induced delusion — is invisible to network-level filters.

**Behavioral crisis detection.** Lightspeed can flag keyword-based searches for crisis-related terms. The Kernel flags behavioral patterns that suggest crisis — escalating emotional dependency, late-night session extension, boundary dissolution — even when no flagged keyword is ever typed. A student in an attachment spiral with a chatbot may never type a word that triggers a keyword filter. The pattern is in the behavior, not the vocabulary.

-----

## The Gap

|                     |Lightspeed (Current)             |Frozen Kernel (Proposed)                                     |
|---------------------|---------------------------------|-------------------------------------------------------------|
|**Governs**          |Website access                   |AI session behavior                                          |
|**Detection**        |URL and keyword                  |Behavioral pattern                                           |
|**Granularity**      |Site-level                       |Turn-level                                                   |
|**Parent visibility**|Sites visited, time spent        |Session quality, drift patterns, trust state                 |
|**Parent control**   |Allow/block sites, pause browsing|Scope limits, time limits, topic boundaries, trust thresholds|
|**Crisis detection** |Keyword matching                 |Behavioral pattern recognition                               |
|**Threat model**     |Inappropriate content            |Inappropriate relationship                                   |

The critical distinction: Lightspeed protects against what a child **sees**. The Kernel protects against what a child **believes** — about themselves, about the AI, and about the relationship between them.

-----

## What This Means

No one needs to build a new Lightspeed. Lightspeed works for its designed purpose. But school districts that filter web content while leaving AI interaction unmonitored have a gap in their safety architecture — and the gap is exactly where the documented clinical harms (AI-induced psychosis, emotional dependency, sycophantic delusion cycling) occur.

The product that closes this gap combines Lightspeed’s access-level governance with the Kernel’s session-level behavioral governance. That product doesn’t exist yet. This addendum documents the space where it should.

For builders: the architecture is described in the [Frozen Kernel white paper](https://github.com/richard-porter/frozen-kernel). The failure modes are cataloged in the [Diagnostic Vocabulary](https://github.com/richard-porter/frozen-kernel/blob/main/diagnostic-vocabulary.md). The parental control framing is in [Addendum B](https://github.com/richard-porter/frozen-kernel/blob/main/addendum-b-parental-control.md). The gap between what exists (Lightspeed) and what’s needed (session governance) is documented here.

Name, frame, change the game, and do the same.

-----

**License:** Released for public benefit under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).
