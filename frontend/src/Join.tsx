import { useEffect, useState } from 'react'
import { CreateGame, FetchRulesets } from "./axios.ts"
import './Join.css'
import './utils/color.ts'
import complementaryHsl from './utils/color.ts';
import Team from './models/Team.ts';
//TODO add enter key functionality to join stuff again?

function App() {

	const [code, setCode] = useState<string>("");
	const [ruleset, setRuleset] = useState("Random");
	const [rulesets, setRulesets] = useState<JSX.Element[]>([])
	const [rulesetDescriptions, setRulesetDescriptions] = useState<Record<string, string>>({})
	const [team, setTeam] = useState<Team>(Team.Black)

	// per-character inline styles for the header text
	const [charStyles, setCharStyles] = useState<React.CSSProperties[]>([])

	const HEADER_TEXT = "RANDOCHESS"

	function randomColor() {
		return `hsl(${Math.floor(Math.random() * 360)}, ${50 + Math.floor(Math.random() * 40)}%, ${30 + Math.floor(Math.random() * 40)}%)`
	}

	function generateStyle(): React.CSSProperties {
		const pick = Math.random()
		const color = randomColor()
		const textShadowColor = complementaryHsl(color)
		const rotate = (Math.random() - 0.5) * 20 // -10deg to 10deg // decreasing range to avoid char overlap
		const translateY = (Math.random() - 0.5) * 8 // -4px to 4px

		switch (true) {
			// heavy shadow
			case pick < 0.25:
				return {
					color,
					textShadow: `2px 2px 6px ${textShadowColor}`,
					transform: `rotate(${rotate}deg) translateY(${translateY}px)`,
				}
			// inverted / filled background
			/* Commenting out for now, may restrict the color selection later if this feature is determined to be usable without destroying readibility
			case pick < 0.45:
				return {
					color: '#fff',
					backgroundColor: color,
					padding: '0 2px',
					transform: `rotate(${rotate / 2}deg)`,
					filter: 'contrast(120%)',
				}
			*/
			// neon / glow
			case pick < 0.7:
				return {
					color,
					textShadow: `0 0 4px ${textShadowColor}, 0 0 12px ${textShadowColor}`,
					transform: `translateY(${translateY}px)`,
				}
			// subtle color change
			default:
				return {
					color,
					transform: `rotate(${rotate / 3}deg)`,
				}
		}
	}

	function changeTeam() {
		if (team == Team.White) {
			setTeam(Team.Black)

		} else {
			setTeam(Team.White)
		}
	}

	useEffect(() => {
		// generate a random style per character on mount
		const styles: React.CSSProperties[] = HEADER_TEXT.split("").map(() => generateStyle());
		if (styles)
			setCharStyles(styles);
	}, [])

	async function create_game() {
		let codes = await CreateGame(ruleset, team);
		console.log(codes);
		localStorage.setItem("guest_code", codes.guestCode);
		window.location.href = `/play?code=${codes.hostCode}`;
	}
	async function join_game() {
		window.location.href = `/play?code=${code}`;
	}

	async function get_rulesets(): Promise<JSX.Element[]> {
		// Fetch rulesets from backend (each ruleset has { name, description })
		const rulesResponse = await FetchRulesets()

		// Extract names and sort them
		const names = rulesResponse.rulesets.map(r => r.name).sort((a, b) => a.localeCompare(b))

		// build description map
		const descMap: Record<string, string> = {}
		rulesResponse.rulesets.forEach((r: any) => {
			descMap[r.name] = r.description || ''
		})
		// provide a default description for Random
		descMap['Random'] = descMap['Random'] || 'Creates a game using a randomly selected ruleset.'
		setRulesetDescriptions(descMap)

		// Put Random first
		names.unshift("Random")

		const html: JSX.Element[] = []
		names.forEach((element) => {
			// ensure unique key for each option
			html.push(
				<option key={element} value={element}>{element}</option>
			)
		})

		return html
	}



	useEffect(() => {
		async function init() {
			setRulesets(await get_rulesets())
		}
		init()
	}, [])

	return (
		<>
			<div className='siteHeader'>
				<button className='headerTextBox'>
					<h1>
						{HEADER_TEXT.split("").map((ch, i) => (
							<span key={i} className={`char char-${i}`} style={charStyles[i] || {}}>
								{ch}
							</span>
						))}
					</h1>
				</button>
			</div>
			<div className='joinContainer'>

				<div className='createGame'>

					<div className='selectGamerule'>
						<label htmlFor="ruleset" className="rulesetLabel">Select a Ruleset:</label>
						<br />
						<select name="ruleset" id="ruleset" onChange={(e) => setRuleset(e.target.value)}>
							{rulesets}
						</select>

						{/* dynamic description for selected ruleset */}
						<div className='rulesetDescription' aria-live="polite">
							{rulesetDescriptions[ruleset] || (ruleset === 'Random' ? 'Creates a game using a randomly selected ruleset.' : 'No description available.')}
						</div>

						<div className='buttonColumn'>
							<button onClick={create_game} className="createGameButton" type="button" aria-label="Create Game">
								<span>Create Game</span>
							</button>
							<button onClick={changeTeam} className={team == Team.White ? "changeTeamButtonWhite" : "changeTeamButtonBlack"}>
								<span>Toggle Team</span>
							</button>
						</div>
					</div>


				</div>


				<div className='joinGame'>
					<input placeholder="Or enter a join code" type="number" onChange={(e) => { setCode(e.target.value) }} className='joinGameTextbox' />
					<button onClick={join_game} className='joinGameButton'>
						Join Game
					</button>

				</div>
			</div>
		</>
	)
}

export default App