# OriginLab — Design Guide

> Vibe: hand-inked, calm, earthy. A single bold black brush line on warm paper. Craft coffee, not corporate tech.

Source of truth: the OriginLab logo (cup + two leaves + saucer, black ink on warm gray).

---

## 1. Brand Feel

- **Hand-made, not polished.** Lines are thick, slightly wobbly, imperfect on purpose.
- **Quiet and minimal.** One color of ink, lots of breathing room.
- **Natural / origin story.** Leaf = plant, farm, source. Cup = the ritual.
- **Warm, not cold.** Never pure white or pure black.

Keywords: _organic, ink, paper, calm, roasted, honest, small-batch._

---

## 2. Color

Hex values are eyeballed from the logo. Tweak after sampling the real file.

| Token          | Hex       | Use                              |
| -------------- | --------- | -------------------------------- |
| `--paper`      | `#E0DFDA` | Page background (logo bg)        |
| `--paper-soft` | `#ECEBE6` | Cards, raised surfaces           |
| `--paper-deep` | `#D3D2CC` | Borders, dividers, hover bg      |
| `--ink`        | `#1C1C1B` | Text, icons, primary buttons     |
| `--ink-soft`   | `#4A4947` | Secondary text                   |
| `--ink-faint`  | `#8A8985` | Captions, placeholders, disabled |

Optional accent (new, not from the logo, pick or drop it):

| Token     | Hex       | Use                                         |
| --------- | --------- | ------------------------------------------- |
| `--leaf`  | `#4F5D3A` | Success states, "in stock", tiny highlights |
| `--roast` | `#8B5A3C` | Sale/limited badges, rare accent            |

Rules:

- Default is **ink on paper**. Accents are rare, never full-page.
- Never use `#FFFFFF` or `#000000`.
- Dark mode (if added): swap to `--paper: #1C1C1B`, `--ink: #E0DFDA`.

```css
:root {
	--paper: #e0dfda;
	--paper-soft: #ecebe6;
	--paper-deep: #d3d2cc;
	--ink: #1c1c1b;
	--ink-soft: #4a4947;
	--ink-faint: #8a8985;
	--leaf: #4f5d3a;
	--roast: #8b5a3c;
}
```

---

## 3. Typography

Fonts below are suggestions (free on Google Fonts). Pick your own if you prefer.

| Role            | Suggestion                       | Why                                   |
| --------------- | -------------------------------- | ------------------------------------- |
| Headings        | **Fraunces** (soft, wonky serif) | Warm, slightly hand-made feel         |
| Body            | **Inter** or **DM Sans**         | Clean, readable                       |
| Accent / labels | **Caveat** or none               | Use sparingly for handwritten touches |

Scale:

- **Display:** 48-64px, tight line-height (1.05), weight 600
- **H1:** 36-40px
- **H2:** 28px
- **H3:** 20px
- **Body:** 16-17px, line-height 1.6
- **Caption:** 13px, `--ink-faint`

Rules:

- Sentence case for headings. No ALL CAPS except tiny labels (letter-spacing 0.08em).
- Max text width ~65ch.

---

## 4. Shape and Line

The logo is all about **thick rounded strokes**. Carry that into the UI.

- **Borders:** `2px solid var(--ink)` on key elements (buttons, inputs, product cards)
- **Radius:** `12px` default, `999px` for pills/tags. No sharp corners.
- **Shadows:** avoid blurry drop shadows. If needed, use a flat offset: `4px 4px 0 var(--ink)`
- **Dividers:** `1.5px` lines in `--paper-deep`, or a hand-drawn squiggle SVG for section breaks
- **Icons:** thick outline style (stroke 2-2.5px), rounded caps and joins, ink color only

---

## 5. Components

**Buttons**

- Primary: `--ink` bg, `--paper` text, 2px ink border, radius 12px
- Secondary: transparent bg, 2px ink border, ink text
- Hover: shift up 2px with flat offset shadow (feels like a stamped card)
- Min height 44px

**Cards (product)**

- `--paper-soft` bg, 2px ink border, radius 16px
- Product photo on top, warm and natural light
- Name in heading font, price in body font, bold

**Inputs**

- `--paper-soft` bg, 2px ink border, radius 12px
- Focus: ink border plus a 3px `--paper-deep` outer ring

**Badges / tags**

- Pill shape, 1.5px ink border, small text
- Use `--leaf` or `--roast` fill only for status

**Nav**

- Paper bg, logo left, minimal text links, cart icon right
- No heavy bars. A single thin ink line below is enough.

---

## 6. Layout and Spacing

- 8px base grid. Common steps: 8, 16, 24, 32, 48, 80, 120
- Generous whitespace. When unsure, add more.
- Max content width: 1120px
- Hero: centered, big heading, short line, one button, lots of empty paper
- Mobile first. Single column under 768px.

---

## 7. Imagery

- **Photos:** natural light, warm tones, raw textures (beans, burlap, ceramic, steam). Slightly desaturated.
- **No** stock-photo smiles, neon, or glossy 3D renders.
- **Illustration:** same hand-inked black line style as the logo. Leaves, beans, cups, hills.
- Crop with rounded corners (16px) or an organic blob mask.

---

## 8. Motion

- Gentle and slow: 150-250ms, `ease-out`
- Hover: small lift (2px) or line "redraw" effect on icons
- No bouncy or flashy animation. Respect `prefers-reduced-motion`.

---

## 9. Voice and Copy

- Friendly, short, honest. Like a barista who knows the farmer.
- Talk about origin, roast level, tasting notes.
- Example: _"Roasted small, shipped fresh."_ not _"Premium artisanal coffee solutions."_

---

## 10. Logo Usage

- Ink logo on `--paper` or `--paper-soft` only
- On dark bg, use the paper-colored version
- Keep clear space equal to the cup handle width on all sides
- Don't recolor, outline, add shadows, or stretch

---

## 11. Do / Don't

| Do                               | Don't                         |
| -------------------------------- | ----------------------------- |
| Thick ink lines, rounded corners | Thin hairlines, sharp corners |
| Warm off-white paper             | Pure white backgrounds        |
| One ink color, rare accent       | Rainbow palettes, gradients   |
| Lots of whitespace               | Dense, cluttered layouts      |
| Hand-drawn touches               | Glossy or corporate-tech look |
