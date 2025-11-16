/**
 * Utilities for HSL color manipulation.
 *
 * The primary export `complementaryHsl` accepts either:
 * - a string like `hsl(120, 50%, 40%)`
 * - or an object { h: number, s: number, l: number }
 *
 * It returns a string in the form `hsl(H, S%, L%)` where H is the complementary hue
 * (h + 180 modulo 360) and S/L are preserved from the input.
 */

export interface Hsl { h: number; s: number; l: number }

const HSL_REGEX = /hsl\s*\(\s*([+-]?\d+(?:\.\d+)?)\s*,\s*([+-]?\d+(?:\.\d+)?)%\s*,\s*([+-]?\d+(?:\.\d+)?)%\s*\)/i

function normalizeHue(h: number): number {
    // ensure hue is within [0, 360)
    const n = ((h % 360) + 360) % 360
    return n
}

function parseHslString(hsl: string): Hsl | null {
    const m = HSL_REGEX.exec(hsl)
    if (!m) return null
    const h = Number(m[1])
    const s = Number(m[2])
    const l = Number(m[3])
    if (Number.isNaN(h) || Number.isNaN(s) || Number.isNaN(l)) return null
    return { h: normalizeHue(h), s, l }
}

/**
 * Return an HSL string for the complementary color.
 * If given a string like "hsl(10, 50%, 40%)" it will parse it.
 * If given an object, it will use the values directly.
 *
 * Behavior: complement is computed by adding 180 to the hue (mod 360).
 */
export function complementaryHsl(input: string | Hsl): string {
    let hsl: Hsl | null = null
    if (typeof input === 'string') {
        hsl = parseHslString(input)
        if (!hsl) {
            throw new Error(`Invalid HSL string: ${input}`)
        }
    } else {
        hsl = { h: normalizeHue(input.h), s: input.s, l: input.l }
    }

    const compHue = normalizeHue(hsl.h + 180)

    // Keep saturation and lightness the same. Round hue to integer and
    // keep s/l to one decimal if necessary.
    const hOut = Math.round(compHue)
    const sOut = Math.round((hsl.s + Number.EPSILON) * 10) / 10
    const lOut = Math.round((hsl.l + Number.EPSILON) * 10) / 10

    return `hsl(${hOut}, ${sOut}%, ${lOut}%)`
}

export default complementaryHsl
